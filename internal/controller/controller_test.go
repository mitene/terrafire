package controller

import (
	"github.com/mitene/terrafire/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestController_RefreshProject(t *testing.T) {
	config := &internal.Config{
		Projects: map[string]*internal.Project{
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

	git := &internal.GitMock{}
	handler := &internal.HandlerMock{}
	executor := &internal.ExecutorMock{}
	ctrl := New(config, handler, executor, git, dir)

	git.On("Fetch", mock.Anything, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		err := ioutil.WriteFile(filepath.Join(args.String(0), "app.hcl"), []byte(manifest), 0644)
		assert.NoError(t, err)
	}).Return("some_commit_hash", nil)

	actions := make(chan *internal.Action)
	receive := make(chan mock.Arguments)

	handler.On("GetActions").Return(actions)
	handler.On("UpdateProjectInfo", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		receive <- args
	}).Return(nil)

	go func() { assert.NoError(t, ctrl.Start()) }()
	defer func() { assert.NoError(t, ctrl.Stop()) }()

	actions <- &internal.Action{
		Type:    internal.ActionTypeRefresh,
		Project: "dev",
	}

	select {
	case args := <-receive:
		{
			assert.Equal(t, "dev", args.String(0))
			assert.Equal(t, &internal.ProjectInfo{
				Project: config.Projects["dev"],
				Manifest: &internal.Manifest{
					Workspaces: map[string]*internal.Workspace{
						"app": {
							Name: "app",
							Source: &internal.Source{
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
