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
	tests := []struct {
		projects       map[string]*api.Project
		manifestFiles  map[string]string
		workspaceFiles map[string]string

		argProject   string
		argWorkspace string

		wantProject                   *api.Project
		wantWorkspace                 *api.Workspace
		wantProjectRepoUrl            string
		wantProjectRepoBranch         string
		wantWorkspaceRepoUrl          string
		wantWorkspaceRepoRef          string
		wantTerraformEnvs             []string
		wantTerraformVars             []string
		wantTerraformVarFilesContents []string
		wantArtifactFiles             map[string]string
		wantUpdateJobStatusRequest    []*api.UpdateJobStatusRequest
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
			manifestFiles: map[string]string{
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
    "ws1-vf1",
  ]
}
`,
				"pj1-path/ws1-vf1": "k1 = v1",
				".git/config":      "",
			},
			workspaceFiles: map[string]string{
				"ws1-path/main.tf": "terraform config",
				".git/config":      "",
			},

			argProject:   "pj1",
			argWorkspace: "ws1",

			wantProject: &api.Project{
				Name:   "pj1",
				Repo:   "https://github.com/pj1-owner/pj1-repo",
				Branch: "pj1-br",
				Path:   "pj1-path",
				Envs:   []*api.Pair{{Key: "pj1-env-k1", Value: "pj1-env-v1"}},
			},
			wantWorkspace: &api.Workspace{
				Name: "ws1",
				Source: &api.Source{
					Type:  api.Source_github,
					Owner: "ws1-owner",
					Repo:  "ws1-repo",
					Path:  "ws1-path",
					Ref:   "ws1-ref",
				},
				Workspace: "ws1-ws",
				Vars:      []*api.Pair{{Key: "ws1-k1", Value: "ws1-v1"}},
				VarFiles:  []string{"ws1-vf1"},
			},
			wantProjectRepoUrl:            "https://github.com/pj1-owner/pj1-repo",
			wantProjectRepoBranch:         "pj1-br",
			wantWorkspaceRepoUrl:          "https://github.com/ws1-owner/ws1-repo",
			wantWorkspaceRepoRef:          "ws1-ref",
			wantTerraformEnvs:             []string{"pj1-env-k1=pj1-env-v1"},
			wantTerraformVars:             []string{"ws1-k1=ws1-v1"},
			wantTerraformVarFilesContents: []string{"k1 = v1"},
			wantArtifactFiles: map[string]string{
				"main.tf": "terraform config",
			},
			wantUpdateJobStatusRequest: []*api.UpdateJobStatusRequest{
				{Project: "pj1", Workspace: "ws1", Status: api.Job_PlanInProgress},
				{Project: "pj1", Workspace: "ws1", Status: api.Job_ReviewRequired, Result: "plan result"},
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

			pj := tt.wantProject
			ws := tt.wantWorkspace

			var updateJobStatusHistory []*api.UpdateJobStatusRequest
			client.On("UpdateJobStatus", mock.Anything, mock.Anything, []grpc.CallOption(nil)).
				Return(&api.UpdateJobStatusResponse{}, nil).
				Run(func(args mock.Arguments) {
					updateJobStatusHistory = append(updateJobStatusHistory, args.Get(1).(*api.UpdateJobStatusRequest))
				})

			client.On("UpdateJobLog", mock.Anything, &api.UpdateJobLogRequest{
				Project:   pj.Name,
				Workspace: ws.Name,
				Phase:     api.Phase_Plan,
				Log:       "",
			}, []grpc.CallOption(nil)).Return(&api.UpdateJobLogResponse{}, nil)

			git.On("Fetch", mock.Anything, tt.wantProjectRepoUrl, tt.wantProjectRepoBranch).
				Return("pj-commit", nil).Run(func(args mock.Arguments) {
				dir := args.String(0)
				for name, body := range tt.manifestFiles {
					p := filepath.Join(dir, name)

					err := os.MkdirAll(filepath.Dir(p), 0755)
					assert.NoError(t, err)

					err = ioutil.WriteFile(p, []byte(body), 0644)
					assert.NoError(t, err)
				}
			})

			git.On("Fetch", mock.Anything, tt.wantWorkspaceRepoUrl, tt.wantWorkspaceRepoRef).
				Return("ws-commit", nil).Run(func(args mock.Arguments) {
				dir := args.String(0)
				for name, body := range tt.workspaceFiles {
					p := filepath.Join(dir, name)

					err := os.MkdirAll(filepath.Dir(p), 0755)
					assert.NoError(t, err)

					err = ioutil.WriteFile(p, []byte(body), 0644)
					assert.NoError(t, err)
				}
			})

			tf.On("Plan", mock.Anything, ws.Workspace, mock.Anything, mock.Anything).
				Return([]byte("plan result"), nil).
				Run(func(args mock.Arguments) {
					assert.Equal(t, tt.wantTerraformEnvs, args.Get(0).(TerraformOption).envs)
					assert.Equal(t, tt.wantTerraformVars, args.Get(2).([]string))

					for i, varFile := range args.Get(3).([]string) {
						c, err := ioutil.ReadFile(varFile)
						assert.NoError(t, err)
						assert.Equal(t, tt.wantTerraformVarFilesContents[i], string(c))
					}
				})

			blob.On("Put", pj.Name, ws.Name, mock.Anything).
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
			for i, _ := range tt.wantUpdateJobStatusRequest {
				assert.Equal(t, tt.wantUpdateJobStatusRequest[i].String(), updateJobStatusHistory[i].String())
			}
		})
	}
}

func TestRunner_Apply(t *testing.T) {
	tests := []struct {
		projects      map[string]*api.Project
		artifactFiles map[string]string

		argProject   string
		argWorkspace string

		wantProjectName            string
		wantWorkspaceName          string
		wantTerraformEnvs          []string
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
				"main.tf": "terraform config",
			},

			argProject:   "pj1",
			argWorkspace: "ws1",

			wantProjectName:   "pj1",
			wantWorkspaceName: "ws1",
			wantTerraformEnvs: []string{"pj1-env-k1=pj1-env-v1"},
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

			tf.On("Apply", mock.Anything).
				Return(nil).
				Run(func(args mock.Arguments) {
					opts := args.Get(0).(TerraformOption)
					assert.Equal(t, tt.wantTerraformEnvs, opts.envs)
					for name, body := range tt.artifactFiles {
						b, err := ioutil.ReadFile(filepath.Join(opts.dir, name))
						assert.NoError(t, err)
						assert.Equal(t, body, string(b))
					}
				})

			err := runner.Apply(tt.argProject, tt.argWorkspace)
			assert.NoError(t, err)

			assert.Equal(t, len(tt.wantUpdateJobStatusRequest), len(updateJobStatusHistory))
			for i, _ := range tt.wantUpdateJobStatusRequest {
				assert.Equal(t, tt.wantUpdateJobStatusRequest[i].String(), updateJobStatusHistory[i].String())
			}
		})
	}
}
