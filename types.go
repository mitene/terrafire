package terrafire

import (
	"time"
)

type (
	Config struct {
		Address    string
		DataDir    string
		Projects   map[string]*Project
		Repos      map[string]*GitCredential
		NumWorkers int
		DbDriver   string
		DbSource   string
	}

	GitCredential struct {
		Name     string
		Protocol string
		Host     string
		User     string
		Password string
	}

	ProjectRepository map[string]*ProjectInfo

	ProjectInfo struct {
		Project  *Project  `json:"project"`
		Manifest *Manifest `json:"-"`
		Commit   string    `json:"commit"`
		Error    string    `json:"error"`
	}

	Project struct {
		Name   string            `json:"name"`
		Repo   string            `json:"repo"`
		Branch string            `json:"branch"`
		Path   string            `json:"path"`
		Envs   map[string]string `json:"-"`
	}

	Manifest struct {
		Workspaces map[string]*Workspace
	}

	WorkspaceInfo struct {
		Project   *ProjectInfo `json:"project"`
		Workspace *Workspace   `json:"workspace"`
		LastJob   *Job         `json:"last_job,omitempty"`
	}

	Workspace struct {
		Name      string            `json:"name"`
		Source    *Source           `json:"source"`
		Workspace string            `json:"workspace"`
		Vars      map[string]string `json:"vars"`
		VarFiles  []string          `json:"var_files"`
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

	Action struct {
		Type      ActionType
		Project   string
		Workspace string
	}
	ActionType int

	Handler interface {
		GetActions() chan *Action
		GetProjects() map[string]*Project
		UpdateProjectInfo(project string, info *ProjectInfo) error
		RefreshProject(project string) error
		GetWorkspaces(project string) (map[string]*Workspace, error)
		GetWorkspaceInfo(project string, workspace string) (*WorkspaceInfo, error)
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

	Executor interface {
		Plan(payload *ExecutorPayload) error
		Apply(payload *ExecutorPayload) error
	}

	ExecutorPayload struct {
		Workspace *Workspace
		Project   *Project
	}

	Blob interface {
		New(project string, workspace string) (string, error)
		Get(project string, workspace string) (string, error)
		Put(project string, workspace string) error
	}

	DB interface {
		CreateJob(project *Project, workspace *Workspace) (JobId, error)
		GetJobs(project string, workspace string) ([]*Job, error)
		GetJob(project string, workspace string) (*Job, error)
		UpdateJobStatusPlanInProgress(project string, workspace string) error
		UpdateJobStatusReviewRequired(project string, workspace string, result string) error
		UpdateJobStatusApplyPending(project string, workspace string) error
		UpdateJobStatusApplyInProgress(project string, workspace string) error
		UpdateJobStatusSucceeded(project string, workspace string) error
		UpdateJobStatusPlanFailed(project string, workspace string, err error) error
		UpdateJobStatusApplyFailed(project string, workspace string, err error) error
		SavePlanLog(project string, workspace string, log string) error
		SaveApplyLog(project string, workspace string, log string) error
		GetJobHistory(jobId JobId) (*Job, error)
	}

	Git interface {
		Init(credentials map[string]*GitCredential) error
		Fetch(dir string, repo string, branch string) (string, error)
	}
)

const (
	JobStatusPending JobStatus = iota
	JobStatusPlanInProgress
	JobStatusReviewRequired
	JobStatusApplyPending
	JobStatusApplyInProgress
	JobStatusSucceeded
	JobStatusPlanFailed
	JobStatusApplyFailed
)

const (
	ActionTypeRefresh ActionType = iota
	ActionTypeRefreshAll
	ActionTypeSubmit
	ActionTypeApprove
)
