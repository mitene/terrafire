package terrafire

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

type TerraformClient interface {
	Plan(dir string, params *ConfigTerraformDeployParams) (string, error)
	Apply(dir string, params *ConfigTerraformDeployParams, autoApprove bool) error
}

type TerraformClientImpl struct {
}

func NewTerraformClient() TerraformClient {
	return &TerraformClientImpl{}
}

func (t *TerraformClientImpl) Plan(dir string, params *ConfigTerraformDeployParams) (string, error) {
	err := t.init(dir, params)
	if err != nil {
		return "", err
	}

	planResult, err := ioutil.TempFile("", "")
	if err != nil {
		return "", err
	}
	defer os.Remove(planResult.Name())

	err = t.run(dir, "plan", append(t.makeArgs(params), "-no-color", "--out="+planResult.Name())...)
	if err != nil {
		return "", err
	}

	cmd := exec.Command("terraform", "show", "-no-color", planResult.Name())
	cmd.Dir = dir
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(out), nil
}

func (t *TerraformClientImpl) Apply(dir string, params *ConfigTerraformDeployParams, autoApprove bool) error {
	err := t.init(dir, params)
	if err != nil {
		return err
	}

	args := t.makeArgs(params)
	args = append(args, "-no-color")
	if autoApprove {
		args = append(args, "-auto-approve")
	}
	return t.run(dir, "apply", args...)
}

func (t *TerraformClientImpl) init(dir string, params *ConfigTerraformDeployParams) error {
	err := t.run(dir, "init", "-no-color")
	if err != nil {
		return err
	}
	if params == nil || params.Workspace == "" {
		return nil
	}
	err = t.run(dir, "workspace", "select", "-no-color", params.Workspace)
	if err != nil {
		err = t.run(dir, "workspace", "new", "-no-color", params.Workspace)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *TerraformClientImpl) run(dir string, command string, arg ...string) error {
	args := append([]string{command}, arg...)
	cmd := exec.Command("terraform", args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (*TerraformClientImpl) makeArgs(params *ConfigTerraformDeployParams) []string {
	var args []string

	if params != nil && params.VarFiles != nil {
		for _, vf := range *params.VarFiles {
			args = append(args, "-var-file="+vf)
		}
	}

	if params != nil && params.Vars != nil {
		for k, v := range *params.Vars {
			args = append(args, fmt.Sprintf("-var=%s=%s", k, v))
		}
	}

	return args
}
