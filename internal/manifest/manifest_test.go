package manifest

import (
	"github.com/mitene/terrafire/internal/api"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

func TestLoadManifest(t *testing.T) {
	tests := []struct {
		argFiles     map[string]string
		wantManifest []*api.Workspace
	}{
		{
			argFiles: map[string]string{
				"f1.hcl": `
workspace "ws1" {
  source "github" {
    owner = "ws1-owner"
    repo  = "ws1-repo"
    path  = "ws1-path"
    ref   = "ws1-ref"
  }
  workspace = "ws1-workspace"
  vars = {
    "ws1-k1" = "ws1-v1"
    "ws1-k2" = ["ws1-v2-1", "ws1-v2-2"]
    "ws1-k3" = {
      "ws1-v3-k1" = "ws1-v3-v1"
      "ws1-v3-k2" = "ws1-v3-v2"
    }
  }
  var_files = [
    "ws1-f1.tfvars",
  ]
}
`,
				"f2.hcl": `
workspace "ws2" {
  source "github" {
    owner = "ws2-owner"
    repo  = "ws2-repo"
    path  = "ws2-path"
    ref   = "ws2-ref"
  }
}
`,
				"ws1-f1.tfvars": `
ws1-f1-k1 = "ws1-f1-v1"
`,
			},
			wantManifest: []*api.Workspace{
				{
					Name: "ws1",
					Source: &api.Source{
						Type:  api.Source_github,
						Owner: "ws1-owner",
						Repo:  "ws1-repo",
						Path:  "ws1-path",
						Ref:   "ws1-ref",
					},
					Workspace: "ws1-workspace",
					Vars: []*api.Pair{
						{Key: "ws1-k1", Value: "ws1-v1"},
						{Key: "ws1-k2", Value: "[\"ws1-v2-1\",\"ws1-v2-2\"]"},
						{Key: "ws1-k3", Value: "{\"ws1-v3-k1\":\"ws1-v3-v1\",\"ws1-v3-k2\":\"ws1-v3-v2\"}"},
					},
					VarFiles: []*api.Pair{
						{Key: "ws1-f1.tfvars", Value: "\nws1-f1-k1 = \"ws1-f1-v1\"\n"},
					},
				},
				{
					Name: "ws2",
					Source: &api.Source{
						Type:  api.Source_github,
						Owner: "ws2-owner",
						Repo:  "ws2-repo",
						Path:  "ws2-path",
						Ref:   "ws2-ref",
					},
				},
			},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			dir, err := ioutil.TempDir("", "")
			assert.NoError(t, err)
			defer func() { _ = os.RemoveAll(dir) }()

			for fp, body := range tt.argFiles {
				fp = filepath.Join(dir, fp)
				assert.NoError(t, os.MkdirAll(filepath.Dir(fp), 0755))
				assert.NoError(t, ioutil.WriteFile(fp, []byte(body), 0644))
			}

			man, err := Load(dir)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantManifest, man)
		})
	}
}
