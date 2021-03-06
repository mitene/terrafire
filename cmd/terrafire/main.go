package main

import (
	"fmt"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/mitene/terrafire/internal/api"
	"github.com/mitene/terrafire/internal/controller"
	"github.com/mitene/terrafire/internal/runner"
	"github.com/mitene/terrafire/internal/server"
	"github.com/mitene/terrafire/internal/utils"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"time"
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
	go func() { utils.LogError(srv.StartJobObserver()) }()

	return srv.StartWeb(fmt.Sprintf(":%d", config.ServerPort))
}

func startController(_ *cobra.Command, _ []string) error {
	config, err := GetCtrlConfig()
	if err != nil {
		return err
	}

	client, err := newSchedulerClient(config.SchedulerAddress)
	if err != nil {
		return err
	}

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

	client, err := newSchedulerClient(config.SchedulerAddress)
	if err != nil {
		return err
	}

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

func newSchedulerClient(address string) (api.SchedulerClient, error) {
	retryOpts := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(func(attempt uint) time.Duration {
			if attempt > 10 {
				return 30 * time.Second
			} else {
				return time.Duration(attempt) * 2 * time.Second
			}
		}),
	}

	conn, err := grpc.Dial(address,
		grpc.WithStreamInterceptor(grpc_retry.StreamClientInterceptor(retryOpts...)),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(retryOpts...)),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	client := api.NewSchedulerClient(conn)
	return client, nil
}
