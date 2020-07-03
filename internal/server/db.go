package server

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mitene/terrafire/internal/api"
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
	db.AutoMigrate(&database.Workspace{}, &database.Job{}, &database.Lock{}, &database.Queue{})

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
	err := s.Select("job_id").Take(ws, ws).Error
	if err != nil {
		return nil, err
	}

	return &database.Job{Model: gorm.Model{ID: ws.JobId}}, nil
}

func (s *DB) updateLastJob(project string, workspace string, jobId uint) error {
	return s.Model(&database.Workspace{}).Where(&database.Workspace{
		Project:   project,
		Workspace: workspace,
	}).Update(&database.Workspace{
		LastJobId: &jobId,
	}).Error
}

func (s *DB) enqueue(m *api.GetActionResponse) error {
	return s.Create(&database.Queue{
		Project:   m.GetProject(),
		Workspace: m.GetWorkspace(),
		Action:    int32(m.GetType()),
	}).Error
}

func (s *DB) dequeue() (*api.GetActionResponse, error) {
	for {
		q := &database.Queue{}
		r := s.First(q)
		if r.RecordNotFound() {
			return nil, nil
		}
		if r.Error != nil {
			return nil, r.Error
		}

		r = s.Delete(q)
		if r.Error != nil {
			return nil, r.Error
		}
		if r.RowsAffected == 0 {
			continue
		}

		return &api.GetActionResponse{
			Type:      api.GetActionResponse_Type(q.Action),
			Project:   q.Project,
			Workspace: q.Workspace,
		}, nil
	}
}
