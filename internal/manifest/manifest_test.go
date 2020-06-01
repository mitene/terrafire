package manifest

import (
	"github.com/mitene/terrafire/internal/api"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestLoadManifest(t *testing.T) {
	cwd, _ := os.Getwd()
	dirPath := filepath.Join(cwd, "manifest_test")

	v, err := Load(dirPath)
	assert.NoError(t, err)

	assert.Equal(t, &api.Manifest{
		Workspaces: []*api.Workspace{
			{
				Name: "app",
				Source: &api.Source{
					Type:  api.Source_github,
					Owner: "terrafire",
					Repo:  "terraform",
					Path:  "app/",
					Ref:   "xxxx",
				},
				Workspace: "dev",
				Vars: []*api.Pair{
					{Key: "ami_list", Value: "[\"ami-abc123\",\"ami-def456\"]"},
					{Key: "foo_revision", Value: "xxx"},
					{Key: "region_map", Value: "{\"us-east-1\":\"ami-abc123\",\"us-east-2\":\"ami-def456\"}"},
				},
				VarFiles: []string{
					"app/variables.tfvars",
					"app/secrets.tfvars.enc",
				},
			},
			{
				Name: "system",
				Source: &api.Source{
					Type:  api.Source_github,
					Owner: "terrafire",
					Repo:  "terraform",
					Path:  "system/",
					Ref:   "xxxx",
				},
				Workspace: "dev",
				Vars: []*api.Pair{
					{Key: "package_revision", Value: "xxx"},
				},
			},
		},
	}, v)
}
