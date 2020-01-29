// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gateway/api/v1/virtual_service.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	matchers "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/core/matchers"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//
// The **VirtualService** is the root routing object for the Gloo Gateway.
// A virtual service describes the set of routes to match for a set of domains.
//
// It defines:
// - a set of domains
// - the root set of routes for those domains
// - an optional SSL configuration for server TLS Termination
// - VirtualHostOptions that will apply configuration to all routes that live on the VirtualService.
//
// Domains must be unique across all virtual services within a gateway (i.e. no overlap between sets).
//
// VirtualServices can delegate routing behavior to the RouteTable resource by using the `delegateAction` on routes.
//
// An example configuration using two VirtualServices (one with TLS termination and one without) which share
// a RouteTable looks as follows:
//
// ```yaml
// # HTTP VirtualService:
// apiVersion: gateway.solo.io/v1
// kind: VirtualService
// metadata:
//   name: 'http'
//   namespace: 'usernamespace'
// spec:
//   virtualHost:
//     domains:
//     - '*.mydomain.com'
//     - 'mydomain.com'
//     routes:
//     - matchers:
//       - prefix: '/'
//       # delegate all traffic to the `shared-routes` RouteTable
//       delegateAction:
//         ref:
//           name: 'shared-routes'
//           namespace: 'usernamespace'
//
// ```
//
// ```yaml
// # HTTPS VirtualService:
// apiVersion: gateway.solo.io/v1
// kind: VirtualService
// metadata:
//   name: 'https'
//   namespace: 'usernamespace'
// spec:
//   virtualHost:
//     domains:
//     - '*.mydomain.com'
//     - 'mydomain.com'
//     routes:
//     - matchers:
//       - prefix: '/'
//       # delegate all traffic to the `shared-routes` RouteTable
//       delegateAction:
//         ref:
//           name: 'shared-routes'
//           namespace: 'usernamespace'
//   sslConfig:
//     secretRef:
//       name: gateway-tls
//       namespace: gloo-system
//
// ```
//
// ```yaml
// # the RouteTable shared by both VirtualServices:
// apiVersion: gateway.solo.io/v1
// kind: RouteTable
// metadata:
//   name: 'shared-routes'
//   namespace: 'usernamespace'
// spec:
//   routes:
//     - matchers:
//       - prefix: '/some-route'
//       routeAction:
//         single:
//           upstream:
//             name: 'some-upstream'
//      ...
// ```
//
// **Delegated Routes** are routes that use the `delegateAction` routing action. Delegated Routes obey the following
// constraints:
//
// - delegate routes must use `prefix` path matchers
// - delegated routes cannot specify header, query, or methods portion of the normal route matcher.
// - `routeOptions` configuration will be inherited from parent routes, but can be overridden by the child
//
type VirtualService struct {
	// The VirtualHost contains the
	// The list of HTTP routes define routing actions to be taken
	// for incoming HTTP requests whose host header matches
	// this virtual host. If the request matches more than one route in the list, the first route matched will be selected.
	// If the list of routes is empty, the virtual host will be ignored by Gloo.
	VirtualHost *VirtualHost `protobuf:"bytes,1,opt,name=virtual_host,json=virtualHost,proto3" json:"virtual_host,omitempty"`
	// If provided, the Gateway will serve TLS/SSL traffic for this set of routes
	SslConfig *v1.SslConfig `protobuf:"bytes,2,opt,name=ssl_config,json=sslConfig,proto3" json:"ssl_config,omitempty"`
	// Display only, optional descriptive name.
	// Unlike metadata.name, DisplayName can be any string
	// and can be changed after creating the resource.
	DisplayName string `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Status indicates the validation status of this resource.
	// Status is read-only by clients, and set by gloo during validation
	Status core.Status `protobuf:"bytes,6,opt,name=status,proto3" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata             core.Metadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *VirtualService) Reset()         { *m = VirtualService{} }
func (m *VirtualService) String() string { return proto.CompactTextString(m) }
func (*VirtualService) ProtoMessage()    {}
func (*VirtualService) Descriptor() ([]byte, []int) {
	return fileDescriptor_93fa9472926a2049, []int{0}
}
func (m *VirtualService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualService.Unmarshal(m, b)
}
func (m *VirtualService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualService.Marshal(b, m, deterministic)
}
func (m *VirtualService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualService.Merge(m, src)
}
func (m *VirtualService) XXX_Size() int {
	return xxx_messageInfo_VirtualService.Size(m)
}
func (m *VirtualService) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualService.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualService proto.InternalMessageInfo

func (m *VirtualService) GetVirtualHost() *VirtualHost {
	if m != nil {
		return m.VirtualHost
	}
	return nil
}

func (m *VirtualService) GetSslConfig() *v1.SslConfig {
	if m != nil {
		return m.SslConfig
	}
	return nil
}

func (m *VirtualService) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *VirtualService) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *VirtualService) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

//
//Virtual Hosts serve an ordered list of routes for a set of domains.
//
//An HTTP request is first matched to a virtual host based on its host header, then to a route within the virtual host.
//
//If a request is not matched to any virtual host or a route therein, the target proxy will reply with a 404.
//
//Unlike the [Gloo Virtual Host]({{< ref "/api/github.com/solo-io/gloo/projects/gloo/api/v1/proxy.proto.sk.md" >}}/#virtualhost),
//Gateway* Virtual Hosts can **delegate** their routes to `RouteTables`.
//
type VirtualHost struct {
	// The list of domains (i.e.: matching the `Host` header of a request) that belong to this virtual host.
	// Note that the wildcard will not match the empty string. e.g. “*-bar.foo.com” will match “baz-bar.foo.com”
	// but not “-bar.foo.com”. Additionally, a special entry “*” is allowed which will match any host/authority header.
	// Only a single virtual host on a gateway can match on “*”. A domain must be unique across all
	// virtual hosts on a gateway or the config will be invalidated by Gloo
	// Domains on virtual hosts obey the same rules as [Envoy Virtual Hosts](https://github.com/envoyproxy/envoy/blob/master/api/envoy/api/v2/route/route.proto)
	Domains []string `protobuf:"bytes,2,rep,name=domains,proto3" json:"domains,omitempty"`
	// The list of HTTP routes define routing actions to be taken for incoming HTTP requests whose host header matches
	// this virtual host. If the request matches more than one route in the list, the first route matched will be selected.
	// If the list of routes is empty, the virtual host will be ignored by Gloo.
	Routes []*Route `protobuf:"bytes,3,rep,name=routes,proto3" json:"routes,omitempty"`
	// Virtual host options contain additional configuration to be applied to all traffic served by the Virtual Host.
	// Some configuration here can be overridden by Route Options.
	Options              *v1.VirtualHostOptions `protobuf:"bytes,4,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *VirtualHost) Reset()         { *m = VirtualHost{} }
func (m *VirtualHost) String() string { return proto.CompactTextString(m) }
func (*VirtualHost) ProtoMessage()    {}
func (*VirtualHost) Descriptor() ([]byte, []int) {
	return fileDescriptor_93fa9472926a2049, []int{1}
}
func (m *VirtualHost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualHost.Unmarshal(m, b)
}
func (m *VirtualHost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualHost.Marshal(b, m, deterministic)
}
func (m *VirtualHost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualHost.Merge(m, src)
}
func (m *VirtualHost) XXX_Size() int {
	return xxx_messageInfo_VirtualHost.Size(m)
}
func (m *VirtualHost) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualHost.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualHost proto.InternalMessageInfo

func (m *VirtualHost) GetDomains() []string {
	if m != nil {
		return m.Domains
	}
	return nil
}

func (m *VirtualHost) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

func (m *VirtualHost) GetOptions() *v1.VirtualHostOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

//
// A route specifies how to match a request and what action to take when the request is matched.
//
// When a request matches on a route, the route can perform one of the following actions:
// - *Route* the request to a destination
// - Reply with a *Direct Response*
// - Send a *Redirect* response to the client
// - *Delegate* the action for the request to one or more top-level [`RouteTable`]({{< ref "/api/github.com/solo-io/gloo/projects/gateway/api/v1/route_table.proto.sk.md" >}}) resources
// DelegateActions can be used to delegate the behavior for a set out routes with a given *prefix* to
// top-level `RouteTable` resources.
type Route struct {
	// Matchers contain parameters for matching requests (i.e., based on HTTP path, headers, etc.)
	// If empty, the route will match all requests (i.e, a single "/" path prefix matcher)
	// For delegated routes, the matcher must contain only a `prefix` path matcher and no other config
	Matchers []*matchers.Matcher `protobuf:"bytes,1,rep,name=matchers,proto3" json:"matchers,omitempty"`
	// The Route Action Defines what action the proxy should take when a request matches the route.
	//
	// Types that are valid to be assigned to Action:
	//	*Route_RouteAction
	//	*Route_RedirectAction
	//	*Route_DirectResponseAction
	//	*Route_DelegateAction
	Action isRoute_Action `protobuf_oneof:"action"`
	// Route Options extend the behavior of routes.
	// Route options include configuration such as retries, rate limiting, and request/response transformation.
	// RouteOption behavior will be inherited by delegated routes which do not specify their own `options`
	Options              *v1.RouteOptions `protobuf:"bytes,6,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_93fa9472926a2049, []int{2}
}
func (m *Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route.Unmarshal(m, b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route.Marshal(b, m, deterministic)
}
func (m *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(m, src)
}
func (m *Route) XXX_Size() int {
	return xxx_messageInfo_Route.Size(m)
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

type isRoute_Action interface {
	isRoute_Action()
	Equal(interface{}) bool
}

type Route_RouteAction struct {
	RouteAction *v1.RouteAction `protobuf:"bytes,2,opt,name=route_action,json=routeAction,proto3,oneof" json:"route_action,omitempty"`
}
type Route_RedirectAction struct {
	RedirectAction *v1.RedirectAction `protobuf:"bytes,3,opt,name=redirect_action,json=redirectAction,proto3,oneof" json:"redirect_action,omitempty"`
}
type Route_DirectResponseAction struct {
	DirectResponseAction *v1.DirectResponseAction `protobuf:"bytes,4,opt,name=direct_response_action,json=directResponseAction,proto3,oneof" json:"direct_response_action,omitempty"`
}
type Route_DelegateAction struct {
	DelegateAction *DelegateAction `protobuf:"bytes,5,opt,name=delegate_action,json=delegateAction,proto3,oneof" json:"delegate_action,omitempty"`
}

func (*Route_RouteAction) isRoute_Action()          {}
func (*Route_RedirectAction) isRoute_Action()       {}
func (*Route_DirectResponseAction) isRoute_Action() {}
func (*Route_DelegateAction) isRoute_Action()       {}

func (m *Route) GetAction() isRoute_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (m *Route) GetMatchers() []*matchers.Matcher {
	if m != nil {
		return m.Matchers
	}
	return nil
}

func (m *Route) GetRouteAction() *v1.RouteAction {
	if x, ok := m.GetAction().(*Route_RouteAction); ok {
		return x.RouteAction
	}
	return nil
}

func (m *Route) GetRedirectAction() *v1.RedirectAction {
	if x, ok := m.GetAction().(*Route_RedirectAction); ok {
		return x.RedirectAction
	}
	return nil
}

func (m *Route) GetDirectResponseAction() *v1.DirectResponseAction {
	if x, ok := m.GetAction().(*Route_DirectResponseAction); ok {
		return x.DirectResponseAction
	}
	return nil
}

func (m *Route) GetDelegateAction() *DelegateAction {
	if x, ok := m.GetAction().(*Route_DelegateAction); ok {
		return x.DelegateAction
	}
	return nil
}

func (m *Route) GetOptions() *v1.RouteOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Route) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Route_RouteAction)(nil),
		(*Route_RedirectAction)(nil),
		(*Route_DirectResponseAction)(nil),
		(*Route_DelegateAction)(nil),
	}
}

// DelegateActions are used to delegate routing decisions to Route Tables.
type DelegateAction struct {
	// The name of the Route Table to delegate to.
	// Deprecated: these fields have been added for backwards-compatibility. Please use the `single` field. If `name`
	// and/or `namespace` have been specified, Gloo will ignore `single` and `selector`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` // Deprecated: Do not use.
	// The namespace of the Route Table to delegate to.
	// Deprecated: these fields have been added for backwards-compatibility. Please use the `single` field. If `name`
	// and/or `namespace` have been specified, Gloo will ignore `single` and `selector`.
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"` // Deprecated: Do not use.
	// Types that are valid to be assigned to DelegationType:
	//	*DelegateAction_Ref
	//	*DelegateAction_Selector
	DelegationType       isDelegateAction_DelegationType `protobuf_oneof:"delegation_type"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *DelegateAction) Reset()         { *m = DelegateAction{} }
func (m *DelegateAction) String() string { return proto.CompactTextString(m) }
func (*DelegateAction) ProtoMessage()    {}
func (*DelegateAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_93fa9472926a2049, []int{3}
}
func (m *DelegateAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelegateAction.Unmarshal(m, b)
}
func (m *DelegateAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelegateAction.Marshal(b, m, deterministic)
}
func (m *DelegateAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelegateAction.Merge(m, src)
}
func (m *DelegateAction) XXX_Size() int {
	return xxx_messageInfo_DelegateAction.Size(m)
}
func (m *DelegateAction) XXX_DiscardUnknown() {
	xxx_messageInfo_DelegateAction.DiscardUnknown(m)
}

var xxx_messageInfo_DelegateAction proto.InternalMessageInfo

type isDelegateAction_DelegationType interface {
	isDelegateAction_DelegationType()
	Equal(interface{}) bool
}

type DelegateAction_Ref struct {
	Ref *core.ResourceRef `protobuf:"bytes,3,opt,name=ref,proto3,oneof" json:"ref,omitempty"`
}
type DelegateAction_Selector struct {
	Selector *RouteTableSelector `protobuf:"bytes,4,opt,name=selector,proto3,oneof" json:"selector,omitempty"`
}

func (*DelegateAction_Ref) isDelegateAction_DelegationType()      {}
func (*DelegateAction_Selector) isDelegateAction_DelegationType() {}

func (m *DelegateAction) GetDelegationType() isDelegateAction_DelegationType {
	if m != nil {
		return m.DelegationType
	}
	return nil
}

// Deprecated: Do not use.
func (m *DelegateAction) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Deprecated: Do not use.
func (m *DelegateAction) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *DelegateAction) GetRef() *core.ResourceRef {
	if x, ok := m.GetDelegationType().(*DelegateAction_Ref); ok {
		return x.Ref
	}
	return nil
}

func (m *DelegateAction) GetSelector() *RouteTableSelector {
	if x, ok := m.GetDelegationType().(*DelegateAction_Selector); ok {
		return x.Selector
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*DelegateAction) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*DelegateAction_Ref)(nil),
		(*DelegateAction_Selector)(nil),
	}
}

// Select route tables for delegation by namespace, labels, or both.
type RouteTableSelector struct {
	// Delegate to Route Tables in these namespaces. If omitted, Gloo will only select Route Tables in the same namespace
	// as the resource (Virtual Service or Route Table) that owns this selector. The reserved value "*" can be used to
	// select Route Tables in all namespaces watched by Gloo.
	Namespaces []string `protobuf:"bytes,1,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	// Delegate to Route Tables whose labels match the ones specified here.
	Labels               map[string]string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RouteTableSelector) Reset()         { *m = RouteTableSelector{} }
func (m *RouteTableSelector) String() string { return proto.CompactTextString(m) }
func (*RouteTableSelector) ProtoMessage()    {}
func (*RouteTableSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_93fa9472926a2049, []int{4}
}
func (m *RouteTableSelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteTableSelector.Unmarshal(m, b)
}
func (m *RouteTableSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteTableSelector.Marshal(b, m, deterministic)
}
func (m *RouteTableSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteTableSelector.Merge(m, src)
}
func (m *RouteTableSelector) XXX_Size() int {
	return xxx_messageInfo_RouteTableSelector.Size(m)
}
func (m *RouteTableSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteTableSelector.DiscardUnknown(m)
}

var xxx_messageInfo_RouteTableSelector proto.InternalMessageInfo

func (m *RouteTableSelector) GetNamespaces() []string {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

func (m *RouteTableSelector) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func init() {
	proto.RegisterType((*VirtualService)(nil), "gateway.solo.io.VirtualService")
	proto.RegisterType((*VirtualHost)(nil), "gateway.solo.io.VirtualHost")
	proto.RegisterType((*Route)(nil), "gateway.solo.io.Route")
	proto.RegisterType((*DelegateAction)(nil), "gateway.solo.io.DelegateAction")
	proto.RegisterType((*RouteTableSelector)(nil), "gateway.solo.io.RouteTableSelector")
	proto.RegisterMapType((map[string]string)(nil), "gateway.solo.io.RouteTableSelector.LabelsEntry")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gateway/api/v1/virtual_service.proto", fileDescriptor_93fa9472926a2049)
}

var fileDescriptor_93fa9472926a2049 = []byte{
	// 822 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xc1, 0x6e, 0xdb, 0x46,
	0x10, 0x35, 0x45, 0x5a, 0x96, 0x46, 0x86, 0x9d, 0x2c, 0x04, 0x97, 0x16, 0x52, 0x59, 0xa0, 0x51,
	0xc4, 0x97, 0x90, 0x68, 0x52, 0x04, 0xa9, 0x80, 0x36, 0x88, 0x9a, 0x20, 0x42, 0xdb, 0xb4, 0xc0,
	0xba, 0xe8, 0x21, 0x17, 0x61, 0x45, 0xae, 0x68, 0xd6, 0x14, 0x97, 0xd8, 0x5d, 0xa9, 0xd6, 0x35,
	0xbf, 0xd0, 0x9f, 0xe8, 0x27, 0xa4, 0x7f, 0xd0, 0x5b, 0xcf, 0xbd, 0xe4, 0xd0, 0x3f, 0x48, 0x81,
	0xde, 0x0b, 0xee, 0x2e, 0x29, 0x51, 0xb1, 0x80, 0x9c, 0xb4, 0x3b, 0xef, 0xcd, 0x9b, 0xe1, 0x9b,
	0x21, 0x05, 0x2f, 0xe2, 0x44, 0x5e, 0x2d, 0xa6, 0x7e, 0xc8, 0xe6, 0x81, 0x60, 0x29, 0x7b, 0x90,
	0xb0, 0x20, 0x4e, 0x19, 0x0b, 0x72, 0xce, 0x7e, 0xa1, 0xa1, 0x14, 0x41, 0x4c, 0x24, 0xfd, 0x95,
	0xac, 0x02, 0x92, 0x27, 0xc1, 0xf2, 0xf3, 0x60, 0x99, 0x70, 0xb9, 0x20, 0xe9, 0x44, 0x50, 0xbe,
	0x4c, 0x42, 0xea, 0xe7, 0x9c, 0x49, 0x86, 0x8e, 0x0d, 0xcb, 0x2f, 0x34, 0xfc, 0x84, 0xf5, 0xba,
	0x31, 0x8b, 0x99, 0xc2, 0x82, 0xe2, 0xa4, 0x69, 0x3d, 0x44, 0x6f, 0xa4, 0x0e, 0xd2, 0x1b, 0x69,
	0x62, 0x7d, 0x55, 0xf6, 0x3a, 0x91, 0x65, 0x85, 0x39, 0x95, 0x24, 0x22, 0x92, 0x18, 0xfc, 0xde,
	0x36, 0x2e, 0x24, 0x91, 0x0b, 0x61, 0xd0, 0xd3, 0x6d, 0x94, 0xd3, 0xd9, 0x2e, 0xe1, 0xf2, 0x6e,
	0xf0, 0xf3, 0xad, 0xe7, 0x2c, 0x6e, 0x25, 0x53, 0xa4, 0x86, 0xf4, 0xd9, 0x6e, 0x52, 0xce, 0xd9,
	0xcd, 0xca, 0xd0, 0xee, 0xef, 0xa6, 0xb1, 0x5c, 0x26, 0x2c, 0x2b, 0xfb, 0x7d, 0xbc, 0x9b, 0x18,
	0x32, 0x4e, 0x83, 0x39, 0x91, 0xe1, 0x15, 0xe5, 0xa2, 0x3a, 0xe8, 0x3c, 0xef, 0xef, 0x06, 0x1c,
	0xfd, 0xac, 0xad, 0xbf, 0xd4, 0xce, 0xa3, 0xa7, 0x70, 0x58, 0x0e, 0xe3, 0x8a, 0x09, 0xe9, 0x5a,
	0x03, 0xeb, 0xa2, 0xf3, 0xf0, 0x9e, 0xbf, 0x35, 0x0a, 0xdf, 0xa4, 0x8d, 0x99, 0x90, 0xb8, 0xb3,
	0x5c, 0x5f, 0xd0, 0x63, 0x00, 0x21, 0xd2, 0x49, 0xc8, 0xb2, 0x59, 0x12, 0xbb, 0x0d, 0x95, 0xfe,
	0x89, 0x5f, 0xb4, 0x54, 0xe5, 0x5e, 0x8a, 0xf4, 0x1b, 0x05, 0xe3, 0xb6, 0x28, 0x8f, 0xe8, 0x3e,
	0x1c, 0x46, 0x89, 0xc8, 0x53, 0xb2, 0x9a, 0x64, 0x64, 0x4e, 0x5d, 0x7b, 0x60, 0x5d, 0xb4, 0x47,
	0xce, 0xdb, 0xff, 0x1c, 0x0b, 0x77, 0x0c, 0xf2, 0x03, 0x99, 0x53, 0xf4, 0x1d, 0x34, 0xf5, 0xb0,
	0xdc, 0xa6, 0x12, 0xef, 0xfa, 0xc5, 0x33, 0xae, 0xc5, 0x15, 0x36, 0xfa, 0xb4, 0x48, 0xfc, 0xf3,
	0xdd, 0xd9, 0xde, 0xbf, 0xef, 0xce, 0xee, 0x4a, 0x2a, 0x64, 0x94, 0xcc, 0x66, 0x43, 0x2f, 0x89,
	0x33, 0xc6, 0xa9, 0x87, 0x8d, 0x04, 0x7a, 0x02, 0xad, 0x72, 0x33, 0xdc, 0x03, 0x25, 0x77, 0x52,
	0x97, 0x7b, 0x65, 0xd0, 0x91, 0x53, 0x88, 0xe1, 0x8a, 0x3d, 0xec, 0xbf, 0x79, 0xef, 0x38, 0xd0,
	0x58, 0x8a, 0x37, 0xef, 0x1d, 0x84, 0xee, 0x6c, 0x6d, 0xb0, 0xf0, 0x7e, 0xb3, 0xa0, 0xb3, 0x61,
	0x12, 0x72, 0xe1, 0x20, 0x62, 0x73, 0x92, 0x64, 0xc2, 0x6d, 0x0c, 0xec, 0x8b, 0x36, 0x2e, 0xaf,
	0xc8, 0x87, 0x26, 0x67, 0x0b, 0x49, 0x85, 0x6b, 0x0f, 0x6c, 0xd5, 0xc1, 0xb6, 0xd9, 0xb8, 0x80,
	0xb1, 0x61, 0xa1, 0x21, 0x1c, 0x98, 0xf1, 0xbb, 0x8e, 0x6a, 0x79, 0x50, 0xb7, 0x77, 0xa3, 0xea,
	0x8f, 0x9a, 0x87, 0xcb, 0x04, 0xef, 0x0f, 0x1b, 0xf6, 0x95, 0x1a, 0x7a, 0x0a, 0xad, 0x72, 0x1b,
	0x5c, 0x4b, 0xd5, 0x3d, 0xf7, 0xab, 0xf5, 0x50, 0x16, 0xd4, 0x44, 0x5f, 0x69, 0x08, 0x57, 0x49,
	0xe8, 0x6b, 0x38, 0x54, 0x0d, 0x4d, 0x48, 0x58, 0x68, 0x9b, 0x51, 0x9f, 0xd6, 0xd3, 0x54, 0xad,
	0x67, 0x8a, 0x30, 0xde, 0xc3, 0x1d, 0xbe, 0xbe, 0xa2, 0x97, 0x70, 0xcc, 0x69, 0x94, 0x70, 0x1a,
	0xca, 0x52, 0xc2, 0x2e, 0x97, 0xad, 0x26, 0x61, 0x48, 0x95, 0xca, 0x11, 0xaf, 0x45, 0xd0, 0x6b,
	0x38, 0x31, 0x32, 0x9c, 0x8a, 0x9c, 0x65, 0xa2, 0x6a, 0x49, 0xdb, 0xe3, 0xd5, 0xf5, 0x9e, 0x2b,
	0x2e, 0x36, 0xd4, 0x4a, 0xb5, 0x1b, 0xdd, 0x12, 0x47, 0xdf, 0xc2, 0x71, 0x44, 0x53, 0x5a, 0x0c,
	0xa4, 0x14, 0xdd, 0x57, 0xa2, 0x67, 0x1f, 0x0c, 0xe9, 0xb9, 0xe1, 0xad, 0xfb, 0x8c, 0x6a, 0x11,
	0xf4, 0xc5, 0x7a, 0x6e, 0x7a, 0x73, 0x7b, 0xb7, 0x78, 0xb5, 0x3d, 0xb1, 0x51, 0x0b, 0x9a, 0xba,
	0xb0, 0xf7, 0x97, 0x05, 0x47, 0xf5, 0x22, 0xe8, 0x04, 0x1c, 0xf5, 0xb2, 0x58, 0xea, 0x65, 0x69,
	0xb8, 0x16, 0x56, 0x77, 0x34, 0x80, 0x76, 0xf1, 0x2b, 0x72, 0x12, 0x52, 0x35, 0x18, 0x0d, 0xae,
	0x83, 0xe8, 0x01, 0xd8, 0x9c, 0xce, 0x8c, 0xe3, 0xa7, 0xf5, 0x9d, 0xc7, 0x54, 0xb0, 0x05, 0x0f,
	0x29, 0xa6, 0xb3, 0xf1, 0x1e, 0x2e, 0x78, 0xe8, 0x19, 0xb4, 0x04, 0x4d, 0x69, 0x28, 0x19, 0x37,
	0xae, 0x9e, 0xdf, 0xbe, 0xa5, 0x3f, 0x91, 0x69, 0x4a, 0x2f, 0x0d, 0x75, 0xbc, 0x87, 0xab, 0xb4,
	0xd1, 0xdd, 0xca, 0xca, 0x84, 0x65, 0x13, 0xb9, 0xca, 0xa9, 0xf7, 0xd6, 0x02, 0xf4, 0x61, 0x16,
	0xea, 0x03, 0x54, 0x8d, 0xea, 0xe5, 0x6c, 0xe3, 0x8d, 0x08, 0x7a, 0x09, 0xcd, 0x94, 0x4c, 0x69,
	0xaa, 0xdf, 0xa4, 0xce, 0xc3, 0xe0, 0x23, 0x5a, 0xf1, 0xbf, 0x57, 0x19, 0x2f, 0x32, 0xc9, 0x57,
	0xd8, 0xa4, 0xf7, 0xbe, 0x84, 0xce, 0x46, 0x18, 0xdd, 0x01, 0xfb, 0x9a, 0xae, 0xb4, 0x99, 0xb8,
	0x38, 0xa2, 0x2e, 0xec, 0x2f, 0x49, 0xba, 0x30, 0x1e, 0x62, 0x7d, 0x19, 0x36, 0x9e, 0x58, 0xa3,
	0xaf, 0x8a, 0xef, 0xcb, 0xef, 0xff, 0xf4, 0xad, 0xd7, 0x8f, 0x3e, 0xfa, 0xcf, 0x2e, 0xbf, 0x8e,
	0xcd, 0x67, 0x79, 0xda, 0x54, 0x1f, 0xe0, 0x47, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0x24, 0x29,
	0x5d, 0x6d, 0x2a, 0x07, 0x00, 0x00,
}

func (this *VirtualService) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualService)
	if !ok {
		that2, ok := that.(VirtualService)
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
	if !this.VirtualHost.Equal(that1.VirtualHost) {
		return false
	}
	if !this.SslConfig.Equal(that1.SslConfig) {
		return false
	}
	if this.DisplayName != that1.DisplayName {
		return false
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *VirtualHost) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualHost)
	if !ok {
		that2, ok := that.(VirtualHost)
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
	if len(this.Domains) != len(that1.Domains) {
		return false
	}
	for i := range this.Domains {
		if this.Domains[i] != that1.Domains[i] {
			return false
		}
	}
	if len(this.Routes) != len(that1.Routes) {
		return false
	}
	for i := range this.Routes {
		if !this.Routes[i].Equal(that1.Routes[i]) {
			return false
		}
	}
	if !this.Options.Equal(that1.Options) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Route) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Route)
	if !ok {
		that2, ok := that.(Route)
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
	if len(this.Matchers) != len(that1.Matchers) {
		return false
	}
	for i := range this.Matchers {
		if !this.Matchers[i].Equal(that1.Matchers[i]) {
			return false
		}
	}
	if that1.Action == nil {
		if this.Action != nil {
			return false
		}
	} else if this.Action == nil {
		return false
	} else if !this.Action.Equal(that1.Action) {
		return false
	}
	if !this.Options.Equal(that1.Options) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Route_RouteAction) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Route_RouteAction)
	if !ok {
		that2, ok := that.(Route_RouteAction)
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
	if !this.RouteAction.Equal(that1.RouteAction) {
		return false
	}
	return true
}
func (this *Route_RedirectAction) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Route_RedirectAction)
	if !ok {
		that2, ok := that.(Route_RedirectAction)
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
	if !this.RedirectAction.Equal(that1.RedirectAction) {
		return false
	}
	return true
}
func (this *Route_DirectResponseAction) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Route_DirectResponseAction)
	if !ok {
		that2, ok := that.(Route_DirectResponseAction)
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
	if !this.DirectResponseAction.Equal(that1.DirectResponseAction) {
		return false
	}
	return true
}
func (this *Route_DelegateAction) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Route_DelegateAction)
	if !ok {
		that2, ok := that.(Route_DelegateAction)
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
	if !this.DelegateAction.Equal(that1.DelegateAction) {
		return false
	}
	return true
}
func (this *DelegateAction) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DelegateAction)
	if !ok {
		that2, ok := that.(DelegateAction)
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
	if this.Name != that1.Name {
		return false
	}
	if this.Namespace != that1.Namespace {
		return false
	}
	if that1.DelegationType == nil {
		if this.DelegationType != nil {
			return false
		}
	} else if this.DelegationType == nil {
		return false
	} else if !this.DelegationType.Equal(that1.DelegationType) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *DelegateAction_Ref) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DelegateAction_Ref)
	if !ok {
		that2, ok := that.(DelegateAction_Ref)
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
	return true
}
func (this *DelegateAction_Selector) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DelegateAction_Selector)
	if !ok {
		that2, ok := that.(DelegateAction_Selector)
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
	if !this.Selector.Equal(that1.Selector) {
		return false
	}
	return true
}
func (this *RouteTableSelector) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RouteTableSelector)
	if !ok {
		that2, ok := that.(RouteTableSelector)
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
	if len(this.Namespaces) != len(that1.Namespaces) {
		return false
	}
	for i := range this.Namespaces {
		if this.Namespaces[i] != that1.Namespaces[i] {
			return false
		}
	}
	if len(this.Labels) != len(that1.Labels) {
		return false
	}
	for i := range this.Labels {
		if this.Labels[i] != that1.Labels[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
