package core

import (
	"fmt"
	"log"
)

type Service struct {
	projects        map[string]*Project
	projectServices map[string]*projectService
	config          *Config
	runner          JobRunner
	db              DB
	git             Git
}

func NewService(config *Config, runner JobRunner, db DB, git Git) *Service {
	projectServices := map[string]*projectService{}
	for name, pj := range config.Projects {
		projectServices[name] = newProjectService(pj, git)
	}

	return &Service{
		projects:        config.Projects,
		projectServices: projectServices,
		config:          config,
		runner:          runner,
		db:              db,
		git:             git,
	}
}

func (s *Service) Start() error {
	for _, svc := range s.projectServices {
		err := svc.start()
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) Close() (err error) {
	for _, svc := range s.projectServices {
		err = svc.close()
	}
	return
}

func (s *Service) GetProjects() map[string]*Project {
	return s.projects
}

func (s *Service) GetProject(project string) (*Project, error) {
	pj, ok := s.projects[project]
	if !ok {
		return nil, fmt.Errorf("project %s is not defined", project)
	}
	return pj, nil
}

func (s *Service) GetWorkspaces(project string) (map[string]*Workspace, error) {
	pj, ok := s.projectServices[project]
	if !ok {
		return nil, fmt.Errorf("project %s is not defiend", project)
	}

	m := pj.manifest
	if m == nil {
		return nil, fmt.Errorf("failed to load project manifest")
	}

	return m.Workspaces, nil
}

func (s *Service) GetWorkspace(project string, workspace string) (*Workspace, error) {
	wss, err := s.GetWorkspaces(project)
	if err != nil {
		return nil, err
	}

	ws, ok := wss[workspace]
	if !ok {
		return nil, fmt.Errorf("workspace %s is not defined", workspace)
	}

	job, err := s.db.GetWorkspaceJob(project, workspace)
	if err != nil {
		return nil, err
	}
	ws.LastJob = job

	return ws, nil
}

func (s *Service) SubmitJob(project string, workspace string) (job *Job, err error) {
	pj, err := s.GetProject(project)
	if err != nil {
		return nil, err
	}

	ws, err := s.GetWorkspace(project, workspace)
	if err != nil {
		return nil, err
	}

	jobId, err := s.db.CreateJob(pj, ws)
	if err != nil {
		return nil, err
	}

	err = s.runner.Plan(project, ws)
	if err != nil {
		err1 := s.UpdateJobStatusPlanFailed(project, workspace, err)
		log.Printf("ERROR: %s\n", err1)

		return nil, err
	}

	return &Job{
		Id: jobId,
	}, nil
}

func (s *Service) ApproveJob(project string, workspace string) (err error) {
	ws, err := s.GetWorkspace(project, workspace)
	if err != nil {
		return err
	}

	err = s.UpdateJobStatusApplyInProgress(project, workspace)
	if err != nil {
		return err
	}

	err = s.runner.Apply(project, ws)
	if err != nil {
		err1 := s.UpdateJobStatusApplyFailed(project, workspace, err)
		log.Printf("ERROR: %s\n", err1)

		return err
	}

	return nil
}

func (s *Service) GetJobs(project string, workspace string) ([]*Job, error) {
	return s.db.GetJobs(project, workspace)
}

func (s *Service) GetJob(jobId JobId) (*Job, error) {
	return s.db.GetJob(jobId)
}

func (s *Service) UpdateJobStatusPlanInProgress(project string, workspace string) error {
	return s.db.UpdateJobStatusPlanInProgress(project, workspace)
}

func (s *Service) UpdateJobStatusReviewRequired(project string, workspace string, result string) error {
	return s.db.UpdateJobStatusReviewRequired(project, workspace, result)
}

func (s *Service) UpdateJobStatusApplyInProgress(project string, workspace string) error {
	return s.db.UpdateJobStatusApplyInProgress(project, workspace)
}

func (s *Service) UpdateJobStatusSucceeded(project string, workspace string) error {
	return s.db.UpdateJobStatusSucceeded(project, workspace)
}

func (s *Service) UpdateJobStatusPlanFailed(project string, workspace string, err error) error {
	return s.db.UpdateJobStatusPlanFailed(project, workspace, err)
}

func (s *Service) UpdateJobStatusApplyFailed(project string, workspace string, err error) error {
	return s.db.UpdateJobStatusApplyFailed(project, workspace, err)
}

func (s *Service) SavePlanLog(project string, workspace string, log string) error {
	return s.db.SavePlanLog(project, workspace, log)
}

func (s *Service) SaveApplyLog(project string, workspace string, log string) error {
	return s.db.SaveApplyLog(project, workspace, log)
}
