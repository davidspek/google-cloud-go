// Copyright 2023 Google LLC
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
// source: google/cloud/aiplatform/v1/evaluated_annotation.proto

package aiplatformpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Describes the type of the EvaluatedAnnotation. The type is determined
type EvaluatedAnnotation_EvaluatedAnnotationType int32

const (
	// Invalid value.
	EvaluatedAnnotation_EVALUATED_ANNOTATION_TYPE_UNSPECIFIED EvaluatedAnnotation_EvaluatedAnnotationType = 0
	// The EvaluatedAnnotation is a true positive. It has a prediction created
	// by the Model and a ground truth Annotation which the prediction matches.
	EvaluatedAnnotation_TRUE_POSITIVE EvaluatedAnnotation_EvaluatedAnnotationType = 1
	// The EvaluatedAnnotation is false positive. It has a prediction created by
	// the Model which does not match any ground truth annotation.
	EvaluatedAnnotation_FALSE_POSITIVE EvaluatedAnnotation_EvaluatedAnnotationType = 2
	// The EvaluatedAnnotation is false negative. It has a ground truth
	// annotation which is not matched by any of the model created predictions.
	EvaluatedAnnotation_FALSE_NEGATIVE EvaluatedAnnotation_EvaluatedAnnotationType = 3
)

// Enum value maps for EvaluatedAnnotation_EvaluatedAnnotationType.
var (
	EvaluatedAnnotation_EvaluatedAnnotationType_name = map[int32]string{
		0: "EVALUATED_ANNOTATION_TYPE_UNSPECIFIED",
		1: "TRUE_POSITIVE",
		2: "FALSE_POSITIVE",
		3: "FALSE_NEGATIVE",
	}
	EvaluatedAnnotation_EvaluatedAnnotationType_value = map[string]int32{
		"EVALUATED_ANNOTATION_TYPE_UNSPECIFIED": 0,
		"TRUE_POSITIVE":                         1,
		"FALSE_POSITIVE":                        2,
		"FALSE_NEGATIVE":                        3,
	}
)

func (x EvaluatedAnnotation_EvaluatedAnnotationType) Enum() *EvaluatedAnnotation_EvaluatedAnnotationType {
	p := new(EvaluatedAnnotation_EvaluatedAnnotationType)
	*p = x
	return p
}

func (x EvaluatedAnnotation_EvaluatedAnnotationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EvaluatedAnnotation_EvaluatedAnnotationType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_aiplatform_v1_evaluated_annotation_proto_enumTypes[0].Descriptor()
}

func (EvaluatedAnnotation_EvaluatedAnnotationType) Type() protoreflect.EnumType {
	return &file_google_cloud_aiplatform_v1_evaluated_annotation_proto_enumTypes[0]
}

func (x EvaluatedAnnotation_EvaluatedAnnotationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EvaluatedAnnotation_EvaluatedAnnotationType.Descriptor instead.
func (EvaluatedAnnotation_EvaluatedAnnotationType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_evaluated_annotation_proto_rawDescGZIP(), []int{0, 0}
}

// The query type used for finding the attributed items.
type ErrorAnalysisAnnotation_QueryType int32

const (
	// Unspecified query type for model error analysis.
	ErrorAnalysisAnnotation_QUERY_TYPE_UNSPECIFIED ErrorAnalysisAnnotation_QueryType = 0
	// Query similar samples across all classes in the dataset.
	ErrorAnalysisAnnotation_ALL_SIMILAR ErrorAnalysisAnnotation_QueryType = 1
	// Query similar samples from the same class of the input sample.
	ErrorAnalysisAnnotation_SAME_CLASS_SIMILAR ErrorAnalysisAnnotation_QueryType = 2
	// Query dissimilar samples from the same class of the input sample.
	ErrorAnalysisAnnotation_SAME_CLASS_DISSIMILAR ErrorAnalysisAnnotation_QueryType = 3
)

// Enum value maps for ErrorAnalysisAnnotation_QueryType.
var (
	ErrorAnalysisAnnotation_QueryType_name = map[int32]string{
		0: "QUERY_TYPE_UNSPECIFIED",
		1: "ALL_SIMILAR",
		2: "SAME_CLASS_SIMILAR",
		3: "SAME_CLASS_DISSIMILAR",
	}
	ErrorAnalysisAnnotation_QueryType_value = map[string]int32{
		"QUERY_TYPE_UNSPECIFIED": 0,
		"ALL_SIMILAR":            1,
		"SAME_CLASS_SIMILAR":     2,
		"SAME_CLASS_DISSIMILAR":  3,
	}
)

func (x ErrorAnalysisAnnotation_QueryType) Enum() *ErrorAnalysisAnnotation_QueryType {
	p := new(ErrorAnalysisAnnotation_QueryType)
	*p = x
	return p
}

func (x ErrorAnalysisAnnotation_QueryType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorAnalysisAnnotation_QueryType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_aiplatform_v1_evaluated_annotation_proto_enumTypes[1].Descriptor()
}

func (ErrorAnalysisAnnotation_QueryType) Type() protoreflect.EnumType {
	return &file_google_cloud_aiplatform_v1_evaluated_annotation_proto_enumTypes[1]
}

func (x ErrorAnalysisAnnotation_QueryType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorAnalysisAnnotation_QueryType.Descriptor instead.
func (ErrorAnalysisAnnotation_QueryType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_evaluated_annotation_proto_rawDescGZIP(), []int{2, 0}
}

// True positive, false positive, or false negative.
//
// EvaluatedAnnotation is only available under ModelEvaluationSlice with slice
// of `annotationSpec` dimension.
type EvaluatedAnnotation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Type of the EvaluatedAnnotation.
	Type EvaluatedAnnotation_EvaluatedAnnotationType `protobuf:"varint,1,opt,name=type,proto3,enum=google.cloud.aiplatform.v1.EvaluatedAnnotation_EvaluatedAnnotationType" json:"type,omitempty"`
	// Output only. The model predicted annotations.
	//
	// For true positive, there is one and only one prediction, which matches the
	// only one ground truth annotation in
	// [ground_truths][google.cloud.aiplatform.v1.EvaluatedAnnotation.ground_truths].
	//
	// For false positive, there is one and only one prediction, which doesn't
	// match any ground truth annotation of the corresponding
	// [data_item_view_id][EvaluatedAnnotation.data_item_view_id].
	//
	// For false negative, there are zero or more predictions which are similar to
	// the only ground truth annotation in
	// [ground_truths][google.cloud.aiplatform.v1.EvaluatedAnnotation.ground_truths]
	// but not enough for a match.
	//
	// The schema of the prediction is stored in
	// [ModelEvaluation.annotation_schema_uri][google.cloud.aiplatform.v1.ModelEvaluation.annotation_schema_uri]
	Predictions []*structpb.Value `protobuf:"bytes,2,rep,name=predictions,proto3" json:"predictions,omitempty"`
	// Output only. The ground truth Annotations, i.e. the Annotations that exist
	// in the test data the Model is evaluated on.
	//
	// For true positive, there is one and only one ground truth annotation, which
	// matches the only prediction in
	// [predictions][google.cloud.aiplatform.v1.EvaluatedAnnotation.predictions].
	//
	// For false positive, there are zero or more ground truth annotations that
	// are similar to the only prediction in
	// [predictions][google.cloud.aiplatform.v1.EvaluatedAnnotation.predictions],
	// but not enough for a match.
	//
	// For false negative, there is one and only one ground truth annotation,
	// which doesn't match any predictions created by the model.
	//
	// The schema of the ground truth is stored in
	// [ModelEvaluation.annotation_schema_uri][google.cloud.aiplatform.v1.ModelEvaluation.annotation_schema_uri]
	GroundTruths []*structpb.Value `protobuf:"bytes,3,rep,name=ground_truths,json=groundTruths,proto3" json:"ground_truths,omitempty"`
	// Output only. The data item payload that the Model predicted this
	// EvaluatedAnnotation on.
	DataItemPayload *structpb.Value `protobuf:"bytes,5,opt,name=data_item_payload,json=dataItemPayload,proto3" json:"data_item_payload,omitempty"`
	// Output only. ID of the EvaluatedDataItemView under the same ancestor
	// ModelEvaluation. The EvaluatedDataItemView consists of all ground truths
	// and predictions on
	// [data_item_payload][google.cloud.aiplatform.v1.EvaluatedAnnotation.data_item_payload].
	//
	// Can be passed in
	// [GetEvaluatedDataItemView's][ModelService.GetEvaluatedDataItemView][]
	// [id][GetEvaluatedDataItemViewRequest.id].
	EvaluatedDataItemViewId string `protobuf:"bytes,6,opt,name=evaluated_data_item_view_id,json=evaluatedDataItemViewId,proto3" json:"evaluated_data_item_view_id,omitempty"`
	// Explanations of
	// [predictions][google.cloud.aiplatform.v1.EvaluatedAnnotation.predictions].
	// Each element of the explanations indicates the explanation for one
	// explanation Method.
	//
	// The attributions list in the
	// [EvaluatedAnnotationExplanation.explanation][google.cloud.aiplatform.v1.EvaluatedAnnotationExplanation.explanation]
	// object corresponds to the
	// [predictions][google.cloud.aiplatform.v1.EvaluatedAnnotation.predictions]
	// list. For example, the second element in the attributions list explains the
	// second element in the predictions list.
	Explanations []*EvaluatedAnnotationExplanation `protobuf:"bytes,8,rep,name=explanations,proto3" json:"explanations,omitempty"`
	// Annotations of model error analysis results.
	ErrorAnalysisAnnotations []*ErrorAnalysisAnnotation `protobuf:"bytes,9,rep,name=error_analysis_annotations,json=errorAnalysisAnnotations,proto3" json:"error_analysis_annotations,omitempty"`
}

func (x *EvaluatedAnnotation) Reset() {
	*x = EvaluatedAnnotation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_evaluated_annotation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvaluatedAnnotation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvaluatedAnnotation) ProtoMessage() {}

func (x *EvaluatedAnnotation) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_evaluated_annotation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvaluatedAnnotation.ProtoReflect.Descriptor instead.
func (*EvaluatedAnnotation) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_evaluated_annotation_proto_rawDescGZIP(), []int{0}
}

func (x *EvaluatedAnnotation) GetType() EvaluatedAnnotation_EvaluatedAnnotationType {
	if x != nil {
		return x.Type
	}
	return EvaluatedAnnotation_EVALUATED_ANNOTATION_TYPE_UNSPECIFIED
}

func (x *EvaluatedAnnotation) GetPredictions() []*structpb.Value {
	if x != nil {
		return x.Predictions
	}
	return nil
}

func (x *EvaluatedAnnotation) GetGroundTruths() []*structpb.Value {
	if x != nil {
		return x.GroundTruths
	}
	return nil
}

func (x *EvaluatedAnnotation) GetDataItemPayload() *structpb.Value {
	if x != nil {
		return x.DataItemPayload
	}
	return nil
}

func (x *EvaluatedAnnotation) GetEvaluatedDataItemViewId() string {
	if x != nil {
		return x.EvaluatedDataItemViewId
	}
	return ""
}

func (x *EvaluatedAnnotation) GetExplanations() []*EvaluatedAnnotationExplanation {
	if x != nil {
		return x.Explanations
	}
	return nil
}

func (x *EvaluatedAnnotation) GetErrorAnalysisAnnotations() []*ErrorAnalysisAnnotation {
	if x != nil {
		return x.ErrorAnalysisAnnotations
	}
	return nil
}

// Explanation result of the prediction produced by the Model.
type EvaluatedAnnotationExplanation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Explanation type.
	//
	// For AutoML Image Classification models, possible values are:
	//
	//   * `image-integrated-gradients`
	//   * `image-xrai`
	ExplanationType string `protobuf:"bytes,1,opt,name=explanation_type,json=explanationType,proto3" json:"explanation_type,omitempty"`
	// Explanation attribution response details.
	Explanation *Explanation `protobuf:"bytes,2,opt,name=explanation,proto3" json:"explanation,omitempty"`
}

func (x *EvaluatedAnnotationExplanation) Reset() {
	*x = EvaluatedAnnotationExplanation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_evaluated_annotation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvaluatedAnnotationExplanation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvaluatedAnnotationExplanation) ProtoMessage() {}

func (x *EvaluatedAnnotationExplanation) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_evaluated_annotation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvaluatedAnnotationExplanation.ProtoReflect.Descriptor instead.
func (*EvaluatedAnnotationExplanation) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_evaluated_annotation_proto_rawDescGZIP(), []int{1}
}

func (x *EvaluatedAnnotationExplanation) GetExplanationType() string {
	if x != nil {
		return x.ExplanationType
	}
	return ""
}

func (x *EvaluatedAnnotationExplanation) GetExplanation() *Explanation {
	if x != nil {
		return x.Explanation
	}
	return nil
}

// Model error analysis for each annotation.
type ErrorAnalysisAnnotation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Attributed items for a given annotation, typically representing neighbors
	// from the training sets constrained by the query type.
	AttributedItems []*ErrorAnalysisAnnotation_AttributedItem `protobuf:"bytes,1,rep,name=attributed_items,json=attributedItems,proto3" json:"attributed_items,omitempty"`
	// The query type used for finding the attributed items.
	QueryType ErrorAnalysisAnnotation_QueryType `protobuf:"varint,2,opt,name=query_type,json=queryType,proto3,enum=google.cloud.aiplatform.v1.ErrorAnalysisAnnotation_QueryType" json:"query_type,omitempty"`
	// The outlier score of this annotated item. Usually defined as the min of all
	// distances from attributed items.
	OutlierScore float64 `protobuf:"fixed64,3,opt,name=outlier_score,json=outlierScore,proto3" json:"outlier_score,omitempty"`
	// The threshold used to determine if this annotation is an outlier or not.
	OutlierThreshold float64 `protobuf:"fixed64,4,opt,name=outlier_threshold,json=outlierThreshold,proto3" json:"outlier_threshold,omitempty"`
}

func (x *ErrorAnalysisAnnotation) Reset() {
	*x = ErrorAnalysisAnnotation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_evaluated_annotation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorAnalysisAnnotation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorAnalysisAnnotation) ProtoMessage() {}

func (x *ErrorAnalysisAnnotation) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_evaluated_annotation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorAnalysisAnnotation.ProtoReflect.Descriptor instead.
func (*ErrorAnalysisAnnotation) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_evaluated_annotation_proto_rawDescGZIP(), []int{2}
}

func (x *ErrorAnalysisAnnotation) GetAttributedItems() []*ErrorAnalysisAnnotation_AttributedItem {
	if x != nil {
		return x.AttributedItems
	}
	return nil
}

func (x *ErrorAnalysisAnnotation) GetQueryType() ErrorAnalysisAnnotation_QueryType {
	if x != nil {
		return x.QueryType
	}
	return ErrorAnalysisAnnotation_QUERY_TYPE_UNSPECIFIED
}

func (x *ErrorAnalysisAnnotation) GetOutlierScore() float64 {
	if x != nil {
		return x.OutlierScore
	}
	return 0
}

func (x *ErrorAnalysisAnnotation) GetOutlierThreshold() float64 {
	if x != nil {
		return x.OutlierThreshold
	}
	return 0
}

// Attributed items for a given annotation, typically representing neighbors
// from the training sets constrained by the query type.
type ErrorAnalysisAnnotation_AttributedItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique ID for each annotation. Used by FE to allocate the annotation
	// in DB.
	AnnotationResourceName string `protobuf:"bytes,1,opt,name=annotation_resource_name,json=annotationResourceName,proto3" json:"annotation_resource_name,omitempty"`
	// The distance of this item to the annotation.
	Distance float64 `protobuf:"fixed64,2,opt,name=distance,proto3" json:"distance,omitempty"`
}

func (x *ErrorAnalysisAnnotation_AttributedItem) Reset() {
	*x = ErrorAnalysisAnnotation_AttributedItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_evaluated_annotation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorAnalysisAnnotation_AttributedItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorAnalysisAnnotation_AttributedItem) ProtoMessage() {}

func (x *ErrorAnalysisAnnotation_AttributedItem) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_evaluated_annotation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorAnalysisAnnotation_AttributedItem.ProtoReflect.Descriptor instead.
func (*ErrorAnalysisAnnotation_AttributedItem) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_evaluated_annotation_proto_rawDescGZIP(), []int{2, 0}
}

func (x *ErrorAnalysisAnnotation_AttributedItem) GetAnnotationResourceName() string {
	if x != nil {
		return x.AnnotationResourceName
	}
	return ""
}

func (x *ErrorAnalysisAnnotation_AttributedItem) GetDistance() float64 {
	if x != nil {
		return x.Distance
	}
	return 0
}

var File_google_cloud_aiplatform_v1_evaluated_annotation_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1_evaluated_annotation_proto_rawDesc = []byte{
	0x0a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x61,
	0x6c, 0x75, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x78, 0x70, 0x6c, 0x61, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xd8, 0x05, 0x0a, 0x13, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x64, 0x41, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x60, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x47, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x64, 0x41, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3d, 0x0a, 0x0b, 0x70, 0x72,
	0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0b, 0x70, 0x72,
	0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x40, 0x0a, 0x0d, 0x67, 0x72, 0x6f,
	0x75, 0x6e, 0x64, 0x5f, 0x74, 0x72, 0x75, 0x74, 0x68, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0c, 0x67,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x54, 0x72, 0x75, 0x74, 0x68, 0x73, 0x12, 0x47, 0x0a, 0x11, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x12, 0x41, 0x0a, 0x1b, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x76, 0x69, 0x65, 0x77,
	0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x17,
	0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x49, 0x74, 0x65,
	0x6d, 0x56, 0x69, 0x65, 0x77, 0x49, 0x64, 0x12, 0x5e, 0x0a, 0x0c, 0x65, 0x78, 0x70, 0x6c, 0x61,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78,
	0x70, 0x6c, 0x61, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x65, 0x78, 0x70, 0x6c, 0x61,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x71, 0x0a, 0x1a, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x5f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x41, 0x6e,
	0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x18, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x41,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x7f, 0x0a, 0x17, 0x45, 0x76,
	0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x64, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x29, 0x0a, 0x25, 0x45, 0x56, 0x41, 0x4c, 0x55, 0x41, 0x54,
	0x45, 0x44, 0x5f, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x11, 0x0a, 0x0d, 0x54, 0x52, 0x55, 0x45, 0x5f, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x49, 0x56,
	0x45, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x46, 0x41, 0x4c, 0x53, 0x45, 0x5f, 0x50, 0x4f, 0x53,
	0x49, 0x54, 0x49, 0x56, 0x45, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x46, 0x41, 0x4c, 0x53, 0x45,
	0x5f, 0x4e, 0x45, 0x47, 0x41, 0x54, 0x49, 0x56, 0x45, 0x10, 0x03, 0x22, 0x96, 0x01, 0x0a, 0x1e,
	0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x64, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x45, 0x78, 0x70, 0x6c, 0x61, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29,
	0x0a, 0x10, 0x65, 0x78, 0x70, 0x6c, 0x61, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x78, 0x70, 0x6c, 0x61, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x49, 0x0a, 0x0b, 0x65, 0x78, 0x70,
	0x6c, 0x61, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x6c,
	0x61, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x65, 0x78, 0x70, 0x6c, 0x61, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x8d, 0x04, 0x0a, 0x17, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x41, 0x6e,
	0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x6d, 0x0a, 0x10, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x5f, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x41, 0x6e, 0x61,
	0x6c, 0x79, 0x73, 0x69, 0x73, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0f,
	0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12,
	0x5c, 0x0a, 0x0a, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x3d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x41, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x09, 0x71, 0x75, 0x65, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x6f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x6f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x6f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72, 0x5f, 0x74, 0x68,
	0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x6f,
	0x75, 0x74, 0x6c, 0x69, 0x65, 0x72, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x1a,
	0x66, 0x0a, 0x0e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x38, 0x0a, 0x18, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x16, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x64,
	0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x6b, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x4c, 0x4c, 0x5f, 0x53, 0x49, 0x4d, 0x49, 0x4c, 0x41, 0x52, 0x10,
	0x01, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x41, 0x4d, 0x45, 0x5f, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x5f,
	0x53, 0x49, 0x4d, 0x49, 0x4c, 0x41, 0x52, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x41, 0x4d,
	0x45, 0x5f, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x5f, 0x44, 0x49, 0x53, 0x53, 0x49, 0x4d, 0x49, 0x4c,
	0x41, 0x52, 0x10, 0x03, 0x42, 0xd6, 0x01, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x18, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0x3b, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x70, 0x62, 0xaa, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x1d,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41,
	0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1_evaluated_annotation_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1_evaluated_annotation_proto_rawDescData = file_google_cloud_aiplatform_v1_evaluated_annotation_proto_rawDesc
)

func file_google_cloud_aiplatform_v1_evaluated_annotation_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1_evaluated_annotation_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1_evaluated_annotation_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1_evaluated_annotation_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1_evaluated_annotation_proto_rawDescData
}

var file_google_cloud_aiplatform_v1_evaluated_annotation_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_cloud_aiplatform_v1_evaluated_annotation_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_aiplatform_v1_evaluated_annotation_proto_goTypes = []interface{}{
	(EvaluatedAnnotation_EvaluatedAnnotationType)(0), // 0: google.cloud.aiplatform.v1.EvaluatedAnnotation.EvaluatedAnnotationType
	(ErrorAnalysisAnnotation_QueryType)(0),           // 1: google.cloud.aiplatform.v1.ErrorAnalysisAnnotation.QueryType
	(*EvaluatedAnnotation)(nil),                      // 2: google.cloud.aiplatform.v1.EvaluatedAnnotation
	(*EvaluatedAnnotationExplanation)(nil),           // 3: google.cloud.aiplatform.v1.EvaluatedAnnotationExplanation
	(*ErrorAnalysisAnnotation)(nil),                  // 4: google.cloud.aiplatform.v1.ErrorAnalysisAnnotation
	(*ErrorAnalysisAnnotation_AttributedItem)(nil),   // 5: google.cloud.aiplatform.v1.ErrorAnalysisAnnotation.AttributedItem
	(*structpb.Value)(nil),                           // 6: google.protobuf.Value
	(*Explanation)(nil),                              // 7: google.cloud.aiplatform.v1.Explanation
}
var file_google_cloud_aiplatform_v1_evaluated_annotation_proto_depIdxs = []int32{
	0, // 0: google.cloud.aiplatform.v1.EvaluatedAnnotation.type:type_name -> google.cloud.aiplatform.v1.EvaluatedAnnotation.EvaluatedAnnotationType
	6, // 1: google.cloud.aiplatform.v1.EvaluatedAnnotation.predictions:type_name -> google.protobuf.Value
	6, // 2: google.cloud.aiplatform.v1.EvaluatedAnnotation.ground_truths:type_name -> google.protobuf.Value
	6, // 3: google.cloud.aiplatform.v1.EvaluatedAnnotation.data_item_payload:type_name -> google.protobuf.Value
	3, // 4: google.cloud.aiplatform.v1.EvaluatedAnnotation.explanations:type_name -> google.cloud.aiplatform.v1.EvaluatedAnnotationExplanation
	4, // 5: google.cloud.aiplatform.v1.EvaluatedAnnotation.error_analysis_annotations:type_name -> google.cloud.aiplatform.v1.ErrorAnalysisAnnotation
	7, // 6: google.cloud.aiplatform.v1.EvaluatedAnnotationExplanation.explanation:type_name -> google.cloud.aiplatform.v1.Explanation
	5, // 7: google.cloud.aiplatform.v1.ErrorAnalysisAnnotation.attributed_items:type_name -> google.cloud.aiplatform.v1.ErrorAnalysisAnnotation.AttributedItem
	1, // 8: google.cloud.aiplatform.v1.ErrorAnalysisAnnotation.query_type:type_name -> google.cloud.aiplatform.v1.ErrorAnalysisAnnotation.QueryType
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1_evaluated_annotation_proto_init() }
func file_google_cloud_aiplatform_v1_evaluated_annotation_proto_init() {
	if File_google_cloud_aiplatform_v1_evaluated_annotation_proto != nil {
		return
	}
	file_google_cloud_aiplatform_v1_explanation_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_aiplatform_v1_evaluated_annotation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvaluatedAnnotation); i {
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
		file_google_cloud_aiplatform_v1_evaluated_annotation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvaluatedAnnotationExplanation); i {
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
		file_google_cloud_aiplatform_v1_evaluated_annotation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorAnalysisAnnotation); i {
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
		file_google_cloud_aiplatform_v1_evaluated_annotation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorAnalysisAnnotation_AttributedItem); i {
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
			RawDescriptor: file_google_cloud_aiplatform_v1_evaluated_annotation_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1_evaluated_annotation_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1_evaluated_annotation_proto_depIdxs,
		EnumInfos:         file_google_cloud_aiplatform_v1_evaluated_annotation_proto_enumTypes,
		MessageInfos:      file_google_cloud_aiplatform_v1_evaluated_annotation_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1_evaluated_annotation_proto = out.File
	file_google_cloud_aiplatform_v1_evaluated_annotation_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1_evaluated_annotation_proto_goTypes = nil
	file_google_cloud_aiplatform_v1_evaluated_annotation_proto_depIdxs = nil
}
