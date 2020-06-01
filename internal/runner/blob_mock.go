package runner

import (
	"github.com/stretchr/testify/mock"
	"io"
)

type BlobMock struct {
	mock.Mock
}

func NewBlobMock() *BlobMock {
	return &BlobMock{}
}

func (m *BlobMock) Get(project string, workspace string) (io.ReadCloser, error) {
	args := m.Called(project, workspace)
	return args.Get(0).(io.ReadCloser), args.Error(1)
}

func (m *BlobMock) Put(project string, workspace string, source io.ReadSeeker) error {
	return m.Called(project, workspace, source).Error(0)
}
