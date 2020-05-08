package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/mitene/terrafire/core"
	"os"
	"path/filepath"
)

type DB struct {
	db         *gorm.DB
	Workspaces *Workspaces
}

func NewDB(config *core.Config) (*DB, error) {
	path := filepath.Join(config.DataDir, "db", "sqlite3.db")

	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	// schema migration
	db.AutoMigrate(&Workspace{}, &Job{}, &Lock{})

	return &DB{
		db:         db,
		Workspaces: &Workspaces{db: db},
	}, nil
}

func (db *DB) Close() error {
	return db.db.Close()
}
