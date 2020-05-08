package server

import (
	"encoding/json"
	"github.com/mitene/terrafire/core"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_GET_Projects(t *testing.T) {
	svc := &core.ServiceMock{}
	svc.On("GetProjects").Return(map[string]*core.Project{
		"dev": {
			Name:   "dev",
			Repo:   "https://github.com/foo/bar",
			Branch: "master",
			Path:   "path",
		},
	})

	s := NewServer(&core.Config{}, svc)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/projects", nil)
	rec := httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)

	resp := map[string]map[string]*core.Project{}
	err := json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, map[string]map[string]*core.Project{
		"projects": {
			"dev": {
				Name:   "dev",
				Repo:   "https://github.com/foo/bar",
				Branch: "master",
				Path:   "path",
			},
		},
	}, resp)
}

func TestServer_GET_Workspace(t *testing.T) {
	svc := &core.ServiceMock{}
	svc.On("GetWorkspaces", "dev").Return(map[string]*core.Workspace{
		"app": {
			Name: "app",
		},
	}, nil)

	s := NewServer(&core.Config{}, svc)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/projects/dev/workspaces", nil)
	rec := httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)

	resp := map[string]map[string]*core.Workspace{}
	err := json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, map[string]map[string]*core.Workspace{
		"workspaces": {
			"app": {
				Name: "app",
			},
		},
	}, resp)
}
