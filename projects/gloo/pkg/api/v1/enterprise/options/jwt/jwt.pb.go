// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/options/jwt/jwt.proto

package jwt

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	math "math"
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

type VhostExtension struct {
	// Auth providers can be used instead of the fields above where more than one is required.
	// if this list is provided the fields above are ignored.
	Providers            map[string]*Provider `protobuf:"bytes,4,rep,name=providers,proto3" json:"providers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *VhostExtension) Reset()         { *m = VhostExtension{} }
func (m *VhostExtension) String() string { return proto.CompactTextString(m) }
func (*VhostExtension) ProtoMessage()    {}
func (*VhostExtension) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d83f6c4a43394a0, []int{0}
}
func (m *VhostExtension) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VhostExtension.Unmarshal(m, b)
}
func (m *VhostExtension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VhostExtension.Marshal(b, m, deterministic)
}
func (m *VhostExtension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VhostExtension.Merge(m, src)
}
func (m *VhostExtension) XXX_Size() int {
	return xxx_messageInfo_VhostExtension.Size(m)
}
func (m *VhostExtension) XXX_DiscardUnknown() {
	xxx_messageInfo_VhostExtension.DiscardUnknown(m)
}

var xxx_messageInfo_VhostExtension proto.InternalMessageInfo

func (m *VhostExtension) GetProviders() map[string]*Provider {
	if m != nil {
		return m.Providers
	}
	return nil
}

type RouteExtension struct {
	// Disable JWT checks on this route.
	Disable              bool     `protobuf:"varint,1,opt,name=disable,proto3" json:"disable,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteExtension) Reset()         { *m = RouteExtension{} }
func (m *RouteExtension) String() string { return proto.CompactTextString(m) }
func (*RouteExtension) ProtoMessage()    {}
func (*RouteExtension) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d83f6c4a43394a0, []int{1}
}
func (m *RouteExtension) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteExtension.Unmarshal(m, b)
}
func (m *RouteExtension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteExtension.Marshal(b, m, deterministic)
}
func (m *RouteExtension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteExtension.Merge(m, src)
}
func (m *RouteExtension) XXX_Size() int {
	return xxx_messageInfo_RouteExtension.Size(m)
}
func (m *RouteExtension) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteExtension.DiscardUnknown(m)
}

var xxx_messageInfo_RouteExtension proto.InternalMessageInfo

func (m *RouteExtension) GetDisable() bool {
	if m != nil {
		return m.Disable
	}
	return false
}

type Provider struct {
	// The source for the keys to validate JWTs.
	Jwks *Jwks `protobuf:"bytes,1,opt,name=jwks,proto3" json:"jwks,omitempty"`
	// An incoming JWT must have an 'aud' claim and it must be in this list.
	Audiences []string `protobuf:"bytes,2,rep,name=audiences,proto3" json:"audiences,omitempty"`
	// Issuer of the JWT. the 'iss' claim of the JWT must match this.
	Issuer string `protobuf:"bytes,3,opt,name=issuer,proto3" json:"issuer,omitempty"`
	// Where to find the JWT of the current provider.
	TokenSource *TokenSource `protobuf:"bytes,4,opt,name=token_source,json=tokenSource,proto3" json:"token_source,omitempty"`
	// Should the token forwarded upstream. if false, the header containing the token will be removed.
	KeepToken bool `protobuf:"varint,5,opt,name=keep_token,json=keepToken,proto3" json:"keep_token,omitempty"`
	// What claims should be copied to upstream headers.
	ClaimsToHeaders      []*ClaimToHeader `protobuf:"bytes,6,rep,name=claims_to_headers,json=claimsToHeaders,proto3" json:"claims_to_headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Provider) Reset()         { *m = Provider{} }
func (m *Provider) String() string { return proto.CompactTextString(m) }
func (*Provider) ProtoMessage()    {}
func (*Provider) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d83f6c4a43394a0, []int{2}
}
func (m *Provider) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Provider.Unmarshal(m, b)
}
func (m *Provider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Provider.Marshal(b, m, deterministic)
}
func (m *Provider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Provider.Merge(m, src)
}
func (m *Provider) XXX_Size() int {
	return xxx_messageInfo_Provider.Size(m)
}
func (m *Provider) XXX_DiscardUnknown() {
	xxx_messageInfo_Provider.DiscardUnknown(m)
}

var xxx_messageInfo_Provider proto.InternalMessageInfo

func (m *Provider) GetJwks() *Jwks {
	if m != nil {
		return m.Jwks
	}
	return nil
}

func (m *Provider) GetAudiences() []string {
	if m != nil {
		return m.Audiences
	}
	return nil
}

func (m *Provider) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *Provider) GetTokenSource() *TokenSource {
	if m != nil {
		return m.TokenSource
	}
	return nil
}

func (m *Provider) GetKeepToken() bool {
	if m != nil {
		return m.KeepToken
	}
	return false
}

func (m *Provider) GetClaimsToHeaders() []*ClaimToHeader {
	if m != nil {
		return m.ClaimsToHeaders
	}
	return nil
}

type Jwks struct {
	// Types that are valid to be assigned to Jwks:
	//	*Jwks_Remote
	//	*Jwks_Local
	Jwks                 isJwks_Jwks `protobuf_oneof:"jwks"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Jwks) Reset()         { *m = Jwks{} }
func (m *Jwks) String() string { return proto.CompactTextString(m) }
func (*Jwks) ProtoMessage()    {}
func (*Jwks) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d83f6c4a43394a0, []int{3}
}
func (m *Jwks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Jwks.Unmarshal(m, b)
}
func (m *Jwks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Jwks.Marshal(b, m, deterministic)
}
func (m *Jwks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Jwks.Merge(m, src)
}
func (m *Jwks) XXX_Size() int {
	return xxx_messageInfo_Jwks.Size(m)
}
func (m *Jwks) XXX_DiscardUnknown() {
	xxx_messageInfo_Jwks.DiscardUnknown(m)
}

var xxx_messageInfo_Jwks proto.InternalMessageInfo

type isJwks_Jwks interface {
	isJwks_Jwks()
	Equal(interface{}) bool
}

type Jwks_Remote struct {
	Remote *RemoteJwks `protobuf:"bytes,1,opt,name=remote,proto3,oneof" json:"remote,omitempty"`
}
type Jwks_Local struct {
	Local *LocalJwks `protobuf:"bytes,2,opt,name=local,proto3,oneof" json:"local,omitempty"`
}

func (*Jwks_Remote) isJwks_Jwks() {}
func (*Jwks_Local) isJwks_Jwks()  {}

func (m *Jwks) GetJwks() isJwks_Jwks {
	if m != nil {
		return m.Jwks
	}
	return nil
}

func (m *Jwks) GetRemote() *RemoteJwks {
	if x, ok := m.GetJwks().(*Jwks_Remote); ok {
		return x.Remote
	}
	return nil
}

func (m *Jwks) GetLocal() *LocalJwks {
	if x, ok := m.GetJwks().(*Jwks_Local); ok {
		return x.Local
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Jwks) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Jwks_Remote)(nil),
		(*Jwks_Local)(nil),
	}
}

type RemoteJwks struct {
	// The url used when accessing the upstream for Json Web Key Set.
	// This is used to set the host and path in the request
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// The Upstream representing the Json Web Key Set server
	UpstreamRef *core.ResourceRef `protobuf:"bytes,2,opt,name=upstream_ref,json=upstreamRef,proto3" json:"upstream_ref,omitempty"`
	// Duration after which the cached JWKS should be expired.
	// If not specified, default cache duration is 5 minutes.
	CacheDuration        *types.Duration `protobuf:"bytes,4,opt,name=cache_duration,json=cacheDuration,proto3" json:"cache_duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RemoteJwks) Reset()         { *m = RemoteJwks{} }
func (m *RemoteJwks) String() string { return proto.CompactTextString(m) }
func (*RemoteJwks) ProtoMessage()    {}
func (*RemoteJwks) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d83f6c4a43394a0, []int{4}
}
func (m *RemoteJwks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoteJwks.Unmarshal(m, b)
}
func (m *RemoteJwks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoteJwks.Marshal(b, m, deterministic)
}
func (m *RemoteJwks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteJwks.Merge(m, src)
}
func (m *RemoteJwks) XXX_Size() int {
	return xxx_messageInfo_RemoteJwks.Size(m)
}
func (m *RemoteJwks) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteJwks.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteJwks proto.InternalMessageInfo

func (m *RemoteJwks) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *RemoteJwks) GetUpstreamRef() *core.ResourceRef {
	if m != nil {
		return m.UpstreamRef
	}
	return nil
}

func (m *RemoteJwks) GetCacheDuration() *types.Duration {
	if m != nil {
		return m.CacheDuration
	}
	return nil
}

type LocalJwks struct {
	// Inline key. this can be json web key, key-set or PEM format.
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocalJwks) Reset()         { *m = LocalJwks{} }
func (m *LocalJwks) String() string { return proto.CompactTextString(m) }
func (*LocalJwks) ProtoMessage()    {}
func (*LocalJwks) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d83f6c4a43394a0, []int{5}
}
func (m *LocalJwks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalJwks.Unmarshal(m, b)
}
func (m *LocalJwks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalJwks.Marshal(b, m, deterministic)
}
func (m *LocalJwks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalJwks.Merge(m, src)
}
func (m *LocalJwks) XXX_Size() int {
	return xxx_messageInfo_LocalJwks.Size(m)
}
func (m *LocalJwks) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalJwks.DiscardUnknown(m)
}

var xxx_messageInfo_LocalJwks proto.InternalMessageInfo

func (m *LocalJwks) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

// Describes the location of a JWT token
type TokenSource struct {
	// Try to retrieve token from these headers
	Headers []*TokenSource_HeaderSource `protobuf:"bytes,1,rep,name=headers,proto3" json:"headers,omitempty"`
	// Try to retrieve token from these query params
	QueryParams          []string `protobuf:"bytes,2,rep,name=query_params,json=queryParams,proto3" json:"query_params,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TokenSource) Reset()         { *m = TokenSource{} }
func (m *TokenSource) String() string { return proto.CompactTextString(m) }
func (*TokenSource) ProtoMessage()    {}
func (*TokenSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d83f6c4a43394a0, []int{6}
}
func (m *TokenSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenSource.Unmarshal(m, b)
}
func (m *TokenSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenSource.Marshal(b, m, deterministic)
}
func (m *TokenSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenSource.Merge(m, src)
}
func (m *TokenSource) XXX_Size() int {
	return xxx_messageInfo_TokenSource.Size(m)
}
func (m *TokenSource) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenSource.DiscardUnknown(m)
}

var xxx_messageInfo_TokenSource proto.InternalMessageInfo

func (m *TokenSource) GetHeaders() []*TokenSource_HeaderSource {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *TokenSource) GetQueryParams() []string {
	if m != nil {
		return m.QueryParams
	}
	return nil
}

// Describes how to retrieve a JWT from a header
type TokenSource_HeaderSource struct {
	// The name of the header. for example, "authorization"
	Header string `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// Prefix before the token. for example, "Bearer "
	Prefix               string   `protobuf:"bytes,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TokenSource_HeaderSource) Reset()         { *m = TokenSource_HeaderSource{} }
func (m *TokenSource_HeaderSource) String() string { return proto.CompactTextString(m) }
func (*TokenSource_HeaderSource) ProtoMessage()    {}
func (*TokenSource_HeaderSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d83f6c4a43394a0, []int{6, 0}
}
func (m *TokenSource_HeaderSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenSource_HeaderSource.Unmarshal(m, b)
}
func (m *TokenSource_HeaderSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenSource_HeaderSource.Marshal(b, m, deterministic)
}
func (m *TokenSource_HeaderSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenSource_HeaderSource.Merge(m, src)
}
func (m *TokenSource_HeaderSource) XXX_Size() int {
	return xxx_messageInfo_TokenSource_HeaderSource.Size(m)
}
func (m *TokenSource_HeaderSource) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenSource_HeaderSource.DiscardUnknown(m)
}

var xxx_messageInfo_TokenSource_HeaderSource proto.InternalMessageInfo

func (m *TokenSource_HeaderSource) GetHeader() string {
	if m != nil {
		return m.Header
	}
	return ""
}

func (m *TokenSource_HeaderSource) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

// Allows copying verified claims to headers sent upstream
type ClaimToHeader struct {
	// Claim name. for example, "sub"
	Claim string `protobuf:"bytes,1,opt,name=claim,proto3" json:"claim,omitempty"`
	// The header the claim will be copied to. for example, "x-sub".
	Header string `protobuf:"bytes,2,opt,name=header,proto3" json:"header,omitempty"`
	// If the header exists, append to it (true), or overwrite it (false).
	Append               bool     `protobuf:"varint,4,opt,name=append,proto3" json:"append,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClaimToHeader) Reset()         { *m = ClaimToHeader{} }
func (m *ClaimToHeader) String() string { return proto.CompactTextString(m) }
func (*ClaimToHeader) ProtoMessage()    {}
func (*ClaimToHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d83f6c4a43394a0, []int{7}
}
func (m *ClaimToHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClaimToHeader.Unmarshal(m, b)
}
func (m *ClaimToHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClaimToHeader.Marshal(b, m, deterministic)
}
func (m *ClaimToHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimToHeader.Merge(m, src)
}
func (m *ClaimToHeader) XXX_Size() int {
	return xxx_messageInfo_ClaimToHeader.Size(m)
}
func (m *ClaimToHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimToHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimToHeader proto.InternalMessageInfo

func (m *ClaimToHeader) GetClaim() string {
	if m != nil {
		return m.Claim
	}
	return ""
}

func (m *ClaimToHeader) GetHeader() string {
	if m != nil {
		return m.Header
	}
	return ""
}

func (m *ClaimToHeader) GetAppend() bool {
	if m != nil {
		return m.Append
	}
	return false
}

func init() {
	proto.RegisterType((*VhostExtension)(nil), "jwt.options.gloo.solo.io.VhostExtension")
	proto.RegisterMapType((map[string]*Provider)(nil), "jwt.options.gloo.solo.io.VhostExtension.ProvidersEntry")
	proto.RegisterType((*RouteExtension)(nil), "jwt.options.gloo.solo.io.RouteExtension")
	proto.RegisterType((*Provider)(nil), "jwt.options.gloo.solo.io.Provider")
	proto.RegisterType((*Jwks)(nil), "jwt.options.gloo.solo.io.Jwks")
	proto.RegisterType((*RemoteJwks)(nil), "jwt.options.gloo.solo.io.RemoteJwks")
	proto.RegisterType((*LocalJwks)(nil), "jwt.options.gloo.solo.io.LocalJwks")
	proto.RegisterType((*TokenSource)(nil), "jwt.options.gloo.solo.io.TokenSource")
	proto.RegisterType((*TokenSource_HeaderSource)(nil), "jwt.options.gloo.solo.io.TokenSource.HeaderSource")
	proto.RegisterType((*ClaimToHeader)(nil), "jwt.options.gloo.solo.io.ClaimToHeader")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/options/jwt/jwt.proto", fileDescriptor_3d83f6c4a43394a0)
}

var fileDescriptor_3d83f6c4a43394a0 = []byte{
	// 686 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x26, 0xfd, 0xdb, 0x7a, 0xda, 0x15, 0xb0, 0x26, 0x94, 0x55, 0x6c, 0x2a, 0x01, 0x44, 0x85,
	0x20, 0x11, 0xe5, 0x82, 0x09, 0xd0, 0x84, 0x06, 0x13, 0x13, 0xda, 0xa4, 0xc9, 0xdb, 0xb8, 0xe0,
	0xa6, 0xa4, 0xe9, 0x69, 0x9b, 0x35, 0x8d, 0x83, 0xed, 0xec, 0xe7, 0x19, 0x78, 0x07, 0xae, 0xb9,
	0xe6, 0x01, 0xb8, 0xe7, 0x35, 0x78, 0x12, 0x64, 0xc7, 0xa1, 0x1d, 0x10, 0xb4, 0x8b, 0xaa, 0xfe,
	0x8e, 0xcf, 0xf7, 0x1d, 0xfb, 0x7c, 0xc7, 0x81, 0xfd, 0x71, 0x28, 0x27, 0xe9, 0xc0, 0x0d, 0xd8,
	0xcc, 0x13, 0x2c, 0x62, 0x8f, 0x43, 0xe6, 0x8d, 0x23, 0xc6, 0xbc, 0x84, 0xb3, 0x13, 0x0c, 0xa4,
	0xc8, 0x90, 0x9f, 0x84, 0xde, 0xe9, 0x13, 0x0f, 0x63, 0x89, 0x3c, 0xe1, 0xa1, 0x40, 0x8f, 0x25,
	0x32, 0x64, 0xb1, 0xf0, 0x4e, 0xce, 0xa4, 0xfa, 0xb9, 0x09, 0x67, 0x92, 0x11, 0x5b, 0x2d, 0xcd,
	0x96, 0xab, 0x98, 0xae, 0x12, 0x75, 0x43, 0xd6, 0x7e, 0xf4, 0x8f, 0x42, 0xfa, 0x7f, 0x1a, 0xca,
	0x5c, 0x9e, 0xe3, 0x28, 0xd3, 0x69, 0xaf, 0x8e, 0xd9, 0x98, 0xe9, 0xa5, 0xa7, 0x56, 0x26, 0xba,
	0x31, 0x66, 0x6c, 0x1c, 0xa1, 0xa7, 0xd1, 0x20, 0x1d, 0x79, 0xc3, 0x94, 0xfb, 0xaa, 0x56, 0xb6,
	0xef, 0xfc, 0xb0, 0xa0, 0xf5, 0x7e, 0xc2, 0x84, 0xdc, 0x39, 0x97, 0x18, 0x8b, 0x90, 0xc5, 0xe4,
	0x18, 0xea, 0x09, 0x67, 0xa7, 0xe1, 0x10, 0xb9, 0xb0, 0x2b, 0x9d, 0x72, 0xb7, 0xd1, 0x7b, 0xe6,
	0x16, 0x1d, 0xd2, 0xbd, 0x4c, 0x76, 0x0f, 0x72, 0xe6, 0x4e, 0x2c, 0xf9, 0x05, 0x9d, 0x2b, 0xb5,
	0x3f, 0x42, 0xeb, 0xf2, 0x26, 0xb9, 0x01, 0xe5, 0x29, 0x5e, 0xd8, 0x56, 0xc7, 0xea, 0xd6, 0xa9,
	0x5a, 0x92, 0x4d, 0xa8, 0x9e, 0xfa, 0x51, 0x8a, 0x76, 0xa9, 0x63, 0x75, 0x1b, 0x3d, 0xa7, 0xb8,
	0x6c, 0x2e, 0x45, 0x33, 0xc2, 0xf3, 0xd2, 0xa6, 0xe5, 0x3c, 0x84, 0x16, 0x65, 0xa9, 0xc4, 0xf9,
	0x55, 0x6c, 0x58, 0x1a, 0x86, 0xc2, 0x1f, 0x44, 0xa8, 0xab, 0x2c, 0xd3, 0x1c, 0x3a, 0xdf, 0x4a,
	0xb0, 0x9c, 0x6b, 0x90, 0x1e, 0x54, 0x4e, 0xce, 0xa6, 0x42, 0xe7, 0x34, 0x7a, 0x1b, 0xc5, 0x55,
	0xdf, 0x9d, 0x4d, 0x05, 0xd5, 0xb9, 0xe4, 0x36, 0xd4, 0xfd, 0x74, 0x18, 0x62, 0x1c, 0xa0, 0xb0,
	0x4b, 0x9d, 0x72, 0xb7, 0x4e, 0xe7, 0x01, 0x72, 0x0b, 0x6a, 0xa1, 0x10, 0x29, 0x72, 0xbb, 0xac,
	0x6f, 0x67, 0x10, 0xd9, 0x85, 0xa6, 0x64, 0x53, 0x8c, 0xfb, 0x82, 0xa5, 0x3c, 0x40, 0xbb, 0xa2,
	0x2b, 0xde, 0x2f, 0xae, 0x78, 0xa4, 0xb2, 0x0f, 0x75, 0x32, 0x6d, 0xc8, 0x39, 0x20, 0xeb, 0x00,
	0x53, 0xc4, 0xa4, 0xaf, 0x63, 0x76, 0x55, 0xdf, 0xae, 0xae, 0x22, 0x9a, 0x41, 0x0e, 0xe1, 0x66,
	0x10, 0xf9, 0xe1, 0x4c, 0xf4, 0x25, 0xeb, 0x4f, 0xd0, 0xd7, 0x66, 0xd6, 0xb4, 0x99, 0x0f, 0x8a,
	0xab, 0xbd, 0x56, 0x94, 0x23, 0xb6, 0xab, 0xf3, 0xe9, 0xf5, 0x4c, 0x21, 0xc7, 0xc2, 0xf9, 0x6c,
	0x41, 0x45, 0xb5, 0x80, 0x6c, 0x41, 0x8d, 0xe3, 0x8c, 0x49, 0x34, 0x2d, 0xbb, 0x57, 0x2c, 0x49,
	0x75, 0x9e, 0x62, 0xed, 0x5e, 0xa3, 0x86, 0x45, 0x5e, 0x40, 0x35, 0x62, 0x81, 0x1f, 0x19, 0x9f,
	0xef, 0x16, 0xd3, 0xf7, 0x54, 0x9a, 0x61, 0x67, 0x9c, 0xed, 0x5a, 0xe6, 0x96, 0xf3, 0xc5, 0x02,
	0x98, 0xab, 0xab, 0x69, 0x4a, 0x79, 0x94, 0x4f, 0x53, 0xca, 0x23, 0xf2, 0x12, 0x9a, 0x69, 0x22,
	0x24, 0x47, 0x7f, 0xd6, 0xe7, 0x38, 0x32, 0xc5, 0xd6, 0xdc, 0x80, 0x71, 0x5c, 0x38, 0x5f, 0x66,
	0x05, 0xc5, 0x11, 0x6d, 0xe4, 0xe9, 0x14, 0x47, 0xe4, 0x15, 0xb4, 0x02, 0x3f, 0x98, 0x60, 0x3f,
	0x7f, 0x31, 0xc6, 0xac, 0x35, 0x37, 0x7b, 0x52, 0x6e, 0xfe, 0xa4, 0xdc, 0x37, 0x26, 0x81, 0xae,
	0x68, 0x42, 0x0e, 0x9d, 0x75, 0xa8, 0xff, 0x3e, 0xfe, 0xdf, 0xc3, 0xee, 0x7c, 0xb7, 0xa0, 0xb1,
	0x60, 0x2f, 0xd9, 0x83, 0xa5, 0xdc, 0x28, 0x4b, 0x1b, 0xd5, 0xbb, 0xd2, 0x58, 0xb8, 0x99, 0x3b,
	0x66, 0x46, 0x72, 0x09, 0x72, 0x07, 0x9a, 0x9f, 0x52, 0xe4, 0x17, 0xfd, 0xc4, 0xe7, 0xfe, 0x2c,
	0x1f, 0xd1, 0x86, 0x8e, 0x1d, 0xe8, 0x50, 0x7b, 0x0b, 0x9a, 0x8b, 0x5c, 0x35, 0xb4, 0x19, 0xdb,
	0x9c, 0xd2, 0x20, 0x15, 0x4f, 0x38, 0x8e, 0xc2, 0x73, 0xdd, 0xc1, 0x3a, 0x35, 0xc8, 0x39, 0x86,
	0x95, 0x4b, 0x03, 0x43, 0x56, 0xa1, 0xaa, 0x47, 0xc6, 0xf0, 0x33, 0xb0, 0x20, 0x5b, 0xfa, 0x53,
	0xd6, 0x4f, 0x12, 0x8c, 0x87, 0xba, 0xb1, 0xcb, 0xd4, 0xa0, 0xed, 0xfd, 0xaf, 0x3f, 0x37, 0xac,
	0x0f, 0x6f, 0xaf, 0xf6, 0x95, 0x4d, 0xa6, 0xe3, 0xff, 0x7f, 0x69, 0x07, 0x35, 0xed, 0xd3, 0xd3,
	0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x43, 0x7d, 0x46, 0xc7, 0xb7, 0x05, 0x00, 0x00,
}

func (this *VhostExtension) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VhostExtension)
	if !ok {
		that2, ok := that.(VhostExtension)
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
	if len(this.Providers) != len(that1.Providers) {
		return false
	}
	for i := range this.Providers {
		if !this.Providers[i].Equal(that1.Providers[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RouteExtension) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RouteExtension)
	if !ok {
		that2, ok := that.(RouteExtension)
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
	if this.Disable != that1.Disable {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Provider) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Provider)
	if !ok {
		that2, ok := that.(Provider)
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
	if !this.Jwks.Equal(that1.Jwks) {
		return false
	}
	if len(this.Audiences) != len(that1.Audiences) {
		return false
	}
	for i := range this.Audiences {
		if this.Audiences[i] != that1.Audiences[i] {
			return false
		}
	}
	if this.Issuer != that1.Issuer {
		return false
	}
	if !this.TokenSource.Equal(that1.TokenSource) {
		return false
	}
	if this.KeepToken != that1.KeepToken {
		return false
	}
	if len(this.ClaimsToHeaders) != len(that1.ClaimsToHeaders) {
		return false
	}
	for i := range this.ClaimsToHeaders {
		if !this.ClaimsToHeaders[i].Equal(that1.ClaimsToHeaders[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Jwks) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Jwks)
	if !ok {
		that2, ok := that.(Jwks)
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
	if that1.Jwks == nil {
		if this.Jwks != nil {
			return false
		}
	} else if this.Jwks == nil {
		return false
	} else if !this.Jwks.Equal(that1.Jwks) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Jwks_Remote) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Jwks_Remote)
	if !ok {
		that2, ok := that.(Jwks_Remote)
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
	if !this.Remote.Equal(that1.Remote) {
		return false
	}
	return true
}
func (this *Jwks_Local) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Jwks_Local)
	if !ok {
		that2, ok := that.(Jwks_Local)
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
	if !this.Local.Equal(that1.Local) {
		return false
	}
	return true
}
func (this *RemoteJwks) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RemoteJwks)
	if !ok {
		that2, ok := that.(RemoteJwks)
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
	if this.Url != that1.Url {
		return false
	}
	if !this.UpstreamRef.Equal(that1.UpstreamRef) {
		return false
	}
	if !this.CacheDuration.Equal(that1.CacheDuration) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *LocalJwks) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LocalJwks)
	if !ok {
		that2, ok := that.(LocalJwks)
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
	if this.Key != that1.Key {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *TokenSource) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TokenSource)
	if !ok {
		that2, ok := that.(TokenSource)
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
	if len(this.Headers) != len(that1.Headers) {
		return false
	}
	for i := range this.Headers {
		if !this.Headers[i].Equal(that1.Headers[i]) {
			return false
		}
	}
	if len(this.QueryParams) != len(that1.QueryParams) {
		return false
	}
	for i := range this.QueryParams {
		if this.QueryParams[i] != that1.QueryParams[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *TokenSource_HeaderSource) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TokenSource_HeaderSource)
	if !ok {
		that2, ok := that.(TokenSource_HeaderSource)
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
	if this.Header != that1.Header {
		return false
	}
	if this.Prefix != that1.Prefix {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ClaimToHeader) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ClaimToHeader)
	if !ok {
		that2, ok := that.(ClaimToHeader)
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
	if this.Claim != that1.Claim {
		return false
	}
	if this.Header != that1.Header {
		return false
	}
	if this.Append != that1.Append {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
