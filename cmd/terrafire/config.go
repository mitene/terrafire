package main

import (
	"fmt"
	"github.com/mitene/terrafire/internal/api"
	"github.com/mitene/terrafire/internal/controller"
	"github.com/mitene/terrafire/internal/runner"
	log "github.com/sirupsen/logrus"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
Server Config
*/
type srvConfig struct {
	ServerPort    int
	SchedulerPort int
	DbDriver      string
	DbAddress     string

	Projects map[string]*api.Project
	Repos    map[string]*api.GitRepository
}

func GetSrvConfig() (config *srvConfig, err error) {
	config = &srvConfig{}

	envs := GetEnvs()

	config.ServerPort = envs.GetInt("TERRAFIRE_SERVER_PORT", 8080)
	config.SchedulerPort = envs.GetInt("TERRAFIRE_SCHEDULER_PORT", 8081)
	config.DbDriver = envs.Get("TERRAFIRE_DB_DRIVER", "sqlite3")
	config.DbAddress = envs.Get("TERRAFIRE_DB_ADDRESS", ":memory:")

	config.Projects, err = envs.GetProjects()
	if err != nil {
		return nil, err
	}

	config.Repos, err = envs.GetRepos()
	if err != nil {
		return nil, err
	}

	return config, nil
}

/*
Controller Config
*/
type ctrlConfig struct {
	SchedulerAddress string
	Concurrency      int
	Executor         controller.Executor
}

func GetCtrlConfig() (config *ctrlConfig, err error) {
	config = &ctrlConfig{}

	envs := GetEnvs()

	config.SchedulerAddress = envs.Get("TERRAFIRE_SCHEDULER_ADDRESS", "localhost:8081")

	config.Concurrency = envs.GetInt("TERRAFIRE_CONCURRENCY", 1)

	switch t := envs.Get("TERRAFIRE_EXECUTOR_TYPE", "local"); t {
	case "local":
		config.Executor = controller.NewLocalExecutor(os.Args[0])
	case "ecs":
		cfg := &controller.ExecutorEcsConfig{
			Cluster:          envs.Get("TERRAFIRE_EXECUTOR_ECS_CLUSTER", ""),
			TaskDefinition:   envs.Get("TERRAFIRE_EXECUTOR_ECS_TASK_DEFINITION", ""),
			ContainerName:    envs.Get("TERRAFIRE_EXECUTOR_ECS_CONTAINER_NAME", "terrafire"),
			CapacityProvider: envs.Get("TERRAFIRE_EXECUTOR_ECS_CAPACITY_PROVIDER", "FARGATE"),
			Subnets:          envs.GetSlice("TERRAFIRE_EXECUTOR_ECS_SUBNETS", ""),
			SecurityGroups:   envs.GetSlice("TERRAFIRE_EXECUTOR_ECS_SECURITY_GROUPS", ""),
			AssignPublicIp:   envs.Get("TERRAFIRE_EXECUTOR_ECS_ASSIGN_PUBLIC_IP", "true") == "true",
		}
		if cfg.Cluster == "" {
			return nil, fmt.Errorf("ecs cluster is not defined")
		}
		if cfg.TaskDefinition == "" {
			return nil, fmt.Errorf("ecs task definiton is not defined")
		}

		config.Executor = controller.NewEcsExecutor(nil)
	default:
		return nil, fmt.Errorf("invalid executor type: %s", t)
	}

	return config, nil
}

/*
Runner Config
*/
type runnerConfig struct {
	SchedulerAddress string

	Projects map[string]*api.Project
	Repos    map[string]*api.GitRepository

	Blob runner.Blob
}

func GetRunnerConfig() (config *runnerConfig, err error) {
	config = &runnerConfig{}

	envs := GetEnvs()

	config.SchedulerAddress = envs.Get("TERRAFIRE_SCHEDULER_ADDRESS", "localhost:8081")

	switch bt := envs.Get("TERRAFIRE_BLOB_TYPE", "local"); bt {
	case "local":
		root := envs.Get("TERRAFIRE_BLOB_LOCAL_ROOT", "")
		if root == "" {
			return nil, fmt.Errorf("TERRAFIRE_BLOB_LOCAL_ROOT is required")
		}
		config.Blob = runner.NewBlobLocal(root)

	case "s3":
		bucket := envs.Get("TERRAFIRE_BLOB_S3_BUCKET", "")
		if bucket == "" {
			return nil, fmt.Errorf("TERRAFIRE_BLOB_S3_BUCKET is required")

		}
		prefix := envs.Get("TERRAFIRE_BLOB_S3_PREFIX", "")
		config.Blob = runner.NewBlobS3(bucket, prefix)

	default:
		return nil, fmt.Errorf("invalid blob type: %s", bt)
	}

	config.Projects, err = envs.GetProjects()
	if err != nil {
		return nil, err
	}

	config.Repos, err = envs.GetRepos()
	if err != nil {
		return nil, err
	}

	return config, nil
}

/*
Utility Functions
*/
type Envs map[string]string

func GetEnvs() Envs {
	envs := map[string]string{}
	for _, e := range os.Environ() {
		vs := strings.SplitN(e, "=", 2)

		key := vs[0]
		val := ""
		if len(vs) > 1 {
			val = vs[1]
		}
		envs[key] = val
	}
	return envs
}

func (e Envs) Get(key string, default_ string) string {
	if v, ok := e[key]; ok {
		return v
	} else {
		return default_
	}
}

func (e Envs) GetSlice(key string, default_ string) []string {
	v := e.Get(key, default_)
	if v == "" {
		return nil
	} else {
		return strings.Split(",", v)
	}
}

func (e Envs) GetInt(key string, default_ int) int {
	if v, ok := e[key]; ok {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Warnf("invalid integer value: %s, use default: %d", v, default_)
			return default_
		}
		return i
	} else {
		return default_
	}
}

func (e Envs) GetProjects() (map[string]*api.Project, error) {
	ret := map[string]*api.Project{}
	re := regexp.MustCompile(`\ATERRAFIRE_PROJECT_([^_]+)(|_BRANCH|_PATH|_ENV_.*)\z`)
	for key, val := range e {
		m := re.FindStringSubmatch(key)
		if m == nil {
			continue
		}
		name := m[1]
		attr := m[2]
		if _, ok := ret[name]; !ok {
			ret[name] = &api.Project{
				Name: name,
				Envs: []*api.Pair{},
			}
		}

		switch attr {
		case "":
			ret[name].Repo = val
		case "_BRANCH":
			ret[name].Branch = val
		case "_PATH":
			ret[name].Path = val
		default:
			if strings.HasPrefix(attr, "_ENV_") {
				key := strings.TrimPrefix(attr, "_ENV_")
				ret[name].Envs = append(ret[name].Envs, &api.Pair{Key: key, Value: val})
			}
		}
	}

	for name, prj := range ret {
		if prj.Repo == "" {
			return nil, fmt.Errorf("repository URL is not set for prject %s", name)
		}
		if prj.Branch == "" {
			prj.Branch = "master"
		}
	}
	return ret, nil
}

func (e Envs) GetRepos() (map[string]*api.GitRepository, error) {
	ret := map[string]*api.GitRepository{}
	re := regexp.MustCompile(`\ATERRAFIRE_GIT_CREDENTIAL_([^_]+)(_PROTOCOL|_HOST|_USER|_PASSWORD)\z`)
	for key, val := range e {
		m := re.FindStringSubmatch(key)
		if m == nil {
			continue
		}
		name := m[1]
		attr := m[2]
		if _, ok := ret[name]; !ok {
			ret[name] = &api.GitRepository{
				Name: name,
			}
		}

		switch attr {
		case "_PROTOCOL":
			ret[name].Protocol = val
		case "_HOST":
			ret[name].Host = val
		case "_USER":
			ret[name].User = val
		case "_PASSWORD":
			ret[name].Password = val
		}
	}

	for name, prj := range ret {
		if prj.Protocol == "" || prj.Host == "" {
			return nil, fmt.Errorf("protocol and host must not be empty in git credential %s", name)
		}
	}
	return ret, nil
}
