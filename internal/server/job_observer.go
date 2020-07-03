package server

import (
	"github.com/mitene/terrafire/internal/api"
	"github.com/mitene/terrafire/internal/database"
	log "github.com/sirupsen/logrus"
	"time"
)

type JobObserver struct {
	db *DB
}

func (s *JobObserver) Start() {
	for {
		err := s.Check()
		if err != nil {
			log.Error(err)
		}

		time.Sleep(5 * time.Second)
	}
}

func (s *JobObserver) Check() error {
	var jobs []*database.Job
	err := s.db.Where("id IN ? AND status IN (?) AND updated_at < ?",
		s.db.Model(&database.Workspace{}).Select("job_id").SubQuery(),
		[]int32{int32(api.Job_PlanInProgress), int32(api.Job_ApplyInProgress)},
		time.Now().Add(-30*time.Second),
	).Find(&jobs).Error
	if err != nil {
		return err
	}

	for _, j := range jobs {
		j.Error = "The job runner is lost. Try again."
		var hook func(tx *DB) error
		div := `
------------------------------------------------------------------------
`
		switch j.Status {
		case int32(api.Job_PlanInProgress):
			j.PlanLog += div + j.Error
			j.Status = int32(api.Job_PlanFailed)
		case int32(api.Job_ApplyInProgress):
			j.ApplyLog += div + j.Error
			j.Status = int32(api.Job_ApplyFailed)
			hook = func(tx *DB) error {
				return tx.updateLastJob(j.Project, j.Workspace, j.ID)
			}
		}

		err = s.db.Transaction(func(tx *DB) error {
			err := tx.Save(j).Error
			if err != nil {
				return err
			}

			if hook != nil {
				err = hook(tx)
				if err != nil {
					return err
				}
			}

			err = tx.unlock(j.Project, j.Workspace)
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return err
		}
	}

	if len(jobs) > 0 {
		log.Infof("%d jobs are out of date. Canceled these jobs.", len(jobs))
	}

	return nil
}
