package terrafire

import (
	"io/ioutil"
	"os"
)

type Runner interface {
	Plan() error
}

type RunnerImpl struct {
	github GithubClient
	terraform TerraformClient
}

func NewRunner(github GithubClient, terraform TerraformClient) Runner {
	return &RunnerImpl{
		github,
		terraform,
	}
}

func (r *RunnerImpl) Plan() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	cfg, err := LoadConfig(cwd)
	if err != nil {
		return err
	}

	for _, deploy := range cfg.TerraformDeploy {
		tmpDir, err := ioutil.TempDir("", "")
		if err != nil {
			return err
		}
		err = r.github.GetSource(deploy.Source.Owner, deploy.Source.Repo, deploy.Source.Revision, deploy.Source.Path, tmpDir)
		if err != nil {
			return err
		}

		err = r.terraform.Plan(tmpDir)
		if err != nil {
			return err
		}
	}

	return nil
}
