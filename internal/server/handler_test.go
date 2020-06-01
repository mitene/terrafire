package server

import (
	"context"
	"github.com/mitene/terrafire/internal/api"
	"github.com/mitene/terrafire/internal/database"
	"github.com/mitene/terrafire/internal/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"
)

func TestHandler_Projects(t *testing.T) {
	type gitFetchCall struct {
		repo     string
		branch   string
		contents map[string]string
		commit   string
		err      error
	}

	tests := []struct {
		projects      map[string]*api.Project
		gitFetchCalls []gitFetchCall

		projectName string

		wantProjectList *api.ListProjectsResponse
		wantProject     *api.GetProjectResponse
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
			wantProject: &api.GetProjectResponse{
				Project: &api.Project{
					Name: "pj1",
					Workspaces: []*api.Workspace{
						{
							Name: "ws1",
							Source: &api.Source{
								Type:  api.Source_github,
								Owner: "ws1-owner",
								Repo:  "ws1-repo",
								Path:  "ws1-path",
								Ref:   "ws1-ref",
							},
						},
					},
					Version: "pj1-commit",
					Repo:    "pj1-repo",
					Branch:  "pj1-br",
					Path:    "pj1-path",
					Envs: []*api.Pair{
						{Key: "pj1-env-k1", Value: "pj1-env-v1"},
					},
					Error: "",
				},
			},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			db, err := database.NewDB("sqlite3", ":memory:")
			assert.NoError(t, err)
			git := utils.NewGitMock()
			handler := NewHandler(tt.projects, db, git)

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

			_, err = handler.RefreshProject(context.Background(), &api.RefreshProjectRequest{Project: tt.projectName})
			assert.NoError(t, err)

			resp1, err := handler.ListProjects(context.Background(), &api.ListProjectsRequest{})
			assert.NoError(t, err)
			assert.Equal(t, tt.wantProjectList, resp1)

			resp2, err := handler.GetProject(context.Background(), &api.GetProjectRequest{Project: tt.projectName})
			assert.NoError(t, err)
			assert.Equal(t, tt.wantProject, resp2)
		})
	}
}

func TestHandler_Actions(t *testing.T) {
	type call struct {
		request interface {
			GetProject() string
			GetWorkspace() string
		}
		wantAction *api.GetActionResponse
		wantJob    *api.Job
	}
	tests := []struct {
		projects map[string]*api.Project

		calls []call
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
					Workspaces: []*api.Workspace{
						{
							Name: "ws1",
						},
					},
				},
			},

			calls: []call{
				{
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
						Workspace: "ws1",
						Status:    api.Job_Pending,
					},
				},
				{
					request: &api.UpdateJobStatusRequest{
						Project:   "pj1",
						Workspace: "ws1",
						Status:    api.Job_ReviewRequired,
					},
					wantJob: &api.Job{
						Project:   "pj1",
						Workspace: "ws1",
						Status:    api.Job_ReviewRequired,
					},
				},
				{
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
						Project:   "pj1",
						Workspace: "ws1",
						Status:    api.Job_ApplyPending,
					},
				},
				{
					request: &api.UpdateJobStatusRequest{
						Project:   "pj1",
						Workspace: "ws1",
						Status:    api.Job_Succeeded,
					},
					wantJob: &api.Job{
						Project:   "pj1",
						Workspace: "ws1",
						Status:    api.Job_Succeeded,
					},
				},
			},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			db, err := database.NewDB("sqlite3", ":memory:")
			assert.NoError(t, err)
			git := utils.NewGitMock()
			handler := NewHandler(tt.projects, db, git)
			for k, v := range tt.projects {
				handler.workspaces[k] = map[string]*api.Workspace{}
				for _, v1 := range v.Workspaces {
					handler.workspaces[k][v1.GetName()] = v1
				}
			}

			for _, c := range tt.calls {
				pj := c.request.GetProject()
				ws := c.request.GetWorkspace()

				switch req := c.request.(type) {
				case *api.SubmitJobRequest:
					_, err = handler.SubmitJob(context.Background(), req)
				case *api.ApproveJobRequest:
					_, err = handler.ApproveJob(context.Background(), req)
				case *api.UpdateJobStatusRequest:
					_, err = handler.UpdateJobStatus(context.Background(), req)
				default:
					assert.FailNow(t, "invalid request type")
				}
				assert.NoError(t, err)

				if c.wantAction != nil {
					ctx1, _ := context.WithTimeout(context.Background(), 3*time.Second)
					resp1, err := handler.GetAction(ctx1, &api.GetActionRequest{})
					assert.NoError(t, err)
					assert.Equal(t, c.wantAction, resp1)
				}

				resp2, err := handler.GetJob(context.Background(), &api.GetJobRequest{
					Project:   pj,
					Workspace: ws,
				})
				assert.NoError(t, err)
				// ignore Id and StartedAt
				c.wantJob.Id = resp2.Job.Id
				c.wantJob.StartedAt = resp2.Job.StartedAt
				assert.Equal(t, c.wantJob, resp2.Job)
			}
		})
	}
}
