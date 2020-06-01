package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
	"path/filepath"
)

type DB struct {
	db         *gorm.DB
	Workspaces *Workspaces
}

func NewDB(driver string, address string) (*DB, error) {
	switch driver {
	case "sqlite3":
		{
			if address != ":memory:" {
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
