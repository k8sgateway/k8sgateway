// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/settings.proto

package v1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = equality.Equalizer(nil)
	_ = proto.Message(nil)
)

// Equal function
func (m *Settings) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Settings)
	if !ok {
		that2, ok := that.(Settings)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetDiscoveryNamespace(), target.GetDiscoveryNamespace()) != 0 {
		return false
	}

	if len(m.GetWatchNamespaces()) != len(target.GetWatchNamespaces()) {
		return false
	}
	for idx, v := range m.GetWatchNamespaces() {

		if strings.Compare(v, target.GetWatchNamespaces()[idx]) != 0 {
			return false
		}

	}

	if h, ok := interface{}(m.GetRefreshRate()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRefreshRate()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRefreshRate(), target.GetRefreshRate()) {
			return false
		}
	}

	if m.GetDevMode() != target.GetDevMode() {
		return false
	}

	if m.GetLinkerd() != target.GetLinkerd() {
		return false
	}

	if h, ok := interface{}(m.GetKnative()).(equality.Equalizer); ok {
		if !h.Equal(target.GetKnative()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetKnative(), target.GetKnative()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetDiscovery()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDiscovery()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDiscovery(), target.GetDiscovery()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetGloo()).(equality.Equalizer); ok {
		if !h.Equal(target.GetGloo()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetGloo(), target.GetGloo()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetGateway()).(equality.Equalizer); ok {
		if !h.Equal(target.GetGateway()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetGateway(), target.GetGateway()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetConsul()).(equality.Equalizer); ok {
		if !h.Equal(target.GetConsul()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetConsul(), target.GetConsul()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetConsulDiscovery()).(equality.Equalizer); ok {
		if !h.Equal(target.GetConsulDiscovery()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetConsulDiscovery(), target.GetConsulDiscovery()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetKubernetes()).(equality.Equalizer); ok {
		if !h.Equal(target.GetKubernetes()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetKubernetes(), target.GetKubernetes()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetExtensions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetExtensions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetExtensions(), target.GetExtensions()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetRatelimit()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRatelimit()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRatelimit(), target.GetRatelimit()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetRatelimitServer()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRatelimitServer()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRatelimitServer(), target.GetRatelimitServer()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetRbac()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRbac()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRbac(), target.GetRbac()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetExtauth()).(equality.Equalizer); ok {
		if !h.Equal(target.GetExtauth()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetExtauth(), target.GetExtauth()) {
			return false
		}
	}

	if len(m.GetNamedExtauth()) != len(target.GetNamedExtauth()) {
		return false
	}
	for k, v := range m.GetNamedExtauth() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetNamedExtauth()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetNamedExtauth()[k]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetMetadata()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMetadata()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMetadata(), target.GetMetadata()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetStatus()).(equality.Equalizer); ok {
		if !h.Equal(target.GetStatus()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetStatus(), target.GetStatus()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetObservabilityOptions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetObservabilityOptions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetObservabilityOptions(), target.GetObservabilityOptions()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetUpstreamOptions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetUpstreamOptions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetUpstreamOptions(), target.GetUpstreamOptions()) {
			return false
		}
	}

	switch m.ConfigSource.(type) {

	case *Settings_KubernetesConfigSource:
		if _, ok := target.ConfigSource.(*Settings_KubernetesConfigSource); !ok {
			return false
		}

		if h, ok := interface{}(m.GetKubernetesConfigSource()).(equality.Equalizer); ok {
			if !h.Equal(target.GetKubernetesConfigSource()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetKubernetesConfigSource(), target.GetKubernetesConfigSource()) {
				return false
			}
		}

	case *Settings_DirectoryConfigSource:
		if _, ok := target.ConfigSource.(*Settings_DirectoryConfigSource); !ok {
			return false
		}

		if h, ok := interface{}(m.GetDirectoryConfigSource()).(equality.Equalizer); ok {
			if !h.Equal(target.GetDirectoryConfigSource()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetDirectoryConfigSource(), target.GetDirectoryConfigSource()) {
				return false
			}
		}

	case *Settings_ConsulKvSource:
		if _, ok := target.ConfigSource.(*Settings_ConsulKvSource); !ok {
			return false
		}

		if h, ok := interface{}(m.GetConsulKvSource()).(equality.Equalizer); ok {
			if !h.Equal(target.GetConsulKvSource()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetConsulKvSource(), target.GetConsulKvSource()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.ConfigSource != target.ConfigSource {
			return false
		}
	}

	switch m.SecretSource.(type) {

	case *Settings_KubernetesSecretSource:
		if _, ok := target.SecretSource.(*Settings_KubernetesSecretSource); !ok {
			return false
		}

		if h, ok := interface{}(m.GetKubernetesSecretSource()).(equality.Equalizer); ok {
			if !h.Equal(target.GetKubernetesSecretSource()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetKubernetesSecretSource(), target.GetKubernetesSecretSource()) {
				return false
			}
		}

	case *Settings_VaultSecretSource:
		if _, ok := target.SecretSource.(*Settings_VaultSecretSource); !ok {
			return false
		}

		if h, ok := interface{}(m.GetVaultSecretSource()).(equality.Equalizer); ok {
			if !h.Equal(target.GetVaultSecretSource()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetVaultSecretSource(), target.GetVaultSecretSource()) {
				return false
			}
		}

	case *Settings_DirectorySecretSource:
		if _, ok := target.SecretSource.(*Settings_DirectorySecretSource); !ok {
			return false
		}

		if h, ok := interface{}(m.GetDirectorySecretSource()).(equality.Equalizer); ok {
			if !h.Equal(target.GetDirectorySecretSource()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetDirectorySecretSource(), target.GetDirectorySecretSource()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.SecretSource != target.SecretSource {
			return false
		}
	}

	switch m.ArtifactSource.(type) {

	case *Settings_KubernetesArtifactSource:
		if _, ok := target.ArtifactSource.(*Settings_KubernetesArtifactSource); !ok {
			return false
		}

		if h, ok := interface{}(m.GetKubernetesArtifactSource()).(equality.Equalizer); ok {
			if !h.Equal(target.GetKubernetesArtifactSource()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetKubernetesArtifactSource(), target.GetKubernetesArtifactSource()) {
				return false
			}
		}

	case *Settings_DirectoryArtifactSource:
		if _, ok := target.ArtifactSource.(*Settings_DirectoryArtifactSource); !ok {
			return false
		}

		if h, ok := interface{}(m.GetDirectoryArtifactSource()).(equality.Equalizer); ok {
			if !h.Equal(target.GetDirectoryArtifactSource()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetDirectoryArtifactSource(), target.GetDirectoryArtifactSource()) {
				return false
			}
		}

	case *Settings_ConsulKvArtifactSource:
		if _, ok := target.ArtifactSource.(*Settings_ConsulKvArtifactSource); !ok {
			return false
		}

		if h, ok := interface{}(m.GetConsulKvArtifactSource()).(equality.Equalizer); ok {
			if !h.Equal(target.GetConsulKvArtifactSource()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetConsulKvArtifactSource(), target.GetConsulKvArtifactSource()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.ArtifactSource != target.ArtifactSource {
			return false
		}
	}

	return true
}

// Equal function
func (m *UpstreamOptions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*UpstreamOptions)
	if !ok {
		that2, ok := that.(UpstreamOptions)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetSslParameters()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSslParameters()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSslParameters(), target.GetSslParameters()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetOneWayTls()).(equality.Equalizer); ok {
		if !h.Equal(target.GetOneWayTls()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetOneWayTls(), target.GetOneWayTls()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *GlooOptions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GlooOptions)
	if !ok {
		that2, ok := that.(GlooOptions)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetXdsBindAddr(), target.GetXdsBindAddr()) != 0 {
		return false
	}

	if strings.Compare(m.GetValidationBindAddr(), target.GetValidationBindAddr()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetCircuitBreakers()).(equality.Equalizer); ok {
		if !h.Equal(target.GetCircuitBreakers()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetCircuitBreakers(), target.GetCircuitBreakers()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetEndpointsWarmingTimeout()).(equality.Equalizer); ok {
		if !h.Equal(target.GetEndpointsWarmingTimeout()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetEndpointsWarmingTimeout(), target.GetEndpointsWarmingTimeout()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetAwsOptions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetAwsOptions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetAwsOptions(), target.GetAwsOptions()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetInvalidConfigPolicy()).(equality.Equalizer); ok {
		if !h.Equal(target.GetInvalidConfigPolicy()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetInvalidConfigPolicy(), target.GetInvalidConfigPolicy()) {
			return false
		}
	}

	if m.GetDisableKubernetesDestinations() != target.GetDisableKubernetesDestinations() {
		return false
	}

	if h, ok := interface{}(m.GetDisableGrpcWeb()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDisableGrpcWeb()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDisableGrpcWeb(), target.GetDisableGrpcWeb()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetDisableProxyGarbageCollection()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDisableProxyGarbageCollection()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDisableProxyGarbageCollection(), target.GetDisableProxyGarbageCollection()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetRegexMaxProgramSize()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRegexMaxProgramSize()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRegexMaxProgramSize(), target.GetRegexMaxProgramSize()) {
			return false
		}
	}

	if strings.Compare(m.GetRestXdsBindAddr(), target.GetRestXdsBindAddr()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetEnableRestEds()).(equality.Equalizer); ok {
		if !h.Equal(target.GetEnableRestEds()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetEnableRestEds(), target.GetEnableRestEds()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetFailoverUpstreamDnsPollingInterval()).(equality.Equalizer); ok {
		if !h.Equal(target.GetFailoverUpstreamDnsPollingInterval()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetFailoverUpstreamDnsPollingInterval(), target.GetFailoverUpstreamDnsPollingInterval()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *VirtualServiceOptions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualServiceOptions)
	if !ok {
		that2, ok := that.(VirtualServiceOptions)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetOneWayTls()).(equality.Equalizer); ok {
		if !h.Equal(target.GetOneWayTls()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetOneWayTls(), target.GetOneWayTls()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *GatewayOptions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GatewayOptions)
	if !ok {
		that2, ok := that.(GatewayOptions)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetValidationServerAddr(), target.GetValidationServerAddr()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetValidation()).(equality.Equalizer); ok {
		if !h.Equal(target.GetValidation()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetValidation(), target.GetValidation()) {
			return false
		}
	}

	if m.GetReadGatewaysFromAllNamespaces() != target.GetReadGatewaysFromAllNamespaces() {
		return false
	}

	if m.GetAlwaysSortRouteTableRoutes() != target.GetAlwaysSortRouteTableRoutes() {
		return false
	}

	if m.GetCompressedProxySpec() != target.GetCompressedProxySpec() {
		return false
	}

	if h, ok := interface{}(m.GetVirtualServiceOptions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetVirtualServiceOptions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetVirtualServiceOptions(), target.GetVirtualServiceOptions()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *Settings_KubernetesCrds) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Settings_KubernetesCrds)
	if !ok {
		that2, ok := that.(Settings_KubernetesCrds)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	return true
}

// Equal function
func (m *Settings_KubernetesSecrets) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Settings_KubernetesSecrets)
	if !ok {
		that2, ok := that.(Settings_KubernetesSecrets)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	return true
}

// Equal function
func (m *Settings_VaultSecrets) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Settings_VaultSecrets)
	if !ok {
		that2, ok := that.(Settings_VaultSecrets)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetToken(), target.GetToken()) != 0 {
		return false
	}

	if strings.Compare(m.GetAddress(), target.GetAddress()) != 0 {
		return false
	}

	if strings.Compare(m.GetCaCert(), target.GetCaCert()) != 0 {
		return false
	}

	if strings.Compare(m.GetCaPath(), target.GetCaPath()) != 0 {
		return false
	}

	if strings.Compare(m.GetClientCert(), target.GetClientCert()) != 0 {
		return false
	}

	if strings.Compare(m.GetClientKey(), target.GetClientKey()) != 0 {
		return false
	}

	if strings.Compare(m.GetTlsServerName(), target.GetTlsServerName()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetInsecure()).(equality.Equalizer); ok {
		if !h.Equal(target.GetInsecure()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetInsecure(), target.GetInsecure()) {
			return false
		}
	}

	if strings.Compare(m.GetRootKey(), target.GetRootKey()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *Settings_ConsulKv) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Settings_ConsulKv)
	if !ok {
		that2, ok := that.(Settings_ConsulKv)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetRootKey(), target.GetRootKey()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *Settings_KubernetesConfigmaps) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Settings_KubernetesConfigmaps)
	if !ok {
		that2, ok := that.(Settings_KubernetesConfigmaps)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	return true
}

// Equal function
func (m *Settings_Directory) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Settings_Directory)
	if !ok {
		that2, ok := that.(Settings_Directory)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetDirectory(), target.GetDirectory()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *Settings_KnativeOptions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Settings_KnativeOptions)
	if !ok {
		that2, ok := that.(Settings_KnativeOptions)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetClusterIngressProxyAddress(), target.GetClusterIngressProxyAddress()) != 0 {
		return false
	}

	if strings.Compare(m.GetKnativeExternalProxyAddress(), target.GetKnativeExternalProxyAddress()) != 0 {
		return false
	}

	if strings.Compare(m.GetKnativeInternalProxyAddress(), target.GetKnativeInternalProxyAddress()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *Settings_DiscoveryOptions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Settings_DiscoveryOptions)
	if !ok {
		that2, ok := that.(Settings_DiscoveryOptions)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.GetFdsMode() != target.GetFdsMode() {
		return false
	}

	return true
}

// Equal function
func (m *Settings_ConsulConfiguration) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Settings_ConsulConfiguration)
	if !ok {
		that2, ok := that.(Settings_ConsulConfiguration)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetAddress(), target.GetAddress()) != 0 {
		return false
	}

	if strings.Compare(m.GetDatacenter(), target.GetDatacenter()) != 0 {
		return false
	}

	if strings.Compare(m.GetUsername(), target.GetUsername()) != 0 {
		return false
	}

	if strings.Compare(m.GetPassword(), target.GetPassword()) != 0 {
		return false
	}

	if strings.Compare(m.GetToken(), target.GetToken()) != 0 {
		return false
	}

	if strings.Compare(m.GetCaFile(), target.GetCaFile()) != 0 {
		return false
	}

	if strings.Compare(m.GetCaPath(), target.GetCaPath()) != 0 {
		return false
	}

	if strings.Compare(m.GetCertFile(), target.GetCertFile()) != 0 {
		return false
	}

	if strings.Compare(m.GetKeyFile(), target.GetKeyFile()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetInsecureSkipVerify()).(equality.Equalizer); ok {
		if !h.Equal(target.GetInsecureSkipVerify()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetInsecureSkipVerify(), target.GetInsecureSkipVerify()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetWaitTime()).(equality.Equalizer); ok {
		if !h.Equal(target.GetWaitTime()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetWaitTime(), target.GetWaitTime()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetServiceDiscovery()).(equality.Equalizer); ok {
		if !h.Equal(target.GetServiceDiscovery()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetServiceDiscovery(), target.GetServiceDiscovery()) {
			return false
		}
	}

	if strings.Compare(m.GetHttpAddress(), target.GetHttpAddress()) != 0 {
		return false
	}

	if strings.Compare(m.GetDnsAddress(), target.GetDnsAddress()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetDnsPollingInterval()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDnsPollingInterval()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDnsPollingInterval(), target.GetDnsPollingInterval()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *Settings_ConsulUpstreamDiscoveryConfiguration) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Settings_ConsulUpstreamDiscoveryConfiguration)
	if !ok {
		that2, ok := that.(Settings_ConsulUpstreamDiscoveryConfiguration)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.GetUseTlsTagging() != target.GetUseTlsTagging() {
		return false
	}

	if strings.Compare(m.GetTlsTagName(), target.GetTlsTagName()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetRootCa()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRootCa()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRootCa(), target.GetRootCa()) {
			return false
		}
	}

	if m.GetSplitTlsServices() != target.GetSplitTlsServices() {
		return false
	}

	return true
}

// Equal function
func (m *Settings_KubernetesConfiguration) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Settings_KubernetesConfiguration)
	if !ok {
		that2, ok := that.(Settings_KubernetesConfiguration)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetRateLimits()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRateLimits()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRateLimits(), target.GetRateLimits()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *Settings_ObservabilityOptions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Settings_ObservabilityOptions)
	if !ok {
		that2, ok := that.(Settings_ObservabilityOptions)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetGrafanaIntegration()).(equality.Equalizer); ok {
		if !h.Equal(target.GetGrafanaIntegration()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetGrafanaIntegration(), target.GetGrafanaIntegration()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *Settings_ConsulConfiguration_ServiceDiscoveryOptions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Settings_ConsulConfiguration_ServiceDiscoveryOptions)
	if !ok {
		that2, ok := that.(Settings_ConsulConfiguration_ServiceDiscoveryOptions)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetDataCenters()) != len(target.GetDataCenters()) {
		return false
	}
	for idx, v := range m.GetDataCenters() {

		if strings.Compare(v, target.GetDataCenters()[idx]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *Settings_KubernetesConfiguration_RateLimits) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Settings_KubernetesConfiguration_RateLimits)
	if !ok {
		that2, ok := that.(Settings_KubernetesConfiguration_RateLimits)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.GetQPS() != target.GetQPS() {
		return false
	}

	if m.GetBurst() != target.GetBurst() {
		return false
	}

	return true
}

// Equal function
func (m *Settings_ObservabilityOptions_GrafanaIntegration) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Settings_ObservabilityOptions_GrafanaIntegration)
	if !ok {
		that2, ok := that.(Settings_ObservabilityOptions_GrafanaIntegration)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetDefaultDashboardFolderId()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDefaultDashboardFolderId()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDefaultDashboardFolderId(), target.GetDefaultDashboardFolderId()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *GlooOptions_AWSOptions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GlooOptions_AWSOptions)
	if !ok {
		that2, ok := that.(GlooOptions_AWSOptions)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	switch m.CredentialsFetcher.(type) {

	case *GlooOptions_AWSOptions_EnableCredentialsDiscovey:
		if _, ok := target.CredentialsFetcher.(*GlooOptions_AWSOptions_EnableCredentialsDiscovey); !ok {
			return false
		}

		if m.GetEnableCredentialsDiscovey() != target.GetEnableCredentialsDiscovey() {
			return false
		}

	case *GlooOptions_AWSOptions_ServiceAccountCredentials:
		if _, ok := target.CredentialsFetcher.(*GlooOptions_AWSOptions_ServiceAccountCredentials); !ok {
			return false
		}

		if h, ok := interface{}(m.GetServiceAccountCredentials()).(equality.Equalizer); ok {
			if !h.Equal(target.GetServiceAccountCredentials()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetServiceAccountCredentials(), target.GetServiceAccountCredentials()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.CredentialsFetcher != target.CredentialsFetcher {
			return false
		}
	}

	return true
}

// Equal function
func (m *GlooOptions_InvalidConfigPolicy) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GlooOptions_InvalidConfigPolicy)
	if !ok {
		that2, ok := that.(GlooOptions_InvalidConfigPolicy)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.GetReplaceInvalidRoutes() != target.GetReplaceInvalidRoutes() {
		return false
	}

	if m.GetInvalidRouteResponseCode() != target.GetInvalidRouteResponseCode() {
		return false
	}

	if strings.Compare(m.GetInvalidRouteResponseBody(), target.GetInvalidRouteResponseBody()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *GatewayOptions_ValidationOptions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GatewayOptions_ValidationOptions)
	if !ok {
		that2, ok := that.(GatewayOptions_ValidationOptions)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetProxyValidationServerAddr(), target.GetProxyValidationServerAddr()) != 0 {
		return false
	}

	if strings.Compare(m.GetValidationWebhookTlsCert(), target.GetValidationWebhookTlsCert()) != 0 {
		return false
	}

	if strings.Compare(m.GetValidationWebhookTlsKey(), target.GetValidationWebhookTlsKey()) != 0 {
		return false
	}

	if m.GetIgnoreGlooValidationFailure() != target.GetIgnoreGlooValidationFailure() {
		return false
	}

	if h, ok := interface{}(m.GetAlwaysAccept()).(equality.Equalizer); ok {
		if !h.Equal(target.GetAlwaysAccept()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetAlwaysAccept(), target.GetAlwaysAccept()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetAllowWarnings()).(equality.Equalizer); ok {
		if !h.Equal(target.GetAllowWarnings()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetAllowWarnings(), target.GetAllowWarnings()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetWarnRouteShortCircuiting()).(equality.Equalizer); ok {
		if !h.Equal(target.GetWarnRouteShortCircuiting()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetWarnRouteShortCircuiting(), target.GetWarnRouteShortCircuiting()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetDisableTransformationValidation()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDisableTransformationValidation()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDisableTransformationValidation(), target.GetDisableTransformationValidation()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetValidationServerGrpcMaxSize()).(equality.Equalizer); ok {
		if !h.Equal(target.GetValidationServerGrpcMaxSize()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetValidationServerGrpcMaxSize(), target.GetValidationServerGrpcMaxSize()) {
			return false
		}
	}

	return true
}
