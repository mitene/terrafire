package controller

import (
	"github.com/mitene/terrafire"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestController_RefreshProject(t *testing.T) {
	config := &terrafire.Config{
		Projects: map[string]*terrafire.Project{
			"dev": {
				Name: "dev",
			},
		},
	}
	manifest := `
workspace "app" {
  source "github" {
    owner = "foo"
    repo  = "bar"
  }
  workspace = "dev"
}
`

	dir, _ := ioutil.TempDir("", "")
	defer func() { _ = os.RemoveAll(dir) }()

	git := &terrafire.GitMock{}
	handler := &terrafire.HandlerMock{}
	executor := &terrafire.ExecutorMock{}
	ctrl := New(config, handler, executor, git, dir)

	git.On("Fetch", mock.Anything, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		err := ioutil.WriteFile(filepath.Join(args.String(0), "app.hcl"), []byte(manifest), 0644)
		assert.NoError(t, err)
	}).Return("some_commit_hash", nil)

	actions := make(chan *terrafire.Action)
	receive := make(chan mock.Arguments)

	handler.On("GetActions").Return(actions)
	handler.On("UpdateProjectInfo", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		receive <- args
	}).Return(nil)

	go func() { assert.NoError(t, ctrl.Start()) }()
	defer func() { assert.NoError(t, ctrl.Stop()) }()

	actions <- &terrafire.Action{
		Type:    terrafire.ActionTypeRefresh,
		Project: "dev",
	}

	select {
	case args := <-receive:
		{
			assert.Equal(t, "dev", args.String(0))
			assert.Equal(t, &terrafire.ProjectInfo{
				Project: config.Projects["dev"],
				Manifest: &terrafire.Manifest{
					Workspaces: map[string]*terrafire.Workspace{
						"app": {
							Name: "app",
							Source: &terrafire.Source{
								Type:  "github",
								Owner: "foo",
								Repo:  "bar",
							},
							Workspace: "dev",
						},
					},
				},
				Commit: "some_commit_hash",
				Error:  "",
			}, args.Get(1))
		}
	case <-time.After(1 * time.Second):
		assert.Fail(t, "timeout")
	}
}
