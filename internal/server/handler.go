package server

import (
	"fmt"
	"github.com/mitene/terrafire/internal"
)

type Handler struct {
	actions  chan *internal.Action
	projects internal.ProjectRepository
	config   *internal.Config
	db       internal.DB
}

func NewHandler(config *internal.Config, db internal.DB) *Handler {
	return &Handler{
		actions:  make(chan *internal.Action, 100),
		projects: internal.ProjectRepository{},
		config:   config,
		db:       db,
	}
}

// Actions
func (s *Handler) GetActions() chan *internal.Action {
	return s.actions
}

// Projects

func (s *Handler) GetProjects() map[string]*internal.Project {
	ret := map[string]*internal.Project{}
	for name, pj := range s.projects {
		ret[name] = pj.Project
	}
	return ret
}

func (s *Handler) GetProject(project string) (*internal.Project, error) {
	pj, err := s.projects.GetProject(project)
	if err != nil {
		return nil, err
	}
	return pj.Project, nil
}

func (s *Handler) UpdateProjectInfo(project string, info *internal.ProjectInfo) error {
	s.projects[project] = info
	return nil
}

func (s *Handler) RefreshProject(project string) error {
	s.actions <- &internal.Action{
		Type:    internal.ActionTypeRefresh,
		Project: project,
	}
	return nil
}

func (s *Handler) RefreshAllProjects() error {
	s.actions <- &internal.Action{
		Type: internal.ActionTypeRefreshAll,
	}
	return nil
}

// Workspace

func (s *Handler) GetWorkspaces(project string) (map[string]*internal.Workspace, error) {
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

func (s *Handler) GetWorkspaceInfo(project string, workspace string) (*internal.WorkspaceInfo, error) {
	pj, ws, err := s.projects.GetWorkspace(project, workspace)
	if err != nil {
		return nil, err
	}

	job, err := s.db.GetJob(project, workspace)
	if err != nil {
		return nil, err
	}

	return &internal.WorkspaceInfo{
		Project:   pj,
		Workspace: ws,
		LastJob:   job,
	}, nil
}

// Job

func (s *Handler) SubmitJob(project string, workspace string) (job *internal.Job, err error) {
	pj, ws, err := s.projects.GetWorkspace(project, workspace)
	if err != nil {
		return nil, err
	}

	jobId, err := s.db.CreateJob(pj.Project, ws)
	if err != nil {
		return nil, err
	}

	s.actions <- &internal.Action{
		Type:      internal.ActionTypeSubmit,
		Project:   project,
		Workspace: workspace,
	}

	return &internal.Job{
		Id: jobId,
	}, nil
}

func (s *Handler) ApproveJob(project string, workspace string) error {
	err := s.UpdateJobStatusApplyPending(project, workspace)
	if err != nil {
		return err
	}

	s.actions <- &internal.Action{
		Type:      internal.ActionTypeApprove,
		Project:   project,
		Workspace: workspace,
	}

	return nil
}

func (s *Handler) GetJobs(project string, workspace string) ([]*internal.Job, error) {
	return s.db.GetJobs(project, workspace)
}

func (s *Handler) GetJob(jobId internal.JobId) (*internal.Job, error) {
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
