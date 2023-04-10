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
// source: google/api/servicecontrol/v1/distribution.proto

package servicecontrolpb

import (
	reflect "reflect"
	sync "sync"

	distribution "google.golang.org/genproto/googleapis/api/distribution"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Distribution represents a frequency distribution of double-valued sample
// points. It contains the size of the population of sample points plus
// additional optional information:
//
// * the arithmetic mean of the samples
// * the minimum and maximum of the samples
// * the sum-squared-deviation of the samples, used to compute variance
// * a histogram of the values of the sample points
type Distribution struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The total number of samples in the distribution. Must be >= 0.
	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	// The arithmetic mean of the samples in the distribution. If `count` is
	// zero then this field must be zero.
	Mean float64 `protobuf:"fixed64,2,opt,name=mean,proto3" json:"mean,omitempty"`
	// The minimum of the population of values. Ignored if `count` is zero.
	Minimum float64 `protobuf:"fixed64,3,opt,name=minimum,proto3" json:"minimum,omitempty"`
	// The maximum of the population of values. Ignored if `count` is zero.
	Maximum float64 `protobuf:"fixed64,4,opt,name=maximum,proto3" json:"maximum,omitempty"`
	// The sum of squared deviations from the mean:
	//   Sum[i=1..count]((x_i - mean)^2)
	// where each x_i is a sample values. If `count` is zero then this field
	// must be zero, otherwise validation of the request fails.
	SumOfSquaredDeviation float64 `protobuf:"fixed64,5,opt,name=sum_of_squared_deviation,json=sumOfSquaredDeviation,proto3" json:"sum_of_squared_deviation,omitempty"`
	// The number of samples in each histogram bucket. `bucket_counts` are
	// optional. If present, they must sum to the `count` value.
	//
	// The buckets are defined below in `bucket_option`. There are N buckets.
	// `bucket_counts[0]` is the number of samples in the underflow bucket.
	// `bucket_counts[1]` to `bucket_counts[N-1]` are the numbers of samples
	// in each of the finite buckets. And `bucket_counts[N] is the number
	// of samples in the overflow bucket. See the comments of `bucket_option`
	// below for more details.
	//
	// Any suffix of trailing zeros may be omitted.
	BucketCounts []int64 `protobuf:"varint,6,rep,packed,name=bucket_counts,json=bucketCounts,proto3" json:"bucket_counts,omitempty"`
	// Defines the buckets in the histogram. `bucket_option` and `bucket_counts`
	// must be both set, or both unset.
	//
	// Buckets are numbered in the range of [0, N], with a total of N+1 buckets.
	// There must be at least two buckets (a single-bucket histogram gives
	// no information that isn't already provided by `count`).
	//
	// The first bucket is the underflow bucket which has a lower bound
	// of -inf. The last bucket is the overflow bucket which has an
	// upper bound of +inf. All other buckets (if any) are called "finite"
	// buckets because they have finite lower and upper bounds. As described
	// below, there are three ways to define the finite buckets.
	//
	//   (1) Buckets with constant width.
	//   (2) Buckets with exponentially growing widths.
	//   (3) Buckets with arbitrary user-provided widths.
	//
	// In all cases, the buckets cover the entire real number line (-inf,
	// +inf). Bucket upper bounds are exclusive and lower bounds are
	// inclusive. The upper bound of the underflow bucket is equal to the
	// lower bound of the smallest finite bucket; the lower bound of the
	// overflow bucket is equal to the upper bound of the largest finite
	// bucket.
	//
	// Types that are assignable to BucketOption:
	//	*Distribution_LinearBuckets_
	//	*Distribution_ExponentialBuckets_
	//	*Distribution_ExplicitBuckets_
	BucketOption isDistribution_BucketOption `protobuf_oneof:"bucket_option"`
	// Example points. Must be in increasing order of `value` field.
	Exemplars []*distribution.Distribution_Exemplar `protobuf:"bytes,10,rep,name=exemplars,proto3" json:"exemplars,omitempty"`
}

func (x *Distribution) Reset() {
	*x = Distribution{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_servicecontrol_v1_distribution_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Distribution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Distribution) ProtoMessage() {}

func (x *Distribution) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_servicecontrol_v1_distribution_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Distribution.ProtoReflect.Descriptor instead.
func (*Distribution) Descriptor() ([]byte, []int) {
	return file_google_api_servicecontrol_v1_distribution_proto_rawDescGZIP(), []int{0}
}

func (x *Distribution) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *Distribution) GetMean() float64 {
	if x != nil {
		return x.Mean
	}
	return 0
}

func (x *Distribution) GetMinimum() float64 {
	if x != nil {
		return x.Minimum
	}
	return 0
}

func (x *Distribution) GetMaximum() float64 {
	if x != nil {
		return x.Maximum
	}
	return 0
}

func (x *Distribution) GetSumOfSquaredDeviation() float64 {
	if x != nil {
		return x.SumOfSquaredDeviation
	}
	return 0
}

func (x *Distribution) GetBucketCounts() []int64 {
	if x != nil {
		return x.BucketCounts
	}
	return nil
}

func (m *Distribution) GetBucketOption() isDistribution_BucketOption {
	if m != nil {
		return m.BucketOption
	}
	return nil
}

func (x *Distribution) GetLinearBuckets() *Distribution_LinearBuckets {
	if x, ok := x.GetBucketOption().(*Distribution_LinearBuckets_); ok {
		return x.LinearBuckets
	}
	return nil
}

func (x *Distribution) GetExponentialBuckets() *Distribution_ExponentialBuckets {
	if x, ok := x.GetBucketOption().(*Distribution_ExponentialBuckets_); ok {
		return x.ExponentialBuckets
	}
	return nil
}

func (x *Distribution) GetExplicitBuckets() *Distribution_ExplicitBuckets {
	if x, ok := x.GetBucketOption().(*Distribution_ExplicitBuckets_); ok {
		return x.ExplicitBuckets
	}
	return nil
}

func (x *Distribution) GetExemplars() []*distribution.Distribution_Exemplar {
	if x != nil {
		return x.Exemplars
	}
	return nil
}

type isDistribution_BucketOption interface {
	isDistribution_BucketOption()
}

type Distribution_LinearBuckets_ struct {
	// Buckets with constant width.
	LinearBuckets *Distribution_LinearBuckets `protobuf:"bytes,7,opt,name=linear_buckets,json=linearBuckets,proto3,oneof"`
}

type Distribution_ExponentialBuckets_ struct {
	// Buckets with exponentially growing width.
	ExponentialBuckets *Distribution_ExponentialBuckets `protobuf:"bytes,8,opt,name=exponential_buckets,json=exponentialBuckets,proto3,oneof"`
}

type Distribution_ExplicitBuckets_ struct {
	// Buckets with arbitrary user-provided width.
	ExplicitBuckets *Distribution_ExplicitBuckets `protobuf:"bytes,9,opt,name=explicit_buckets,json=explicitBuckets,proto3,oneof"`
}

func (*Distribution_LinearBuckets_) isDistribution_BucketOption() {}

func (*Distribution_ExponentialBuckets_) isDistribution_BucketOption() {}

func (*Distribution_ExplicitBuckets_) isDistribution_BucketOption() {}

// Describing buckets with constant width.
type Distribution_LinearBuckets struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of finite buckets. With the underflow and overflow buckets,
	// the total number of buckets is `num_finite_buckets` + 2.
	// See comments on `bucket_options` for details.
	NumFiniteBuckets int32 `protobuf:"varint,1,opt,name=num_finite_buckets,json=numFiniteBuckets,proto3" json:"num_finite_buckets,omitempty"`
	// The i'th linear bucket covers the interval
	//   [offset + (i-1) * width, offset + i * width)
	// where i ranges from 1 to num_finite_buckets, inclusive.
	// Must be strictly positive.
	Width float64 `protobuf:"fixed64,2,opt,name=width,proto3" json:"width,omitempty"`
	// The i'th linear bucket covers the interval
	//   [offset + (i-1) * width, offset + i * width)
	// where i ranges from 1 to num_finite_buckets, inclusive.
	Offset float64 `protobuf:"fixed64,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *Distribution_LinearBuckets) Reset() {
	*x = Distribution_LinearBuckets{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_servicecontrol_v1_distribution_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Distribution_LinearBuckets) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Distribution_LinearBuckets) ProtoMessage() {}

func (x *Distribution_LinearBuckets) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_servicecontrol_v1_distribution_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Distribution_LinearBuckets.ProtoReflect.Descriptor instead.
func (*Distribution_LinearBuckets) Descriptor() ([]byte, []int) {
	return file_google_api_servicecontrol_v1_distribution_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Distribution_LinearBuckets) GetNumFiniteBuckets() int32 {
	if x != nil {
		return x.NumFiniteBuckets
	}
	return 0
}

func (x *Distribution_LinearBuckets) GetWidth() float64 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Distribution_LinearBuckets) GetOffset() float64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

// Describing buckets with exponentially growing width.
type Distribution_ExponentialBuckets struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of finite buckets. With the underflow and overflow buckets,
	// the total number of buckets is `num_finite_buckets` + 2.
	// See comments on `bucket_options` for details.
	NumFiniteBuckets int32 `protobuf:"varint,1,opt,name=num_finite_buckets,json=numFiniteBuckets,proto3" json:"num_finite_buckets,omitempty"`
	// The i'th exponential bucket covers the interval
	//   [scale * growth_factor^(i-1), scale * growth_factor^i)
	// where i ranges from 1 to num_finite_buckets inclusive.
	// Must be larger than 1.0.
	GrowthFactor float64 `protobuf:"fixed64,2,opt,name=growth_factor,json=growthFactor,proto3" json:"growth_factor,omitempty"`
	// The i'th exponential bucket covers the interval
	//   [scale * growth_factor^(i-1), scale * growth_factor^i)
	// where i ranges from 1 to num_finite_buckets inclusive.
	// Must be > 0.
	Scale float64 `protobuf:"fixed64,3,opt,name=scale,proto3" json:"scale,omitempty"`
}

func (x *Distribution_ExponentialBuckets) Reset() {
	*x = Distribution_ExponentialBuckets{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_servicecontrol_v1_distribution_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Distribution_ExponentialBuckets) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Distribution_ExponentialBuckets) ProtoMessage() {}

func (x *Distribution_ExponentialBuckets) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_servicecontrol_v1_distribution_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Distribution_ExponentialBuckets.ProtoReflect.Descriptor instead.
func (*Distribution_ExponentialBuckets) Descriptor() ([]byte, []int) {
	return file_google_api_servicecontrol_v1_distribution_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Distribution_ExponentialBuckets) GetNumFiniteBuckets() int32 {
	if x != nil {
		return x.NumFiniteBuckets
	}
	return 0
}

func (x *Distribution_ExponentialBuckets) GetGrowthFactor() float64 {
	if x != nil {
		return x.GrowthFactor
	}
	return 0
}

func (x *Distribution_ExponentialBuckets) GetScale() float64 {
	if x != nil {
		return x.Scale
	}
	return 0
}

// Describing buckets with arbitrary user-provided width.
type Distribution_ExplicitBuckets struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 'bound' is a list of strictly increasing boundaries between
	// buckets. Note that a list of length N-1 defines N buckets because
	// of fenceposting. See comments on `bucket_options` for details.
	//
	// The i'th finite bucket covers the interval
	//   [bound[i-1], bound[i])
	// where i ranges from 1 to bound_size() - 1. Note that there are no
	// finite buckets at all if 'bound' only contains a single element; in
	// that special case the single bound defines the boundary between the
	// underflow and overflow buckets.
	//
	// bucket number                   lower bound    upper bound
	//  i == 0 (underflow)              -inf           bound[i]
	//  0 < i < bound_size()            bound[i-1]     bound[i]
	//  i == bound_size() (overflow)    bound[i-1]     +inf
	Bounds []float64 `protobuf:"fixed64,1,rep,packed,name=bounds,proto3" json:"bounds,omitempty"`
}

func (x *Distribution_ExplicitBuckets) Reset() {
	*x = Distribution_ExplicitBuckets{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_servicecontrol_v1_distribution_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Distribution_ExplicitBuckets) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Distribution_ExplicitBuckets) ProtoMessage() {}

func (x *Distribution_ExplicitBuckets) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_servicecontrol_v1_distribution_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Distribution_ExplicitBuckets.ProtoReflect.Descriptor instead.
func (*Distribution_ExplicitBuckets) Descriptor() ([]byte, []int) {
	return file_google_api_servicecontrol_v1_distribution_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Distribution_ExplicitBuckets) GetBounds() []float64 {
	if x != nil {
		return x.Bounds
	}
	return nil
}

var File_google_api_servicecontrol_v1_distribution_proto protoreflect.FileDescriptor

var file_google_api_servicecontrol_v1_distribution_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x64,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x1a,
	0x1d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf1,
	0x06, 0x0a, 0x0c, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x61, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x04, 0x6d, 0x65, 0x61, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x69, 0x6e,
	0x69, 0x6d, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x6d, 0x69, 0x6e, 0x69,
	0x6d, 0x75, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x12, 0x37, 0x0a,
	0x18, 0x73, 0x75, 0x6d, 0x5f, 0x6f, 0x66, 0x5f, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x64, 0x5f,
	0x64, 0x65, 0x76, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x15, 0x73, 0x75, 0x6d, 0x4f, 0x66, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x64, 0x44, 0x65, 0x76,
	0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0c, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x61, 0x0a, 0x0e, 0x6c,
	0x69, 0x6e, 0x65, 0x61, 0x72, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x4c, 0x69, 0x6e, 0x65, 0x61, 0x72, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x48, 0x00, 0x52,
	0x0d, 0x6c, 0x69, 0x6e, 0x65, 0x61, 0x72, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x70,
	0x0a, 0x13, 0x65, 0x78, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x48, 0x00, 0x52, 0x12, 0x65, 0x78,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73,
	0x12, 0x67, 0x0a, 0x10, 0x65, 0x78, 0x70, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x5f, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x42,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x48, 0x00, 0x52, 0x0f, 0x65, 0x78, 0x70, 0x6c, 0x69, 0x63,
	0x69, 0x74, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x3f, 0x0a, 0x09, 0x65, 0x78, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x72, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x78, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x72, 0x52,
	0x09, 0x65, 0x78, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x72, 0x73, 0x1a, 0x6b, 0x0a, 0x0d, 0x4c, 0x69,
	0x6e, 0x65, 0x61, 0x72, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x6e,
	0x75, 0x6d, 0x5f, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x65, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x6e, 0x75, 0x6d, 0x46, 0x69, 0x6e, 0x69,
	0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64,
	0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x1a, 0x7d, 0x0a, 0x12, 0x45, 0x78, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x2c, 0x0a,
	0x12, 0x6e, 0x75, 0x6d, 0x5f, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x65, 0x5f, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x6e, 0x75, 0x6d, 0x46, 0x69,
	0x6e, 0x69, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x67,
	0x72, 0x6f, 0x77, 0x74, 0x68, 0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0c, 0x67, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x05, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x1a, 0x29, 0x0a, 0x0f, 0x45, 0x78, 0x70, 0x6c, 0x69, 0x63,
	0x69, 0x74, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x6f, 0x75,
	0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x01, 0x52, 0x06, 0x62, 0x6f, 0x75, 0x6e, 0x64,
	0x73, 0x42, 0x0f, 0x0a, 0x0d, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0xec, 0x01, 0x0a, 0x20, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x62, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x62, 0xf8, 0x01, 0x01, 0xaa, 0x02, 0x1e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1e, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x21,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_api_servicecontrol_v1_distribution_proto_rawDescOnce sync.Once
	file_google_api_servicecontrol_v1_distribution_proto_rawDescData = file_google_api_servicecontrol_v1_distribution_proto_rawDesc
)

func file_google_api_servicecontrol_v1_distribution_proto_rawDescGZIP() []byte {
	file_google_api_servicecontrol_v1_distribution_proto_rawDescOnce.Do(func() {
		file_google_api_servicecontrol_v1_distribution_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_api_servicecontrol_v1_distribution_proto_rawDescData)
	})
	return file_google_api_servicecontrol_v1_distribution_proto_rawDescData
}

var file_google_api_servicecontrol_v1_distribution_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_api_servicecontrol_v1_distribution_proto_goTypes = []interface{}{
	(*Distribution)(nil),                       // 0: google.api.servicecontrol.v1.Distribution
	(*Distribution_LinearBuckets)(nil),         // 1: google.api.servicecontrol.v1.Distribution.LinearBuckets
	(*Distribution_ExponentialBuckets)(nil),    // 2: google.api.servicecontrol.v1.Distribution.ExponentialBuckets
	(*Distribution_ExplicitBuckets)(nil),       // 3: google.api.servicecontrol.v1.Distribution.ExplicitBuckets
	(*distribution.Distribution_Exemplar)(nil), // 4: google.api.Distribution.Exemplar
}
var file_google_api_servicecontrol_v1_distribution_proto_depIdxs = []int32{
	1, // 0: google.api.servicecontrol.v1.Distribution.linear_buckets:type_name -> google.api.servicecontrol.v1.Distribution.LinearBuckets
	2, // 1: google.api.servicecontrol.v1.Distribution.exponential_buckets:type_name -> google.api.servicecontrol.v1.Distribution.ExponentialBuckets
	3, // 2: google.api.servicecontrol.v1.Distribution.explicit_buckets:type_name -> google.api.servicecontrol.v1.Distribution.ExplicitBuckets
	4, // 3: google.api.servicecontrol.v1.Distribution.exemplars:type_name -> google.api.Distribution.Exemplar
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_api_servicecontrol_v1_distribution_proto_init() }
func file_google_api_servicecontrol_v1_distribution_proto_init() {
	if File_google_api_servicecontrol_v1_distribution_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_api_servicecontrol_v1_distribution_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Distribution); i {
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
		file_google_api_servicecontrol_v1_distribution_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Distribution_LinearBuckets); i {
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
		file_google_api_servicecontrol_v1_distribution_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Distribution_ExponentialBuckets); i {
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
		file_google_api_servicecontrol_v1_distribution_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Distribution_ExplicitBuckets); i {
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
	file_google_api_servicecontrol_v1_distribution_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Distribution_LinearBuckets_)(nil),
		(*Distribution_ExponentialBuckets_)(nil),
		(*Distribution_ExplicitBuckets_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_api_servicecontrol_v1_distribution_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_api_servicecontrol_v1_distribution_proto_goTypes,
		DependencyIndexes: file_google_api_servicecontrol_v1_distribution_proto_depIdxs,
		MessageInfos:      file_google_api_servicecontrol_v1_distribution_proto_msgTypes,
	}.Build()
	File_google_api_servicecontrol_v1_distribution_proto = out.File
	file_google_api_servicecontrol_v1_distribution_proto_rawDesc = nil
	file_google_api_servicecontrol_v1_distribution_proto_goTypes = nil
	file_google_api_servicecontrol_v1_distribution_proto_depIdxs = nil
}
