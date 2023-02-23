// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.9
// source: google/cloud/tasks/v2beta2/task.proto

package cloudtaskspb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The view specifies a subset of [Task][google.cloud.tasks.v2beta2.Task] data.
//
// When a task is returned in a response, not all
// information is retrieved by default because some data, such as
// payloads, might be desirable to return only when needed because
// of its large size or because of the sensitivity of data that it
// contains.
type Task_View int32

const (
	// Unspecified. Defaults to BASIC.
	Task_VIEW_UNSPECIFIED Task_View = 0
	// The basic view omits fields which can be large or can contain
	// sensitive data.
	//
	// This view does not include the
	// ([payload in AppEngineHttpRequest][google.cloud.tasks.v2beta2.AppEngineHttpRequest] and
	// [payload in PullMessage][google.cloud.tasks.v2beta2.PullMessage.payload]). These payloads are
	// desirable to return only when needed, because they can be large
	// and because of the sensitivity of the data that you choose to
	// store in it.
	Task_BASIC Task_View = 1
	// All information is returned.
	//
	// Authorization for [FULL][google.cloud.tasks.v2beta2.Task.View.FULL] requires
	// `cloudtasks.tasks.fullView` [Google IAM](https://cloud.google.com/iam/)
	// permission on the [Queue][google.cloud.tasks.v2beta2.Queue] resource.
	Task_FULL Task_View = 2
)

// Enum value maps for Task_View.
var (
	Task_View_name = map[int32]string{
		0: "VIEW_UNSPECIFIED",
		1: "BASIC",
		2: "FULL",
	}
	Task_View_value = map[string]int32{
		"VIEW_UNSPECIFIED": 0,
		"BASIC":            1,
		"FULL":             2,
	}
)

func (x Task_View) Enum() *Task_View {
	p := new(Task_View)
	*p = x
	return p
}

func (x Task_View) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Task_View) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_tasks_v2beta2_task_proto_enumTypes[0].Descriptor()
}

func (Task_View) Type() protoreflect.EnumType {
	return &file_google_cloud_tasks_v2beta2_task_proto_enumTypes[0]
}

func (x Task_View) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Task_View.Descriptor instead.
func (Task_View) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_tasks_v2beta2_task_proto_rawDescGZIP(), []int{0, 0}
}

// A unit of scheduled work.
type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optionally caller-specified in [CreateTask][google.cloud.tasks.v2beta2.CloudTasks.CreateTask].
	//
	// The task name.
	//
	// The task name must have the following format:
	// `projects/PROJECT_ID/locations/LOCATION_ID/queues/QUEUE_ID/tasks/TASK_ID`
	//
	//   - `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]),
	//     hyphens (-), colons (:), or periods (.).
	//     For more information, see
	//     [Identifying
	//     projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects)
	//   - `LOCATION_ID` is the canonical ID for the task's location.
	//     The list of available locations can be obtained by calling
	//     [ListLocations][google.cloud.location.Locations.ListLocations].
	//     For more information, see https://cloud.google.com/about/locations/.
	//   - `QUEUE_ID` can contain letters ([A-Za-z]), numbers ([0-9]), or
	//     hyphens (-). The maximum length is 100 characters.
	//   - `TASK_ID` can contain only letters ([A-Za-z]), numbers ([0-9]),
	//     hyphens (-), or underscores (_). The maximum length is 500 characters.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required.
	//
	// The task's payload is used by the task's target to process the task.
	// A payload is valid only if it is compatible with the queue's target.
	//
	// Types that are assignable to PayloadType:
	//
	//	*Task_AppEngineHttpRequest
	//	*Task_PullMessage
	PayloadType isTask_PayloadType `protobuf_oneof:"payload_type"`
	// The time when the task is scheduled to be attempted.
	//
	// For App Engine queues, this is when the task will be attempted or retried.
	//
	// For pull queues, this is the time when the task is available to
	// be leased; if a task is currently leased, this is the time when
	// the current lease expires, that is, the time that the task was
	// leased plus the [lease_duration][google.cloud.tasks.v2beta2.LeaseTasksRequest.lease_duration].
	//
	// `schedule_time` will be truncated to the nearest microsecond.
	ScheduleTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=schedule_time,json=scheduleTime,proto3" json:"schedule_time,omitempty"`
	// Output only. The time that the task was created.
	//
	// `create_time` will be truncated to the nearest second.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The task status.
	Status *TaskStatus `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	// Output only. The view specifies which subset of the [Task][google.cloud.tasks.v2beta2.Task] has
	// been returned.
	View Task_View `protobuf:"varint,8,opt,name=view,proto3,enum=google.cloud.tasks.v2beta2.Task_View" json:"view,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_tasks_v2beta2_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_tasks_v2beta2_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_google_cloud_tasks_v2beta2_task_proto_rawDescGZIP(), []int{0}
}

func (x *Task) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (m *Task) GetPayloadType() isTask_PayloadType {
	if m != nil {
		return m.PayloadType
	}
	return nil
}

func (x *Task) GetAppEngineHttpRequest() *AppEngineHttpRequest {
	if x, ok := x.GetPayloadType().(*Task_AppEngineHttpRequest); ok {
		return x.AppEngineHttpRequest
	}
	return nil
}

func (x *Task) GetPullMessage() *PullMessage {
	if x, ok := x.GetPayloadType().(*Task_PullMessage); ok {
		return x.PullMessage
	}
	return nil
}

func (x *Task) GetScheduleTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ScheduleTime
	}
	return nil
}

func (x *Task) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Task) GetStatus() *TaskStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *Task) GetView() Task_View {
	if x != nil {
		return x.View
	}
	return Task_VIEW_UNSPECIFIED
}

type isTask_PayloadType interface {
	isTask_PayloadType()
}

type Task_AppEngineHttpRequest struct {
	// App Engine HTTP request that is sent to the task's target. Can
	// be set only if
	// [app_engine_http_target][google.cloud.tasks.v2beta2.Queue.app_engine_http_target] is set
	// on the queue.
	//
	// An App Engine task is a task that has [AppEngineHttpRequest][google.cloud.tasks.v2beta2.AppEngineHttpRequest] set.
	AppEngineHttpRequest *AppEngineHttpRequest `protobuf:"bytes,3,opt,name=app_engine_http_request,json=appEngineHttpRequest,proto3,oneof"`
}

type Task_PullMessage struct {
	// [LeaseTasks][google.cloud.tasks.v2beta2.CloudTasks.LeaseTasks] to process the task. Can be
	// set only if [pull_target][google.cloud.tasks.v2beta2.Queue.pull_target] is set on the queue.
	//
	// A pull task is a task that has [PullMessage][google.cloud.tasks.v2beta2.PullMessage] set.
	PullMessage *PullMessage `protobuf:"bytes,4,opt,name=pull_message,json=pullMessage,proto3,oneof"`
}

func (*Task_AppEngineHttpRequest) isTask_PayloadType() {}

func (*Task_PullMessage) isTask_PayloadType() {}

// Status of the task.
type TaskStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The number of attempts dispatched.
	//
	// This count includes attempts which have been dispatched but haven't
	// received a response.
	AttemptDispatchCount int32 `protobuf:"varint,1,opt,name=attempt_dispatch_count,json=attemptDispatchCount,proto3" json:"attempt_dispatch_count,omitempty"`
	// Output only. The number of attempts which have received a response.
	//
	// This field is not calculated for [pull tasks][google.cloud.tasks.v2beta2.PullMessage].
	AttemptResponseCount int32 `protobuf:"varint,2,opt,name=attempt_response_count,json=attemptResponseCount,proto3" json:"attempt_response_count,omitempty"`
	// Output only. The status of the task's first attempt.
	//
	// Only [dispatch_time][google.cloud.tasks.v2beta2.AttemptStatus.dispatch_time] will be set.
	// The other [AttemptStatus][google.cloud.tasks.v2beta2.AttemptStatus] information is not retained by Cloud Tasks.
	//
	// This field is not calculated for [pull tasks][google.cloud.tasks.v2beta2.PullMessage].
	FirstAttemptStatus *AttemptStatus `protobuf:"bytes,3,opt,name=first_attempt_status,json=firstAttemptStatus,proto3" json:"first_attempt_status,omitempty"`
	// Output only. The status of the task's last attempt.
	//
	// This field is not calculated for [pull tasks][google.cloud.tasks.v2beta2.PullMessage].
	LastAttemptStatus *AttemptStatus `protobuf:"bytes,4,opt,name=last_attempt_status,json=lastAttemptStatus,proto3" json:"last_attempt_status,omitempty"`
}

func (x *TaskStatus) Reset() {
	*x = TaskStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_tasks_v2beta2_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskStatus) ProtoMessage() {}

func (x *TaskStatus) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_tasks_v2beta2_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskStatus.ProtoReflect.Descriptor instead.
func (*TaskStatus) Descriptor() ([]byte, []int) {
	return file_google_cloud_tasks_v2beta2_task_proto_rawDescGZIP(), []int{1}
}

func (x *TaskStatus) GetAttemptDispatchCount() int32 {
	if x != nil {
		return x.AttemptDispatchCount
	}
	return 0
}

func (x *TaskStatus) GetAttemptResponseCount() int32 {
	if x != nil {
		return x.AttemptResponseCount
	}
	return 0
}

func (x *TaskStatus) GetFirstAttemptStatus() *AttemptStatus {
	if x != nil {
		return x.FirstAttemptStatus
	}
	return nil
}

func (x *TaskStatus) GetLastAttemptStatus() *AttemptStatus {
	if x != nil {
		return x.LastAttemptStatus
	}
	return nil
}

// The status of a task attempt.
type AttemptStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The time that this attempt was scheduled.
	//
	// `schedule_time` will be truncated to the nearest microsecond.
	ScheduleTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=schedule_time,json=scheduleTime,proto3" json:"schedule_time,omitempty"`
	// Output only. The time that this attempt was dispatched.
	//
	// `dispatch_time` will be truncated to the nearest microsecond.
	DispatchTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=dispatch_time,json=dispatchTime,proto3" json:"dispatch_time,omitempty"`
	// Output only. The time that this attempt response was received.
	//
	// `response_time` will be truncated to the nearest microsecond.
	ResponseTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=response_time,json=responseTime,proto3" json:"response_time,omitempty"`
	// Output only. The response from the target for this attempt.
	//
	// If the task has not been attempted or the task is currently running
	// then the response status is unset.
	ResponseStatus *status.Status `protobuf:"bytes,4,opt,name=response_status,json=responseStatus,proto3" json:"response_status,omitempty"`
}

func (x *AttemptStatus) Reset() {
	*x = AttemptStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_tasks_v2beta2_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttemptStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttemptStatus) ProtoMessage() {}

func (x *AttemptStatus) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_tasks_v2beta2_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttemptStatus.ProtoReflect.Descriptor instead.
func (*AttemptStatus) Descriptor() ([]byte, []int) {
	return file_google_cloud_tasks_v2beta2_task_proto_rawDescGZIP(), []int{2}
}

func (x *AttemptStatus) GetScheduleTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ScheduleTime
	}
	return nil
}

func (x *AttemptStatus) GetDispatchTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DispatchTime
	}
	return nil
}

func (x *AttemptStatus) GetResponseTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ResponseTime
	}
	return nil
}

func (x *AttemptStatus) GetResponseStatus() *status.Status {
	if x != nil {
		return x.ResponseStatus
	}
	return nil
}

var File_google_cloud_tasks_v2beta2_task_proto protoreflect.FileDescriptor

var file_google_cloud_tasks_v2beta2_task_proto_rawDesc = []byte{
	0x0a, 0x25, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74,
	0x61, 0x73, 0x6b, 0x73, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2f, 0x74, 0x61, 0x73,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x32, 0x62, 0x65,
	0x74, 0x61, 0x32, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x61, 0x73,
	0x6b, 0x73, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2f, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xf9, 0x04, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x69,
	0x0a, 0x17, 0x61, 0x70, 0x70, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5f, 0x68, 0x74, 0x74,
	0x70, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74,
	0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2e, 0x41, 0x70, 0x70,
	0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x48, 0x00, 0x52, 0x14, 0x61, 0x70, 0x70, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x48, 0x74,
	0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4c, 0x0a, 0x0c, 0x70, 0x75, 0x6c,
	0x6c, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74,
	0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2e, 0x50, 0x75, 0x6c,
	0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x70, 0x75, 0x6c, 0x6c,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3f, 0x0a, 0x0d, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74,
	0x61, 0x32, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x04, 0x76, 0x69, 0x65, 0x77, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x32,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x52, 0x04, 0x76, 0x69, 0x65, 0x77,
	0x22, 0x31, 0x0a, 0x04, 0x56, 0x69, 0x65, 0x77, 0x12, 0x14, 0x0a, 0x10, 0x56, 0x49, 0x45, 0x57,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09,
	0x0a, 0x05, 0x42, 0x41, 0x53, 0x49, 0x43, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x55, 0x4c,
	0x4c, 0x10, 0x02, 0x3a, 0x68, 0xea, 0x41, 0x65, 0x0a, 0x1e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x74,
	0x61, 0x73, 0x6b, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x43, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x7d, 0x2f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x73, 0x2f, 0x7b, 0x71, 0x75, 0x65, 0x75, 0x65, 0x7d,
	0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2f, 0x7b, 0x74, 0x61, 0x73, 0x6b, 0x7d, 0x42, 0x0e, 0x0a,
	0x0c, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0xb0, 0x02,
	0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x34, 0x0a, 0x16,
	0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x14, 0x61, 0x74,
	0x74, 0x65, 0x6d, 0x70, 0x74, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x34, 0x0a, 0x16, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x14, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x5b, 0x0a, 0x14, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x32, 0x62, 0x65,
	0x74, 0x61, 0x32, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x12, 0x66, 0x69, 0x72, 0x73, 0x74, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x59, 0x0a, 0x13, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x61, 0x74,
	0x74, 0x65, 0x6d, 0x70, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2e,
	0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x11, 0x6c,
	0x61, 0x73, 0x74, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x8f, 0x02, 0x0a, 0x0d, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x3f, 0x0a, 0x0d, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x3f, 0x0a, 0x0d, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x3f, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x42, 0x63, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x32, 0x62,
	0x65, 0x74, 0x61, 0x32, 0x42, 0x09, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x34, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2f, 0x61, 0x70, 0x69,
	0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x70, 0x62, 0x3b,
	0x74, 0x61, 0x73, 0x6b, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_tasks_v2beta2_task_proto_rawDescOnce sync.Once
	file_google_cloud_tasks_v2beta2_task_proto_rawDescData = file_google_cloud_tasks_v2beta2_task_proto_rawDesc
)

func file_google_cloud_tasks_v2beta2_task_proto_rawDescGZIP() []byte {
	file_google_cloud_tasks_v2beta2_task_proto_rawDescOnce.Do(func() {
		file_google_cloud_tasks_v2beta2_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_tasks_v2beta2_task_proto_rawDescData)
	})
	return file_google_cloud_tasks_v2beta2_task_proto_rawDescData
}

var file_google_cloud_tasks_v2beta2_task_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_tasks_v2beta2_task_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_tasks_v2beta2_task_proto_goTypes = []interface{}{
	(Task_View)(0),                // 0: google.cloud.tasks.v2beta2.Task.View
	(*Task)(nil),                  // 1: google.cloud.tasks.v2beta2.Task
	(*TaskStatus)(nil),            // 2: google.cloud.tasks.v2beta2.TaskStatus
	(*AttemptStatus)(nil),         // 3: google.cloud.tasks.v2beta2.AttemptStatus
	(*AppEngineHttpRequest)(nil),  // 4: google.cloud.tasks.v2beta2.AppEngineHttpRequest
	(*PullMessage)(nil),           // 5: google.cloud.tasks.v2beta2.PullMessage
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*status.Status)(nil),         // 7: google.rpc.Status
}
var file_google_cloud_tasks_v2beta2_task_proto_depIdxs = []int32{
	4,  // 0: google.cloud.tasks.v2beta2.Task.app_engine_http_request:type_name -> google.cloud.tasks.v2beta2.AppEngineHttpRequest
	5,  // 1: google.cloud.tasks.v2beta2.Task.pull_message:type_name -> google.cloud.tasks.v2beta2.PullMessage
	6,  // 2: google.cloud.tasks.v2beta2.Task.schedule_time:type_name -> google.protobuf.Timestamp
	6,  // 3: google.cloud.tasks.v2beta2.Task.create_time:type_name -> google.protobuf.Timestamp
	2,  // 4: google.cloud.tasks.v2beta2.Task.status:type_name -> google.cloud.tasks.v2beta2.TaskStatus
	0,  // 5: google.cloud.tasks.v2beta2.Task.view:type_name -> google.cloud.tasks.v2beta2.Task.View
	3,  // 6: google.cloud.tasks.v2beta2.TaskStatus.first_attempt_status:type_name -> google.cloud.tasks.v2beta2.AttemptStatus
	3,  // 7: google.cloud.tasks.v2beta2.TaskStatus.last_attempt_status:type_name -> google.cloud.tasks.v2beta2.AttemptStatus
	6,  // 8: google.cloud.tasks.v2beta2.AttemptStatus.schedule_time:type_name -> google.protobuf.Timestamp
	6,  // 9: google.cloud.tasks.v2beta2.AttemptStatus.dispatch_time:type_name -> google.protobuf.Timestamp
	6,  // 10: google.cloud.tasks.v2beta2.AttemptStatus.response_time:type_name -> google.protobuf.Timestamp
	7,  // 11: google.cloud.tasks.v2beta2.AttemptStatus.response_status:type_name -> google.rpc.Status
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_google_cloud_tasks_v2beta2_task_proto_init() }
func file_google_cloud_tasks_v2beta2_task_proto_init() {
	if File_google_cloud_tasks_v2beta2_task_proto != nil {
		return
	}
	file_google_cloud_tasks_v2beta2_target_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_tasks_v2beta2_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_google_cloud_tasks_v2beta2_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskStatus); i {
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
		file_google_cloud_tasks_v2beta2_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttemptStatus); i {
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
	file_google_cloud_tasks_v2beta2_task_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Task_AppEngineHttpRequest)(nil),
		(*Task_PullMessage)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_tasks_v2beta2_task_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_tasks_v2beta2_task_proto_goTypes,
		DependencyIndexes: file_google_cloud_tasks_v2beta2_task_proto_depIdxs,
		EnumInfos:         file_google_cloud_tasks_v2beta2_task_proto_enumTypes,
		MessageInfos:      file_google_cloud_tasks_v2beta2_task_proto_msgTypes,
	}.Build()
	File_google_cloud_tasks_v2beta2_task_proto = out.File
	file_google_cloud_tasks_v2beta2_task_proto_rawDesc = nil
	file_google_cloud_tasks_v2beta2_task_proto_goTypes = nil
	file_google_cloud_tasks_v2beta2_task_proto_depIdxs = nil
}
