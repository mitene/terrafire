package terrafire

import (
	"testing"
)

func TestRunnerImpl_Plan(t *testing.T) {
	var githubArgs struct {
		owner  string
		repo   string
		ref    string
		subDir string
		dest   string
	}
	var terraformArgs struct {
		dir string
	}
	r := NewRunner(
		&GithubClientMock{
			getSource: func(owner string, repo string, ref string, subDir string, dest string) error {
				githubArgs.owner = owner
				githubArgs.repo = repo
				githubArgs.ref = ref
				githubArgs.subDir = subDir
				githubArgs.dest = dest
				return nil
			},
		},
		&TerraformClientMock{
			plan: func(dir string, params *ConfigTerraformDeployParams) error {
				terraformArgs.dir = dir
				return nil
			},
		},
	)
	err := r.Plan("sample")
	if err != nil {
		t.Fatal(err)
	}

	if githubArgs.owner != "terrafire" {
		t.Fatalf("github.getSource: want terrafire, got %s", githubArgs.owner)
	}
	if githubArgs.repo != "terraform" {
		t.Fatalf("github.getSource: want terraform, got %s", githubArgs.repo)
	}
	if githubArgs.ref != "xxxx" {
		t.Fatalf("github.getSource: want xxxx, got %s", githubArgs.ref)
	}
	if githubArgs.subDir != "system/" {
		t.Fatalf("github.getSource: want system/, got %s", githubArgs.subDir)
	}

	if terraformArgs.dir != githubArgs.dest {
		t.Fatalf("terraform.Plan: want %s, got %s", githubArgs.dest, terraformArgs.dir)
	}
}

func TestRunnerImpl_Apply(t *testing.T) {
	var githubArgs struct {
		owner  string
		repo   string
		ref    string
		subDir string
		dest   string
	}
	var terraformArgs struct {
		dir string
	}
	r := NewRunner(
		&GithubClientMock{
			getSource: func(owner string, repo string, ref string, subDir string, dest string) error {
				githubArgs.owner = owner
				githubArgs.repo = repo
				githubArgs.ref = ref
				githubArgs.subDir = subDir
				githubArgs.dest = dest
				return nil
			},
		},
		&TerraformClientMock{
			apply: func(dir string, params *ConfigTerraformDeployParams) error {
				terraformArgs.dir = dir
				return nil
			},
		},
	)
	err := r.Apply("sample")
	if err != nil {
		t.Fatal(err)
	}

	if githubArgs.owner != "terrafire" {
		t.Fatalf("github.getSource: want terrafire, got %s", githubArgs.owner)
	}
	if githubArgs.repo != "terraform" {
		t.Fatalf("github.getSource: want terraform, got %s", githubArgs.repo)
	}
	if githubArgs.ref != "xxxx" {
		t.Fatalf("github.getSource: want xxxx, got %s", githubArgs.ref)
	}
	if githubArgs.subDir != "system/" {
		t.Fatalf("github.getSource: want system/, got %s", githubArgs.subDir)
	}

	if terraformArgs.dir != githubArgs.dest {
		t.Fatalf("terraform.Apply: want %s, got %s", githubArgs.dest, terraformArgs.dir)
	}
}
