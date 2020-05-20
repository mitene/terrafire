package core

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

type projectService struct {
	project  *Project
	manifest *Manifest
	mux      sync.Mutex
	dir      string
	git      Git
}

func newProjectService(project *Project, git Git) *projectService {
	return &projectService{
		project:  project,
		manifest: nil,

		mux: sync.Mutex{},
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

	err = s.refresh()
	if err != nil {
		return err
	}

	return nil
}

func (s *projectService) close() error {
	if s.dir != "" {
		return os.RemoveAll(s.dir)
	}
	return nil
}

func (s *projectService) refresh() error {
	s.mux.Lock()
	defer s.mux.Unlock()

	pj := s.project
	commit, err := s.git.Fetch(s.dir, pj.Repo, pj.Branch)
	if err != nil {
		return err
	}

	pj.Commit = commit

	m, err := LoadManifest(filepath.Join(s.dir, pj.Path))
	if err != nil {
		return err
	}

	s.manifest = m
	return nil
}
