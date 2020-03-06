package terrafire

import (
	"fmt"
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

	cmd := exec.Command("terraform", "plan")
	cmd.Dir = t.dir
	out, err := cmd.Output()
	fmt.Println(string(out))
	return err
}

func (t *TerraformClientImpl) init() error {
	cmd := exec.Command("terraform", "init")
	cmd.Dir = t.dir
	out, err := cmd.Output()
	fmt.Println(string(out))
	return err
}

// NEXT: 綺麗にしよう
func (t *TerraformClientImpl) run(command string, arg ...string) error {
	cmd := exec.Command("terraform", []string{command, arg ...} ...)
	cmd.Dir = t.dir
	out, err := cmd.Output()
	fmt.Println(string(out))
	return err
}

func (t *TerraformClientImpl) Apply() error {
	return nil
}