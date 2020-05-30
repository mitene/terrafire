package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mitene/terrafire/internal"
	"net/http"
	"strconv"
)

type Server struct {
	config  *internal.Config
	handler internal.Handler
	echo    *echo.Echo
}

func NewServer(config *internal.Config, handler internal.Handler) *Server {
	s := &Server{
		config:  config,
		handler: handler,
		echo:    echo.New(),
	}

	s.echo.Use(middleware.Logger())
	//s.echo.Logger.SetLevel(log.INFO)

	s.echo.GET("/api/v1/projects", s.listProjects)
	s.echo.POST("/api/v1/projects/:name/refresh", s.refreshProject)
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
	return s.echo.Start(s.config.Address)
}

func (s *Server) Stop() error {
	return s.echo.Close()
}

func (s *Server) listProjects(c echo.Context) error {
	pjs := s.handler.GetProjects()
	return c.JSON(http.StatusOK, map[string]interface{}{"projects": pjs})
}

func (s *Server) refreshProject(c echo.Context) error {
	name := c.Param("name")
	err := s.handler.RefreshProject(name)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, nil)
}

func (s *Server) listWorkspaces(c echo.Context) error {
	project := c.Param("name")

	ws, err := s.handler.GetWorkspaces(project)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"workspaces": ws})
}

func (s *Server) getWorkspace(c echo.Context) error {
	project := c.Param("project")
	workspace := c.Param("workspace")

	ws, err := s.handler.GetWorkspaceInfo(project, workspace)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, ws)
}

func (s *Server) submitJob(c echo.Context) error {
	project := c.Param("project_name")
	workspace := c.Param("workspace_name")

	job, err := s.handler.SubmitJob(project, workspace)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, job)
}

func (s *Server) approveJob(c echo.Context) error {
	project := c.Param("project_name")
	workspace := c.Param("workspace_name")

	err := s.handler.ApproveJob(project, workspace)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, nil)
}

func (s *Server) getJobs(c echo.Context) error {
	project := c.Param("project_name")
	workspace := c.Param("workspace_name")

	jobs, err := s.handler.GetJobs(project, workspace)
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

	job, err := s.handler.GetJob(internal.JobId(jobId))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, job)
}
