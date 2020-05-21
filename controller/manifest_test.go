package controller

import (
	"os"
	"path"
	"path/filepath"
	"testing"
)

func TestLoadManifest(t *testing.T) {
	cwd, _ := os.Getwd()
	dirPath := path.Clean(path.Join(cwd, "../sample"))
	v, err := LoadManifest(dirPath)
	if err != nil {
		t.Fatal(err)
	}

	if v.Workspaces["app"].Source.Owner != "terrafire" {
		t.Fatalf("terraform_deploy[0].source.owner: want terrafire, got %s", v.Workspaces["app"].Source.Owner)
	}

	if varfile := v.Workspaces["app"].VarFiles[0]; varfile != filepath.Join(dirPath, "app/variables.tfvars") {
		t.Fatalf("terraformDeploy[0].VarFiles[0]: want %s, got %s", filepath.Join(dirPath, "app/variables.tfvars"), varfile)
	}

	if vars := v.Workspaces["app"].Vars; vars["foo_revision"] != "xxx" {
		t.Fatalf("terraform_deploy[0].vars[\"foo_revision\"]: want xxx, got %s", vars["foo_revision"])
	}
	if vars := v.Workspaces["app"].Vars; vars["ami_list"] != "[\"ami-abc123\",\"ami-def456\"]" {
		t.Fatalf("terraform_deploy[0].vars[\"ami_list\"]: want [\"ami-abc123\",\"ami-def456\"], got %s", vars["ami_list"])
	}
	if vars := v.Workspaces["app"].Vars; vars["region_map"] != "{\"us-east-1\":\"ami-abc123\",\"us-east-2\":\"ami-def456\"}" {
		t.Fatalf("terraform_deploy[0].vars[\"region_map\"]: want {\"us-east-1\":\"ami-abc123\",\"us-east-2\":\"ami-def456\"}, got %s", vars["region_map"])
	}

	if v.Workspaces["system"].Source.Owner != "terrafire" {
		t.Fatalf("terraform_deploy[1].source.owner: want terrafire, got %s", v.Workspaces["system"].Source.Owner)
	}
	if vars := v.Workspaces["system"].Vars; vars["package_revision"] != "xxx" {
		t.Fatalf("terraform_deploy[1].vars[\"package_revision\"]: want xxx, got %s", vars["package_revision"])
	}
}
