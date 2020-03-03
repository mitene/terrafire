package terrafire

import (
	"github.com/hashicorp/hcl/v2/hclsimple"
)

type Config struct {
	Terrafire *struct {
		Backend struct {
			Name   string `hcl:"name,label"`
			Bucket string `hcl:"bucket"`
			Key    string `hcl:"key"`
		} `hcl:"backend,block"`
	} `hcl:"terrafire,block"`
	TerraformDeploy []struct {
		Name   string `hcl:"name,label"`
		Source struct {
			Owner   string `hcl:"owner"`
			Repo string `hcl:"repo"`
			Path    string `hcl:"path"`
			Revision string `hcl:"revision"`
		} `hcl:"source,block"`
		Workspace string `hcl:"workspace"`
		AllowDestroy *bool `hcl:"allow_destroy"`
		Vars map[string]string `hcl:"vars"`
		VarFiles *[]string `hcl:"var_files"`
	} `hcl:"terraform_deploy,block"`
}

func DecodeFile(path string) (*Config, error) {
	var config Config
	err := hclsimple.DecodeFile(path, nil, &config)
	if err != nil {
		return nil, err
	}

	err = hclsimple.DecodeFile("sample/system.hcl", nil, &config)
	if err != nil {
		return nil, err
	}

	err = hclsimple.DecodeFile("sample/app.hcl", nil, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
