package controller

import "github.com/stretchr/testify/mock"

type ExecutorMock struct {
	mock.Mock
}

func NewExecutorMock() *ExecutorMock {
	return &ExecutorMock{}
}

func (m *ExecutorMock) Plan(project string, workspace string) (Process, error) {
	args := m.Called(project, workspace)
	return args.Get(0).(*mockProcess), args.Error(1)
}

func (m *ExecutorMock) Apply(project string, workspace string) (Process, error) {
	args := m.Called(project, workspace)
	return args.Get(0).(*mockProcess), args.Error(1)
}

func (m *ExecutorMock) Cancel(project string, workspace string) error {
	return m.Called(project, workspace).Error(0)
}

type mockProcess struct {
	mock.Mock
}

func (m *mockProcess) wait() error {
	return m.Called().Error(0)
}

func (m *mockProcess) cancel() error {
	return m.Called().Error(0)
}
