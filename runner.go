package terrafire

import (
	"io/ioutil"
)

type Runner interface {
	Plan(dir string) error
}

type RunnerImpl struct {
	github    GithubClient
	terraform TerraformClient
}

func NewRunner(github GithubClient, terraform TerraformClient) Runner {
	return &RunnerImpl{
		github,
		terraform,
	}
}

func (r *RunnerImpl) Plan(dir string) error {
	cfg, err := LoadConfig(dir)
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
