package terrafire

import (
	"github.com/stretchr/testify/mock"
)

type HandlerMock struct {
	mock.Mock
}

func (s *HandlerMock) GetActions() chan *Action {
	return s.Called().Get(0).(chan *Action)
}

func (s *HandlerMock) GetProjects() map[string]*Project {
	return s.Called().Get(0).(map[string]*Project)
}

func (s *HandlerMock) UpdateProjectInfo(project string, info *ProjectInfo) error {
	return s.Called(project, info).Error(0)
}

func (s *HandlerMock) RefreshProject(project string) error {
	return s.Called(project).Error(0)
}

func (s *HandlerMock) GetWorkspaces(projectName string) (map[string]*Workspace, error) {
	args := s.Called(projectName)
	return args.Get(0).(map[string]*Workspace), args.Error(1)
}

func (s *HandlerMock) GetWorkspace(project string, workspace string) (*Workspace, error) {
	args := s.Called(project, workspace)
	return args.Get(0).(*Workspace), args.Error(1)
}

func (s *HandlerMock) GetWorkspaceInfo(project string, workspace string) (*WorkspaceInfo, error) {
	args := s.Called(project, workspace)
	return args.Get(0).(*WorkspaceInfo), args.Error(1)
}

func (s *HandlerMock) SubmitJob(project string, workspace string) (*Job, error) {
	args := s.Called(project, workspace)
	return args.Get(0).(*Job), args.Error(1)
}

func (s *HandlerMock) ApproveJob(project string, workspace string) error {
	args := s.Called(project, workspace)
	return args.Error(0)
}

func (s *HandlerMock) GetJobs(project string, workspace string) ([]*Job, error) {
	args := s.Called(project, workspace)
	return args.Get(0).([]*Job), args.Error(1)
}

func (s *HandlerMock) GetJob(jobId JobId) (*Job, error) {
	args := s.Called(jobId)
	return args.Get(0).(*Job), args.Error(1)
}

func (s *HandlerMock) UpdateJobStatusPlanInProgress(project string, workspace string) error {
	return s.Called(project, workspace).Error(0)
}

func (s *HandlerMock) UpdateJobStatusReviewRequired(project string, workspace string, result string) error {
	return s.Called(project, workspace, result).Error(0)
}

func (s *HandlerMock) UpdateJobStatusApplyInProgress(project string, workspace string) error {
	return s.Called(project, workspace).Error(0)
}

func (s *HandlerMock) UpdateJobStatusSucceeded(project string, workspace string) error {
	return s.Called(project, workspace).Error(0)
}

func (s *HandlerMock) UpdateJobStatusPlanFailed(project string, workspace string, err error) error {
	return s.Called(project, workspace, err).Error(0)
}

func (s *HandlerMock) UpdateJobStatusApplyFailed(project string, workspace string, err error) error {
	return s.Called(project, workspace, err).Error(0)
}

func (s *HandlerMock) SavePlanLog(project string, workspace string, log string) error {
	return s.Called(project, workspace, log).Error(0)
}

func (s *HandlerMock) SaveApplyLog(project string, workspace string, log string) error {
	return s.Called(project, workspace, log).Error(0)
}

// Executor

type ExecutorMock struct {
	mock.Mock
}

func (m *ExecutorMock) Plan(payload *ExecutorPayload) error {
	return m.Called(payload).Error(0)
}

func (m *ExecutorMock) Apply(payload *ExecutorPayload) error {
	return m.Called(payload).Error(0)
}

// Git

type GitMock struct {
	mock.Mock
}

func (m *GitMock) Init(credentials map[string]*GitCredential) error {
	return m.Called(credentials).Error(0)
}

func (m *GitMock) Clean() error {
	return m.Called().Error(0)
}

func (m *GitMock) Fetch(dir string, repo string, branch string) (string, error) {
	args := m.Called(dir, repo, branch)
	return args.String(0), args.Error(1)
}
