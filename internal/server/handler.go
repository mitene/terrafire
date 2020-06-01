package server

import (
	"context"
	"fmt"
	"github.com/mitene/terrafire/internal/api"
	"github.com/mitene/terrafire/internal/database"
	"github.com/mitene/terrafire/internal/manifest"
	"github.com/mitene/terrafire/internal/utils"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"path/filepath"
	"sync"
)

type Handler struct {
	actions        chan *api.GetActionResponse
	actionControls chan *api.GetActionControlResponse
	projects       map[string]*api.Project
	workspaces     map[string]map[string]*api.Workspace
	db             *database.DB
	git            utils.Git

	mux sync.Mutex
}

func NewHandler(projects map[string]*api.Project, db *database.DB, git utils.Git) *Handler {
	pjs := map[string]*api.Project{}
	for name, pj := range projects {
		pjs[name] = &(*pj) // copy project
	}

	return &Handler{
		actions:        make(chan *api.GetActionResponse, 100),
		actionControls: make(chan *api.GetActionControlResponse, 100),
		projects:       pjs,
		workspaces:     map[string]map[string]*api.Workspace{},
		db:             db,
		git:            git,

		mux: sync.Mutex{},
	}
}

/*
Projects API
*/
func (h *Handler) RefreshProject(_ context.Context, req *api.RefreshProjectRequest) (*api.RefreshProjectResponse, error) {
	project := req.GetProject()

	log.WithFields(log.Fields{"project": project}).Info("refresh project")

	pj, ok := h.projects[project]
	if !ok {
		return nil, fmt.Errorf("project is not defined: %s", project)
	}

	dir, err := utils.TempDir()
	if err != nil {
		return nil, err
	}
	defer utils.TempClean(dir)

	commit, err := h.git.Fetch(dir, pj.Repo, pj.Branch)
	if err != nil {
		return nil, err
	}

	man, err := manifest.Load(filepath.Join(dir, pj.Path))
	if err != nil {
		pj.Error = err.Error()
		return &api.RefreshProjectResponse{}, nil
	}

	h.mux.Lock()
	defer h.mux.Unlock()

	pj.Error = ""
	pj.Version = commit
	pj.Workspaces = man.Workspaces

	h.workspaces[project] = map[string]*api.Workspace{}
	for _, ws := range man.Workspaces {
		h.workspaces[project][ws.Name] = ws
	}

	return &api.RefreshProjectResponse{}, nil
}

func (h *Handler) ListProjects(_ context.Context, _ *api.ListProjectsRequest) (*api.ListProjectsResponse, error) {
	projects := make([]*api.ListProjectsResponse_Project, 0, len(h.projects))
	for _, pj := range h.projects {
		projects = append(projects, &api.ListProjectsResponse_Project{
			Name: pj.GetName(),
		})
	}

	return &api.ListProjectsResponse{
		Projects: projects,
	}, nil
}

func (h *Handler) GetProject(_ context.Context, req *api.GetProjectRequest) (*api.GetProjectResponse, error) {
	pj, ok := h.projects[req.GetProject()]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "project %s not found", pj)
	}
	return &api.GetProjectResponse{
		Project: pj,
	}, nil
}

/*
Jobs API
*/
func (h *Handler) SubmitJob(_ context.Context, req *api.SubmitJobRequest) (*api.SubmitJobResponse, error) {
	project := req.GetProject()
	workspace := req.GetWorkspace()

	pj, ok := h.projects[project]
	if !ok {
		return nil, fmt.Errorf("project %s is not defined", project)
	}

	var ws *api.Workspace
	if _, ok := h.workspaces[project]; ok {
		if _, ok := h.workspaces[project][workspace]; ok {
			ws = h.workspaces[project][workspace]
		}
	}
	if ws == nil {
		return nil, fmt.Errorf("workspace %s/%s is not defined", project, workspace)
	}

	_, err := h.db.CreateJob(pj, ws)
	if err != nil {
		return nil, err
	}

	h.actions <- &api.GetActionResponse{
		Type:      api.GetActionResponse_SUBMIT,
		Project:   project,
		Workspace: workspace,
	}

	return &api.SubmitJobResponse{}, nil
}

func (h *Handler) ApproveJob(_ context.Context, req *api.ApproveJobRequest) (*api.ApproveJobResponse, error) {
	project := req.GetProject()
	workspace := req.GetWorkspace()

	err := h.db.UpdateJobStatusApplyPending(project, workspace)
	if err != nil {
		return nil, err
	}

	h.actions <- &api.GetActionResponse{
		Type:      api.GetActionResponse_APPROVE,
		Project:   project,
		Workspace: workspace,
	}

	return &api.ApproveJobResponse{}, nil
}

func (h *Handler) GetJob(_ context.Context, req *api.GetJobRequest) (*api.GetJobResponse, error) {
	job, err := h.db.GetJob(req.GetProject(), req.GetWorkspace())
	if err != nil {
		return nil, err
	}
	return &api.GetJobResponse{
		Job: job,
	}, nil
}

func (h *Handler) GetJobs(project string, workspace string) ([]*api.Job, error) {
	return h.db.GetJobs(project, workspace)
}

func (h *Handler) GetJobHistory(jobId uint) (*api.Job, error) {
	return h.db.GetJobHistory(jobId)
}

/*
Scheduler API
*/
func (h *Handler) GetAction(ctx context.Context, req *api.GetActionRequest) (*api.GetActionResponse, error) {
	select {
	case a := <-h.actions:
		return a, nil
	case <-ctx.Done():
		log.Info("connection closed")
		return nil, status.Errorf(codes.Canceled, "connection closed")
	}
}

func (h *Handler) GetActionControl(ctx context.Context, req *api.GetActionControlRequest) (*api.GetActionControlResponse, error) {
	select {
	case c := <-h.actionControls:
		return c, nil
	case <-ctx.Done():
		log.Info("connection closed")
		return nil, status.Errorf(codes.Canceled, "connection closed")
	}
}

func (h *Handler) UpdateJobStatus(_ context.Context, req *api.UpdateJobStatusRequest) (*api.UpdateJobStatusResponse, error) {
	project := req.GetProject()
	workspace := req.GetWorkspace()
	var err error
	switch req.GetStatus() {
	case api.Job_Pending:
		err = fmt.Errorf("unsupported job status: %s", req.GetStatus().String())
	case api.Job_PlanInProgress:
		err = h.db.UpdateJobStatusPlanInProgress(project, workspace)
	case api.Job_ReviewRequired:
		err = h.db.UpdateJobStatusReviewRequired(project, workspace, req.GetResult())
	case api.Job_ApplyPending:
		err = h.db.UpdateJobStatusApplyPending(project, workspace)
	case api.Job_ApplyInProgress:
		err = h.db.UpdateJobStatusApplyInProgress(project, workspace)
	case api.Job_Succeeded:
		err = h.db.UpdateJobStatusSucceeded(project, workspace)
	case api.Job_PlanFailed:
		err = h.db.UpdateJobStatusPlanFailed(project, workspace, req.GetError())
	case api.Job_ApplyFailed:
		err = h.db.UpdateJobStatusApplyFailed(project, workspace, req.GetError())
	default:
		err = fmt.Errorf("unknown job status: %s", req.GetStatus().String())
	}
	if err != nil {
		return nil, err
	}
	return &api.UpdateJobStatusResponse{}, nil
}

func (h *Handler) UpdateJobLog(_ context.Context, req *api.UpdateJobLogRequest) (*api.UpdateJobLogResponse, error) {
	project := req.GetProject()
	workspace := req.GetWorkspace()
	jobLog := req.GetLog()

	var err error
	switch req.GetPhase() {
	case api.Phase_Plan:
		err = h.db.UpdateJobLogPlan(project, workspace, jobLog)
	case api.Phase_Apply:
		err = h.db.UpdateJobLogApply(project, workspace, jobLog)
	default:
		err = fmt.Errorf("invalid job phase: %s", req.GetPhase().String())
	}
	if err != nil {
		return nil, err
	}
	return &api.UpdateJobLogResponse{}, nil
}

/*
Helper Functions
*/

func (h *Handler) refreshAllProject() error {
	cnt := 0
	for project := range h.projects {
		_, err := h.RefreshProject(context.Background(), &api.RefreshProjectRequest{
			Project: project,
		})
		if err != nil {
			utils.LogError(err)
			cnt += 1
		}
	}

	if cnt > 0 {
		return fmt.Errorf("faield to refresh %d projects", cnt)
	}
	return nil
}
