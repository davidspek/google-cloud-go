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
// source: google/cloud/channel/v1/entitlement_changes.proto

package channelpb

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

// Specifies the type of change action
type EntitlementChange_ChangeType int32

const (
	// Not used.
	EntitlementChange_CHANGE_TYPE_UNSPECIFIED EntitlementChange_ChangeType = 0
	// New Entitlement was created.
	EntitlementChange_CREATED EntitlementChange_ChangeType = 1
	// Price plan associated with an Entitlement was changed.
	EntitlementChange_PRICE_PLAN_SWITCHED EntitlementChange_ChangeType = 3
	// Number of seats committed for a commitment Entitlement was changed.
	EntitlementChange_COMMITMENT_CHANGED EntitlementChange_ChangeType = 4
	// An annual Entitlement was renewed.
	EntitlementChange_RENEWED EntitlementChange_ChangeType = 5
	// Entitlement was suspended.
	EntitlementChange_SUSPENDED EntitlementChange_ChangeType = 6
	// Entitlement was activated.
	EntitlementChange_ACTIVATED EntitlementChange_ChangeType = 7
	// Entitlement was cancelled.
	EntitlementChange_CANCELLED EntitlementChange_ChangeType = 8
	// Entitlement was upgraded or downgraded for ex. from Google Workspace
	// Business Standard to Google Workspace Business Plus.
	EntitlementChange_SKU_CHANGED EntitlementChange_ChangeType = 9
	// The settings for renewal of an Entitlement have changed.
	EntitlementChange_RENEWAL_SETTING_CHANGED EntitlementChange_ChangeType = 10
	// Use for Google Workspace subscription.
	// Either a trial was converted to a paid subscription or a new subscription
	// with no trial is created.
	EntitlementChange_PAID_SUBSCRIPTION_STARTED EntitlementChange_ChangeType = 11
	// License cap was changed for the entitlement.
	EntitlementChange_LICENSE_CAP_CHANGED EntitlementChange_ChangeType = 12
	// The suspension details have changed (but it is still suspended).
	EntitlementChange_SUSPENSION_DETAILS_CHANGED EntitlementChange_ChangeType = 13
	// The trial end date was extended.
	EntitlementChange_TRIAL_END_DATE_EXTENDED EntitlementChange_ChangeType = 14
	// Entitlement started trial.
	EntitlementChange_TRIAL_STARTED EntitlementChange_ChangeType = 15
)

// Enum value maps for EntitlementChange_ChangeType.
var (
	EntitlementChange_ChangeType_name = map[int32]string{
		0:  "CHANGE_TYPE_UNSPECIFIED",
		1:  "CREATED",
		3:  "PRICE_PLAN_SWITCHED",
		4:  "COMMITMENT_CHANGED",
		5:  "RENEWED",
		6:  "SUSPENDED",
		7:  "ACTIVATED",
		8:  "CANCELLED",
		9:  "SKU_CHANGED",
		10: "RENEWAL_SETTING_CHANGED",
		11: "PAID_SUBSCRIPTION_STARTED",
		12: "LICENSE_CAP_CHANGED",
		13: "SUSPENSION_DETAILS_CHANGED",
		14: "TRIAL_END_DATE_EXTENDED",
		15: "TRIAL_STARTED",
	}
	EntitlementChange_ChangeType_value = map[string]int32{
		"CHANGE_TYPE_UNSPECIFIED":    0,
		"CREATED":                    1,
		"PRICE_PLAN_SWITCHED":        3,
		"COMMITMENT_CHANGED":         4,
		"RENEWED":                    5,
		"SUSPENDED":                  6,
		"ACTIVATED":                  7,
		"CANCELLED":                  8,
		"SKU_CHANGED":                9,
		"RENEWAL_SETTING_CHANGED":    10,
		"PAID_SUBSCRIPTION_STARTED":  11,
		"LICENSE_CAP_CHANGED":        12,
		"SUSPENSION_DETAILS_CHANGED": 13,
		"TRIAL_END_DATE_EXTENDED":    14,
		"TRIAL_STARTED":              15,
	}
)

func (x EntitlementChange_ChangeType) Enum() *EntitlementChange_ChangeType {
	p := new(EntitlementChange_ChangeType)
	*p = x
	return p
}

func (x EntitlementChange_ChangeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EntitlementChange_ChangeType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_channel_v1_entitlement_changes_proto_enumTypes[0].Descriptor()
}

func (EntitlementChange_ChangeType) Type() protoreflect.EnumType {
	return &file_google_cloud_channel_v1_entitlement_changes_proto_enumTypes[0]
}

func (x EntitlementChange_ChangeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EntitlementChange_ChangeType.Descriptor instead.
func (EntitlementChange_ChangeType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_channel_v1_entitlement_changes_proto_rawDescGZIP(), []int{0, 0}
}

// Specifies the type of operator responsible for the change
type EntitlementChange_OperatorType int32

const (
	// Not used.
	EntitlementChange_OPERATOR_TYPE_UNSPECIFIED EntitlementChange_OperatorType = 0
	// Customer service representative.
	EntitlementChange_CUSTOMER_SERVICE_REPRESENTATIVE EntitlementChange_OperatorType = 1
	// System auto job.
	EntitlementChange_SYSTEM EntitlementChange_OperatorType = 2
	// Customer user.
	EntitlementChange_CUSTOMER EntitlementChange_OperatorType = 3
	// Reseller user.
	EntitlementChange_RESELLER EntitlementChange_OperatorType = 4
)

// Enum value maps for EntitlementChange_OperatorType.
var (
	EntitlementChange_OperatorType_name = map[int32]string{
		0: "OPERATOR_TYPE_UNSPECIFIED",
		1: "CUSTOMER_SERVICE_REPRESENTATIVE",
		2: "SYSTEM",
		3: "CUSTOMER",
		4: "RESELLER",
	}
	EntitlementChange_OperatorType_value = map[string]int32{
		"OPERATOR_TYPE_UNSPECIFIED":       0,
		"CUSTOMER_SERVICE_REPRESENTATIVE": 1,
		"SYSTEM":                          2,
		"CUSTOMER":                        3,
		"RESELLER":                        4,
	}
)

func (x EntitlementChange_OperatorType) Enum() *EntitlementChange_OperatorType {
	p := new(EntitlementChange_OperatorType)
	*p = x
	return p
}

func (x EntitlementChange_OperatorType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EntitlementChange_OperatorType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_channel_v1_entitlement_changes_proto_enumTypes[1].Descriptor()
}

func (EntitlementChange_OperatorType) Type() protoreflect.EnumType {
	return &file_google_cloud_channel_v1_entitlement_changes_proto_enumTypes[1]
}

func (x EntitlementChange_OperatorType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EntitlementChange_OperatorType.Descriptor instead.
func (EntitlementChange_OperatorType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_channel_v1_entitlement_changes_proto_rawDescGZIP(), []int{0, 1}
}

// Cancellation reason for the entitlement
type EntitlementChange_CancellationReason int32

const (
	// Not used.
	EntitlementChange_CANCELLATION_REASON_UNSPECIFIED EntitlementChange_CancellationReason = 0
	// Reseller triggered a cancellation of the service.
	EntitlementChange_SERVICE_TERMINATED EntitlementChange_CancellationReason = 1
	// Relationship between the reseller and customer has ended due to a
	// transfer.
	EntitlementChange_RELATIONSHIP_ENDED EntitlementChange_CancellationReason = 2
	// Entitlement transferred away from reseller while still keeping other
	// entitlement(s) with the reseller.
	EntitlementChange_PARTIAL_TRANSFER EntitlementChange_CancellationReason = 3
)

// Enum value maps for EntitlementChange_CancellationReason.
var (
	EntitlementChange_CancellationReason_name = map[int32]string{
		0: "CANCELLATION_REASON_UNSPECIFIED",
		1: "SERVICE_TERMINATED",
		2: "RELATIONSHIP_ENDED",
		3: "PARTIAL_TRANSFER",
	}
	EntitlementChange_CancellationReason_value = map[string]int32{
		"CANCELLATION_REASON_UNSPECIFIED": 0,
		"SERVICE_TERMINATED":              1,
		"RELATIONSHIP_ENDED":              2,
		"PARTIAL_TRANSFER":                3,
	}
)

func (x EntitlementChange_CancellationReason) Enum() *EntitlementChange_CancellationReason {
	p := new(EntitlementChange_CancellationReason)
	*p = x
	return p
}

func (x EntitlementChange_CancellationReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EntitlementChange_CancellationReason) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_channel_v1_entitlement_changes_proto_enumTypes[2].Descriptor()
}

func (EntitlementChange_CancellationReason) Type() protoreflect.EnumType {
	return &file_google_cloud_channel_v1_entitlement_changes_proto_enumTypes[2]
}

func (x EntitlementChange_CancellationReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EntitlementChange_CancellationReason.Descriptor instead.
func (EntitlementChange_CancellationReason) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_channel_v1_entitlement_changes_proto_rawDescGZIP(), []int{0, 2}
}

// The Entitlement's activation reason
type EntitlementChange_ActivationReason int32

const (
	// Not used.
	EntitlementChange_ACTIVATION_REASON_UNSPECIFIED EntitlementChange_ActivationReason = 0
	// Reseller reactivated a suspended Entitlement.
	EntitlementChange_RESELLER_REVOKED_SUSPENSION EntitlementChange_ActivationReason = 1
	// Customer accepted pending terms of service.
	EntitlementChange_CUSTOMER_ACCEPTED_PENDING_TOS EntitlementChange_ActivationReason = 2
	// Reseller updated the renewal settings on an entitlement that was
	// suspended due to cancellation, and this update reactivated the
	// entitlement.
	EntitlementChange_RENEWAL_SETTINGS_CHANGED EntitlementChange_ActivationReason = 3
	// Other reasons (Activated temporarily for cancellation, added a payment
	// plan to a trial entitlement, etc.)
	EntitlementChange_OTHER_ACTIVATION_REASON EntitlementChange_ActivationReason = 100
)

// Enum value maps for EntitlementChange_ActivationReason.
var (
	EntitlementChange_ActivationReason_name = map[int32]string{
		0:   "ACTIVATION_REASON_UNSPECIFIED",
		1:   "RESELLER_REVOKED_SUSPENSION",
		2:   "CUSTOMER_ACCEPTED_PENDING_TOS",
		3:   "RENEWAL_SETTINGS_CHANGED",
		100: "OTHER_ACTIVATION_REASON",
	}
	EntitlementChange_ActivationReason_value = map[string]int32{
		"ACTIVATION_REASON_UNSPECIFIED": 0,
		"RESELLER_REVOKED_SUSPENSION":   1,
		"CUSTOMER_ACCEPTED_PENDING_TOS": 2,
		"RENEWAL_SETTINGS_CHANGED":      3,
		"OTHER_ACTIVATION_REASON":       100,
	}
)

func (x EntitlementChange_ActivationReason) Enum() *EntitlementChange_ActivationReason {
	p := new(EntitlementChange_ActivationReason)
	*p = x
	return p
}

func (x EntitlementChange_ActivationReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EntitlementChange_ActivationReason) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_channel_v1_entitlement_changes_proto_enumTypes[3].Descriptor()
}

func (EntitlementChange_ActivationReason) Type() protoreflect.EnumType {
	return &file_google_cloud_channel_v1_entitlement_changes_proto_enumTypes[3]
}

func (x EntitlementChange_ActivationReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EntitlementChange_ActivationReason.Descriptor instead.
func (EntitlementChange_ActivationReason) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_channel_v1_entitlement_changes_proto_rawDescGZIP(), []int{0, 3}
}

// Change event entry for Entitlement order history
type EntitlementChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The reason the change was made
	//
	// Types that are assignable to ChangeReason:
	//	*EntitlementChange_SuspensionReason
	//	*EntitlementChange_CancellationReason_
	//	*EntitlementChange_ActivationReason_
	//	*EntitlementChange_OtherChangeReason
	ChangeReason isEntitlementChange_ChangeReason `protobuf_oneof:"change_reason"`
	// Required. Resource name of an entitlement in the form:
	// accounts/{account_id}/customers/{customer_id}/entitlements/{entitlement_id}
	Entitlement string `protobuf:"bytes,1,opt,name=entitlement,proto3" json:"entitlement,omitempty"`
	// Required. Resource name of the Offer at the time of change.
	// Takes the form: accounts/{account_id}/offers/{offer_id}.
	Offer string `protobuf:"bytes,2,opt,name=offer,proto3" json:"offer,omitempty"`
	// Service provisioned for an Entitlement.
	ProvisionedService *ProvisionedService `protobuf:"bytes,3,opt,name=provisioned_service,json=provisionedService,proto3" json:"provisioned_service,omitempty"`
	// The change action type.
	ChangeType EntitlementChange_ChangeType `protobuf:"varint,4,opt,name=change_type,json=changeType,proto3,enum=google.cloud.channel.v1.EntitlementChange_ChangeType" json:"change_type,omitempty"`
	// The submitted time of the change.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Operator type responsible for the change.
	OperatorType EntitlementChange_OperatorType `protobuf:"varint,6,opt,name=operator_type,json=operatorType,proto3,enum=google.cloud.channel.v1.EntitlementChange_OperatorType" json:"operator_type,omitempty"`
	// Extended parameters, such as:
	// purchase_order_number, gcp_details;
	// internal_correlation_id, long_running_operation_id, order_id;
	// etc.
	Parameters []*Parameter `protobuf:"bytes,8,rep,name=parameters,proto3" json:"parameters,omitempty"`
	// Human-readable identifier that shows what operator made a change.
	// When the operator_type is RESELLER, this is the user's email address.
	// For all other operator types, this is empty.
	Operator string `protobuf:"bytes,12,opt,name=operator,proto3" json:"operator,omitempty"`
}

func (x *EntitlementChange) Reset() {
	*x = EntitlementChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_channel_v1_entitlement_changes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntitlementChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntitlementChange) ProtoMessage() {}

func (x *EntitlementChange) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_channel_v1_entitlement_changes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntitlementChange.ProtoReflect.Descriptor instead.
func (*EntitlementChange) Descriptor() ([]byte, []int) {
	return file_google_cloud_channel_v1_entitlement_changes_proto_rawDescGZIP(), []int{0}
}

func (m *EntitlementChange) GetChangeReason() isEntitlementChange_ChangeReason {
	if m != nil {
		return m.ChangeReason
	}
	return nil
}

func (x *EntitlementChange) GetSuspensionReason() Entitlement_SuspensionReason {
	if x, ok := x.GetChangeReason().(*EntitlementChange_SuspensionReason); ok {
		return x.SuspensionReason
	}
	return Entitlement_SUSPENSION_REASON_UNSPECIFIED
}

func (x *EntitlementChange) GetCancellationReason() EntitlementChange_CancellationReason {
	if x, ok := x.GetChangeReason().(*EntitlementChange_CancellationReason_); ok {
		return x.CancellationReason
	}
	return EntitlementChange_CANCELLATION_REASON_UNSPECIFIED
}

func (x *EntitlementChange) GetActivationReason() EntitlementChange_ActivationReason {
	if x, ok := x.GetChangeReason().(*EntitlementChange_ActivationReason_); ok {
		return x.ActivationReason
	}
	return EntitlementChange_ACTIVATION_REASON_UNSPECIFIED
}

func (x *EntitlementChange) GetOtherChangeReason() string {
	if x, ok := x.GetChangeReason().(*EntitlementChange_OtherChangeReason); ok {
		return x.OtherChangeReason
	}
	return ""
}

func (x *EntitlementChange) GetEntitlement() string {
	if x != nil {
		return x.Entitlement
	}
	return ""
}

func (x *EntitlementChange) GetOffer() string {
	if x != nil {
		return x.Offer
	}
	return ""
}

func (x *EntitlementChange) GetProvisionedService() *ProvisionedService {
	if x != nil {
		return x.ProvisionedService
	}
	return nil
}

func (x *EntitlementChange) GetChangeType() EntitlementChange_ChangeType {
	if x != nil {
		return x.ChangeType
	}
	return EntitlementChange_CHANGE_TYPE_UNSPECIFIED
}

func (x *EntitlementChange) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *EntitlementChange) GetOperatorType() EntitlementChange_OperatorType {
	if x != nil {
		return x.OperatorType
	}
	return EntitlementChange_OPERATOR_TYPE_UNSPECIFIED
}

func (x *EntitlementChange) GetParameters() []*Parameter {
	if x != nil {
		return x.Parameters
	}
	return nil
}

func (x *EntitlementChange) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

type isEntitlementChange_ChangeReason interface {
	isEntitlementChange_ChangeReason()
}

type EntitlementChange_SuspensionReason struct {
	// Suspension reason for the Entitlement.
	SuspensionReason Entitlement_SuspensionReason `protobuf:"varint,9,opt,name=suspension_reason,json=suspensionReason,proto3,enum=google.cloud.channel.v1.Entitlement_SuspensionReason,oneof"`
}

type EntitlementChange_CancellationReason_ struct {
	// Cancellation reason for the Entitlement.
	CancellationReason EntitlementChange_CancellationReason `protobuf:"varint,10,opt,name=cancellation_reason,json=cancellationReason,proto3,enum=google.cloud.channel.v1.EntitlementChange_CancellationReason,oneof"`
}

type EntitlementChange_ActivationReason_ struct {
	// The Entitlement's activation reason
	ActivationReason EntitlementChange_ActivationReason `protobuf:"varint,11,opt,name=activation_reason,json=activationReason,proto3,enum=google.cloud.channel.v1.EntitlementChange_ActivationReason,oneof"`
}

type EntitlementChange_OtherChangeReason struct {
	// e.g. purchase_number change reason, entered by CRS.
	OtherChangeReason string `protobuf:"bytes,100,opt,name=other_change_reason,json=otherChangeReason,proto3,oneof"`
}

func (*EntitlementChange_SuspensionReason) isEntitlementChange_ChangeReason() {}

func (*EntitlementChange_CancellationReason_) isEntitlementChange_ChangeReason() {}

func (*EntitlementChange_ActivationReason_) isEntitlementChange_ChangeReason() {}

func (*EntitlementChange_OtherChangeReason) isEntitlementChange_ChangeReason() {}

var File_google_cloud_channel_v1_entitlement_changes_proto protoreflect.FileDescriptor

var file_google_cloud_channel_v1_entitlement_changes_proto_rawDesc = []byte{
	0x0a, 0x31, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x76,
	0x31, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xed, 0x0d, 0x0a, 0x11, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x64, 0x0a, 0x11, 0x73,
	0x75, 0x73, 0x70, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x75, 0x73,
	0x70, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x48, 0x00, 0x52,
	0x10, 0x73, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x12, 0x70, 0x0a, 0x13, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x48, 0x00, 0x52,
	0x12, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x12, 0x6a, 0x0a, 0x11, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x10, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12,
	0x30, 0x0a, 0x13, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x11,
	0x6f, 0x74, 0x68, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x12, 0x51, 0x0a, 0x0b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2f, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x29, 0x0a, 0x27,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x3f, 0x0a, 0x05, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x29, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x05,
	0x6f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x5c, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x12, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0a, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x5c, 0x0a, 0x0d, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x42, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x52, 0x0a,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x22, 0xd7, 0x02, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12,
	0x17, 0x0a, 0x13, 0x50, 0x52, 0x49, 0x43, 0x45, 0x5f, 0x50, 0x4c, 0x41, 0x4e, 0x5f, 0x53, 0x57,
	0x49, 0x54, 0x43, 0x48, 0x45, 0x44, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x4f, 0x4d, 0x4d,
	0x49, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0x04,
	0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x4e, 0x45, 0x57, 0x45, 0x44, 0x10, 0x05, 0x12, 0x0d, 0x0a,
	0x09, 0x53, 0x55, 0x53, 0x50, 0x45, 0x4e, 0x44, 0x45, 0x44, 0x10, 0x06, 0x12, 0x0d, 0x0a, 0x09,
	0x41, 0x43, 0x54, 0x49, 0x56, 0x41, 0x54, 0x45, 0x44, 0x10, 0x07, 0x12, 0x0d, 0x0a, 0x09, 0x43,
	0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x08, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x4b,
	0x55, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0x09, 0x12, 0x1b, 0x0a, 0x17, 0x52,
	0x45, 0x4e, 0x45, 0x57, 0x41, 0x4c, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x43,
	0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0x0a, 0x12, 0x1d, 0x0a, 0x19, 0x50, 0x41, 0x49, 0x44,
	0x5f, 0x53, 0x55, 0x42, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54,
	0x41, 0x52, 0x54, 0x45, 0x44, 0x10, 0x0b, 0x12, 0x17, 0x0a, 0x13, 0x4c, 0x49, 0x43, 0x45, 0x4e,
	0x53, 0x45, 0x5f, 0x43, 0x41, 0x50, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0x0c,
	0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x55, 0x53, 0x50, 0x45, 0x4e, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x44,
	0x45, 0x54, 0x41, 0x49, 0x4c, 0x53, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0x0d,
	0x12, 0x1b, 0x0a, 0x17, 0x54, 0x52, 0x49, 0x41, 0x4c, 0x5f, 0x45, 0x4e, 0x44, 0x5f, 0x44, 0x41,
	0x54, 0x45, 0x5f, 0x45, 0x58, 0x54, 0x45, 0x4e, 0x44, 0x45, 0x44, 0x10, 0x0e, 0x12, 0x11, 0x0a,
	0x0d, 0x54, 0x52, 0x49, 0x41, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10, 0x0f,
	0x22, 0x7a, 0x0a, 0x0c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1d, 0x0a, 0x19, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x23, 0x0a, 0x1f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x5f, 0x53, 0x45, 0x52, 0x56,
	0x49, 0x43, 0x45, 0x5f, 0x52, 0x45, 0x50, 0x52, 0x45, 0x53, 0x45, 0x4e, 0x54, 0x41, 0x54, 0x49,
	0x56, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x10, 0x02,
	0x12, 0x0c, 0x0a, 0x08, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x10, 0x03, 0x12, 0x0c,
	0x0a, 0x08, 0x52, 0x45, 0x53, 0x45, 0x4c, 0x4c, 0x45, 0x52, 0x10, 0x04, 0x22, 0x7f, 0x0a, 0x12,
	0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x1f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x45, 0x52, 0x56, 0x49,
	0x43, 0x45, 0x5f, 0x54, 0x45, 0x52, 0x4d, 0x49, 0x4e, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12,
	0x16, 0x0a, 0x12, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x48, 0x49, 0x50, 0x5f,
	0x45, 0x4e, 0x44, 0x45, 0x44, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x41, 0x52, 0x54, 0x49,
	0x41, 0x4c, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x10, 0x03, 0x22, 0xb4, 0x01,
	0x0a, 0x10, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x1d, 0x41, 0x43, 0x54, 0x49, 0x56, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1f, 0x0a, 0x1b, 0x52, 0x45, 0x53, 0x45, 0x4c, 0x4c, 0x45,
	0x52, 0x5f, 0x52, 0x45, 0x56, 0x4f, 0x4b, 0x45, 0x44, 0x5f, 0x53, 0x55, 0x53, 0x50, 0x45, 0x4e,
	0x53, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x21, 0x0a, 0x1d, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d,
	0x45, 0x52, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44, 0x5f, 0x50, 0x45, 0x4e, 0x44,
	0x49, 0x4e, 0x47, 0x5f, 0x54, 0x4f, 0x53, 0x10, 0x02, 0x12, 0x1c, 0x0a, 0x18, 0x52, 0x45, 0x4e,
	0x45, 0x57, 0x41, 0x4c, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x53, 0x5f, 0x43, 0x48,
	0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0x03, 0x12, 0x1b, 0x0a, 0x17, 0x4f, 0x54, 0x48, 0x45, 0x52,
	0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x41, 0x53,
	0x4f, 0x4e, 0x10, 0x64, 0x42, 0x0f, 0x0a, 0x0d, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x72,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x42, 0x6f, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x2e, 0x76, 0x31, 0x42, 0x17, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x35, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x61, 0x70, 0x69,
	0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x70, 0x62, 0x3b, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_channel_v1_entitlement_changes_proto_rawDescOnce sync.Once
	file_google_cloud_channel_v1_entitlement_changes_proto_rawDescData = file_google_cloud_channel_v1_entitlement_changes_proto_rawDesc
)

func file_google_cloud_channel_v1_entitlement_changes_proto_rawDescGZIP() []byte {
	file_google_cloud_channel_v1_entitlement_changes_proto_rawDescOnce.Do(func() {
		file_google_cloud_channel_v1_entitlement_changes_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_channel_v1_entitlement_changes_proto_rawDescData)
	})
	return file_google_cloud_channel_v1_entitlement_changes_proto_rawDescData
}

var file_google_cloud_channel_v1_entitlement_changes_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_google_cloud_channel_v1_entitlement_changes_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_channel_v1_entitlement_changes_proto_goTypes = []interface{}{
	(EntitlementChange_ChangeType)(0),         // 0: google.cloud.channel.v1.EntitlementChange.ChangeType
	(EntitlementChange_OperatorType)(0),       // 1: google.cloud.channel.v1.EntitlementChange.OperatorType
	(EntitlementChange_CancellationReason)(0), // 2: google.cloud.channel.v1.EntitlementChange.CancellationReason
	(EntitlementChange_ActivationReason)(0),   // 3: google.cloud.channel.v1.EntitlementChange.ActivationReason
	(*EntitlementChange)(nil),                 // 4: google.cloud.channel.v1.EntitlementChange
	(Entitlement_SuspensionReason)(0),         // 5: google.cloud.channel.v1.Entitlement.SuspensionReason
	(*ProvisionedService)(nil),                // 6: google.cloud.channel.v1.ProvisionedService
	(*timestamppb.Timestamp)(nil),             // 7: google.protobuf.Timestamp
	(*Parameter)(nil),                         // 8: google.cloud.channel.v1.Parameter
}
var file_google_cloud_channel_v1_entitlement_changes_proto_depIdxs = []int32{
	5, // 0: google.cloud.channel.v1.EntitlementChange.suspension_reason:type_name -> google.cloud.channel.v1.Entitlement.SuspensionReason
	2, // 1: google.cloud.channel.v1.EntitlementChange.cancellation_reason:type_name -> google.cloud.channel.v1.EntitlementChange.CancellationReason
	3, // 2: google.cloud.channel.v1.EntitlementChange.activation_reason:type_name -> google.cloud.channel.v1.EntitlementChange.ActivationReason
	6, // 3: google.cloud.channel.v1.EntitlementChange.provisioned_service:type_name -> google.cloud.channel.v1.ProvisionedService
	0, // 4: google.cloud.channel.v1.EntitlementChange.change_type:type_name -> google.cloud.channel.v1.EntitlementChange.ChangeType
	7, // 5: google.cloud.channel.v1.EntitlementChange.create_time:type_name -> google.protobuf.Timestamp
	1, // 6: google.cloud.channel.v1.EntitlementChange.operator_type:type_name -> google.cloud.channel.v1.EntitlementChange.OperatorType
	8, // 7: google.cloud.channel.v1.EntitlementChange.parameters:type_name -> google.cloud.channel.v1.Parameter
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_google_cloud_channel_v1_entitlement_changes_proto_init() }
func file_google_cloud_channel_v1_entitlement_changes_proto_init() {
	if File_google_cloud_channel_v1_entitlement_changes_proto != nil {
		return
	}
	file_google_cloud_channel_v1_entitlements_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_channel_v1_entitlement_changes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntitlementChange); i {
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
	file_google_cloud_channel_v1_entitlement_changes_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*EntitlementChange_SuspensionReason)(nil),
		(*EntitlementChange_CancellationReason_)(nil),
		(*EntitlementChange_ActivationReason_)(nil),
		(*EntitlementChange_OtherChangeReason)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_channel_v1_entitlement_changes_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_channel_v1_entitlement_changes_proto_goTypes,
		DependencyIndexes: file_google_cloud_channel_v1_entitlement_changes_proto_depIdxs,
		EnumInfos:         file_google_cloud_channel_v1_entitlement_changes_proto_enumTypes,
		MessageInfos:      file_google_cloud_channel_v1_entitlement_changes_proto_msgTypes,
	}.Build()
	File_google_cloud_channel_v1_entitlement_changes_proto = out.File
	file_google_cloud_channel_v1_entitlement_changes_proto_rawDesc = nil
	file_google_cloud_channel_v1_entitlement_changes_proto_goTypes = nil
	file_google_cloud_channel_v1_entitlement_changes_proto_depIdxs = nil
}
