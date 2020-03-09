package terrafire

import (
	"os"
	"os/exec"
)

type TerraformClient interface {
	Plan() error
	Apply() error
}

type TerraformClientImpl struct {
	dir string
}

func NewTerraformClient(dir string) TerraformClient {
	return &TerraformClientImpl{
		dir: dir,
	}
}

func (t *TerraformClientImpl) Plan() error {
	err := t.init()
	if err != nil {
		return err
	}

	return t.run("plan")
}

func (t *TerraformClientImpl) Apply() error {
	err := t.init()
	if err != nil {
		return err
	}

	return t.run("apply")
}

func (t *TerraformClientImpl) init() error {
	return t.run("init")
}

func (t *TerraformClientImpl) run(command string, arg ...string) error {
	args := append([]string{command}, arg...)
	cmd := exec.Command("terraform", args...)
	cmd.Dir = t.dir
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
