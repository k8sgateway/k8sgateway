// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/hcm/hcm.proto

package hcm

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

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_protocol_upgrade "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/protocol_upgrade"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_tracing "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/tracing"
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
func (m *HttpConnectionManagerSettings) Clone() proto.Message {
	var target *HttpConnectionManagerSettings
	if m == nil {
		return target
	}
	target = &HttpConnectionManagerSettings{}

	target.SkipXffAppend = m.GetSkipXffAppend()

	target.Via = m.GetVia()

	target.XffNumTrustedHops = m.GetXffNumTrustedHops()

	if h, ok := interface{}(m.GetUseRemoteAddress()).(clone.Cloner); ok {
		target.UseRemoteAddress = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.UseRemoteAddress = proto.Clone(m.GetUseRemoteAddress()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetGenerateRequestId()).(clone.Cloner); ok {
		target.GenerateRequestId = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.GenerateRequestId = proto.Clone(m.GetGenerateRequestId()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	target.Proxy_100Continue = m.GetProxy_100Continue()

	if h, ok := interface{}(m.GetStreamIdleTimeout()).(clone.Cloner); ok {
		target.StreamIdleTimeout = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.StreamIdleTimeout = proto.Clone(m.GetStreamIdleTimeout()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	if h, ok := interface{}(m.GetIdleTimeout()).(clone.Cloner); ok {
		target.IdleTimeout = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.IdleTimeout = proto.Clone(m.GetIdleTimeout()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	if h, ok := interface{}(m.GetMaxRequestHeadersKb()).(clone.Cloner); ok {
		target.MaxRequestHeadersKb = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	} else {
		target.MaxRequestHeadersKb = proto.Clone(m.GetMaxRequestHeadersKb()).(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	}

	if h, ok := interface{}(m.GetRequestTimeout()).(clone.Cloner); ok {
		target.RequestTimeout = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.RequestTimeout = proto.Clone(m.GetRequestTimeout()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	if h, ok := interface{}(m.GetDrainTimeout()).(clone.Cloner); ok {
		target.DrainTimeout = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.DrainTimeout = proto.Clone(m.GetDrainTimeout()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	if h, ok := interface{}(m.GetDelayedCloseTimeout()).(clone.Cloner); ok {
		target.DelayedCloseTimeout = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.DelayedCloseTimeout = proto.Clone(m.GetDelayedCloseTimeout()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	target.ServerName = m.GetServerName()

	target.AcceptHttp_10 = m.GetAcceptHttp_10()

	target.DefaultHostForHttp_10 = m.GetDefaultHostForHttp_10()

	target.AllowChunkedLength = m.GetAllowChunkedLength()

	target.EnableTrailers = m.GetEnableTrailers()

	if h, ok := interface{}(m.GetTracing()).(clone.Cloner); ok {
		target.Tracing = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_tracing.ListenerTracingSettings)
	} else {
		target.Tracing = proto.Clone(m.GetTracing()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_tracing.ListenerTracingSettings)
	}

	target.ForwardClientCertDetails = m.GetForwardClientCertDetails()

	if h, ok := interface{}(m.GetSetCurrentClientCertDetails()).(clone.Cloner); ok {
		target.SetCurrentClientCertDetails = h.Clone().(*HttpConnectionManagerSettings_SetCurrentClientCertDetails)
	} else {
		target.SetCurrentClientCertDetails = proto.Clone(m.GetSetCurrentClientCertDetails()).(*HttpConnectionManagerSettings_SetCurrentClientCertDetails)
	}

	target.PreserveExternalRequestId = m.GetPreserveExternalRequestId()

	if m.GetUpgrades() != nil {
		target.Upgrades = make([]*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_protocol_upgrade.ProtocolUpgradeConfig, len(m.GetUpgrades()))
		for idx, v := range m.GetUpgrades() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Upgrades[idx] = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_protocol_upgrade.ProtocolUpgradeConfig)
			} else {
				target.Upgrades[idx] = proto.Clone(v).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_protocol_upgrade.ProtocolUpgradeConfig)
			}

		}
	}

	if h, ok := interface{}(m.GetMaxConnectionDuration()).(clone.Cloner); ok {
		target.MaxConnectionDuration = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.MaxConnectionDuration = proto.Clone(m.GetMaxConnectionDuration()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	if h, ok := interface{}(m.GetMaxStreamDuration()).(clone.Cloner); ok {
		target.MaxStreamDuration = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.MaxStreamDuration = proto.Clone(m.GetMaxStreamDuration()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	if h, ok := interface{}(m.GetMaxHeadersCount()).(clone.Cloner); ok {
		target.MaxHeadersCount = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	} else {
		target.MaxHeadersCount = proto.Clone(m.GetMaxHeadersCount()).(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	}

	target.HeadersWithUnderscoresAction = m.GetHeadersWithUnderscoresAction()

	if h, ok := interface{}(m.GetMaxRequestsPerConnection()).(clone.Cloner); ok {
		target.MaxRequestsPerConnection = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	} else {
		target.MaxRequestsPerConnection = proto.Clone(m.GetMaxRequestsPerConnection()).(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	}

	target.ServerHeaderTransformation = m.GetServerHeaderTransformation()

	target.PathWithEscapedSlashesAction = m.GetPathWithEscapedSlashesAction()

	target.CodecType = m.GetCodecType()

	target.MergeSlashes = m.GetMergeSlashes()

	if h, ok := interface{}(m.GetNormalizePath()).(clone.Cloner); ok {
		target.NormalizePath = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.NormalizePath = proto.Clone(m.GetNormalizePath()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	switch m.HeaderFormat.(type) {

	case *HttpConnectionManagerSettings_ProperCaseHeaderKeyFormat:

		target.HeaderFormat = &HttpConnectionManagerSettings_ProperCaseHeaderKeyFormat{
			ProperCaseHeaderKeyFormat: m.GetProperCaseHeaderKeyFormat(),
		}

	case *HttpConnectionManagerSettings_PreserveCaseHeaderKeyFormat:

		target.HeaderFormat = &HttpConnectionManagerSettings_PreserveCaseHeaderKeyFormat{
			PreserveCaseHeaderKeyFormat: m.GetPreserveCaseHeaderKeyFormat(),
		}

	}

	return target
}

// Clone function
func (m *HttpConnectionManagerSettings_SetCurrentClientCertDetails) Clone() proto.Message {
	var target *HttpConnectionManagerSettings_SetCurrentClientCertDetails
	if m == nil {
		return target
	}
	target = &HttpConnectionManagerSettings_SetCurrentClientCertDetails{}

	if h, ok := interface{}(m.GetSubject()).(clone.Cloner); ok {
		target.Subject = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.Subject = proto.Clone(m.GetSubject()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	target.Cert = m.GetCert()

	target.Chain = m.GetChain()

	target.Dns = m.GetDns()

	target.Uri = m.GetUri()

	return target
}
