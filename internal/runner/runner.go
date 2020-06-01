package runner

import (
	"bytes"
	"context"
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
		Project:   project,
		Workspace: workspace,
		Status:    api.Job_ReviewRequired,
		Result:    result,
	})
	return err
}

func (r *Runner) doPlan(project string, workspace string, output io.Writer) (result string, err error) {
	log.WithFields(logrus.Fields{
		"project":   project,
		"workspace": workspace,
	}).Info("start plan")

	pj, ok := r.projects[project]
	if !ok {
		return "", fmt.Errorf("project is not defined: %s", project)
	}

	// fetch terrafire manifest
	dirM, err := utils.TempDir()
	if err != nil {
		return "", err
	}
	defer utils.TempClean(dirM)

	man, err := r.fetchManifest(dirM, pj)
	if err != nil {
		return "", err
	}

	ws, err := r.findWorkspace(man, workspace)
	if err != nil {
		return "", err
	}

	// fetch terraform config
	dirT, err := utils.TempDir()
	if err != nil {
		return "", err
	}
	defer utils.TempClean(dirT)

	_, err = r.fetchModule(dirT, ws)
	if err != nil {
		return "", err
	}

	// run terraform
	opts := TerraformOption{
		dir:  dirT,
		out:  output,
		envs: make([]string, 0, len(pj.Envs)),
	}
	for _, kv := range pj.Envs {
		opts.envs = append(opts.envs, kv.GetKey()+"="+kv.GetValue())
	}

	varfiles := make([]string, 0, len(ws.VarFiles))
	for _, vf := range ws.VarFiles {
		varfiles = append(varfiles, filepath.Join(dirM, vf))
	}

	vars := make([]string, 0, len(ws.Vars))
	for _, v := range ws.Vars {
		vars = append(vars, v.Key+"="+v.Value)
	}

	out, err := r.tf.Plan(opts, ws.Workspace, vars, varfiles)
	if err != nil {
		return "", err
	}

	err = r.saveArtifacts(project, workspace, dirT)
	if err != nil {
		return "", err
	}

	return string(out), nil
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

	opts := TerraformOption{
		dir:  dir,
		out:  output,
		envs: make([]string, 0, len(pj.Envs)),
	}
	for _, kv := range pj.Envs {
		opts.envs = append(opts.envs, kv.GetKey()+"="+kv.GetValue())
	}

	return r.tf.Apply(opts)
}

/*
Helper Functions
*/

// download sources

func (r *Runner) fetchManifest(dest string, pj *api.Project) (*api.Manifest, error) {
	log.WithFields(logrus.Fields{
		"repo":   pj.Repo,
		"branch": pj.Branch,
		"path":   pj.Path,
	}).Info("fetch manifest")

	_, err := r.fetch(dest, pj.Repo, pj.Branch, pj.Path)
	if err != nil {
		return nil, err
	}

	man, err := manifest.Load(dest)
	if err != nil {
		return nil, err
	}

	return man, nil
}

func (r *Runner) fetchModule(dest string, ws *api.Workspace) (string, error) {
	src := ws.Source

	var repo string
	switch src.Type {
	case api.Source_github:
		repo = fmt.Sprintf("https://github.com/%s/%s", src.Owner, src.Repo)
	default:
		return "", fmt.Errorf("invalid source type: %d", src.Type)
	}

	log.WithFields(logrus.Fields{
		"repo": repo,
		"ref":  src.Ref,
		"path": src.Path,
	}).Info("fetch terraform module")

	return r.fetch(dest, repo, src.Ref, src.Path)
}

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

func (r *Runner) findWorkspace(manifest *api.Manifest, workspace string) (*api.Workspace, error) {
	for _, ws := range manifest.Workspaces {
		if ws.Name == workspace {
			return ws, nil
		}
	}
	return nil, fmt.Errorf("workspace is not defined: %s", workspace)
}

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
