package main

import (
	"fmt"
	"github.com/mitene/terrafire/internal/api"
	"github.com/mitene/terrafire/internal/controller"
	"github.com/mitene/terrafire/internal/runner"
	"github.com/mitene/terrafire/internal/server"
	"github.com/mitene/terrafire/internal/utils"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func main() {
	cmd := &cobra.Command{
		Use: "terrafire [command]",
	}

	wrap := func(f func(*cobra.Command, []string) error) func(*cobra.Command, []string) {
		return func(cmd *cobra.Command, args []string) {
			utils.LogFatal(f(cmd, args))
		}
	}

	cmd.AddCommand(
		&cobra.Command{
			Use: "server",
			Run: wrap(startServer),
		},

		&cobra.Command{
			Use: "controller",
			Run: wrap(startController),
		},

		&cobra.Command{
			Use: "run PHASE PROJECT WORKSPACE",
			Run: wrap(startRunner),
		},
	)

	utils.LogFatal(cmd.Execute())
}

func startServer(_ *cobra.Command, _ []string) error {
	config, err := GetSrvConfig()
	if err != nil {
		return err
	}

	git := utils.NewGit(config.Repos)

	db, err := server.NewDB(config.DbDriver, config.DbAddress)
	if err != nil {
		return err
	}

	srv := server.New(config.Projects, db, git)

	go func() { utils.LogError(srv.StartScheduler(fmt.Sprintf(":%d", config.SchedulerPort))) }()

	return srv.StartWeb(fmt.Sprintf(":%d", config.ServerPort))
}

func startController(_ *cobra.Command, _ []string) error {
	config, err := GetCtrlConfig()
	if err != nil {
		return err
	}

	conn, err := grpc.Dial(config.SchedulerAddress, grpc.WithInsecure())
	if err != nil {
		return err
	}
	client := api.NewSchedulerClient(conn)

	ctrl := controller.New(client, config.Executor, config.Concurrency)

	return ctrl.Start()
}

func startRunner(_ *cobra.Command, args []string) error {
	config, err := GetRunnerConfig()
	if err != nil {
		return err
	}

	phase := args[0]
	project := args[1]
	workspace := args[2]

	conn, err := grpc.Dial(config.SchedulerAddress, grpc.WithInsecure())
	if err != nil {
		return err
	}
	client := api.NewSchedulerClient(conn)

	tf := runner.NewTerraform()
	git := utils.NewGit(config.Repos)

	runner_ := runner.NewRunner(config.Projects, client, git, tf, config.Blob)

	switch phase {
	case "plan":
		return runner_.Plan(project, workspace)
	case "apply":
		return runner_.Apply(project, workspace)
	default:
		return fmt.Errorf("invalid command: %s", phase)
	}
}
