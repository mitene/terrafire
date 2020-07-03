package server

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/mitene/terrafire/internal/api"
	"github.com/mitene/terrafire/internal/database"
	"github.com/mitene/terrafire/internal/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"
)

func TestServer_Projects(t *testing.T) {
	type gitFetchCall struct {
		repo     string
		branch   string
		contents map[string]string
		commit   string
		err      error
	}

	tests := []struct {
		projects      map[string]*api.Project
		dbRecords     []interface{}
		gitFetchCalls []gitFetchCall

		projectName string

		wantProjectList       *api.ListProjectsResponse
		wantWorkspaceList     *api.ListWorkspacesResponse
		wantWorkspaceVersions map[string]*api.GetWorkspaceVersionResponse
	}{
		{
			projects: map[string]*api.Project{
				"pj1": {
					Name:   "pj1",
					Repo:   "pj1-repo",
					Branch: "pj1-br",
					Path:   "pj1-path",
					Envs: []*api.Pair{
						{Key: "pj1-env-k1", Value: "pj1-env-v1"},
					},
				},
			},
			dbRecords: []interface{}{
				&database.Workspace{Project: "pj1", Workspace: "ws0"},
				&database.Workspace{Project: "pj1", Workspace: "ws1", LastJobId: func() *uint { var v uint = 1000; return &v }()},
				&database.Job{
					Model:            gorm.Model{ID: 1000},
					Project:          "pj1",
					Workspace:        "ws1",
					ProjectVersion:   "pj1-commit",
					WorkspaceVersion: "ws1-commit",
				},
			},
			gitFetchCalls: []gitFetchCall{
				{
					repo:   "pj1-repo",
					branch: "pj1-br",
					contents: map[string]string{
						"pj1-path/main.hcl": `
workspace "ws1" {
  source "github" {
    owner = "ws1-owner"
    repo = "ws1-repo"
    path = "ws1-path"
    ref = "ws1-ref"
  }
}
workspace "ws2" {
  source "github" {
    owner = "ws2-owner"
    repo = "ws2-repo"
    path = "ws2-path"
    ref = "ws2-ref"
  }
}
`,
					},
					commit: "pj1-commit",
					err:    nil,
				},
			},

			projectName: "pj1",

			wantProjectList: &api.ListProjectsResponse{
				Projects: []*api.ListProjectsResponse_Project{
					{Name: "pj1"},
				},
			},
			wantWorkspaceList: &api.ListWorkspacesResponse{
				Workspaces: []*api.ListWorkspacesResponse_Workspace{
					{Name: "ws0"},
					{Name: "ws1"},
					{Name: "ws2"},
				},
			},
			wantWorkspaceVersions: map[string]*api.GetWorkspaceVersionResponse{
				"ws0": nil,
				"ws1": {
					ProjectVersion:   "pj1-commit",
					WorkspaceVersion: "ws1-commit",
				},
				"ws2": nil,
			},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			db := NewTestDB(t)
			git := utils.NewGitMock()
			srv := New(tt.projects, db, git)

			db.createRecords(t, tt.dbRecords)

			gitFetchCalls := tt.gitFetchCalls
			gitFetchCall := git.On("Fetch", mock.Anything, mock.Anything, mock.Anything)
			gitFetchCall.Run(func(args mock.Arguments) {
				c := gitFetchCalls[0]
				gitFetchCalls = gitFetchCalls[1:]

				dir := args.String(0)
				assert.Equal(t, c.repo, args.String(1))
				assert.Equal(t, c.branch, args.String(2))
				gitFetchCall.Return(c.commit, c.err)

				for name, body := range c.contents {
					p := filepath.Join(dir, name)
					err := os.MkdirAll(filepath.Dir(p), 0755)
					assert.NoError(t, err)
					err = ioutil.WriteFile(p, []byte(body), 0644)
					assert.NoError(t, err)
				}
			})

			_, err := srv.web.RefreshProject(context.Background(), &api.RefreshProjectRequest{Project: tt.projectName})
			assert.NoError(t, err)

			resp1, err := srv.web.ListProjects(context.Background(), &api.ListProjectsRequest{})
			assert.NoError(t, err)
			assert.Equal(t, tt.wantProjectList, resp1)

			resp2, err := srv.web.ListWorkspaces(context.Background(), &api.ListWorkspacesRequest{Project: tt.projectName})
			assert.NoError(t, err)
			assert.Equal(t, tt.wantWorkspaceList, resp2)

			for _, ws := range resp2.GetWorkspaces() {
				resp3, _ := srv.scheduler.GetWorkspaceVersion(context.Background(), &api.GetWorkspaceVersionRequest{
					Project:   tt.projectName,
					Workspace: ws.GetName(),
				})
				assert.Equal(t, tt.wantWorkspaceVersions[ws.GetName()], resp3)
			}
		})
	}
}

func TestServer_Actions(t *testing.T) {
	type call struct {
		name    string
		request interface {
			GetProject() string
			GetWorkspace() string
		}
		wantError            string
		wantAction           *api.GetActionResponse
		wantActionControl    *api.GetActionControlResponse
		wantJob              *api.Job
		wantWorkspaceVersion *api.GetWorkspaceVersionResponse
	}
	tests := []struct {
		projects   map[string]*api.Project
		workspaces map[string]map[string]*api.Workspace

		calls []call
	}{
		// test normal sequence
		{
			projects: map[string]*api.Project{
				"pj1": {
					Name:   "pj1",
					Repo:   "pj1-repo",
					Branch: "pj1-br",
					Path:   "pj1-path",
					Envs: []*api.Pair{
						{Key: "pj1-env-k1", Value: "pj1-env-v1"},
					},
				},
			},
			workspaces: map[string]map[string]*api.Workspace{
				"pj1": {
					"ws1": &api.Workspace{Name: "ws1"},
				},
			},

			calls: []call{
				{
					name: "submit",
					request: &api.SubmitJobRequest{
						Project:   "pj1",
						Workspace: "ws1",
					},
					wantAction: &api.GetActionResponse{
						Type:      api.GetActionResponse_SUBMIT,
						Project:   "pj1",
						Workspace: "ws1",
					},
					wantJob: &api.Job{
						Project:   "pj1",
						Workspace: &api.Workspace{Name: "ws1"},
						Status:    api.Job_Pending,
					},
				},
				{
					name: "plan in progress",
					request: &api.UpdateJobStatusRequest{
						Project:   "pj1",
						Workspace: "ws1",
						Status:    api.Job_PlanInProgress,
					},
					wantJob: &api.Job{
						Project:   "pj1",
						Workspace: &api.Workspace{Name: "ws1"},
						Status:    api.Job_PlanInProgress,
					},
				},
				{
					name: "review required",
					request: &api.UpdateJobStatusRequest{
						Project:          "pj1",
						Workspace:        "ws1",
						Status:           api.Job_ReviewRequired,
						ProjectVersion:   "pj1-commit",
						WorkspaceVersion: "ws1-commit",
					},
					wantJob: &api.Job{
						Project:          "pj1",
						Workspace:        &api.Workspace{Name: "ws1"},
						Status:           api.Job_ReviewRequired,
						ProjectVersion:   "pj1-commit",
						WorkspaceVersion: "ws1-commit",
					},
				},
				{
					name: "approve",
					request: &api.ApproveJobRequest{
						Project:   "pj1",
						Workspace: "ws1",
					},
					wantAction: &api.GetActionResponse{
						Type:      api.GetActionResponse_APPROVE,
						Project:   "pj1",
						Workspace: "ws1",
					},
					wantJob: &api.Job{
						Project:          "pj1",
						Workspace:        &api.Workspace{Name: "ws1"},
						Status:           api.Job_ApplyPending,
						ProjectVersion:   "pj1-commit",
						WorkspaceVersion: "ws1-commit",
					},
				},
				{
					name: "apply in progress",
					request: &api.UpdateJobStatusRequest{
						Project:   "pj1",
						Workspace: "ws1",
						Status:    api.Job_ApplyInProgress,
					},
					wantJob: &api.Job{
						Project:          "pj1",
						Workspace:        &api.Workspace{Name: "ws1"},
						Status:           api.Job_ApplyInProgress,
						ProjectVersion:   "pj1-commit",
						WorkspaceVersion: "ws1-commit",
					},
				},
				{
					name: "succeeded",
					request: &api.UpdateJobStatusRequest{
						Project:   "pj1",
						Workspace: "ws1",
						Status:    api.Job_Succeeded,
					},
					wantJob: &api.Job{
						Project:          "pj1",
						Workspace:        &api.Workspace{Name: "ws1"},
						Status:           api.Job_Succeeded,
						ProjectVersion:   "pj1-commit",
						WorkspaceVersion: "ws1-commit",
					},
					wantWorkspaceVersion: &api.GetWorkspaceVersionResponse{
						ProjectVersion:   "pj1-commit",
						WorkspaceVersion: "ws1-commit",
					},
				},
			},
		},

		// test invalid state transition
		{
			projects: map[string]*api.Project{
				"pj1": {
					Name:   "pj1",
					Repo:   "pj1-repo",
					Branch: "pj1-br",
					Path:   "pj1-path",
					Envs: []*api.Pair{
						{Key: "pj1-env-k1", Value: "pj1-env-v1"},
					},
				},
			},
			workspaces: map[string]map[string]*api.Workspace{
				"pj1": {
					"ws1": &api.Workspace{Name: "ws1"},
				},
			},

			calls: []call{
				{
					name: "submit",
					request: &api.SubmitJobRequest{
						Project:   "pj1",
						Workspace: "ws1",
					},
					wantAction: &api.GetActionResponse{
						Type:      api.GetActionResponse_SUBMIT,
						Project:   "pj1",
						Workspace: "ws1",
					},
					wantJob: &api.Job{
						Project:   "pj1",
						Workspace: &api.Workspace{Name: "ws1"},
						Status:    api.Job_Pending,
					},
				},
				{
					name: "plan in progress",
					request: &api.UpdateJobStatusRequest{
						Project:   "pj1",
						Workspace: "ws1",
						Status:    api.Job_PlanInProgress,
					},
					wantJob: &api.Job{
						Project:   "pj1",
						Workspace: &api.Workspace{Name: "ws1"},
						Status:    api.Job_PlanInProgress,
					},
				},
				{
					name: "plan in progress (duplicate)",
					request: &api.UpdateJobStatusRequest{
						Project:   "pj1",
						Workspace: "ws1",
						Status:    api.Job_PlanInProgress,
					},
					wantError: status.Error(codes.Aborted, "invalid state transition").Error(),
					wantJob: &api.Job{
						Project:   "pj1",
						Workspace: &api.Workspace{Name: "ws1"},
						Status:    api.Job_PlanInProgress,
					},
				},
				{
					name: "review required",
					request: &api.UpdateJobStatusRequest{
						Project:          "pj1",
						Workspace:        "ws1",
						Status:           api.Job_ReviewRequired,
						ProjectVersion:   "pj1-commit",
						WorkspaceVersion: "ws1-commit",
					},
					wantJob: &api.Job{
						Project:          "pj1",
						Workspace:        &api.Workspace{Name: "ws1"},
						Status:           api.Job_ReviewRequired,
						ProjectVersion:   "pj1-commit",
						WorkspaceVersion: "ws1-commit",
					},
				},
				{
					name: "apply in progress",
					request: &api.UpdateJobStatusRequest{
						Project:   "pj1",
						Workspace: "ws1",
						Status:    api.Job_ApplyInProgress,
					},
					wantError: status.Error(codes.Aborted, "invalid state transition").Error(),
					wantJob: &api.Job{
						Project:          "pj1",
						Workspace:        &api.Workspace{Name: "ws1"},
						Status:           api.Job_ReviewRequired,
						ProjectVersion:   "pj1-commit",
						WorkspaceVersion: "ws1-commit",
					},
				},
			},
		},

		// test normal sequence
		{
			projects: map[string]*api.Project{
				"pj1": {
					Name:   "pj1",
					Repo:   "pj1-repo",
					Branch: "pj1-br",
					Path:   "pj1-path",
					Envs: []*api.Pair{
						{Key: "pj1-env-k1", Value: "pj1-env-v1"},
					},
				},
			},
			workspaces: map[string]map[string]*api.Workspace{
				"pj1": {
					"ws1": &api.Workspace{Name: "ws1"},
				},
			},

			calls: []call{
				{
					name: "submit",
					request: &api.SubmitJobRequest{
						Project:   "pj1",
						Workspace: "ws1",
					},
					wantJob: &api.Job{
						Project:   "pj1",
						Workspace: &api.Workspace{Name: "ws1"},
						Status:    api.Job_Pending,
					},
				},
				{
					name: "cancel",
					request: &api.CancelJobRequest{
						Project:   "pj1",
						Workspace: "ws1",
					},
					wantAction: &api.GetActionResponse{
						Type: api.GetActionResponse_NONE,
					},
					wantJob: &api.Job{
						Project:   "pj1",
						Workspace: &api.Workspace{Name: "ws1"},
						Status:    api.Job_PlanFailed,
					},
				},
				{
					name: "submit (2nd)",
					request: &api.SubmitJobRequest{
						Project:   "pj1",
						Workspace: "ws1",
					},
					wantJob: &api.Job{
						Project:   "pj1",
						Workspace: &api.Workspace{Name: "ws1"},
						Status:    api.Job_Pending,
					},
				},
				{
					name: "plan in progress",
					request: &api.UpdateJobStatusRequest{
						Project:   "pj1",
						Workspace: "ws1",
						Status:    api.Job_PlanInProgress,
					},
					wantJob: &api.Job{
						Project:   "pj1",
						Workspace: &api.Workspace{Name: "ws1"},
						Status:    api.Job_PlanInProgress,
					},
				},
				{
					name: "cancel (2nd)",
					request: &api.CancelJobRequest{
						Project:   "pj1",
						Workspace: "ws1",
					},
					wantActionControl: &api.GetActionControlResponse{
						Type:      api.GetActionControlResponse_CANCEL,
						Project:   "pj1",
						Workspace: "ws1",
					},
					wantJob: &api.Job{
						Project:   "pj1",
						Workspace: &api.Workspace{Name: "ws1"},
						Status:    api.Job_PlanInProgress,
					},
				},
				{
					name: "review required",
					request: &api.UpdateJobStatusRequest{
						Project:          "pj1",
						Workspace:        "ws1",
						Status:           api.Job_ReviewRequired,
						ProjectVersion:   "pj1-commit",
						WorkspaceVersion: "ws1-commit",
					},
					wantJob: &api.Job{
						Project:          "pj1",
						Workspace:        &api.Workspace{Name: "ws1"},
						Status:           api.Job_ReviewRequired,
						ProjectVersion:   "pj1-commit",
						WorkspaceVersion: "ws1-commit",
					},
				},
				{
					name: "approve",
					request: &api.ApproveJobRequest{
						Project:   "pj1",
						Workspace: "ws1",
					},
					wantJob: &api.Job{
						Project:          "pj1",
						Workspace:        &api.Workspace{Name: "ws1"},
						Status:           api.Job_ApplyPending,
						ProjectVersion:   "pj1-commit",
						WorkspaceVersion: "ws1-commit",
					},
				},
				{
					name: "cancel (3rd)",
					request: &api.CancelJobRequest{
						Project:   "pj1",
						Workspace: "ws1",
					},
					wantAction: &api.GetActionResponse{
						Type: api.GetActionResponse_NONE,
					},
					wantJob: &api.Job{
						Project:          "pj1",
						Workspace:        &api.Workspace{Name: "ws1"},
						Status:           api.Job_ApplyFailed,
						ProjectVersion:   "pj1-commit",
						WorkspaceVersion: "ws1-commit",
					},
				},
			},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			db := NewTestDB(t)
			git := utils.NewGitMock()
			srv := New(tt.projects, db, git)

			for _, c := range tt.calls {
				pj := c.request.GetProject()
				ws := c.request.GetWorkspace()

				var err error
				switch req := c.request.(type) {
				case *api.SubmitJobRequest:
					_, err = srv.web.SubmitJob(context.Background(), req)
				case *api.ApproveJobRequest:
					_, err = srv.web.ApproveJob(context.Background(), req)
				case *api.CancelJobRequest:
					_, err = srv.web.CancelJob(context.Background(), req)
				case *api.UpdateJobStatusRequest:
					_, err = srv.scheduler.UpdateJobStatus(context.Background(), req)
				default:
					assert.FailNow(t, "invalid request type")
				}
				if c.wantError == "" {
					assert.NoError(t, err, c.name)
				} else {
					assert.EqualError(t, err, c.wantError, c.name)
				}

				if c.wantAction != nil {
					ctx1, _ := context.WithTimeout(context.Background(), 10*time.Millisecond)
					resp1, err := srv.scheduler.GetAction(ctx1, &api.GetActionRequest{})
					if st, ok := status.FromError(err); ok && st.Code() == codes.Canceled {
						resp1 = &api.GetActionResponse{Type: api.GetActionResponse_NONE}
					} else {
						assert.NoError(t, err, c.name)
					}
					assert.Equal(t, c.wantAction, resp1, c.name)
				}

				if c.wantActionControl != nil {
					ctx2, _ := context.WithTimeout(context.Background(), 10*time.Millisecond)
					resp2, err := srv.scheduler.GetActionControl(ctx2, &api.GetActionControlRequest{})
					if st, ok := status.FromError(err); ok && st.Code() == codes.Canceled {
						resp2 = &api.GetActionControlResponse{Type: api.GetActionControlResponse_NONE}
					} else {
						assert.NoError(t, err, c.name)
					}
					assert.Equal(t, c.wantActionControl, resp2, c.name)
				}

				resp2, err := srv.web.GetJob(context.Background(), &api.GetJobRequest{
					Project:   pj,
					Workspace: ws,
				})
				assert.NoError(t, err, c.name)
				// ignore Id and StartedAt
				c.wantJob.Id = resp2.Job.Id
				c.wantJob.StartedAt = resp2.Job.StartedAt
				assert.Equal(t, c.wantJob, resp2.Job, c.name)

				resp3, _ := srv.scheduler.GetWorkspaceVersion(context.Background(), &api.GetWorkspaceVersionRequest{
					Project:   pj,
					Workspace: ws,
				})
				assert.Equal(t, c.wantWorkspaceVersion, resp3, c.name)
			}
		})
	}
}
