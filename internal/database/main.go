package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/mitene/terrafire/internal"
	"os"
	"path/filepath"
)

type DB struct {
	db         *gorm.DB
	Workspaces *Workspaces
}

func NewDB(config *internal.Config) (*DB, error) {
	var source string
	switch config.DbDriver {
	case "sqlite3":
		{
			source = config.DbSource

			if source == "" {
				source = filepath.Join(config.DataDir, "db", "sqlite3.db")
			}

			if source != ":memory:" {
				err := os.MkdirAll(filepath.Dir(source), 0755)
				if err != nil {
					return nil, err
				}
			}
		}
	default:
		return nil, fmt.Errorf("invalid db driver: %s", config.DbDriver)
	}

	db, err := gorm.Open("sqlite3", source)
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
