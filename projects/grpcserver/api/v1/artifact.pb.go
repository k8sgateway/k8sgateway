// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-projects/projects/grpcserver/api/v1/artifact.proto

package v1

import (
	bytes "bytes"
	context "context"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type GetArtifactRequest struct {
	Ref                  *core.ResourceRef `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetArtifactRequest) Reset()         { *m = GetArtifactRequest{} }
func (m *GetArtifactRequest) String() string { return proto.CompactTextString(m) }
func (*GetArtifactRequest) ProtoMessage()    {}
func (*GetArtifactRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_31231f9377f1df05, []int{0}
}
func (m *GetArtifactRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetArtifactRequest.Unmarshal(m, b)
}
func (m *GetArtifactRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetArtifactRequest.Marshal(b, m, deterministic)
}
func (m *GetArtifactRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetArtifactRequest.Merge(m, src)
}
func (m *GetArtifactRequest) XXX_Size() int {
	return xxx_messageInfo_GetArtifactRequest.Size(m)
}
func (m *GetArtifactRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetArtifactRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetArtifactRequest proto.InternalMessageInfo

func (m *GetArtifactRequest) GetRef() *core.ResourceRef {
	if m != nil {
		return m.Ref
	}
	return nil
}

type GetArtifactResponse struct {
	Artifact             *v1.Artifact `protobuf:"bytes,1,opt,name=artifact,proto3" json:"artifact,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetArtifactResponse) Reset()         { *m = GetArtifactResponse{} }
func (m *GetArtifactResponse) String() string { return proto.CompactTextString(m) }
func (*GetArtifactResponse) ProtoMessage()    {}
func (*GetArtifactResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_31231f9377f1df05, []int{1}
}
func (m *GetArtifactResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetArtifactResponse.Unmarshal(m, b)
}
func (m *GetArtifactResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetArtifactResponse.Marshal(b, m, deterministic)
}
func (m *GetArtifactResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetArtifactResponse.Merge(m, src)
}
func (m *GetArtifactResponse) XXX_Size() int {
	return xxx_messageInfo_GetArtifactResponse.Size(m)
}
func (m *GetArtifactResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetArtifactResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetArtifactResponse proto.InternalMessageInfo

func (m *GetArtifactResponse) GetArtifact() *v1.Artifact {
	if m != nil {
		return m.Artifact
	}
	return nil
}

type ListArtifactsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListArtifactsRequest) Reset()         { *m = ListArtifactsRequest{} }
func (m *ListArtifactsRequest) String() string { return proto.CompactTextString(m) }
func (*ListArtifactsRequest) ProtoMessage()    {}
func (*ListArtifactsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_31231f9377f1df05, []int{2}
}
func (m *ListArtifactsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListArtifactsRequest.Unmarshal(m, b)
}
func (m *ListArtifactsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListArtifactsRequest.Marshal(b, m, deterministic)
}
func (m *ListArtifactsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListArtifactsRequest.Merge(m, src)
}
func (m *ListArtifactsRequest) XXX_Size() int {
	return xxx_messageInfo_ListArtifactsRequest.Size(m)
}
func (m *ListArtifactsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListArtifactsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListArtifactsRequest proto.InternalMessageInfo

type ListArtifactsResponse struct {
	Artifacts            []*v1.Artifact `protobuf:"bytes,1,rep,name=artifacts,proto3" json:"artifacts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListArtifactsResponse) Reset()         { *m = ListArtifactsResponse{} }
func (m *ListArtifactsResponse) String() string { return proto.CompactTextString(m) }
func (*ListArtifactsResponse) ProtoMessage()    {}
func (*ListArtifactsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_31231f9377f1df05, []int{3}
}
func (m *ListArtifactsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListArtifactsResponse.Unmarshal(m, b)
}
func (m *ListArtifactsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListArtifactsResponse.Marshal(b, m, deterministic)
}
func (m *ListArtifactsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListArtifactsResponse.Merge(m, src)
}
func (m *ListArtifactsResponse) XXX_Size() int {
	return xxx_messageInfo_ListArtifactsResponse.Size(m)
}
func (m *ListArtifactsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListArtifactsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListArtifactsResponse proto.InternalMessageInfo

func (m *ListArtifactsResponse) GetArtifacts() []*v1.Artifact {
	if m != nil {
		return m.Artifacts
	}
	return nil
}

type CreateArtifactRequest struct {
	Artifact             *v1.Artifact `protobuf:"bytes,3,opt,name=artifact,proto3" json:"artifact,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateArtifactRequest) Reset()         { *m = CreateArtifactRequest{} }
func (m *CreateArtifactRequest) String() string { return proto.CompactTextString(m) }
func (*CreateArtifactRequest) ProtoMessage()    {}
func (*CreateArtifactRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_31231f9377f1df05, []int{4}
}
func (m *CreateArtifactRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateArtifactRequest.Unmarshal(m, b)
}
func (m *CreateArtifactRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateArtifactRequest.Marshal(b, m, deterministic)
}
func (m *CreateArtifactRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateArtifactRequest.Merge(m, src)
}
func (m *CreateArtifactRequest) XXX_Size() int {
	return xxx_messageInfo_CreateArtifactRequest.Size(m)
}
func (m *CreateArtifactRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateArtifactRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateArtifactRequest proto.InternalMessageInfo

func (m *CreateArtifactRequest) GetArtifact() *v1.Artifact {
	if m != nil {
		return m.Artifact
	}
	return nil
}

type CreateArtifactResponse struct {
	Artifact             *v1.Artifact `protobuf:"bytes,1,opt,name=artifact,proto3" json:"artifact,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateArtifactResponse) Reset()         { *m = CreateArtifactResponse{} }
func (m *CreateArtifactResponse) String() string { return proto.CompactTextString(m) }
func (*CreateArtifactResponse) ProtoMessage()    {}
func (*CreateArtifactResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_31231f9377f1df05, []int{5}
}
func (m *CreateArtifactResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateArtifactResponse.Unmarshal(m, b)
}
func (m *CreateArtifactResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateArtifactResponse.Marshal(b, m, deterministic)
}
func (m *CreateArtifactResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateArtifactResponse.Merge(m, src)
}
func (m *CreateArtifactResponse) XXX_Size() int {
	return xxx_messageInfo_CreateArtifactResponse.Size(m)
}
func (m *CreateArtifactResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateArtifactResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateArtifactResponse proto.InternalMessageInfo

func (m *CreateArtifactResponse) GetArtifact() *v1.Artifact {
	if m != nil {
		return m.Artifact
	}
	return nil
}

type UpdateArtifactRequest struct {
	Artifact             *v1.Artifact `protobuf:"bytes,3,opt,name=artifact,proto3" json:"artifact,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpdateArtifactRequest) Reset()         { *m = UpdateArtifactRequest{} }
func (m *UpdateArtifactRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateArtifactRequest) ProtoMessage()    {}
func (*UpdateArtifactRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_31231f9377f1df05, []int{6}
}
func (m *UpdateArtifactRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateArtifactRequest.Unmarshal(m, b)
}
func (m *UpdateArtifactRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateArtifactRequest.Marshal(b, m, deterministic)
}
func (m *UpdateArtifactRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateArtifactRequest.Merge(m, src)
}
func (m *UpdateArtifactRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateArtifactRequest.Size(m)
}
func (m *UpdateArtifactRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateArtifactRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateArtifactRequest proto.InternalMessageInfo

func (m *UpdateArtifactRequest) GetArtifact() *v1.Artifact {
	if m != nil {
		return m.Artifact
	}
	return nil
}

type UpdateArtifactResponse struct {
	Artifact             *v1.Artifact `protobuf:"bytes,1,opt,name=artifact,proto3" json:"artifact,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpdateArtifactResponse) Reset()         { *m = UpdateArtifactResponse{} }
func (m *UpdateArtifactResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateArtifactResponse) ProtoMessage()    {}
func (*UpdateArtifactResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_31231f9377f1df05, []int{7}
}
func (m *UpdateArtifactResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateArtifactResponse.Unmarshal(m, b)
}
func (m *UpdateArtifactResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateArtifactResponse.Marshal(b, m, deterministic)
}
func (m *UpdateArtifactResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateArtifactResponse.Merge(m, src)
}
func (m *UpdateArtifactResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateArtifactResponse.Size(m)
}
func (m *UpdateArtifactResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateArtifactResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateArtifactResponse proto.InternalMessageInfo

func (m *UpdateArtifactResponse) GetArtifact() *v1.Artifact {
	if m != nil {
		return m.Artifact
	}
	return nil
}

type DeleteArtifactRequest struct {
	Ref                  *core.ResourceRef `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DeleteArtifactRequest) Reset()         { *m = DeleteArtifactRequest{} }
func (m *DeleteArtifactRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteArtifactRequest) ProtoMessage()    {}
func (*DeleteArtifactRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_31231f9377f1df05, []int{8}
}
func (m *DeleteArtifactRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteArtifactRequest.Unmarshal(m, b)
}
func (m *DeleteArtifactRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteArtifactRequest.Marshal(b, m, deterministic)
}
func (m *DeleteArtifactRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteArtifactRequest.Merge(m, src)
}
func (m *DeleteArtifactRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteArtifactRequest.Size(m)
}
func (m *DeleteArtifactRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteArtifactRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteArtifactRequest proto.InternalMessageInfo

func (m *DeleteArtifactRequest) GetRef() *core.ResourceRef {
	if m != nil {
		return m.Ref
	}
	return nil
}

type DeleteArtifactResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteArtifactResponse) Reset()         { *m = DeleteArtifactResponse{} }
func (m *DeleteArtifactResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteArtifactResponse) ProtoMessage()    {}
func (*DeleteArtifactResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_31231f9377f1df05, []int{9}
}
func (m *DeleteArtifactResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteArtifactResponse.Unmarshal(m, b)
}
func (m *DeleteArtifactResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteArtifactResponse.Marshal(b, m, deterministic)
}
func (m *DeleteArtifactResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteArtifactResponse.Merge(m, src)
}
func (m *DeleteArtifactResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteArtifactResponse.Size(m)
}
func (m *DeleteArtifactResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteArtifactResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteArtifactResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetArtifactRequest)(nil), "glooeeapi.solo.io.GetArtifactRequest")
	proto.RegisterType((*GetArtifactResponse)(nil), "glooeeapi.solo.io.GetArtifactResponse")
	proto.RegisterType((*ListArtifactsRequest)(nil), "glooeeapi.solo.io.ListArtifactsRequest")
	proto.RegisterType((*ListArtifactsResponse)(nil), "glooeeapi.solo.io.ListArtifactsResponse")
	proto.RegisterType((*CreateArtifactRequest)(nil), "glooeeapi.solo.io.CreateArtifactRequest")
	proto.RegisterType((*CreateArtifactResponse)(nil), "glooeeapi.solo.io.CreateArtifactResponse")
	proto.RegisterType((*UpdateArtifactRequest)(nil), "glooeeapi.solo.io.UpdateArtifactRequest")
	proto.RegisterType((*UpdateArtifactResponse)(nil), "glooeeapi.solo.io.UpdateArtifactResponse")
	proto.RegisterType((*DeleteArtifactRequest)(nil), "glooeeapi.solo.io.DeleteArtifactRequest")
	proto.RegisterType((*DeleteArtifactResponse)(nil), "glooeeapi.solo.io.DeleteArtifactResponse")
}

func init() {
	proto.RegisterFile("github.com/solo-io/solo-projects/projects/grpcserver/api/v1/artifact.proto", fileDescriptor_31231f9377f1df05)
}

var fileDescriptor_31231f9377f1df05 = []byte{
	// 428 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xdf, 0x8e, 0xd2, 0x40,
	0x18, 0xc5, 0x69, 0x30, 0x46, 0x3f, 0xa2, 0x89, 0x23, 0x6d, 0xb0, 0x17, 0x86, 0x34, 0x51, 0x31,
	0xea, 0x34, 0xa2, 0x77, 0x5c, 0x55, 0x49, 0x8c, 0x82, 0x37, 0x4d, 0xbc, 0xf1, 0xc2, 0xa4, 0x94,
	0xaf, 0x75, 0x94, 0xdd, 0x99, 0x9d, 0x19, 0x78, 0x8c, 0x7d, 0x8e, 0x7d, 0xae, 0x7d, 0x92, 0x4d,
	0xe9, 0x9f, 0x6d, 0xcb, 0xec, 0xc2, 0x86, 0xbd, 0x62, 0x18, 0xce, 0x77, 0x7e, 0x87, 0xce, 0xe9,
	0xc0, 0x8f, 0x94, 0xe9, 0xbf, 0xeb, 0x05, 0x8d, 0xf9, 0x89, 0xaf, 0xf8, 0x8a, 0x7f, 0x60, 0x3c,
	0xff, 0x14, 0x92, 0xff, 0xc3, 0x58, 0x2b, 0xbf, 0x5a, 0xa4, 0x52, 0xc4, 0x0a, 0xe5, 0x06, 0xa5,
	0x1f, 0x09, 0xe6, 0x6f, 0x3e, 0xfa, 0x91, 0xd4, 0x2c, 0x89, 0x62, 0x4d, 0x85, 0xe4, 0x9a, 0x93,
	0x67, 0xe9, 0x8a, 0x73, 0xc4, 0x48, 0x30, 0x9a, 0x59, 0x50, 0xc6, 0xdd, 0x7e, 0xca, 0x53, 0xbe,
	0xfd, 0xd5, 0xcf, 0x56, 0xb9, 0xd0, 0x9d, 0x18, 0xa0, 0xd9, 0x6c, 0x8d, 0x95, 0x7d, 0x33, 0x52,
	0xdc, 0xf7, 0x37, 0x25, 0xfe, 0xcf, 0x74, 0x39, 0x22, 0x31, 0xc9, 0xd5, 0x5e, 0x00, 0xe4, 0x1b,
	0xea, 0xa0, 0xb0, 0x08, 0xf1, 0x6c, 0x8d, 0x4a, 0x93, 0x77, 0xd0, 0x95, 0x98, 0x0c, 0xac, 0xa1,
	0x35, 0xea, 0x8d, 0x5f, 0xd0, 0x98, 0x4b, 0x2c, 0x23, 0xd3, 0x10, 0x15, 0x5f, 0xcb, 0x18, 0x43,
	0x4c, 0xc2, 0x4c, 0xe5, 0x7d, 0x87, 0xe7, 0x0d, 0x0b, 0x25, 0xf8, 0xa9, 0x42, 0x32, 0x86, 0x47,
	0x65, 0xb2, 0xc2, 0xc8, 0xa1, 0x59, 0xec, 0xca, 0xa8, 0x9a, 0xa8, 0x74, 0x9e, 0x03, 0xfd, 0x39,
	0x53, 0x95, 0x97, 0x2a, 0xf2, 0x78, 0x3f, 0xc1, 0x6e, 0xed, 0x17, 0x90, 0xcf, 0xf0, 0xb8, 0x1c,
	0x56, 0x03, 0x6b, 0xd8, 0xbd, 0x85, 0x72, 0x2d, 0xf4, 0x66, 0x60, 0x7f, 0x95, 0x18, 0x69, 0x6c,
	0xff, 0xef, 0x7a, 0xe6, 0xee, 0x81, 0x99, 0xe7, 0xe0, 0xb4, 0xcd, 0x8e, 0x78, 0x02, 0x33, 0xb0,
	0x7f, 0x89, 0xe5, 0xfd, 0x45, 0x6b, 0x9b, 0x1d, 0x11, 0x6d, 0x0a, 0xf6, 0x14, 0x57, 0xb8, 0x1b,
	0xed, 0x4e, 0x6d, 0x19, 0x80, 0xd3, 0x76, 0xc9, 0x33, 0x8d, 0xcf, 0x1f, 0x40, 0xaf, 0xdc, 0x0c,
	0x04, 0x23, 0x7f, 0xa0, 0x57, 0xeb, 0x15, 0x79, 0x45, 0x77, 0x5e, 0x1f, 0xba, 0x5b, 0x5d, 0xf7,
	0xf5, 0x3e, 0x59, 0x4e, 0xf3, 0x3a, 0x64, 0x09, 0x4f, 0x1a, 0xa5, 0x22, 0x6f, 0x0c, 0xa3, 0xa6,
	0x3a, 0xba, 0xa3, 0xfd, 0xc2, 0x8a, 0x92, 0xc2, 0xd3, 0x66, 0x3d, 0x88, 0x69, 0xda, 0x58, 0x47,
	0xf7, 0xed, 0x01, 0xca, 0x3a, 0xa8, 0x79, 0xd8, 0x46, 0x90, 0xb1, 0x5c, 0x46, 0x90, 0xb9, 0x39,
	0x39, 0xa8, 0x79, 0x82, 0x46, 0x90, 0xb1, 0x2a, 0x46, 0x90, 0xb9, 0x0e, 0x5e, 0xe7, 0x4b, 0x70,
	0x71, 0xf9, 0xd2, 0xfa, 0x3d, 0x39, 0xe2, 0x06, 0x5e, 0x3c, 0xdc, 0xde, 0x72, 0x9f, 0xae, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xde, 0xd2, 0x7f, 0x25, 0xc7, 0x05, 0x00, 0x00,
}

func (this *GetArtifactRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetArtifactRequest)
	if !ok {
		that2, ok := that.(GetArtifactRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Ref.Equal(that1.Ref) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *GetArtifactResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetArtifactResponse)
	if !ok {
		that2, ok := that.(GetArtifactResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Artifact.Equal(that1.Artifact) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ListArtifactsRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ListArtifactsRequest)
	if !ok {
		that2, ok := that.(ListArtifactsRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ListArtifactsResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ListArtifactsResponse)
	if !ok {
		that2, ok := that.(ListArtifactsResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Artifacts) != len(that1.Artifacts) {
		return false
	}
	for i := range this.Artifacts {
		if !this.Artifacts[i].Equal(that1.Artifacts[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *CreateArtifactRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CreateArtifactRequest)
	if !ok {
		that2, ok := that.(CreateArtifactRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Artifact.Equal(that1.Artifact) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *CreateArtifactResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CreateArtifactResponse)
	if !ok {
		that2, ok := that.(CreateArtifactResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Artifact.Equal(that1.Artifact) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *UpdateArtifactRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpdateArtifactRequest)
	if !ok {
		that2, ok := that.(UpdateArtifactRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Artifact.Equal(that1.Artifact) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *UpdateArtifactResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpdateArtifactResponse)
	if !ok {
		that2, ok := that.(UpdateArtifactResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Artifact.Equal(that1.Artifact) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *DeleteArtifactRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DeleteArtifactRequest)
	if !ok {
		that2, ok := that.(DeleteArtifactRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Ref.Equal(that1.Ref) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *DeleteArtifactResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DeleteArtifactResponse)
	if !ok {
		that2, ok := that.(DeleteArtifactResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ArtifactApiClient is the client API for ArtifactApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ArtifactApiClient interface {
	GetArtifact(ctx context.Context, in *GetArtifactRequest, opts ...grpc.CallOption) (*GetArtifactResponse, error)
	ListArtifacts(ctx context.Context, in *ListArtifactsRequest, opts ...grpc.CallOption) (*ListArtifactsResponse, error)
	CreateArtifact(ctx context.Context, in *CreateArtifactRequest, opts ...grpc.CallOption) (*CreateArtifactResponse, error)
	UpdateArtifact(ctx context.Context, in *UpdateArtifactRequest, opts ...grpc.CallOption) (*UpdateArtifactResponse, error)
	DeleteArtifact(ctx context.Context, in *DeleteArtifactRequest, opts ...grpc.CallOption) (*DeleteArtifactResponse, error)
}

type artifactApiClient struct {
	cc *grpc.ClientConn
}

func NewArtifactApiClient(cc *grpc.ClientConn) ArtifactApiClient {
	return &artifactApiClient{cc}
}

func (c *artifactApiClient) GetArtifact(ctx context.Context, in *GetArtifactRequest, opts ...grpc.CallOption) (*GetArtifactResponse, error) {
	out := new(GetArtifactResponse)
	err := c.cc.Invoke(ctx, "/glooeeapi.solo.io.ArtifactApi/GetArtifact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactApiClient) ListArtifacts(ctx context.Context, in *ListArtifactsRequest, opts ...grpc.CallOption) (*ListArtifactsResponse, error) {
	out := new(ListArtifactsResponse)
	err := c.cc.Invoke(ctx, "/glooeeapi.solo.io.ArtifactApi/ListArtifacts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactApiClient) CreateArtifact(ctx context.Context, in *CreateArtifactRequest, opts ...grpc.CallOption) (*CreateArtifactResponse, error) {
	out := new(CreateArtifactResponse)
	err := c.cc.Invoke(ctx, "/glooeeapi.solo.io.ArtifactApi/CreateArtifact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactApiClient) UpdateArtifact(ctx context.Context, in *UpdateArtifactRequest, opts ...grpc.CallOption) (*UpdateArtifactResponse, error) {
	out := new(UpdateArtifactResponse)
	err := c.cc.Invoke(ctx, "/glooeeapi.solo.io.ArtifactApi/UpdateArtifact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactApiClient) DeleteArtifact(ctx context.Context, in *DeleteArtifactRequest, opts ...grpc.CallOption) (*DeleteArtifactResponse, error) {
	out := new(DeleteArtifactResponse)
	err := c.cc.Invoke(ctx, "/glooeeapi.solo.io.ArtifactApi/DeleteArtifact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArtifactApiServer is the server API for ArtifactApi service.
type ArtifactApiServer interface {
	GetArtifact(context.Context, *GetArtifactRequest) (*GetArtifactResponse, error)
	ListArtifacts(context.Context, *ListArtifactsRequest) (*ListArtifactsResponse, error)
	CreateArtifact(context.Context, *CreateArtifactRequest) (*CreateArtifactResponse, error)
	UpdateArtifact(context.Context, *UpdateArtifactRequest) (*UpdateArtifactResponse, error)
	DeleteArtifact(context.Context, *DeleteArtifactRequest) (*DeleteArtifactResponse, error)
}

// UnimplementedArtifactApiServer can be embedded to have forward compatible implementations.
type UnimplementedArtifactApiServer struct {
}

func (*UnimplementedArtifactApiServer) GetArtifact(ctx context.Context, req *GetArtifactRequest) (*GetArtifactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArtifact not implemented")
}
func (*UnimplementedArtifactApiServer) ListArtifacts(ctx context.Context, req *ListArtifactsRequest) (*ListArtifactsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListArtifacts not implemented")
}
func (*UnimplementedArtifactApiServer) CreateArtifact(ctx context.Context, req *CreateArtifactRequest) (*CreateArtifactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArtifact not implemented")
}
func (*UnimplementedArtifactApiServer) UpdateArtifact(ctx context.Context, req *UpdateArtifactRequest) (*UpdateArtifactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateArtifact not implemented")
}
func (*UnimplementedArtifactApiServer) DeleteArtifact(ctx context.Context, req *DeleteArtifactRequest) (*DeleteArtifactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArtifact not implemented")
}

func RegisterArtifactApiServer(s *grpc.Server, srv ArtifactApiServer) {
	s.RegisterService(&_ArtifactApi_serviceDesc, srv)
}

func _ArtifactApi_GetArtifact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArtifactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactApiServer).GetArtifact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/glooeeapi.solo.io.ArtifactApi/GetArtifact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactApiServer).GetArtifact(ctx, req.(*GetArtifactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactApi_ListArtifacts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListArtifactsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactApiServer).ListArtifacts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/glooeeapi.solo.io.ArtifactApi/ListArtifacts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactApiServer).ListArtifacts(ctx, req.(*ListArtifactsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactApi_CreateArtifact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateArtifactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactApiServer).CreateArtifact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/glooeeapi.solo.io.ArtifactApi/CreateArtifact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactApiServer).CreateArtifact(ctx, req.(*CreateArtifactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactApi_UpdateArtifact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateArtifactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactApiServer).UpdateArtifact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/glooeeapi.solo.io.ArtifactApi/UpdateArtifact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactApiServer).UpdateArtifact(ctx, req.(*UpdateArtifactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactApi_DeleteArtifact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteArtifactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactApiServer).DeleteArtifact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/glooeeapi.solo.io.ArtifactApi/DeleteArtifact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactApiServer).DeleteArtifact(ctx, req.(*DeleteArtifactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ArtifactApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "glooeeapi.solo.io.ArtifactApi",
	HandlerType: (*ArtifactApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetArtifact",
			Handler:    _ArtifactApi_GetArtifact_Handler,
		},
		{
			MethodName: "ListArtifacts",
			Handler:    _ArtifactApi_ListArtifacts_Handler,
		},
		{
			MethodName: "CreateArtifact",
			Handler:    _ArtifactApi_CreateArtifact_Handler,
		},
		{
			MethodName: "UpdateArtifact",
			Handler:    _ArtifactApi_UpdateArtifact_Handler,
		},
		{
			MethodName: "DeleteArtifact",
			Handler:    _ArtifactApi_DeleteArtifact_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/solo-io/solo-projects/projects/grpcserver/api/v1/artifact.proto",
}
