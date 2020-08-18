package server

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"
	"github.com/mitene/terrafire/internal/api"
	"github.com/mitene/terrafire/internal/database"
	"github.com/mitene/terrafire/internal/manifest"
	"github.com/mitene/terrafire/internal/utils"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"path/filepath"
	"sort"
)

type Web struct {
	projects       map[string]*api.Project
	workspaces     map[string]map[string]*api.Workspace
	actionControls chan *api.GetActionControlResponse
	db             *DB
	git            utils.Git
}

func (s *Web) RefreshProject(_ context.Context, req *api.RefreshProjectRequest) (*api.RefreshProjectResponse, error) {
	project := req.GetProject()

	log.WithFields(log.Fields{"project": project}).Info("refresh project")

	pj, ok := s.projects[project]
	if !ok {
		return nil, fmt.Errorf("project is not defined: %s", project)
	}

	dir, err := utils.TempDir()
	if err != nil {
		return nil, err
	}
	defer utils.TempClean(dir)

	_, err = s.git.Fetch(dir, pj.Repo, pj.Branch)
	if err != nil {
		return nil, err
	}

	wss, err := manifest.Load(filepath.Join(dir, pj.Path))
	if err != nil {
		return &api.RefreshProjectResponse{}, nil
	}

	m := map[string]*api.Workspace{}
	for _, ws := range wss {
		m[ws.Name] = ws
	}
	s.workspaces[project] = m

	return &api.RefreshProjectResponse{}, nil
}

func (s *Web) ListProjects(_ context.Context, _ *api.ListProjectsRequest) (*api.ListProjectsResponse, error) {
	projects := make([]*api.ListProjectsResponse_Project, 0, len(s.projects))
	for _, pj := range s.projects {
		projects = append(projects, &api.ListProjectsResponse_Project{
			Name: pj.GetName(),
		})
	}

	return &api.ListProjectsResponse{
		Projects: projects,
	}, nil
}

func (s *Web) ListWorkspaces(_ context.Context, req *api.ListWorkspacesRequest) (*api.ListWorkspacesResponse, error) {
	project := req.GetProject()

	wss, ok := s.workspaces[project]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "project %s is not found", project)
	}

	var wssDb []string
	err := s.db.Model(&database.Workspace{}).
		Select("workspace AS name").
		Where("project = ?", project).Pluck("workspace", &wssDb).Error
	if err != nil {
		return nil, err
	}

	// unique sort
	ret := make([]*api.ListWorkspacesResponse_Workspace, 0, len(wss))
	for name := range wss {
		ret = append(ret, &api.ListWorkspacesResponse_Workspace{Name: name})
	}
	for _, name := range wssDb {
		if _, ok := wss[name]; !ok {
			ret = append(ret, &api.ListWorkspacesResponse_Workspace{Name: name})
		}
	}
	sort.Slice(ret, func(i, j int) bool { return ret[i].GetName() < ret[j].GetName() })

	return &api.ListWorkspacesResponse{Workspaces: ret}, nil
}

func (s *Web) SubmitJob(_ context.Context, req *api.SubmitJobRequest) (*api.SubmitJobResponse, error) {
	project := req.GetProject()
	workspace := req.GetWorkspace()

	err := s.db.Transaction(func(tx *DB) error {
		err := tx.lock(project, workspace)
		if err != nil {
			return err
		}

		j := &database.Job{
			Project:   project,
			Workspace: workspace,
			Status:    int32(api.Job_Pending),
		}
		err = tx.Create(j).Error
		if err != nil {
			return err
		}

		ws := &database.Workspace{}
		err = tx.FirstOrCreate(ws, database.Workspace{
			Project:   project,
			Workspace: workspace,
		}).Error
		if err != nil {
			return err
		}

		err = tx.enqueue(&api.GetActionResponse{
			Type:      api.GetActionResponse_SUBMIT,
			Project:   project,
			Workspace: workspace,
		})
		if err != nil {
			return err
		}

		return tx.Model(ws).Update("job_id", j.ID).Error
	})
	if err != nil {
		return nil, err
	}

	return &api.SubmitJobResponse{}, nil
}

func (s *Web) ApproveJob(_ context.Context, req *api.ApproveJobRequest) (*api.ApproveJobResponse, error) {
	project := req.GetProject()
	workspace := req.GetWorkspace()

	err := s.db.Transaction(func(tx *DB) error {
		err := tx.lock(project, workspace)
		if err != nil {
			return err
		}

		j, err := tx.getJob(project, workspace)
		if err != nil {
			return err
		}

		err = tx.enqueue(&api.GetActionResponse{
			Type:      api.GetActionResponse_APPROVE,
			Project:   project,
			Workspace: workspace,
		})
		if err != nil {
			return err
		}

		return tx.Model(j).Update("status", api.Job_ApplyPending).Error
	})
	if err != nil {
		return nil, err
	}

	return &api.ApproveJobResponse{}, nil
}

func (s *Web) CancelJob(_ context.Context, req *api.CancelJobRequest) (*api.CancelJobResponse, error) {
	project := req.GetProject()
	workspace := req.GetWorkspace()

	err := s.db.Transaction(func(tx *DB) error {
		j, err := tx.getJob(project, workspace)
		if err != nil {
			return err
		}

		// reset job status if the job is pending.
		res1 := tx.Model(j).Where("status = ?", api.Job_Pending).Update("status", api.Job_PlanFailed)
		if res1.Error != nil {
			return res1.Error
		}
		res2 := tx.Model(j).Where("status = ?", api.Job_ApplyPending).Update("status", api.Job_ApplyFailed)
		if res2.Error != nil {
			return res2.Error
		}
		if res1.RowsAffected > 0 || res2.RowsAffected > 0 {
			err := tx.unlock(project, workspace)
			if err != nil {
				return err
			}

			err = tx.Delete(&database.Queue{}, "project = ? AND workspace = ?", project, workspace).Error
			if err != nil {
				return err
			}

			return nil
		}

		// send cancel signal if the job is in progress
		err = tx.Select("status").Take(j, j.ID).Error
		if err != nil {
			return err
		}
		if j.Status == int32(api.Job_PlanInProgress) || j.Status == int32(api.Job_ApplyInProgress) {
			s.actionControls <- &api.GetActionControlResponse{
				Type:      api.GetActionControlResponse_CANCEL,
				Project:   project,
				Workspace: workspace,
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &api.CancelJobResponse{}, nil
}

func (s *Web) GetJob(_ context.Context, req *api.GetJobRequest) (*api.GetJobResponse, error) {
	project := req.GetProject()
	workspace := req.GetWorkspace()

	j, err := s.db.getJob(project, workspace)
	if err == gorm.ErrRecordNotFound {
		return &api.GetJobResponse{Job: nil}, nil
	}
	if err != nil {
		return nil, err
	}

	err = s.db.First(j, j.ID).Error
	if err != nil {
		return nil, err
	}

	t, err := ptypes.TimestampProto(j.CreatedAt)
	if err != nil {
		return nil, err
	}

	if _, ok := api.Job_Status_name[j.Status]; !ok {
		return nil, fmt.Errorf("invalid job status code: %d", j.Status)
	}

	return &api.GetJobResponse{
		Job: &api.Job{
			Id:        uint64(j.ID),
			StartedAt: t,
			Project:   j.Project,
			Workspace: &api.Workspace{
				Name: j.Workspace,
			},
			Status:           api.Job_Status(j.Status),
			PlanResult:       j.PlanResult,
			Error:            j.Error,
			PlanLog:          j.PlanLog,
			ApplyLog:         j.ApplyLog,
			ProjectVersion:   j.ProjectVersion,
			WorkspaceVersion: j.WorkspaceVersion,
			Destroy:          j.Destroy,
		},
	}, nil
}

func (s *Web) refreshAllProject() error {
	cnt := 0
	for project := range s.projects {
		_, err := s.RefreshProject(context.Background(), &api.RefreshProjectRequest{
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
