// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/load_balancer.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"
	time "time"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	lbhash "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/lbhash"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// LoadBalancerConfig is the settings for the load balancer used to send request to the Upstream
// endpoints.
type LoadBalancerConfig struct {
	// Configures envoy's panic threshold Percent between 0-100. Once the number of non health hosts
	// reaches this percentage, envoy disregards health information.
	// see more info [here](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/load_balancing/panic_threshold#arch-overview-load-balancing-panic-threshold).
	HealthyPanicThreshold *types.DoubleValue `protobuf:"bytes,1,opt,name=healthy_panic_threshold,json=healthyPanicThreshold,proto3" json:"healthy_panic_threshold,omitempty"`
	// This allows batch updates of endpoints health/weight/metadata that happen during a time window.
	// this help lower cpu usage when endpoint change rate is high. defaults to 1 second.
	// Set to 0 to disable and have changes applied immediately.
	UpdateMergeWindow *time.Duration `protobuf:"bytes,2,opt,name=update_merge_window,json=updateMergeWindow,proto3,stdduration" json:"update_merge_window,omitempty"`
	// Types that are valid to be assigned to Type:
	//	*LoadBalancerConfig_RoundRobin_
	//	*LoadBalancerConfig_LeastRequest_
	//	*LoadBalancerConfig_Random_
	//	*LoadBalancerConfig_RingHash_
	//	*LoadBalancerConfig_Maglev_
	Type                 isLoadBalancerConfig_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *LoadBalancerConfig) Reset()         { *m = LoadBalancerConfig{} }
func (m *LoadBalancerConfig) String() string { return proto.CompactTextString(m) }
func (*LoadBalancerConfig) ProtoMessage()    {}
func (*LoadBalancerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaa1c019b03e4b0f, []int{0}
}
func (m *LoadBalancerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadBalancerConfig.Unmarshal(m, b)
}
func (m *LoadBalancerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadBalancerConfig.Marshal(b, m, deterministic)
}
func (m *LoadBalancerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadBalancerConfig.Merge(m, src)
}
func (m *LoadBalancerConfig) XXX_Size() int {
	return xxx_messageInfo_LoadBalancerConfig.Size(m)
}
func (m *LoadBalancerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadBalancerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_LoadBalancerConfig proto.InternalMessageInfo

type isLoadBalancerConfig_Type interface {
	isLoadBalancerConfig_Type()
	Equal(interface{}) bool
}

type LoadBalancerConfig_RoundRobin_ struct {
	RoundRobin *LoadBalancerConfig_RoundRobin `protobuf:"bytes,3,opt,name=round_robin,json=roundRobin,proto3,oneof"`
}
type LoadBalancerConfig_LeastRequest_ struct {
	LeastRequest *LoadBalancerConfig_LeastRequest `protobuf:"bytes,4,opt,name=least_request,json=leastRequest,proto3,oneof"`
}
type LoadBalancerConfig_Random_ struct {
	Random *LoadBalancerConfig_Random `protobuf:"bytes,5,opt,name=random,proto3,oneof"`
}
type LoadBalancerConfig_RingHash_ struct {
	RingHash *LoadBalancerConfig_RingHash `protobuf:"bytes,6,opt,name=ring_hash,json=ringHash,proto3,oneof"`
}
type LoadBalancerConfig_Maglev_ struct {
	Maglev *LoadBalancerConfig_Maglev `protobuf:"bytes,7,opt,name=maglev,proto3,oneof"`
}

func (*LoadBalancerConfig_RoundRobin_) isLoadBalancerConfig_Type()   {}
func (*LoadBalancerConfig_LeastRequest_) isLoadBalancerConfig_Type() {}
func (*LoadBalancerConfig_Random_) isLoadBalancerConfig_Type()       {}
func (*LoadBalancerConfig_RingHash_) isLoadBalancerConfig_Type()     {}
func (*LoadBalancerConfig_Maglev_) isLoadBalancerConfig_Type()       {}

func (m *LoadBalancerConfig) GetType() isLoadBalancerConfig_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *LoadBalancerConfig) GetHealthyPanicThreshold() *types.DoubleValue {
	if m != nil {
		return m.HealthyPanicThreshold
	}
	return nil
}

func (m *LoadBalancerConfig) GetUpdateMergeWindow() *time.Duration {
	if m != nil {
		return m.UpdateMergeWindow
	}
	return nil
}

func (m *LoadBalancerConfig) GetRoundRobin() *LoadBalancerConfig_RoundRobin {
	if x, ok := m.GetType().(*LoadBalancerConfig_RoundRobin_); ok {
		return x.RoundRobin
	}
	return nil
}

func (m *LoadBalancerConfig) GetLeastRequest() *LoadBalancerConfig_LeastRequest {
	if x, ok := m.GetType().(*LoadBalancerConfig_LeastRequest_); ok {
		return x.LeastRequest
	}
	return nil
}

func (m *LoadBalancerConfig) GetRandom() *LoadBalancerConfig_Random {
	if x, ok := m.GetType().(*LoadBalancerConfig_Random_); ok {
		return x.Random
	}
	return nil
}

func (m *LoadBalancerConfig) GetRingHash() *LoadBalancerConfig_RingHash {
	if x, ok := m.GetType().(*LoadBalancerConfig_RingHash_); ok {
		return x.RingHash
	}
	return nil
}

func (m *LoadBalancerConfig) GetMaglev() *LoadBalancerConfig_Maglev {
	if x, ok := m.GetType().(*LoadBalancerConfig_Maglev_); ok {
		return x.Maglev
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*LoadBalancerConfig) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _LoadBalancerConfig_OneofMarshaler, _LoadBalancerConfig_OneofUnmarshaler, _LoadBalancerConfig_OneofSizer, []interface{}{
		(*LoadBalancerConfig_RoundRobin_)(nil),
		(*LoadBalancerConfig_LeastRequest_)(nil),
		(*LoadBalancerConfig_Random_)(nil),
		(*LoadBalancerConfig_RingHash_)(nil),
		(*LoadBalancerConfig_Maglev_)(nil),
	}
}

func _LoadBalancerConfig_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*LoadBalancerConfig)
	// type
	switch x := m.Type.(type) {
	case *LoadBalancerConfig_RoundRobin_:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RoundRobin); err != nil {
			return err
		}
	case *LoadBalancerConfig_LeastRequest_:
		_ = b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LeastRequest); err != nil {
			return err
		}
	case *LoadBalancerConfig_Random_:
		_ = b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Random); err != nil {
			return err
		}
	case *LoadBalancerConfig_RingHash_:
		_ = b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RingHash); err != nil {
			return err
		}
	case *LoadBalancerConfig_Maglev_:
		_ = b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Maglev); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("LoadBalancerConfig.Type has unexpected type %T", x)
	}
	return nil
}

func _LoadBalancerConfig_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*LoadBalancerConfig)
	switch tag {
	case 3: // type.round_robin
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LoadBalancerConfig_RoundRobin)
		err := b.DecodeMessage(msg)
		m.Type = &LoadBalancerConfig_RoundRobin_{msg}
		return true, err
	case 4: // type.least_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LoadBalancerConfig_LeastRequest)
		err := b.DecodeMessage(msg)
		m.Type = &LoadBalancerConfig_LeastRequest_{msg}
		return true, err
	case 5: // type.random
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LoadBalancerConfig_Random)
		err := b.DecodeMessage(msg)
		m.Type = &LoadBalancerConfig_Random_{msg}
		return true, err
	case 6: // type.ring_hash
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LoadBalancerConfig_RingHash)
		err := b.DecodeMessage(msg)
		m.Type = &LoadBalancerConfig_RingHash_{msg}
		return true, err
	case 7: // type.maglev
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LoadBalancerConfig_Maglev)
		err := b.DecodeMessage(msg)
		m.Type = &LoadBalancerConfig_Maglev_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _LoadBalancerConfig_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*LoadBalancerConfig)
	// type
	switch x := m.Type.(type) {
	case *LoadBalancerConfig_RoundRobin_:
		s := proto.Size(x.RoundRobin)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LoadBalancerConfig_LeastRequest_:
		s := proto.Size(x.LeastRequest)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LoadBalancerConfig_Random_:
		s := proto.Size(x.Random)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LoadBalancerConfig_RingHash_:
		s := proto.Size(x.RingHash)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LoadBalancerConfig_Maglev_:
		s := proto.Size(x.Maglev)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type LoadBalancerConfig_RoundRobin struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoadBalancerConfig_RoundRobin) Reset()         { *m = LoadBalancerConfig_RoundRobin{} }
func (m *LoadBalancerConfig_RoundRobin) String() string { return proto.CompactTextString(m) }
func (*LoadBalancerConfig_RoundRobin) ProtoMessage()    {}
func (*LoadBalancerConfig_RoundRobin) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaa1c019b03e4b0f, []int{0, 0}
}
func (m *LoadBalancerConfig_RoundRobin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadBalancerConfig_RoundRobin.Unmarshal(m, b)
}
func (m *LoadBalancerConfig_RoundRobin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadBalancerConfig_RoundRobin.Marshal(b, m, deterministic)
}
func (m *LoadBalancerConfig_RoundRobin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadBalancerConfig_RoundRobin.Merge(m, src)
}
func (m *LoadBalancerConfig_RoundRobin) XXX_Size() int {
	return xxx_messageInfo_LoadBalancerConfig_RoundRobin.Size(m)
}
func (m *LoadBalancerConfig_RoundRobin) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadBalancerConfig_RoundRobin.DiscardUnknown(m)
}

var xxx_messageInfo_LoadBalancerConfig_RoundRobin proto.InternalMessageInfo

type LoadBalancerConfig_LeastRequest struct {
	// How many choices to take into account. defaults to 2.
	ChoiceCount          uint32   `protobuf:"varint,1,opt,name=choice_count,json=choiceCount,proto3" json:"choice_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoadBalancerConfig_LeastRequest) Reset()         { *m = LoadBalancerConfig_LeastRequest{} }
func (m *LoadBalancerConfig_LeastRequest) String() string { return proto.CompactTextString(m) }
func (*LoadBalancerConfig_LeastRequest) ProtoMessage()    {}
func (*LoadBalancerConfig_LeastRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaa1c019b03e4b0f, []int{0, 1}
}
func (m *LoadBalancerConfig_LeastRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadBalancerConfig_LeastRequest.Unmarshal(m, b)
}
func (m *LoadBalancerConfig_LeastRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadBalancerConfig_LeastRequest.Marshal(b, m, deterministic)
}
func (m *LoadBalancerConfig_LeastRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadBalancerConfig_LeastRequest.Merge(m, src)
}
func (m *LoadBalancerConfig_LeastRequest) XXX_Size() int {
	return xxx_messageInfo_LoadBalancerConfig_LeastRequest.Size(m)
}
func (m *LoadBalancerConfig_LeastRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadBalancerConfig_LeastRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoadBalancerConfig_LeastRequest proto.InternalMessageInfo

func (m *LoadBalancerConfig_LeastRequest) GetChoiceCount() uint32 {
	if m != nil {
		return m.ChoiceCount
	}
	return 0
}

type LoadBalancerConfig_Random struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoadBalancerConfig_Random) Reset()         { *m = LoadBalancerConfig_Random{} }
func (m *LoadBalancerConfig_Random) String() string { return proto.CompactTextString(m) }
func (*LoadBalancerConfig_Random) ProtoMessage()    {}
func (*LoadBalancerConfig_Random) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaa1c019b03e4b0f, []int{0, 2}
}
func (m *LoadBalancerConfig_Random) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadBalancerConfig_Random.Unmarshal(m, b)
}
func (m *LoadBalancerConfig_Random) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadBalancerConfig_Random.Marshal(b, m, deterministic)
}
func (m *LoadBalancerConfig_Random) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadBalancerConfig_Random.Merge(m, src)
}
func (m *LoadBalancerConfig_Random) XXX_Size() int {
	return xxx_messageInfo_LoadBalancerConfig_Random.Size(m)
}
func (m *LoadBalancerConfig_Random) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadBalancerConfig_Random.DiscardUnknown(m)
}

var xxx_messageInfo_LoadBalancerConfig_Random proto.InternalMessageInfo

// Customizes the parameters used in the hashing algorithm to refine performance or resource usage.
type LoadBalancerConfig_RingHashConfig struct {
	// Minimum hash ring size. The larger the ring is (that is, the more hashes there are for each provided host)
	// the better the request distribution will reflect the desired weights. Defaults to 1024 entries, and limited
	// to 8M entries.
	MinimumRingSize uint64 `protobuf:"varint,1,opt,name=minimum_ring_size,json=minimumRingSize,proto3" json:"minimum_ring_size,omitempty"`
	// Maximum hash ring size. Defaults to 8M entries, and limited to 8M entries, but can be lowered to further
	// constrain resource use.
	MaximumRingSize uint64 `protobuf:"varint,2,opt,name=maximum_ring_size,json=maximumRingSize,proto3" json:"maximum_ring_size,omitempty"`
	// Optional, if set, routes to this upstream will use this hash policy, unless a policy has been set on the route
	// Gloo configures Envoy with the first available RouteActionHashConfig among the following ordered list of providers:
	// - route, upstream, virtual service
	DefaultHashPolicy    *lbhash.RouteActionHashConfig `protobuf:"bytes,12,opt,name=default_hash_policy,json=defaultHashPolicy,proto3" json:"default_hash_policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *LoadBalancerConfig_RingHashConfig) Reset()         { *m = LoadBalancerConfig_RingHashConfig{} }
func (m *LoadBalancerConfig_RingHashConfig) String() string { return proto.CompactTextString(m) }
func (*LoadBalancerConfig_RingHashConfig) ProtoMessage()    {}
func (*LoadBalancerConfig_RingHashConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaa1c019b03e4b0f, []int{0, 3}
}
func (m *LoadBalancerConfig_RingHashConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadBalancerConfig_RingHashConfig.Unmarshal(m, b)
}
func (m *LoadBalancerConfig_RingHashConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadBalancerConfig_RingHashConfig.Marshal(b, m, deterministic)
}
func (m *LoadBalancerConfig_RingHashConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadBalancerConfig_RingHashConfig.Merge(m, src)
}
func (m *LoadBalancerConfig_RingHashConfig) XXX_Size() int {
	return xxx_messageInfo_LoadBalancerConfig_RingHashConfig.Size(m)
}
func (m *LoadBalancerConfig_RingHashConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadBalancerConfig_RingHashConfig.DiscardUnknown(m)
}

var xxx_messageInfo_LoadBalancerConfig_RingHashConfig proto.InternalMessageInfo

func (m *LoadBalancerConfig_RingHashConfig) GetMinimumRingSize() uint64 {
	if m != nil {
		return m.MinimumRingSize
	}
	return 0
}

func (m *LoadBalancerConfig_RingHashConfig) GetMaximumRingSize() uint64 {
	if m != nil {
		return m.MaximumRingSize
	}
	return 0
}

func (m *LoadBalancerConfig_RingHashConfig) GetDefaultHashPolicy() *lbhash.RouteActionHashConfig {
	if m != nil {
		return m.DefaultHashPolicy
	}
	return nil
}

type LoadBalancerConfig_RingHash struct {
	// Optional, customizes the parameters used in the hashing algorithm
	RingHashConfig       *LoadBalancerConfig_RingHashConfig `protobuf:"bytes,1,opt,name=ring_hash_config,json=ringHashConfig,proto3" json:"ring_hash_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *LoadBalancerConfig_RingHash) Reset()         { *m = LoadBalancerConfig_RingHash{} }
func (m *LoadBalancerConfig_RingHash) String() string { return proto.CompactTextString(m) }
func (*LoadBalancerConfig_RingHash) ProtoMessage()    {}
func (*LoadBalancerConfig_RingHash) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaa1c019b03e4b0f, []int{0, 4}
}
func (m *LoadBalancerConfig_RingHash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadBalancerConfig_RingHash.Unmarshal(m, b)
}
func (m *LoadBalancerConfig_RingHash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadBalancerConfig_RingHash.Marshal(b, m, deterministic)
}
func (m *LoadBalancerConfig_RingHash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadBalancerConfig_RingHash.Merge(m, src)
}
func (m *LoadBalancerConfig_RingHash) XXX_Size() int {
	return xxx_messageInfo_LoadBalancerConfig_RingHash.Size(m)
}
func (m *LoadBalancerConfig_RingHash) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadBalancerConfig_RingHash.DiscardUnknown(m)
}

var xxx_messageInfo_LoadBalancerConfig_RingHash proto.InternalMessageInfo

func (m *LoadBalancerConfig_RingHash) GetRingHashConfig() *LoadBalancerConfig_RingHashConfig {
	if m != nil {
		return m.RingHashConfig
	}
	return nil
}

type LoadBalancerConfig_Maglev struct {
	// Optional, customizes the parameters used in the hashing algorithm
	RingHashConfig       *LoadBalancerConfig_RingHashConfig `protobuf:"bytes,1,opt,name=ring_hash_config,json=ringHashConfig,proto3" json:"ring_hash_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *LoadBalancerConfig_Maglev) Reset()         { *m = LoadBalancerConfig_Maglev{} }
func (m *LoadBalancerConfig_Maglev) String() string { return proto.CompactTextString(m) }
func (*LoadBalancerConfig_Maglev) ProtoMessage()    {}
func (*LoadBalancerConfig_Maglev) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaa1c019b03e4b0f, []int{0, 5}
}
func (m *LoadBalancerConfig_Maglev) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadBalancerConfig_Maglev.Unmarshal(m, b)
}
func (m *LoadBalancerConfig_Maglev) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadBalancerConfig_Maglev.Marshal(b, m, deterministic)
}
func (m *LoadBalancerConfig_Maglev) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadBalancerConfig_Maglev.Merge(m, src)
}
func (m *LoadBalancerConfig_Maglev) XXX_Size() int {
	return xxx_messageInfo_LoadBalancerConfig_Maglev.Size(m)
}
func (m *LoadBalancerConfig_Maglev) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadBalancerConfig_Maglev.DiscardUnknown(m)
}

var xxx_messageInfo_LoadBalancerConfig_Maglev proto.InternalMessageInfo

func (m *LoadBalancerConfig_Maglev) GetRingHashConfig() *LoadBalancerConfig_RingHashConfig {
	if m != nil {
		return m.RingHashConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*LoadBalancerConfig)(nil), "gloo.solo.io.LoadBalancerConfig")
	proto.RegisterType((*LoadBalancerConfig_RoundRobin)(nil), "gloo.solo.io.LoadBalancerConfig.RoundRobin")
	proto.RegisterType((*LoadBalancerConfig_LeastRequest)(nil), "gloo.solo.io.LoadBalancerConfig.LeastRequest")
	proto.RegisterType((*LoadBalancerConfig_Random)(nil), "gloo.solo.io.LoadBalancerConfig.Random")
	proto.RegisterType((*LoadBalancerConfig_RingHashConfig)(nil), "gloo.solo.io.LoadBalancerConfig.RingHashConfig")
	proto.RegisterType((*LoadBalancerConfig_RingHash)(nil), "gloo.solo.io.LoadBalancerConfig.RingHash")
	proto.RegisterType((*LoadBalancerConfig_Maglev)(nil), "gloo.solo.io.LoadBalancerConfig.Maglev")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/load_balancer.proto", fileDescriptor_aaa1c019b03e4b0f)
}

var fileDescriptor_aaa1c019b03e4b0f = []byte{
	// 609 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0x4d, 0x6f, 0xd4, 0x3c,
	0x10, 0xc7, 0x77, 0xfb, 0xec, 0xb3, 0x2c, 0xee, 0xb6, 0xd0, 0x14, 0x44, 0x88, 0x50, 0x79, 0xb9,
	0xf0, 0xa6, 0x26, 0xb4, 0x48, 0x9c, 0xe9, 0x96, 0x43, 0x0e, 0x2d, 0x54, 0xa6, 0x02, 0xc1, 0x25,
	0x72, 0x12, 0xd7, 0x31, 0x38, 0x9e, 0xe0, 0xd8, 0x2d, 0xed, 0x27, 0xe1, 0xc6, 0x95, 0x0f, 0xc3,
	0x77, 0x40, 0xe2, 0x93, 0x20, 0x3b, 0xde, 0xbe, 0x50, 0xa1, 0x2e, 0x07, 0x4e, 0xf1, 0xd8, 0xf3,
	0xff, 0x65, 0xc6, 0x33, 0x63, 0xf4, 0x9c, 0x71, 0x5d, 0x99, 0x3c, 0x2e, 0xa0, 0x4e, 0x5a, 0x10,
	0xb0, 0xca, 0x21, 0x61, 0x02, 0x20, 0x69, 0x14, 0x7c, 0xa0, 0x85, 0x6e, 0x3b, 0x8b, 0x34, 0x3c,
	0xd9, 0x5f, 0x4b, 0x04, 0x90, 0x32, 0xcb, 0x89, 0x20, 0xb2, 0xa0, 0x2a, 0x6e, 0x14, 0x68, 0x08,
	0xc6, 0xd6, 0x21, 0xb6, 0xda, 0x98, 0x43, 0x94, 0xfe, 0x15, 0xaf, 0x11, 0x86, 0x71, 0xd9, 0x26,
	0x22, 0xaf, 0x48, 0x5b, 0xf9, 0x4f, 0xc7, 0x8d, 0xae, 0x31, 0x60, 0xe0, 0x96, 0x89, 0x5d, 0xf9,
	0xdd, 0x15, 0x06, 0xc0, 0x04, 0x4d, 0x9c, 0x95, 0x9b, 0xbd, 0xa4, 0x34, 0x8a, 0x68, 0x0e, 0xf2,
	0x4f, 0xe7, 0x07, 0x8a, 0x34, 0x0d, 0x55, 0x6d, 0x77, 0x7e, 0xef, 0xeb, 0x08, 0x05, 0x5b, 0x40,
	0xca, 0x89, 0x4f, 0x62, 0x13, 0xe4, 0x1e, 0x67, 0xc1, 0x2e, 0xba, 0x51, 0x51, 0x22, 0x74, 0x75,
	0x98, 0x35, 0x44, 0xf2, 0x22, 0xd3, 0x95, 0xa2, 0x6d, 0x05, 0xa2, 0x0c, 0xfb, 0x77, 0xfa, 0x0f,
	0xe6, 0xd7, 0x6f, 0xc5, 0x1d, 0x38, 0x9e, 0x82, 0xe3, 0x17, 0x60, 0x72, 0x41, 0xdf, 0x10, 0x61,
	0x28, 0xbe, 0xee, 0xc5, 0x3b, 0x56, 0xbb, 0x3b, 0x95, 0x06, 0xaf, 0xd0, 0xb2, 0x69, 0x4a, 0xa2,
	0x69, 0x56, 0x53, 0xc5, 0x68, 0x76, 0xc0, 0x65, 0x09, 0x07, 0xe1, 0x9c, 0x23, 0xde, 0x3c, 0x4f,
	0xf4, 0xa9, 0x4c, 0x06, 0x5f, 0x7e, 0xdc, 0xee, 0xe3, 0xa5, 0x4e, 0xbb, 0x6d, 0xa5, 0x6f, 0x9d,
	0x32, 0x78, 0x89, 0xe6, 0x15, 0x18, 0x59, 0x66, 0x0a, 0x72, 0x2e, 0xc3, 0xff, 0x1c, 0xe8, 0x71,
	0x7c, 0xba, 0x02, 0xf1, 0xf9, 0xec, 0x62, 0x6c, 0x35, 0xd8, 0x4a, 0xd2, 0x1e, 0x46, 0xea, 0xd8,
	0x0a, 0x76, 0xd1, 0x82, 0xa0, 0xa4, 0xd5, 0x99, 0xa2, 0x9f, 0x0c, 0x6d, 0x75, 0x38, 0x70, 0xc4,
	0xd5, 0x0b, 0x89, 0x5b, 0x56, 0x85, 0x3b, 0x51, 0xda, 0xc3, 0x63, 0x71, 0xca, 0x0e, 0x36, 0xd0,
	0x50, 0x11, 0x59, 0x42, 0x1d, 0xfe, 0xef, 0x70, 0xf7, 0x2f, 0x0e, 0xd0, 0xb9, 0xa7, 0x3d, 0xec,
	0x85, 0x41, 0x8a, 0x2e, 0x2b, 0x2e, 0x59, 0x66, 0xfb, 0x21, 0x1c, 0x3a, 0xca, 0xc3, 0x8b, 0x29,
	0x5c, 0xb2, 0x94, 0xb4, 0x55, 0xda, 0xc3, 0x23, 0xe5, 0xd7, 0x36, 0x98, 0x9a, 0x30, 0x41, 0xf7,
	0xc3, 0x4b, 0x33, 0x06, 0xb3, 0xed, 0xdc, 0x6d, 0x30, 0x9d, 0x30, 0x1a, 0x23, 0x74, 0x72, 0x83,
	0xd1, 0x1a, 0x1a, 0x9f, 0xce, 0x3e, 0xb8, 0x8b, 0xc6, 0x45, 0x05, 0xbc, 0xa0, 0x59, 0x01, 0x46,
	0x6a, 0xd7, 0x2f, 0x0b, 0x78, 0xbe, 0xdb, 0xdb, 0xb4, 0x5b, 0xd1, 0x08, 0x0d, 0xbb, 0x0c, 0xa3,
	0xef, 0x7d, 0xb4, 0x38, 0x0d, 0xd3, 0xb7, 0xde, 0x23, 0xb4, 0x54, 0x73, 0xc9, 0x6b, 0x53, 0x67,
	0x2e, 0xe5, 0x96, 0x1f, 0x51, 0x07, 0x19, 0xe0, 0x2b, 0xfe, 0xc0, 0x2a, 0x5e, 0xf3, 0x23, 0xea,
	0x7c, 0xc9, 0xe7, 0xdf, 0x7c, 0xe7, 0xbc, 0x6f, 0x77, 0x70, 0xec, 0x9b, 0xa3, 0xe5, 0x92, 0xee,
	0x11, 0x23, 0xb4, 0xbb, 0xc5, 0xac, 0x01, 0xc1, 0x8b, 0xc3, 0x70, 0xec, 0x6e, 0x61, 0x3d, 0x9e,
	0xce, 0x5a, 0x37, 0x81, 0x67, 0x2f, 0x05, 0x83, 0xd1, 0x74, 0xa3, 0xb0, 0xbd, 0x78, 0x12, 0x28,
	0x5e, 0xf2, 0x38, 0xbb, 0xb5, 0xe3, 0x60, 0x11, 0x45, 0xa3, 0x69, 0x36, 0xc1, 0x3b, 0x74, 0xf5,
	0xb8, 0x64, 0x59, 0xe1, 0x24, 0x7e, 0x76, 0x92, 0x99, 0x2b, 0xe7, 0xff, 0xb4, 0xa8, 0xce, 0xd8,
	0x51, 0x81, 0x86, 0x5d, 0x51, 0xfe, 0xe1, 0x4f, 0x26, 0x43, 0x34, 0xd0, 0x87, 0x0d, 0x9d, 0x3c,
	0xfb, 0xf6, 0x73, 0xa5, 0xff, 0xfe, 0xc9, 0x6c, 0xef, 0x58, 0xf3, 0x91, 0xf9, 0xb7, 0x2c, 0x1f,
	0xba, 0x39, 0x7e, 0xfa, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x07, 0x11, 0x11, 0x4d, 0x52, 0x05, 0x00,
	0x00,
}

func (this *LoadBalancerConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LoadBalancerConfig)
	if !ok {
		that2, ok := that.(LoadBalancerConfig)
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
	if !this.HealthyPanicThreshold.Equal(that1.HealthyPanicThreshold) {
		return false
	}
	if this.UpdateMergeWindow != nil && that1.UpdateMergeWindow != nil {
		if *this.UpdateMergeWindow != *that1.UpdateMergeWindow {
			return false
		}
	} else if this.UpdateMergeWindow != nil {
		return false
	} else if that1.UpdateMergeWindow != nil {
		return false
	}
	if that1.Type == nil {
		if this.Type != nil {
			return false
		}
	} else if this.Type == nil {
		return false
	} else if !this.Type.Equal(that1.Type) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *LoadBalancerConfig_RoundRobin_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LoadBalancerConfig_RoundRobin_)
	if !ok {
		that2, ok := that.(LoadBalancerConfig_RoundRobin_)
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
	if !this.RoundRobin.Equal(that1.RoundRobin) {
		return false
	}
	return true
}
func (this *LoadBalancerConfig_LeastRequest_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LoadBalancerConfig_LeastRequest_)
	if !ok {
		that2, ok := that.(LoadBalancerConfig_LeastRequest_)
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
	if !this.LeastRequest.Equal(that1.LeastRequest) {
		return false
	}
	return true
}
func (this *LoadBalancerConfig_Random_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LoadBalancerConfig_Random_)
	if !ok {
		that2, ok := that.(LoadBalancerConfig_Random_)
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
	if !this.Random.Equal(that1.Random) {
		return false
	}
	return true
}
func (this *LoadBalancerConfig_RingHash_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LoadBalancerConfig_RingHash_)
	if !ok {
		that2, ok := that.(LoadBalancerConfig_RingHash_)
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
	if !this.RingHash.Equal(that1.RingHash) {
		return false
	}
	return true
}
func (this *LoadBalancerConfig_Maglev_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LoadBalancerConfig_Maglev_)
	if !ok {
		that2, ok := that.(LoadBalancerConfig_Maglev_)
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
	if !this.Maglev.Equal(that1.Maglev) {
		return false
	}
	return true
}
func (this *LoadBalancerConfig_RoundRobin) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LoadBalancerConfig_RoundRobin)
	if !ok {
		that2, ok := that.(LoadBalancerConfig_RoundRobin)
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
func (this *LoadBalancerConfig_LeastRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LoadBalancerConfig_LeastRequest)
	if !ok {
		that2, ok := that.(LoadBalancerConfig_LeastRequest)
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
	if this.ChoiceCount != that1.ChoiceCount {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *LoadBalancerConfig_Random) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LoadBalancerConfig_Random)
	if !ok {
		that2, ok := that.(LoadBalancerConfig_Random)
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
func (this *LoadBalancerConfig_RingHashConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LoadBalancerConfig_RingHashConfig)
	if !ok {
		that2, ok := that.(LoadBalancerConfig_RingHashConfig)
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
	if this.MinimumRingSize != that1.MinimumRingSize {
		return false
	}
	if this.MaximumRingSize != that1.MaximumRingSize {
		return false
	}
	if !this.DefaultHashPolicy.Equal(that1.DefaultHashPolicy) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *LoadBalancerConfig_RingHash) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LoadBalancerConfig_RingHash)
	if !ok {
		that2, ok := that.(LoadBalancerConfig_RingHash)
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
	if !this.RingHashConfig.Equal(that1.RingHashConfig) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *LoadBalancerConfig_Maglev) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LoadBalancerConfig_Maglev)
	if !ok {
		that2, ok := that.(LoadBalancerConfig_Maglev)
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
	if !this.RingHashConfig.Equal(that1.RingHashConfig) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
