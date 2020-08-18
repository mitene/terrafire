package runner

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

type Terraform interface {
	Plan(option TerraformOption, workspace string, vars []string, varfiles []string, destroy bool) ([]byte, error)
	Apply(option TerraformOption, destroy bool) error
}

type TerraformOption struct {
	dir  string
	out  io.Writer
	envs []string
}

type TerraformImpl struct {
}

func NewTerraform() Terraform {
	return &TerraformImpl{}
}

func (t *TerraformImpl) Plan(option TerraformOption, workspace string, vars []string, varfiles []string, destroy bool) ([]byte, error) {
	_, _ = fmt.Fprintln(option.out, "---- init --------------------------------------------------------------")

	err := t.newCmd(option, "init", "-input=false", "-no-color").Run()
	if err != nil {
		return nil, fmt.Errorf("terraform init failed: %w", err)
	}

	if workspace != "" {
		_, _ = fmt.Fprintln(option.out, "\n---- workspace select/new ----------------------------------------------")
		err = t.newCmd(option, "workspace", "select", "-no-color", workspace).Run()
		if err != nil {
			err = t.newCmd(option, "workspace", "new", "-no-color", workspace).Run()
			if err != nil {
				return nil, fmt.Errorf("terraform workspace new failed: %w", err)
			}
		}
	}

	_, _ = fmt.Fprintln(option.out, "\n---- plan --------------------------------------------------------------")
	args := []string{"plan", "-no-color", "-input=false", "-out=tfplan"}
	for _, vf := range varfiles {
		args = append(args, "-var-file="+vf)
	}
	for _, v := range vars {
		args = append(args, "-var="+v)
	}
	if destroy {
		args = append(args, "-destroy")
	}
	err = t.newCmd(option, args...).Run()
	if err != nil {
		return nil, fmt.Errorf("terraform plan failed: %w", err)
	}

	_, _ = fmt.Fprintln(option.out, "\n---- show --------------------------------------------------------------")
	cmd := t.newCmd(option, "show", "-no-color", "tfplan")
	cmd.Stdout = nil
	result, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("terraform show failed: %w", err)
	}

	return result, nil
}

func (t *TerraformImpl) Apply(option TerraformOption, destroy bool) error {
	_, _ = fmt.Fprintln(option.out, "\n---- apply -------------------------------------------------------------")
	err := t.newCmd(option, "apply", "-no-color", "-input=false", "tfplan").Run()
	if !destroy && err != nil {
		return fmt.Errorf("terraform apply failed: %w", err)
	}

	if destroy {
		_, _ = fmt.Fprintln(option.out, "\n---- check resources ---------------------------------------------------")
		cmd := t.newCmd(option, "state", "list")
		cmd.Stdout = nil
		result, err := cmd.Output()
		if err != nil {
			return err
		}
		if len(bytes.TrimSpace(result)) != 0 {
			_, _ = fmt.Fprintf(option.out, "resources remain left in workspace: %s\n", result)
			return fmt.Errorf("resources remain left in workspace: %s", result)
		}

		cmd = t.newCmd(option, "workspace", "show")
		cmd.Stdout = nil
		result, err = cmd.Output()
		if err != nil {
			return err
		}
		current := string(bytes.TrimSpace(result))

		if current == "default" {
			// TODO: cannot delete default workspace
			_, _ = fmt.Println(option.out, "Cannot delete default workspace.")
			_, _ = fmt.Println(option.out, "Remote terraform state may remain.")
		} else {
			err = t.newCmd(option, "workspace", "select", "default").Run()
			if err != nil {
				return err
			}

			err = t.newCmd(option, "workspace", "delete", "-force", current).Run()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (t *TerraformImpl) newCmd(option TerraformOption, arg ...string) *exec.Cmd {
	cmd := exec.Command("terraform", arg...)
	cmd.Dir = option.dir
	cmd.Env = append(os.Environ(), "TF_IN_AUTOMATION=true")
	cmd.Env = append(cmd.Env, option.envs...)
	cmd.Stdout = option.out
	cmd.Stderr = option.out
	return cmd
}
