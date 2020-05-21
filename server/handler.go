package server

import (
	"fmt"
	"github.com/mitene/terrafire"
)

type Handler struct {
	actions  chan *terrafire.Action
	projects terrafire.ProjectRepository
	config   *terrafire.Config
	db       terrafire.DB
}

func NewHandler(config *terrafire.Config, db terrafire.DB) *Handler {
	return &Handler{
		actions:  make(chan *terrafire.Action, 100),
		projects: terrafire.ProjectRepository{},
		config:   config,
		db:       db,
	}
}

// Actions
func (s *Handler) GetActions() chan *terrafire.Action {
	return s.actions
}

// Projects

func (s *Handler) GetProjects() map[string]*terrafire.Project {
	ret := map[string]*terrafire.Project{}
	for name, pj := range s.projects {
		ret[name] = pj.Project
	}
	return ret
}

func (s *Handler) GetProject(project string) (*terrafire.Project, error) {
	pj, err := s.projects.GetProject(project)
	if err != nil {
		return nil, err
	}
	return pj.Project, nil
}

func (s *Handler) UpdateProjectInfo(project string, info *terrafire.ProjectInfo) error {
	s.projects[project] = info
	return nil
}

func (s *Handler) RefreshProject(project string) error {
	s.actions <- &terrafire.Action{
		Type:    terrafire.ActionTypeRefresh,
		Project: project,
	}
	return nil
}

func (s *Handler) RefreshAllProjects() error {
	s.actions <- &terrafire.Action{
		Type: terrafire.ActionTypeRefreshAll,
	}
	return nil
}

// Workspace

func (s *Handler) GetWorkspaces(project string) (map[string]*terrafire.Workspace, error) {
	pj, err := s.projects.GetProject(project)
	if err != nil {
		return nil, err
	}

	m := pj.Manifest
	if m == nil {
		return nil, fmt.Errorf("failed to load project manifest")
	}

	return m.Workspaces, nil
}

func (s *Handler) GetWorkspaceInfo(project string, workspace string) (*terrafire.WorkspaceInfo, error) {
	pj, ws, err := s.projects.GetWorkspace(project, workspace)
	if err != nil {
		return nil, err
	}

	job, err := s.db.GetJob(project, workspace)
	if err != nil {
		return nil, err
	}

	return &terrafire.WorkspaceInfo{
		Project:   pj,
		Workspace: ws,
		LastJob:   job,
	}, nil
}

// Job

func (s *Handler) SubmitJob(project string, workspace string) (job *terrafire.Job, err error) {
	pj, ws, err := s.projects.GetWorkspace(project, workspace)
	if err != nil {
		return nil, err
	}

	jobId, err := s.db.CreateJob(pj.Project, ws)
	if err != nil {
		return nil, err
	}

	s.actions <- &terrafire.Action{
		Type:      terrafire.ActionTypeSubmit,
		Project:   project,
		Workspace: workspace,
	}

	return &terrafire.Job{
		Id: jobId,
	}, nil
}

func (s *Handler) ApproveJob(project string, workspace string) error {
	err := s.UpdateJobStatusApplyPending(project, workspace)
	if err != nil {
		return err
	}

	s.actions <- &terrafire.Action{
		Type:      terrafire.ActionTypeApprove,
		Project:   project,
		Workspace: workspace,
	}

	return nil
}

func (s *Handler) GetJobs(project string, workspace string) ([]*terrafire.Job, error) {
	return s.db.GetJobs(project, workspace)
}

func (s *Handler) GetJob(jobId terrafire.JobId) (*terrafire.Job, error) {
	return s.db.GetJobHistory(jobId)
}

func (s *Handler) UpdateJobStatusPlanInProgress(project string, workspace string) error {
	return s.db.UpdateJobStatusPlanInProgress(project, workspace)
}

func (s *Handler) UpdateJobStatusReviewRequired(project string, workspace string, result string) error {
	return s.db.UpdateJobStatusReviewRequired(project, workspace, result)
}

func (s *Handler) UpdateJobStatusApplyPending(project string, workspace string) error {
	return s.db.UpdateJobStatusApplyPending(project, workspace)
}

func (s *Handler) UpdateJobStatusApplyInProgress(project string, workspace string) error {
	return s.db.UpdateJobStatusApplyInProgress(project, workspace)
}

func (s *Handler) UpdateJobStatusSucceeded(project string, workspace string) error {
	return s.db.UpdateJobStatusSucceeded(project, workspace)
}

func (s *Handler) UpdateJobStatusPlanFailed(project string, workspace string, err error) error {
	return s.db.UpdateJobStatusPlanFailed(project, workspace, err)
}

func (s *Handler) UpdateJobStatusApplyFailed(project string, workspace string, err error) error {
	return s.db.UpdateJobStatusApplyFailed(project, workspace, err)
}

func (s *Handler) SavePlanLog(project string, workspace string, log string) error {
	return s.db.SavePlanLog(project, workspace, log)
}

func (s *Handler) SaveApplyLog(project string, workspace string, log string) error {
	return s.db.SaveApplyLog(project, workspace, log)
}
