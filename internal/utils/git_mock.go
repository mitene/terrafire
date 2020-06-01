package utils

import "github.com/stretchr/testify/mock"

type GitMock struct {
	mock.Mock
}

func NewGitMock() *GitMock {
	return &GitMock{}
}

func (m *GitMock) Fetch(dir string, repo string, branch string) (string, error) {
	args := m.Called(dir, repo, branch)
	return args.String(0), args.Error(1)
}
