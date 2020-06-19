package server

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mitene/terrafire/internal/database"
	"os"
	"path/filepath"
	"strings"
)

type DB struct {
	*gorm.DB
}

func NewDB(driver string, address string) (*DB, error) {
	switch driver {
	case "sqlite3":
		{
			if address != ":memory:" && !strings.Contains(address, "mode=memory") {
				err := os.MkdirAll(filepath.Dir(address), 0755)
				if err != nil {
					return nil, err
				}
			}
		}
	default:
		return nil, fmt.Errorf("invalid db driver: %s", driver)
	}

	db, err := gorm.Open(driver, address)
	if err != nil {
		return nil, err
	}
	//db.LogMode(true)

	// schema migration
	db.AutoMigrate(&database.Workspace{}, &database.Job{}, &database.Lock{})

	return &DB{db}, nil
}

func (s *DB) Transaction(fc func(tx *DB) error) (err error) {
	return s.DB.Transaction(func(tx *gorm.DB) error {
		return fc(&DB{tx})
	})
}

func (s *DB) lock(project string, workspace string) error {
	err := s.Create(&database.Lock{Project: project, Workspace: workspace}).Error
	if err != nil {
		return fmt.Errorf("failed to acquire lock: %w", err)
	}
	return nil
}

func (s *DB) unlock(project string, workspace string) error {
	return s.Delete(&database.Lock{Project: project, Workspace: workspace}).Error
}

func (s *DB) getJob(project string, workspace string) (*database.Job, error) {
	ws := &database.Workspace{Project: project, Workspace: workspace}
	r := s.Select("job_id").First(ws, ws)
	if r.Error != nil {
		return nil, r.Error
	}

	return &database.Job{Model: gorm.Model{ID: ws.JobId}}, nil
}
