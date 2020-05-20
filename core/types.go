package core

import "time"

type (
	Config struct {
		Address    string
		DataDir    string
		Projects   map[string]*Project
		Repos      map[string]*GitCredential
		NumWorkers int
	}

	GitCredential struct {
		Name     string
		Protocol string
		Host     string
		User     string
		Password string
	}

	Project struct {
		Name   string `json:"name"`
		Repo   string `json:"repo"`
		Branch string `json:"branch"`
		Path   string `json:"path"`
		Commit string `json:"commit,omitempty"`
		Envs   map[string]string
	}

	Manifest struct {
		Workspaces map[string]*Workspace
	}

	Workspace struct {
		Name      string            `json:"name"`
		Source    *Source           `json:"source"`
		Workspace string            `json:"workspace"`
		Vars      map[string]string `json:"vars"`
		VarFiles  []string          `json:"var_files"`
		LastJob   *Job              `json:"last_job,omitempty"`
		Project   *Project          `json:"project"`
	}

	Source struct {
		Type  string `json:"source"`
		Owner string `json:"owner"`
		Repo  string `json:"repo"`
		Path  string `json:"path"`
		Ref   string `json:"ref"`
	}

	Job struct {
		Id         JobId     `json:"id"`
		StartedAt  time.Time `json:"started_at"`
		Project    string    `json:"project"`
		Workspace  string    `json:"workspace"`
		Status     JobStatus `json:"status"`
		PlanResult string    `json:"plan_result"`
		Error      string    `json:"error"`
		PlanLog    string    `json:"plan_log"`
		ApplyLog   string    `json:"apply_log"`
	}

	JobId     uint
	JobStatus int

	ServiceProvider interface {
		GetProjects() map[string]*Project
		RefreshProject(project string) error
		GetWorkspaces(project string) (map[string]*Workspace, error)
		GetWorkspace(project string, workspace string) (*Workspace, error)
		SubmitJob(project string, workspace string) (*Job, error)
		ApproveJob(project string, workspace string) error
		GetJobs(project string, workspace string) ([]*Job, error)
		GetJob(jobId JobId) (*Job, error)
		UpdateJobStatusPlanInProgress(project string, workspace string) error
		UpdateJobStatusReviewRequired(project string, workspace string, result string) error
		UpdateJobStatusApplyInProgress(project string, workspace string) error
		UpdateJobStatusSucceeded(project string, workspace string) error
		UpdateJobStatusPlanFailed(project string, workspace string, errorInfo error) error
		UpdateJobStatusApplyFailed(project string, workspace string, errorInfo error) error
		SavePlanLog(project string, workspace string, log string) error
		SaveApplyLog(project string, workspace string, log string) error
	}

	JobRunner interface {
		Plan(project string, workspace *Workspace) error
		Apply(project string, workspace *Workspace) error
	}

	DB interface {
		CreateJob(project *Project, workspace *Workspace) (JobId, error)
		GetJobs(project string, workspace string) ([]*Job, error)
		GetJob(jobId JobId) (*Job, error)
		GetWorkspaceJob(project string, workspace string) (*Job, error)
		UpdateJobStatusPlanInProgress(project string, workspace string) error
		UpdateJobStatusReviewRequired(project string, workspace string, result string) error
		UpdateJobStatusApplyInProgress(project string, workspace string) error
		UpdateJobStatusSucceeded(project string, workspace string) error
		UpdateJobStatusPlanFailed(project string, workspace string, err error) error
		UpdateJobStatusApplyFailed(project string, workspace string, err error) error
		SavePlanLog(project string, workspace string, log string) error
		SaveApplyLog(project string, workspace string, log string) error
	}

	Git interface {
		Init(credentials map[string]*GitCredential) error
		Clean() error
		Fetch(dir string, repo string, branch string) (string, error)
	}
)

const (
	JobStatusPending JobStatus = iota
	JobStatusPlanInProgress
	JobStatusReviewRequired
	JobStatusApplyInProgress
	JobStatusSucceeded
	JobStatusPlanFailed
	JobStatusApplyFailed
)
