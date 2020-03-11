package terrafire

import (
	"fmt"
	"os"
	"os/exec"
)

type TerraformClient interface {
	Plan(dir string, params *ConfigTerraformDeployParams) error
	Apply(dir string, params *ConfigTerraformDeployParams) error
}

type TerraformClientImpl struct {
}

func NewTerraformClient() TerraformClient {
	return &TerraformClientImpl{}
}

func (t *TerraformClientImpl) Plan(dir string, params *ConfigTerraformDeployParams) error {
	err := t.init(dir, params)
	if err != nil {
		return err
	}

	return t.run(dir, "plan", t.makeArgs(params)...)
}

func (t *TerraformClientImpl) Apply(dir string, params *ConfigTerraformDeployParams) error {
	err := t.init(dir, params)
	if err != nil {
		return err
	}

	return t.run(dir, "apply", t.makeArgs(params)...)
}

func (t *TerraformClientImpl) init(dir string, params *ConfigTerraformDeployParams) error {
	err := t.run(dir, "init")
	if err != nil {
		return err
	}
	if params == nil || params.Workspace == "" {
		return nil
	}
	err = t.run(dir, "workspace", "select", params.Workspace)
	if err != nil {
		err = t.run(dir, "workspace", "new", params.Workspace)
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
