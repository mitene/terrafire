package controller

import (
	"context"
	"fmt"
	"github.com/mitene/terrafire/internal/api"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

type Controller struct {
	client      api.SchedulerClient
	executor    Executor
	concurrency int

	ps      map[key]Process
	stopped chan struct{}
	logger  *log.Entry
}

type key struct {
	project   string
	workspace string
}

func New(client api.SchedulerClient, executor Executor, concurrency int) *Controller {
	return &Controller{
		client:      client,
		executor:    executor,
		concurrency: concurrency,

		ps:      map[key]Process{},
		stopped: make(chan struct{}),
		logger:  log.WithField("name", "controller"),
	}
}

func (c *Controller) Start() error {
	var wg sync.WaitGroup

	for i := 0; i < c.concurrency; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			c.startWorker("worker", c.handleAction)
		}(i)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		c.startWorker("control worker", c.handleActionControl)
	}()

	wg.Wait()
	return nil
}

func (c *Controller) startWorker(name string, handler func(context.Context) (func() error, error)) {
	c.logger.Infof("start %s", name)
	for {
		ctx, cancel := context.WithCancel(context.Background())
		ch := make(chan func() error)

		go func() {
			f, err := handler(ctx)
			if err != nil {
				ch <- func() error {
					time.Sleep(3 * time.Second)
					return err
				}
			} else {
				ch <- f
			}
		}()

		select {
		case f := <-ch:
			err := f()
			if err != nil {
				c.logger.Error(err)
			}
			continue

		case <-c.stopped:
			cancel()
			return
		}
	}
}

func (c *Controller) Stop() error {
	log.Info("controller is stopping")
	close(c.stopped)
	return nil
}

func (c *Controller) handleAction(ctx context.Context) (func() error, error) {
	action, err := c.client.GetAction(ctx, &api.GetActionRequest{})
	if err != nil {
		return nil, err
	}

	return func() error {
		switch action.Type {
		case api.GetActionResponse_NONE:
			return nil
		case api.GetActionResponse_SUBMIT:
			return c.SubmitJob(action.GetProject(), action.GetWorkspace())
		case api.GetActionResponse_APPROVE:
			return c.ApproveJob(action.GetProject(), action.GetWorkspace())
		default:
			return fmt.Errorf("invalid aciton type: %d", action.Type)
		}
	}, nil
}

func (c *Controller) handleActionControl(ctx context.Context) (func() error, error) {
	action, err := c.client.GetActionControl(ctx, &api.GetActionControlRequest{})
	if err != nil {
		return nil, err
	}

	return func() error {
		switch action.Type {
		case api.GetActionControlResponse_NONE:
			return nil
		case api.GetActionControlResponse_CANCEL:
			return c.CancelJob(action.GetProject(), action.GetWorkspace())
		default:
			return fmt.Errorf("invalid aciton type: %d", action.Type)
		}
	}, nil
}

func (c *Controller) SubmitJob(project string, workspace string) error {
	p, err := c.executor.Plan(project, workspace)
	if err != nil {
		return c.wrapError("plan", err)
	}

	k := key{project: project, workspace: workspace}
	c.ps[k] = p
	defer delete(c.ps, k)

	return c.wrapError("plan", p.wait())
}

func (c *Controller) ApproveJob(project string, workspace string) error {
	p, err := c.executor.Apply(project, workspace)
	if err != nil {
		return c.wrapError("apply", err)
	}

	k := key{project: project, workspace: workspace}
	c.ps[k] = p
	defer delete(c.ps, k)

	return c.wrapError("apply", p.wait())
}

func (c *Controller) CancelJob(project string, workspace string) error {
	k := key{project: project, workspace: workspace}

	p, ok := c.ps[k]
	if !ok {
		return fmt.Errorf("no job running in %s/%s", project, workspace)
	}

	return p.cancel()
}

func (*Controller) wrapError(phase string, err error) error {
	if err != nil {
		return fmt.Errorf("terrafire runner %s failed: %w", phase, err)
	}
	return nil
}
