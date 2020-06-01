package database

import (
	"database/sql"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"
	"github.com/mitene/terrafire/internal/api"
	"github.com/mitene/terrafire/internal/utils"
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
}

type Lock struct {
	Project   string `gorm:"primary_key"`
	Workspace string `gorm:"primary_key"`
}

func (db *DB) CreateJob(project *api.Project, workspace *api.Workspace) (uint, error) {
	if err := db.mustLock(project.Name, workspace.Name); err != nil {
		return 0, err
	}

	j := &Job{
		Project:   project.Name,
		Workspace: workspace.Name,
		Status:    int32(api.Job_Pending),
	}

	err := db.db.Create(j).Error
	if err != nil {
		db.unlock(project.Name, workspace.Name)
		return 0, err
	}

	return j.ID, nil
}

func (db *DB) GetJobs(project string, workspace string) ([]*api.Job, error) {
	var out []*Job
	err := db.db.Where("project = ? AND workspace = ?", project, workspace).Order("id desc").Find(&out).Error
	if err != nil {
		return nil, err
	}

	ret := make([]*api.Job, len(out))
	for i, j := range out {
		ret[i], err = formatJob(j)
		if err != nil {
			return nil, err
		}
	}
	return ret, nil
}

func (db *DB) GetJob(project string, workspace string) (*api.Job, error) {
	j, err := db.getLastJob(project, workspace)
	if err != nil {
		return nil, err
	}
	if j == nil {
		return nil, nil
	}

	err = db.db.First(j).Error
	if err != nil {
		return nil, err
	}

	return formatJob(j)
}

func (db *DB) GetJobHistory(jobId uint) (*api.Job, error) {
	out := &Job{Model: gorm.Model{ID: jobId}}
	err := db.db.First(&out).Error
	if err != nil {
		return nil, err
	}

	return formatJob(out)
}

func (db *DB) UpdateJobStatusPlanInProgress(project string, workspace string) error {
	j, err := db.mustGetLastJob(project, workspace)
	if err != nil {
		return err
	}

	return db.db.Model(j).Update("status", api.Job_PlanInProgress).Error
}

func (db *DB) UpdateJobStatusReviewRequired(project string, workspace string, result string) error {
	j, err := db.mustGetLastJob(project, workspace)
	if err != nil {
		return err
	}

	err = db.db.Model(j).Updates(map[string]interface{}{
		"status":      api.Job_ReviewRequired,
		"plan_result": result,
	}).Error
	if err != nil {
		return err
	}

	db.unlock(project, workspace)
	return nil
}

func (db *DB) UpdateJobStatusApplyPending(project string, workspace string) error {
	if err := db.mustLock(project, workspace); err != nil {
		return err
	}

	j, err := db.mustGetLastJob(project, workspace)
	if err != nil {
		return err
	}

	return db.db.Model(j).Update("status", api.Job_ApplyPending).Error
}

func (db *DB) UpdateJobStatusApplyInProgress(project string, workspace string) error {
	j, err := db.mustGetLastJob(project, workspace)
	if err != nil {
		return err
	}

	return db.db.Model(j).Update("status", api.Job_ApplyInProgress).Error
}

func (db *DB) UpdateJobStatusSucceeded(project string, workspace string) error {
	db.unlock(project, workspace)

	j, err := db.mustGetLastJob(project, workspace)
	if err != nil {
		return err
	}

	err = db.db.Model(j).Update("status", api.Job_Succeeded).Error
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) UpdateJobStatusPlanFailed(project string, workspace string, errInfo string) error {
	db.unlock(project, workspace)

	j, err := db.mustGetLastJob(project, workspace)
	if err != nil {
		return err
	}

	err = db.db.Model(j).Updates(map[string]interface{}{
		"status": api.Job_PlanFailed,
		"error":  errInfo,
	}).Error
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) UpdateJobStatusApplyFailed(project string, workspace string, errInfo string) error {
	db.unlock(project, workspace)

	j, err := db.mustGetLastJob(project, workspace)
	if err != nil {
		return err
	}

	err = db.db.Model(j).Updates(map[string]interface{}{
		"status": api.Job_ApplyFailed,
		"error":  errInfo,
	}).Error
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) UpdateJobLogPlan(project string, workspace string, log string) error {
	j, err := db.mustGetLastJob(project, workspace)
	if err != nil {
		return err
	}

	return db.db.Model(j).Update("plan_log", log).Error
}

func (db *DB) UpdateJobLogApply(project string, workspace string, log string) error {
	j, err := db.mustGetLastJob(project, workspace)
	if err != nil {
		return err
	}

	return db.db.Model(j).Update("apply_log", log).Error
}

/*
Helper Functions
*/
func (db *DB) getLastJob(project string, workspace string) (*Job, error) {
	row := db.db.Model(&Job{}).Where("project = ? AND workspace = ?", project, workspace).Select("MAX(id)").Row()

	var id sql.NullInt64
	err := row.Scan(&id)
	if err != nil {
		return nil, err
	}

	if !id.Valid {
		return nil, nil
	}

	j := &Job{}
	j.ID = uint(id.Int64)
	return j, nil
}

func (db *DB) mustGetLastJob(project string, workspace string) (*Job, error) {
	j, err := db.getLastJob(project, workspace)
	if j == nil {
		return nil, fmt.Errorf("no job found for %s/%s", project, workspace)
	}
	return j, err
}

func (db *DB) lock(project string, workspace string) bool {
	return db.db.Create(&Lock{project, workspace}).Error == nil
}

func (db *DB) mustLock(project string, workspace string) error {
	if !db.lock(project, workspace) {
		return fmt.Errorf("another job is running in %s/%s", project, workspace)
	}
	return nil
}

func (db *DB) unlock(project string, workspace string) {
	utils.LogError(db.db.Delete(&Lock{project, workspace}).Error)
}

func formatJob(job *Job) (*api.Job, error) {
	t, err := ptypes.TimestampProto(job.CreatedAt)
	if err != nil {
		return nil, err
	}

	if _, ok := api.Job_Status_name[job.Status]; !ok {
		return nil, fmt.Errorf("invalid job status code: %d", job.Status)
	}

	return &api.Job{
		Id:         uint64(job.ID),
		StartedAt:  t,
		Project:    job.Project,
		Workspace:  job.Workspace,
		Status:     api.Job_Status(job.Status),
		PlanResult: job.PlanResult,
		Error:      job.Error,
		PlanLog:    job.PlanLog,
		ApplyLog:   job.ApplyLog,
	}, nil
}
