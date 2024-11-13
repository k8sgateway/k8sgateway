// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gateway/api/v1/matchable_http_gateway.proto

package v1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_config_core_v3 "github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/config/core/v3"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_ssl "github.com/solo-io/gloo/projects/controllerrollerroller/pkg/api/v1/ssl"

	github_com_solo_io_solo_kit_pkg_api_v1_resources_core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
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
func (m *MatchableHttpGateway) Clone() proto.Message {
	var target *MatchableHttpGateway
	if m == nil {
		return target
	}
	target = &MatchableHttpGateway{}

	if h, ok := interface{}(m.GetNamespacedStatuses()).(clone.Cloner); ok {
		target.NamespacedStatuses = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.NamespacedStatuses)
	} else {
		target.NamespacedStatuses = proto.Clone(m.GetNamespacedStatuses()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.NamespacedStatuses)
	}

	if h, ok := interface{}(m.GetMetadata()).(clone.Cloner); ok {
		target.Metadata = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.Metadata)
	} else {
		target.Metadata = proto.Clone(m.GetMetadata()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.Metadata)
	}

	if h, ok := interface{}(m.GetMatcher()).(clone.Cloner); ok {
		target.Matcher = h.Clone().(*MatchableHttpGateway_Matcher)
	} else {
		target.Matcher = proto.Clone(m.GetMatcher()).(*MatchableHttpGateway_Matcher)
	}

	if h, ok := interface{}(m.GetHttpGateway()).(clone.Cloner); ok {
		target.HttpGateway = h.Clone().(*HttpGateway)
	} else {
		target.HttpGateway = proto.Clone(m.GetHttpGateway()).(*HttpGateway)
	}

	return target
}

// Clone function
func (m *MatchableHttpGateway_Matcher) Clone() proto.Message {
	var target *MatchableHttpGateway_Matcher
	if m == nil {
		return target
	}
	target = &MatchableHttpGateway_Matcher{}

	if m.GetSourcePrefixRanges() != nil {
		target.SourcePrefixRanges = make([]*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_config_core_v3.CidrRange, len(m.GetSourcePrefixRanges()))
		for idx, v := range m.GetSourcePrefixRanges() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.SourcePrefixRanges[idx] = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_config_core_v3.CidrRange)
			} else {
				target.SourcePrefixRanges[idx] = proto.Clone(v).(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_config_core_v3.CidrRange)
			}

		}
	}

	if h, ok := interface{}(m.GetSslConfig()).(clone.Cloner); ok {
		target.SslConfig = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_ssl.SslConfig)
	} else {
		target.SslConfig = proto.Clone(m.GetSslConfig()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_ssl.SslConfig)
	}

	return target
}
