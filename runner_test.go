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
		&TerraformClientMock{},
	)
	err := r.Plan("sample/terraform")
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
}
