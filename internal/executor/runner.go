package executor

import (
	"bytes"
	"fmt"
	"github.com/mitene/terrafire/internal"
	"github.com/mitene/terrafire/internal/utils"
	"io"
	"os"
	"os/exec"
	"time"
)

type Runner struct {
	handler internal.Handler
	blob    internal.Blob
}

func NewRunner(handler internal.Handler, blob internal.Blob) *Runner {
	return &Runner{
		handler: handler,
		blob:    blob,
	}
}

func (r *Runner) Plan(payload *internal.ExecutorPayload) {
	project, workspace := payload.Project.Name, payload.Workspace.Name

	utils.LogError(r.handler.UpdateJobStatusPlanInProgress(project, workspace))

	output := bytes.NewBuffer(nil)
	closer := r.withInterval(5*time.Second, func() {
		utils.LogError(r.handler.SavePlanLog(project, workspace, output.String()))
	})
	defer closer()

	result, err := r.doPlan(payload, output)
	if err != nil {
		utils.LogError(r.handler.UpdateJobStatusPlanFailed(project, workspace, err))
		return
	}

	utils.LogError(r.handler.UpdateJobStatusReviewRequired(project, workspace, result))
}

func (r *Runner) doPlan(payload *internal.ExecutorPayload, output io.Writer) (result string, err error) {
	pj, ws := payload.Project, payload.Workspace
	project, workspace := pj.Name, ws.Name
	envs := pj.Envs

	dir, err := r.blob.New(project, workspace)
	if err != nil {
		return "", err
	}

	src, err := r.formatModuleAddr(ws)
	if err != nil {
		return "", err
	}

	// download source code and initialize
	err = r.run(dir, output, envs, "init", "-from-module="+src, "-input=false", "-no-color")
	if err != nil {
		return "", err
	}

	// select workspace
	if ws.Workspace != "" {
		err = r.run(dir, output, envs, "workspace", "select", "-no-color", ws.Workspace)
		if err != nil {
			err = r.run(dir, output, envs, "workspace", "new", "-no-color", ws.Workspace)
			if err != nil {
				return "", err
			}
		}
	}

	// plan
	args := []string{"plan", "-no-color", "-out=tfplan", "-input=false"}
	for _, vf := range ws.VarFiles {
		args = append(args, "-var-file="+vf)
	}
	for k, v := range ws.Vars {
		args = append(args, fmt.Sprintf("-var=%s=%s", k, v))
	}
	err = r.run(dir, output, envs, args...)
	if err != nil {
		return "", err
	}

	out, err := r.output(dir, output, envs, "show", "-no-color", "tfplan")
	if err != nil {
		return "", err
	}

	err = r.blob.Put(project, workspace)
	if err != nil {
		return "", err
	}

	return string(out), nil
}

func (r *Runner) Apply(payload *internal.ExecutorPayload) {
	project, workspace := payload.Project.Name, payload.Workspace.Name

	utils.LogError(r.handler.UpdateJobStatusApplyInProgress(project, workspace))

	output := bytes.NewBuffer(nil)
	closer := r.withInterval(5*time.Second, func() {
		utils.LogError(r.handler.SaveApplyLog(project, workspace, output.String()))
	})
	defer closer()

	err := r.doApply(payload, output)
	if err != nil {
		utils.LogError(r.handler.UpdateJobStatusApplyFailed(project, workspace, err))
		return
	}

	utils.LogError(r.handler.UpdateJobStatusSucceeded(project, workspace))
}

func (r *Runner) doApply(payload *internal.ExecutorPayload, output io.Writer) (err error) {
	pj, ws := payload.Project, payload.Workspace
	project, workspace := pj.Name, ws.Name
	envs := pj.Envs

	dir, err := r.blob.Get(project, workspace)
	if err != nil {
		return
	}

	return r.run(dir, output, envs, "apply", "-no-color", "-input=false", "tfplan")
}

func (r *Runner) formatModuleAddr(workspace *internal.Workspace) (string, error) {
	switch src := workspace.Source; src.Type {
	case "github":
		if src.Owner == "" || src.Repo == "" {
			return "", fmt.Errorf("owner and repo must be specified")
		}

		addr := fmt.Sprintf("git::https://github.com/%s/%s", src.Owner, src.Repo)
		if src.Path != "" {
			addr += "//" + src.Path
		}
		if src.Ref != "" {
			addr += "?ref=" + src.Ref
		}

		return addr, nil
	default:
		return "", fmt.Errorf("invalid source type: %s", src.Type)
	}
}

func (r *Runner) run(dir string, output io.Writer, envs map[string]string, arg ...string) error {
	cmd := exec.Command("terraform", arg...)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), "TF_IN_AUTOMATION=true")
	for k, v := range envs {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", k, v))
	}

	if output == nil {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	} else {
		cmd.Stdout = output
		cmd.Stderr = output
	}

	return cmd.Run()
}

func (r *Runner) output(dir string, output io.Writer, envs map[string]string, arg ...string) ([]byte, error) {
	cmd := exec.Command("terraform", arg...)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), "TF_IN_AUTOMATION=true")
	for k, v := range envs {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", k, v))
	}
	if output == nil {
		cmd.Stderr = os.Stderr
	} else {
		cmd.Stderr = output
	}
	return cmd.Output()
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
