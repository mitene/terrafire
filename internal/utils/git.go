package utils

import (
	"github.com/mitene/terrafire/internal/api"
	"os"
	"os/exec"
	"path/filepath"
)

type Git interface {
	Fetch(dir string, repo string, branch string) (string, error)
}

type git struct {
	repos map[string]*api.GitRepository
}

func NewGit(repos map[string]*api.GitRepository) Git {
	return &git{repos: repos}
}

func (g *git) Fetch(dir string, repo string, branch string) (string, error) {
	if _, err := os.Stat(filepath.Join(dir, ".git")); err == nil {
		err := g.run(dir, "fetch", "origin", branch, "--depth=1")
		if err != nil {
			return "", err
		}

		err = g.run(dir, "reset", "--hard", "origin/"+branch)
		if err != nil {
			return "", err
		}
	} else {
		err = g.run(dir, "clone", repo, ".", "--depth=1", "--branch="+branch, "--single-branch")
		if err != nil {
			return "", err
		}
	}

	rev, err := g.output(dir, "rev-parse", "HEAD")
	if err != nil {
		return "", err
	}
	return string(rev), nil
}

func (g *git) run(dir string, arg ...string) (err error) {
	err1 := g.withCredentials(func(credArgs []string) {
		cmd := exec.Command("git", append(credArgs, arg...)...)
		cmd.Dir = dir
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
	})
	if err1 != nil {
		return err1
	}

	return
}

func (g *git) output(dir string, arg ...string) (out []byte, err error) {
	err1 := g.withCredentials(func(credArgs []string) {
		cmd := exec.Command("git", append(credArgs, arg...)...)
		cmd.Dir = dir
		out, err = cmd.Output()
	})
	if err1 != nil {
		return nil, err1
	}

	return
}

func (g *git) withCredentials(f func(credArgs []string)) error {
	credfile, err := TempFile()
	if err != nil {
		return err
	}
	defer LogDefer(func() error { return os.Remove(credfile.Name()) })

	err = credfile.Chmod(0400)
	if err != nil {
		return err
	}

	for _, c := range g.repos {
		proto, host := c.Protocol, c.Host
		if c.User != "" || c.Password != "" {
			host = c.User + ":" + c.Password + "@" + host
		}

		_, err = credfile.WriteString(proto + "://" + host + "\n")
		if err != nil {
			return err
		}
	}

	err = credfile.Close()
	if err != nil {
		return err
	}

	f([]string{
		"-c", "credential.helper=store --file " + credfile.Name(),
		"-c", "core.askpass=echo", // disable askpass
	})

	return nil
}
