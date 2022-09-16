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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.5
// source: google/cloud/gkeconnect/gateway/v1beta1/gateway.proto

package gatewaypb

import (
	context "context"
	reflect "reflect"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_google_cloud_gkeconnect_gateway_v1beta1_gateway_proto protoreflect.FileDescriptor

var file_google_cloud_gkeconnect_gateway_v1beta1_gateway_proto_rawDesc = []byte{
	0x0a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67,
	0x6b, 0x65, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6b, 0x65, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x62, 0x6f, 0x64, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0xf9, 0x03, 0x0a, 0x0e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x6f, 0x64, 0x79,
	0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x2a, 0x2a, 0x12, 0x4f, 0x0a, 0x0c, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x1a, 0x14, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x6f, 0x64,
	0x79, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x22, 0x0b, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2f, 0x2a, 0x2a, 0x12, 0x51, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x1a, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70,
	0x42, 0x6f, 0x64, 0x79, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x2a, 0x0b, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x2a, 0x2a, 0x12, 0x4e, 0x0a, 0x0b, 0x50, 0x75, 0x74,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x1a, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70,
	0x42, 0x6f, 0x64, 0x79, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x1a, 0x0b, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x2a, 0x2a, 0x12, 0x50, 0x0a, 0x0d, 0x50, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x6f, 0x64, 0x79,
	0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x74,
	0x74, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x32, 0x0b,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x2a, 0x2a, 0x1a, 0x51, 0xca, 0x41, 0x1d,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x2e,
	0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0x8f,
	0x02, 0x0a, 0x2b, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6b, 0x65, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0c,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6b, 0x65, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0xaa, 0x02,
	0x27, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x47, 0x6b,
	0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x27, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x47, 0x6b, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x5c, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0xea, 0x02, 0x2b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x3a, 0x3a, 0x47, 0x6b, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x3a, 0x3a,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_google_cloud_gkeconnect_gateway_v1beta1_gateway_proto_goTypes = []interface{}{
	(*httpbody.HttpBody)(nil), // 0: google.api.HttpBody
}
var file_google_cloud_gkeconnect_gateway_v1beta1_gateway_proto_depIdxs = []int32{
	0, // 0: google.cloud.gkeconnect.gateway.v1beta1.GatewayService.GetResource:input_type -> google.api.HttpBody
	0, // 1: google.cloud.gkeconnect.gateway.v1beta1.GatewayService.PostResource:input_type -> google.api.HttpBody
	0, // 2: google.cloud.gkeconnect.gateway.v1beta1.GatewayService.DeleteResource:input_type -> google.api.HttpBody
	0, // 3: google.cloud.gkeconnect.gateway.v1beta1.GatewayService.PutResource:input_type -> google.api.HttpBody
	0, // 4: google.cloud.gkeconnect.gateway.v1beta1.GatewayService.PatchResource:input_type -> google.api.HttpBody
	0, // 5: google.cloud.gkeconnect.gateway.v1beta1.GatewayService.GetResource:output_type -> google.api.HttpBody
	0, // 6: google.cloud.gkeconnect.gateway.v1beta1.GatewayService.PostResource:output_type -> google.api.HttpBody
	0, // 7: google.cloud.gkeconnect.gateway.v1beta1.GatewayService.DeleteResource:output_type -> google.api.HttpBody
	0, // 8: google.cloud.gkeconnect.gateway.v1beta1.GatewayService.PutResource:output_type -> google.api.HttpBody
	0, // 9: google.cloud.gkeconnect.gateway.v1beta1.GatewayService.PatchResource:output_type -> google.api.HttpBody
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_gkeconnect_gateway_v1beta1_gateway_proto_init() }
func file_google_cloud_gkeconnect_gateway_v1beta1_gateway_proto_init() {
	if File_google_cloud_gkeconnect_gateway_v1beta1_gateway_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_gkeconnect_gateway_v1beta1_gateway_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_gkeconnect_gateway_v1beta1_gateway_proto_goTypes,
		DependencyIndexes: file_google_cloud_gkeconnect_gateway_v1beta1_gateway_proto_depIdxs,
	}.Build()
	File_google_cloud_gkeconnect_gateway_v1beta1_gateway_proto = out.File
	file_google_cloud_gkeconnect_gateway_v1beta1_gateway_proto_rawDesc = nil
	file_google_cloud_gkeconnect_gateway_v1beta1_gateway_proto_goTypes = nil
	file_google_cloud_gkeconnect_gateway_v1beta1_gateway_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GatewayServiceClient is the client API for GatewayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GatewayServiceClient interface {
	// GetResource performs an HTTP GET request on the Kubernetes API Server.
	GetResource(ctx context.Context, in *httpbody.HttpBody, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
	// PostResource performs an HTTP POST on the Kubernetes API Server.
	PostResource(ctx context.Context, in *httpbody.HttpBody, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
	// DeleteResource performs an HTTP DELETE on the Kubernetes API Server.
	DeleteResource(ctx context.Context, in *httpbody.HttpBody, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
	// PutResource performs an HTTP PUT on the Kubernetes API Server.
	PutResource(ctx context.Context, in *httpbody.HttpBody, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
	// PatchResource performs an HTTP PATCH on the Kubernetes API Server.
	PatchResource(ctx context.Context, in *httpbody.HttpBody, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
}

type gatewayServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayServiceClient(cc grpc.ClientConnInterface) GatewayServiceClient {
	return &gatewayServiceClient{cc}
}

func (c *gatewayServiceClient) GetResource(ctx context.Context, in *httpbody.HttpBody, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, "/google.cloud.gkeconnect.gateway.v1beta1.GatewayService/GetResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) PostResource(ctx context.Context, in *httpbody.HttpBody, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, "/google.cloud.gkeconnect.gateway.v1beta1.GatewayService/PostResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) DeleteResource(ctx context.Context, in *httpbody.HttpBody, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, "/google.cloud.gkeconnect.gateway.v1beta1.GatewayService/DeleteResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) PutResource(ctx context.Context, in *httpbody.HttpBody, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, "/google.cloud.gkeconnect.gateway.v1beta1.GatewayService/PutResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) PatchResource(ctx context.Context, in *httpbody.HttpBody, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, "/google.cloud.gkeconnect.gateway.v1beta1.GatewayService/PatchResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServiceServer is the server API for GatewayService service.
type GatewayServiceServer interface {
	// GetResource performs an HTTP GET request on the Kubernetes API Server.
	GetResource(context.Context, *httpbody.HttpBody) (*httpbody.HttpBody, error)
	// PostResource performs an HTTP POST on the Kubernetes API Server.
	PostResource(context.Context, *httpbody.HttpBody) (*httpbody.HttpBody, error)
	// DeleteResource performs an HTTP DELETE on the Kubernetes API Server.
	DeleteResource(context.Context, *httpbody.HttpBody) (*httpbody.HttpBody, error)
	// PutResource performs an HTTP PUT on the Kubernetes API Server.
	PutResource(context.Context, *httpbody.HttpBody) (*httpbody.HttpBody, error)
	// PatchResource performs an HTTP PATCH on the Kubernetes API Server.
	PatchResource(context.Context, *httpbody.HttpBody) (*httpbody.HttpBody, error)
}

// UnimplementedGatewayServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGatewayServiceServer struct {
}

func (*UnimplementedGatewayServiceServer) GetResource(context.Context, *httpbody.HttpBody) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResource not implemented")
}
func (*UnimplementedGatewayServiceServer) PostResource(context.Context, *httpbody.HttpBody) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostResource not implemented")
}
func (*UnimplementedGatewayServiceServer) DeleteResource(context.Context, *httpbody.HttpBody) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteResource not implemented")
}
func (*UnimplementedGatewayServiceServer) PutResource(context.Context, *httpbody.HttpBody) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutResource not implemented")
}
func (*UnimplementedGatewayServiceServer) PatchResource(context.Context, *httpbody.HttpBody) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchResource not implemented")
}

func RegisterGatewayServiceServer(s *grpc.Server, srv GatewayServiceServer) {
	s.RegisterService(&_GatewayService_serviceDesc, srv)
}

func _GatewayService_GetResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(httpbody.HttpBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).GetResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.gkeconnect.gateway.v1beta1.GatewayService/GetResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).GetResource(ctx, req.(*httpbody.HttpBody))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_PostResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(httpbody.HttpBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).PostResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.gkeconnect.gateway.v1beta1.GatewayService/PostResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).PostResource(ctx, req.(*httpbody.HttpBody))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_DeleteResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(httpbody.HttpBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).DeleteResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.gkeconnect.gateway.v1beta1.GatewayService/DeleteResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).DeleteResource(ctx, req.(*httpbody.HttpBody))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_PutResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(httpbody.HttpBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).PutResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.gkeconnect.gateway.v1beta1.GatewayService/PutResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).PutResource(ctx, req.(*httpbody.HttpBody))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_PatchResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(httpbody.HttpBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).PatchResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.gkeconnect.gateway.v1beta1.GatewayService/PatchResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).PatchResource(ctx, req.(*httpbody.HttpBody))
	}
	return interceptor(ctx, in, info, handler)
}

var _GatewayService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.gkeconnect.gateway.v1beta1.GatewayService",
	HandlerType: (*GatewayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetResource",
			Handler:    _GatewayService_GetResource_Handler,
		},
		{
			MethodName: "PostResource",
			Handler:    _GatewayService_PostResource_Handler,
		},
		{
			MethodName: "DeleteResource",
			Handler:    _GatewayService_DeleteResource_Handler,
		},
		{
			MethodName: "PutResource",
			Handler:    _GatewayService_PutResource_Handler,
		},
		{
			MethodName: "PatchResource",
			Handler:    _GatewayService_PatchResource_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/gkeconnect/gateway/v1beta1/gateway.proto",
}
