package database

import (
	"github.com/jinzhu/gorm"
)

type Job struct {
	gorm.Model

	Project   string `gorm:"index:idx_ws"`
	Workspace string `gorm:"index:idx_ws"`

	Status     int32
	PlanResult string `gorm:"size:21844"`
	Error      string

	PlanLog  string `gorm:"size:21844"`
	ApplyLog string `gorm:"size:21844"`

	ProjectVersion   string
	WorkspaceVersion string
	Destroy          bool
}

type Lock struct {
	Project   string `gorm:"primary_key"`
	Workspace string `gorm:"primary_key"`
}

type Workspace struct {
	gorm.Model

	Project   string `gorm:"unique_index:idx_name"`
	Workspace string `gorm:"unique_index:idx_name"`

	JobId     uint
	LastJobId *uint
}
