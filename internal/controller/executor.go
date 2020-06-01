package controller

type Executor interface {
	Plan(project string, workspace string) (Process, error)
	Apply(project string, workspace string) (Process, error)
}

type Process interface {
	cancel() error
	wait() error
}
