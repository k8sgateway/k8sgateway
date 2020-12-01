// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/ssl.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
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

type SslParameters_ProtocolVersion int32

const (
	// Envoy will choose the optimal TLS version.
	SslParameters_TLS_AUTO SslParameters_ProtocolVersion = 0
	// TLS 1.0
	SslParameters_TLSv1_0 SslParameters_ProtocolVersion = 1
	// TLS 1.1
	SslParameters_TLSv1_1 SslParameters_ProtocolVersion = 2
	// TLS 1.2
	SslParameters_TLSv1_2 SslParameters_ProtocolVersion = 3
	// TLS 1.3
	SslParameters_TLSv1_3 SslParameters_ProtocolVersion = 4
)

var SslParameters_ProtocolVersion_name = map[int32]string{
	0: "TLS_AUTO",
	1: "TLSv1_0",
	2: "TLSv1_1",
	3: "TLSv1_2",
	4: "TLSv1_3",
}

var SslParameters_ProtocolVersion_value = map[string]int32{
	"TLS_AUTO": 0,
	"TLSv1_0":  1,
	"TLSv1_1":  2,
	"TLSv1_2":  3,
	"TLSv1_3":  4,
}

func (x SslParameters_ProtocolVersion) String() string {
	return proto.EnumName(SslParameters_ProtocolVersion_name, int32(x))
}

func (SslParameters_ProtocolVersion) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c4a65e8067d81add, []int{5, 0}
}

// SslConfig contains the options necessary to configure a virtual host or listener to use TLS
type SslConfig struct {
	// Types that are valid to be assigned to SslSecrets:
	//	*SslConfig_SecretRef
	//	*SslConfig_SslFiles
	//	*SslConfig_Sds
	SslSecrets isSslConfig_SslSecrets `protobuf_oneof:"ssl_secrets"`
	// optional. the SNI domains that should be considered for TLS connections
	SniDomains []string `protobuf:"bytes,3,rep,name=sni_domains,json=sniDomains,proto3" json:"sni_domains,omitempty"`
	// Verify that the Subject Alternative Name in the peer certificate is one of the specified values.
	// note that a root_ca must be provided if this option is used.
	VerifySubjectAltName []string       `protobuf:"bytes,5,rep,name=verify_subject_alt_name,json=verifySubjectAltName,proto3" json:"verify_subject_alt_name,omitempty"`
	Parameters           *SslParameters `protobuf:"bytes,6,opt,name=parameters,proto3" json:"parameters,omitempty"`
	// Set Application Level Protocol Negotiation
	// If empty, defaults to ["h2", "http/1.1"].
	AlpnProtocols        []string `protobuf:"bytes,7,rep,name=alpn_protocols,json=alpnProtocols,proto3" json:"alpn_protocols,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SslConfig) Reset()         { *m = SslConfig{} }
func (m *SslConfig) String() string { return proto.CompactTextString(m) }
func (*SslConfig) ProtoMessage()    {}
func (*SslConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4a65e8067d81add, []int{0}
}
func (m *SslConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SslConfig.Unmarshal(m, b)
}
func (m *SslConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SslConfig.Marshal(b, m, deterministic)
}
func (m *SslConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SslConfig.Merge(m, src)
}
func (m *SslConfig) XXX_Size() int {
	return xxx_messageInfo_SslConfig.Size(m)
}
func (m *SslConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_SslConfig.DiscardUnknown(m)
}

var xxx_messageInfo_SslConfig proto.InternalMessageInfo

type isSslConfig_SslSecrets interface {
	isSslConfig_SslSecrets()
	Equal(interface{}) bool
}

type SslConfig_SecretRef struct {
	SecretRef *core.ResourceRef `protobuf:"bytes,1,opt,name=secret_ref,json=secretRef,proto3,oneof" json:"secret_ref,omitempty"`
}
type SslConfig_SslFiles struct {
	SslFiles *SSLFiles `protobuf:"bytes,2,opt,name=ssl_files,json=sslFiles,proto3,oneof" json:"ssl_files,omitempty"`
}
type SslConfig_Sds struct {
	Sds *SDSConfig `protobuf:"bytes,4,opt,name=sds,proto3,oneof" json:"sds,omitempty"`
}

func (*SslConfig_SecretRef) isSslConfig_SslSecrets() {}
func (*SslConfig_SslFiles) isSslConfig_SslSecrets()  {}
func (*SslConfig_Sds) isSslConfig_SslSecrets()       {}

func (m *SslConfig) GetSslSecrets() isSslConfig_SslSecrets {
	if m != nil {
		return m.SslSecrets
	}
	return nil
}

func (m *SslConfig) GetSecretRef() *core.ResourceRef {
	if x, ok := m.GetSslSecrets().(*SslConfig_SecretRef); ok {
		return x.SecretRef
	}
	return nil
}

func (m *SslConfig) GetSslFiles() *SSLFiles {
	if x, ok := m.GetSslSecrets().(*SslConfig_SslFiles); ok {
		return x.SslFiles
	}
	return nil
}

func (m *SslConfig) GetSds() *SDSConfig {
	if x, ok := m.GetSslSecrets().(*SslConfig_Sds); ok {
		return x.Sds
	}
	return nil
}

func (m *SslConfig) GetSniDomains() []string {
	if m != nil {
		return m.SniDomains
	}
	return nil
}

func (m *SslConfig) GetVerifySubjectAltName() []string {
	if m != nil {
		return m.VerifySubjectAltName
	}
	return nil
}

func (m *SslConfig) GetParameters() *SslParameters {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *SslConfig) GetAlpnProtocols() []string {
	if m != nil {
		return m.AlpnProtocols
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SslConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SslConfig_SecretRef)(nil),
		(*SslConfig_SslFiles)(nil),
		(*SslConfig_Sds)(nil),
	}
}

// SSLFiles reference paths to certificates which can be read by the proxy off of its local filesystem
type SSLFiles struct {
	TlsCert string `protobuf:"bytes,1,opt,name=tls_cert,json=tlsCert,proto3" json:"tls_cert,omitempty"`
	TlsKey  string `protobuf:"bytes,2,opt,name=tls_key,json=tlsKey,proto3" json:"tls_key,omitempty"`
	// for client cert validation. optional
	RootCa               string   `protobuf:"bytes,3,opt,name=root_ca,json=rootCa,proto3" json:"root_ca,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SSLFiles) Reset()         { *m = SSLFiles{} }
func (m *SSLFiles) String() string { return proto.CompactTextString(m) }
func (*SSLFiles) ProtoMessage()    {}
func (*SSLFiles) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4a65e8067d81add, []int{1}
}
func (m *SSLFiles) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SSLFiles.Unmarshal(m, b)
}
func (m *SSLFiles) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SSLFiles.Marshal(b, m, deterministic)
}
func (m *SSLFiles) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSLFiles.Merge(m, src)
}
func (m *SSLFiles) XXX_Size() int {
	return xxx_messageInfo_SSLFiles.Size(m)
}
func (m *SSLFiles) XXX_DiscardUnknown() {
	xxx_messageInfo_SSLFiles.DiscardUnknown(m)
}

var xxx_messageInfo_SSLFiles proto.InternalMessageInfo

func (m *SSLFiles) GetTlsCert() string {
	if m != nil {
		return m.TlsCert
	}
	return ""
}

func (m *SSLFiles) GetTlsKey() string {
	if m != nil {
		return m.TlsKey
	}
	return ""
}

func (m *SSLFiles) GetRootCa() string {
	if m != nil {
		return m.RootCa
	}
	return ""
}

// SslConfig contains the options necessary to configure a virtual host or listener to use TLS
type UpstreamSslConfig struct {
	// Types that are valid to be assigned to SslSecrets:
	//	*UpstreamSslConfig_SecretRef
	//	*UpstreamSslConfig_SslFiles
	//	*UpstreamSslConfig_Sds
	SslSecrets isUpstreamSslConfig_SslSecrets `protobuf_oneof:"ssl_secrets"`
	// optional. the SNI domains that should be considered for TLS connections
	Sni string `protobuf:"bytes,3,opt,name=sni,proto3" json:"sni,omitempty"`
	// Verify that the Subject Alternative Name in the peer certificate is one of the specified values.
	// note that a root_ca must be provided if this option is used.
	VerifySubjectAltName []string       `protobuf:"bytes,5,rep,name=verify_subject_alt_name,json=verifySubjectAltName,proto3" json:"verify_subject_alt_name,omitempty"`
	Parameters           *SslParameters `protobuf:"bytes,7,opt,name=parameters,proto3" json:"parameters,omitempty"`
	// Set Application Level Protocol Negotiation.
	// If empty, it is not set.
	AlpnProtocols        []string `protobuf:"bytes,8,rep,name=alpn_protocols,json=alpnProtocols,proto3" json:"alpn_protocols,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpstreamSslConfig) Reset()         { *m = UpstreamSslConfig{} }
func (m *UpstreamSslConfig) String() string { return proto.CompactTextString(m) }
func (*UpstreamSslConfig) ProtoMessage()    {}
func (*UpstreamSslConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4a65e8067d81add, []int{2}
}
func (m *UpstreamSslConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpstreamSslConfig.Unmarshal(m, b)
}
func (m *UpstreamSslConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpstreamSslConfig.Marshal(b, m, deterministic)
}
func (m *UpstreamSslConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpstreamSslConfig.Merge(m, src)
}
func (m *UpstreamSslConfig) XXX_Size() int {
	return xxx_messageInfo_UpstreamSslConfig.Size(m)
}
func (m *UpstreamSslConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_UpstreamSslConfig.DiscardUnknown(m)
}

var xxx_messageInfo_UpstreamSslConfig proto.InternalMessageInfo

type isUpstreamSslConfig_SslSecrets interface {
	isUpstreamSslConfig_SslSecrets()
	Equal(interface{}) bool
}

type UpstreamSslConfig_SecretRef struct {
	SecretRef *core.ResourceRef `protobuf:"bytes,1,opt,name=secret_ref,json=secretRef,proto3,oneof" json:"secret_ref,omitempty"`
}
type UpstreamSslConfig_SslFiles struct {
	SslFiles *SSLFiles `protobuf:"bytes,2,opt,name=ssl_files,json=sslFiles,proto3,oneof" json:"ssl_files,omitempty"`
}
type UpstreamSslConfig_Sds struct {
	Sds *SDSConfig `protobuf:"bytes,4,opt,name=sds,proto3,oneof" json:"sds,omitempty"`
}

func (*UpstreamSslConfig_SecretRef) isUpstreamSslConfig_SslSecrets() {}
func (*UpstreamSslConfig_SslFiles) isUpstreamSslConfig_SslSecrets()  {}
func (*UpstreamSslConfig_Sds) isUpstreamSslConfig_SslSecrets()       {}

func (m *UpstreamSslConfig) GetSslSecrets() isUpstreamSslConfig_SslSecrets {
	if m != nil {
		return m.SslSecrets
	}
	return nil
}

func (m *UpstreamSslConfig) GetSecretRef() *core.ResourceRef {
	if x, ok := m.GetSslSecrets().(*UpstreamSslConfig_SecretRef); ok {
		return x.SecretRef
	}
	return nil
}

func (m *UpstreamSslConfig) GetSslFiles() *SSLFiles {
	if x, ok := m.GetSslSecrets().(*UpstreamSslConfig_SslFiles); ok {
		return x.SslFiles
	}
	return nil
}

func (m *UpstreamSslConfig) GetSds() *SDSConfig {
	if x, ok := m.GetSslSecrets().(*UpstreamSslConfig_Sds); ok {
		return x.Sds
	}
	return nil
}

func (m *UpstreamSslConfig) GetSni() string {
	if m != nil {
		return m.Sni
	}
	return ""
}

func (m *UpstreamSslConfig) GetVerifySubjectAltName() []string {
	if m != nil {
		return m.VerifySubjectAltName
	}
	return nil
}

func (m *UpstreamSslConfig) GetParameters() *SslParameters {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *UpstreamSslConfig) GetAlpnProtocols() []string {
	if m != nil {
		return m.AlpnProtocols
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*UpstreamSslConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*UpstreamSslConfig_SecretRef)(nil),
		(*UpstreamSslConfig_SslFiles)(nil),
		(*UpstreamSslConfig_Sds)(nil),
	}
}

type SDSConfig struct {
	// Target uri for the sds channel. currently only a unix domain socket is supported.
	TargetUri string `protobuf:"bytes,1,opt,name=target_uri,json=targetUri,proto3" json:"target_uri,omitempty"`
	// Types that are valid to be assigned to SdsBuilder:
	//	*SDSConfig_CallCredentials
	//	*SDSConfig_ClusterName
	SdsBuilder isSDSConfig_SdsBuilder `protobuf_oneof:"sds_builder"`
	// The name of the secret containing the certificate
	CertificatesSecretName string `protobuf:"bytes,3,opt,name=certificates_secret_name,json=certificatesSecretName,proto3" json:"certificates_secret_name,omitempty"`
	// The name of secret containing the validation context (i.e. root ca)
	ValidationContextName string   `protobuf:"bytes,4,opt,name=validation_context_name,json=validationContextName,proto3" json:"validation_context_name,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *SDSConfig) Reset()         { *m = SDSConfig{} }
func (m *SDSConfig) String() string { return proto.CompactTextString(m) }
func (*SDSConfig) ProtoMessage()    {}
func (*SDSConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4a65e8067d81add, []int{3}
}
func (m *SDSConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SDSConfig.Unmarshal(m, b)
}
func (m *SDSConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SDSConfig.Marshal(b, m, deterministic)
}
func (m *SDSConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SDSConfig.Merge(m, src)
}
func (m *SDSConfig) XXX_Size() int {
	return xxx_messageInfo_SDSConfig.Size(m)
}
func (m *SDSConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_SDSConfig.DiscardUnknown(m)
}

var xxx_messageInfo_SDSConfig proto.InternalMessageInfo

type isSDSConfig_SdsBuilder interface {
	isSDSConfig_SdsBuilder()
	Equal(interface{}) bool
}

type SDSConfig_CallCredentials struct {
	CallCredentials *CallCredentials `protobuf:"bytes,2,opt,name=call_credentials,json=callCredentials,proto3,oneof" json:"call_credentials,omitempty"`
}
type SDSConfig_ClusterName struct {
	ClusterName string `protobuf:"bytes,5,opt,name=cluster_name,json=clusterName,proto3,oneof" json:"cluster_name,omitempty"`
}

func (*SDSConfig_CallCredentials) isSDSConfig_SdsBuilder() {}
func (*SDSConfig_ClusterName) isSDSConfig_SdsBuilder()     {}

func (m *SDSConfig) GetSdsBuilder() isSDSConfig_SdsBuilder {
	if m != nil {
		return m.SdsBuilder
	}
	return nil
}

func (m *SDSConfig) GetTargetUri() string {
	if m != nil {
		return m.TargetUri
	}
	return ""
}

func (m *SDSConfig) GetCallCredentials() *CallCredentials {
	if x, ok := m.GetSdsBuilder().(*SDSConfig_CallCredentials); ok {
		return x.CallCredentials
	}
	return nil
}

func (m *SDSConfig) GetClusterName() string {
	if x, ok := m.GetSdsBuilder().(*SDSConfig_ClusterName); ok {
		return x.ClusterName
	}
	return ""
}

func (m *SDSConfig) GetCertificatesSecretName() string {
	if m != nil {
		return m.CertificatesSecretName
	}
	return ""
}

func (m *SDSConfig) GetValidationContextName() string {
	if m != nil {
		return m.ValidationContextName
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SDSConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SDSConfig_CallCredentials)(nil),
		(*SDSConfig_ClusterName)(nil),
	}
}

type CallCredentials struct {
	// Call credentials are coming from a file,
	FileCredentialSource *CallCredentials_FileCredentialSource `protobuf:"bytes,1,opt,name=file_credential_source,json=fileCredentialSource,proto3" json:"file_credential_source,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *CallCredentials) Reset()         { *m = CallCredentials{} }
func (m *CallCredentials) String() string { return proto.CompactTextString(m) }
func (*CallCredentials) ProtoMessage()    {}
func (*CallCredentials) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4a65e8067d81add, []int{4}
}
func (m *CallCredentials) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallCredentials.Unmarshal(m, b)
}
func (m *CallCredentials) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallCredentials.Marshal(b, m, deterministic)
}
func (m *CallCredentials) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallCredentials.Merge(m, src)
}
func (m *CallCredentials) XXX_Size() int {
	return xxx_messageInfo_CallCredentials.Size(m)
}
func (m *CallCredentials) XXX_DiscardUnknown() {
	xxx_messageInfo_CallCredentials.DiscardUnknown(m)
}

var xxx_messageInfo_CallCredentials proto.InternalMessageInfo

func (m *CallCredentials) GetFileCredentialSource() *CallCredentials_FileCredentialSource {
	if m != nil {
		return m.FileCredentialSource
	}
	return nil
}

type CallCredentials_FileCredentialSource struct {
	// File containing auth token.
	TokenFileName string `protobuf:"bytes,1,opt,name=token_file_name,json=tokenFileName,proto3" json:"token_file_name,omitempty"`
	// Header to carry the token.
	Header               string   `protobuf:"bytes,2,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallCredentials_FileCredentialSource) Reset()         { *m = CallCredentials_FileCredentialSource{} }
func (m *CallCredentials_FileCredentialSource) String() string { return proto.CompactTextString(m) }
func (*CallCredentials_FileCredentialSource) ProtoMessage()    {}
func (*CallCredentials_FileCredentialSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4a65e8067d81add, []int{4, 0}
}
func (m *CallCredentials_FileCredentialSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallCredentials_FileCredentialSource.Unmarshal(m, b)
}
func (m *CallCredentials_FileCredentialSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallCredentials_FileCredentialSource.Marshal(b, m, deterministic)
}
func (m *CallCredentials_FileCredentialSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallCredentials_FileCredentialSource.Merge(m, src)
}
func (m *CallCredentials_FileCredentialSource) XXX_Size() int {
	return xxx_messageInfo_CallCredentials_FileCredentialSource.Size(m)
}
func (m *CallCredentials_FileCredentialSource) XXX_DiscardUnknown() {
	xxx_messageInfo_CallCredentials_FileCredentialSource.DiscardUnknown(m)
}

var xxx_messageInfo_CallCredentials_FileCredentialSource proto.InternalMessageInfo

func (m *CallCredentials_FileCredentialSource) GetTokenFileName() string {
	if m != nil {
		return m.TokenFileName
	}
	return ""
}

func (m *CallCredentials_FileCredentialSource) GetHeader() string {
	if m != nil {
		return m.Header
	}
	return ""
}

// General TLS parameters. See the [envoy docs](https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/auth/cert.proto#envoy-api-enum-auth-tlsparameters-tlsprotocol)
// for more information on the meaning of these values.
type SslParameters struct {
	MinimumProtocolVersion SslParameters_ProtocolVersion `protobuf:"varint,1,opt,name=minimum_protocol_version,json=minimumProtocolVersion,proto3,enum=gloo.solo.io.SslParameters_ProtocolVersion" json:"minimum_protocol_version,omitempty"`
	MaximumProtocolVersion SslParameters_ProtocolVersion `protobuf:"varint,2,opt,name=maximum_protocol_version,json=maximumProtocolVersion,proto3,enum=gloo.solo.io.SslParameters_ProtocolVersion" json:"maximum_protocol_version,omitempty"`
	CipherSuites           []string                      `protobuf:"bytes,3,rep,name=cipher_suites,json=cipherSuites,proto3" json:"cipher_suites,omitempty"`
	EcdhCurves             []string                      `protobuf:"bytes,4,rep,name=ecdh_curves,json=ecdhCurves,proto3" json:"ecdh_curves,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                      `json:"-"`
	XXX_unrecognized       []byte                        `json:"-"`
	XXX_sizecache          int32                         `json:"-"`
}

func (m *SslParameters) Reset()         { *m = SslParameters{} }
func (m *SslParameters) String() string { return proto.CompactTextString(m) }
func (*SslParameters) ProtoMessage()    {}
func (*SslParameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4a65e8067d81add, []int{5}
}
func (m *SslParameters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SslParameters.Unmarshal(m, b)
}
func (m *SslParameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SslParameters.Marshal(b, m, deterministic)
}
func (m *SslParameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SslParameters.Merge(m, src)
}
func (m *SslParameters) XXX_Size() int {
	return xxx_messageInfo_SslParameters.Size(m)
}
func (m *SslParameters) XXX_DiscardUnknown() {
	xxx_messageInfo_SslParameters.DiscardUnknown(m)
}

var xxx_messageInfo_SslParameters proto.InternalMessageInfo

func (m *SslParameters) GetMinimumProtocolVersion() SslParameters_ProtocolVersion {
	if m != nil {
		return m.MinimumProtocolVersion
	}
	return SslParameters_TLS_AUTO
}

func (m *SslParameters) GetMaximumProtocolVersion() SslParameters_ProtocolVersion {
	if m != nil {
		return m.MaximumProtocolVersion
	}
	return SslParameters_TLS_AUTO
}

func (m *SslParameters) GetCipherSuites() []string {
	if m != nil {
		return m.CipherSuites
	}
	return nil
}

func (m *SslParameters) GetEcdhCurves() []string {
	if m != nil {
		return m.EcdhCurves
	}
	return nil
}

func init() {
	proto.RegisterEnum("gloo.solo.io.SslParameters_ProtocolVersion", SslParameters_ProtocolVersion_name, SslParameters_ProtocolVersion_value)
	proto.RegisterType((*SslConfig)(nil), "gloo.solo.io.SslConfig")
	proto.RegisterType((*SSLFiles)(nil), "gloo.solo.io.SSLFiles")
	proto.RegisterType((*UpstreamSslConfig)(nil), "gloo.solo.io.UpstreamSslConfig")
	proto.RegisterType((*SDSConfig)(nil), "gloo.solo.io.SDSConfig")
	proto.RegisterType((*CallCredentials)(nil), "gloo.solo.io.CallCredentials")
	proto.RegisterType((*CallCredentials_FileCredentialSource)(nil), "gloo.solo.io.CallCredentials.FileCredentialSource")
	proto.RegisterType((*SslParameters)(nil), "gloo.solo.io.SslParameters")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/ssl.proto", fileDescriptor_c4a65e8067d81add)
}

var fileDescriptor_c4a65e8067d81add = []byte{
	// 822 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x55, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0xf6, 0x4f, 0x48, 0xec, 0xe3, 0xb8, 0x31, 0xa3, 0xe0, 0x6c, 0x82, 0x0a, 0x95, 0x2b, 0x50,
	0xa4, 0x8a, 0x75, 0x93, 0xaa, 0x15, 0x2a, 0x57, 0xad, 0x2b, 0x64, 0x41, 0x04, 0xd5, 0x6e, 0x52,
	0x24, 0x6e, 0x46, 0x93, 0xf1, 0xb1, 0x3d, 0x64, 0x76, 0xc7, 0x9a, 0x99, 0xb5, 0x92, 0x77, 0xe1,
	0x01, 0x78, 0x04, 0xde, 0x83, 0x37, 0xe0, 0x05, 0xb8, 0x02, 0x71, 0x89, 0x66, 0x66, 0xb7, 0x4e,
	0xac, 0xa8, 0x8a, 0x80, 0x9b, 0xde, 0xcd, 0xf9, 0xce, 0xf9, 0xe6, 0xcc, 0x7c, 0xe7, 0xdb, 0x59,
	0x78, 0x36, 0x13, 0x76, 0x5e, 0x9c, 0xc7, 0x5c, 0x65, 0x43, 0xa3, 0xa4, 0xfa, 0x42, 0xa8, 0xe1,
	0x4c, 0x2a, 0x35, 0x5c, 0x68, 0xf5, 0x13, 0x72, 0x6b, 0x42, 0xc4, 0x16, 0x62, 0xb8, 0x3c, 0x1a,
	0x1a, 0x23, 0xe3, 0x85, 0x56, 0x56, 0x91, 0x6d, 0x07, 0xc7, 0x8e, 0x11, 0x0b, 0x75, 0xb0, 0x3b,
	0x53, 0x33, 0xe5, 0x13, 0x43, 0xb7, 0x0a, 0x35, 0x07, 0x04, 0x2f, 0x6d, 0x00, 0xf1, 0xd2, 0x96,
	0xd8, 0xbe, 0x6f, 0x72, 0x21, 0x6c, 0xb5, 0xa5, 0xc6, 0x69, 0x48, 0x0d, 0xfe, 0x6a, 0x40, 0x3b,
	0x35, 0x72, 0xa4, 0xf2, 0xa9, 0x98, 0x91, 0xe7, 0x00, 0x06, 0xb9, 0x46, 0x4b, 0x35, 0x4e, 0xa3,
	0xfa, 0x83, 0xfa, 0x61, 0xe7, 0x78, 0x3f, 0xe6, 0x4a, 0x63, 0xd5, 0x35, 0x4e, 0xd0, 0xa8, 0x42,
	0x73, 0x4c, 0x70, 0x3a, 0xae, 0x25, 0xed, 0x50, 0x9e, 0xe0, 0x94, 0x3c, 0x85, 0xb6, 0x31, 0x92,
	0x4e, 0x85, 0x44, 0x13, 0x35, 0x3c, 0xb5, 0x1f, 0x5f, 0x3f, 0x70, 0x9c, 0xa6, 0x27, 0x5f, 0xbb,
	0xec, 0xb8, 0x96, 0xb4, 0x8c, 0x91, 0x7e, 0x4d, 0x1e, 0x41, 0xd3, 0x4c, 0x4c, 0xb4, 0xe1, 0x09,
	0x7b, 0x6b, 0x84, 0x57, 0x69, 0x38, 0xd8, 0xb8, 0x96, 0xb8, 0x2a, 0xf2, 0x29, 0x74, 0x4c, 0x2e,
	0xe8, 0x44, 0x65, 0x4c, 0xe4, 0x26, 0x6a, 0x3e, 0x68, 0x1e, 0xb6, 0x13, 0x30, 0xb9, 0x78, 0x15,
	0x10, 0xf2, 0x14, 0xf6, 0x96, 0xa8, 0xc5, 0xf4, 0x8a, 0x9a, 0xe2, 0xdc, 0x49, 0x49, 0x99, 0xb4,
	0x34, 0x67, 0x19, 0x46, 0x1f, 0xf8, 0xe2, 0xdd, 0x90, 0x4e, 0x43, 0xf6, 0x85, 0xb4, 0xdf, 0xb1,
	0x0c, 0xc9, 0x57, 0x00, 0x0b, 0xa6, 0x59, 0x86, 0x16, 0xb5, 0x89, 0x36, 0xfd, 0x59, 0x3e, 0x5e,
	0x3b, 0x8b, 0x91, 0xaf, 0xdf, 0x96, 0x24, 0xd7, 0xca, 0xc9, 0x67, 0x70, 0x8f, 0xc9, 0x45, 0x4e,
	0xbd, 0xa0, 0x5c, 0x49, 0x13, 0x6d, 0xf9, 0x56, 0x5d, 0x87, 0xbe, 0xae, 0xc0, 0x97, 0x5d, 0xe8,
	0x38, 0x7d, 0x82, 0x60, 0x66, 0xf0, 0x03, 0xb4, 0x2a, 0x3d, 0xc8, 0x3e, 0xb4, 0xac, 0x34, 0x94,
	0xa3, 0xb6, 0x5e, 0xf4, 0x76, 0xb2, 0x65, 0xa5, 0x19, 0xa1, 0xb6, 0x64, 0x0f, 0xdc, 0x92, 0x5e,
	0xe0, 0x95, 0xd7, 0xb4, 0x9d, 0x6c, 0x5a, 0x69, 0xbe, 0xc5, 0x2b, 0x97, 0xd0, 0x4a, 0x59, 0xca,
	0x59, 0xd4, 0x0c, 0x09, 0x17, 0x8e, 0xd8, 0xe0, 0x8f, 0x06, 0x7c, 0x78, 0xb6, 0x30, 0x56, 0x23,
	0xcb, 0xde, 0x9f, 0xc9, 0xf6, 0xa0, 0x69, 0x72, 0x51, 0x5e, 0xc5, 0x2d, 0xff, 0x9f, 0x51, 0x6e,
	0xfd, 0xd7, 0x51, 0xb6, 0xee, 0x30, 0xca, 0x9f, 0xdd, 0x37, 0x54, 0x5d, 0x88, 0xdc, 0x07, 0xb0,
	0x4c, 0xcf, 0xd0, 0xd2, 0x42, 0x8b, 0x72, 0x9c, 0xed, 0x80, 0x9c, 0x69, 0x41, 0xbe, 0x81, 0x1e,
	0x67, 0x52, 0x52, 0xae, 0x71, 0x82, 0xb9, 0x15, 0x4c, 0x56, 0x9a, 0xde, 0xbf, 0x79, 0xca, 0x11,
	0x93, 0x72, 0xb4, 0x2a, 0x1a, 0xd7, 0x92, 0x1d, 0x7e, 0x13, 0x22, 0x0f, 0x61, 0x9b, 0xcb, 0xc2,
	0x58, 0xd4, 0x95, 0x2e, 0xf5, 0xc3, 0xf6, 0xb8, 0x96, 0x74, 0x4a, 0xd4, 0x0b, 0xf2, 0x25, 0x44,
	0xce, 0x58, 0x62, 0x2a, 0x38, 0xb3, 0x68, 0xca, 0x53, 0x07, 0x42, 0x90, 0xbb, 0x7f, 0x3d, 0x9f,
	0xfa, 0xb4, 0x67, 0x3e, 0x83, 0xbd, 0x25, 0x93, 0x62, 0xc2, 0xac, 0x50, 0x39, 0xe5, 0x2a, 0xb7,
	0x78, 0x59, 0x12, 0x37, 0x3c, 0xf1, 0xa3, 0x55, 0x7a, 0x14, 0xb2, 0x8e, 0xe7, 0xe5, 0x99, 0x18,
	0x7a, 0x5e, 0x08, 0x39, 0x41, 0x3d, 0xf8, 0xad, 0x0e, 0x3b, 0x6b, 0x97, 0x21, 0x73, 0xe8, 0x3b,
	0x3b, 0x5d, 0x53, 0x81, 0x06, 0xf3, 0x95, 0xd6, 0x3c, 0x7e, 0xa7, 0x16, 0xb1, 0x33, 0xd8, 0x2a,
	0x4e, 0x83, 0x6d, 0x77, 0xa7, 0xb7, 0xa0, 0x07, 0x6f, 0x60, 0xf7, 0xb6, 0x6a, 0xf2, 0x39, 0xec,
	0x58, 0x75, 0x81, 0xb9, 0xb7, 0x75, 0xb8, 0x54, 0x98, 0x55, 0xd7, 0xc3, 0x8e, 0xe3, 0x45, 0xe8,
	0xc3, 0xe6, 0x1c, 0xd9, 0x04, 0x75, 0xf5, 0xfd, 0x85, 0x68, 0xf0, 0x77, 0x03, 0xba, 0x37, 0x8c,
	0x44, 0x10, 0xa2, 0x4c, 0xe4, 0x22, 0x2b, 0xb2, 0xb7, 0xfe, 0xa1, 0x4b, 0xd4, 0x46, 0xa8, 0xdc,
	0x6f, 0x7d, 0xef, 0xf8, 0xd1, 0x3b, 0x7c, 0x18, 0x57, 0xf6, 0x7a, 0x13, 0x28, 0x49, 0xbf, 0xdc,
	0x6c, 0x0d, 0xf7, 0x6d, 0xd8, 0xe5, 0xed, 0x6d, 0x1a, 0xff, 0xa6, 0x4d, 0xd8, 0x6c, 0xbd, 0xcd,
	0x43, 0xe8, 0x72, 0xb1, 0x98, 0xa3, 0xa6, 0xa6, 0x10, 0x16, 0xab, 0xc7, 0x76, 0x3b, 0x80, 0xa9,
	0xc7, 0xdc, 0x7b, 0x8c, 0x7c, 0x32, 0xa7, 0xbc, 0xd0, 0x4b, 0x74, 0x9f, 0xba, 0x7f, 0x8f, 0x1d,
	0x34, 0xf2, 0xc8, 0x20, 0x85, 0x9d, 0xf5, 0x8d, 0xb7, 0xa1, 0x75, 0x7a, 0x92, 0xd2, 0x17, 0x67,
	0xa7, 0xdf, 0xf7, 0x6a, 0xa4, 0x03, 0x5b, 0xa7, 0x27, 0xe9, 0xf2, 0x88, 0x3e, 0xee, 0xd5, 0x57,
	0xc1, 0x51, 0xaf, 0xb1, 0x0a, 0x8e, 0x7b, 0xcd, 0x55, 0xf0, 0xa4, 0xb7, 0xf1, 0xf2, 0xf9, 0xaf,
	0x7f, 0x6e, 0xd4, 0x7f, 0xf9, 0xfd, 0x93, 0xfa, 0x8f, 0x8f, 0xef, 0xf6, 0x23, 0x5d, 0x5c, 0xcc,
	0xca, 0x3f, 0xdf, 0xf9, 0xa6, 0xd7, 0xec, 0xc9, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x84,
	0x4e, 0x50, 0x83, 0x07, 0x00, 0x00,
}

func (this *SslConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SslConfig)
	if !ok {
		that2, ok := that.(SslConfig)
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
	if that1.SslSecrets == nil {
		if this.SslSecrets != nil {
			return false
		}
	} else if this.SslSecrets == nil {
		return false
	} else if !this.SslSecrets.Equal(that1.SslSecrets) {
		return false
	}
	if len(this.SniDomains) != len(that1.SniDomains) {
		return false
	}
	for i := range this.SniDomains {
		if this.SniDomains[i] != that1.SniDomains[i] {
			return false
		}
	}
	if len(this.VerifySubjectAltName) != len(that1.VerifySubjectAltName) {
		return false
	}
	for i := range this.VerifySubjectAltName {
		if this.VerifySubjectAltName[i] != that1.VerifySubjectAltName[i] {
			return false
		}
	}
	if !this.Parameters.Equal(that1.Parameters) {
		return false
	}
	if len(this.AlpnProtocols) != len(that1.AlpnProtocols) {
		return false
	}
	for i := range this.AlpnProtocols {
		if this.AlpnProtocols[i] != that1.AlpnProtocols[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *SslConfig_SecretRef) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SslConfig_SecretRef)
	if !ok {
		that2, ok := that.(SslConfig_SecretRef)
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
	if !this.SecretRef.Equal(that1.SecretRef) {
		return false
	}
	return true
}
func (this *SslConfig_SslFiles) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SslConfig_SslFiles)
	if !ok {
		that2, ok := that.(SslConfig_SslFiles)
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
	if !this.SslFiles.Equal(that1.SslFiles) {
		return false
	}
	return true
}
func (this *SslConfig_Sds) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SslConfig_Sds)
	if !ok {
		that2, ok := that.(SslConfig_Sds)
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
	if !this.Sds.Equal(that1.Sds) {
		return false
	}
	return true
}
func (this *SSLFiles) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SSLFiles)
	if !ok {
		that2, ok := that.(SSLFiles)
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
	if this.TlsCert != that1.TlsCert {
		return false
	}
	if this.TlsKey != that1.TlsKey {
		return false
	}
	if this.RootCa != that1.RootCa {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *UpstreamSslConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSslConfig)
	if !ok {
		that2, ok := that.(UpstreamSslConfig)
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
	if that1.SslSecrets == nil {
		if this.SslSecrets != nil {
			return false
		}
	} else if this.SslSecrets == nil {
		return false
	} else if !this.SslSecrets.Equal(that1.SslSecrets) {
		return false
	}
	if this.Sni != that1.Sni {
		return false
	}
	if len(this.VerifySubjectAltName) != len(that1.VerifySubjectAltName) {
		return false
	}
	for i := range this.VerifySubjectAltName {
		if this.VerifySubjectAltName[i] != that1.VerifySubjectAltName[i] {
			return false
		}
	}
	if !this.Parameters.Equal(that1.Parameters) {
		return false
	}
	if len(this.AlpnProtocols) != len(that1.AlpnProtocols) {
		return false
	}
	for i := range this.AlpnProtocols {
		if this.AlpnProtocols[i] != that1.AlpnProtocols[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *UpstreamSslConfig_SecretRef) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSslConfig_SecretRef)
	if !ok {
		that2, ok := that.(UpstreamSslConfig_SecretRef)
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
	if !this.SecretRef.Equal(that1.SecretRef) {
		return false
	}
	return true
}
func (this *UpstreamSslConfig_SslFiles) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSslConfig_SslFiles)
	if !ok {
		that2, ok := that.(UpstreamSslConfig_SslFiles)
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
	if !this.SslFiles.Equal(that1.SslFiles) {
		return false
	}
	return true
}
func (this *UpstreamSslConfig_Sds) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSslConfig_Sds)
	if !ok {
		that2, ok := that.(UpstreamSslConfig_Sds)
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
	if !this.Sds.Equal(that1.Sds) {
		return false
	}
	return true
}
func (this *SDSConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SDSConfig)
	if !ok {
		that2, ok := that.(SDSConfig)
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
	if this.TargetUri != that1.TargetUri {
		return false
	}
	if that1.SdsBuilder == nil {
		if this.SdsBuilder != nil {
			return false
		}
	} else if this.SdsBuilder == nil {
		return false
	} else if !this.SdsBuilder.Equal(that1.SdsBuilder) {
		return false
	}
	if this.CertificatesSecretName != that1.CertificatesSecretName {
		return false
	}
	if this.ValidationContextName != that1.ValidationContextName {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *SDSConfig_CallCredentials) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SDSConfig_CallCredentials)
	if !ok {
		that2, ok := that.(SDSConfig_CallCredentials)
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
	if !this.CallCredentials.Equal(that1.CallCredentials) {
		return false
	}
	return true
}
func (this *SDSConfig_ClusterName) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SDSConfig_ClusterName)
	if !ok {
		that2, ok := that.(SDSConfig_ClusterName)
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
	if this.ClusterName != that1.ClusterName {
		return false
	}
	return true
}
func (this *CallCredentials) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CallCredentials)
	if !ok {
		that2, ok := that.(CallCredentials)
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
	if !this.FileCredentialSource.Equal(that1.FileCredentialSource) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *CallCredentials_FileCredentialSource) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CallCredentials_FileCredentialSource)
	if !ok {
		that2, ok := that.(CallCredentials_FileCredentialSource)
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
	if this.TokenFileName != that1.TokenFileName {
		return false
	}
	if this.Header != that1.Header {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *SslParameters) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SslParameters)
	if !ok {
		that2, ok := that.(SslParameters)
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
	if this.MinimumProtocolVersion != that1.MinimumProtocolVersion {
		return false
	}
	if this.MaximumProtocolVersion != that1.MaximumProtocolVersion {
		return false
	}
	if len(this.CipherSuites) != len(that1.CipherSuites) {
		return false
	}
	for i := range this.CipherSuites {
		if this.CipherSuites[i] != that1.CipherSuites[i] {
			return false
		}
	}
	if len(this.EcdhCurves) != len(that1.EcdhCurves) {
		return false
	}
	for i := range this.EcdhCurves {
		if this.EcdhCurves[i] != that1.EcdhCurves[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
