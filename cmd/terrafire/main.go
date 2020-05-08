package main

import (
	"github.com/mitene/terrafire/core"
	"github.com/mitene/terrafire/database"
	"github.com/mitene/terrafire/runner"
	"github.com/mitene/terrafire/server"
	"github.com/mitene/terrafire/utils"
	"log"
)

func main() {
	config, err := core.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	tf := utils.NewTerraform()

	git := utils.NewGit()
	err = git.Init(config.Repos)
	defer git.Clean()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.NewDB(config)
	if err != nil {
		log.Fatal(err)
	}

	runner_ := runner.NewLocalRunner(config.NumWorkers, tf)
	defer runner_.Clean()

	service := core.NewService(config, runner_, db, git)
	defer service.Close()

	server_ := server.NewServer(config, service)

	err = service.Start()
	if err != nil {
		log.Fatalln(err)
	}

	err = runner_.Start(service)
	if err != nil {
		log.Fatalln(err)
	}

	log.Fatalln(server_.Start())
}
