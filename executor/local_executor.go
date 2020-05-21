package executor

import (
	"github.com/mitene/terrafire"
)

type LocalExecutor struct {
	handler terrafire.Handler
	lock    chan interface{}
	runner  *Runner
}

func NewLocalExecutor(handler terrafire.Handler, runner *Runner, workerNum int) *LocalExecutor {
	if workerNum < 1 {
		workerNum = 1
	}

	return &LocalExecutor{
		handler: handler,
		lock:    make(chan interface{}, workerNum),
		runner:  runner,
	}
}

func (r *LocalExecutor) Plan(payload *terrafire.ExecutorPayload) error {
	r.withLock(func() {
		r.runner.Plan(payload)
	})
	return nil
}

func (r *LocalExecutor) Apply(payload *terrafire.ExecutorPayload) error {
	r.withLock(func() {
		r.runner.Apply(payload)
	})
	return nil
}

func (r *LocalExecutor) withLock(f func()) {
	go func() {
		r.lock <- nil
		defer func() { <-r.lock }()
		f()
	}()
}
