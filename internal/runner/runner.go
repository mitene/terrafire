package runner

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/mitene/terrafire/internal/api"
	"github.com/mitene/terrafire/internal/manifest"
	"github.com/mitene/terrafire/internal/utils"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type Runner struct {
	projects map[string]*api.Project
	client   api.SchedulerClient
	git      utils.Git
	tf       Terraform
	blob     Blob
}

type runnerContext struct {
	Destroy bool
}

func NewRunner(projects map[string]*api.Project, client api.SchedulerClient, git utils.Git, tf Terraform, blob Blob) *Runner {
	return &Runner{
		projects: projects,
		client:   client,
		git:      git,
		tf:       tf,
		blob:     blob,
	}
}

/*
Plan
*/

func (r *Runner) Plan(project string, workspace string) error {
	_, err := r.client.UpdateJobStatus(context.Background(), &api.UpdateJobStatusRequest{
		Project:   project,
		Workspace: workspace,
		Status:    api.Job_PlanInProgress,
	})
	if err != nil {
		return err
	}

	output, closer := r.makeJobLogger(project, workspace, api.Phase_Plan)
	defer closer()

	result, err := r.doPlan(project, workspace, output)
	if err != nil {
		_, err1 := r.client.UpdateJobStatus(context.Background(), &api.UpdateJobStatusRequest{
			Project:   project,
			Workspace: workspace,
			Status:    api.Job_PlanFailed,
			Error:     err.Error(),
		})
		utils.LogError(err1)
		return err
	}

	_, err = r.client.UpdateJobStatus(context.Background(), &api.UpdateJobStatusRequest{
		Project:          project,
		Workspace:        workspace,
		Status:           api.Job_ReviewRequired,
		Result:           result.result,
		ProjectVersion:   result.projectVersion,
		WorkspaceVersion: result.workspaceVersion,
		Destroy:          result.destroy,
	})
	return err
}

type planResult struct {
	result           string
	projectVersion   string
	workspaceVersion string
	destroy          bool
}

func (r *Runner) doPlan(project string, workspace string, output io.Writer) (*planResult, error) {
	log.WithFields(logrus.Fields{
		"project":   project,
		"workspace": workspace,
	}).Info("start plan")

	pj, ok := r.projects[project]
	if !ok {
		return nil, fmt.Errorf("project is not defined: %s", project)
	}

	ws, pjVersion, err := r.fetchWorkspace(pj.Repo, pj.Branch, pj.Path, workspace)
	if err != nil {
		return nil, err
	}

	var wsVersion string
	var destroy bool
	if ws != nil {
		wsVersion = ws.Source.Ref
		destroy = false
	} else {
		resp, err := r.client.GetWorkspaceVersion(context.Background(), &api.GetWorkspaceVersionRequest{
			Project:   project,
			Workspace: workspace,
		})
		if err != nil {
			return nil, err
		}

		ws, pjVersion, err = r.fetchWorkspace(pj.Repo, resp.ProjectVersion, pj.Path, workspace)
		if err != nil {
			return nil, err
		}
		if ws == nil {
			return nil, fmt.Errorf("workspace %s/%s not found", project, workspace)
		}

		wsVersion = resp.WorkspaceVersion
		destroy = true
	}

	// fetch terraform module
	dir, err := utils.TempDir()
	if err != nil {
		return nil, err
	}
	defer utils.TempClean(dir)

	var repo string
	switch ws.Source.Type {
	case api.Source_github:
		repo = fmt.Sprintf("https://github.com/%s/%s", ws.Source.Owner, ws.Source.Repo)
	default:
		return nil, fmt.Errorf("invalid source type: %d", ws.Source.Type)
	}

	log.WithFields(logrus.Fields{
		"repo": repo,
		"ref":  ws.Source.Ref,
		"path": ws.Source.Path,
	}).Info("fetch terraform module")
	wsVersion, err = r.fetch(dir, repo, wsVersion, ws.Source.Path)
	if err != nil {
		return nil, err
	}

	// run terraform
	opts := TerraformOption{
		dir:  dir,
		out:  output,
		envs: make([]string, len(pj.Envs)),
	}
	for i, kv := range pj.Envs {
		opts.envs[i] = kv.GetKey() + "=" + kv.GetValue()
	}

	varfiles := make([]string, len(ws.VarFiles))
	for i, v := range ws.VarFiles {
		err := r.makeVarFile(dir, v.Key, v.Value)
		if err != nil {
			return nil, err
		}
		varfiles[i] = v.Key
	}

	vars := make([]string, len(ws.Vars))
	for i, v := range ws.Vars {
		vars[i] = v.Key + "=" + v.Value
	}

	out, err := r.tf.Plan(opts, ws.Workspace, vars, varfiles, destroy)
	if err != nil {
		return nil, err
	}

	err = r.saveContext(dir, &runnerContext{Destroy: destroy})
	if err != nil {
		return nil, err
	}

	err = r.saveArtifacts(project, workspace, dir)
	if err != nil {
		return nil, err
	}

	return &planResult{
		result:           string(out),
		projectVersion:   pjVersion,
		workspaceVersion: wsVersion,
		destroy:          destroy,
	}, nil
}

func (r *Runner) fetchWorkspace(repo, ref, path, workspace string) (*api.Workspace, string, error) {
	// fetch terrafire manifest
	dir, err := utils.TempDir()
	if err != nil {
		return nil, "", err
	}
	defer utils.TempClean(dir)

	log.WithFields(logrus.Fields{
		"repo": repo,
		"ref":  ref,
		"path": path,
	}).Info("fetch manifest")

	commit, err := r.fetch(dir, repo, ref, path)
	if err != nil {
		return nil, "", err
	}

	man, err := manifest.Load(dir)
	if err != nil {
		return nil, "", err
	}

	for _, ws := range man {
		if ws.Name == workspace {
			return ws, commit, nil
		}
	}
	return nil, commit, nil
}

func (*Runner) makeVarFile(dir string, filename string, body string) error {
	fp := filepath.Join(dir, filename)

	err := os.MkdirAll(filepath.Dir(fp), 0755)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fp, []byte(body), 0400)
	if err != nil {
		return err
	}

	return nil
}

/*
Apply
*/

func (r *Runner) Apply(project string, workspace string) error {
	_, err := r.client.UpdateJobStatus(context.Background(), &api.UpdateJobStatusRequest{
		Project:   project,
		Workspace: workspace,
		Status:    api.Job_ApplyInProgress,
	})
	if err != nil {
		return err
	}

	output, closer := r.makeJobLogger(project, workspace, api.Phase_Apply)
	defer closer()

	err = r.doApply(project, workspace, output)
	if err != nil {
		_, err1 := r.client.UpdateJobStatus(context.Background(), &api.UpdateJobStatusRequest{
			Project:   project,
			Workspace: workspace,
			Status:    api.Job_ApplyFailed,
			Error:     err.Error(),
		})
		utils.LogError(err1)
		return err
	}

	_, err = r.client.UpdateJobStatus(context.Background(), &api.UpdateJobStatusRequest{
		Project:   project,
		Workspace: workspace,
		Status:    api.Job_Succeeded,
	})
	return err
}

func (r *Runner) doApply(project string, workspace string, output io.Writer) (err error) {
	log.WithFields(logrus.Fields{
		"project":   project,
		"workspace": workspace,
	}).Info("start apply")

	pj, ok := r.projects[project]
	if !ok {
		return fmt.Errorf("project is not defined: %s", project)
	}

	dir, err := utils.TempDir()
	if err != nil {
		return err
	}
	defer utils.TempClean(dir)

	err = r.loadArtifacts(project, workspace, dir)
	if err != nil {
		return
	}

	ctx, err := r.loadContext(dir)
	if err != nil {
		return err
	}

	opts := TerraformOption{
		dir:  dir,
		out:  output,
		envs: make([]string, 0, len(pj.Envs)),
	}
	for _, kv := range pj.Envs {
		opts.envs = append(opts.envs, kv.GetKey()+"="+kv.GetValue())
	}

	return r.tf.Apply(opts, ctx.Destroy)
}

/*
Helper Functions
*/

func (r *Runner) fetch(dest string, repo string, ref string, path string) (string, error) {
	dir, err := utils.TempDir()
	if err != nil {
		return "", err
	}
	defer utils.TempClean(dir)

	commit, err := r.git.Fetch(dir, repo, ref)
	if err != nil {
		return "", err
	}

	err = os.RemoveAll(filepath.Join(dir, ".git"))
	if err != nil {
		return "", err
	}

	path = filepath.Clean(path)
	if strings.HasPrefix(path, "..") {
		return "", fmt.Errorf("invalid source path: %s", path)
	}

	fs, err := ioutil.ReadDir(filepath.Join(dir, path))
	if err != nil {
		return "", err
	}

	for _, f := range fs {
		err = os.Rename(filepath.Join(dir, path, f.Name()), filepath.Join(dest, f.Name()))
		if err != nil {
			return "", err
		}
	}

	return commit, nil
}

// runner context

func (*Runner) saveContext(dir string, ctx *runnerContext) error {
	fp := filepath.Join(dir, ".terrafire")
	f, err := os.Create(fp)
	if err != nil {
		return err
	}
	defer utils.LogDefer(f.Close)

	return json.NewEncoder(f).Encode(ctx)
}

func (*Runner) loadContext(dir string) (*runnerContext, error) {
	fp := filepath.Join(dir, ".terrafire")
	f, err := os.Open(fp)
	if err != nil {
		return nil, err
	}
	defer utils.LogDefer(f.Close)

	ctx := &runnerContext{}
	err = json.NewDecoder(f).Decode(ctx)
	if err != nil {
		return nil, err
	}
	return ctx, nil
}

// artifacts

func (r *Runner) saveArtifacts(project string, workspace string, dir string) error {
	buf, err := utils.TempFile()
	if err != nil {
		return err
	}
	defer utils.TempClean(buf.Name())

	err = Zip(buf, dir)
	if err != nil {
		return err
	}

	err = buf.Sync()
	if err != nil {
		return err
	}

	_, err = buf.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	return r.blob.Put(project, workspace, buf)
}

func (r *Runner) loadArtifacts(project string, workspace string, dir string) error {
	src, err := r.blob.Get(project, workspace)
	if err != nil {
		return err
	}
	defer utils.LogDefer(src.Close)

	err = Unzip(src, dir)
	if err != nil {
		return err
	}

	return nil
}

// log

func (r *Runner) makeJobLogger(project string, workspace string, phase api.Phase) (*bytes.Buffer, func()) {
	output := bytes.NewBuffer(nil)
	mux := sync.Mutex{}

	sendLog := func() {
		mux.Lock()
		defer mux.Unlock()

		_, err1 := r.client.UpdateJobLog(context.Background(), &api.UpdateJobLogRequest{
			Project:   project,
			Workspace: workspace,
			Phase:     phase,
			Log:       output.String(),
		})
		utils.LogError(err1)
	}

	closer := r.withInterval(5*time.Second, sendLog)

	stop := func() {
		closer()
		sendLog() // assume that log is completed sent when process finishes
	}

	return output, stop
}

// misc

func (*Runner) withInterval(d time.Duration, f func()) func() {
	finish := false
	go func() {
		for !finish {
			f()
			time.Sleep(d)
		}
		f()
	}()
	return func() {
		finish = true
	}
}
