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

	var args []string

	// -var-file="testing.tfvars"
	if params != nil && params.VarFiles != nil {
		for _, vf := range *params.VarFiles {
			args = append(args, "-var-file=" + vf  )
		}
	}

	//-var="image_id=ami-abc123"
	//-var='image_id_list=["ami-abc123","ami-def456"]'
	//-var='image_id_map={"us-east-1":"ami-abc123","us-east-2":"ami-def456"}'
	if params != nil && params.Vars != nil {
		for _, v := range *params.Vars {
			fmt.Println(v)
			//args = append(args, "-var-file=" + vf  )
		}
	}

	return t.run(dir, "plan", args...)
}

func (t *TerraformClientImpl) Apply(dir string, params *ConfigTerraformDeployParams) error {
	err := t.init(dir, params)
	if err != nil {
		return err
	}

	var args []string
	return t.run(dir, "apply", args...)
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
		err = t.run(dir, "workspace", "init")
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
	return cmd.Run()
}
