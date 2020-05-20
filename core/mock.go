package core

import "github.com/stretchr/testify/mock"

type ServiceMock struct {
	mock.Mock
}

func (s *ServiceMock) GetProjects() map[string]*Project {
	return s.Called().Get(0).(map[string]*Project)
}

func (s *ServiceMock) RefreshProject(project string) error {
	return s.Called(project).Error(0)
}

func (s *ServiceMock) GetWorkspaces(projectName string) (map[string]*Workspace, error) {
	args := s.Called(projectName)
	return args.Get(0).(map[string]*Workspace), args.Error(1)
}

func (s *ServiceMock) GetWorkspace(project string, workspace string) (*Workspace, error) {
	args := s.Called(project, workspace)
	return args.Get(0).(*Workspace), args.Error(1)
}

func (s *ServiceMock) SubmitJob(project string, workspace string) (*Job, error) {
	args := s.Called(project, workspace)
	return args.Get(0).(*Job), args.Error(1)
}

func (s *ServiceMock) ApproveJob(project string, workspace string) error {
	args := s.Called(project, workspace)
	return args.Error(0)
}

func (s *ServiceMock) GetJobs(project string, workspace string) ([]*Job, error) {
	args := s.Called(project, workspace)
	return args.Get(0).([]*Job), args.Error(1)
}

func (s *ServiceMock) GetJob(jobId JobId) (*Job, error) {
	args := s.Called(jobId)
	return args.Get(0).(*Job), args.Error(1)
}

func (s *ServiceMock) UpdateJobStatusPlanInProgress(project string, workspace string) error {
	return s.Called(project, workspace).Error(0)
}

func (s *ServiceMock) UpdateJobStatusReviewRequired(project string, workspace string, result string) error {
	return s.Called(project, workspace, result).Error(0)
}

func (s *ServiceMock) UpdateJobStatusApplyInProgress(project string, workspace string) error {
	return s.Called(project, workspace).Error(0)
}

func (s *ServiceMock) UpdateJobStatusSucceeded(project string, workspace string) error {
	return s.Called(project, workspace).Error(0)
}

func (s *ServiceMock) UpdateJobStatusPlanFailed(project string, workspace string, err error) error {
	return s.Called(project, workspace, err).Error(0)
}

func (s *ServiceMock) UpdateJobStatusApplyFailed(project string, workspace string, err error) error {
	return s.Called(project, workspace, err).Error(0)
}

func (s *ServiceMock) SavePlanLog(project string, workspace string, log string) error {
	return s.Called(project, workspace, log).Error(0)
}

func (s *ServiceMock) SaveApplyLog(project string, workspace string, log string) error {
	return s.Called(project, workspace, log).Error(0)
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
