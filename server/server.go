package server

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mitene/terrafire/core"
	"net/http"
	"strconv"
)

type Server struct {
	config  *core.Config
	service core.ServiceProvider
	echo    *echo.Echo
}

func NewServer(config *core.Config, service core.ServiceProvider) *Server {
	s := &Server{
		config:  config,
		service: service,
		echo:    echo.New(),
	}

	s.echo.Use(middleware.Logger())
	//s.echo.Logger.SetLevel(log.INFO)

	s.echo.GET("/api/v1/projects", s.listProjects)
	s.echo.GET("/api/v1/projects/:name/workspaces", s.listWorkspaces)
	s.echo.GET("/api/v1/projects/:project/workspaces/:workspace", s.getWorkspace)
	s.echo.GET("/api/v1/projects/:project_name/workspaces/:workspace_name/jobs", s.getJobs)
	s.echo.POST("/api/v1/projects/:project_name/workspaces/:workspace_name/jobs", s.submitJob)
	s.echo.POST("/api/v1/projects/:project_name/workspaces/:workspace_name/approve", s.approveJob)
	s.echo.GET("/api/v1/jobs/:job_id", s.getJob)

	s.echo.GET("/*", s.assets)

	return s
}

func (s *Server) Start() error {
	return s.echo.Start(fmt.Sprintf(":%d", s.config.Port))
}

func (s *Server) listProjects(c echo.Context) error {
	pjs := s.service.GetProjects()
	return c.JSON(http.StatusOK, map[string]interface{}{"projects": pjs})
}

func (s *Server) listWorkspaces(c echo.Context) error {
	project := c.Param("name")

	ws, err := s.service.GetWorkspaces(project)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"workspaces": ws})
}

func (s *Server) getWorkspace(c echo.Context) error {
	project := c.Param("project")
	workspace := c.Param("workspace")

	ws, err := s.service.GetWorkspace(project, workspace)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, ws)
}

func (s *Server) submitJob(c echo.Context) error {
	project := c.Param("project_name")
	workspace := c.Param("workspace_name")

	job, err := s.service.SubmitJob(project, workspace)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, job)
}

func (s *Server) approveJob(c echo.Context) error {
	project := c.Param("project_name")
	workspace := c.Param("workspace_name")

	err := s.service.ApproveJob(project, workspace)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, nil)
}

func (s *Server) getJobs(c echo.Context) error {
	project := c.Param("project_name")
	workspace := c.Param("workspace_name")

	jobs, err := s.service.GetJobs(project, workspace)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"jobs": jobs})
}

func (s *Server) getJob(c echo.Context) error {
	jobId, err := strconv.ParseUint(c.Param("job_id"), 0, 0)
	if err != nil {
		return err
	}

	job, err := s.service.GetJob(core.JobId(jobId))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, job)
}
