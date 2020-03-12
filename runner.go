package terrafire

import (
	"io/ioutil"
	"os"
)

type Runner interface {
	Plan(dir string) error
	Apply(dir string) error
}

type RunnerImpl struct {
	github    GithubClient
	terraform TerraformClient
}

type PlanResult struct {
	Name  string
	Body  string
	Error error
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

	results := map[string]PlanResult{}
	for _, deploy := range cfg.TerraformDeploy {
		result, err := r.planSingle(deploy)
		results[deploy.Name] = PlanResult{
			Name:  deploy.Name,
			Body:  result,
			Error: err,
		}
	}

	return nil
}

func (r *RunnerImpl) planSingle(deploy ConfigTerraformDeploy) (string, error) {
	tmpDir, err := ioutil.TempDir("", "")
	if err != nil {
		return "", err
	}
	defer os.RemoveAll(tmpDir)

	err = r.github.GetSource(deploy.Source.Owner, deploy.Source.Repo, deploy.Source.Revision, deploy.Source.Path, tmpDir)
	if err != nil {
		return "", err
	}

	result, err := r.terraform.Plan(tmpDir, deploy.Params)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (r *RunnerImpl) Apply(dir string) error {
	cfg, err := LoadConfig(dir)
	if err != nil {
		return err
	}

	for _, deploy := range cfg.TerraformDeploy {
		err := r.applySingle(deploy)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *RunnerImpl) applySingle(deploy ConfigTerraformDeploy) error {
	tmpDir, err := ioutil.TempDir("", "")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmpDir)

	err = r.github.GetSource(deploy.Source.Owner, deploy.Source.Repo, deploy.Source.Revision, deploy.Source.Path, tmpDir)
	if err != nil {
		return err
	}

	err = r.terraform.Apply(tmpDir, deploy.Params)
	if err != nil {
		return err
	}

	return nil
}
