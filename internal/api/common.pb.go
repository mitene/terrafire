// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.2
// source: common.proto

package api

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Phase int32

const (
	Phase_Plan  Phase = 0
	Phase_Apply Phase = 1
)

// Enum value maps for Phase.
var (
	Phase_name = map[int32]string{
		0: "Plan",
		1: "Apply",
	}
	Phase_value = map[string]int32{
		"Plan":  0,
		"Apply": 1,
	}
)

func (x Phase) Enum() *Phase {
	p := new(Phase)
	*p = x
	return p
}

func (x Phase) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Phase) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[0].Descriptor()
}

func (Phase) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[0]
}

func (x Phase) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Phase.Descriptor instead.
func (Phase) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

type Source_Type int32

const (
	Source_github Source_Type = 0
)

// Enum value maps for Source_Type.
var (
	Source_Type_name = map[int32]string{
		0: "github",
	}
	Source_Type_value = map[string]int32{
		"github": 0,
	}
)

func (x Source_Type) Enum() *Source_Type {
	p := new(Source_Type)
	*p = x
	return p
}

func (x Source_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Source_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[1].Descriptor()
}

func (Source_Type) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[1]
}

func (x Source_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Source_Type.Descriptor instead.
func (Source_Type) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2, 0}
}

type Job_Status int32

const (
	Job_Pending         Job_Status = 0
	Job_PlanInProgress  Job_Status = 1
	Job_ReviewRequired  Job_Status = 2
	Job_ApplyPending    Job_Status = 3
	Job_ApplyInProgress Job_Status = 4
	Job_Succeeded       Job_Status = 5
	Job_PlanFailed      Job_Status = 6
	Job_ApplyFailed     Job_Status = 7
)

// Enum value maps for Job_Status.
var (
	Job_Status_name = map[int32]string{
		0: "Pending",
		1: "PlanInProgress",
		2: "ReviewRequired",
		3: "ApplyPending",
		4: "ApplyInProgress",
		5: "Succeeded",
		6: "PlanFailed",
		7: "ApplyFailed",
	}
	Job_Status_value = map[string]int32{
		"Pending":         0,
		"PlanInProgress":  1,
		"ReviewRequired":  2,
		"ApplyPending":    3,
		"ApplyInProgress": 4,
		"Succeeded":       5,
		"PlanFailed":      6,
		"ApplyFailed":     7,
	}
)

func (x Job_Status) Enum() *Job_Status {
	p := new(Job_Status)
	*p = x
	return p
}

func (x Job_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Job_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[2].Descriptor()
}

func (Job_Status) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[2]
}

func (x Job_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Job_Status.Descriptor instead.
func (Job_Status) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{4, 0}
}

type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Repo   string  `protobuf:"bytes,2,opt,name=repo,proto3" json:"repo,omitempty"`
	Branch string  `protobuf:"bytes,3,opt,name=branch,proto3" json:"branch,omitempty"`
	Path   string  `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Envs   []*Pair `protobuf:"bytes,5,rep,name=envs,proto3" json:"envs,omitempty"`
}

func (x *Project) Reset() {
	*x = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *Project) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Project) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

func (x *Project) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

func (x *Project) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Project) GetEnvs() []*Pair {
	if x != nil {
		return x.Envs
	}
	return nil
}

type Workspace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Source    *Source `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Workspace string  `protobuf:"bytes,3,opt,name=workspace,proto3" json:"workspace,omitempty"`
	Vars      []*Pair `protobuf:"bytes,4,rep,name=vars,proto3" json:"vars,omitempty"`
	VarFiles  []*Pair `protobuf:"bytes,5,rep,name=var_files,json=varFiles,proto3" json:"var_files,omitempty"`
}

func (x *Workspace) Reset() {
	*x = Workspace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Workspace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Workspace) ProtoMessage() {}

func (x *Workspace) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Workspace.ProtoReflect.Descriptor instead.
func (*Workspace) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

func (x *Workspace) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Workspace) GetSource() *Source {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *Workspace) GetWorkspace() string {
	if x != nil {
		return x.Workspace
	}
	return ""
}

func (x *Workspace) GetVars() []*Pair {
	if x != nil {
		return x.Vars
	}
	return nil
}

func (x *Workspace) GetVarFiles() []*Pair {
	if x != nil {
		return x.VarFiles
	}
	return nil
}

type Source struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  Source_Type `protobuf:"varint,1,opt,name=type,proto3,enum=Source_Type" json:"type,omitempty"`
	Owner string      `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Repo  string      `protobuf:"bytes,3,opt,name=repo,proto3" json:"repo,omitempty"`
	Path  string      `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Ref   string      `protobuf:"bytes,5,opt,name=Ref,proto3" json:"Ref,omitempty"`
}

func (x *Source) Reset() {
	*x = Source{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Source) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Source) ProtoMessage() {}

func (x *Source) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Source.ProtoReflect.Descriptor instead.
func (*Source) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

func (x *Source) GetType() Source_Type {
	if x != nil {
		return x.Type
	}
	return Source_github
}

func (x *Source) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *Source) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

func (x *Source) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Source) GetRef() string {
	if x != nil {
		return x.Ref
	}
	return ""
}

type GitRepository struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Protocol string `protobuf:"bytes,2,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Host     string `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	User     string `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	Password string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *GitRepository) Reset() {
	*x = GitRepository{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GitRepository) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GitRepository) ProtoMessage() {}

func (x *GitRepository) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GitRepository.ProtoReflect.Descriptor instead.
func (*GitRepository) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{3}
}

func (x *GitRepository) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GitRepository) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *GitRepository) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *GitRepository) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *GitRepository) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Job struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               uint64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	StartedAt        *timestamp.Timestamp `protobuf:"bytes,2,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	Project          string               `protobuf:"bytes,3,opt,name=Project,proto3" json:"Project,omitempty"`
	Workspace        *Workspace           `protobuf:"bytes,4,opt,name=Workspace,proto3" json:"Workspace,omitempty"`
	Status           Job_Status           `protobuf:"varint,5,opt,name=status,proto3,enum=Job_Status" json:"status,omitempty"`
	PlanResult       string               `protobuf:"bytes,6,opt,name=plan_result,json=planResult,proto3" json:"plan_result,omitempty"`
	Error            string               `protobuf:"bytes,7,opt,name=error,proto3" json:"error,omitempty"`
	PlanLog          string               `protobuf:"bytes,8,opt,name=plan_log,json=planLog,proto3" json:"plan_log,omitempty"`
	ApplyLog         string               `protobuf:"bytes,9,opt,name=apply_log,json=applyLog,proto3" json:"apply_log,omitempty"`
	ProjectVersion   string               `protobuf:"bytes,10,opt,name=project_version,json=projectVersion,proto3" json:"project_version,omitempty"`
	WorkspaceVersion string               `protobuf:"bytes,11,opt,name=workspace_version,json=workspaceVersion,proto3" json:"workspace_version,omitempty"`
	Destroy          bool                 `protobuf:"varint,12,opt,name=destroy,proto3" json:"destroy,omitempty"`
}

func (x *Job) Reset() {
	*x = Job{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Job) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job) ProtoMessage() {}

func (x *Job) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Job.ProtoReflect.Descriptor instead.
func (*Job) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{4}
}

func (x *Job) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Job) GetStartedAt() *timestamp.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *Job) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *Job) GetWorkspace() *Workspace {
	if x != nil {
		return x.Workspace
	}
	return nil
}

func (x *Job) GetStatus() Job_Status {
	if x != nil {
		return x.Status
	}
	return Job_Pending
}

func (x *Job) GetPlanResult() string {
	if x != nil {
		return x.PlanResult
	}
	return ""
}

func (x *Job) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *Job) GetPlanLog() string {
	if x != nil {
		return x.PlanLog
	}
	return ""
}

func (x *Job) GetApplyLog() string {
	if x != nil {
		return x.ApplyLog
	}
	return ""
}

func (x *Job) GetProjectVersion() string {
	if x != nil {
		return x.ProjectVersion
	}
	return ""
}

func (x *Job) GetWorkspaceVersion() string {
	if x != nil {
		return x.WorkspaceVersion
	}
	return ""
}

func (x *Job) GetDestroy() bool {
	if x != nil {
		return x.Destroy
	}
	return false
}

type Pair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Pair) Reset() {
	*x = Pair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pair) ProtoMessage() {}

func (x *Pair) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pair.ProtoReflect.Descriptor instead.
func (*Pair) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{5}
}

func (x *Pair) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Pair) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x78, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x65,
	0x70, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x19,
	0x0a, 0x04, 0x65, 0x6e, 0x76, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x50,
	0x61, 0x69, 0x72, 0x52, 0x04, 0x65, 0x6e, 0x76, 0x73, 0x22, 0x9d, 0x01, 0x0a, 0x09, 0x57, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x76, 0x61,
	0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x52,
	0x04, 0x76, 0x61, 0x72, 0x73, 0x12, 0x22, 0x0a, 0x09, 0x76, 0x61, 0x72, 0x5f, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x52,
	0x08, 0x76, 0x61, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x8e, 0x01, 0x0a, 0x06, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x65, 0x70, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x65, 0x70, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x52, 0x65, 0x66, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x52, 0x65, 0x66, 0x22, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a,
	0x0a, 0x06, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x10, 0x00, 0x22, 0x83, 0x01, 0x0a, 0x0d, 0x47,
	0x69, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0xaf, 0x04, 0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x28, 0x0a,
	0x09, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x09, 0x57, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x4a, 0x6f, 0x62, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x6c, 0x6f, 0x67, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x4c, 0x6f, 0x67, 0x12, 0x1b,
	0x0a, 0x09, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x4c, 0x6f, 0x67, 0x12, 0x27, 0x0a, 0x0f, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x11, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x22, 0x94, 0x01, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x6e, 0x50, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x41,
	0x70, 0x70, 0x6c, 0x79, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x10, 0x03, 0x12, 0x13, 0x0a,
	0x0f, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x49, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x75, 0x63, 0x63, 0x65, 0x65, 0x64, 0x65, 0x64, 0x10,
	0x05, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x6e, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10,
	0x06, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64,
	0x10, 0x07, 0x22, 0x2e, 0x0a, 0x04, 0x50, 0x61, 0x69, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x2a, 0x1c, 0x0a, 0x05, 0x50, 0x68, 0x61, 0x73, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x50,
	0x6c, 0x61, 0x6e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x10, 0x01,
	0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d,
	0x69, 0x74, 0x65, 0x6e, 0x65, 0x2f, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x69, 0x72, 0x65, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_common_proto_goTypes = []interface{}{
	(Phase)(0),                  // 0: Phase
	(Source_Type)(0),            // 1: Source.Type
	(Job_Status)(0),             // 2: Job.Status
	(*Project)(nil),             // 3: Project
	(*Workspace)(nil),           // 4: Workspace
	(*Source)(nil),              // 5: Source
	(*GitRepository)(nil),       // 6: GitRepository
	(*Job)(nil),                 // 7: Job
	(*Pair)(nil),                // 8: Pair
	(*timestamp.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_common_proto_depIdxs = []int32{
	8, // 0: Project.envs:type_name -> Pair
	5, // 1: Workspace.source:type_name -> Source
	8, // 2: Workspace.vars:type_name -> Pair
	8, // 3: Workspace.var_files:type_name -> Pair
	1, // 4: Source.type:type_name -> Source.Type
	9, // 5: Job.started_at:type_name -> google.protobuf.Timestamp
	4, // 6: Job.Workspace:type_name -> Workspace
	2, // 7: Job.status:type_name -> Job.Status
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project); i {
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
		file_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Workspace); i {
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
		file_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Source); i {
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
		file_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GitRepository); i {
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
		file_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Job); i {
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
		file_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pair); i {
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
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		EnumInfos:         file_common_proto_enumTypes,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}
