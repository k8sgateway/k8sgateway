// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gateway2/api/v1alpha1/kube/autoscaling.proto

package kube

import (
	reflect "reflect"
	sync "sync"

	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Kubernetes autoscaling configuration.
type Autoscaling struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HorizontalPodAutoscaler *HorizontalPodAutoscaler `protobuf:"bytes,1,opt,name=horizontal_pod_autoscaler,json=horizontalPodAutoscaler,proto3" json:"horizontal_pod_autoscaler,omitempty"`
}

func (x *Autoscaling) Reset() {
	*x = Autoscaling{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_autoscaling_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Autoscaling) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Autoscaling) ProtoMessage() {}

func (x *Autoscaling) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_autoscaling_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Autoscaling.ProtoReflect.Descriptor instead.
func (*Autoscaling) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_autoscaling_proto_rawDescGZIP(), []int{0}
}

func (x *Autoscaling) GetHorizontalPodAutoscaler() *HorizontalPodAutoscaler {
	if x != nil {
		return x.HorizontalPodAutoscaler
	}
	return nil
}

// Horizontal pod autoscaling configuration. See
// https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/
// for details.
type HorizontalPodAutoscaler struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The lower limit for the number of replicas to which the autoscaler can
	// scale down. Defaults to 1.
	MinReplicas *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=min_replicas,json=minReplicas,proto3" json:"min_replicas,omitempty"`
	// The upper limit for the number of replicas to which the autoscaler can
	// scale up. Cannot be less than `minReplicas`. Defaults to 100.
	MaxReplicas *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=max_replicas,json=maxReplicas,proto3" json:"max_replicas,omitempty"`
	// The target value of the average CPU utilization across all relevant pods,
	// represented as a percentage of the requested value of the resource for the
	// pods. Defaults to 80.
	TargetCpuUtilizationPercentage *wrappers.UInt32Value `protobuf:"bytes,3,opt,name=target_cpu_utilization_percentage,json=targetCpuUtilizationPercentage,proto3" json:"target_cpu_utilization_percentage,omitempty"`
	// The target value of the average memory utilization across all relevant
	// pods, represented as a percentage of the requested value of the resource
	// for the pods. Defaults to 80.
	TargetMemoryUtilizationPercentage *wrappers.UInt32Value `protobuf:"bytes,4,opt,name=target_memory_utilization_percentage,json=targetMemoryUtilizationPercentage,proto3" json:"target_memory_utilization_percentage,omitempty"`
}

func (x *HorizontalPodAutoscaler) Reset() {
	*x = HorizontalPodAutoscaler{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_autoscaling_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HorizontalPodAutoscaler) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HorizontalPodAutoscaler) ProtoMessage() {}

func (x *HorizontalPodAutoscaler) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_autoscaling_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HorizontalPodAutoscaler.ProtoReflect.Descriptor instead.
func (*HorizontalPodAutoscaler) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_autoscaling_proto_rawDescGZIP(), []int{1}
}

func (x *HorizontalPodAutoscaler) GetMinReplicas() *wrappers.UInt32Value {
	if x != nil {
		return x.MinReplicas
	}
	return nil
}

func (x *HorizontalPodAutoscaler) GetMaxReplicas() *wrappers.UInt32Value {
	if x != nil {
		return x.MaxReplicas
	}
	return nil
}

func (x *HorizontalPodAutoscaler) GetTargetCpuUtilizationPercentage() *wrappers.UInt32Value {
	if x != nil {
		return x.TargetCpuUtilizationPercentage
	}
	return nil
}

func (x *HorizontalPodAutoscaler) GetTargetMemoryUtilizationPercentage() *wrappers.UInt32Value {
	if x != nil {
		return x.TargetMemoryUtilizationPercentage
	}
	return nil
}

var File_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_autoscaling_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_autoscaling_proto_rawDesc = []byte{
	0x0a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x2f, 0x61, 0x75,
	0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x19, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7d,
	0x0a, 0x0b, 0x41, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x12, 0x6e, 0x0a,
	0x19, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x6f, 0x64, 0x5f,
	0x61, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x32, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x48, 0x6f, 0x72,
	0x69, 0x7a, 0x6f, 0x6e, 0x74, 0x61, 0x6c, 0x50, 0x6f, 0x64, 0x41, 0x75, 0x74, 0x6f, 0x73, 0x63,
	0x61, 0x6c, 0x65, 0x72, 0x52, 0x17, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x74, 0x61, 0x6c,
	0x50, 0x6f, 0x64, 0x41, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x22, 0xf3, 0x02,
	0x0a, 0x17, 0x48, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x74, 0x61, 0x6c, 0x50, 0x6f, 0x64, 0x41,
	0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x0c, 0x6d, 0x69, 0x6e,
	0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x6d,
	0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x12, 0x3f, 0x0a, 0x0c, 0x6d, 0x61,
	0x78, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b,
	0x6d, 0x61, 0x78, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x12, 0x67, 0x0a, 0x21, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x70, 0x75, 0x5f, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x1e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x70, 0x75, 0x55,
	0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x61, 0x67, 0x65, 0x12, 0x6d, 0x0a, 0x24, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x6d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x21, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x55, 0x74,
	0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x61, 0x67, 0x65, 0x42, 0x5e, 0xb8, 0xf5, 0x04, 0x01, 0xc0, 0xf5, 0x04, 0x01, 0x5a, 0x54, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69,
	0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x32, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6b,
	0x75, 0x62, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_autoscaling_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_autoscaling_proto_rawDescData = file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_autoscaling_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_autoscaling_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_autoscaling_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_autoscaling_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_autoscaling_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_autoscaling_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_autoscaling_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_autoscaling_proto_goTypes = []interface{}{
	(*Autoscaling)(nil),             // 0: kube.gateway.gloo.solo.io.Autoscaling
	(*HorizontalPodAutoscaler)(nil), // 1: kube.gateway.gloo.solo.io.HorizontalPodAutoscaler
	(*wrappers.UInt32Value)(nil),    // 2: google.protobuf.UInt32Value
}
var file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_autoscaling_proto_depIdxs = []int32{
	1, // 0: kube.gateway.gloo.solo.io.Autoscaling.horizontal_pod_autoscaler:type_name -> kube.gateway.gloo.solo.io.HorizontalPodAutoscaler
	2, // 1: kube.gateway.gloo.solo.io.HorizontalPodAutoscaler.min_replicas:type_name -> google.protobuf.UInt32Value
	2, // 2: kube.gateway.gloo.solo.io.HorizontalPodAutoscaler.max_replicas:type_name -> google.protobuf.UInt32Value
	2, // 3: kube.gateway.gloo.solo.io.HorizontalPodAutoscaler.target_cpu_utilization_percentage:type_name -> google.protobuf.UInt32Value
	2, // 4: kube.gateway.gloo.solo.io.HorizontalPodAutoscaler.target_memory_utilization_percentage:type_name -> google.protobuf.UInt32Value
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_autoscaling_proto_init()
}
func file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_autoscaling_proto_init() {
	if File_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_autoscaling_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_autoscaling_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Autoscaling); i {
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
		file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_autoscaling_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HorizontalPodAutoscaler); i {
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
			RawDescriptor: file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_autoscaling_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_autoscaling_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_autoscaling_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_autoscaling_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_autoscaling_proto = out.File
	file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_autoscaling_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_autoscaling_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_autoscaling_proto_depIdxs = nil
}
