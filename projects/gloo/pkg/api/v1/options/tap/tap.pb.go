// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/options/tap/tap.proto

package tap

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Tap filter: a filter that copies the contents of HTTP requests and responses
// to an external tap server. The full HTTP headers and bodies are reported in
// full to the configured address, and data can be reported using either over
// HTTP or GRPC.
type Tap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Sinks to which tap data should be output. Currently, only a single sink
	// is supported.
	Sinks []*Sink `protobuf:"bytes,1,rep,name=sinks,proto3" json:"sinks,omitempty"`
	// For buffered tapping, the maximum amount of received body that will be buffered
	// prior to truncation. If truncation occurs, the truncated field will be set.
	// If not specified, the default is 1KiB.
	MaxBufferedRxBytes *wrapperspb.UInt32Value `protobuf:"bytes,2,opt,name=max_buffered_rx_bytes,json=maxBufferedRxBytes,proto3" json:"max_buffered_rx_bytes,omitempty"`
	// For buffered tapping, the maximum amount of transmitted body that will be buffered
	// prior to truncation. If truncation occurs, the truncated field will be set.
	// If not specified, the default is 1KiB.
	MaxBufferedTxBytes *wrapperspb.UInt32Value `protobuf:"bytes,3,opt,name=max_buffered_tx_bytes,json=maxBufferedTxBytes,proto3" json:"max_buffered_tx_bytes,omitempty"`
	// Indicates whether tap filter records the time stamp for request/response headers.
	// Request headers time stamp is stored after receiving request headers.
	// Response headers time stamp is stored after receiving response headers.
	RecordHeadersReceivedTime *wrapperspb.BoolValue `protobuf:"bytes,4,opt,name=record_headers_received_time,json=recordHeadersReceivedTime,proto3" json:"record_headers_received_time,omitempty"`
	// Indicates whether report downstream connection info
	RecordDownstreamConnection *wrapperspb.BoolValue `protobuf:"bytes,5,opt,name=record_downstream_connection,json=recordDownstreamConnection,proto3" json:"record_downstream_connection,omitempty"`
}

func (x *Tap) Reset() {
	*x = Tap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_tap_tap_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tap) ProtoMessage() {}

func (x *Tap) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_tap_tap_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tap.ProtoReflect.Descriptor instead.
func (*Tap) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_tap_tap_proto_rawDescGZIP(), []int{0}
}

func (x *Tap) GetSinks() []*Sink {
	if x != nil {
		return x.Sinks
	}
	return nil
}

func (x *Tap) GetMaxBufferedRxBytes() *wrapperspb.UInt32Value {
	if x != nil {
		return x.MaxBufferedRxBytes
	}
	return nil
}

func (x *Tap) GetMaxBufferedTxBytes() *wrapperspb.UInt32Value {
	if x != nil {
		return x.MaxBufferedTxBytes
	}
	return nil
}

func (x *Tap) GetRecordHeadersReceivedTime() *wrapperspb.BoolValue {
	if x != nil {
		return x.RecordHeadersReceivedTime
	}
	return nil
}

func (x *Tap) GetRecordDownstreamConnection() *wrapperspb.BoolValue {
	if x != nil {
		return x.RecordDownstreamConnection
	}
	return nil
}

type Sink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of the output sink to which tap data should be written
	//
	// Types that are assignable to SinkType:
	//
	//	*Sink_GrpcService
	//	*Sink_HttpService
	SinkType isSink_SinkType `protobuf_oneof:"SinkType"`
}

func (x *Sink) Reset() {
	*x = Sink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_tap_tap_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sink) ProtoMessage() {}

func (x *Sink) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_tap_tap_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sink.ProtoReflect.Descriptor instead.
func (*Sink) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_tap_tap_proto_rawDescGZIP(), []int{1}
}

func (m *Sink) GetSinkType() isSink_SinkType {
	if m != nil {
		return m.SinkType
	}
	return nil
}

func (x *Sink) GetGrpcService() *GrpcService {
	if x, ok := x.GetSinkType().(*Sink_GrpcService); ok {
		return x.GrpcService
	}
	return nil
}

func (x *Sink) GetHttpService() *HttpService {
	if x, ok := x.GetSinkType().(*Sink_HttpService); ok {
		return x.HttpService
	}
	return nil
}

type isSink_SinkType interface {
	isSink_SinkType()
}

type Sink_GrpcService struct {
	// Write tap data out to a GRPC service
	GrpcService *GrpcService `protobuf:"bytes,1,opt,name=grpc_service,json=grpcService,proto3,oneof"`
}

type Sink_HttpService struct {
	// Write tap data out to a HTTP service
	HttpService *HttpService `protobuf:"bytes,2,opt,name=http_service,json=httpService,proto3,oneof"`
}

func (*Sink_GrpcService) isSink_SinkType() {}

func (*Sink_HttpService) isSink_SinkType() {}

// A tap sink over a GRPC service
type GrpcService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Upstream reference to the tap server
	TapServer *core.ResourceRef `protobuf:"bytes,1,opt,name=tap_server,json=tapServer,proto3" json:"tap_server,omitempty"`
}

func (x *GrpcService) Reset() {
	*x = GrpcService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_tap_tap_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcService) ProtoMessage() {}

func (x *GrpcService) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_tap_tap_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcService.ProtoReflect.Descriptor instead.
func (*GrpcService) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_tap_tap_proto_rawDescGZIP(), []int{2}
}

func (x *GrpcService) GetTapServer() *core.ResourceRef {
	if x != nil {
		return x.TapServer
	}
	return nil
}

// A tap sink over a HTTP service
type HttpService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Upstream reference to the tap server
	TapServer *core.ResourceRef `protobuf:"bytes,1,opt,name=tap_server,json=tapServer,proto3" json:"tap_server,omitempty"`
	// Connection timeout
	Timeout *durationpb.Duration `protobuf:"bytes,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *HttpService) Reset() {
	*x = HttpService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_tap_tap_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpService) ProtoMessage() {}

func (x *HttpService) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_tap_tap_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpService.ProtoReflect.Descriptor instead.
func (*HttpService) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_tap_tap_proto_rawDescGZIP(), []int{3}
}

func (x *HttpService) GetTapServer() *core.ResourceRef {
	if x != nil {
		return x.TapServer
	}
	return nil
}

func (x *HttpService) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_tap_tap_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_tap_tap_proto_rawDesc = []byte{
	0x0a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x74, 0x61, 0x70, 0x2f, 0x74, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x18, 0x74, 0x61, 0x70, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa4, 0x03, 0x0a, 0x03, 0x54, 0x61, 0x70, 0x12, 0x40, 0x0a, 0x05, 0x73, 0x69, 0x6e, 0x6b,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x53, 0x69, 0x6e, 0x6b, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x92, 0x01, 0x04, 0x08,
	0x01, 0x10, 0x01, 0x52, 0x05, 0x73, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x4f, 0x0a, 0x15, 0x6d, 0x61,
	0x78, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x65, 0x64, 0x5f, 0x72, 0x78, 0x5f, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74,
	0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x12, 0x6d, 0x61, 0x78, 0x42, 0x75, 0x66, 0x66,
	0x65, 0x72, 0x65, 0x64, 0x52, 0x78, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x4f, 0x0a, 0x15, 0x6d,
	0x61, 0x78, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x78, 0x5f, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e,
	0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x12, 0x6d, 0x61, 0x78, 0x42, 0x75, 0x66,
	0x66, 0x65, 0x72, 0x65, 0x64, 0x54, 0x78, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x5b, 0x0a, 0x1c,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x19,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x5c, 0x0a, 0x1c, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x1a, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x44, 0x6f, 0x77, 0x6e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xaf, 0x01, 0x0a, 0x04, 0x53, 0x69, 0x6e, 0x6b,
	0x12, 0x4a, 0x0a, 0x0c, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48, 0x00, 0x52,
	0x0b, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x0c,
	0x68, 0x74, 0x74, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x48, 0x74,
	0x74, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x68, 0x74, 0x74,
	0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x0f, 0x0a, 0x08, 0x53, 0x69, 0x6e, 0x6b,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x22, 0x51, 0x0a, 0x0b, 0x47, 0x72, 0x70,
	0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x0a, 0x74, 0x61, 0x70, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10,
	0x01, 0x52, 0x09, 0x74, 0x61, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x90, 0x01, 0x0a,
	0x0b, 0x48, 0x74, 0x74, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x0a,
	0x74, 0x61, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x09, 0x74, 0x61, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x12, 0x3d, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x42,
	0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x61, 0x70, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_tap_tap_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_tap_tap_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_tap_tap_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_tap_tap_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_tap_tap_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_tap_tap_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_tap_tap_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_tap_tap_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_tap_tap_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_tap_tap_proto_goTypes = []interface{}{
	(*Tap)(nil),                    // 0: tap.options.gloo.solo.io.Tap
	(*Sink)(nil),                   // 1: tap.options.gloo.solo.io.Sink
	(*GrpcService)(nil),            // 2: tap.options.gloo.solo.io.GrpcService
	(*HttpService)(nil),            // 3: tap.options.gloo.solo.io.HttpService
	(*wrapperspb.UInt32Value)(nil), // 4: google.protobuf.UInt32Value
	(*wrapperspb.BoolValue)(nil),   // 5: google.protobuf.BoolValue
	(*core.ResourceRef)(nil),       // 6: core.solo.io.ResourceRef
	(*durationpb.Duration)(nil),    // 7: google.protobuf.Duration
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_tap_tap_proto_depIdxs = []int32{
	1,  // 0: tap.options.gloo.solo.io.Tap.sinks:type_name -> tap.options.gloo.solo.io.Sink
	4,  // 1: tap.options.gloo.solo.io.Tap.max_buffered_rx_bytes:type_name -> google.protobuf.UInt32Value
	4,  // 2: tap.options.gloo.solo.io.Tap.max_buffered_tx_bytes:type_name -> google.protobuf.UInt32Value
	5,  // 3: tap.options.gloo.solo.io.Tap.record_headers_received_time:type_name -> google.protobuf.BoolValue
	5,  // 4: tap.options.gloo.solo.io.Tap.record_downstream_connection:type_name -> google.protobuf.BoolValue
	2,  // 5: tap.options.gloo.solo.io.Sink.grpc_service:type_name -> tap.options.gloo.solo.io.GrpcService
	3,  // 6: tap.options.gloo.solo.io.Sink.http_service:type_name -> tap.options.gloo.solo.io.HttpService
	6,  // 7: tap.options.gloo.solo.io.GrpcService.tap_server:type_name -> core.solo.io.ResourceRef
	6,  // 8: tap.options.gloo.solo.io.HttpService.tap_server:type_name -> core.solo.io.ResourceRef
	7,  // 9: tap.options.gloo.solo.io.HttpService.timeout:type_name -> google.protobuf.Duration
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_tap_tap_proto_init()
}
func file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_tap_tap_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_tap_tap_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_tap_tap_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tap); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_tap_tap_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sink); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_tap_tap_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcService); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_tap_tap_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpService); i {
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
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_tap_tap_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Sink_GrpcService)(nil),
		(*Sink_HttpService)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_tap_tap_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_tap_tap_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_tap_tap_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_tap_tap_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_tap_tap_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_tap_tap_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_tap_tap_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_tap_tap_proto_depIdxs = nil
}
