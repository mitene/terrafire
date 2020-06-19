package api

import (
	"context"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type SchedulerClientMock struct {
	mock.Mock
}

func NewSchedulerClientMock() *SchedulerClientMock {
	return &SchedulerClientMock{}
}

func (m *SchedulerClientMock) GetAction(ctx context.Context, in *GetActionRequest, opts ...grpc.CallOption) (*GetActionResponse, error) {
	args := m.Called(ctx, in, opts)
	return args.Get(0).(*GetActionResponse), args.Error(1)
}

func (m *SchedulerClientMock) GetActionControl(ctx context.Context, in *GetActionControlRequest, opts ...grpc.CallOption) (*GetActionControlResponse, error) {
	args := m.Called(ctx, in, opts)
	return args.Get(0).(*GetActionControlResponse), args.Error(1)
}

func (m *SchedulerClientMock) UpdateJobStatus(ctx context.Context, in *UpdateJobStatusRequest, opts ...grpc.CallOption) (*UpdateJobStatusResponse, error) {
	args := m.Called(ctx, in, opts)
	return args.Get(0).(*UpdateJobStatusResponse), args.Error(1)
}

func (m *SchedulerClientMock) UpdateJobLog(ctx context.Context, in *UpdateJobLogRequest, opts ...grpc.CallOption) (*UpdateJobLogResponse, error) {
	args := m.Called(ctx, in, opts)
	return args.Get(0).(*UpdateJobLogResponse), args.Error(1)
}

func (m *SchedulerClientMock) GetWorkspaceVersion(ctx context.Context, in *GetWorkspaceVersionRequest, opts ...grpc.CallOption) (*GetWorkspaceVersionResponse, error) {
	args := m.Called(ctx, in, opts)
	return args.Get(0).(*GetWorkspaceVersionResponse), args.Error(1)
}
