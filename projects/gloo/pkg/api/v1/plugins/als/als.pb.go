// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/plugins/als/als.proto

package als

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

// Contains various settings for Envoy's access logging service.
// See here for more information: https://www.envoyproxy.io/docs/envoy/latest/api-v2/config/filter/accesslog/v2/accesslog.proto#envoy-api-msg-config-filter-accesslog-v2-accesslog
type AccessLoggingService struct {
	AccessLog            []*AccessLog `protobuf:"bytes,1,rep,name=access_log,json=accessLog,proto3" json:"access_log,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *AccessLoggingService) Reset()         { *m = AccessLoggingService{} }
func (m *AccessLoggingService) String() string { return proto.CompactTextString(m) }
func (*AccessLoggingService) ProtoMessage()    {}
func (*AccessLoggingService) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd8d2602efe636cc, []int{0}
}
func (m *AccessLoggingService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessLoggingService.Unmarshal(m, b)
}
func (m *AccessLoggingService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessLoggingService.Marshal(b, m, deterministic)
}
func (m *AccessLoggingService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessLoggingService.Merge(m, src)
}
func (m *AccessLoggingService) XXX_Size() int {
	return xxx_messageInfo_AccessLoggingService.Size(m)
}
func (m *AccessLoggingService) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessLoggingService.DiscardUnknown(m)
}

var xxx_messageInfo_AccessLoggingService proto.InternalMessageInfo

func (m *AccessLoggingService) GetAccessLog() []*AccessLog {
	if m != nil {
		return m.AccessLog
	}
	return nil
}

type AccessLog struct {
	// type of Access Logging service to implement
	//
	// Types that are valid to be assigned to OutputDestination:
	//	*AccessLog_FileSink
	OutputDestination    isAccessLog_OutputDestination `protobuf_oneof:"OutputDestination"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *AccessLog) Reset()         { *m = AccessLog{} }
func (m *AccessLog) String() string { return proto.CompactTextString(m) }
func (*AccessLog) ProtoMessage()    {}
func (*AccessLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd8d2602efe636cc, []int{1}
}
func (m *AccessLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessLog.Unmarshal(m, b)
}
func (m *AccessLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessLog.Marshal(b, m, deterministic)
}
func (m *AccessLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessLog.Merge(m, src)
}
func (m *AccessLog) XXX_Size() int {
	return xxx_messageInfo_AccessLog.Size(m)
}
func (m *AccessLog) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessLog.DiscardUnknown(m)
}

var xxx_messageInfo_AccessLog proto.InternalMessageInfo

type isAccessLog_OutputDestination interface {
	isAccessLog_OutputDestination()
	Equal(interface{}) bool
}

type AccessLog_FileSink struct {
	FileSink *FileSink `protobuf:"bytes,2,opt,name=file_sink,json=fileSink,proto3,oneof" json:"file_sink,omitempty"`
}

func (*AccessLog_FileSink) isAccessLog_OutputDestination() {}

func (m *AccessLog) GetOutputDestination() isAccessLog_OutputDestination {
	if m != nil {
		return m.OutputDestination
	}
	return nil
}

func (m *AccessLog) GetFileSink() *FileSink {
	if x, ok := m.GetOutputDestination().(*AccessLog_FileSink); ok {
		return x.FileSink
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AccessLog) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AccessLog_FileSink)(nil),
	}
}

type FileSink struct {
	// the file path to which the file access logging service will sink
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// the format which the logs should be outputted by
	//
	// Types that are valid to be assigned to OutputFormat:
	//	*FileSink_StringFormat
	//	*FileSink_JsonFormat
	OutputFormat         isFileSink_OutputFormat `protobuf_oneof:"output_format"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *FileSink) Reset()         { *m = FileSink{} }
func (m *FileSink) String() string { return proto.CompactTextString(m) }
func (*FileSink) ProtoMessage()    {}
func (*FileSink) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd8d2602efe636cc, []int{2}
}
func (m *FileSink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileSink.Unmarshal(m, b)
}
func (m *FileSink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileSink.Marshal(b, m, deterministic)
}
func (m *FileSink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileSink.Merge(m, src)
}
func (m *FileSink) XXX_Size() int {
	return xxx_messageInfo_FileSink.Size(m)
}
func (m *FileSink) XXX_DiscardUnknown() {
	xxx_messageInfo_FileSink.DiscardUnknown(m)
}

var xxx_messageInfo_FileSink proto.InternalMessageInfo

type isFileSink_OutputFormat interface {
	isFileSink_OutputFormat()
	Equal(interface{}) bool
}

type FileSink_StringFormat struct {
	StringFormat string `protobuf:"bytes,2,opt,name=string_format,json=stringFormat,proto3,oneof" json:"string_format,omitempty"`
}
type FileSink_JsonFormat struct {
	JsonFormat *types.Struct `protobuf:"bytes,3,opt,name=json_format,json=jsonFormat,proto3,oneof" json:"json_format,omitempty"`
}

func (*FileSink_StringFormat) isFileSink_OutputFormat() {}
func (*FileSink_JsonFormat) isFileSink_OutputFormat()   {}

func (m *FileSink) GetOutputFormat() isFileSink_OutputFormat {
	if m != nil {
		return m.OutputFormat
	}
	return nil
}

func (m *FileSink) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *FileSink) GetStringFormat() string {
	if x, ok := m.GetOutputFormat().(*FileSink_StringFormat); ok {
		return x.StringFormat
	}
	return ""
}

func (m *FileSink) GetJsonFormat() *types.Struct {
	if x, ok := m.GetOutputFormat().(*FileSink_JsonFormat); ok {
		return x.JsonFormat
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FileSink) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FileSink_StringFormat)(nil),
		(*FileSink_JsonFormat)(nil),
	}
}

func init() {
	proto.RegisterType((*AccessLoggingService)(nil), "als.plugins.gloo.solo.io.AccessLoggingService")
	proto.RegisterType((*AccessLog)(nil), "als.plugins.gloo.solo.io.AccessLog")
	proto.RegisterType((*FileSink)(nil), "als.plugins.gloo.solo.io.FileSink")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/plugins/als/als.proto", fileDescriptor_dd8d2602efe636cc)
}

var fileDescriptor_dd8d2602efe636cc = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x1b, 0x2b, 0xd2, 0x6c, 0x2d, 0x62, 0x2c, 0x18, 0x8a, 0x48, 0x89, 0x08, 0xbd, 0xb8,
	0x8b, 0xf5, 0x26, 0x5e, 0x1a, 0xa4, 0xf4, 0x20, 0x08, 0xe9, 0xad, 0x97, 0x92, 0x86, 0xcd, 0x76,
	0xda, 0xed, 0x4e, 0xc8, 0x6e, 0xfa, 0x20, 0x3e, 0x85, 0xcf, 0xe5, 0x93, 0x48, 0x76, 0x9b, 0x9e,
	0x2c, 0x78, 0x58, 0xf8, 0xf6, 0x9b, 0xdf, 0xfc, 0x61, 0x86, 0xc4, 0x02, 0xcc, 0xba, 0x5a, 0xd1,
	0x0c, 0x77, 0x4c, 0xa3, 0xc4, 0x27, 0x40, 0x26, 0x24, 0x22, 0x2b, 0x4a, 0xdc, 0xf0, 0xcc, 0x68,
	0xf7, 0x4b, 0x0b, 0x60, 0xfb, 0x67, 0x56, 0xc8, 0x4a, 0x80, 0xd2, 0x2c, 0x95, 0xf6, 0xd1, 0xa2,
	0x44, 0x83, 0x41, 0x68, 0xa5, 0x0b, 0xd1, 0x1a, 0xa7, 0x75, 0x25, 0x0a, 0x38, 0xe8, 0x0b, 0x14,
	0x68, 0x21, 0x56, 0x2b, 0xc7, 0x0f, 0xee, 0x04, 0xa2, 0x90, 0x9c, 0xd9, 0xdf, 0xaa, 0xca, 0x99,
	0x36, 0x65, 0x95, 0x19, 0x17, 0x8d, 0x16, 0xa4, 0x3f, 0xc9, 0x32, 0xae, 0xf5, 0x07, 0x0a, 0x01,
	0x4a, 0xcc, 0x79, 0xb9, 0x87, 0x8c, 0x07, 0x31, 0x21, 0xa9, 0xf5, 0x97, 0x12, 0x45, 0xe8, 0x0d,
	0xdb, 0xa3, 0xee, 0xf8, 0x81, 0x9e, 0x6a, 0x4d, 0x8f, 0x35, 0x12, 0x3f, 0x6d, 0x64, 0x94, 0x11,
	0xff, 0xe8, 0x07, 0x13, 0xe2, 0xe7, 0x20, 0xf9, 0x52, 0x83, 0xda, 0x86, 0x67, 0x43, 0x6f, 0xd4,
	0x1d, 0x47, 0xa7, 0xeb, 0x4d, 0x41, 0xf2, 0x39, 0xa8, 0xed, 0xac, 0x95, 0x74, 0xf2, 0x83, 0x8e,
	0x6f, 0xc8, 0xf5, 0x67, 0x65, 0x8a, 0xca, 0xbc, 0x73, 0x6d, 0x40, 0xa5, 0x06, 0x50, 0x45, 0x5f,
	0x1e, 0xe9, 0x34, 0x74, 0x10, 0x90, 0xf3, 0x22, 0x35, 0xeb, 0xd0, 0x1b, 0x7a, 0x23, 0x3f, 0xb1,
	0x3a, 0x78, 0x24, 0x3d, 0x6d, 0x4a, 0x50, 0x62, 0x99, 0x63, 0xb9, 0x4b, 0x8d, 0x6d, 0xee, 0xcf,
	0x5a, 0xc9, 0xa5, 0xb3, 0xa7, 0xd6, 0x0d, 0x5e, 0x49, 0x77, 0xa3, 0x51, 0x35, 0x50, 0xdb, 0x4e,
	0x78, 0x4b, 0xdd, 0xf2, 0x68, 0xb3, 0x3c, 0x3a, 0xb7, 0xcb, 0x9b, 0xb5, 0x12, 0x52, 0xd3, 0x2e,
	0x37, 0xbe, 0x22, 0x3d, 0xb4, 0x83, 0x1d, 0xb2, 0xe3, 0xf8, 0xfb, 0xe7, 0xde, 0x5b, 0xbc, 0xfd,
	0xef, 0xda, 0xc5, 0x56, 0xfc, 0x71, 0xf1, 0xd5, 0x85, 0xed, 0xf9, 0xf2, 0x1b, 0x00, 0x00, 0xff,
	0xff, 0x45, 0x47, 0xf1, 0xb0, 0x34, 0x02, 0x00, 0x00,
}

func (this *AccessLoggingService) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AccessLoggingService)
	if !ok {
		that2, ok := that.(AccessLoggingService)
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
	if len(this.AccessLog) != len(that1.AccessLog) {
		return false
	}
	for i := range this.AccessLog {
		if !this.AccessLog[i].Equal(that1.AccessLog[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *AccessLog) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AccessLog)
	if !ok {
		that2, ok := that.(AccessLog)
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
	if that1.OutputDestination == nil {
		if this.OutputDestination != nil {
			return false
		}
	} else if this.OutputDestination == nil {
		return false
	} else if !this.OutputDestination.Equal(that1.OutputDestination) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *AccessLog_FileSink) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AccessLog_FileSink)
	if !ok {
		that2, ok := that.(AccessLog_FileSink)
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
	if !this.FileSink.Equal(that1.FileSink) {
		return false
	}
	return true
}
func (this *FileSink) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FileSink)
	if !ok {
		that2, ok := that.(FileSink)
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
	if this.Path != that1.Path {
		return false
	}
	if that1.OutputFormat == nil {
		if this.OutputFormat != nil {
			return false
		}
	} else if this.OutputFormat == nil {
		return false
	} else if !this.OutputFormat.Equal(that1.OutputFormat) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *FileSink_StringFormat) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FileSink_StringFormat)
	if !ok {
		that2, ok := that.(FileSink_StringFormat)
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
	if this.StringFormat != that1.StringFormat {
		return false
	}
	return true
}
func (this *FileSink_JsonFormat) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FileSink_JsonFormat)
	if !ok {
		that2, ok := that.(FileSink_JsonFormat)
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
	if !this.JsonFormat.Equal(that1.JsonFormat) {
		return false
	}
	return true
}
