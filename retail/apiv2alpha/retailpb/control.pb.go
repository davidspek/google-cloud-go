// Copyright 2021 Google LLC
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
// source: google/cloud/retail/v2alpha/control.proto

package retailpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Configures dynamic metadata that can be linked to a
// [ServingConfig][google.cloud.retail.v2alpha.ServingConfig] and affect search
// or recommendation results at serving time.
type Control struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The behavior/type of the control
	//
	// A behavior/type must be specified on creation. Type cannot be changed once
	// specified (e.g. A Rule control will always be a Rule control.). An
	// INVALID_ARGUMENT will be returned if either condition is violated.
	//
	// Types that are assignable to Control:
	//	*Control_FacetSpec
	//	*Control_Rule
	Control isControl_Control `protobuf_oneof:"control"`
	// Immutable. Fully qualified name
	// `projects/*/locations/global/catalogs/*/controls/*`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The human readable control display name. Used in Retail UI.
	//
	// This field must be a UTF-8 encoded string with a length limit of 128
	// characters. Otherwise, an INVALID_ARGUMENT error is thrown.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Output only. List of [serving
	// config][google.cloud.retail.v2alpha.ServingConfig] ids that are associated
	// with this control in the same
	// [Catalog][google.cloud.retail.v2alpha.Catalog].
	//
	// Note the association is managed via the
	// [ServingConfig][google.cloud.retail.v2alpha.ServingConfig], this is an
	// output only denormalized view.
	AssociatedServingConfigIds []string `protobuf:"bytes,5,rep,name=associated_serving_config_ids,json=associatedServingConfigIds,proto3" json:"associated_serving_config_ids,omitempty"`
	// Required. Immutable. The solution types that the control is used for.
	// Currently we support setting only one type of solution at creation time.
	//
	// Only `SOLUTION_TYPE_SEARCH` value is supported at the moment.
	// If no solution type is provided at creation time, will default to
	// [SOLUTION_TYPE_SEARCH][google.cloud.retail.v2alpha.SolutionType.SOLUTION_TYPE_SEARCH].
	SolutionTypes []SolutionType `protobuf:"varint,6,rep,packed,name=solution_types,json=solutionTypes,proto3,enum=google.cloud.retail.v2alpha.SolutionType" json:"solution_types,omitempty"`
	// Specifies the use case for the control.
	// Affects what condition fields can be set.
	// Only settable by search controls.
	// Will default to
	// [SEARCH_SOLUTION_USE_CASE_SEARCH][google.cloud.retail.v2alpha.SearchSolutionUseCase.SEARCH_SOLUTION_USE_CASE_SEARCH]
	// if not specified. Currently only allow one search_solution_use_case per
	// control.
	SearchSolutionUseCase []SearchSolutionUseCase `protobuf:"varint,7,rep,packed,name=search_solution_use_case,json=searchSolutionUseCase,proto3,enum=google.cloud.retail.v2alpha.SearchSolutionUseCase" json:"search_solution_use_case,omitempty"`
}

func (x *Control) Reset() {
	*x = Control{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2alpha_control_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Control) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Control) ProtoMessage() {}

func (x *Control) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2alpha_control_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Control.ProtoReflect.Descriptor instead.
func (*Control) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2alpha_control_proto_rawDescGZIP(), []int{0}
}

func (m *Control) GetControl() isControl_Control {
	if m != nil {
		return m.Control
	}
	return nil
}

// Deprecated: Marked as deprecated in google/cloud/retail/v2alpha/control.proto.
func (x *Control) GetFacetSpec() *SearchRequest_FacetSpec {
	if x, ok := x.GetControl().(*Control_FacetSpec); ok {
		return x.FacetSpec
	}
	return nil
}

func (x *Control) GetRule() *Rule {
	if x, ok := x.GetControl().(*Control_Rule); ok {
		return x.Rule
	}
	return nil
}

func (x *Control) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Control) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Control) GetAssociatedServingConfigIds() []string {
	if x != nil {
		return x.AssociatedServingConfigIds
	}
	return nil
}

func (x *Control) GetSolutionTypes() []SolutionType {
	if x != nil {
		return x.SolutionTypes
	}
	return nil
}

func (x *Control) GetSearchSolutionUseCase() []SearchSolutionUseCase {
	if x != nil {
		return x.SearchSolutionUseCase
	}
	return nil
}

type isControl_Control interface {
	isControl_Control()
}

type Control_FacetSpec struct {
	// A facet specification to perform faceted search.
	//
	// Note that this field is deprecated and will throw NOT_IMPLEMENTED if
	// used for creating a control.
	//
	// Deprecated: Marked as deprecated in google/cloud/retail/v2alpha/control.proto.
	FacetSpec *SearchRequest_FacetSpec `protobuf:"bytes,3,opt,name=facet_spec,json=facetSpec,proto3,oneof"`
}

type Control_Rule struct {
	// A rule control - a condition-action pair.
	// Enacts a set action when the condition is triggered.
	// For example: Boost "gShoe" when query full matches "Running Shoes".
	Rule *Rule `protobuf:"bytes,4,opt,name=rule,proto3,oneof"`
}

func (*Control_FacetSpec) isControl_Control() {}

func (*Control_Rule) isControl_Control() {}

var File_google_cloud_retail_v2alpha_control_proto protoreflect.FileDescriptor

var file_google_cloud_retail_v2alpha_control_proto_rawDesc = []byte{
	0x0a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xeb, 0x04, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x59, 0x0a, 0x0a,
	0x66, 0x61, 0x63, 0x65, 0x74, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x61, 0x63,
	0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x42, 0x02, 0x18, 0x01, 0x48, 0x00, 0x52, 0x09, 0x66, 0x61,
	0x63, 0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12, 0x37, 0x0a, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x04, 0x72, 0x75, 0x6c, 0x65,
	0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x05, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0c, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x46, 0x0a, 0x1d, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69,
	0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x1a, 0x61,
	0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x64, 0x73, 0x12, 0x58, 0x0a, 0x0e, 0x73, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0e, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x06, 0xe0, 0x41,
	0x02, 0xe0, 0x41, 0x05, 0x52, 0x0d, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x12, 0x6b, 0x0a, 0x18, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x73, 0x6f,
	0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x55, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x15, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65,
	0x3a, 0x71, 0xea, 0x41, 0x6e, 0x0a, 0x1d, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x12, 0x4d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x63, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x7b, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x7d,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x7d, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x42, 0xd0,
	0x01, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x42, 0x0c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x37, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x70, 0x62, 0x3b, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x70, 0x62, 0xa2, 0x02, 0x06, 0x52, 0x45,
	0x54, 0x41, 0x49, 0x4c, 0xaa, 0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x56, 0x32, 0x41, 0x6c, 0x70,
	0x68, 0x61, 0xca, 0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x5c, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5c, 0x56, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0xea, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x3a, 0x3a, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x3a, 0x3a, 0x56, 0x32, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_retail_v2alpha_control_proto_rawDescOnce sync.Once
	file_google_cloud_retail_v2alpha_control_proto_rawDescData = file_google_cloud_retail_v2alpha_control_proto_rawDesc
)

func file_google_cloud_retail_v2alpha_control_proto_rawDescGZIP() []byte {
	file_google_cloud_retail_v2alpha_control_proto_rawDescOnce.Do(func() {
		file_google_cloud_retail_v2alpha_control_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_retail_v2alpha_control_proto_rawDescData)
	})
	return file_google_cloud_retail_v2alpha_control_proto_rawDescData
}

var file_google_cloud_retail_v2alpha_control_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_retail_v2alpha_control_proto_goTypes = []interface{}{
	(*Control)(nil),                 // 0: google.cloud.retail.v2alpha.Control
	(*SearchRequest_FacetSpec)(nil), // 1: google.cloud.retail.v2alpha.SearchRequest.FacetSpec
	(*Rule)(nil),                    // 2: google.cloud.retail.v2alpha.Rule
	(SolutionType)(0),               // 3: google.cloud.retail.v2alpha.SolutionType
	(SearchSolutionUseCase)(0),      // 4: google.cloud.retail.v2alpha.SearchSolutionUseCase
}
var file_google_cloud_retail_v2alpha_control_proto_depIdxs = []int32{
	1, // 0: google.cloud.retail.v2alpha.Control.facet_spec:type_name -> google.cloud.retail.v2alpha.SearchRequest.FacetSpec
	2, // 1: google.cloud.retail.v2alpha.Control.rule:type_name -> google.cloud.retail.v2alpha.Rule
	3, // 2: google.cloud.retail.v2alpha.Control.solution_types:type_name -> google.cloud.retail.v2alpha.SolutionType
	4, // 3: google.cloud.retail.v2alpha.Control.search_solution_use_case:type_name -> google.cloud.retail.v2alpha.SearchSolutionUseCase
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_cloud_retail_v2alpha_control_proto_init() }
func file_google_cloud_retail_v2alpha_control_proto_init() {
	if File_google_cloud_retail_v2alpha_control_proto != nil {
		return
	}
	file_google_cloud_retail_v2alpha_common_proto_init()
	file_google_cloud_retail_v2alpha_search_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_retail_v2alpha_control_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Control); i {
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
	file_google_cloud_retail_v2alpha_control_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Control_FacetSpec)(nil),
		(*Control_Rule)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_retail_v2alpha_control_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_retail_v2alpha_control_proto_goTypes,
		DependencyIndexes: file_google_cloud_retail_v2alpha_control_proto_depIdxs,
		MessageInfos:      file_google_cloud_retail_v2alpha_control_proto_msgTypes,
	}.Build()
	File_google_cloud_retail_v2alpha_control_proto = out.File
	file_google_cloud_retail_v2alpha_control_proto_rawDesc = nil
	file_google_cloud_retail_v2alpha_control_proto_goTypes = nil
	file_google_cloud_retail_v2alpha_control_proto_depIdxs = nil
}
