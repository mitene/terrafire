package manifest

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/mitene/terrafire/internal/api"
	"github.com/zclconf/go-cty/cty"
	ctyjson "github.com/zclconf/go-cty/cty/json"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"
)

// Parse all `*.hcl` files in the given directory.
func Load(dirPath string) ([]*api.Workspace, error) {
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
			Vars      *cty.Value `hcl:"vars,optional"`
			VarFiles  []string   `hcl:"var_files,optional"`
		} `hcl:"workspace,block"`
	}
	diags = gohcl.DecodeBody(hcl.MergeBodies(bodies), nil, &config)
	if diags.HasErrors() {
		return nil, diags
	}

	out := make([]*api.Workspace, 0, len(config.Workspaces))

	for _, d := range config.Workspaces {
		vars, err := convertRawVars(d.Vars)
		if err != nil {
			return nil, err
		}

		varFiles, err := resolveVarFiles(dirPath, d.VarFiles)
		if err != nil {
			return nil, err
		}

		sourceType, ok := api.Source_Type_value[d.Source.Type]
		if !ok {
			return nil, fmt.Errorf("invalid source type: %s", d.Source.Type)
		}

		out = append(out, &api.Workspace{
			Name: d.Name,
			Source: &api.Source{
				Type:  api.Source_Type(sourceType),
				Owner: d.Source.Owner,
				Repo:  d.Source.Repo,
				Path:  d.Source.Path,
				Ref:   d.Source.Ref,
			},
			Workspace: d.Workspace,
			Vars:      vars,
			VarFiles:  varFiles,
		})
	}

	return out, nil
}

func convertRawVars(value *cty.Value) ([]*api.Pair, error) {
	if value == nil {
		return nil, nil
	}

	m := value.AsValueMap()

	ret := make([]*api.Pair, 0, len(m))

	for k, v := range m {
		var v1 string
		if v.Type() == cty.String {
			v1 = v.AsString()
		} else {
			j, err := ctyjson.Marshal(v, v.Type())
			if err != nil {
				return nil, err
			}
			v1 = string(j)
		}
		ret = append(ret, &api.Pair{Key: k, Value: v1})
	}

	sort.Slice(ret, func(i, j int) bool { return ret[i].Key < ret[j].Key }) // sort by key

	return ret, nil
}

func resolveVarFiles(dir string, varFiles []string) ([]*api.Pair, error) {
	if varFiles == nil {
		return nil, nil
	}

	ret := make([]*api.Pair, len(varFiles))
	for i, f := range varFiles {
		fp := filepath.Clean(f)
		if strings.HasPrefix("..", fp) {
			return nil, fmt.Errorf("cannot read parent directory: %s", fp)
		}

		body, err := ioutil.ReadFile(filepath.Join(dir, fp))
		if err != nil {
			return nil, err
		}

		ret[i] = &api.Pair{Key: fp, Value: string(body)}
	}

	return ret, nil
}
