// Copyright 2022 Google LLC
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
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: google/cloud/documentai/v1/processor.proto

package documentaipb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// The possible states of the processor version.
type ProcessorVersion_State int32

const (
	// The processor version is in an unspecified state.
	ProcessorVersion_STATE_UNSPECIFIED ProcessorVersion_State = 0
	// The processor version is deployed and can be used for processing.
	ProcessorVersion_DEPLOYED ProcessorVersion_State = 1
	// The processor version is being deployed.
	ProcessorVersion_DEPLOYING ProcessorVersion_State = 2
	// The processor version is not deployed and cannot be used for processing.
	ProcessorVersion_UNDEPLOYED ProcessorVersion_State = 3
	// The processor version is being undeployed.
	ProcessorVersion_UNDEPLOYING ProcessorVersion_State = 4
	// The processor version is being created.
	ProcessorVersion_CREATING ProcessorVersion_State = 5
	// The processor version is being deleted.
	ProcessorVersion_DELETING ProcessorVersion_State = 6
	// The processor version failed and is in an indeterminate state.
	ProcessorVersion_FAILED ProcessorVersion_State = 7
)

// Enum value maps for ProcessorVersion_State.
var (
	ProcessorVersion_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "DEPLOYED",
		2: "DEPLOYING",
		3: "UNDEPLOYED",
		4: "UNDEPLOYING",
		5: "CREATING",
		6: "DELETING",
		7: "FAILED",
	}
	ProcessorVersion_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"DEPLOYED":          1,
		"DEPLOYING":         2,
		"UNDEPLOYED":        3,
		"UNDEPLOYING":       4,
		"CREATING":          5,
		"DELETING":          6,
		"FAILED":            7,
	}
)

func (x ProcessorVersion_State) Enum() *ProcessorVersion_State {
	p := new(ProcessorVersion_State)
	*p = x
	return p
}

func (x ProcessorVersion_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProcessorVersion_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_documentai_v1_processor_proto_enumTypes[0].Descriptor()
}

func (ProcessorVersion_State) Type() protoreflect.EnumType {
	return &file_google_cloud_documentai_v1_processor_proto_enumTypes[0]
}

func (x ProcessorVersion_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProcessorVersion_State.Descriptor instead.
func (ProcessorVersion_State) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_documentai_v1_processor_proto_rawDescGZIP(), []int{0, 0}
}

// The possible states of the processor.
type Processor_State int32

const (
	// The processor is in an unspecified state.
	Processor_STATE_UNSPECIFIED Processor_State = 0
	// The processor is enabled, i.e., has an enabled version which can
	// currently serve processing requests and all the feature dependencies have
	// been successfully initialized.
	Processor_ENABLED Processor_State = 1
	// The processor is disabled.
	Processor_DISABLED Processor_State = 2
	// The processor is being enabled, will become `ENABLED` if successful.
	Processor_ENABLING Processor_State = 3
	// The processor is being disabled, will become `DISABLED` if successful.
	Processor_DISABLING Processor_State = 4
	// The processor is being created, will become either `ENABLED` (for
	// successful creation) or `FAILED` (for failed ones).
	// Once a processor is in this state, it can then be used for document
	// processing, but the feature dependencies of the processor might not be
	// fully created yet.
	Processor_CREATING Processor_State = 5
	// The processor failed during creation or initialization of feature
	// dependencies. The user should delete the processor and recreate one as
	// all the functionalities of the processor are disabled.
	Processor_FAILED Processor_State = 6
	// The processor is being deleted, will be removed if successful.
	Processor_DELETING Processor_State = 7
)

// Enum value maps for Processor_State.
var (
	Processor_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "ENABLED",
		2: "DISABLED",
		3: "ENABLING",
		4: "DISABLING",
		5: "CREATING",
		6: "FAILED",
		7: "DELETING",
	}
	Processor_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"ENABLED":           1,
		"DISABLED":          2,
		"ENABLING":          3,
		"DISABLING":         4,
		"CREATING":          5,
		"FAILED":            6,
		"DELETING":          7,
	}
)

func (x Processor_State) Enum() *Processor_State {
	p := new(Processor_State)
	*p = x
	return p
}

func (x Processor_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Processor_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_documentai_v1_processor_proto_enumTypes[1].Descriptor()
}

func (Processor_State) Type() protoreflect.EnumType {
	return &file_google_cloud_documentai_v1_processor_proto_enumTypes[1]
}

func (x Processor_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Processor_State.Descriptor instead.
func (Processor_State) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_documentai_v1_processor_proto_rawDescGZIP(), []int{1, 0}
}

// A processor version is an implementation of a processor. Each processor
// can have multiple versions, pre-trained by Google internally or up-trained
// by the customer. At a time, a processor can only have one default version
// version. So the processor's behavior (when processing documents) is defined
// by a default version
type ProcessorVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the processor version.
	// Format:
	// `projects/{project}/locations/{location}/processors/{processor}/processorVersions/{processor_version}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The display name of the processor version.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The schema of the processor version. Describes the output.
	DocumentSchema *DocumentSchema `protobuf:"bytes,12,opt,name=document_schema,json=documentSchema,proto3" json:"document_schema,omitempty"`
	// The state of the processor version.
	State ProcessorVersion_State `protobuf:"varint,6,opt,name=state,proto3,enum=google.cloud.documentai.v1.ProcessorVersion_State" json:"state,omitempty"`
	// The time the processor version was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The most recently invoked evaluation for the processor version.
	LatestEvaluation *EvaluationReference `protobuf:"bytes,8,opt,name=latest_evaluation,json=latestEvaluation,proto3" json:"latest_evaluation,omitempty"`
	// The KMS key name used for encryption.
	KmsKeyName string `protobuf:"bytes,9,opt,name=kms_key_name,json=kmsKeyName,proto3" json:"kms_key_name,omitempty"`
	// The KMS key version with which data is encrypted.
	KmsKeyVersionName string `protobuf:"bytes,10,opt,name=kms_key_version_name,json=kmsKeyVersionName,proto3" json:"kms_key_version_name,omitempty"`
	// Denotes that this ProcessorVersion is managed by google.
	GoogleManaged bool `protobuf:"varint,11,opt,name=google_managed,json=googleManaged,proto3" json:"google_managed,omitempty"`
	// If set, information about the eventual deprecation of this version.
	DeprecationInfo *ProcessorVersion_DeprecationInfo `protobuf:"bytes,13,opt,name=deprecation_info,json=deprecationInfo,proto3" json:"deprecation_info,omitempty"`
}

func (x *ProcessorVersion) Reset() {
	*x = ProcessorVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_documentai_v1_processor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessorVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessorVersion) ProtoMessage() {}

func (x *ProcessorVersion) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_documentai_v1_processor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessorVersion.ProtoReflect.Descriptor instead.
func (*ProcessorVersion) Descriptor() ([]byte, []int) {
	return file_google_cloud_documentai_v1_processor_proto_rawDescGZIP(), []int{0}
}

func (x *ProcessorVersion) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProcessorVersion) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *ProcessorVersion) GetDocumentSchema() *DocumentSchema {
	if x != nil {
		return x.DocumentSchema
	}
	return nil
}

func (x *ProcessorVersion) GetState() ProcessorVersion_State {
	if x != nil {
		return x.State
	}
	return ProcessorVersion_STATE_UNSPECIFIED
}

func (x *ProcessorVersion) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *ProcessorVersion) GetLatestEvaluation() *EvaluationReference {
	if x != nil {
		return x.LatestEvaluation
	}
	return nil
}

func (x *ProcessorVersion) GetKmsKeyName() string {
	if x != nil {
		return x.KmsKeyName
	}
	return ""
}

func (x *ProcessorVersion) GetKmsKeyVersionName() string {
	if x != nil {
		return x.KmsKeyVersionName
	}
	return ""
}

func (x *ProcessorVersion) GetGoogleManaged() bool {
	if x != nil {
		return x.GoogleManaged
	}
	return false
}

func (x *ProcessorVersion) GetDeprecationInfo() *ProcessorVersion_DeprecationInfo {
	if x != nil {
		return x.DeprecationInfo
	}
	return nil
}

// The first-class citizen for Document AI. Each processor defines how to
// extract structural information from a document.
type Processor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Immutable. The resource name of the processor.
	// Format: `projects/{project}/locations/{location}/processors/{processor}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The processor type, e.g., `OCR_PROCESSOR`, `INVOICE_PROCESSOR`, etc.
	// To get a list of processors types, see
	// [FetchProcessorTypes][google.cloud.documentai.v1.DocumentProcessorService.FetchProcessorTypes].
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// The display name of the processor.
	DisplayName string `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Output only. The state of the processor.
	State Processor_State `protobuf:"varint,4,opt,name=state,proto3,enum=google.cloud.documentai.v1.Processor_State" json:"state,omitempty"`
	// The default processor version.
	DefaultProcessorVersion string `protobuf:"bytes,9,opt,name=default_processor_version,json=defaultProcessorVersion,proto3" json:"default_processor_version,omitempty"`
	// Output only. Immutable. The http endpoint that can be called to invoke
	// processing.
	ProcessEndpoint string `protobuf:"bytes,6,opt,name=process_endpoint,json=processEndpoint,proto3" json:"process_endpoint,omitempty"`
	// The time the processor was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The KMS key used for encryption/decryption in CMEK scenarios.
	// See https://cloud.google.com/security-key-management.
	KmsKeyName string `protobuf:"bytes,8,opt,name=kms_key_name,json=kmsKeyName,proto3" json:"kms_key_name,omitempty"`
}

func (x *Processor) Reset() {
	*x = Processor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_documentai_v1_processor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Processor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Processor) ProtoMessage() {}

func (x *Processor) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_documentai_v1_processor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Processor.ProtoReflect.Descriptor instead.
func (*Processor) Descriptor() ([]byte, []int) {
	return file_google_cloud_documentai_v1_processor_proto_rawDescGZIP(), []int{1}
}

func (x *Processor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Processor) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Processor) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Processor) GetState() Processor_State {
	if x != nil {
		return x.State
	}
	return Processor_STATE_UNSPECIFIED
}

func (x *Processor) GetDefaultProcessorVersion() string {
	if x != nil {
		return x.DefaultProcessorVersion
	}
	return ""
}

func (x *Processor) GetProcessEndpoint() string {
	if x != nil {
		return x.ProcessEndpoint
	}
	return ""
}

func (x *Processor) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Processor) GetKmsKeyName() string {
	if x != nil {
		return x.KmsKeyName
	}
	return ""
}

// Information about the upcoming deprecation of this processor version.
type ProcessorVersion_DeprecationInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The time at which this processor version will be deprecated.
	DeprecationTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=deprecation_time,json=deprecationTime,proto3" json:"deprecation_time,omitempty"`
	// If set, the processor version that will be used as a replacement.
	ReplacementProcessorVersion string `protobuf:"bytes,2,opt,name=replacement_processor_version,json=replacementProcessorVersion,proto3" json:"replacement_processor_version,omitempty"`
}

func (x *ProcessorVersion_DeprecationInfo) Reset() {
	*x = ProcessorVersion_DeprecationInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_documentai_v1_processor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessorVersion_DeprecationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessorVersion_DeprecationInfo) ProtoMessage() {}

func (x *ProcessorVersion_DeprecationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_documentai_v1_processor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessorVersion_DeprecationInfo.ProtoReflect.Descriptor instead.
func (*ProcessorVersion_DeprecationInfo) Descriptor() ([]byte, []int) {
	return file_google_cloud_documentai_v1_processor_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ProcessorVersion_DeprecationInfo) GetDeprecationTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DeprecationTime
	}
	return nil
}

func (x *ProcessorVersion_DeprecationInfo) GetReplacementProcessorVersion() string {
	if x != nil {
		return x.ReplacementProcessorVersion
	}
	return ""
}

var File_google_cloud_documentai_v1_processor_proto protoreflect.FileDescriptor

var file_google_cloud_documentai_v1_processor_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd6, 0x08, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x6f, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x53, 0x0a, 0x0f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x0e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x48, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x5c, 0x0a, 0x11,
	0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x10, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74,
	0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0c, 0x6b, 0x6d,
	0x73, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x6b, 0x6d, 0x73, 0x4b, 0x65, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x14,
	0x6b, 0x6d, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6b, 0x6d, 0x73, 0x4b,
	0x65, 0x79, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a,
	0x0e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x64, 0x12, 0x67, 0x0a, 0x10, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x6f, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x70,
	0x72, 0x65, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x64, 0x65,
	0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0xcd, 0x01,
	0x0a, 0x0f, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x45, 0x0a, 0x10, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x73, 0x0a, 0x1d, 0x72, 0x65, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f,
	0x72, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x2f, 0xfa, 0x41, 0x2c, 0x0a, 0x2a, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x1b, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x84, 0x01,
	0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c,
	0x0a, 0x08, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09,
	0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x55,
	0x4e, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x55,
	0x4e, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08,
	0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45,
	0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x06, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c,
	0x45, 0x44, 0x10, 0x07, 0x3a, 0x96, 0x01, 0xea, 0x41, 0x92, 0x01, 0x0a, 0x2a, 0x64, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x64, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d,
	0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x7d, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f,
	0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x6f, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x7d, 0x22, 0x8f, 0x05,
	0x0a, 0x09, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe0, 0x41, 0x05, 0xe0, 0x41,
	0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x46,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x6f, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x6b, 0x0a, 0x19, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2f, 0xfa, 0x41, 0x2c, 0x0a, 0x2a,
	0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x6f, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x17, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe0,
	0x41, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x6b, 0x6d, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6b, 0x6d, 0x73, 0x4b, 0x65,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x7e, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15,
	0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44,
	0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x02,
	0x12, 0x0c, 0x0a, 0x08, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x0d,
	0x0a, 0x09, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x0c, 0x0a,
	0x08, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x12, 0x0a, 0x0a, 0x06, 0x46,
	0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45, 0x4c, 0x45, 0x54,
	0x49, 0x4e, 0x47, 0x10, 0x07, 0x3a, 0x68, 0xea, 0x41, 0x65, 0x0a, 0x23, 0x64, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x12,
	0x3e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x6f, 0x72, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x7d, 0x42,
	0xd1, 0x01, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x2e,
	0x76, 0x31, 0x42, 0x13, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x69, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x50, 0x01, 0x5a, 0x3e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f,
	0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x70, 0x62, 0x3b, 0x64, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x70, 0x62, 0xaa, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x41, 0x49, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x49,
	0x5c, 0x56, 0x31, 0xea, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x49, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_documentai_v1_processor_proto_rawDescOnce sync.Once
	file_google_cloud_documentai_v1_processor_proto_rawDescData = file_google_cloud_documentai_v1_processor_proto_rawDesc
)

func file_google_cloud_documentai_v1_processor_proto_rawDescGZIP() []byte {
	file_google_cloud_documentai_v1_processor_proto_rawDescOnce.Do(func() {
		file_google_cloud_documentai_v1_processor_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_documentai_v1_processor_proto_rawDescData)
	})
	return file_google_cloud_documentai_v1_processor_proto_rawDescData
}

var file_google_cloud_documentai_v1_processor_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_cloud_documentai_v1_processor_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_documentai_v1_processor_proto_goTypes = []interface{}{
	(ProcessorVersion_State)(0),              // 0: google.cloud.documentai.v1.ProcessorVersion.State
	(Processor_State)(0),                     // 1: google.cloud.documentai.v1.Processor.State
	(*ProcessorVersion)(nil),                 // 2: google.cloud.documentai.v1.ProcessorVersion
	(*Processor)(nil),                        // 3: google.cloud.documentai.v1.Processor
	(*ProcessorVersion_DeprecationInfo)(nil), // 4: google.cloud.documentai.v1.ProcessorVersion.DeprecationInfo
	(*DocumentSchema)(nil),                   // 5: google.cloud.documentai.v1.DocumentSchema
	(*timestamppb.Timestamp)(nil),            // 6: google.protobuf.Timestamp
	(*EvaluationReference)(nil),              // 7: google.cloud.documentai.v1.EvaluationReference
}
var file_google_cloud_documentai_v1_processor_proto_depIdxs = []int32{
	5, // 0: google.cloud.documentai.v1.ProcessorVersion.document_schema:type_name -> google.cloud.documentai.v1.DocumentSchema
	0, // 1: google.cloud.documentai.v1.ProcessorVersion.state:type_name -> google.cloud.documentai.v1.ProcessorVersion.State
	6, // 2: google.cloud.documentai.v1.ProcessorVersion.create_time:type_name -> google.protobuf.Timestamp
	7, // 3: google.cloud.documentai.v1.ProcessorVersion.latest_evaluation:type_name -> google.cloud.documentai.v1.EvaluationReference
	4, // 4: google.cloud.documentai.v1.ProcessorVersion.deprecation_info:type_name -> google.cloud.documentai.v1.ProcessorVersion.DeprecationInfo
	1, // 5: google.cloud.documentai.v1.Processor.state:type_name -> google.cloud.documentai.v1.Processor.State
	6, // 6: google.cloud.documentai.v1.Processor.create_time:type_name -> google.protobuf.Timestamp
	6, // 7: google.cloud.documentai.v1.ProcessorVersion.DeprecationInfo.deprecation_time:type_name -> google.protobuf.Timestamp
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_google_cloud_documentai_v1_processor_proto_init() }
func file_google_cloud_documentai_v1_processor_proto_init() {
	if File_google_cloud_documentai_v1_processor_proto != nil {
		return
	}
	file_google_cloud_documentai_v1_document_schema_proto_init()
	file_google_cloud_documentai_v1_evaluation_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_documentai_v1_processor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessorVersion); i {
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
		file_google_cloud_documentai_v1_processor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Processor); i {
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
		file_google_cloud_documentai_v1_processor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessorVersion_DeprecationInfo); i {
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
			RawDescriptor: file_google_cloud_documentai_v1_processor_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_documentai_v1_processor_proto_goTypes,
		DependencyIndexes: file_google_cloud_documentai_v1_processor_proto_depIdxs,
		EnumInfos:         file_google_cloud_documentai_v1_processor_proto_enumTypes,
		MessageInfos:      file_google_cloud_documentai_v1_processor_proto_msgTypes,
	}.Build()
	File_google_cloud_documentai_v1_processor_proto = out.File
	file_google_cloud_documentai_v1_processor_proto_rawDesc = nil
	file_google_cloud_documentai_v1_processor_proto_goTypes = nil
	file_google_cloud_documentai_v1_processor_proto_depIdxs = nil
}
