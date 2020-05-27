package main

import (
	"github.com/mitene/terrafire"
	"github.com/mitene/terrafire/controller"
	"github.com/mitene/terrafire/database"
	"github.com/mitene/terrafire/executor"
	"github.com/mitene/terrafire/server"
	"github.com/mitene/terrafire/utils"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	utils.LogFatal(runServer())
}

func runServer() error {
	config, err := terrafire.GetConfig()
	if err != nil {
		return err
	}

	tmp, err := newTempDir()
	if err != nil {
		return err
	}
	defer func() { utils.LogError(tmp.Delete()) }()

	gitDir, err := tmp.Create("git")
	if err != nil {
		return err
	}
	git := utils.NewGit(gitDir)
	err = git.Init(config.Repos)
	if err != nil {
		return err
	}

	db, err := database.NewDB(config)
	if err != nil {
		return err
	}

	handler := server.NewHandler(config, db)
	srv := server.NewServer(config, handler)

	blob := executor.NewLocalBlob(filepath.Join(config.DataDir, "blob"))
	runner := executor.NewRunner(handler, blob)
	exe := executor.NewLocalExecutor(handler, runner, config.NumWorkers)

	ctrlDir, err := tmp.Create("controller")
	if err != nil {
		return err
	}
	ctrl := controller.New(config, handler, exe, git, ctrlDir)

	go func() { utils.LogError(ctrl.Start()) }()
	defer func() { utils.LogError(ctrl.Stop()) }()

	utils.LogError(handler.RefreshAllProjects())

	return srv.Start()
}

type tempDir struct {
	root string
}

func newTempDir() (*tempDir, error) {
	dir, err := ioutil.TempDir("", "terrafire-")
	if err != nil {
		return nil, err
	}
	return &tempDir{root: dir}, nil
}

func (d *tempDir) Create(name string) (string, error) {
	p := filepath.Join(d.root, name)
	err := os.MkdirAll(p, 0755)
	if err != nil {
		return "", err
	}
	return p, nil
}

func (d *tempDir) Delete() error {
	return os.RemoveAll(d.root)
}
