package core

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type projectService struct {
	project  *Project
	manifest *Manifest
	ch       chan func()
	dir      string
	git      Git
}

func newProjectService(project *Project, git Git) *projectService {
	return &projectService{
		project:  project,
		manifest: nil,

		ch:  make(chan func()),
		dir: "", // initialized when start
		git: git,
	}
}

func (s *projectService) start() error {
	dir, err := ioutil.TempDir("", s.project.Name)
	if err != nil {
		return err
	}
	s.dir = dir

	err = s.refreshNow()
	if err != nil {
		return err
	}

	go func() {
		for range s.ch {
			if err := s.refreshNow(); err != nil {
				log.Println("ERROR: " + err.Error())
			}
		}
	}()

	return nil
}

func (s *projectService) close() error {
	if s.dir != "" {
		return os.RemoveAll(s.dir)
	}
	return nil
}

func (s *projectService) refresh() {
	s.ch <- nil
}

func (s *projectService) refreshNow() error {
	pj := s.project
	err := s.git.Fetch(s.dir, pj.Repo, pj.Branch)
	if err != nil {
		return err
	}

	m, err := LoadManifest(filepath.Join(s.dir, pj.Path))
	if err != nil {
		return err
	}

	s.manifest = m
	return nil
}
