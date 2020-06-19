package server

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/mitene/terrafire/internal/api"
	"github.com/mitene/terrafire/internal/database"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Scheduler struct {
	actions        chan *api.GetActionResponse
	actionControls chan *api.GetActionControlResponse
	db             *DB
}

func (s *Scheduler) GetAction(ctx context.Context, _ *api.GetActionRequest) (*api.GetActionResponse, error) {
	select {
	case a := <-s.actions:
		return a, nil
	case <-ctx.Done():
		log.Info("connection closed")
		return nil, status.Errorf(codes.Canceled, "connection closed")
	}
}

func (s *Scheduler) GetActionControl(ctx context.Context, _ *api.GetActionControlRequest) (*api.GetActionControlResponse, error) {
	select {
	case c := <-s.actionControls:
		return c, nil
	case <-ctx.Done():
		log.Info("connection closed")
		return nil, status.Errorf(codes.Canceled, "connection closed")
	}
}

func (s *Scheduler) UpdateJobStatus(_ context.Context, req *api.UpdateJobStatusRequest) (*api.UpdateJobStatusResponse, error) {
	project := req.GetProject()
	workspace := req.GetWorkspace()

	values := map[string]interface{}{"status": req.GetStatus()}
	unlock := false
	hook := func(tx *DB, job *database.Job) error { return nil }

	switch req.GetStatus() {
	case api.Job_Pending, api.Job_ApplyPending:
		return nil, fmt.Errorf("unsupported job status: %s", req.GetStatus().String())

	case api.Job_PlanInProgress:

	case api.Job_ReviewRequired:
		values["plan_result"] = req.GetResult()
		values["project_version"] = req.GetProjectVersion()
		values["workspace_version"] = req.GetWorkspaceVersion()
		values["destroy"] = req.GetDestroy()
		unlock = true

	case api.Job_PlanFailed:
		values["error"] = req.GetError()
		unlock = true

	case api.Job_ApplyInProgress:

	case api.Job_Succeeded:
		unlock = true
		hook = func(tx *DB, job *database.Job) error {
			err := tx.Take(job, job.ID).Error
			if err != nil {
				return err
			}

			if job.Destroy {
				return tx.Unscoped().Delete(&database.Workspace{}, &database.Workspace{
					Project:   project,
					Workspace: workspace,
				}).Error
			} else {
				return tx.Model(&database.Workspace{}).Where(&database.Workspace{
					Project:   project,
					Workspace: workspace,
				}).Update(&database.Workspace{
					LastJobId: &job.ID,
				}).Error
			}
		}

	case api.Job_ApplyFailed:
		values["error"] = req.GetError()
		unlock = true
		hook = func(tx *DB, job *database.Job) error {
			return tx.Model(&database.Workspace{}).Where(&database.Workspace{
				Project:   project,
				Workspace: workspace,
			}).Update(&database.Workspace{
				LastJobId: &job.ID,
			}).Error
		}

	default:
		return nil, fmt.Errorf("unknown job status: %s", req.GetStatus().String())
	}

	err := s.db.Transaction(func(tx *DB) error {
		j, err := tx.getJob(project, workspace)
		if err != nil {
			return err
		}

		err = tx.Model(j).Updates(values).Error
		if err != nil {
			return err
		}

		err = hook(tx, j)
		if err != nil {
			return err
		}

		if unlock {
			err := tx.unlock(project, workspace)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &api.UpdateJobStatusResponse{}, nil
}

func (s *Scheduler) UpdateJobLog(_ context.Context, req *api.UpdateJobLogRequest) (*api.UpdateJobLogResponse, error) {
	project := req.GetProject()
	workspace := req.GetWorkspace()
	jobLog := req.GetLog()

	j, err := s.db.getJob(project, workspace)
	if err != nil {
		return nil, err
	}

	var attr string
	switch req.GetPhase() {
	case api.Phase_Plan:
		attr = "plan_log"
	case api.Phase_Apply:
		attr = "apply_log"
	default:
		err = fmt.Errorf("invalid job phase: %s", req.GetPhase().String())
	}

	err = s.db.Model(j).Update(attr, jobLog).Error
	if err != nil {
		return nil, err
	}

	return &api.UpdateJobLogResponse{}, nil
}

func (s *Scheduler) GetWorkspaceVersion(_ context.Context, req *api.GetWorkspaceVersionRequest) (*api.GetWorkspaceVersionResponse, error) {
	project := req.GetProject()
	workspace := req.GetWorkspace()

	ws := &database.Workspace{}
	err := s.db.First(&ws, database.Workspace{Project: project, Workspace: workspace}).Error
	if err == gorm.ErrRecordNotFound {
		return nil, status.Errorf(codes.NotFound, "workspace version is not yet registered: %s/%s", project, workspace)
	}
	if err != nil {
		return nil, err
	}

	if ws.LastJobId == nil {
		return nil, fmt.Errorf("no finished job for workspace %s/%s", project, workspace)
	}

	j := &database.Job{}
	err = s.db.First(j, *ws.LastJobId).Error
	if err != nil {
		return nil, err
	}

	return &api.GetWorkspaceVersionResponse{
		ProjectVersion:   j.ProjectVersion,
		WorkspaceVersion: j.WorkspaceVersion,
	}, nil
}
