package controller

import (
	"fmt"
	"github.com/mitene/terrafire/internal"
	"github.com/mitene/terrafire/internal/utils"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"sync"
)

type Controller struct {
	config   *internal.Config
	handler  internal.Handler
	executor internal.Executor
	git      internal.Git

	projects internal.ProjectRepository
	done     chan interface{}
	mux      sync.Mutex
	dir      string
}

func New(config *internal.Config, handler internal.Handler, executor internal.Executor, git internal.Git, dir string) *Controller {
	return &Controller{
		config:   config,
		handler:  handler,
		executor: executor,
		git:      git,
		dir:      dir,

		projects: internal.ProjectRepository{},
		done:     make(chan interface{}),
		mux:      sync.Mutex{},
	}
}

func (c *Controller) Start() error {
	ch := c.handler.GetActions()
	for {
		select {
		case action := <-ch:
			{
				var err error

				switch action.Type {
				case internal.ActionTypeRefresh:
					err = c.RefreshProject(action.Project)
				case internal.ActionTypeRefreshAll:
					err = c.RefreshAllProjects()
				case internal.ActionTypeSubmit:
					err = c.SubmitJob(action.Project, action.Workspace)
				case internal.ActionTypeApprove:
					err = c.ApproveJob(action.Project, action.Workspace)
				default:
					err = fmt.Errorf("invalid aciton type: %d", action.Type)
				}

				utils.LogError(err)
			}
		case _, ok := <-c.done:
			{
				if !ok {
					log.Error("controller stopped")
					return nil
				}
			}
		}
	}
}

func (c *Controller) Stop() error {
	log.Info("controller is stopping")
	close(c.done)
	return nil
}

func (c *Controller) RefreshAllProjects() error {
	errCount := 0
	for project := range c.config.Projects {
		err := c.RefreshProject(project)
		if err != nil {
			utils.LogError(err)
			errCount += 1
		}
	}

	if errCount > 0 {
		return fmt.Errorf("failed to refresh %d projects", errCount)
	}
	return nil
}

func (c *Controller) RefreshProject(project string) (err error) {
	c.mux.Lock()
	defer c.mux.Unlock()

	log.WithField("project", project).Info("start refresh")

	if _, ok := c.projects[project]; !ok {
		c.projects[project] = &internal.ProjectInfo{}
	}
	info, _ := c.projects[project]

	func() {
		var ok bool

		dir := filepath.Join(c.dir, "project", project)
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return
		}

		info.Project, ok = c.config.Projects[project]
		if !ok {
			err = fmt.Errorf("project is not defined: %s", project)
			return
		}

		info.Commit, err = c.git.Fetch(dir, info.Project.Repo, info.Project.Branch)
		if err != nil {
			return
		}

		info.Manifest, err = LoadManifest(filepath.Join(dir, info.Project.Path))
		if err != nil {
			return
		}
	}()
	if err != nil {
		utils.LogError(err)
		info.Error = err.Error()
	}

	err = c.handler.UpdateProjectInfo(project, info)
	if err != nil {
		return err
	}

	return nil
}

func (c *Controller) SubmitJob(project string, workspace string) error {
	pj, ws, err := c.projects.GetWorkspace(project, workspace)
	if err != nil {
		return err
	}

	err = c.executor.Plan(&internal.ExecutorPayload{
		Project:   pj.Project,
		Workspace: ws,
	})
	if err != nil {
		utils.LogError(c.handler.UpdateJobStatusPlanFailed(project, workspace, err))
		return err
	}

	return nil
}

func (c *Controller) ApproveJob(project string, workspace string) error {
	pj, ws, err := c.projects.GetWorkspace(project, workspace)
	if err != nil {
		return err
	}

	err = c.executor.Apply(&internal.ExecutorPayload{
		Project:   pj.Project,
		Workspace: ws,
	})
	if err != nil {
		utils.LogError(c.handler.UpdateJobStatusApplyFailed(project, workspace, err))
		return err
	}

	return nil
}
