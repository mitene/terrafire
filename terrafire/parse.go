package terrafire

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	"io/ioutil"
	"path/filepath"
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

// Parse all `*.hcl` files in the given directory.
func LoadConfig(dirPath string) (*Config, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	var bodies []hcl.Body
	diags := hcl.Diagnostics{}
	parser := hclparse.NewParser()
	for _, file := range files {
		fileName := file.Name()
		suffix := filepath.Ext(fileName)
		if suffix != ".hcl" {
			continue
		}

		f, d := parser.ParseHCLFile(filepath.Join(dirPath, fileName))
		if f != nil {
			bodies = append(bodies, f.Body)
		}
		diags = diags.Extend(d)
	}
	if diags.HasErrors() {
		return nil, diags
	}

	var config Config

	merged := hcl.MergeBodies(bodies)
	diags = gohcl.DecodeBody(merged, nil, &config)
	if diags.HasErrors() {
		return nil, diags
	}

	return &config, nil
}