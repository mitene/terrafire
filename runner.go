package terrafire

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Runner interface {
	Plan(dir string, reportType ReportType) error
	Apply(dir string, autoApprove bool) error
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

type PlanResults map[string]PlanResult

type ReportType int

const (
	ReportTypeNone = iota
	ReportTypeGithub
)

func NewRunner(github GithubClient, terraform TerraformClient) Runner {
	return &RunnerImpl{
		github,
		terraform,
	}
}

func (r *RunnerImpl) Plan(dir string, reportTo ReportType) error {
	cfg, err := LoadConfig(dir)
	if err != nil {
		return err
	}

	results := PlanResults{}
	failed := 0
	for _, deploy := range cfg.TerraformDeploy {
		result, err := r.planSingle(deploy)
		results[deploy.Name] = PlanResult{
			Name:  deploy.Name,
			Body:  result,
			Error: err,
		}
		if err != nil {
			log.Print(err.Error())
			failed++
		}
	}

	if reportTo == ReportTypeGithub {
		err = NewReporterGithub(r.github).Report(results)
		if err != nil {
			return err
		}
	}

	if failed > 0 {
		return fmt.Errorf("plan failed: %d of %d", failed, len(cfg.TerraformDeploy))
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

func (r *RunnerImpl) Apply(dir string, autoApprove bool) error {
	cfg, err := LoadConfig(dir)
	if err != nil {
		return err
	}

	failed := 0
	for _, deploy := range cfg.TerraformDeploy {
		err := r.applySingle(deploy, autoApprove)
		if err != nil {
			log.Print(err.Error())
			failed++
		}
	}

	if failed > 0 {
		return fmt.Errorf("apply failed: %d of %d", failed, len(cfg.TerraformDeploy))
	}
	return nil
}

func (r *RunnerImpl) applySingle(deploy ConfigTerraformDeploy, autoApprove bool) error {
	tmpDir, err := ioutil.TempDir("", "")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmpDir)

	err = r.github.GetSource(deploy.Source.Owner, deploy.Source.Repo, deploy.Source.Revision, deploy.Source.Path, tmpDir)
	if err != nil {
		return err
	}

	err = r.terraform.Apply(tmpDir, deploy.Params, autoApprove)
	if err != nil {
		return err
	}

	return nil
}
