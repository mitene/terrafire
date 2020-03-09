package terrafire

import (
	"io/ioutil"
	"os"
)

type Runner interface {
	Plan() error
}

type RunnerImpl struct {
}

func NewRunner() Runner {
	return &RunnerImpl{}
}

// config をパースしたい
// config をもとにソースコードをもってくる
// その上でコマンドを実行する
func (r *RunnerImpl) Plan() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	cfg, err := LoadConfig(cwd)
	if err != nil {
		return err
	}

	client := NewGithubClient()

	for _, deploy := range cfg.TerraformDeploy {
		tmpDir, err := ioutil.TempDir("", "")
		if err != nil {
			return err
		}
		err = client.GetSource(deploy.Source.Owner, deploy.Source.Repo, deploy.Source.Revision, deploy.Source.Path, tmpDir)
		if err != nil {
			return err
		}

		tc := NewTerraformClient(tmpDir)
		err = tc.Plan()
		if err != nil {
			return err
		}
	}

	return nil
}
