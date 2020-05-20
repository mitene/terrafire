package core

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func GetConfig() (config *Config, err error) {
	config = &Config{}

	envs := getEnvs()

	config.Address = getEnvWithDefault("TERRAFIRE_ADDRESS", "127.0.0.1:8080")
	config.DataDir = getEnvWithDefault("TERRAFIRE_DATA_DIR", "/usr/local/var/lib/terrafire")

	config.Projects, err = getProjectConfig(envs)
	if err != nil {
		return nil, err
	}

	config.Repos, err = getRepoConfig(envs)
	if err != nil {
		return nil, err
	}

	num := getEnvWithDefault("TERRAFIRE_NUM_WORKERS", "4")
	config.NumWorkers, err = strconv.Atoi(num)
	if err != nil {
		return nil, fmt.Errorf("invalid num_worker config: %s", num)
	}

	return config, nil
}

func getEnvs() map[string]string {
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

func getEnvWithDefault(key string, default_ string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	} else {
		return default_
	}
}

func getProjectConfig(envs map[string]string) (map[string]*Project, error) {
	ret := map[string]*Project{}
	re := regexp.MustCompile(`\ATERRAFIRE_PROJECT_([^_]+)(|_BRANCH|_PATH|_ENV_.*)\z`)
	for key, val := range envs {
		m := re.FindStringSubmatch(key)
		if m == nil {
			continue
		}
		name := m[1]
		attr := m[2]
		if _, ok := ret[name]; !ok {
			ret[name] = &Project{
				Name: name,
				Envs: map[string]string{},
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
				ret[name].Envs[key] = val
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

func getRepoConfig(envs map[string]string) (map[string]*GitCredential, error) {
	ret := map[string]*GitCredential{}
	re := regexp.MustCompile(`\ATERRAFIRE_GIT_CREDENTIAL_([^_]+)(_PROTOCOL|_HOST|_USER|_PASSWORD)\z`)
	for key, val := range envs {
		m := re.FindStringSubmatch(key)
		if m == nil {
			continue
		}
		name := m[1]
		attr := m[2]
		if _, ok := ret[name]; !ok {
			ret[name] = &GitCredential{
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
