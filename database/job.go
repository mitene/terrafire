package database

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/mitene/terrafire/core"
	"log"
)

type Job struct {
	gorm.Model

	Project   string `gorm:"index:idx_ws"`
	Workspace string `gorm:"index:idx_ws"`

	Status     core.JobStatus
	PlanResult string `gorm:"size:21844"`
	Error      string

	PlanLog  string `gorm:"size:21844"`
	ApplyLog string `gorm:"size:21844"`
}

type Lock struct {
	Project   string `gorm:"primary_key"`
	Workspace string `gorm:"primary_key"`
}

func (db *DB) CreateJob(project *core.Project, workspace *core.Workspace) (core.JobId, error) {
	if !db.lock(project.Name, workspace.Name) {
		return 0, fmt.Errorf("another job is running in %s/%s", project.Name, workspace.Name)
	}

	j := &Job{
		Project:   project.Name,
		Workspace: workspace.Name,
		Status:    core.JobStatusPending,
	}

	err := db.db.Create(j).Error
	if err != nil {
		db.unlock(project.Name, workspace.Name)
		return 0, err
	}

	return core.JobId(j.ID), nil
}

func (db *DB) GetJobs(project string, workspace string) ([]*core.Job, error) {
	var out []*Job
	err := db.db.Where("project = ? AND workspace = ?", project, workspace).Order("id desc").Find(&out).Error
	if err != nil {
		return nil, err
	}

	ret := make([]*core.Job, len(out))
	for i, j := range out {
		ret[i] = formatJob(j)
	}
	return ret, nil
}

func (db *DB) GetJob(jobId core.JobId) (*core.Job, error) {
	out := &Job{}
	err := db.db.Where("id = ?", jobId).First(&out).Error
	if err != nil {
		return nil, err
	}

	return formatJob(out), nil
}

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

func (db *DB) GetWorkspaceJob(project string, workspace string) (*core.Job, error) {
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

	return formatJob(j), nil
}

func (db *DB) UpdateJobStatusPlanInProgress(project string, workspace string) error {
	j, err := db.getLastJob(project, workspace)
	if err != nil {
		return err
	}
	if j == nil {
		return fmt.Errorf("no job found for %s/%s", project, workspace)
	}

	return db.db.Model(j).Update("status", core.JobStatusPlanInProgress).Error
}

func (db *DB) UpdateJobStatusReviewRequired(project string, workspace string, result string) error {
	j, err := db.getLastJob(project, workspace)
	if err != nil {
		return err
	}
	if j == nil {
		return fmt.Errorf("no job found for %s/%s", project, workspace)
	}

	err = db.db.Model(j).Updates(map[string]interface{}{
		"status":      core.JobStatusReviewRequired,
		"plan_result": result,
	}).Error
	if err != nil {
		return err
	}

	db.unlock(project, workspace)
	return nil
}

func (db *DB) UpdateJobStatusApplyInProgress(project string, workspace string) error {
	if !db.lock(project, workspace) {
		return fmt.Errorf("another job is running in %s/%s", project, workspace)
	}

	j, err := db.getLastJob(project, workspace)
	if err != nil {
		return err
	}
	if j == nil {
		return fmt.Errorf("no job found for %s/%s", project, workspace)
	}

	return db.db.Model(j).Update("status", core.JobStatusApplyInProgress).Error
}

func (db *DB) UpdateJobStatusSucceeded(project string, workspace string) error {
	j, err := db.getLastJob(project, workspace)
	if err != nil {
		return err
	}

	err = db.db.Model(j).Update("status", core.JobStatusSucceeded).Error
	if err != nil {
		return err
	}
	if j == nil {
		return fmt.Errorf("no job found for %s/%s", project, workspace)
	}

	db.unlock(project, workspace)
	return nil
}

func (db *DB) UpdateJobStatusPlanFailed(project string, workspace string, errorInfo error) error {
	j, err := db.getLastJob(project, workspace)
	if err != nil {
		return err
	}
	if j == nil {
		return fmt.Errorf("no job found for %s/%s", project, workspace)
	}

	err = db.db.Model(j).Updates(map[string]interface{}{
		"status": core.JobStatusPlanFailed,
		"error":  errorInfo.Error(),
	}).Error
	if err != nil {
		return err
	}

	db.unlock(project, workspace)
	return nil
}

func (db *DB) UpdateJobStatusApplyFailed(project string, workspace string, errorInfo error) error {
	j, err := db.getLastJob(project, workspace)
	if err != nil {
		return err
	}
	if j == nil {
		return fmt.Errorf("no job found for %s/%s", project, workspace)
	}

	err = db.db.Model(j).Updates(map[string]interface{}{
		"status": core.JobStatusApplyFailed,
		"error":  errorInfo.Error(),
	}).Error
	if err != nil {
		return err
	}

	db.unlock(project, workspace)
	return nil
}

func (db *DB) SavePlanLog(project string, workspace string, log string) error {
	j, err := db.getLastJob(project, workspace)
	if err != nil {
		return err
	}
	if j == nil {
		return fmt.Errorf("no job found for %s/%s", project, workspace)
	}

	return db.db.Model(j).Updates(map[string]interface{}{
		"plan_log": log,
	}).Error
}

func (db *DB) SaveApplyLog(project string, workspace string, log string) error {
	j, err := db.getLastJob(project, workspace)
	if err != nil {
		return err
	}
	if j == nil {
		return fmt.Errorf("no job found for %s/%s", project, workspace)
	}

	return db.db.Model(j).Updates(map[string]interface{}{
		"apply_log": log,
	}).Error
}

func (db *DB) lock(project string, workspace string) bool {
	return db.db.Create(&Lock{project, workspace}).Error == nil
}

func (db *DB) unlock(project string, workspace string) {
	err := db.db.Delete(&Lock{project, workspace}).Error
	if err != nil {
		log.Fatal("ERROR: " + err.Error())
	}
}

func formatJob(job *Job) *core.Job {
	return &core.Job{
		Id:         core.JobId(job.ID),
		StartedAt:  job.CreatedAt,
		Project:    job.Project,
		Workspace:  job.Workspace,
		Status:     job.Status,
		PlanResult: job.PlanResult,
		Error:      job.Error,
		PlanLog:    job.PlanLog,
		ApplyLog:   job.ApplyLog,
	}
}
