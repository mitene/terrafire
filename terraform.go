package terrafire

import (
	"os"
	"os/exec"
)

type TerraformClient interface {
	Plan(dir string) error
	Apply(dir string) error
}

type TerraformClientImpl struct {
}

func NewTerraformClient() TerraformClient {
	return &TerraformClientImpl{}
}

func (t *TerraformClientImpl) Plan(dir string) error {
	err := t.init(dir)
	if err != nil {
		return err
	}

	return t.run(dir, "plan")
}

func (t *TerraformClientImpl) Apply(dir string) error {
	err := t.init(dir)
	if err != nil {
		return err
	}

	return t.run(dir, "apply")
}

func (t *TerraformClientImpl) init(dir string) error {
	return t.run(dir, "init")
}

func (t *TerraformClientImpl) run(dir string, command string, arg ...string) error {
	args := append([]string{command}, arg...)
	cmd := exec.Command("terraform", args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
