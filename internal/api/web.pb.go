// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.2
// source: web.proto

package api

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ListProjectsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListProjectsRequest) Reset() {
	*x = ListProjectsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProjectsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProjectsRequest) ProtoMessage() {}

func (x *ListProjectsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProjectsRequest.ProtoReflect.Descriptor instead.
func (*ListProjectsRequest) Descriptor() ([]byte, []int) {
	return file_web_proto_rawDescGZIP(), []int{0}
}

type ListProjectsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Projects []*ListProjectsResponse_Project `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty"`
}

func (x *ListProjectsResponse) Reset() {
	*x = ListProjectsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProjectsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProjectsResponse) ProtoMessage() {}

func (x *ListProjectsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProjectsResponse.ProtoReflect.Descriptor instead.
func (*ListProjectsResponse) Descriptor() ([]byte, []int) {
	return file_web_proto_rawDescGZIP(), []int{1}
}

func (x *ListProjectsResponse) GetProjects() []*ListProjectsResponse_Project {
	if x != nil {
		return x.Projects
	}
	return nil
}

type RefreshProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
}

func (x *RefreshProjectRequest) Reset() {
	*x = RefreshProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshProjectRequest) ProtoMessage() {}

func (x *RefreshProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshProjectRequest.ProtoReflect.Descriptor instead.
func (*RefreshProjectRequest) Descriptor() ([]byte, []int) {
	return file_web_proto_rawDescGZIP(), []int{2}
}

func (x *RefreshProjectRequest) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

type RefreshProjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RefreshProjectResponse) Reset() {
	*x = RefreshProjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshProjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshProjectResponse) ProtoMessage() {}

func (x *RefreshProjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshProjectResponse.ProtoReflect.Descriptor instead.
func (*RefreshProjectResponse) Descriptor() ([]byte, []int) {
	return file_web_proto_rawDescGZIP(), []int{3}
}

type GetProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
}

func (x *GetProjectRequest) Reset() {
	*x = GetProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProjectRequest) ProtoMessage() {}

func (x *GetProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProjectRequest.ProtoReflect.Descriptor instead.
func (*GetProjectRequest) Descriptor() ([]byte, []int) {
	return file_web_proto_rawDescGZIP(), []int{4}
}

func (x *GetProjectRequest) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

type GetProjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project *Project `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
}

func (x *GetProjectResponse) Reset() {
	*x = GetProjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProjectResponse) ProtoMessage() {}

func (x *GetProjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProjectResponse.ProtoReflect.Descriptor instead.
func (*GetProjectResponse) Descriptor() ([]byte, []int) {
	return file_web_proto_rawDescGZIP(), []int{5}
}

func (x *GetProjectResponse) GetProject() *Project {
	if x != nil {
		return x.Project
	}
	return nil
}

type GetJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project   string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	Workspace string `protobuf:"bytes,2,opt,name=workspace,proto3" json:"workspace,omitempty"`
}

func (x *GetJobRequest) Reset() {
	*x = GetJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJobRequest) ProtoMessage() {}

func (x *GetJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJobRequest.ProtoReflect.Descriptor instead.
func (*GetJobRequest) Descriptor() ([]byte, []int) {
	return file_web_proto_rawDescGZIP(), []int{6}
}

func (x *GetJobRequest) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *GetJobRequest) GetWorkspace() string {
	if x != nil {
		return x.Workspace
	}
	return ""
}

type GetJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Job *Job `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
}

func (x *GetJobResponse) Reset() {
	*x = GetJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJobResponse) ProtoMessage() {}

func (x *GetJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJobResponse.ProtoReflect.Descriptor instead.
func (*GetJobResponse) Descriptor() ([]byte, []int) {
	return file_web_proto_rawDescGZIP(), []int{7}
}

func (x *GetJobResponse) GetJob() *Job {
	if x != nil {
		return x.Job
	}
	return nil
}

type SubmitJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project   string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	Workspace string `protobuf:"bytes,2,opt,name=workspace,proto3" json:"workspace,omitempty"`
}

func (x *SubmitJobRequest) Reset() {
	*x = SubmitJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitJobRequest) ProtoMessage() {}

func (x *SubmitJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitJobRequest.ProtoReflect.Descriptor instead.
func (*SubmitJobRequest) Descriptor() ([]byte, []int) {
	return file_web_proto_rawDescGZIP(), []int{8}
}

func (x *SubmitJobRequest) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *SubmitJobRequest) GetWorkspace() string {
	if x != nil {
		return x.Workspace
	}
	return ""
}

type SubmitJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubmitJobResponse) Reset() {
	*x = SubmitJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitJobResponse) ProtoMessage() {}

func (x *SubmitJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitJobResponse.ProtoReflect.Descriptor instead.
func (*SubmitJobResponse) Descriptor() ([]byte, []int) {
	return file_web_proto_rawDescGZIP(), []int{9}
}

type ApproveJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project   string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	Workspace string `protobuf:"bytes,2,opt,name=workspace,proto3" json:"workspace,omitempty"`
}

func (x *ApproveJobRequest) Reset() {
	*x = ApproveJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApproveJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApproveJobRequest) ProtoMessage() {}

func (x *ApproveJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApproveJobRequest.ProtoReflect.Descriptor instead.
func (*ApproveJobRequest) Descriptor() ([]byte, []int) {
	return file_web_proto_rawDescGZIP(), []int{10}
}

func (x *ApproveJobRequest) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *ApproveJobRequest) GetWorkspace() string {
	if x != nil {
		return x.Workspace
	}
	return ""
}

type ApproveJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ApproveJobResponse) Reset() {
	*x = ApproveJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApproveJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApproveJobResponse) ProtoMessage() {}

func (x *ApproveJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApproveJobResponse.ProtoReflect.Descriptor instead.
func (*ApproveJobResponse) Descriptor() ([]byte, []int) {
	return file_web_proto_rawDescGZIP(), []int{11}
}

type ListProjectsResponse_Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ListProjectsResponse_Project) Reset() {
	*x = ListProjectsResponse_Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProjectsResponse_Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProjectsResponse_Project) ProtoMessage() {}

func (x *ListProjectsResponse_Project) ProtoReflect() protoreflect.Message {
	mi := &file_web_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProjectsResponse_Project.ProtoReflect.Descriptor instead.
func (*ListProjectsResponse_Project) Descriptor() ([]byte, []int) {
	return file_web_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ListProjectsResponse_Project) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_web_proto protoreflect.FileDescriptor

var file_web_proto_rawDesc = []byte{
	0x0a, 0x09, 0x77, 0x65, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x15, 0x0a, 0x13, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x70, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x1a, 0x1d, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x31, 0x0a, 0x15, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x18, 0x0a, 0x16, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x2d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x38,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x47, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4a,
	0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x22, 0x28, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x04, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x22, 0x4a, 0x0a, 0x10, 0x53,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x77, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4b, 0x0a, 0x11,
	0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x77,
	0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x41, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xde, 0x02, 0x0a, 0x03, 0x57, 0x65, 0x62, 0x12, 0x3d, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x14, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x16, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x12, 0x0e,
	0x2e, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f,
	0x2e, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x34, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4a, 0x6f, 0x62, 0x12, 0x11,
	0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x41, 0x70, 0x70, 0x72, 0x6f,
	0x76, 0x65, 0x4a, 0x6f, 0x62, 0x12, 0x12, 0x2e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x4a,
	0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x41, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d,
	0x69, 0x74, 0x65, 0x6e, 0x65, 0x2f, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x69, 0x72, 0x65, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_web_proto_rawDescOnce sync.Once
	file_web_proto_rawDescData = file_web_proto_rawDesc
)

func file_web_proto_rawDescGZIP() []byte {
	file_web_proto_rawDescOnce.Do(func() {
		file_web_proto_rawDescData = protoimpl.X.CompressGZIP(file_web_proto_rawDescData)
	})
	return file_web_proto_rawDescData
}

var file_web_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_web_proto_goTypes = []interface{}{
	(*ListProjectsRequest)(nil),          // 0: ListProjectsRequest
	(*ListProjectsResponse)(nil),         // 1: ListProjectsResponse
	(*RefreshProjectRequest)(nil),        // 2: RefreshProjectRequest
	(*RefreshProjectResponse)(nil),       // 3: RefreshProjectResponse
	(*GetProjectRequest)(nil),            // 4: GetProjectRequest
	(*GetProjectResponse)(nil),           // 5: GetProjectResponse
	(*GetJobRequest)(nil),                // 6: GetJobRequest
	(*GetJobResponse)(nil),               // 7: GetJobResponse
	(*SubmitJobRequest)(nil),             // 8: SubmitJobRequest
	(*SubmitJobResponse)(nil),            // 9: SubmitJobResponse
	(*ApproveJobRequest)(nil),            // 10: ApproveJobRequest
	(*ApproveJobResponse)(nil),           // 11: ApproveJobResponse
	(*ListProjectsResponse_Project)(nil), // 12: ListProjectsResponse.Project
	(*Project)(nil),                      // 13: Project
	(*Job)(nil),                          // 14: Job
}
var file_web_proto_depIdxs = []int32{
	12, // 0: ListProjectsResponse.projects:type_name -> ListProjectsResponse.Project
	13, // 1: GetProjectResponse.project:type_name -> Project
	14, // 2: GetJobResponse.job:type_name -> Job
	0,  // 3: Web.ListProjects:input_type -> ListProjectsRequest
	2,  // 4: Web.RefreshProject:input_type -> RefreshProjectRequest
	4,  // 5: Web.GetProject:input_type -> GetProjectRequest
	6,  // 6: Web.GetJob:input_type -> GetJobRequest
	8,  // 7: Web.SubmitJob:input_type -> SubmitJobRequest
	10, // 8: Web.ApproveJob:input_type -> ApproveJobRequest
	1,  // 9: Web.ListProjects:output_type -> ListProjectsResponse
	3,  // 10: Web.RefreshProject:output_type -> RefreshProjectResponse
	5,  // 11: Web.GetProject:output_type -> GetProjectResponse
	7,  // 12: Web.GetJob:output_type -> GetJobResponse
	9,  // 13: Web.SubmitJob:output_type -> SubmitJobResponse
	11, // 14: Web.ApproveJob:output_type -> ApproveJobResponse
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_web_proto_init() }
func file_web_proto_init() {
	if File_web_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_web_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProjectsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_web_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProjectsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_web_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshProjectRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_web_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshProjectResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_web_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProjectRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_web_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProjectResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_web_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJobRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_web_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJobResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_web_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitJobRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_web_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitJobResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_web_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApproveJobRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_web_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApproveJobResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_web_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProjectsResponse_Project); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_web_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_web_proto_goTypes,
		DependencyIndexes: file_web_proto_depIdxs,
		MessageInfos:      file_web_proto_msgTypes,
	}.Build()
	File_web_proto = out.File
	file_web_proto_rawDesc = nil
	file_web_proto_goTypes = nil
	file_web_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WebClient is the client API for Web service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WebClient interface {
	ListProjects(ctx context.Context, in *ListProjectsRequest, opts ...grpc.CallOption) (*ListProjectsResponse, error)
	RefreshProject(ctx context.Context, in *RefreshProjectRequest, opts ...grpc.CallOption) (*RefreshProjectResponse, error)
	GetProject(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*GetProjectResponse, error)
	GetJob(ctx context.Context, in *GetJobRequest, opts ...grpc.CallOption) (*GetJobResponse, error)
	SubmitJob(ctx context.Context, in *SubmitJobRequest, opts ...grpc.CallOption) (*SubmitJobResponse, error)
	ApproveJob(ctx context.Context, in *ApproveJobRequest, opts ...grpc.CallOption) (*ApproveJobResponse, error)
}

type webClient struct {
	cc grpc.ClientConnInterface
}

func NewWebClient(cc grpc.ClientConnInterface) WebClient {
	return &webClient{cc}
}

func (c *webClient) ListProjects(ctx context.Context, in *ListProjectsRequest, opts ...grpc.CallOption) (*ListProjectsResponse, error) {
	out := new(ListProjectsResponse)
	err := c.cc.Invoke(ctx, "/Web/ListProjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webClient) RefreshProject(ctx context.Context, in *RefreshProjectRequest, opts ...grpc.CallOption) (*RefreshProjectResponse, error) {
	out := new(RefreshProjectResponse)
	err := c.cc.Invoke(ctx, "/Web/RefreshProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webClient) GetProject(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*GetProjectResponse, error) {
	out := new(GetProjectResponse)
	err := c.cc.Invoke(ctx, "/Web/GetProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webClient) GetJob(ctx context.Context, in *GetJobRequest, opts ...grpc.CallOption) (*GetJobResponse, error) {
	out := new(GetJobResponse)
	err := c.cc.Invoke(ctx, "/Web/GetJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webClient) SubmitJob(ctx context.Context, in *SubmitJobRequest, opts ...grpc.CallOption) (*SubmitJobResponse, error) {
	out := new(SubmitJobResponse)
	err := c.cc.Invoke(ctx, "/Web/SubmitJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webClient) ApproveJob(ctx context.Context, in *ApproveJobRequest, opts ...grpc.CallOption) (*ApproveJobResponse, error) {
	out := new(ApproveJobResponse)
	err := c.cc.Invoke(ctx, "/Web/ApproveJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebServer is the server API for Web service.
type WebServer interface {
	ListProjects(context.Context, *ListProjectsRequest) (*ListProjectsResponse, error)
	RefreshProject(context.Context, *RefreshProjectRequest) (*RefreshProjectResponse, error)
	GetProject(context.Context, *GetProjectRequest) (*GetProjectResponse, error)
	GetJob(context.Context, *GetJobRequest) (*GetJobResponse, error)
	SubmitJob(context.Context, *SubmitJobRequest) (*SubmitJobResponse, error)
	ApproveJob(context.Context, *ApproveJobRequest) (*ApproveJobResponse, error)
}

// UnimplementedWebServer can be embedded to have forward compatible implementations.
type UnimplementedWebServer struct {
}

func (*UnimplementedWebServer) ListProjects(context.Context, *ListProjectsRequest) (*ListProjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProjects not implemented")
}
func (*UnimplementedWebServer) RefreshProject(context.Context, *RefreshProjectRequest) (*RefreshProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshProject not implemented")
}
func (*UnimplementedWebServer) GetProject(context.Context, *GetProjectRequest) (*GetProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProject not implemented")
}
func (*UnimplementedWebServer) GetJob(context.Context, *GetJobRequest) (*GetJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJob not implemented")
}
func (*UnimplementedWebServer) SubmitJob(context.Context, *SubmitJobRequest) (*SubmitJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitJob not implemented")
}
func (*UnimplementedWebServer) ApproveJob(context.Context, *ApproveJobRequest) (*ApproveJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveJob not implemented")
}

func RegisterWebServer(s *grpc.Server, srv WebServer) {
	s.RegisterService(&_Web_serviceDesc, srv)
}

func _Web_ListProjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebServer).ListProjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Web/ListProjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebServer).ListProjects(ctx, req.(*ListProjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Web_RefreshProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebServer).RefreshProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Web/RefreshProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebServer).RefreshProject(ctx, req.(*RefreshProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Web_GetProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebServer).GetProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Web/GetProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebServer).GetProject(ctx, req.(*GetProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Web_GetJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebServer).GetJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Web/GetJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebServer).GetJob(ctx, req.(*GetJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Web_SubmitJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebServer).SubmitJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Web/SubmitJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebServer).SubmitJob(ctx, req.(*SubmitJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Web_ApproveJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApproveJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebServer).ApproveJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Web/ApproveJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebServer).ApproveJob(ctx, req.(*ApproveJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Web_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Web",
	HandlerType: (*WebServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListProjects",
			Handler:    _Web_ListProjects_Handler,
		},
		{
			MethodName: "RefreshProject",
			Handler:    _Web_RefreshProject_Handler,
		},
		{
			MethodName: "GetProject",
			Handler:    _Web_GetProject_Handler,
		},
		{
			MethodName: "GetJob",
			Handler:    _Web_GetJob_Handler,
		},
		{
			MethodName: "SubmitJob",
			Handler:    _Web_SubmitJob_Handler,
		},
		{
			MethodName: "ApproveJob",
			Handler:    _Web_ApproveJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "web.proto",
}
