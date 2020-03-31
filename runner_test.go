package terrafire

import (
	"fmt"
	"io"
	"strings"
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
		dir    string
		params *ConfigTerraformDeployParams
	}
	var decryptFileInputs []string
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
			plan: func(dir string, params *ConfigTerraformDeployParams) (string, error) {
				terraformArgs.dir = dir
				terraformArgs.params = params
				return "", nil
			},
		},
		&SopsClientMock{
			decryptFile: func(input string, output io.Writer) error {
				decryptFileInputs = append(decryptFileInputs, input)
				return nil
			},
		},
	)
	err := r.Plan("sample", ReportTypeGithub)
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
	if (*terraformArgs.params.Vars)["package_revision"] != "\"xxx\"" {
		t.Fatalf("terraform.Plan: want \"xxx\", got %s", (*terraformArgs.params.Vars)["package_revision"])
	}

	// TODO .enc ファイルの展開先のテンポラリファイル名が .enc が外されたファイル名であることをテストすること
	//if (*terraformArgs.params.VarFiles)["package_revision"] != "\"xxx\"" {
	//	t.Fatalf("terraform.Plan: want \"xxx\", got %s", (*terraformArgs.params.Vars)["package_revision"])
	//}
	fmt.Println(*terraformArgs.params.VarFiles)
	t.Fatal("hoge")

	if len(decryptFileInputs) != 1 {
		t.Fatalf("sops.DecryptFile: called count want 1, got %d", len(decryptFileInputs))
	}
	if !strings.HasSuffix(decryptFileInputs[0], "sample/app/secrets.tfvars.enc") {
		t.Fatalf("sops.DecryptFile: want \"sample/app/secrets.tfvars.enc\", got %s", decryptFileInputs[0])
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
		&SopsClientMock{},
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
