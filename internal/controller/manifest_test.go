package controller

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestLoadManifest(t *testing.T) {
	cwd, _ := os.Getwd()
	dirPath := filepath.Join(cwd, "manifest_test")

	v, err := LoadManifest(dirPath)
	assert.NoError(t, err)

	assert.Equal(t, "terrafire", v.Workspaces["app"].Source.Owner)
	assert.Equal(t, filepath.Join(dirPath, "app/variables.tfvars"), v.Workspaces["app"].VarFiles[0])
	assert.Equal(t, "xxx", v.Workspaces["app"].Vars["foo_revision"])
	assert.Equal(t, "[\"ami-abc123\",\"ami-def456\"]", v.Workspaces["app"].Vars["ami_list"])
	assert.Equal(t, "{\"us-east-1\":\"ami-abc123\",\"us-east-2\":\"ami-def456\"}", v.Workspaces["app"].Vars["region_map"])
	assert.Equal(t, "terrafire", v.Workspaces["system"].Source.Owner)
	assert.Equal(t, v.Workspaces["system"].Vars["package_revision"], "xxx")
}
