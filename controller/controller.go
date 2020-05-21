package controller

import (
	"fmt"
	"github.com/mitene/terrafire"
	"github.com/mitene/terrafire/utils"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"sync"
)

type Controller struct {
	config   *terrafire.Config
	handler  terrafire.Handler
	executor terrafire.Executor
	git      terrafire.Git

	projects terrafire.ProjectRepository
	done     chan struct{}
	mux      sync.Mutex
	dir      string
}

func New(config *terrafire.Config, handler terrafire.Handler, executor terrafire.Executor, git terrafire.Git, dir string) *Controller {
	return &Controller{
		config:   config,
		handler:  handler,
		executor: executor,
		git:      git,
		dir:      dir,

		projects: terrafire.ProjectRepository{},
		done:     make(chan struct{}),
		mux:      sync.Mutex{},
	}
}

func (c *Controller) Start() error {
	go func() {
		ch := c.handler.GetActions()
		for {
			select {
			case action := <-ch:
				{
					var err error

					switch action.Type {
					case terrafire.ActionTypeRefresh:
						err = c.RefreshProject(action.Project)
					case terrafire.ActionTypeRefreshAll:
						err = c.RefreshAllProjects()
					case terrafire.ActionTypeSubmit:
						err = c.SubmitJob(action.Project, action.Workspace)
					case terrafire.ActionTypeApprove:
						err = c.ApproveJob(action.Project, action.Workspace)
					default:
						utils.LogError(fmt.Errorf("invalid aciton type: %d", action.Type))
					}

					utils.LogError(err)
				}
			case _, ok := <-c.done:
				{
					if !ok {
						break
					}
				}
			}
		}
	}()

	return nil
}

func (c *Controller) Stop() error {
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
		c.projects[project] = &terrafire.ProjectInfo{}
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

	err = c.executor.Plan(&terrafire.ExecutorPayload{
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

	err = c.executor.Apply(&terrafire.ExecutorPayload{
		Project:   pj.Project,
		Workspace: ws,
	})
	if err != nil {
		utils.LogError(c.handler.UpdateJobStatusApplyFailed(project, workspace, err))
		return err
	}

	return nil
}
