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
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: google/cloud/automl/v1beta1/text_sentiment.proto

package automlpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Contains annotation details specific to text sentiment.
type TextSentimentAnnotation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The sentiment with the semantic, as given to the
	// [AutoMl.ImportData][google.cloud.automl.v1beta1.AutoMl.ImportData] when populating the dataset from which the model used
	// for the prediction had been trained.
	// The sentiment values are between 0 and
	// Dataset.text_sentiment_dataset_metadata.sentiment_max (inclusive),
	// with higher value meaning more positive sentiment. They are completely
	// relative, i.e. 0 means least positive sentiment and sentiment_max means
	// the most positive from the sentiments present in the train data. Therefore
	//  e.g. if train data had only negative sentiment, then sentiment_max, would
	// be still negative (although least negative).
	// The sentiment shouldn't be confused with "score" or "magnitude"
	// from the previous Natural Language Sentiment Analysis API.
	Sentiment int32 `protobuf:"varint,1,opt,name=sentiment,proto3" json:"sentiment,omitempty"`
}

func (x *TextSentimentAnnotation) Reset() {
	*x = TextSentimentAnnotation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1beta1_text_sentiment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextSentimentAnnotation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextSentimentAnnotation) ProtoMessage() {}

func (x *TextSentimentAnnotation) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1beta1_text_sentiment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextSentimentAnnotation.ProtoReflect.Descriptor instead.
func (*TextSentimentAnnotation) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1beta1_text_sentiment_proto_rawDescGZIP(), []int{0}
}

func (x *TextSentimentAnnotation) GetSentiment() int32 {
	if x != nil {
		return x.Sentiment
	}
	return 0
}

// Model evaluation metrics for text sentiment problems.
type TextSentimentEvaluationMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Precision.
	Precision float32 `protobuf:"fixed32,1,opt,name=precision,proto3" json:"precision,omitempty"`
	// Output only. Recall.
	Recall float32 `protobuf:"fixed32,2,opt,name=recall,proto3" json:"recall,omitempty"`
	// Output only. The harmonic mean of recall and precision.
	F1Score float32 `protobuf:"fixed32,3,opt,name=f1_score,json=f1Score,proto3" json:"f1_score,omitempty"`
	// Output only. Mean absolute error. Only set for the overall model
	// evaluation, not for evaluation of a single annotation spec.
	MeanAbsoluteError float32 `protobuf:"fixed32,4,opt,name=mean_absolute_error,json=meanAbsoluteError,proto3" json:"mean_absolute_error,omitempty"`
	// Output only. Mean squared error. Only set for the overall model
	// evaluation, not for evaluation of a single annotation spec.
	MeanSquaredError float32 `protobuf:"fixed32,5,opt,name=mean_squared_error,json=meanSquaredError,proto3" json:"mean_squared_error,omitempty"`
	// Output only. Linear weighted kappa. Only set for the overall model
	// evaluation, not for evaluation of a single annotation spec.
	LinearKappa float32 `protobuf:"fixed32,6,opt,name=linear_kappa,json=linearKappa,proto3" json:"linear_kappa,omitempty"`
	// Output only. Quadratic weighted kappa. Only set for the overall model
	// evaluation, not for evaluation of a single annotation spec.
	QuadraticKappa float32 `protobuf:"fixed32,7,opt,name=quadratic_kappa,json=quadraticKappa,proto3" json:"quadratic_kappa,omitempty"`
	// Output only. Confusion matrix of the evaluation.
	// Only set for the overall model evaluation, not for evaluation of a single
	// annotation spec.
	ConfusionMatrix *ClassificationEvaluationMetrics_ConfusionMatrix `protobuf:"bytes,8,opt,name=confusion_matrix,json=confusionMatrix,proto3" json:"confusion_matrix,omitempty"`
	// Output only. The annotation spec ids used for this evaluation.
	// Deprecated .
	//
	// Deprecated: Marked as deprecated in google/cloud/automl/v1beta1/text_sentiment.proto.
	AnnotationSpecId []string `protobuf:"bytes,9,rep,name=annotation_spec_id,json=annotationSpecId,proto3" json:"annotation_spec_id,omitempty"`
}

func (x *TextSentimentEvaluationMetrics) Reset() {
	*x = TextSentimentEvaluationMetrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1beta1_text_sentiment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextSentimentEvaluationMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextSentimentEvaluationMetrics) ProtoMessage() {}

func (x *TextSentimentEvaluationMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1beta1_text_sentiment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextSentimentEvaluationMetrics.ProtoReflect.Descriptor instead.
func (*TextSentimentEvaluationMetrics) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1beta1_text_sentiment_proto_rawDescGZIP(), []int{1}
}

func (x *TextSentimentEvaluationMetrics) GetPrecision() float32 {
	if x != nil {
		return x.Precision
	}
	return 0
}

func (x *TextSentimentEvaluationMetrics) GetRecall() float32 {
	if x != nil {
		return x.Recall
	}
	return 0
}

func (x *TextSentimentEvaluationMetrics) GetF1Score() float32 {
	if x != nil {
		return x.F1Score
	}
	return 0
}

func (x *TextSentimentEvaluationMetrics) GetMeanAbsoluteError() float32 {
	if x != nil {
		return x.MeanAbsoluteError
	}
	return 0
}

func (x *TextSentimentEvaluationMetrics) GetMeanSquaredError() float32 {
	if x != nil {
		return x.MeanSquaredError
	}
	return 0
}

func (x *TextSentimentEvaluationMetrics) GetLinearKappa() float32 {
	if x != nil {
		return x.LinearKappa
	}
	return 0
}

func (x *TextSentimentEvaluationMetrics) GetQuadraticKappa() float32 {
	if x != nil {
		return x.QuadraticKappa
	}
	return 0
}

func (x *TextSentimentEvaluationMetrics) GetConfusionMatrix() *ClassificationEvaluationMetrics_ConfusionMatrix {
	if x != nil {
		return x.ConfusionMatrix
	}
	return nil
}

// Deprecated: Marked as deprecated in google/cloud/automl/v1beta1/text_sentiment.proto.
func (x *TextSentimentEvaluationMetrics) GetAnnotationSpecId() []string {
	if x != nil {
		return x.AnnotationSpecId
	}
	return nil
}

var File_google_cloud_automl_v1beta1_text_sentiment_proto protoreflect.FileDescriptor

var file_google_cloud_automl_v1beta1_text_sentiment_proto_rawDesc = []byte{
	0x0a, 0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x74, 0x65,
	0x78, 0x74, 0x5f, 0x73, 0x65, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a,
	0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x75,
	0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x37, 0x0a, 0x17, 0x54, 0x65, 0x78, 0x74, 0x53, 0x65, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x6e, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x65, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x73, 0x65, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xc6, 0x03, 0x0a, 0x1e, 0x54,
	0x65, 0x78, 0x74, 0x53, 0x65, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x61, 0x6c,
	0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x09, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x63, 0x61, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x72, 0x65, 0x63,
	0x61, 0x6c, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x31, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x66, 0x31, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x2e,
	0x0a, 0x13, 0x6d, 0x65, 0x61, 0x6e, 0x5f, 0x61, 0x62, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x65, 0x5f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x11, 0x6d, 0x65, 0x61,
	0x6e, 0x41, 0x62, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2c,
	0x0a, 0x12, 0x6d, 0x65, 0x61, 0x6e, 0x5f, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x10, 0x6d, 0x65, 0x61, 0x6e,
	0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x21, 0x0a, 0x0c,
	0x6c, 0x69, 0x6e, 0x65, 0x61, 0x72, 0x5f, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0b, 0x6c, 0x69, 0x6e, 0x65, 0x61, 0x72, 0x4b, 0x61, 0x70, 0x70, 0x61, 0x12,
	0x27, 0x0a, 0x0f, 0x71, 0x75, 0x61, 0x64, 0x72, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x6b, 0x61, 0x70,
	0x70, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x71, 0x75, 0x61, 0x64, 0x72, 0x61,
	0x74, 0x69, 0x63, 0x4b, 0x61, 0x70, 0x70, 0x61, 0x12, 0x77, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x66,
	0x75, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x4c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45,
	0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78,
	0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x74, 0x72, 0x69,
	0x78, 0x12, 0x30, 0x0a, 0x12, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x70, 0x65, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x42, 0x02, 0x18,
	0x01, 0x52, 0x10, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65,
	0x63, 0x49, 0x64, 0x42, 0xad, 0x01, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x12, 0x54, 0x65, 0x78, 0x74, 0x53, 0x65, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x37, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x70, 0x62, 0x3b, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x6c, 0x70, 0x62, 0xca, 0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x5c, 0x41, 0x75, 0x74, 0x6f, 0x4d, 0x6c, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0xea, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x3a, 0x3a, 0x41, 0x75, 0x74, 0x6f, 0x4d, 0x4c, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_automl_v1beta1_text_sentiment_proto_rawDescOnce sync.Once
	file_google_cloud_automl_v1beta1_text_sentiment_proto_rawDescData = file_google_cloud_automl_v1beta1_text_sentiment_proto_rawDesc
)

func file_google_cloud_automl_v1beta1_text_sentiment_proto_rawDescGZIP() []byte {
	file_google_cloud_automl_v1beta1_text_sentiment_proto_rawDescOnce.Do(func() {
		file_google_cloud_automl_v1beta1_text_sentiment_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_automl_v1beta1_text_sentiment_proto_rawDescData)
	})
	return file_google_cloud_automl_v1beta1_text_sentiment_proto_rawDescData
}

var file_google_cloud_automl_v1beta1_text_sentiment_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_automl_v1beta1_text_sentiment_proto_goTypes = []interface{}{
	(*TextSentimentAnnotation)(nil),                         // 0: google.cloud.automl.v1beta1.TextSentimentAnnotation
	(*TextSentimentEvaluationMetrics)(nil),                  // 1: google.cloud.automl.v1beta1.TextSentimentEvaluationMetrics
	(*ClassificationEvaluationMetrics_ConfusionMatrix)(nil), // 2: google.cloud.automl.v1beta1.ClassificationEvaluationMetrics.ConfusionMatrix
}
var file_google_cloud_automl_v1beta1_text_sentiment_proto_depIdxs = []int32{
	2, // 0: google.cloud.automl.v1beta1.TextSentimentEvaluationMetrics.confusion_matrix:type_name -> google.cloud.automl.v1beta1.ClassificationEvaluationMetrics.ConfusionMatrix
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_automl_v1beta1_text_sentiment_proto_init() }
func file_google_cloud_automl_v1beta1_text_sentiment_proto_init() {
	if File_google_cloud_automl_v1beta1_text_sentiment_proto != nil {
		return
	}
	file_google_cloud_automl_v1beta1_classification_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_automl_v1beta1_text_sentiment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextSentimentAnnotation); i {
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
		file_google_cloud_automl_v1beta1_text_sentiment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextSentimentEvaluationMetrics); i {
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
			RawDescriptor: file_google_cloud_automl_v1beta1_text_sentiment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_automl_v1beta1_text_sentiment_proto_goTypes,
		DependencyIndexes: file_google_cloud_automl_v1beta1_text_sentiment_proto_depIdxs,
		MessageInfos:      file_google_cloud_automl_v1beta1_text_sentiment_proto_msgTypes,
	}.Build()
	File_google_cloud_automl_v1beta1_text_sentiment_proto = out.File
	file_google_cloud_automl_v1beta1_text_sentiment_proto_rawDesc = nil
	file_google_cloud_automl_v1beta1_text_sentiment_proto_goTypes = nil
	file_google_cloud_automl_v1beta1_text_sentiment_proto_depIdxs = nil
}
