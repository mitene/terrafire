package runner

import (
	"github.com/mitene/terrafire/core"
	"io"
)

type (
	action int

	payload struct {
		action    action
		project   string
		workspace *core.Workspace
	}

	Terraform interface {
		Plan(dir string, workspace *core.Workspace, output io.Writer) (string, error)
		Apply(dir string, workspace *core.Workspace, output io.Writer) error
	}
)

const (
	actionPlan action = iota
	actionApply
)
