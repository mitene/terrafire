package utils

import (
	"github.com/mitene/terrafire/core"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

type git struct {
	dir string
}

func NewGit() core.Git {
	return &git{}
}

func (g *git) Init(credentials map[string]*core.GitCredential) (err error) {
	gitpath, err := exec.LookPath("git")
	if err != nil {
		return err
	}

	g.dir, err = ioutil.TempDir("", "terrafire-git")
	if err != nil {
		return err
	}

	// save credentials
	credfile := filepath.Join(g.dir, "credentials")
	creds := ""
	for _, c := range credentials {
		proto, host := c.Protocol, c.Host
		if c.User != "" || c.Password != "" {
			host = c.User + ":" + c.Password + "@" + host
		}
		creds += proto + "://" + host + "\n"
	}

	err = ioutil.WriteFile(credfile, []byte(creds), 0400)
	if err != nil {
		return err
	}

	// create git alias command
	bindir := filepath.Join(g.dir, "bin")
	err = os.Mkdir(bindir, 0755)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath.Join(bindir, "git"), []byte(
		"#!/bin/sh\n"+
			"exec "+gitpath+
			" -c credential.helper=\"store --file "+credfile+"\""+
			" -c core.askpass=echo"+ // disable askpass
			" \"$@\""), 0755)
	if err != nil {
		return err
	}

	return os.Setenv("PATH", bindir+string(os.PathListSeparator)+os.Getenv("PATH"))
}

func (g *git) Clean() error {
	if g.dir != "" {
		return os.RemoveAll(g.dir)
	}
	return nil
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

	cmd := exec.Command("git", "rev-parse", "HEAD")
	cmd.Dir = dir
	rev, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(rev), nil
}

func (g *git) run(dir string, arg ...string) error {
	cmd := exec.Command("git", arg...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
