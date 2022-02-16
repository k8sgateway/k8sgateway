// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: github.com/solo-io/solo-projects/projects/gloo-fed/api/fed/v1/failover.proto

package types

import (
	reflect "reflect"
	sync "sync"

	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The State of a reconciled object
type FailoverSchemeStatus_State int32

const (
	// Waiting to be processed.
	FailoverSchemeStatus_PENDING FailoverSchemeStatus_State = 0
	// Currently processing.
	FailoverSchemeStatus_PROCESSING FailoverSchemeStatus_State = 1
	// Invalid parameters supplied, will not continue.
	FailoverSchemeStatus_INVALID FailoverSchemeStatus_State = 2
	// Failed during processing.
	FailoverSchemeStatus_FAILED FailoverSchemeStatus_State = 3
	// Finished processing successfully.
	FailoverSchemeStatus_ACCEPTED FailoverSchemeStatus_State = 4
)

// Enum value maps for FailoverSchemeStatus_State.
var (
	FailoverSchemeStatus_State_name = map[int32]string{
		0: "PENDING",
		1: "PROCESSING",
		2: "INVALID",
		3: "FAILED",
		4: "ACCEPTED",
	}
	FailoverSchemeStatus_State_value = map[string]int32{
		"PENDING":    0,
		"PROCESSING": 1,
		"INVALID":    2,
		"FAILED":     3,
		"ACCEPTED":   4,
	}
)

func (x FailoverSchemeStatus_State) Enum() *FailoverSchemeStatus_State {
	p := new(FailoverSchemeStatus_State)
	*p = x
	return p
}

func (x FailoverSchemeStatus_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FailoverSchemeStatus_State) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto_enumTypes[0].Descriptor()
}

func (FailoverSchemeStatus_State) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto_enumTypes[0]
}

func (x FailoverSchemeStatus_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FailoverSchemeStatus_State.Descriptor instead.
func (FailoverSchemeStatus_State) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto_rawDescGZIP(), []int{1, 0}
}

//
//FailoverSpec is the core portion of the API for enabling failover between Gloo Upstreams in gloo-fed.
//This API is heavily inspired by the Failover API present in the Gloo Upstream which can be found in
//`api/gloo/v1/upstream`.
//
//The source Upstream below is the initial primary target of traffic. The type of endpoints vary by the type
//of Upstream specified. Each target specified is then configured as a failover endpoint in the case that
//the prmiary Upstream becomes unhealthy. The priority of the failover endpoints is inferred from the
//order in which the Upstreams are specified. source = [0], targets = [1-n].
//
//Example:
//
//primary:
//cluster: primary
//name: primary
//namespace: primary
//failover_groups:
//- priority_group:
//- cluster: A
//upstreams:
//- name: one
//namespace: one
//- cluster: B
//upstreams:
//- name: two
//namespace: two
//- priority_group:
//- cluster: C
//upstreams:
//- name: one
//namespace: one
//- cluster: D
//upstreams:
//- name: two
//namespace: two
type FailoverSchemeSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The upstream which will be configured for failover.
	Primary        *v1.ClusterObjectRef                    `protobuf:"bytes,1,opt,name=primary,proto3" json:"primary,omitempty"`
	FailoverGroups []*FailoverSchemeSpec_FailoverEndpoints `protobuf:"bytes,2,rep,name=failover_groups,json=failoverGroups,proto3" json:"failover_groups,omitempty"`
}

func (x *FailoverSchemeSpec) Reset() {
	*x = FailoverSchemeSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FailoverSchemeSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FailoverSchemeSpec) ProtoMessage() {}

func (x *FailoverSchemeSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FailoverSchemeSpec.ProtoReflect.Descriptor instead.
func (*FailoverSchemeSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto_rawDescGZIP(), []int{0}
}

func (x *FailoverSchemeSpec) GetPrimary() *v1.ClusterObjectRef {
	if x != nil {
		return x.Primary
	}
	return nil
}

func (x *FailoverSchemeSpec) GetFailoverGroups() []*FailoverSchemeSpec_FailoverEndpoints {
	if x != nil {
		return x.FailoverGroups
	}
	return nil
}

type FailoverSchemeStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The current state of the resource
	State FailoverSchemeStatus_State `protobuf:"varint,1,opt,name=state,proto3,enum=fed.solo.io.FailoverSchemeStatus_State" json:"state,omitempty"`
	// A human readable message about the current state of the object
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// The most recently observed generation of the resource. This value corresponds to the `metadata.generation` of
	// a kubernetes resource
	ObservedGeneration int64 `protobuf:"varint,3,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	// The time at which this status was recorded
	ProcessingTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=processing_time,json=processingTime,proto3" json:"processing_time,omitempty"`
}

func (x *FailoverSchemeStatus) Reset() {
	*x = FailoverSchemeStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FailoverSchemeStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FailoverSchemeStatus) ProtoMessage() {}

func (x *FailoverSchemeStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FailoverSchemeStatus.ProtoReflect.Descriptor instead.
func (*FailoverSchemeStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto_rawDescGZIP(), []int{1}
}

func (x *FailoverSchemeStatus) GetState() FailoverSchemeStatus_State {
	if x != nil {
		return x.State
	}
	return FailoverSchemeStatus_PENDING
}

func (x *FailoverSchemeStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *FailoverSchemeStatus) GetObservedGeneration() int64 {
	if x != nil {
		return x.ObservedGeneration
	}
	return 0
}

func (x *FailoverSchemeStatus) GetProcessingTime() *timestamp.Timestamp {
	if x != nil {
		return x.ProcessingTime
	}
	return nil
}

type FailoverSchemeSpec_FailoverEndpoints struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PriorityGroup []*FailoverSchemeSpec_FailoverEndpoints_LocalityLbTargets `protobuf:"bytes,2,rep,name=priority_group,json=priorityGroup,proto3" json:"priority_group,omitempty"`
}

func (x *FailoverSchemeSpec_FailoverEndpoints) Reset() {
	*x = FailoverSchemeSpec_FailoverEndpoints{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FailoverSchemeSpec_FailoverEndpoints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FailoverSchemeSpec_FailoverEndpoints) ProtoMessage() {}

func (x *FailoverSchemeSpec_FailoverEndpoints) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FailoverSchemeSpec_FailoverEndpoints.ProtoReflect.Descriptor instead.
func (*FailoverSchemeSpec_FailoverEndpoints) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto_rawDescGZIP(), []int{0, 0}
}

func (x *FailoverSchemeSpec_FailoverEndpoints) GetPriorityGroup() []*FailoverSchemeSpec_FailoverEndpoints_LocalityLbTargets {
	if x != nil {
		return x.PriorityGroup
	}
	return nil
}

type FailoverSchemeSpec_FailoverEndpoints_LocalityLbTargets struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (REQUIRED) Cluster on which the endpoints for this Group can be found
	Cluster string `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// A list of Upstream targets, each of these targets must exist on the cluster specified in this message
	Upstreams []*v1.ObjectRef `protobuf:"bytes,2,rep,name=upstreams,proto3" json:"upstreams,omitempty"`
	// (optional) locality load balancing weight assigned to the specified upstreams.
	// Locality load balancing will add a special load balancing weight among all
	// targets within a given priority, who are located in the zame zone.
	// See envoy Locality Weighted Load Balancing for more information:
	// https://www.envoyproxy.io/docs/envoy/v1.14.1/intro/arch_overview/upstream/load_balancing/locality_weight#arch-overview-load-balancing-locality-weighted-lb
	LocalityWeight *wrappers.UInt32Value `protobuf:"bytes,3,opt,name=locality_weight,json=localityWeight,proto3" json:"locality_weight,omitempty"`
}

func (x *FailoverSchemeSpec_FailoverEndpoints_LocalityLbTargets) Reset() {
	*x = FailoverSchemeSpec_FailoverEndpoints_LocalityLbTargets{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FailoverSchemeSpec_FailoverEndpoints_LocalityLbTargets) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FailoverSchemeSpec_FailoverEndpoints_LocalityLbTargets) ProtoMessage() {}

func (x *FailoverSchemeSpec_FailoverEndpoints_LocalityLbTargets) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FailoverSchemeSpec_FailoverEndpoints_LocalityLbTargets.ProtoReflect.Descriptor instead.
func (*FailoverSchemeSpec_FailoverEndpoints_LocalityLbTargets) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *FailoverSchemeSpec_FailoverEndpoints_LocalityLbTargets) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *FailoverSchemeSpec_FailoverEndpoints_LocalityLbTargets) GetUpstreams() []*v1.ObjectRef {
	if x != nil {
		return x.Upstreams
	}
	return nil
}

func (x *FailoverSchemeSpec_FailoverEndpoints_LocalityLbTargets) GetLocalityWeight() *wrappers.UInt32Value {
	if x != nil {
		return x.LocalityWeight
	}
	return nil
}

var File_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto_rawDesc = []byte{
	0x0a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f,
	0x2d, 0x66, 0x65, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x2f,
	0x66, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x66, 0x65, 0x64, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6b, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xe4, 0x03, 0x0a, 0x12, 0x46, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x3d, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x73, 0x6b, 0x76, 0x32, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x52, 0x07, 0x70,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x5a, 0x0a, 0x0f, 0x66, 0x61, 0x69, 0x6c, 0x6f, 0x76,
	0x65, 0x72, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x31, 0x2e, 0x66, 0x65, 0x64, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x46, 0x61,
	0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x53, 0x70, 0x65, 0x63,
	0x2e, 0x46, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x52, 0x0e, 0x66, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x1a, 0xb2, 0x02, 0x0a, 0x11, 0x46, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x6a, 0x0a, 0x0e, 0x70, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x43, 0x2e, 0x66, 0x65, 0x64, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x46,
	0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x53, 0x70, 0x65,
	0x63, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x4c, 0x62, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x73, 0x52, 0x0d, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x1a, 0xb0, 0x01, 0x0a, 0x11, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74,
	0x79, 0x4c, 0x62, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x09, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73,
	0x6b, 0x76, 0x32, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x66, 0x52, 0x09, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73,
	0x12, 0x45, 0x0a, 0x0f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74,
	0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74,
	0x79, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0xb2, 0x02, 0x0a, 0x14, 0x46, 0x61, 0x69, 0x6c,
	0x6f, 0x76, 0x65, 0x72, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x3d, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x27, 0x2e, 0x66, 0x65, 0x64, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x46, 0x61,
	0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x6f, 0x62, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x43, 0x0a, 0x0f, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0x4b, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44,
	0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53,
	0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0c,
	0x0a, 0x08, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44, 0x10, 0x04, 0x42, 0x59, 0x5a, 0x4f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d,
	0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x66,
	0x65, 0x64, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x65, 0x64, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0xc0,
	0xf5, 0x04, 0x01, 0xb8, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto_rawDescData = file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto_rawDesc
)

func file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto_rawDescData
}

var file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto_goTypes = []interface{}{
	(FailoverSchemeStatus_State)(0),                                // 0: fed.solo.io.FailoverSchemeStatus.State
	(*FailoverSchemeSpec)(nil),                                     // 1: fed.solo.io.FailoverSchemeSpec
	(*FailoverSchemeStatus)(nil),                                   // 2: fed.solo.io.FailoverSchemeStatus
	(*FailoverSchemeSpec_FailoverEndpoints)(nil),                   // 3: fed.solo.io.FailoverSchemeSpec.FailoverEndpoints
	(*FailoverSchemeSpec_FailoverEndpoints_LocalityLbTargets)(nil), // 4: fed.solo.io.FailoverSchemeSpec.FailoverEndpoints.LocalityLbTargets
	(*v1.ClusterObjectRef)(nil),                                    // 5: core.skv2.solo.io.ClusterObjectRef
	(*timestamp.Timestamp)(nil),                                    // 6: google.protobuf.Timestamp
	(*v1.ObjectRef)(nil),                                           // 7: core.skv2.solo.io.ObjectRef
	(*wrappers.UInt32Value)(nil),                                   // 8: google.protobuf.UInt32Value
}
var file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto_depIdxs = []int32{
	5, // 0: fed.solo.io.FailoverSchemeSpec.primary:type_name -> core.skv2.solo.io.ClusterObjectRef
	3, // 1: fed.solo.io.FailoverSchemeSpec.failover_groups:type_name -> fed.solo.io.FailoverSchemeSpec.FailoverEndpoints
	0, // 2: fed.solo.io.FailoverSchemeStatus.state:type_name -> fed.solo.io.FailoverSchemeStatus.State
	6, // 3: fed.solo.io.FailoverSchemeStatus.processing_time:type_name -> google.protobuf.Timestamp
	4, // 4: fed.solo.io.FailoverSchemeSpec.FailoverEndpoints.priority_group:type_name -> fed.solo.io.FailoverSchemeSpec.FailoverEndpoints.LocalityLbTargets
	7, // 5: fed.solo.io.FailoverSchemeSpec.FailoverEndpoints.LocalityLbTargets.upstreams:type_name -> core.skv2.solo.io.ObjectRef
	8, // 6: fed.solo.io.FailoverSchemeSpec.FailoverEndpoints.LocalityLbTargets.locality_weight:type_name -> google.protobuf.UInt32Value
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto_init() }
func file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto_init() {
	if File_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FailoverSchemeSpec); i {
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
		file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FailoverSchemeStatus); i {
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
		file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FailoverSchemeSpec_FailoverEndpoints); i {
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
		file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FailoverSchemeSpec_FailoverEndpoints_LocalityLbTargets); i {
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
			RawDescriptor: file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto = out.File
	file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto_rawDesc = nil
	file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto_goTypes = nil
	file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_v1_failover_proto_depIdxs = nil
}
