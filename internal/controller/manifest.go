package controller

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/mitene/terrafire/internal"
	"github.com/zclconf/go-cty/cty"
	ctyjson "github.com/zclconf/go-cty/cty/json"
	"io/ioutil"
	"path/filepath"
)

// Parse all `*.hcl` files in the given directory.
func LoadManifest(dirPath string) (*internal.Manifest, error) {
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

	var config struct {
		Workspaces []*struct {
			Name   string `hcl:"name,label"`
			Source *struct {
				Type  string `hcl:"type,label"`
				Owner string `hcl:"owner"`
				Repo  string `hcl:"repo"`
				Path  string `hcl:"path,optional"`
				Ref   string `hcl:"ref,optional"`
			} `hcl:"source,block"`
			Workspace string     `hcl:"workspace,optional"`
			RawVars   *cty.Value `hcl:"vars,optional"`
			Vars      map[string]string
			VarFiles  []string `hcl:"var_files,optional"`
		} `hcl:"workspace,block"`
	}
	diags = gohcl.DecodeBody(hcl.MergeBodies(bodies), nil, &config)
	if diags.HasErrors() {
		return nil, diags
	}

	out := &internal.Manifest{
		Workspaces: map[string]*internal.Workspace{},
	}
	for _, d := range config.Workspaces {
		vars, err := convertRawVars(d.RawVars)
		if err != nil {
			return nil, err
		}

		varFiles, err := resolveVarFiles(d.VarFiles, dirPath)
		if err != nil {
			return nil, err
		}

		out.Workspaces[d.Name] = &internal.Workspace{
			Name: d.Name,
			Source: &internal.Source{
				Type:  d.Source.Type,
				Owner: d.Source.Owner,
				Repo:  d.Source.Repo,
				Path:  d.Source.Path,
				Ref:   d.Source.Ref,
			},
			Workspace: d.Workspace,
			Vars:      vars,
			VarFiles:  varFiles,
		}
	}

	return out, nil
}

func convertRawVars(value *cty.Value) (map[string]string, error) {
	if value == nil {
		return nil, nil
	}

	ret := map[string]string{}

	for k, v := range value.AsValueMap() {
		if v.Type() == cty.String {
			ret[k] = v.AsString()
		} else {
			j, err := ctyjson.Marshal(v, v.Type())
			if err != nil {
				return nil, err
			}
			ret[k] = string(j)
		}
	}

	return ret, nil
}

func resolveVarFiles(varFiles []string, dirPath string) ([]string, error) {
	if varFiles == nil {
		return nil, nil
	}

	abs, err := filepath.Abs(dirPath)
	if err != nil {
		return nil, err
	}

	files := make([]string, len(varFiles))
	for i, f := range varFiles {
		files[i] = filepath.Clean(filepath.Join(abs, f))
	}

	return files, nil
}
