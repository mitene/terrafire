package utils

import (
	"fmt"
	"github.com/mitene/terrafire/core"
	"io"
	"os"
	"os/exec"
)

type Terraform struct{}

func NewTerraform() *Terraform {
	return &Terraform{}
}

func (t *Terraform) Plan(dir string, workspace *core.Workspace, output io.Writer) (result string, err error) {
	envs := workspace.Project.Envs

	src, err := t.formatModuleAddr(workspace)
	if err != nil {
		return "", err
	}

	// download source code and initialize
	err = t.run(dir, output, envs, "init", "-from-module="+src, "-input=false", "-no-color")
	if err != nil {
		return "", err
	}

	// select workspace
	if workspace.Workspace != "" {
		err = t.run(dir, output, envs, "workspace", "select", "-no-color", workspace.Workspace)
		if err != nil {
			err = t.run(dir, output, envs, "workspace", "new", "-no-color", workspace.Workspace)
			if err != nil {
				return "", err
			}
		}
	}

	// plan
	args := []string{"plan", "-no-color", "-out=tfplan", "-input=false"}
	for _, vf := range workspace.VarFiles {
		args = append(args, "-var-file="+vf)
	}
	for k, v := range workspace.Vars {
		args = append(args, fmt.Sprintf("-var=%s=%s", k, v))
	}
	err = t.run(dir, output, envs, args...)
	if err != nil {
		return "", err
	}

	out, err := t.output(dir, "show", "-no-color", "tfplan")
	if err != nil {
		return "", err
	}

	return string(out), nil
}

func (t *Terraform) Apply(dir string, workspace *core.Workspace, output io.Writer) (err error) {
	envs := workspace.Project.Envs
	return t.run(dir, output, envs, "apply", "-no-color", "-input=false", "tfplan")
}

func (t *Terraform) formatModuleAddr(workspace *core.Workspace) (string, error) {
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

func (t *Terraform) run(dir string, output io.Writer, envs map[string]string, arg ...string) error {
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

func (t *Terraform) output(dir string, arg ...string) ([]byte, error) {
	cmd := exec.Command("terraform", arg...)
	cmd.Dir = dir
	return cmd.Output()
}
