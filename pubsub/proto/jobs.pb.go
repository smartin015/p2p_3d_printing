// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: proto/jobs.proto

package proto

import (
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

type Job struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Protocol string `protobuf:"bytes,3,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Data     []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Job) Reset() {
	*x = Job{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_jobs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Job) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job) ProtoMessage() {}

func (x *Job) ProtoReflect() protoreflect.Message {
	mi := &file_proto_jobs_proto_msgTypes[0]
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
	return file_proto_jobs_proto_rawDescGZIP(), []int{0}
}

func (x *Job) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Job) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *Job) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type SetJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Job *Job `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
}

func (x *SetJobRequest) Reset() {
	*x = SetJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_jobs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetJobRequest) ProtoMessage() {}

func (x *SetJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_jobs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetJobRequest.ProtoReflect.Descriptor instead.
func (*SetJobRequest) Descriptor() ([]byte, []int) {
	return file_proto_jobs_proto_rawDescGZIP(), []int{1}
}

func (x *SetJobRequest) GetJob() *Job {
	if x != nil {
		return x.Job
	}
	return nil
}

type DeleteJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId string `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
}

func (x *DeleteJobRequest) Reset() {
	*x = DeleteJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_jobs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteJobRequest) ProtoMessage() {}

func (x *DeleteJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_jobs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteJobRequest.ProtoReflect.Descriptor instead.
func (*DeleteJobRequest) Descriptor() ([]byte, []int) {
	return file_proto_jobs_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteJobRequest) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

type AcquireJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId string `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
}

func (x *AcquireJobRequest) Reset() {
	*x = AcquireJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_jobs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcquireJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcquireJobRequest) ProtoMessage() {}

func (x *AcquireJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_jobs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcquireJobRequest.ProtoReflect.Descriptor instead.
func (*AcquireJobRequest) Descriptor() ([]byte, []int) {
	return file_proto_jobs_proto_rawDescGZIP(), []int{3}
}

func (x *AcquireJobRequest) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

type ReleaseJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId string `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
}

func (x *ReleaseJobRequest) Reset() {
	*x = ReleaseJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_jobs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleaseJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseJobRequest) ProtoMessage() {}

func (x *ReleaseJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_jobs_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseJobRequest.ProtoReflect.Descriptor instead.
func (*ReleaseJobRequest) Descriptor() ([]byte, []int) {
	return file_proto_jobs_proto_rawDescGZIP(), []int{4}
}

func (x *ReleaseJobRequest) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

// JobMutationResponse is published as a status update when any mutating
// request causes a job to change
type JobMutationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Job *Job `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
}

func (x *JobMutationResponse) Reset() {
	*x = JobMutationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_jobs_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobMutationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobMutationResponse) ProtoMessage() {}

func (x *JobMutationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_jobs_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobMutationResponse.ProtoReflect.Descriptor instead.
func (*JobMutationResponse) Descriptor() ([]byte, []int) {
	return file_proto_jobs_proto_rawDescGZIP(), []int{5}
}

func (x *JobMutationResponse) GetJob() *Job {
	if x != nil {
		return x.Job
	}
	return nil
}

type GetJobsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetJobsRequest) Reset() {
	*x = GetJobsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_jobs_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJobsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJobsRequest) ProtoMessage() {}

func (x *GetJobsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_jobs_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJobsRequest.ProtoReflect.Descriptor instead.
func (*GetJobsRequest) Descriptor() ([]byte, []int) {
	return file_proto_jobs_proto_rawDescGZIP(), []int{6}
}

type GetJobsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jobs []*Job `protobuf:"bytes,1,rep,name=jobs,proto3" json:"jobs,omitempty"`
}

func (x *GetJobsResponse) Reset() {
	*x = GetJobsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_jobs_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJobsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJobsResponse) ProtoMessage() {}

func (x *GetJobsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_jobs_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJobsResponse.ProtoReflect.Descriptor instead.
func (*GetJobsResponse) Descriptor() ([]byte, []int) {
	return file_proto_jobs_proto_rawDescGZIP(), []int{7}
}

func (x *GetJobsResponse) GetJobs() []*Job {
	if x != nil {
		return x.Jobs
	}
	return nil
}

var File_proto_jobs_proto protoreflect.FileDescriptor

var file_proto_jobs_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x22, 0x45, 0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x2c, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x22, 0x29, 0x0a,
	0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x22, 0x2a, 0x0a, 0x11, 0x41, 0x63, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a,
	0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a,
	0x6f, 0x62, 0x49, 0x64, 0x22, 0x2a, 0x0a, 0x11, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4a,
	0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64,
	0x22, 0x32, 0x0a, 0x13, 0x4a, 0x6f, 0x62, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x4a, 0x6f, 0x62, 0x52,
	0x03, 0x6a, 0x6f, 0x62, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x30, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x6a, 0x6f, 0x62,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x4a,
	0x6f, 0x62, 0x52, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x30, 0x31,
	0x35, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x2f, 0x70, 0x75, 0x62, 0x73,
	0x75, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_jobs_proto_rawDescOnce sync.Once
	file_proto_jobs_proto_rawDescData = file_proto_jobs_proto_rawDesc
)

func file_proto_jobs_proto_rawDescGZIP() []byte {
	file_proto_jobs_proto_rawDescOnce.Do(func() {
		file_proto_jobs_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_jobs_proto_rawDescData)
	})
	return file_proto_jobs_proto_rawDescData
}

var file_proto_jobs_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_jobs_proto_goTypes = []interface{}{
	(*Job)(nil),                 // 0: jobs.Job
	(*SetJobRequest)(nil),       // 1: jobs.SetJobRequest
	(*DeleteJobRequest)(nil),    // 2: jobs.DeleteJobRequest
	(*AcquireJobRequest)(nil),   // 3: jobs.AcquireJobRequest
	(*ReleaseJobRequest)(nil),   // 4: jobs.ReleaseJobRequest
	(*JobMutationResponse)(nil), // 5: jobs.JobMutationResponse
	(*GetJobsRequest)(nil),      // 6: jobs.GetJobsRequest
	(*GetJobsResponse)(nil),     // 7: jobs.GetJobsResponse
}
var file_proto_jobs_proto_depIdxs = []int32{
	0, // 0: jobs.SetJobRequest.job:type_name -> jobs.Job
	0, // 1: jobs.JobMutationResponse.job:type_name -> jobs.Job
	0, // 2: jobs.GetJobsResponse.jobs:type_name -> jobs.Job
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_jobs_proto_init() }
func file_proto_jobs_proto_init() {
	if File_proto_jobs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_jobs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_jobs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetJobRequest); i {
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
		file_proto_jobs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteJobRequest); i {
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
		file_proto_jobs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcquireJobRequest); i {
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
		file_proto_jobs_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReleaseJobRequest); i {
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
		file_proto_jobs_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobMutationResponse); i {
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
		file_proto_jobs_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJobsRequest); i {
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
		file_proto_jobs_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJobsResponse); i {
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
			RawDescriptor: file_proto_jobs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_jobs_proto_goTypes,
		DependencyIndexes: file_proto_jobs_proto_depIdxs,
		MessageInfos:      file_proto_jobs_proto_msgTypes,
	}.Build()
	File_proto_jobs_proto = out.File
	file_proto_jobs_proto_rawDesc = nil
	file_proto_jobs_proto_goTypes = nil
	file_proto_jobs_proto_depIdxs = nil
}
