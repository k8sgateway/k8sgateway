// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/settings.proto

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

	github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_aws "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/aws"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_caching "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/caching"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_extauth_v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/extauth/v1"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ratelimit"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_rbac "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/rbac"

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
func (m *Settings) Clone() proto.Message {
	var target *Settings
	if m == nil {
		return target
	}
	target = &Settings{}

	target.DiscoveryNamespace = m.GetDiscoveryNamespace()

	if m.GetWatchNamespaces() != nil {
		target.WatchNamespaces = make([]string, len(m.GetWatchNamespaces()))
		for idx, v := range m.GetWatchNamespaces() {

			target.WatchNamespaces[idx] = v

		}
	}

	if h, ok := interface{}(m.GetRefreshRate()).(clone.Cloner); ok {
		target.RefreshRate = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.RefreshRate = proto.Clone(m.GetRefreshRate()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	target.DevMode = m.GetDevMode()

	target.Linkerd = m.GetLinkerd()

	if h, ok := interface{}(m.GetKnative()).(clone.Cloner); ok {
		target.Knative = h.Clone().(*Settings_KnativeOptions)
	} else {
		target.Knative = proto.Clone(m.GetKnative()).(*Settings_KnativeOptions)
	}

	if h, ok := interface{}(m.GetDiscovery()).(clone.Cloner); ok {
		target.Discovery = h.Clone().(*Settings_DiscoveryOptions)
	} else {
		target.Discovery = proto.Clone(m.GetDiscovery()).(*Settings_DiscoveryOptions)
	}

	if h, ok := interface{}(m.GetGloo()).(clone.Cloner); ok {
		target.Gloo = h.Clone().(*GlooOptions)
	} else {
		target.Gloo = proto.Clone(m.GetGloo()).(*GlooOptions)
	}

	if h, ok := interface{}(m.GetGateway()).(clone.Cloner); ok {
		target.Gateway = h.Clone().(*GatewayOptions)
	} else {
		target.Gateway = proto.Clone(m.GetGateway()).(*GatewayOptions)
	}

	if h, ok := interface{}(m.GetConsul()).(clone.Cloner); ok {
		target.Consul = h.Clone().(*Settings_ConsulConfiguration)
	} else {
		target.Consul = proto.Clone(m.GetConsul()).(*Settings_ConsulConfiguration)
	}

	if h, ok := interface{}(m.GetConsulDiscovery()).(clone.Cloner); ok {
		target.ConsulDiscovery = h.Clone().(*Settings_ConsulUpstreamDiscoveryConfiguration)
	} else {
		target.ConsulDiscovery = proto.Clone(m.GetConsulDiscovery()).(*Settings_ConsulUpstreamDiscoveryConfiguration)
	}

	if h, ok := interface{}(m.GetKubernetes()).(clone.Cloner); ok {
		target.Kubernetes = h.Clone().(*Settings_KubernetesConfiguration)
	} else {
		target.Kubernetes = proto.Clone(m.GetKubernetes()).(*Settings_KubernetesConfiguration)
	}

	if h, ok := interface{}(m.GetExtensions()).(clone.Cloner); ok {
		target.Extensions = h.Clone().(*Extensions)
	} else {
		target.Extensions = proto.Clone(m.GetExtensions()).(*Extensions)
	}

	if h, ok := interface{}(m.GetRatelimit()).(clone.Cloner); ok {
		target.Ratelimit = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit.ServiceSettings)
	} else {
		target.Ratelimit = proto.Clone(m.GetRatelimit()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit.ServiceSettings)
	}

	if h, ok := interface{}(m.GetRatelimitServer()).(clone.Cloner); ok {
		target.RatelimitServer = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit.Settings)
	} else {
		target.RatelimitServer = proto.Clone(m.GetRatelimitServer()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit.Settings)
	}

	if h, ok := interface{}(m.GetRbac()).(clone.Cloner); ok {
		target.Rbac = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_rbac.Settings)
	} else {
		target.Rbac = proto.Clone(m.GetRbac()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_rbac.Settings)
	}

	if h, ok := interface{}(m.GetExtauth()).(clone.Cloner); ok {
		target.Extauth = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_extauth_v1.Settings)
	} else {
		target.Extauth = proto.Clone(m.GetExtauth()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_extauth_v1.Settings)
	}

	if m.GetNamedExtauth() != nil {
		target.NamedExtauth = make(map[string]*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_extauth_v1.Settings, len(m.GetNamedExtauth()))
		for k, v := range m.GetNamedExtauth() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.NamedExtauth[k] = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_extauth_v1.Settings)
			} else {
				target.NamedExtauth[k] = proto.Clone(v).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_extauth_v1.Settings)
			}

		}
	}

	if h, ok := interface{}(m.GetCachingServer()).(clone.Cloner); ok {
		target.CachingServer = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_caching.Settings)
	} else {
		target.CachingServer = proto.Clone(m.GetCachingServer()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_caching.Settings)
	}

	if h, ok := interface{}(m.GetMetadata()).(clone.Cloner); ok {
		target.Metadata = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.Metadata)
	} else {
		target.Metadata = proto.Clone(m.GetMetadata()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.Metadata)
	}

	if h, ok := interface{}(m.GetNamespacedStatuses()).(clone.Cloner); ok {
		target.NamespacedStatuses = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.NamespacedStatuses)
	} else {
		target.NamespacedStatuses = proto.Clone(m.GetNamespacedStatuses()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.NamespacedStatuses)
	}

	if h, ok := interface{}(m.GetObservabilityOptions()).(clone.Cloner); ok {
		target.ObservabilityOptions = h.Clone().(*Settings_ObservabilityOptions)
	} else {
		target.ObservabilityOptions = proto.Clone(m.GetObservabilityOptions()).(*Settings_ObservabilityOptions)
	}

	if h, ok := interface{}(m.GetUpstreamOptions()).(clone.Cloner); ok {
		target.UpstreamOptions = h.Clone().(*UpstreamOptions)
	} else {
		target.UpstreamOptions = proto.Clone(m.GetUpstreamOptions()).(*UpstreamOptions)
	}

	if h, ok := interface{}(m.GetConsoleOptions()).(clone.Cloner); ok {
		target.ConsoleOptions = h.Clone().(*ConsoleOptions)
	} else {
		target.ConsoleOptions = proto.Clone(m.GetConsoleOptions()).(*ConsoleOptions)
	}

	switch m.ConfigSource.(type) {

	case *Settings_KubernetesConfigSource:

		if h, ok := interface{}(m.GetKubernetesConfigSource()).(clone.Cloner); ok {
			target.ConfigSource = &Settings_KubernetesConfigSource{
				KubernetesConfigSource: h.Clone().(*Settings_KubernetesCrds),
			}
		} else {
			target.ConfigSource = &Settings_KubernetesConfigSource{
				KubernetesConfigSource: proto.Clone(m.GetKubernetesConfigSource()).(*Settings_KubernetesCrds),
			}
		}

	case *Settings_DirectoryConfigSource:

		if h, ok := interface{}(m.GetDirectoryConfigSource()).(clone.Cloner); ok {
			target.ConfigSource = &Settings_DirectoryConfigSource{
				DirectoryConfigSource: h.Clone().(*Settings_Directory),
			}
		} else {
			target.ConfigSource = &Settings_DirectoryConfigSource{
				DirectoryConfigSource: proto.Clone(m.GetDirectoryConfigSource()).(*Settings_Directory),
			}
		}

	case *Settings_ConsulKvSource:

		if h, ok := interface{}(m.GetConsulKvSource()).(clone.Cloner); ok {
			target.ConfigSource = &Settings_ConsulKvSource{
				ConsulKvSource: h.Clone().(*Settings_ConsulKv),
			}
		} else {
			target.ConfigSource = &Settings_ConsulKvSource{
				ConsulKvSource: proto.Clone(m.GetConsulKvSource()).(*Settings_ConsulKv),
			}
		}

	}

	switch m.SecretSource.(type) {

	case *Settings_KubernetesSecretSource:

		if h, ok := interface{}(m.GetKubernetesSecretSource()).(clone.Cloner); ok {
			target.SecretSource = &Settings_KubernetesSecretSource{
				KubernetesSecretSource: h.Clone().(*Settings_KubernetesSecrets),
			}
		} else {
			target.SecretSource = &Settings_KubernetesSecretSource{
				KubernetesSecretSource: proto.Clone(m.GetKubernetesSecretSource()).(*Settings_KubernetesSecrets),
			}
		}

	case *Settings_VaultSecretSource:

		if h, ok := interface{}(m.GetVaultSecretSource()).(clone.Cloner); ok {
			target.SecretSource = &Settings_VaultSecretSource{
				VaultSecretSource: h.Clone().(*Settings_VaultSecrets),
			}
		} else {
			target.SecretSource = &Settings_VaultSecretSource{
				VaultSecretSource: proto.Clone(m.GetVaultSecretSource()).(*Settings_VaultSecrets),
			}
		}

	case *Settings_DirectorySecretSource:

		if h, ok := interface{}(m.GetDirectorySecretSource()).(clone.Cloner); ok {
			target.SecretSource = &Settings_DirectorySecretSource{
				DirectorySecretSource: h.Clone().(*Settings_Directory),
			}
		} else {
			target.SecretSource = &Settings_DirectorySecretSource{
				DirectorySecretSource: proto.Clone(m.GetDirectorySecretSource()).(*Settings_Directory),
			}
		}

	}

	switch m.ArtifactSource.(type) {

	case *Settings_KubernetesArtifactSource:

		if h, ok := interface{}(m.GetKubernetesArtifactSource()).(clone.Cloner); ok {
			target.ArtifactSource = &Settings_KubernetesArtifactSource{
				KubernetesArtifactSource: h.Clone().(*Settings_KubernetesConfigmaps),
			}
		} else {
			target.ArtifactSource = &Settings_KubernetesArtifactSource{
				KubernetesArtifactSource: proto.Clone(m.GetKubernetesArtifactSource()).(*Settings_KubernetesConfigmaps),
			}
		}

	case *Settings_DirectoryArtifactSource:

		if h, ok := interface{}(m.GetDirectoryArtifactSource()).(clone.Cloner); ok {
			target.ArtifactSource = &Settings_DirectoryArtifactSource{
				DirectoryArtifactSource: h.Clone().(*Settings_Directory),
			}
		} else {
			target.ArtifactSource = &Settings_DirectoryArtifactSource{
				DirectoryArtifactSource: proto.Clone(m.GetDirectoryArtifactSource()).(*Settings_Directory),
			}
		}

	case *Settings_ConsulKvArtifactSource:

		if h, ok := interface{}(m.GetConsulKvArtifactSource()).(clone.Cloner); ok {
			target.ArtifactSource = &Settings_ConsulKvArtifactSource{
				ConsulKvArtifactSource: h.Clone().(*Settings_ConsulKv),
			}
		} else {
			target.ArtifactSource = &Settings_ConsulKvArtifactSource{
				ConsulKvArtifactSource: proto.Clone(m.GetConsulKvArtifactSource()).(*Settings_ConsulKv),
			}
		}

	}

	return target
}

// Clone function
func (m *UpstreamOptions) Clone() proto.Message {
	var target *UpstreamOptions
	if m == nil {
		return target
	}
	target = &UpstreamOptions{}

	if h, ok := interface{}(m.GetSslParameters()).(clone.Cloner); ok {
		target.SslParameters = h.Clone().(*SslParameters)
	} else {
		target.SslParameters = proto.Clone(m.GetSslParameters()).(*SslParameters)
	}

	return target
}

// Clone function
func (m *GlooOptions) Clone() proto.Message {
	var target *GlooOptions
	if m == nil {
		return target
	}
	target = &GlooOptions{}

	target.XdsBindAddr = m.GetXdsBindAddr()

	target.ValidationBindAddr = m.GetValidationBindAddr()

	if h, ok := interface{}(m.GetCircuitBreakers()).(clone.Cloner); ok {
		target.CircuitBreakers = h.Clone().(*CircuitBreakerConfig)
	} else {
		target.CircuitBreakers = proto.Clone(m.GetCircuitBreakers()).(*CircuitBreakerConfig)
	}

	if h, ok := interface{}(m.GetEndpointsWarmingTimeout()).(clone.Cloner); ok {
		target.EndpointsWarmingTimeout = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.EndpointsWarmingTimeout = proto.Clone(m.GetEndpointsWarmingTimeout()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	if h, ok := interface{}(m.GetAwsOptions()).(clone.Cloner); ok {
		target.AwsOptions = h.Clone().(*GlooOptions_AWSOptions)
	} else {
		target.AwsOptions = proto.Clone(m.GetAwsOptions()).(*GlooOptions_AWSOptions)
	}

	if h, ok := interface{}(m.GetInvalidConfigPolicy()).(clone.Cloner); ok {
		target.InvalidConfigPolicy = h.Clone().(*GlooOptions_InvalidConfigPolicy)
	} else {
		target.InvalidConfigPolicy = proto.Clone(m.GetInvalidConfigPolicy()).(*GlooOptions_InvalidConfigPolicy)
	}

	target.DisableKubernetesDestinations = m.GetDisableKubernetesDestinations()

	if h, ok := interface{}(m.GetDisableGrpcWeb()).(clone.Cloner); ok {
		target.DisableGrpcWeb = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.DisableGrpcWeb = proto.Clone(m.GetDisableGrpcWeb()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetDisableProxyGarbageCollection()).(clone.Cloner); ok {
		target.DisableProxyGarbageCollection = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.DisableProxyGarbageCollection = proto.Clone(m.GetDisableProxyGarbageCollection()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetRegexMaxProgramSize()).(clone.Cloner); ok {
		target.RegexMaxProgramSize = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	} else {
		target.RegexMaxProgramSize = proto.Clone(m.GetRegexMaxProgramSize()).(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	}

	target.RestXdsBindAddr = m.GetRestXdsBindAddr()

	if h, ok := interface{}(m.GetEnableRestEds()).(clone.Cloner); ok {
		target.EnableRestEds = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.EnableRestEds = proto.Clone(m.GetEnableRestEds()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetFailoverUpstreamDnsPollingInterval()).(clone.Cloner); ok {
		target.FailoverUpstreamDnsPollingInterval = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.FailoverUpstreamDnsPollingInterval = proto.Clone(m.GetFailoverUpstreamDnsPollingInterval()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	if h, ok := interface{}(m.GetRemoveUnusedFilters()).(clone.Cloner); ok {
		target.RemoveUnusedFilters = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.RemoveUnusedFilters = proto.Clone(m.GetRemoveUnusedFilters()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	target.ProxyDebugBindAddr = m.GetProxyDebugBindAddr()

	return target
}

// Clone function
func (m *VirtualServiceOptions) Clone() proto.Message {
	var target *VirtualServiceOptions
	if m == nil {
		return target
	}
	target = &VirtualServiceOptions{}

	if h, ok := interface{}(m.GetOneWayTls()).(clone.Cloner); ok {
		target.OneWayTls = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.OneWayTls = proto.Clone(m.GetOneWayTls()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	return target
}

// Clone function
func (m *GatewayOptions) Clone() proto.Message {
	var target *GatewayOptions
	if m == nil {
		return target
	}
	target = &GatewayOptions{}

	target.ValidationServerAddr = m.GetValidationServerAddr()

	if h, ok := interface{}(m.GetValidation()).(clone.Cloner); ok {
		target.Validation = h.Clone().(*GatewayOptions_ValidationOptions)
	} else {
		target.Validation = proto.Clone(m.GetValidation()).(*GatewayOptions_ValidationOptions)
	}

	target.ReadGatewaysFromAllNamespaces = m.GetReadGatewaysFromAllNamespaces()

	target.AlwaysSortRouteTableRoutes = m.GetAlwaysSortRouteTableRoutes()

	target.CompressedProxySpec = m.GetCompressedProxySpec()

	if h, ok := interface{}(m.GetVirtualServiceOptions()).(clone.Cloner); ok {
		target.VirtualServiceOptions = h.Clone().(*VirtualServiceOptions)
	} else {
		target.VirtualServiceOptions = proto.Clone(m.GetVirtualServiceOptions()).(*VirtualServiceOptions)
	}

	if h, ok := interface{}(m.GetPersistProxySpec()).(clone.Cloner); ok {
		target.PersistProxySpec = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.PersistProxySpec = proto.Clone(m.GetPersistProxySpec()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetEnableGatewayController()).(clone.Cloner); ok {
		target.EnableGatewayController = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.EnableGatewayController = proto.Clone(m.GetEnableGatewayController()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	return target
}

// Clone function
func (m *ConsoleOptions) Clone() proto.Message {
	var target *ConsoleOptions
	if m == nil {
		return target
	}
	target = &ConsoleOptions{}

	if h, ok := interface{}(m.GetReadOnly()).(clone.Cloner); ok {
		target.ReadOnly = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.ReadOnly = proto.Clone(m.GetReadOnly()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetApiExplorerEnabled()).(clone.Cloner); ok {
		target.ApiExplorerEnabled = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.ApiExplorerEnabled = proto.Clone(m.GetApiExplorerEnabled()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	return target
}

// Clone function
func (m *Settings_KubernetesCrds) Clone() proto.Message {
	var target *Settings_KubernetesCrds
	if m == nil {
		return target
	}
	target = &Settings_KubernetesCrds{}

	return target
}

// Clone function
func (m *Settings_KubernetesSecrets) Clone() proto.Message {
	var target *Settings_KubernetesSecrets
	if m == nil {
		return target
	}
	target = &Settings_KubernetesSecrets{}

	return target
}

// Clone function
func (m *Settings_VaultSecrets) Clone() proto.Message {
	var target *Settings_VaultSecrets
	if m == nil {
		return target
	}
	target = &Settings_VaultSecrets{}

	target.Token = m.GetToken()

	target.Address = m.GetAddress()

	target.CaCert = m.GetCaCert()

	target.CaPath = m.GetCaPath()

	target.ClientCert = m.GetClientCert()

	target.ClientKey = m.GetClientKey()

	target.TlsServerName = m.GetTlsServerName()

	if h, ok := interface{}(m.GetInsecure()).(clone.Cloner); ok {
		target.Insecure = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.Insecure = proto.Clone(m.GetInsecure()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	target.RootKey = m.GetRootKey()

	target.PathPrefix = m.GetPathPrefix()

	return target
}

// Clone function
func (m *Settings_ConsulKv) Clone() proto.Message {
	var target *Settings_ConsulKv
	if m == nil {
		return target
	}
	target = &Settings_ConsulKv{}

	target.RootKey = m.GetRootKey()

	return target
}

// Clone function
func (m *Settings_KubernetesConfigmaps) Clone() proto.Message {
	var target *Settings_KubernetesConfigmaps
	if m == nil {
		return target
	}
	target = &Settings_KubernetesConfigmaps{}

	return target
}

// Clone function
func (m *Settings_Directory) Clone() proto.Message {
	var target *Settings_Directory
	if m == nil {
		return target
	}
	target = &Settings_Directory{}

	target.Directory = m.GetDirectory()

	return target
}

// Clone function
func (m *Settings_KnativeOptions) Clone() proto.Message {
	var target *Settings_KnativeOptions
	if m == nil {
		return target
	}
	target = &Settings_KnativeOptions{}

	target.ClusterIngressProxyAddress = m.GetClusterIngressProxyAddress()

	target.KnativeExternalProxyAddress = m.GetKnativeExternalProxyAddress()

	target.KnativeInternalProxyAddress = m.GetKnativeInternalProxyAddress()

	return target
}

// Clone function
func (m *Settings_DiscoveryOptions) Clone() proto.Message {
	var target *Settings_DiscoveryOptions
	if m == nil {
		return target
	}
	target = &Settings_DiscoveryOptions{}

	target.FdsMode = m.GetFdsMode()

	if h, ok := interface{}(m.GetUdsOptions()).(clone.Cloner); ok {
		target.UdsOptions = h.Clone().(*Settings_DiscoveryOptions_UdsOptions)
	} else {
		target.UdsOptions = proto.Clone(m.GetUdsOptions()).(*Settings_DiscoveryOptions_UdsOptions)
	}

	return target
}

// Clone function
func (m *Settings_ConsulConfiguration) Clone() proto.Message {
	var target *Settings_ConsulConfiguration
	if m == nil {
		return target
	}
	target = &Settings_ConsulConfiguration{}

	target.Address = m.GetAddress()

	target.Datacenter = m.GetDatacenter()

	target.Username = m.GetUsername()

	target.Password = m.GetPassword()

	target.Token = m.GetToken()

	target.CaFile = m.GetCaFile()

	target.CaPath = m.GetCaPath()

	target.CertFile = m.GetCertFile()

	target.KeyFile = m.GetKeyFile()

	if h, ok := interface{}(m.GetInsecureSkipVerify()).(clone.Cloner); ok {
		target.InsecureSkipVerify = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.InsecureSkipVerify = proto.Clone(m.GetInsecureSkipVerify()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetWaitTime()).(clone.Cloner); ok {
		target.WaitTime = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.WaitTime = proto.Clone(m.GetWaitTime()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	if h, ok := interface{}(m.GetServiceDiscovery()).(clone.Cloner); ok {
		target.ServiceDiscovery = h.Clone().(*Settings_ConsulConfiguration_ServiceDiscoveryOptions)
	} else {
		target.ServiceDiscovery = proto.Clone(m.GetServiceDiscovery()).(*Settings_ConsulConfiguration_ServiceDiscoveryOptions)
	}

	target.HttpAddress = m.GetHttpAddress()

	target.DnsAddress = m.GetDnsAddress()

	if h, ok := interface{}(m.GetDnsPollingInterval()).(clone.Cloner); ok {
		target.DnsPollingInterval = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.DnsPollingInterval = proto.Clone(m.GetDnsPollingInterval()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	return target
}

// Clone function
func (m *Settings_ConsulUpstreamDiscoveryConfiguration) Clone() proto.Message {
	var target *Settings_ConsulUpstreamDiscoveryConfiguration
	if m == nil {
		return target
	}
	target = &Settings_ConsulUpstreamDiscoveryConfiguration{}

	target.UseTlsTagging = m.GetUseTlsTagging()

	target.TlsTagName = m.GetTlsTagName()

	if h, ok := interface{}(m.GetRootCa()).(clone.Cloner); ok {
		target.RootCa = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
	} else {
		target.RootCa = proto.Clone(m.GetRootCa()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
	}

	target.SplitTlsServices = m.GetSplitTlsServices()

	target.ConsistencyMode = m.GetConsistencyMode()

	return target
}

// Clone function
func (m *Settings_KubernetesConfiguration) Clone() proto.Message {
	var target *Settings_KubernetesConfiguration
	if m == nil {
		return target
	}
	target = &Settings_KubernetesConfiguration{}

	if h, ok := interface{}(m.GetRateLimits()).(clone.Cloner); ok {
		target.RateLimits = h.Clone().(*Settings_KubernetesConfiguration_RateLimits)
	} else {
		target.RateLimits = proto.Clone(m.GetRateLimits()).(*Settings_KubernetesConfiguration_RateLimits)
	}

	return target
}

// Clone function
func (m *Settings_ObservabilityOptions) Clone() proto.Message {
	var target *Settings_ObservabilityOptions
	if m == nil {
		return target
	}
	target = &Settings_ObservabilityOptions{}

	if h, ok := interface{}(m.GetGrafanaIntegration()).(clone.Cloner); ok {
		target.GrafanaIntegration = h.Clone().(*Settings_ObservabilityOptions_GrafanaIntegration)
	} else {
		target.GrafanaIntegration = proto.Clone(m.GetGrafanaIntegration()).(*Settings_ObservabilityOptions_GrafanaIntegration)
	}

	if m.GetConfigStatusMetricLabels() != nil {
		target.ConfigStatusMetricLabels = make(map[string]*Settings_ObservabilityOptions_MetricLabels, len(m.GetConfigStatusMetricLabels()))
		for k, v := range m.GetConfigStatusMetricLabels() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.ConfigStatusMetricLabels[k] = h.Clone().(*Settings_ObservabilityOptions_MetricLabels)
			} else {
				target.ConfigStatusMetricLabels[k] = proto.Clone(v).(*Settings_ObservabilityOptions_MetricLabels)
			}

		}
	}

	return target
}

// Clone function
func (m *Settings_DiscoveryOptions_UdsOptions) Clone() proto.Message {
	var target *Settings_DiscoveryOptions_UdsOptions
	if m == nil {
		return target
	}
	target = &Settings_DiscoveryOptions_UdsOptions{}

	if h, ok := interface{}(m.GetEnabled()).(clone.Cloner); ok {
		target.Enabled = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.Enabled = proto.Clone(m.GetEnabled()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if m.GetWatchLabels() != nil {
		target.WatchLabels = make(map[string]string, len(m.GetWatchLabels()))
		for k, v := range m.GetWatchLabels() {

			target.WatchLabels[k] = v

		}
	}

	return target
}

// Clone function
func (m *Settings_ConsulConfiguration_ServiceDiscoveryOptions) Clone() proto.Message {
	var target *Settings_ConsulConfiguration_ServiceDiscoveryOptions
	if m == nil {
		return target
	}
	target = &Settings_ConsulConfiguration_ServiceDiscoveryOptions{}

	if m.GetDataCenters() != nil {
		target.DataCenters = make([]string, len(m.GetDataCenters()))
		for idx, v := range m.GetDataCenters() {

			target.DataCenters[idx] = v

		}
	}

	return target
}

// Clone function
func (m *Settings_KubernetesConfiguration_RateLimits) Clone() proto.Message {
	var target *Settings_KubernetesConfiguration_RateLimits
	if m == nil {
		return target
	}
	target = &Settings_KubernetesConfiguration_RateLimits{}

	target.QPS = m.GetQPS()

	target.Burst = m.GetBurst()

	return target
}

// Clone function
func (m *Settings_ObservabilityOptions_GrafanaIntegration) Clone() proto.Message {
	var target *Settings_ObservabilityOptions_GrafanaIntegration
	if m == nil {
		return target
	}
	target = &Settings_ObservabilityOptions_GrafanaIntegration{}

	if h, ok := interface{}(m.GetDefaultDashboardFolderId()).(clone.Cloner); ok {
		target.DefaultDashboardFolderId = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	} else {
		target.DefaultDashboardFolderId = proto.Clone(m.GetDefaultDashboardFolderId()).(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	}

	return target
}

// Clone function
func (m *Settings_ObservabilityOptions_MetricLabels) Clone() proto.Message {
	var target *Settings_ObservabilityOptions_MetricLabels
	if m == nil {
		return target
	}
	target = &Settings_ObservabilityOptions_MetricLabels{}

	if m.GetLabelToPath() != nil {
		target.LabelToPath = make(map[string]string, len(m.GetLabelToPath()))
		for k, v := range m.GetLabelToPath() {

			target.LabelToPath[k] = v

		}
	}

	return target
}

// Clone function
func (m *GlooOptions_AWSOptions) Clone() proto.Message {
	var target *GlooOptions_AWSOptions
	if m == nil {
		return target
	}
	target = &GlooOptions_AWSOptions{}

	if h, ok := interface{}(m.GetPropagateOriginalRouting()).(clone.Cloner); ok {
		target.PropagateOriginalRouting = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.PropagateOriginalRouting = proto.Clone(m.GetPropagateOriginalRouting()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetCredentialRefreshDelay()).(clone.Cloner); ok {
		target.CredentialRefreshDelay = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.CredentialRefreshDelay = proto.Clone(m.GetCredentialRefreshDelay()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	switch m.CredentialsFetcher.(type) {

	case *GlooOptions_AWSOptions_EnableCredentialsDiscovey:

		target.CredentialsFetcher = &GlooOptions_AWSOptions_EnableCredentialsDiscovey{
			EnableCredentialsDiscovey: m.GetEnableCredentialsDiscovey(),
		}

	case *GlooOptions_AWSOptions_ServiceAccountCredentials:

		if h, ok := interface{}(m.GetServiceAccountCredentials()).(clone.Cloner); ok {
			target.CredentialsFetcher = &GlooOptions_AWSOptions_ServiceAccountCredentials{
				ServiceAccountCredentials: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_aws.AWSLambdaConfig_ServiceAccountCredentials),
			}
		} else {
			target.CredentialsFetcher = &GlooOptions_AWSOptions_ServiceAccountCredentials{
				ServiceAccountCredentials: proto.Clone(m.GetServiceAccountCredentials()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_aws.AWSLambdaConfig_ServiceAccountCredentials),
			}
		}

	}

	return target
}

// Clone function
func (m *GlooOptions_InvalidConfigPolicy) Clone() proto.Message {
	var target *GlooOptions_InvalidConfigPolicy
	if m == nil {
		return target
	}
	target = &GlooOptions_InvalidConfigPolicy{}

	target.ReplaceInvalidRoutes = m.GetReplaceInvalidRoutes()

	target.InvalidRouteResponseCode = m.GetInvalidRouteResponseCode()

	target.InvalidRouteResponseBody = m.GetInvalidRouteResponseBody()

	return target
}

// Clone function
func (m *GatewayOptions_ValidationOptions) Clone() proto.Message {
	var target *GatewayOptions_ValidationOptions
	if m == nil {
		return target
	}
	target = &GatewayOptions_ValidationOptions{}

	target.ProxyValidationServerAddr = m.GetProxyValidationServerAddr()

	target.ValidationWebhookTlsCert = m.GetValidationWebhookTlsCert()

	target.ValidationWebhookTlsKey = m.GetValidationWebhookTlsKey()

	target.IgnoreGlooValidationFailure = m.GetIgnoreGlooValidationFailure()

	if h, ok := interface{}(m.GetAlwaysAccept()).(clone.Cloner); ok {
		target.AlwaysAccept = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.AlwaysAccept = proto.Clone(m.GetAlwaysAccept()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetAllowWarnings()).(clone.Cloner); ok {
		target.AllowWarnings = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.AllowWarnings = proto.Clone(m.GetAllowWarnings()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetWarnRouteShortCircuiting()).(clone.Cloner); ok {
		target.WarnRouteShortCircuiting = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.WarnRouteShortCircuiting = proto.Clone(m.GetWarnRouteShortCircuiting()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetDisableTransformationValidation()).(clone.Cloner); ok {
		target.DisableTransformationValidation = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.DisableTransformationValidation = proto.Clone(m.GetDisableTransformationValidation()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetValidationServerGrpcMaxSizeBytes()).(clone.Cloner); ok {
		target.ValidationServerGrpcMaxSizeBytes = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.Int32Value)
	} else {
		target.ValidationServerGrpcMaxSizeBytes = proto.Clone(m.GetValidationServerGrpcMaxSizeBytes()).(*github_com_golang_protobuf_ptypes_wrappers.Int32Value)
	}

	return target
}
