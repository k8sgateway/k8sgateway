// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/dynamic_forward_proxy/dynamic_forward_proxy.proto

package dynamic_forward_proxy

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_golang_protobuf_ptypes_duration "github.com/golang/protobuf/ptypes/duration"

	github_com_golang_protobuf_ptypes_wrappers "github.com/golang/protobuf/ptypes/wrappers"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = clone.Cloner(nil)
	_ = proto.Message(nil)
)

// Clone function
func (m *FilterConfig) Clone() proto.Message {
	var target *FilterConfig
	if m == nil {
		return target
	}
	target = &FilterConfig{}

	if h, ok := interface{}(m.GetDnsCacheConfig()).(clone.Cloner); ok {
		target.DnsCacheConfig = h.Clone().(*DnsCacheConfig)
	} else {
		target.DnsCacheConfig = proto.Clone(m.GetDnsCacheConfig()).(*DnsCacheConfig)
	}

	return target
}

// Clone function
func (m *DnsCacheCircuitBreakers) Clone() proto.Message {
	var target *DnsCacheCircuitBreakers
	if m == nil {
		return target
	}
	target = &DnsCacheCircuitBreakers{}

	if h, ok := interface{}(m.GetMaxPendingRequests()).(clone.Cloner); ok {
		target.MaxPendingRequests = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	} else {
		target.MaxPendingRequests = proto.Clone(m.GetMaxPendingRequests()).(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	}

	return target
}

// Clone function
func (m *DnsCacheConfig) Clone() proto.Message {
	var target *DnsCacheConfig
	if m == nil {
		return target
	}
	target = &DnsCacheConfig{}

	target.Name = m.GetName()

	if h, ok := interface{}(m.GetDnsRefreshRate()).(clone.Cloner); ok {
		target.DnsRefreshRate = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.DnsRefreshRate = proto.Clone(m.GetDnsRefreshRate()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	if h, ok := interface{}(m.GetHostTtl()).(clone.Cloner); ok {
		target.HostTtl = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.HostTtl = proto.Clone(m.GetHostTtl()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	if h, ok := interface{}(m.GetMaxHosts()).(clone.Cloner); ok {
		target.MaxHosts = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	} else {
		target.MaxHosts = proto.Clone(m.GetMaxHosts()).(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	}

	if h, ok := interface{}(m.GetDnsCacheCircuitBreaker()).(clone.Cloner); ok {
		target.DnsCacheCircuitBreaker = h.Clone().(*DnsCacheCircuitBreakers)
	} else {
		target.DnsCacheCircuitBreaker = proto.Clone(m.GetDnsCacheCircuitBreaker()).(*DnsCacheCircuitBreakers)
	}

	if h, ok := interface{}(m.GetDnsQueryTimeout()).(clone.Cloner); ok {
		target.DnsQueryTimeout = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.DnsQueryTimeout = proto.Clone(m.GetDnsQueryTimeout()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	return target
}

// Clone function
func (m *PerRouteConfig) Clone() proto.Message {
	var target *PerRouteConfig
	if m == nil {
		return target
	}
	target = &PerRouteConfig{}

	switch m.HostRewriteSpecifier.(type) {

	case *PerRouteConfig_HostRewrite:

		target.HostRewriteSpecifier = &PerRouteConfig_HostRewrite{
			HostRewrite: m.GetHostRewrite(),
		}

	case *PerRouteConfig_AutoHostRewriteHeader:

		target.HostRewriteSpecifier = &PerRouteConfig_AutoHostRewriteHeader{
			AutoHostRewriteHeader: m.GetAutoHostRewriteHeader(),
		}

	}

	return target
}
