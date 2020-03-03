package terrafire

import (
	"io/ioutil"
	"path"

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
			Owner    string `hcl:"owner"`
			Repo     string `hcl:"repo"`
			Path     string `hcl:"path"`
			Revision string `hcl:"revision"`
		} `hcl:"source,block"`
		Workspace    string            `hcl:"workspace"`
		AllowDestroy *bool             `hcl:"allow_destroy"`
		Vars         map[string]string `hcl:"vars"`
		VarFiles     *[]string         `hcl:"var_files"`
	} `hcl:"terraform_deploy,block"`
}

func LoadConfig(configPath string) (*Config, error) {
	var config Config

	files, err := ioutil.ReadDir(configPath)
	if err != nil {
		return nil, err
	}

	body := []byte{}
	for _, file := range files {
		content, err := ioutil.ReadFile(path.Join(configPath, file.Name()))
		if err != nil {
			return nil, err
		}
		body = append(body, content...)
		body = append(body, []byte("\n")...)
	}

	err = hclsimple.Decode("dummy.hcl", body, nil, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
