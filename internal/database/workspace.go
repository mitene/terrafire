package database

import "github.com/jinzhu/gorm"

type Workspace struct {
	Project string `gorm:"primary_key"`
	Name    string `gorm:"primary_key"`

	SourceType  string
	SourceOwner string
	SourceRepo  string
	SourcePath  string
	SourceRef   string

	Workspace string
	Vars      string
	VarFiles  string

	Status int
}

type Workspaces struct {
	db *gorm.DB
}

func (d *Workspaces) List(project string) ([]*Workspace, error) {
	var ws []*Workspace

	err := d.db.Where("project = ?", project).Find(ws).Error
	if err != nil {
		return nil, err
	}

	return ws, nil
}

func (d *Workspace) Save(project string, workspace *Workspace) error {
	return nil
}
