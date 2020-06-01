package controller

import (
	"os"
	"os/exec"
)

type ExecutorLocal struct {
	path string
}

func NewLocalExecutor(path string) Executor {
	return &ExecutorLocal{
		path: path,
	}
}

func (r *ExecutorLocal) Plan(project string, workspace string) (Process, error) {
	return r.run("plan", project, workspace)
}

func (r *ExecutorLocal) Apply(project string, workspace string) (Process, error) {
	return r.run("apply", project, workspace)
}

func (r *ExecutorLocal) run(phase string, project string, workspace string) (*localProcess, error) {
	cmd := exec.Command(r.path, "run", phase, project, workspace)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		return nil, err
	}

	return &localProcess{cmd: cmd}, nil
}

type localProcess struct {
	cmd *exec.Cmd
}

func (p *localProcess) wait() error {
	return p.cmd.Wait()
}

func (p *localProcess) cancel() error {
	return p.cmd.Process.Signal(os.Interrupt)
}
