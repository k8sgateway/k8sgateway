//
//Copyright The Kubernetes Authors.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// This is a copy of https://github.com/kubernetes/apimachinery/blob/v0.28.3/pkg/runtime/schema/generated.proto
// with the go_package changed to a gloo path.
// Ideally we should update this proto every time we upgrade our k8s.io/apimachinery dependency.

// This file was autogenerated by go-to-protobuf. Do not edit it manually!

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: github.com/solo-io/gloo/projects/gateway2/api/external/kubernetes/apimachinery/pkg/runtime/schema/generated.proto

package schema

import (
	reflect "reflect"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_github_com_solo_io_gloo_projects_gateway2_api_external_kubernetes_apimachinery_pkg_runtime_schema_generated_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gateway2_api_external_kubernetes_apimachinery_pkg_runtime_schema_generated_proto_rawDesc = []byte{
	0x0a, 0x71, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x26, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x6d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x75, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x42, 0x67, 0x5a, 0x65, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69,
	0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x32, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72,
	0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61,
}

var file_github_com_solo_io_gloo_projects_gateway2_api_external_kubernetes_apimachinery_pkg_runtime_schema_generated_proto_goTypes = []interface{}{}
var file_github_com_solo_io_gloo_projects_gateway2_api_external_kubernetes_apimachinery_pkg_runtime_schema_generated_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_projects_gateway2_api_external_kubernetes_apimachinery_pkg_runtime_schema_generated_proto_init()
}
func file_github_com_solo_io_gloo_projects_gateway2_api_external_kubernetes_apimachinery_pkg_runtime_schema_generated_proto_init() {
	if File_github_com_solo_io_gloo_projects_gateway2_api_external_kubernetes_apimachinery_pkg_runtime_schema_generated_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gateway2_api_external_kubernetes_apimachinery_pkg_runtime_schema_generated_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gateway2_api_external_kubernetes_apimachinery_pkg_runtime_schema_generated_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gateway2_api_external_kubernetes_apimachinery_pkg_runtime_schema_generated_proto_depIdxs,
	}.Build()
	File_github_com_solo_io_gloo_projects_gateway2_api_external_kubernetes_apimachinery_pkg_runtime_schema_generated_proto = out.File
	file_github_com_solo_io_gloo_projects_gateway2_api_external_kubernetes_apimachinery_pkg_runtime_schema_generated_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gateway2_api_external_kubernetes_apimachinery_pkg_runtime_schema_generated_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gateway2_api_external_kubernetes_apimachinery_pkg_runtime_schema_generated_proto_depIdxs = nil
}
