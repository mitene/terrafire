package server

import (
	"github.com/jinzhu/gorm"
	"github.com/mitene/terrafire/internal/api"
	"github.com/mitene/terrafire/internal/database"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func uintRef(i uint) *uint {
	return &i
}

func TestJobObserver_Check(t *testing.T) {
	errMsg := "The job runner is lost. Try again."
	footer := `
------------------------------------------------------------------------
` + errMsg

	tests := []struct {
		name           string
		records        []interface{}
		wantJobs       []*database.Job
		wantWorkspaces []*database.Workspace
	}{
		{
			name: "plan in progress",
			records: []interface{}{
				&database.Job{
					Model: gorm.Model{
						ID:        1,
						UpdatedAt: time.Now().Add(-60 * time.Second),
					},
					Project:   "pj1",
					Workspace: "ws1",
					Status:    int32(api.Job_PlanInProgress),
					PlanLog:   "plan_log_1",
					ApplyLog:  "apply_log_1",
				},
				&database.Workspace{
					Model:     gorm.Model{ID: 1},
					Project:   "pj1",
					Workspace: "ws1",
					JobId:     1,
					LastJobId: uintRef(0),
				},
			},
			wantJobs: []*database.Job{
				{
					Model:    gorm.Model{ID: 1},
					Status:   int32(api.Job_PlanFailed),
					PlanLog:  "plan_log_1" + footer,
					ApplyLog: "apply_log_1",
					Error:    errMsg,
				},
			},
			wantWorkspaces: []*database.Workspace{
				{
					Model:     gorm.Model{ID: 1},
					LastJobId: uintRef(0),
				},
			},
		},

		{
			name: "apply in progress",
			records: []interface{}{
				&database.Job{
					Model: gorm.Model{
						ID:        2,
						UpdatedAt: time.Now().Add(-60 * time.Second),
					},
					Project:   "pj2",
					Workspace: "ws2",
					Status:    int32(api.Job_ApplyInProgress),
					PlanLog:   "plan_log_2",
					ApplyLog:  "apply_log_2",
				},
				&database.Workspace{
					Model:     gorm.Model{ID: 2},
					Project:   "pj2",
					Workspace: "ws2",
					JobId:     2,
					LastJobId: uintRef(1),
				},
			},
			wantJobs: []*database.Job{
				{
					Model:    gorm.Model{ID: 2},
					Status:   int32(api.Job_ApplyFailed),
					PlanLog:  "plan_log_2",
					ApplyLog: "apply_log_2" + footer,
					Error:    errMsg,
				},
			},
			wantWorkspaces: []*database.Workspace{
				{
					Model:     gorm.Model{ID: 2},
					LastJobId: uintRef(2),
				},
			},
		},

		{
			name: "not expired",
			records: []interface{}{
				&database.Job{
					Model: gorm.Model{
						ID:        3,
						UpdatedAt: time.Now().Add(60 * time.Second),
					},
					Project:   "pj3",
					Workspace: "ws3",
					Status:    int32(api.Job_ApplyInProgress),
					PlanLog:   "plan_log_3",
					ApplyLog:  "apply_log_3",
				},
				&database.Workspace{
					Model:     gorm.Model{ID: 3},
					Project:   "pj3",
					Workspace: "ws3",
					JobId:     3,
					LastJobId: uintRef(1),
				},
			},
			wantJobs: []*database.Job{
				{
					Model:    gorm.Model{ID: 3},
					Status:   int32(api.Job_ApplyInProgress),
					PlanLog:  "plan_log_3",
					ApplyLog: "apply_log_3",
					Error:    "",
				},
			},
			wantWorkspaces: []*database.Workspace{
				{
					Model:     gorm.Model{ID: 3},
					LastJobId: uintRef(1),
				},
			},
		},

		{
			name: "not in progress",
			records: []interface{}{
				&database.Job{
					Model: gorm.Model{
						ID:        4,
						UpdatedAt: time.Now().Add(-60 * time.Second),
					},
					Project:   "pj4",
					Workspace: "ws4",
					Status:    int32(api.Job_ApplyPending),
					PlanLog:   "plan_log_4",
					ApplyLog:  "apply_log_4",
				},
				&database.Workspace{
					Model:     gorm.Model{ID: 4},
					Project:   "pj4",
					Workspace: "ws4",
					JobId:     4,
					LastJobId: uintRef(1),
				},
			},
			wantJobs: []*database.Job{
				{
					Model:    gorm.Model{ID: 4},
					Status:   int32(api.Job_ApplyPending),
					PlanLog:  "plan_log_4",
					ApplyLog: "apply_log_4",
					Error:    "",
				},
			},
			wantWorkspaces: []*database.Workspace{
				{
					Model:     gorm.Model{ID: 4},
					LastJobId: uintRef(1),
				},
			},
		},

		{
			name: "no workspace",
			records: []interface{}{
				&database.Job{
					Model: gorm.Model{
						ID:        5,
						UpdatedAt: time.Now().Add(-60 * time.Second),
					},
					Project:   "pj5",
					Workspace: "ws5",
					Status:    int32(api.Job_ApplyInProgress),
					PlanLog:   "plan_log_5",
					ApplyLog:  "apply_log_5",
				},
			},
			wantJobs: []*database.Job{
				{
					Model:    gorm.Model{ID: 5},
					Status:   int32(api.Job_ApplyInProgress),
					PlanLog:  "plan_log_5",
					ApplyLog: "apply_log_5",
					Error:    "",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := NewTestDB(t)
			db.createRecords(t, tt.records)
			s := &JobObserver{db: db}

			err := s.Check()
			assert.NoError(t, err)

			for _, j := range tt.wantJobs {
				j1 := &database.Job{}
				err := db.Take(j1, j.ID).Error
				assert.NoError(t, err)
				assert.Equal(t, j.Status, j1.Status)
				assert.Equal(t, j.PlanLog, j1.PlanLog)
				assert.Equal(t, j.ApplyLog, j1.ApplyLog)
				assert.Equal(t, j.Error, j1.Error)
			}

			for _, ws := range tt.wantWorkspaces {
				ws1 := &database.Workspace{}
				err := db.Take(ws1, ws.ID).Error
				assert.NoError(t, err)
				assert.Equal(t, *ws.LastJobId, *ws1.LastJobId)
			}
		})
	}
}
