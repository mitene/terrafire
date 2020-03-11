package terrafire

import (
	"io/ioutil"
	"path/filepath"

	"github.com/zclconf/go-cty/cty"
	cty_json "github.com/zclconf/go-cty/cty/json"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
)

type Config struct {
	Terrafire *struct {
		Backend struct {
			Name   string `hcl:"name,label"`
			Bucket string `hcl:"bucket"`
			Key    string `hcl:"key"`
		} `hcl:"backend,block"`
	} `hcl:"terrafire,block"`
	TerraformDeploy []ConfigTerraformDeploy `hcl:"terraform_deploy,block"`
}

type ConfigTerraformDeploy struct {
	Name   string `hcl:"name,label"`
	Source struct {
		Owner    string `hcl:"owner"`
		Repo     string `hcl:"repo"`
		Path     string `hcl:"path"`
		Revision string `hcl:"revision"`
	} `hcl:"source,block"`
	Params       *ConfigTerraformDeployParams `hcl:"params,block"`
	AllowDestroy *bool                        `hcl:"allow_destroy"`
}

type ConfigTerraformDeployParams struct {
	Workspace string     `hcl:"workspace"`
	RawVars   *cty.Value `hcl:"vars"`
	Vars      *map[string]string
	VarFiles  *[]string `hcl:"var_files"`
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

		if suffix := filepath.Ext(fileName); suffix != ".hcl" {
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

	for _, d := range config.TerraformDeploy {
		vars, err := convertRawVars(d.Params.RawVars)
		if err != nil {
			return nil, err
		}
		d.Params.Vars = vars

		d.Params.VarFiles, err = resolveVarFiles(d.Params.VarFiles, dirPath)
		if err != nil {
			return nil, err
		}
	}

	return &config, nil
}

func convertRawVars(value *cty.Value) (*map[string]string, error) {
	if value == nil {
		return nil, nil
	}

	ret := map[string]string{}

	for k, v := range value.AsValueMap() {
		j, err := cty_json.Marshal(v, v.Type())
		if err != nil {
			return nil, err
		}
		ret[k] = string(j)
	}

	return &ret, nil
}

func resolveVarFiles(varFiles *[]string, dirPath string) (*[]string, error) {
	if varFiles == nil {
		return nil, nil
	}

	abs, err := filepath.Abs(dirPath)
	if err != nil {
		return nil, err
	}

	files := make([]string, len(*varFiles))
	for i, f := range *varFiles {
		s := filepath.Join(abs, f)
		files[i] = filepath.Clean(s)
	}

	return &files, nil
}
