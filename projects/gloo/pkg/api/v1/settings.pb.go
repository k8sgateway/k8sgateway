// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/settings.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
//@solo-kit:resource.short_name=st
//@solo-kit:resource.plural_name=settings
type Settings struct {
	// namespace to write discovered data
	DiscoveryNamespace string `protobuf:"bytes,1,opt,name=discovery_namespace,json=discoveryNamespace,proto3" json:"discovery_namespace,omitempty"`
	// namespaces to watch for user config as well as services
	// TODO(ilackarms): split out watch_namespaces and service_discovery_namespaces...
	WatchNamespaces []string `protobuf:"bytes,2,rep,name=watch_namespaces,json=watchNamespaces,proto3" json:"watch_namespaces,omitempty"`
	// where to read user config (upstream, proxy) from
	// if nil, use only in memory config
	//
	// Types that are valid to be assigned to ConfigSource:
	//	*Settings_KubernetesConfigSource
	//	*Settings_DirectoryConfigSource
	ConfigSource isSettings_ConfigSource `protobuf_oneof:"config_source"`
	// where to read secrets from (vault, k8s)
	//
	// Types that are valid to be assigned to SecretSource:
	//	*Settings_KubernetesSecretSource
	//	*Settings_VaultSecretSource
	//	*Settings_DirectorySecretSource
	SecretSource isSettings_SecretSource `protobuf_oneof:"secret_source"`
	// where to read artifacts from (configmap, file)
	//
	// Types that are valid to be assigned to ArtifactSource:
	//	*Settings_KubernetesArtifactSource
	//	*Settings_DirectoryArtifactSource
	ArtifactSource isSettings_ArtifactSource `protobuf_oneof:"artifact_source"`
	// where the gloo xds server should bind (should not need configuration by user)
	BindAddr string `protobuf:"bytes,11,opt,name=bind_addr,json=bindAddr,proto3" json:"bind_addr,omitempty"`
	// how frequently to resync watches, etc
	RefreshRate *types.Duration `protobuf:"bytes,12,opt,name=refresh_rate,json=refreshRate,proto3" json:"refresh_rate,omitempty"`
	// enable serving debug data on port 9090
	DevMode bool `protobuf:"varint,13,opt,name=dev_mode,json=devMode,proto3" json:"dev_mode,omitempty"`
	// enable automatic linkerd upstream header addition for easier routing to linkerd services
	Linkerd bool `protobuf:"varint,17,opt,name=linkerd,proto3" json:"linkerd,omitempty"`
	// Default circuit breakers when not set in a specific upstream.
	CircuitBreakers *CircuitBreakerConfig `protobuf:"bytes,3,opt,name=circuit_breakers,json=circuitBreakers,proto3" json:"circuit_breakers,omitempty"`
	// configuration options for the Clusteringress Controller (for Knative)
	Knative *Settings_KnativeOptions `protobuf:"bytes,18,opt,name=knative,proto3" json:"knative,omitempty"`
	// Settings for extensions
	Extensions *Extensions `protobuf:"bytes,16,opt,name=extensions,proto3" json:"extensions,omitempty"`
	// Metadata contains the object metadata for this resource
	Metadata core.Metadata `protobuf:"bytes,14,opt,name=metadata,proto3" json:"metadata"`
	// Status indicates the validation status of this resource.
	// Status is read-only by clients, and set by gloo during validation
	Status               core.Status `protobuf:"bytes,15,opt,name=status,proto3" json:"status" testdiff:"ignore"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Settings) Reset()         { *m = Settings{} }
func (m *Settings) String() string { return proto.CompactTextString(m) }
func (*Settings) ProtoMessage()    {}
func (*Settings) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd7533c2495e1752, []int{0}
}
func (m *Settings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Settings.Unmarshal(m, b)
}
func (m *Settings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Settings.Marshal(b, m, deterministic)
}
func (m *Settings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Settings.Merge(m, src)
}
func (m *Settings) XXX_Size() int {
	return xxx_messageInfo_Settings.Size(m)
}
func (m *Settings) XXX_DiscardUnknown() {
	xxx_messageInfo_Settings.DiscardUnknown(m)
}

var xxx_messageInfo_Settings proto.InternalMessageInfo

type isSettings_ConfigSource interface {
	isSettings_ConfigSource()
	Equal(interface{}) bool
}
type isSettings_SecretSource interface {
	isSettings_SecretSource()
	Equal(interface{}) bool
}
type isSettings_ArtifactSource interface {
	isSettings_ArtifactSource()
	Equal(interface{}) bool
}

type Settings_KubernetesConfigSource struct {
	KubernetesConfigSource *Settings_KubernetesCrds `protobuf:"bytes,4,opt,name=kubernetes_config_source,json=kubernetesConfigSource,proto3,oneof"`
}
type Settings_DirectoryConfigSource struct {
	DirectoryConfigSource *Settings_Directory `protobuf:"bytes,5,opt,name=directory_config_source,json=directoryConfigSource,proto3,oneof"`
}
type Settings_KubernetesSecretSource struct {
	KubernetesSecretSource *Settings_KubernetesSecrets `protobuf:"bytes,6,opt,name=kubernetes_secret_source,json=kubernetesSecretSource,proto3,oneof"`
}
type Settings_VaultSecretSource struct {
	VaultSecretSource *Settings_VaultSecrets `protobuf:"bytes,7,opt,name=vault_secret_source,json=vaultSecretSource,proto3,oneof"`
}
type Settings_DirectorySecretSource struct {
	DirectorySecretSource *Settings_Directory `protobuf:"bytes,8,opt,name=directory_secret_source,json=directorySecretSource,proto3,oneof"`
}
type Settings_KubernetesArtifactSource struct {
	KubernetesArtifactSource *Settings_KubernetesConfigmaps `protobuf:"bytes,9,opt,name=kubernetes_artifact_source,json=kubernetesArtifactSource,proto3,oneof"`
}
type Settings_DirectoryArtifactSource struct {
	DirectoryArtifactSource *Settings_Directory `protobuf:"bytes,10,opt,name=directory_artifact_source,json=directoryArtifactSource,proto3,oneof"`
}

func (*Settings_KubernetesConfigSource) isSettings_ConfigSource()     {}
func (*Settings_DirectoryConfigSource) isSettings_ConfigSource()      {}
func (*Settings_KubernetesSecretSource) isSettings_SecretSource()     {}
func (*Settings_VaultSecretSource) isSettings_SecretSource()          {}
func (*Settings_DirectorySecretSource) isSettings_SecretSource()      {}
func (*Settings_KubernetesArtifactSource) isSettings_ArtifactSource() {}
func (*Settings_DirectoryArtifactSource) isSettings_ArtifactSource()  {}

func (m *Settings) GetConfigSource() isSettings_ConfigSource {
	if m != nil {
		return m.ConfigSource
	}
	return nil
}
func (m *Settings) GetSecretSource() isSettings_SecretSource {
	if m != nil {
		return m.SecretSource
	}
	return nil
}
func (m *Settings) GetArtifactSource() isSettings_ArtifactSource {
	if m != nil {
		return m.ArtifactSource
	}
	return nil
}

func (m *Settings) GetDiscoveryNamespace() string {
	if m != nil {
		return m.DiscoveryNamespace
	}
	return ""
}

func (m *Settings) GetWatchNamespaces() []string {
	if m != nil {
		return m.WatchNamespaces
	}
	return nil
}

func (m *Settings) GetKubernetesConfigSource() *Settings_KubernetesCrds {
	if x, ok := m.GetConfigSource().(*Settings_KubernetesConfigSource); ok {
		return x.KubernetesConfigSource
	}
	return nil
}

func (m *Settings) GetDirectoryConfigSource() *Settings_Directory {
	if x, ok := m.GetConfigSource().(*Settings_DirectoryConfigSource); ok {
		return x.DirectoryConfigSource
	}
	return nil
}

func (m *Settings) GetKubernetesSecretSource() *Settings_KubernetesSecrets {
	if x, ok := m.GetSecretSource().(*Settings_KubernetesSecretSource); ok {
		return x.KubernetesSecretSource
	}
	return nil
}

func (m *Settings) GetVaultSecretSource() *Settings_VaultSecrets {
	if x, ok := m.GetSecretSource().(*Settings_VaultSecretSource); ok {
		return x.VaultSecretSource
	}
	return nil
}

func (m *Settings) GetDirectorySecretSource() *Settings_Directory {
	if x, ok := m.GetSecretSource().(*Settings_DirectorySecretSource); ok {
		return x.DirectorySecretSource
	}
	return nil
}

func (m *Settings) GetKubernetesArtifactSource() *Settings_KubernetesConfigmaps {
	if x, ok := m.GetArtifactSource().(*Settings_KubernetesArtifactSource); ok {
		return x.KubernetesArtifactSource
	}
	return nil
}

func (m *Settings) GetDirectoryArtifactSource() *Settings_Directory {
	if x, ok := m.GetArtifactSource().(*Settings_DirectoryArtifactSource); ok {
		return x.DirectoryArtifactSource
	}
	return nil
}

func (m *Settings) GetBindAddr() string {
	if m != nil {
		return m.BindAddr
	}
	return ""
}

func (m *Settings) GetRefreshRate() *types.Duration {
	if m != nil {
		return m.RefreshRate
	}
	return nil
}

func (m *Settings) GetDevMode() bool {
	if m != nil {
		return m.DevMode
	}
	return false
}

func (m *Settings) GetLinkerd() bool {
	if m != nil {
		return m.Linkerd
	}
	return false
}

func (m *Settings) GetCircuitBreakers() *CircuitBreakerConfig {
	if m != nil {
		return m.CircuitBreakers
	}
	return nil
}

func (m *Settings) GetKnative() *Settings_KnativeOptions {
	if m != nil {
		return m.Knative
	}
	return nil
}

func (m *Settings) GetExtensions() *Extensions {
	if m != nil {
		return m.Extensions
	}
	return nil
}

func (m *Settings) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *Settings) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Settings) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Settings_OneofMarshaler, _Settings_OneofUnmarshaler, _Settings_OneofSizer, []interface{}{
		(*Settings_KubernetesConfigSource)(nil),
		(*Settings_DirectoryConfigSource)(nil),
		(*Settings_KubernetesSecretSource)(nil),
		(*Settings_VaultSecretSource)(nil),
		(*Settings_DirectorySecretSource)(nil),
		(*Settings_KubernetesArtifactSource)(nil),
		(*Settings_DirectoryArtifactSource)(nil),
	}
}

func _Settings_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Settings)
	// config_source
	switch x := m.ConfigSource.(type) {
	case *Settings_KubernetesConfigSource:
		_ = b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.KubernetesConfigSource); err != nil {
			return err
		}
	case *Settings_DirectoryConfigSource:
		_ = b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DirectoryConfigSource); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Settings.ConfigSource has unexpected type %T", x)
	}
	// secret_source
	switch x := m.SecretSource.(type) {
	case *Settings_KubernetesSecretSource:
		_ = b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.KubernetesSecretSource); err != nil {
			return err
		}
	case *Settings_VaultSecretSource:
		_ = b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.VaultSecretSource); err != nil {
			return err
		}
	case *Settings_DirectorySecretSource:
		_ = b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DirectorySecretSource); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Settings.SecretSource has unexpected type %T", x)
	}
	// artifact_source
	switch x := m.ArtifactSource.(type) {
	case *Settings_KubernetesArtifactSource:
		_ = b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.KubernetesArtifactSource); err != nil {
			return err
		}
	case *Settings_DirectoryArtifactSource:
		_ = b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DirectoryArtifactSource); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Settings.ArtifactSource has unexpected type %T", x)
	}
	return nil
}

func _Settings_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Settings)
	switch tag {
	case 4: // config_source.kubernetes_config_source
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Settings_KubernetesCrds)
		err := b.DecodeMessage(msg)
		m.ConfigSource = &Settings_KubernetesConfigSource{msg}
		return true, err
	case 5: // config_source.directory_config_source
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Settings_Directory)
		err := b.DecodeMessage(msg)
		m.ConfigSource = &Settings_DirectoryConfigSource{msg}
		return true, err
	case 6: // secret_source.kubernetes_secret_source
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Settings_KubernetesSecrets)
		err := b.DecodeMessage(msg)
		m.SecretSource = &Settings_KubernetesSecretSource{msg}
		return true, err
	case 7: // secret_source.vault_secret_source
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Settings_VaultSecrets)
		err := b.DecodeMessage(msg)
		m.SecretSource = &Settings_VaultSecretSource{msg}
		return true, err
	case 8: // secret_source.directory_secret_source
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Settings_Directory)
		err := b.DecodeMessage(msg)
		m.SecretSource = &Settings_DirectorySecretSource{msg}
		return true, err
	case 9: // artifact_source.kubernetes_artifact_source
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Settings_KubernetesConfigmaps)
		err := b.DecodeMessage(msg)
		m.ArtifactSource = &Settings_KubernetesArtifactSource{msg}
		return true, err
	case 10: // artifact_source.directory_artifact_source
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Settings_Directory)
		err := b.DecodeMessage(msg)
		m.ArtifactSource = &Settings_DirectoryArtifactSource{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Settings_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Settings)
	// config_source
	switch x := m.ConfigSource.(type) {
	case *Settings_KubernetesConfigSource:
		s := proto.Size(x.KubernetesConfigSource)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Settings_DirectoryConfigSource:
		s := proto.Size(x.DirectoryConfigSource)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	// secret_source
	switch x := m.SecretSource.(type) {
	case *Settings_KubernetesSecretSource:
		s := proto.Size(x.KubernetesSecretSource)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Settings_VaultSecretSource:
		s := proto.Size(x.VaultSecretSource)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Settings_DirectorySecretSource:
		s := proto.Size(x.DirectorySecretSource)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	// artifact_source
	switch x := m.ArtifactSource.(type) {
	case *Settings_KubernetesArtifactSource:
		s := proto.Size(x.KubernetesArtifactSource)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Settings_DirectoryArtifactSource:
		s := proto.Size(x.DirectoryArtifactSource)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// ilackarms(todo: make sure these are configurable)
type Settings_KubernetesCrds struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Settings_KubernetesCrds) Reset()         { *m = Settings_KubernetesCrds{} }
func (m *Settings_KubernetesCrds) String() string { return proto.CompactTextString(m) }
func (*Settings_KubernetesCrds) ProtoMessage()    {}
func (*Settings_KubernetesCrds) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd7533c2495e1752, []int{0, 0}
}
func (m *Settings_KubernetesCrds) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Settings_KubernetesCrds.Unmarshal(m, b)
}
func (m *Settings_KubernetesCrds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Settings_KubernetesCrds.Marshal(b, m, deterministic)
}
func (m *Settings_KubernetesCrds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Settings_KubernetesCrds.Merge(m, src)
}
func (m *Settings_KubernetesCrds) XXX_Size() int {
	return xxx_messageInfo_Settings_KubernetesCrds.Size(m)
}
func (m *Settings_KubernetesCrds) XXX_DiscardUnknown() {
	xxx_messageInfo_Settings_KubernetesCrds.DiscardUnknown(m)
}

var xxx_messageInfo_Settings_KubernetesCrds proto.InternalMessageInfo

type Settings_KubernetesSecrets struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Settings_KubernetesSecrets) Reset()         { *m = Settings_KubernetesSecrets{} }
func (m *Settings_KubernetesSecrets) String() string { return proto.CompactTextString(m) }
func (*Settings_KubernetesSecrets) ProtoMessage()    {}
func (*Settings_KubernetesSecrets) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd7533c2495e1752, []int{0, 1}
}
func (m *Settings_KubernetesSecrets) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Settings_KubernetesSecrets.Unmarshal(m, b)
}
func (m *Settings_KubernetesSecrets) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Settings_KubernetesSecrets.Marshal(b, m, deterministic)
}
func (m *Settings_KubernetesSecrets) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Settings_KubernetesSecrets.Merge(m, src)
}
func (m *Settings_KubernetesSecrets) XXX_Size() int {
	return xxx_messageInfo_Settings_KubernetesSecrets.Size(m)
}
func (m *Settings_KubernetesSecrets) XXX_DiscardUnknown() {
	xxx_messageInfo_Settings_KubernetesSecrets.DiscardUnknown(m)
}

var xxx_messageInfo_Settings_KubernetesSecrets proto.InternalMessageInfo

type Settings_VaultSecrets struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Settings_VaultSecrets) Reset()         { *m = Settings_VaultSecrets{} }
func (m *Settings_VaultSecrets) String() string { return proto.CompactTextString(m) }
func (*Settings_VaultSecrets) ProtoMessage()    {}
func (*Settings_VaultSecrets) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd7533c2495e1752, []int{0, 2}
}
func (m *Settings_VaultSecrets) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Settings_VaultSecrets.Unmarshal(m, b)
}
func (m *Settings_VaultSecrets) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Settings_VaultSecrets.Marshal(b, m, deterministic)
}
func (m *Settings_VaultSecrets) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Settings_VaultSecrets.Merge(m, src)
}
func (m *Settings_VaultSecrets) XXX_Size() int {
	return xxx_messageInfo_Settings_VaultSecrets.Size(m)
}
func (m *Settings_VaultSecrets) XXX_DiscardUnknown() {
	xxx_messageInfo_Settings_VaultSecrets.DiscardUnknown(m)
}

var xxx_messageInfo_Settings_VaultSecrets proto.InternalMessageInfo

type Settings_KubernetesConfigmaps struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Settings_KubernetesConfigmaps) Reset()         { *m = Settings_KubernetesConfigmaps{} }
func (m *Settings_KubernetesConfigmaps) String() string { return proto.CompactTextString(m) }
func (*Settings_KubernetesConfigmaps) ProtoMessage()    {}
func (*Settings_KubernetesConfigmaps) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd7533c2495e1752, []int{0, 3}
}
func (m *Settings_KubernetesConfigmaps) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Settings_KubernetesConfigmaps.Unmarshal(m, b)
}
func (m *Settings_KubernetesConfigmaps) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Settings_KubernetesConfigmaps.Marshal(b, m, deterministic)
}
func (m *Settings_KubernetesConfigmaps) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Settings_KubernetesConfigmaps.Merge(m, src)
}
func (m *Settings_KubernetesConfigmaps) XXX_Size() int {
	return xxx_messageInfo_Settings_KubernetesConfigmaps.Size(m)
}
func (m *Settings_KubernetesConfigmaps) XXX_DiscardUnknown() {
	xxx_messageInfo_Settings_KubernetesConfigmaps.DiscardUnknown(m)
}

var xxx_messageInfo_Settings_KubernetesConfigmaps proto.InternalMessageInfo

type Settings_Directory struct {
	Directory            string   `protobuf:"bytes,1,opt,name=directory,proto3" json:"directory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Settings_Directory) Reset()         { *m = Settings_Directory{} }
func (m *Settings_Directory) String() string { return proto.CompactTextString(m) }
func (*Settings_Directory) ProtoMessage()    {}
func (*Settings_Directory) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd7533c2495e1752, []int{0, 4}
}
func (m *Settings_Directory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Settings_Directory.Unmarshal(m, b)
}
func (m *Settings_Directory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Settings_Directory.Marshal(b, m, deterministic)
}
func (m *Settings_Directory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Settings_Directory.Merge(m, src)
}
func (m *Settings_Directory) XXX_Size() int {
	return xxx_messageInfo_Settings_Directory.Size(m)
}
func (m *Settings_Directory) XXX_DiscardUnknown() {
	xxx_messageInfo_Settings_Directory.DiscardUnknown(m)
}

var xxx_messageInfo_Settings_Directory proto.InternalMessageInfo

func (m *Settings_Directory) GetDirectory() string {
	if m != nil {
		return m.Directory
	}
	return ""
}

type Settings_KnativeOptions struct {
	// address of the clusteringress proxy
	// if empty, it will default to clusteringress-proxy.$POD_NAMESPACE.svc.cluster.local
	ClusterIngressProxyAddress string   `protobuf:"bytes,1,opt,name=cluster_ingress_proxy_address,json=clusterIngressProxyAddress,proto3" json:"cluster_ingress_proxy_address,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *Settings_KnativeOptions) Reset()         { *m = Settings_KnativeOptions{} }
func (m *Settings_KnativeOptions) String() string { return proto.CompactTextString(m) }
func (*Settings_KnativeOptions) ProtoMessage()    {}
func (*Settings_KnativeOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd7533c2495e1752, []int{0, 5}
}
func (m *Settings_KnativeOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Settings_KnativeOptions.Unmarshal(m, b)
}
func (m *Settings_KnativeOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Settings_KnativeOptions.Marshal(b, m, deterministic)
}
func (m *Settings_KnativeOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Settings_KnativeOptions.Merge(m, src)
}
func (m *Settings_KnativeOptions) XXX_Size() int {
	return xxx_messageInfo_Settings_KnativeOptions.Size(m)
}
func (m *Settings_KnativeOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_Settings_KnativeOptions.DiscardUnknown(m)
}

var xxx_messageInfo_Settings_KnativeOptions proto.InternalMessageInfo

func (m *Settings_KnativeOptions) GetClusterIngressProxyAddress() string {
	if m != nil {
		return m.ClusterIngressProxyAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*Settings)(nil), "gloo.solo.io.Settings")
	proto.RegisterType((*Settings_KubernetesCrds)(nil), "gloo.solo.io.Settings.KubernetesCrds")
	proto.RegisterType((*Settings_KubernetesSecrets)(nil), "gloo.solo.io.Settings.KubernetesSecrets")
	proto.RegisterType((*Settings_VaultSecrets)(nil), "gloo.solo.io.Settings.VaultSecrets")
	proto.RegisterType((*Settings_KubernetesConfigmaps)(nil), "gloo.solo.io.Settings.KubernetesConfigmaps")
	proto.RegisterType((*Settings_Directory)(nil), "gloo.solo.io.Settings.Directory")
	proto.RegisterType((*Settings_KnativeOptions)(nil), "gloo.solo.io.Settings.KnativeOptions")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/settings.proto", fileDescriptor_bd7533c2495e1752)
}

var fileDescriptor_bd7533c2495e1752 = []byte{
	// 797 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x6e, 0xba, 0x4b, 0x93, 0x4c, 0xb2, 0xf9, 0x99, 0x96, 0xe2, 0x18, 0xd8, 0x8d, 0x82, 0x90,
	0xb2, 0x42, 0xd8, 0x2c, 0x48, 0x68, 0xc5, 0x8f, 0x50, 0xd2, 0x45, 0x04, 0xa1, 0x02, 0x72, 0x04,
	0x17, 0x7b, 0x81, 0x35, 0xf1, 0x9c, 0xb8, 0x43, 0x12, 0x8f, 0x35, 0x33, 0x0e, 0xdb, 0x37, 0xe2,
	0x51, 0x78, 0x8a, 0xbd, 0xe0, 0x9a, 0x2b, 0x9e, 0x00, 0x79, 0x3c, 0xb6, 0xe3, 0xd0, 0x6e, 0xd3,
	0xab, 0x76, 0xe6, 0x7c, 0x3f, 0x73, 0x4e, 0xce, 0x39, 0x46, 0x5f, 0x86, 0x4c, 0x5d, 0x25, 0x0b,
	0x27, 0xe0, 0x1b, 0x57, 0xf2, 0x35, 0xff, 0x98, 0x71, 0x37, 0x5c, 0x73, 0xee, 0xc6, 0x82, 0xff,
	0x0e, 0x81, 0x92, 0xd9, 0x89, 0xc4, 0xcc, 0xdd, 0x3e, 0x73, 0x25, 0x28, 0xc5, 0xa2, 0x50, 0x3a,
	0xb1, 0xe0, 0x8a, 0xe3, 0x76, 0x1a, 0x73, 0x52, 0x9a, 0xc3, 0xb8, 0x7d, 0x16, 0xf2, 0x90, 0xeb,
	0x80, 0x9b, 0xfe, 0x97, 0x61, 0xec, 0x67, 0x37, 0x18, 0xe8, 0xbf, 0x2b, 0xa6, 0x72, 0xd9, 0x0d,
	0x28, 0x42, 0x89, 0x22, 0x86, 0xe2, 0x1e, 0x40, 0x91, 0x8a, 0xa8, 0xc4, 0xbc, 0xc3, 0xfe, 0xfa,
	0x5e, 0x49, 0xc0, 0x2b, 0x05, 0x91, 0x64, 0x3c, 0xca, 0xe9, 0xd3, 0x7b, 0xd1, 0x03, 0x26, 0x82,
	0x84, 0x29, 0x7f, 0x21, 0x80, 0xac, 0x40, 0x18, 0x8d, 0xc7, 0x21, 0xe7, 0xe1, 0x1a, 0x5c, 0x7d,
	0x5a, 0x24, 0x4b, 0x97, 0x26, 0x82, 0x28, 0xc6, 0xa3, 0x2c, 0x3e, 0xfa, 0xa7, 0x85, 0x1a, 0x73,
	0x53, 0x3d, 0xec, 0xa2, 0x53, 0xca, 0x64, 0xc0, 0xb7, 0x20, 0xae, 0xfd, 0x88, 0x6c, 0x40, 0xc6,
	0x24, 0x00, 0xab, 0x36, 0xac, 0x8d, 0x9b, 0x1e, 0x2e, 0x42, 0x3f, 0xe6, 0x11, 0xfc, 0x14, 0xf5,
	0xfe, 0x20, 0x2a, 0xb8, 0x2a, 0xc1, 0xd2, 0x3a, 0x1e, 0x3e, 0x18, 0x37, 0xbd, 0xae, 0xbe, 0x2f,
	0x90, 0x12, 0x13, 0x64, 0xad, 0x92, 0x05, 0x88, 0x08, 0x14, 0x48, 0x3f, 0xe0, 0xd1, 0x92, 0x85,
	0xbe, 0xe4, 0x89, 0x08, 0xc0, 0x7a, 0x38, 0xac, 0x8d, 0x5b, 0x9f, 0x7e, 0xe8, 0xec, 0xfe, 0x6c,
	0x4e, 0xfe, 0x2a, 0xe7, 0x87, 0x82, 0x76, 0x21, 0xa8, 0x9c, 0x1d, 0x79, 0xe7, 0xa5, 0xd0, 0x85,
	0xd6, 0x99, 0x6b, 0x19, 0xfc, 0x12, 0xbd, 0x43, 0x99, 0x80, 0x40, 0x71, 0x71, 0xbd, 0xe7, 0xf0,
	0x96, 0x76, 0x18, 0xde, 0xe2, 0xf0, 0x22, 0x67, 0xcd, 0x8e, 0xbc, 0xb7, 0x0b, 0x89, 0x8a, 0x36,
	0xad, 0x3c, 0x5f, 0x42, 0x20, 0x40, 0xe5, 0xe2, 0x27, 0x5a, 0x7c, 0x7c, 0xe7, 0xf3, 0xe7, 0x9a,
	0x25, 0x67, 0xb5, 0xdd, 0x0c, 0xb2, 0x4b, 0xe3, 0xf2, 0x0b, 0x3a, 0xdd, 0x92, 0x64, 0xad, 0xf6,
	0x0c, 0xea, 0xda, 0xe0, 0x83, 0x5b, 0x0c, 0x7e, 0x4d, 0x19, 0xa5, 0x76, 0x7f, 0x5b, 0x9e, 0x6f,
	0x2a, 0x4c, 0x55, 0xba, 0x71, 0x60, 0x61, 0x6a, 0x3b, 0x85, 0xa9, 0x68, 0xaf, 0x90, 0xbd, 0x53,
	0x18, 0x22, 0x14, 0x5b, 0x92, 0xa0, 0x90, 0x6f, 0x6a, 0xf9, 0x8f, 0xee, 0xfe, 0x65, 0x75, 0xad,
	0x37, 0x24, 0x96, 0xb3, 0x63, 0x6f, 0xa7, 0xd2, 0x13, 0xa3, 0x67, 0xcc, 0x7e, 0x43, 0x83, 0x32,
	0x91, 0x7d, 0x2f, 0x74, 0x60, 0x2a, 0xc7, 0x5e, 0x59, 0x8d, 0x3d, 0xfd, 0x77, 0x51, 0x73, 0xc1,
	0x22, 0xea, 0x13, 0x4a, 0x85, 0xd5, 0xd2, 0x6d, 0xdf, 0x48, 0x2f, 0x26, 0x94, 0x0a, 0xfc, 0x15,
	0x6a, 0x0b, 0x58, 0x0a, 0x90, 0x57, 0xbe, 0x20, 0x0a, 0xac, 0xb6, 0xf6, 0x1b, 0x38, 0xd9, 0x84,
	0x39, 0xf9, 0x84, 0x39, 0x2f, 0xcc, 0x84, 0x79, 0x2d, 0x03, 0xf7, 0x88, 0x02, 0x3c, 0x40, 0x0d,
	0x0a, 0x5b, 0x7f, 0xc3, 0x29, 0x58, 0x8f, 0x86, 0xb5, 0x71, 0xc3, 0xab, 0x53, 0xd8, 0x5e, 0x72,
	0x0a, 0xd8, 0x42, 0xf5, 0x35, 0x8b, 0x56, 0x20, 0xa8, 0xd5, 0xcf, 0x22, 0xe6, 0x88, 0x2f, 0x51,
	0x6f, 0x6f, 0xac, 0xa5, 0xf5, 0x40, 0xdb, 0x8e, 0xaa, 0x69, 0x5e, 0x64, 0xa8, 0x69, 0x06, 0xca,
	0xaa, 0xe9, 0x75, 0x83, 0xca, 0xad, 0xc4, 0xdf, 0xa0, 0xfa, 0x2a, 0x22, 0x8a, 0x6d, 0xc1, 0xc2,
	0x6f, 0x1e, 0xb9, 0x0c, 0xf5, 0x53, 0x9c, 0xe6, 0x21, 0xbd, 0x9c, 0x85, 0x9f, 0x23, 0x54, 0x6e,
	0x29, 0xab, 0xa7, 0x35, 0xac, 0xaa, 0xc6, 0xb7, 0x45, 0xdc, 0xdb, 0xc1, 0xe2, 0xe7, 0xa8, 0x91,
	0x6f, 0x53, 0xab, 0xa3, 0x79, 0xe7, 0x4e, 0xc0, 0x05, 0x14, 0xbc, 0x4b, 0x13, 0x9d, 0x3e, 0xfc,
	0xeb, 0xf5, 0x93, 0x23, 0xaf, 0x40, 0xe3, 0xef, 0xd0, 0x49, 0xb6, 0x54, 0xad, 0xae, 0xe6, 0x9d,
	0x55, 0x79, 0x73, 0x1d, 0x9b, 0x0e, 0x52, 0xd6, 0xbf, 0xaf, 0x9f, 0xf4, 0x15, 0x48, 0x45, 0xd9,
	0x72, 0xf9, 0xc5, 0x88, 0x85, 0x11, 0x17, 0x30, 0xf2, 0x0c, 0xdd, 0xee, 0xa1, 0x4e, 0x75, 0x95,
	0xd8, 0xa7, 0xa8, 0xff, 0xbf, 0xe9, 0xb4, 0x3b, 0xa8, 0xbd, 0x3b, 0x51, 0xf6, 0x39, 0x3a, 0xbb,
	0xa9, 0x4f, 0xed, 0xa7, 0xa8, 0x59, 0xf4, 0x14, 0x7e, 0x0f, 0x35, 0x8b, 0x9e, 0x32, 0xfb, 0xb2,
	0xbc, 0xb0, 0xe7, 0xa8, 0x53, 0xad, 0x28, 0x9e, 0xa0, 0xf7, 0x83, 0x75, 0x22, 0x15, 0x08, 0x9f,
	0x45, 0xa1, 0x00, 0x29, 0xfd, 0x58, 0xf0, 0x57, 0xd7, 0xba, 0xf3, 0x40, 0x4a, 0xa3, 0x61, 0x1b,
	0xd0, 0xf7, 0x19, 0xe6, 0xe7, 0x14, 0x32, 0xc9, 0x10, 0xd3, 0x2e, 0x7a, 0x54, 0xd9, 0x71, 0xe9,
	0x45, 0x65, 0xb6, 0xa7, 0x7d, 0xd4, 0xdd, 0x9b, 0x91, 0xe9, 0xe7, 0x7f, 0xfe, 0xfd, 0xb8, 0xf6,
	0xf2, 0x93, 0xc3, 0x3e, 0x2c, 0xf1, 0x2a, 0x34, 0x1f, 0x97, 0xc5, 0x89, 0xee, 0xee, 0xcf, 0xfe,
	0x0b, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x84, 0x4f, 0xc7, 0x97, 0x07, 0x00, 0x00,
}

func (this *Settings) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Settings)
	if !ok {
		that2, ok := that.(Settings)
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
	if this.DiscoveryNamespace != that1.DiscoveryNamespace {
		return false
	}
	if len(this.WatchNamespaces) != len(that1.WatchNamespaces) {
		return false
	}
	for i := range this.WatchNamespaces {
		if this.WatchNamespaces[i] != that1.WatchNamespaces[i] {
			return false
		}
	}
	if that1.ConfigSource == nil {
		if this.ConfigSource != nil {
			return false
		}
	} else if this.ConfigSource == nil {
		return false
	} else if !this.ConfigSource.Equal(that1.ConfigSource) {
		return false
	}
	if that1.SecretSource == nil {
		if this.SecretSource != nil {
			return false
		}
	} else if this.SecretSource == nil {
		return false
	} else if !this.SecretSource.Equal(that1.SecretSource) {
		return false
	}
	if that1.ArtifactSource == nil {
		if this.ArtifactSource != nil {
			return false
		}
	} else if this.ArtifactSource == nil {
		return false
	} else if !this.ArtifactSource.Equal(that1.ArtifactSource) {
		return false
	}
	if this.BindAddr != that1.BindAddr {
		return false
	}
	if !this.RefreshRate.Equal(that1.RefreshRate) {
		return false
	}
	if this.DevMode != that1.DevMode {
		return false
	}
	if this.Linkerd != that1.Linkerd {
		return false
	}
	if !this.CircuitBreakers.Equal(that1.CircuitBreakers) {
		return false
	}
	if !this.Knative.Equal(that1.Knative) {
		return false
	}
	if !this.Extensions.Equal(that1.Extensions) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Settings_KubernetesConfigSource) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Settings_KubernetesConfigSource)
	if !ok {
		that2, ok := that.(Settings_KubernetesConfigSource)
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
	if !this.KubernetesConfigSource.Equal(that1.KubernetesConfigSource) {
		return false
	}
	return true
}
func (this *Settings_DirectoryConfigSource) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Settings_DirectoryConfigSource)
	if !ok {
		that2, ok := that.(Settings_DirectoryConfigSource)
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
	if !this.DirectoryConfigSource.Equal(that1.DirectoryConfigSource) {
		return false
	}
	return true
}
func (this *Settings_KubernetesSecretSource) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Settings_KubernetesSecretSource)
	if !ok {
		that2, ok := that.(Settings_KubernetesSecretSource)
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
	if !this.KubernetesSecretSource.Equal(that1.KubernetesSecretSource) {
		return false
	}
	return true
}
func (this *Settings_VaultSecretSource) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Settings_VaultSecretSource)
	if !ok {
		that2, ok := that.(Settings_VaultSecretSource)
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
	if !this.VaultSecretSource.Equal(that1.VaultSecretSource) {
		return false
	}
	return true
}
func (this *Settings_DirectorySecretSource) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Settings_DirectorySecretSource)
	if !ok {
		that2, ok := that.(Settings_DirectorySecretSource)
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
	if !this.DirectorySecretSource.Equal(that1.DirectorySecretSource) {
		return false
	}
	return true
}
func (this *Settings_KubernetesArtifactSource) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Settings_KubernetesArtifactSource)
	if !ok {
		that2, ok := that.(Settings_KubernetesArtifactSource)
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
	if !this.KubernetesArtifactSource.Equal(that1.KubernetesArtifactSource) {
		return false
	}
	return true
}
func (this *Settings_DirectoryArtifactSource) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Settings_DirectoryArtifactSource)
	if !ok {
		that2, ok := that.(Settings_DirectoryArtifactSource)
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
	if !this.DirectoryArtifactSource.Equal(that1.DirectoryArtifactSource) {
		return false
	}
	return true
}
func (this *Settings_KubernetesCrds) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Settings_KubernetesCrds)
	if !ok {
		that2, ok := that.(Settings_KubernetesCrds)
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
func (this *Settings_KubernetesSecrets) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Settings_KubernetesSecrets)
	if !ok {
		that2, ok := that.(Settings_KubernetesSecrets)
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
func (this *Settings_VaultSecrets) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Settings_VaultSecrets)
	if !ok {
		that2, ok := that.(Settings_VaultSecrets)
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
func (this *Settings_KubernetesConfigmaps) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Settings_KubernetesConfigmaps)
	if !ok {
		that2, ok := that.(Settings_KubernetesConfigmaps)
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
func (this *Settings_Directory) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Settings_Directory)
	if !ok {
		that2, ok := that.(Settings_Directory)
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
	if this.Directory != that1.Directory {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Settings_KnativeOptions) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Settings_KnativeOptions)
	if !ok {
		that2, ok := that.(Settings_KnativeOptions)
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
	if this.ClusterIngressProxyAddress != that1.ClusterIngressProxyAddress {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
