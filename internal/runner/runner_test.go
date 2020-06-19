package runner

import (
	"archive/zip"
	"bytes"
	"github.com/mitene/terrafire/internal/api"
	"github.com/mitene/terrafire/internal/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

func TestRunner_Plan(t *testing.T) {
	type gitFetchCall struct {
		repo   string
		ref    string
		files  map[string]string
		commit string
		err    error
	}
	type terraformCall struct {
		envs      []string
		workspace string
		vars      []string
		varFiles  map[string]string
		destroy   bool
	}
	type getWorkspaceVersionCall struct {
		req  *api.GetWorkspaceVersionRequest
		resp *api.GetWorkspaceVersionResponse
		err  error
	}
	tests := []struct {
		projects map[string]*api.Project

		argProject   string
		argWorkspace string

		wantGitFetchCalls           []*gitFetchCall
		wantTerraformCall           *terraformCall
		wantArtifactFiles           map[string]string
		wantGetWorkspaceVersionCall *getWorkspaceVersionCall
		wantUpdateJobStatusRequest  []*api.UpdateJobStatusRequest
	}{
		{
			projects: map[string]*api.Project{
				"pj1": {
					Name:   "pj1",
					Repo:   "https://github.com/pj1-owner/pj1-repo",
					Branch: "pj1-br",
					Path:   "pj1-path",
					Envs:   []*api.Pair{{Key: "pj1-env-k1", Value: "pj1-env-v1"}},
				},
			},

			argProject:   "pj1",
			argWorkspace: "ws1",

			wantGitFetchCalls: []*gitFetchCall{
				{
					repo: "https://github.com/pj1-owner/pj1-repo",
					ref:  "pj1-br",
					files: map[string]string{
						"pj1-path/main.hcl": `
workspace "ws1" {
  source "github" {
    owner = "ws1-owner"
    repo  = "ws1-repo"
    path  = "ws1-path"
    ref   = "ws1-ref"
  }
  workspace = "ws1-ws"
  vars = {
    "ws1-k1" = "ws1-v1"
  }
  var_files = [
    "ws1-vf1.tfvars",
  ]
}
`,
						"pj1-path/ws1-vf1.tfvars": "k1 = v1",
						".git/config":             "",
					},
					commit: "pj1-commit",
				},
				{
					repo: "https://github.com/ws1-owner/ws1-repo",
					ref:  "ws1-ref",
					files: map[string]string{
						"ws1-path/main.tf": "terraform config",
						".git/config":      "",
					},
					commit: "ws1-commit",
				},
			},
			wantTerraformCall: &terraformCall{
				envs:      []string{"pj1-env-k1=pj1-env-v1"},
				workspace: "ws1-ws",
				vars:      []string{"ws1-k1=ws1-v1"},
				varFiles: map[string]string{
					"ws1-vf1.tfvars": "k1 = v1",
				},
				destroy: false,
			},
			wantArtifactFiles: map[string]string{
				"main.tf":        "terraform config",
				"ws1-vf1.tfvars": "k1 = v1",
				".terrafire":     "{\"Destroy\":false}\n",
			},
			wantUpdateJobStatusRequest: []*api.UpdateJobStatusRequest{
				{Project: "pj1", Workspace: "ws1", Status: api.Job_PlanInProgress},
				{Project: "pj1", Workspace: "ws1", Status: api.Job_ReviewRequired, Result: "plan result", ProjectVersion: "pj1-commit", WorkspaceVersion: "ws1-commit"},
			},
		},

		// test for destroy
		{
			projects: map[string]*api.Project{
				"pj1": {
					Name:   "pj1",
					Repo:   "https://github.com/pj1-owner/pj1-repo",
					Branch: "pj1-br",
					Path:   "pj1-path",
				},
			},

			argProject:   "pj1",
			argWorkspace: "ws1",

			wantGitFetchCalls: []*gitFetchCall{
				{
					repo: "https://github.com/pj1-owner/pj1-repo",
					ref:  "pj1-br",
					files: map[string]string{
						"pj1-path/.gitkeep": "",
					},
					commit: "pj1-commit",
				},
				{
					repo: "https://github.com/pj1-owner/pj1-repo",
					ref:  "pj1-commit-removed",
					files: map[string]string{
						"pj1-path/main.hcl": `
workspace "ws1" {
  source "github" {
    owner = "ws1-owner"
    repo  = "ws1-repo"
    path  = "ws1-path"
    ref   = "ws1-ref"
  }
}
`,
					},
					commit: "pj1-commit-removed",
				},
				{
					repo: "https://github.com/ws1-owner/ws1-repo",
					ref:  "ws1-commit-removed",
					files: map[string]string{
						"ws1-path/main.tf": "terraform config",
					},
					commit: "ws1-commit-removed",
				},
			},
			wantTerraformCall: &terraformCall{
				envs:     []string{},
				vars:     []string{},
				varFiles: map[string]string{},
				destroy:  true,
			},
			wantArtifactFiles: map[string]string{
				"main.tf":    "terraform config",
				".terrafire": "{\"Destroy\":true}\n",
			},
			wantGetWorkspaceVersionCall: &getWorkspaceVersionCall{
				req: &api.GetWorkspaceVersionRequest{
					Project:   "pj1",
					Workspace: "ws1",
				},
				resp: &api.GetWorkspaceVersionResponse{
					ProjectVersion:   "pj1-commit-removed",
					WorkspaceVersion: "ws1-commit-removed",
				},
			},
			wantUpdateJobStatusRequest: []*api.UpdateJobStatusRequest{
				{Project: "pj1", Workspace: "ws1", Status: api.Job_PlanInProgress},
				{Project: "pj1", Workspace: "ws1", Status: api.Job_ReviewRequired, Result: "plan result",
					ProjectVersion: "pj1-commit-removed", WorkspaceVersion: "ws1-commit-removed", Destroy: true},
			},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			client := api.NewSchedulerClientMock()
			git := utils.NewGitMock()
			tf := NewTerraformMock()
			blob := NewBlobMock()
			runner := NewRunner(tt.projects, client, git, tf, blob)

			var updateJobStatusHistory []*api.UpdateJobStatusRequest
			client.On("UpdateJobStatus", mock.Anything, mock.Anything, []grpc.CallOption(nil)).
				Return(&api.UpdateJobStatusResponse{}, nil).
				Run(func(args mock.Arguments) {
					updateJobStatusHistory = append(updateJobStatusHistory, args.Get(1).(*api.UpdateJobStatusRequest))
				})

			client.On("UpdateJobLog", mock.Anything, &api.UpdateJobLogRequest{
				Project:   tt.argProject,
				Workspace: tt.argWorkspace,
				Phase:     api.Phase_Plan,
				Log:       "",
			}, []grpc.CallOption(nil)).Return(&api.UpdateJobLogResponse{}, nil)

			if c := tt.wantGetWorkspaceVersionCall; c != nil {
				client.On("GetWorkspaceVersion", mock.Anything, c.req, []grpc.CallOption(nil)).
					Return(c.resp, c.err)
			}

			for _, c := range tt.wantGitFetchCalls {
				func(c *gitFetchCall) {
					git.On("Fetch", mock.Anything, c.repo, c.ref).
						Return(c.commit, c.err).Run(func(args mock.Arguments) {
						dir := args.String(0)
						for name, body := range c.files {
							p := filepath.Join(dir, name)

							err := os.MkdirAll(filepath.Dir(p), 0755)
							assert.NoError(t, err)

							err = ioutil.WriteFile(p, []byte(body), 0644)
							assert.NoError(t, err)
						}
					})
				}(c)
			}

			tf.On("Plan", mock.Anything, tt.wantTerraformCall.workspace, mock.Anything, mock.Anything, tt.wantTerraformCall.destroy).
				Return([]byte("plan result"), nil).
				Run(func(args mock.Arguments) {
					assert.Equal(t, tt.wantTerraformCall.envs, args.Get(0).(TerraformOption).envs)
					assert.Equal(t, tt.wantTerraformCall.vars, args.Get(2).([]string))

					dir := args.Get(0).(TerraformOption).dir
					fs := map[string]string{}
					for _, varFile := range args.Get(3).([]string) {
						c, err := ioutil.ReadFile(filepath.Join(dir, varFile))
						assert.NoError(t, err)
						fs[varFile] = string(c)
					}
					assert.Equal(t, tt.wantTerraformCall.varFiles, fs)
				})

			blob.On("Put", tt.argProject, tt.argWorkspace, mock.Anything).
				Return(nil).
				Run(func(args mock.Arguments) {
					body, err := ioutil.ReadAll(args.Get(2).(io.Reader))
					assert.NoError(t, err)

					files := map[string]string{}
					r := bytes.NewReader(body)
					z, err := zip.NewReader(r, r.Size())
					assert.NoError(t, err)
					for _, f := range z.File {
						func() {
							ff, err := f.Open()
							assert.NoError(t, err)
							defer func() { _ = ff.Close() }()

							body, err := ioutil.ReadAll(ff)
							assert.NoError(t, err)

							files[f.Name] = string(body)
						}()
					}

					assert.Equal(t, tt.wantArtifactFiles, files)
				})

			err := runner.Plan(tt.argProject, tt.argWorkspace)
			assert.NoError(t, err)

			assert.Equal(t, len(tt.wantUpdateJobStatusRequest), len(updateJobStatusHistory))
			for i := range tt.wantUpdateJobStatusRequest {
				assert.Equal(t, tt.wantUpdateJobStatusRequest[i].String(), updateJobStatusHistory[i].String())
			}

			client.AssertExpectations(t)
			git.AssertExpectations(t)
			tf.AssertExpectations(t)
			blob.AssertExpectations(t)
		})
	}
}

func TestRunner_Apply(t *testing.T) {
	type terraformCall struct {
		envs    []string
		destroy bool
	}
	tests := []struct {
		projects      map[string]*api.Project
		artifactFiles map[string]string

		argProject   string
		argWorkspace string

		wantProjectName            string
		wantWorkspaceName          string
		wantTerraformCall          *terraformCall
		wantUpdateJobStatusRequest []*api.UpdateJobStatusRequest
	}{
		{
			projects: map[string]*api.Project{
				"pj1": {
					Name:   "pj1",
					Repo:   "pj1-repo",
					Branch: "pj1-br",
					Path:   "pj1-path",
					Envs:   []*api.Pair{{Key: "pj1-env-k1", Value: "pj1-env-v1"}},
				},
			},
			artifactFiles: map[string]string{
				"main.tf":    "terraform config",
				".terrafire": "{\"Destroy\":true}\n",
			},

			argProject:   "pj1",
			argWorkspace: "ws1",

			wantProjectName:   "pj1",
			wantWorkspaceName: "ws1",
			wantTerraformCall: &terraformCall{
				envs:    []string{"pj1-env-k1=pj1-env-v1"},
				destroy: true,
			},
			wantUpdateJobStatusRequest: []*api.UpdateJobStatusRequest{
				{Project: "pj1", Workspace: "ws1", Status: api.Job_ApplyInProgress},
				{Project: "pj1", Workspace: "ws1", Status: api.Job_Succeeded},
			},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			client := api.NewSchedulerClientMock()
			git := utils.NewGitMock()
			tf := NewTerraformMock()
			blob := NewBlobMock()
			runner := NewRunner(tt.projects, client, git, tf, blob)

			var updateJobStatusHistory []*api.UpdateJobStatusRequest
			client.On("UpdateJobStatus", mock.Anything, mock.Anything, []grpc.CallOption(nil)).
				Return(&api.UpdateJobStatusResponse{}, nil).
				Run(func(args mock.Arguments) {
					updateJobStatusHistory = append(updateJobStatusHistory, args.Get(1).(*api.UpdateJobStatusRequest))
				})

			client.On("UpdateJobLog", mock.Anything, &api.UpdateJobLogRequest{
				Project:   tt.wantProjectName,
				Workspace: tt.wantWorkspaceName,
				Phase:     api.Phase_Apply,
				Log:       "",
			}, []grpc.CallOption(nil)).Return(&api.UpdateJobLogResponse{}, nil)

			buf := bytes.NewBuffer(nil)
			z := zip.NewWriter(buf)
			for name, body := range tt.artifactFiles {
				w, err := z.Create(name)
				assert.NoError(t, err)

				_, err = w.Write([]byte(body))
				assert.NoError(t, err)
			}
			assert.NoError(t, z.Close())
			ar := ioutil.NopCloser(bytes.NewBuffer(buf.Bytes()))
			blob.On("Get", tt.wantProjectName, tt.wantWorkspaceName).Return(ar, nil)

			tf.On("Apply", mock.Anything, tt.wantTerraformCall.destroy).
				Return(nil).
				Run(func(args mock.Arguments) {
					opts := args.Get(0).(TerraformOption)
					assert.Equal(t, tt.wantTerraformCall.envs, opts.envs)
					for name, body := range tt.artifactFiles {
						b, err := ioutil.ReadFile(filepath.Join(opts.dir, name))
						assert.NoError(t, err)
						assert.Equal(t, body, string(b))
					}
				})

			err := runner.Apply(tt.argProject, tt.argWorkspace)
			assert.NoError(t, err)

			assert.Equal(t, len(tt.wantUpdateJobStatusRequest), len(updateJobStatusHistory))
			for i := range tt.wantUpdateJobStatusRequest {
				assert.Equal(t, tt.wantUpdateJobStatusRequest[i].String(), updateJobStatusHistory[i].String())
			}
		})
	}
}
