// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/upstream.proto

package v1

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

	github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_api_v2_cluster "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/api/v2/cluster"

	github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_api_v2_core "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/api/v2/core"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_aws "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/aws"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_aws_ec2 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/aws/ec2"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_azure "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/azure"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_consul "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/consul"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_kubernetes "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/kubernetes"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_pipe "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/pipe"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_static "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/static"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_ssl "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/ssl"

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
func (m *Upstream) Clone() proto.Message {
	var target *Upstream
	if m == nil {
		return target
	}
	target = &Upstream{}

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

	if h, ok := interface{}(m.GetDiscoveryMetadata()).(clone.Cloner); ok {
		target.DiscoveryMetadata = h.Clone().(*DiscoveryMetadata)
	} else {
		target.DiscoveryMetadata = proto.Clone(m.GetDiscoveryMetadata()).(*DiscoveryMetadata)
	}

	if h, ok := interface{}(m.GetSslConfig()).(clone.Cloner); ok {
		target.SslConfig = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_ssl.UpstreamSslConfig)
	} else {
		target.SslConfig = proto.Clone(m.GetSslConfig()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_ssl.UpstreamSslConfig)
	}

	if h, ok := interface{}(m.GetCircuitBreakers()).(clone.Cloner); ok {
		target.CircuitBreakers = h.Clone().(*CircuitBreakerConfig)
	} else {
		target.CircuitBreakers = proto.Clone(m.GetCircuitBreakers()).(*CircuitBreakerConfig)
	}

	if h, ok := interface{}(m.GetLoadBalancerConfig()).(clone.Cloner); ok {
		target.LoadBalancerConfig = h.Clone().(*LoadBalancerConfig)
	} else {
		target.LoadBalancerConfig = proto.Clone(m.GetLoadBalancerConfig()).(*LoadBalancerConfig)
	}

	if m.GetHealthChecks() != nil {
		target.HealthChecks = make([]*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_api_v2_core.HealthCheck, len(m.GetHealthChecks()))
		for idx, v := range m.GetHealthChecks() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.HealthChecks[idx] = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_api_v2_core.HealthCheck)
			} else {
				target.HealthChecks[idx] = proto.Clone(v).(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_api_v2_core.HealthCheck)
			}

		}
	}

	if h, ok := interface{}(m.GetOutlierDetection()).(clone.Cloner); ok {
		target.OutlierDetection = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_api_v2_cluster.OutlierDetection)
	} else {
		target.OutlierDetection = proto.Clone(m.GetOutlierDetection()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_api_v2_cluster.OutlierDetection)
	}

	if h, ok := interface{}(m.GetFailover()).(clone.Cloner); ok {
		target.Failover = h.Clone().(*Failover)
	} else {
		target.Failover = proto.Clone(m.GetFailover()).(*Failover)
	}

	if h, ok := interface{}(m.GetConnectionConfig()).(clone.Cloner); ok {
		target.ConnectionConfig = h.Clone().(*ConnectionConfig)
	} else {
		target.ConnectionConfig = proto.Clone(m.GetConnectionConfig()).(*ConnectionConfig)
	}

	target.ProtocolSelection = m.GetProtocolSelection()

	if h, ok := interface{}(m.GetUseHttp2()).(clone.Cloner); ok {
		target.UseHttp2 = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.UseHttp2 = proto.Clone(m.GetUseHttp2()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetInitialStreamWindowSize()).(clone.Cloner); ok {
		target.InitialStreamWindowSize = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	} else {
		target.InitialStreamWindowSize = proto.Clone(m.GetInitialStreamWindowSize()).(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	}

	if h, ok := interface{}(m.GetInitialConnectionWindowSize()).(clone.Cloner); ok {
		target.InitialConnectionWindowSize = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	} else {
		target.InitialConnectionWindowSize = proto.Clone(m.GetInitialConnectionWindowSize()).(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	}

	if h, ok := interface{}(m.GetMaxConcurrentStreams()).(clone.Cloner); ok {
		target.MaxConcurrentStreams = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	} else {
		target.MaxConcurrentStreams = proto.Clone(m.GetMaxConcurrentStreams()).(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	}

	if h, ok := interface{}(m.GetOverrideStreamErrorOnInvalidHttpMessage()).(clone.Cloner); ok {
		target.OverrideStreamErrorOnInvalidHttpMessage = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.OverrideStreamErrorOnInvalidHttpMessage = proto.Clone(m.GetOverrideStreamErrorOnInvalidHttpMessage()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetHttpProxyHostname()).(clone.Cloner); ok {
		target.HttpProxyHostname = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.StringValue)
	} else {
		target.HttpProxyHostname = proto.Clone(m.GetHttpProxyHostname()).(*github_com_golang_protobuf_ptypes_wrappers.StringValue)
	}

	if h, ok := interface{}(m.GetHttpConnectSslConfig()).(clone.Cloner); ok {
		target.HttpConnectSslConfig = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_ssl.UpstreamSslConfig)
	} else {
		target.HttpConnectSslConfig = proto.Clone(m.GetHttpConnectSslConfig()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_ssl.UpstreamSslConfig)
	}

	if m.GetHttpConnectHeaders() != nil {
		target.HttpConnectHeaders = make([]*HeaderValue, len(m.GetHttpConnectHeaders()))
		for idx, v := range m.GetHttpConnectHeaders() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.HttpConnectHeaders[idx] = h.Clone().(*HeaderValue)
			} else {
				target.HttpConnectHeaders[idx] = proto.Clone(v).(*HeaderValue)
			}

		}
	}

	if h, ok := interface{}(m.GetIgnoreHealthOnHostRemoval()).(clone.Cloner); ok {
		target.IgnoreHealthOnHostRemoval = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.IgnoreHealthOnHostRemoval = proto.Clone(m.GetIgnoreHealthOnHostRemoval()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetRespectDnsTtl()).(clone.Cloner); ok {
		target.RespectDnsTtl = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.RespectDnsTtl = proto.Clone(m.GetRespectDnsTtl()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetDnsRefreshRate()).(clone.Cloner); ok {
		target.DnsRefreshRate = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.DnsRefreshRate = proto.Clone(m.GetDnsRefreshRate()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	if h, ok := interface{}(m.GetProxyProtocolVersion()).(clone.Cloner); ok {
		target.ProxyProtocolVersion = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.Int32Value)
	} else {
		target.ProxyProtocolVersion = proto.Clone(m.GetProxyProtocolVersion()).(*github_com_golang_protobuf_ptypes_wrappers.Int32Value)
	}

	switch m.UpstreamType.(type) {

	case *Upstream_Kube:

		if h, ok := interface{}(m.GetKube()).(clone.Cloner); ok {
			target.UpstreamType = &Upstream_Kube{
				Kube: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_kubernetes.UpstreamSpec),
			}
		} else {
			target.UpstreamType = &Upstream_Kube{
				Kube: proto.Clone(m.GetKube()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_kubernetes.UpstreamSpec),
			}
		}

	case *Upstream_Static:

		if h, ok := interface{}(m.GetStatic()).(clone.Cloner); ok {
			target.UpstreamType = &Upstream_Static{
				Static: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_static.UpstreamSpec),
			}
		} else {
			target.UpstreamType = &Upstream_Static{
				Static: proto.Clone(m.GetStatic()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_static.UpstreamSpec),
			}
		}

	case *Upstream_Pipe:

		if h, ok := interface{}(m.GetPipe()).(clone.Cloner); ok {
			target.UpstreamType = &Upstream_Pipe{
				Pipe: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_pipe.UpstreamSpec),
			}
		} else {
			target.UpstreamType = &Upstream_Pipe{
				Pipe: proto.Clone(m.GetPipe()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_pipe.UpstreamSpec),
			}
		}

	case *Upstream_Aws:

		if h, ok := interface{}(m.GetAws()).(clone.Cloner); ok {
			target.UpstreamType = &Upstream_Aws{
				Aws: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_aws.UpstreamSpec),
			}
		} else {
			target.UpstreamType = &Upstream_Aws{
				Aws: proto.Clone(m.GetAws()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_aws.UpstreamSpec),
			}
		}

	case *Upstream_Azure:

		if h, ok := interface{}(m.GetAzure()).(clone.Cloner); ok {
			target.UpstreamType = &Upstream_Azure{
				Azure: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_azure.UpstreamSpec),
			}
		} else {
			target.UpstreamType = &Upstream_Azure{
				Azure: proto.Clone(m.GetAzure()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_azure.UpstreamSpec),
			}
		}

	case *Upstream_Consul:

		if h, ok := interface{}(m.GetConsul()).(clone.Cloner); ok {
			target.UpstreamType = &Upstream_Consul{
				Consul: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_consul.UpstreamSpec),
			}
		} else {
			target.UpstreamType = &Upstream_Consul{
				Consul: proto.Clone(m.GetConsul()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_consul.UpstreamSpec),
			}
		}

	case *Upstream_AwsEc2:

		if h, ok := interface{}(m.GetAwsEc2()).(clone.Cloner); ok {
			target.UpstreamType = &Upstream_AwsEc2{
				AwsEc2: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_aws_ec2.UpstreamSpec),
			}
		} else {
			target.UpstreamType = &Upstream_AwsEc2{
				AwsEc2: proto.Clone(m.GetAwsEc2()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_aws_ec2.UpstreamSpec),
			}
		}

	}

	return target
}

// Clone function
func (m *DiscoveryMetadata) Clone() proto.Message {
	var target *DiscoveryMetadata
	if m == nil {
		return target
	}
	target = &DiscoveryMetadata{}

	if m.GetLabels() != nil {
		target.Labels = make(map[string]string, len(m.GetLabels()))
		for k, v := range m.GetLabels() {

			target.Labels[k] = v

		}
	}

	return target
}

// Clone function
func (m *HeaderValue) Clone() proto.Message {
	var target *HeaderValue
	if m == nil {
		return target
	}
	target = &HeaderValue{}

	target.Key = m.GetKey()

	target.Value = m.GetValue()

	return target
}
