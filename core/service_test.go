package core

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestService_GetWorkspaces(t *testing.T) {
	git := &GitMock{}
	// mock git fetch
	git.On("Fetch", mock.Anything, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		err := ioutil.WriteFile(filepath.Join(args.String(0), "app.hcl"), []byte(`
workspace "app" {
  source "github" {
    owner = "foo"
    repo  = "bar"
  }
  workspace = "dev"
}
`), 0644)
		assert.NoError(t, err)
	}).Return("some_commit_hash", nil)

	svc := NewService(&Config{
		Projects: map[string]*Project{
			"dev": {
				Name: "dev",
			},
		},
	}, nil, nil, git)
	defer svc.Close()

	err := svc.Start()
	assert.NoError(t, err)

	ws, err := svc.GetWorkspaces("dev")
	assert.NoError(t, err)

	w, ok := ws["app"]
	assert.True(t, ok, "workspace app exists")
	assert.Equal(t, "foo", w.Source.Owner)
	assert.Equal(t, "bar", w.Source.Repo)
	assert.Equal(t, "dev", w.Workspace)
}
