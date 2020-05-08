package core

import (
	"os"
	"reflect"
	"testing"
)

func TestGetEnvs(t *testing.T) {
	envs := getEnvs()

	if _, ok := envs["PATH"]; !ok {
		t.Fatalf("environment variable PATH does not exists")
	}

	if envs["PATH"] != os.Getenv("PATH") {
		t.Fatalf("PATH environment variable, got %s, want %s", envs["PATH"], os.Getenv("PATH"))
	}
}

func TestGetProjectConfig(t *testing.T) {
	prjs, err := getProjectConfig(map[string]string{
		"TERRAFIRE_PROJECT_dev": "https://github.com/foo/bar",

		"TERRAFIRE_PROJECT_dev-1":        "https://github.com/foo/bar1",
		"TERRAFIRE_PROJECT_dev-1_BRANCH": "dev",
		"TERRAFIRE_PROJECT_dev-1_PATH":   "subdir",
	})
	if err != nil {
		t.Fatal(err)
	}

	expected1 := Project{
		Name:   "dev",
		Repo:   "https://github.com/foo/bar",
		Branch: "master",
		Path:   "",
	}
	if !reflect.DeepEqual(*prjs["dev"], expected1) {
		t.Fatalf("project dev, got %q, want %q", *prjs["dev"], expected1)
	}

	expected2 := Project{
		Name:   "dev-1",
		Repo:   "https://github.com/foo/bar1",
		Branch: "dev",
		Path:   "subdir",
	}
	if !reflect.DeepEqual(*prjs["dev-1"], expected2) {
		t.Fatalf("project dev-1, got %q, want %q", *prjs["dev-1"], expected2)
	}
}
