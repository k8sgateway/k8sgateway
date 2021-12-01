// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/ssl.proto

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
func (m *SslConfig) Clone() proto.Message {
	var target *SslConfig
	if m == nil {
		return target
	}
	target = &SslConfig{}

	if m.GetSniDomains() != nil {
		target.SniDomains = make([]string, len(m.GetSniDomains()))
		for idx, v := range m.GetSniDomains() {

			target.SniDomains[idx] = v

		}
	}

	if m.GetVerifySubjectAltName() != nil {
		target.VerifySubjectAltName = make([]string, len(m.GetVerifySubjectAltName()))
		for idx, v := range m.GetVerifySubjectAltName() {

			target.VerifySubjectAltName[idx] = v

		}
	}

	if h, ok := interface{}(m.GetParameters()).(clone.Cloner); ok {
		target.Parameters = h.Clone().(*SslParameters)
	} else {
		target.Parameters = proto.Clone(m.GetParameters()).(*SslParameters)
	}

	if m.GetAlpnProtocols() != nil {
		target.AlpnProtocols = make([]string, len(m.GetAlpnProtocols()))
		for idx, v := range m.GetAlpnProtocols() {

			target.AlpnProtocols[idx] = v

		}
	}

	if h, ok := interface{}(m.GetOneWayTls()).(clone.Cloner); ok {
		target.OneWayTls = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.OneWayTls = proto.Clone(m.GetOneWayTls()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetDisableTlsSessionResumption()).(clone.Cloner); ok {
		target.DisableTlsSessionResumption = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.DisableTlsSessionResumption = proto.Clone(m.GetDisableTlsSessionResumption()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetTransportSocketConnectTimeout()).(clone.Cloner); ok {
		target.TransportSocketConnectTimeout = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.TransportSocketConnectTimeout = proto.Clone(m.GetTransportSocketConnectTimeout()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	switch m.SslSecrets.(type) {

	case *SslConfig_SecretRef:

		if h, ok := interface{}(m.GetSecretRef()).(clone.Cloner); ok {
			target.SslSecrets = &SslConfig_SecretRef{
				SecretRef: h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef),
			}
		} else {
			target.SslSecrets = &SslConfig_SecretRef{
				SecretRef: proto.Clone(m.GetSecretRef()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef),
			}
		}

	case *SslConfig_SslFiles:

		if h, ok := interface{}(m.GetSslFiles()).(clone.Cloner); ok {
			target.SslSecrets = &SslConfig_SslFiles{
				SslFiles: h.Clone().(*SSLFiles),
			}
		} else {
			target.SslSecrets = &SslConfig_SslFiles{
				SslFiles: proto.Clone(m.GetSslFiles()).(*SSLFiles),
			}
		}

	case *SslConfig_Sds:

		if h, ok := interface{}(m.GetSds()).(clone.Cloner); ok {
			target.SslSecrets = &SslConfig_Sds{
				Sds: h.Clone().(*SDSConfig),
			}
		} else {
			target.SslSecrets = &SslConfig_Sds{
				Sds: proto.Clone(m.GetSds()).(*SDSConfig),
			}
		}

	}

	return target
}

// Clone function
func (m *SSLFiles) Clone() proto.Message {
	var target *SSLFiles
	if m == nil {
		return target
	}
	target = &SSLFiles{}

	target.TlsCert = m.GetTlsCert()

	target.TlsKey = m.GetTlsKey()

	target.RootCa = m.GetRootCa()

	return target
}

// Clone function
func (m *UpstreamSslConfig) Clone() proto.Message {
	var target *UpstreamSslConfig
	if m == nil {
		return target
	}
	target = &UpstreamSslConfig{}

	target.Sni = m.GetSni()

	if m.GetVerifySubjectAltName() != nil {
		target.VerifySubjectAltName = make([]string, len(m.GetVerifySubjectAltName()))
		for idx, v := range m.GetVerifySubjectAltName() {

			target.VerifySubjectAltName[idx] = v

		}
	}

	if h, ok := interface{}(m.GetParameters()).(clone.Cloner); ok {
		target.Parameters = h.Clone().(*SslParameters)
	} else {
		target.Parameters = proto.Clone(m.GetParameters()).(*SslParameters)
	}

	if m.GetAlpnProtocols() != nil {
		target.AlpnProtocols = make([]string, len(m.GetAlpnProtocols()))
		for idx, v := range m.GetAlpnProtocols() {

			target.AlpnProtocols[idx] = v

		}
	}

	switch m.SslSecrets.(type) {

	case *UpstreamSslConfig_SecretRef:

		if h, ok := interface{}(m.GetSecretRef()).(clone.Cloner); ok {
			target.SslSecrets = &UpstreamSslConfig_SecretRef{
				SecretRef: h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef),
			}
		} else {
			target.SslSecrets = &UpstreamSslConfig_SecretRef{
				SecretRef: proto.Clone(m.GetSecretRef()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef),
			}
		}

	case *UpstreamSslConfig_SslFiles:

		if h, ok := interface{}(m.GetSslFiles()).(clone.Cloner); ok {
			target.SslSecrets = &UpstreamSslConfig_SslFiles{
				SslFiles: h.Clone().(*SSLFiles),
			}
		} else {
			target.SslSecrets = &UpstreamSslConfig_SslFiles{
				SslFiles: proto.Clone(m.GetSslFiles()).(*SSLFiles),
			}
		}

	case *UpstreamSslConfig_Sds:

		if h, ok := interface{}(m.GetSds()).(clone.Cloner); ok {
			target.SslSecrets = &UpstreamSslConfig_Sds{
				Sds: h.Clone().(*SDSConfig),
			}
		} else {
			target.SslSecrets = &UpstreamSslConfig_Sds{
				Sds: proto.Clone(m.GetSds()).(*SDSConfig),
			}
		}

	}

	return target
}

// Clone function
func (m *SDSConfig) Clone() proto.Message {
	var target *SDSConfig
	if m == nil {
		return target
	}
	target = &SDSConfig{}

	target.TargetUri = m.GetTargetUri()

	target.CertificatesSecretName = m.GetCertificatesSecretName()

	target.ValidationContextName = m.GetValidationContextName()

	switch m.SdsBuilder.(type) {

	case *SDSConfig_CallCredentials:

		if h, ok := interface{}(m.GetCallCredentials()).(clone.Cloner); ok {
			target.SdsBuilder = &SDSConfig_CallCredentials{
				CallCredentials: h.Clone().(*CallCredentials),
			}
		} else {
			target.SdsBuilder = &SDSConfig_CallCredentials{
				CallCredentials: proto.Clone(m.GetCallCredentials()).(*CallCredentials),
			}
		}

	case *SDSConfig_ClusterName:

		target.SdsBuilder = &SDSConfig_ClusterName{
			ClusterName: m.GetClusterName(),
		}

	}

	return target
}

// Clone function
func (m *CallCredentials) Clone() proto.Message {
	var target *CallCredentials
	if m == nil {
		return target
	}
	target = &CallCredentials{}

	if h, ok := interface{}(m.GetFileCredentialSource()).(clone.Cloner); ok {
		target.FileCredentialSource = h.Clone().(*CallCredentials_FileCredentialSource)
	} else {
		target.FileCredentialSource = proto.Clone(m.GetFileCredentialSource()).(*CallCredentials_FileCredentialSource)
	}

	return target
}

// Clone function
func (m *SslParameters) Clone() proto.Message {
	var target *SslParameters
	if m == nil {
		return target
	}
	target = &SslParameters{}

	target.MinimumProtocolVersion = m.GetMinimumProtocolVersion()

	target.MaximumProtocolVersion = m.GetMaximumProtocolVersion()

	if m.GetCipherSuites() != nil {
		target.CipherSuites = make([]string, len(m.GetCipherSuites()))
		for idx, v := range m.GetCipherSuites() {

			target.CipherSuites[idx] = v

		}
	}

	if m.GetEcdhCurves() != nil {
		target.EcdhCurves = make([]string, len(m.GetEcdhCurves()))
		for idx, v := range m.GetEcdhCurves() {

			target.EcdhCurves[idx] = v

		}
	}

	return target
}

// Clone function
func (m *CallCredentials_FileCredentialSource) Clone() proto.Message {
	var target *CallCredentials_FileCredentialSource
	if m == nil {
		return target
	}
	target = &CallCredentials_FileCredentialSource{}

	target.TokenFileName = m.GetTokenFileName()

	target.Header = m.GetHeader()

	return target
}
