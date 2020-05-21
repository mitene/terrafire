package main

import (
	"github.com/mitene/terrafire"
	"github.com/mitene/terrafire/controller"
	"github.com/mitene/terrafire/database"
	"github.com/mitene/terrafire/executor"
	"github.com/mitene/terrafire/server"
	"github.com/mitene/terrafire/utils"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	config, err := terrafire.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	tmp, err := newTempDir()
	utils.LogFatal(err)
	defer utils.LogError(tmp.Delete())

	gitDir, err := tmp.Create("git")
	utils.LogFatal(err)
	git := utils.NewGit(gitDir)
	utils.LogFatal(git.Init(config.Repos))

	db, err := database.NewDB(config)
	utils.LogFatal(err)

	handler := server.NewHandler(config, db)
	srv := server.NewServer(config, handler)

	blob := executor.NewLocalBlob(filepath.Join(config.DataDir, "blob"))
	runner := executor.NewRunner(handler, blob)
	exe := executor.NewLocalExecutor(handler, runner, config.NumWorkers)

	ctrlDir, err := tmp.Create("controller")
	utils.LogFatal(err)
	ctrl := controller.New(config, handler, exe, git, ctrlDir)

	utils.LogFatal(ctrl.Start())
	defer utils.LogError(ctrl.Stop())

	utils.LogError(handler.RefreshAllProjects())

	log.Fatalln(srv.Start())
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
