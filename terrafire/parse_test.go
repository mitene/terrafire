package terrafire

import (
	"os"
	"path"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	cwd, _ := os.Getwd()
	v, err := LoadConfig(path.Join(cwd, "..", "sample"))
	if err != nil {
		t.Fatal(err)
	}

	if v.Terrafire.Backend.Name != "s3" {
		t.Fatalf("terrafire.backend.name: want s3, got %s", v.Terrafire.Backend.Name)
	}
	if v.Terrafire.Backend.Bucket != "state_bucket" {
		t.Fatalf("terrafire.backend.bucket: want state_bucket, got %s", v.Terrafire.Backend.Bucket)
	}
	if v.Terrafire.Backend.Key != "state_file" {
		t.Fatalf("terrafire.backend.key: want state_file, got %s", v.Terrafire.Backend.Key)
	}

	if v.TerraformDeploy[0].Name != "app" {
		t.Fatalf("terraform_deploy[0].name: want app, got %s", v.TerraformDeploy[0].Name)
	}
	if v.TerraformDeploy[0].Source.Owner != "terrafire" {
		t.Fatalf("terraform_deploy[0].source.owner: want terrafire, got %s", v.TerraformDeploy[0].Source.Owner)
	}
	if v.TerraformDeploy[0].Vars["foo_revision"] != "xxx" {
		t.Fatalf("terraform_deploy[0].vars[\"foo_revision\"]: want xxx, got %s", v.TerraformDeploy[0].Vars["foo_revision"])
	}

	if v.TerraformDeploy[1].Name != "system" {
		t.Fatalf("terraform_deploy[1].name: want system, got %s", v.TerraformDeploy[1].Name)
	}
	if v.TerraformDeploy[1].Source.Owner != "terrafire" {
		t.Fatalf("terraform_deploy[1].source.owner: want terrafire, got %s", v.TerraformDeploy[1].Source.Owner)
	}
	if v.TerraformDeploy[1].Vars["package_revision"] != "xxx" {
		t.Fatalf("terraform_deploy[1].vars[\"package_revision\"]: want xxx, got %s", v.TerraformDeploy[1].Vars["package_revision"])
	}
}
