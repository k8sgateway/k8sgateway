// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/upstream.proto

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
func (m *Upstream) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Upstream)
	if !ok {
		that2, ok := that.(Upstream)
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

	if h, ok := interface{}(m.GetNamespacedStatuses()).(equality.Equalizer); ok {
		if !h.Equal(target.GetNamespacedStatuses()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetNamespacedStatuses(), target.GetNamespacedStatuses()) {
			return false
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

	if h, ok := interface{}(m.GetDiscoveryMetadata()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDiscoveryMetadata()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDiscoveryMetadata(), target.GetDiscoveryMetadata()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetSslConfig()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSslConfig()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSslConfig(), target.GetSslConfig()) {
			return false
		}
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

	if h, ok := interface{}(m.GetLoadBalancerConfig()).(equality.Equalizer); ok {
		if !h.Equal(target.GetLoadBalancerConfig()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetLoadBalancerConfig(), target.GetLoadBalancerConfig()) {
			return false
		}
	}

	if len(m.GetHealthChecks()) != len(target.GetHealthChecks()) {
		return false
	}
	for idx, v := range m.GetHealthChecks() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetHealthChecks()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetHealthChecks()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetOutlierDetection()).(equality.Equalizer); ok {
		if !h.Equal(target.GetOutlierDetection()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetOutlierDetection(), target.GetOutlierDetection()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetFailover()).(equality.Equalizer); ok {
		if !h.Equal(target.GetFailover()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetFailover(), target.GetFailover()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetConnectionConfig()).(equality.Equalizer); ok {
		if !h.Equal(target.GetConnectionConfig()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetConnectionConfig(), target.GetConnectionConfig()) {
			return false
		}
	}

	if m.GetProtocolSelection() != target.GetProtocolSelection() {
		return false
	}

	if h, ok := interface{}(m.GetUseHttp2()).(equality.Equalizer); ok {
		if !h.Equal(target.GetUseHttp2()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetUseHttp2(), target.GetUseHttp2()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetInitialStreamWindowSize()).(equality.Equalizer); ok {
		if !h.Equal(target.GetInitialStreamWindowSize()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetInitialStreamWindowSize(), target.GetInitialStreamWindowSize()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetInitialConnectionWindowSize()).(equality.Equalizer); ok {
		if !h.Equal(target.GetInitialConnectionWindowSize()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetInitialConnectionWindowSize(), target.GetInitialConnectionWindowSize()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetMaxConcurrentStreams()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMaxConcurrentStreams()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMaxConcurrentStreams(), target.GetMaxConcurrentStreams()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetOverrideStreamErrorOnInvalidHttpMessage()).(equality.Equalizer); ok {
		if !h.Equal(target.GetOverrideStreamErrorOnInvalidHttpMessage()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetOverrideStreamErrorOnInvalidHttpMessage(), target.GetOverrideStreamErrorOnInvalidHttpMessage()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetHttpProxyHostname()).(equality.Equalizer); ok {
		if !h.Equal(target.GetHttpProxyHostname()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetHttpProxyHostname(), target.GetHttpProxyHostname()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetHttpConnectSslConfig()).(equality.Equalizer); ok {
		if !h.Equal(target.GetHttpConnectSslConfig()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetHttpConnectSslConfig(), target.GetHttpConnectSslConfig()) {
			return false
		}
	}

	if len(m.GetHttpConnectHeaders()) != len(target.GetHttpConnectHeaders()) {
		return false
	}
	for idx, v := range m.GetHttpConnectHeaders() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetHttpConnectHeaders()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetHttpConnectHeaders()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetIgnoreHealthOnHostRemoval()).(equality.Equalizer); ok {
		if !h.Equal(target.GetIgnoreHealthOnHostRemoval()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetIgnoreHealthOnHostRemoval(), target.GetIgnoreHealthOnHostRemoval()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetRespectDnsTtl()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRespectDnsTtl()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRespectDnsTtl(), target.GetRespectDnsTtl()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetDnsRefreshRate()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDnsRefreshRate()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDnsRefreshRate(), target.GetDnsRefreshRate()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetProxyProtocolVersion()).(equality.Equalizer); ok {
		if !h.Equal(target.GetProxyProtocolVersion()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetProxyProtocolVersion(), target.GetProxyProtocolVersion()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetPreconnectPolicy()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPreconnectPolicy()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPreconnectPolicy(), target.GetPreconnectPolicy()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetDisableIstioAutoMtls()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDisableIstioAutoMtls()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDisableIstioAutoMtls(), target.GetDisableIstioAutoMtls()) {
			return false
		}
	}

	switch m.UpstreamType.(type) {

	case *Upstream_Kube:
		if _, ok := target.UpstreamType.(*Upstream_Kube); !ok {
			return false
		}

		if h, ok := interface{}(m.GetKube()).(equality.Equalizer); ok {
			if !h.Equal(target.GetKube()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetKube(), target.GetKube()) {
				return false
			}
		}

	case *Upstream_Static:
		if _, ok := target.UpstreamType.(*Upstream_Static); !ok {
			return false
		}

		if h, ok := interface{}(m.GetStatic()).(equality.Equalizer); ok {
			if !h.Equal(target.GetStatic()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetStatic(), target.GetStatic()) {
				return false
			}
		}

	case *Upstream_Pipe:
		if _, ok := target.UpstreamType.(*Upstream_Pipe); !ok {
			return false
		}

		if h, ok := interface{}(m.GetPipe()).(equality.Equalizer); ok {
			if !h.Equal(target.GetPipe()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetPipe(), target.GetPipe()) {
				return false
			}
		}

	case *Upstream_Aws:
		if _, ok := target.UpstreamType.(*Upstream_Aws); !ok {
			return false
		}

		if h, ok := interface{}(m.GetAws()).(equality.Equalizer); ok {
			if !h.Equal(target.GetAws()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetAws(), target.GetAws()) {
				return false
			}
		}

	case *Upstream_Azure:
		if _, ok := target.UpstreamType.(*Upstream_Azure); !ok {
			return false
		}

		if h, ok := interface{}(m.GetAzure()).(equality.Equalizer); ok {
			if !h.Equal(target.GetAzure()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetAzure(), target.GetAzure()) {
				return false
			}
		}

	case *Upstream_Consul:
		if _, ok := target.UpstreamType.(*Upstream_Consul); !ok {
			return false
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

	case *Upstream_AwsEc2:
		if _, ok := target.UpstreamType.(*Upstream_AwsEc2); !ok {
			return false
		}

		if h, ok := interface{}(m.GetAwsEc2()).(equality.Equalizer); ok {
			if !h.Equal(target.GetAwsEc2()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetAwsEc2(), target.GetAwsEc2()) {
				return false
			}
		}

	case *Upstream_Gcp:
		if _, ok := target.UpstreamType.(*Upstream_Gcp); !ok {
			return false
		}

		if h, ok := interface{}(m.GetGcp()).(equality.Equalizer); ok {
			if !h.Equal(target.GetGcp()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetGcp(), target.GetGcp()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.UpstreamType != target.UpstreamType {
			return false
		}
	}

	return true
}

// Equal function
func (m *DiscoveryMetadata) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*DiscoveryMetadata)
	if !ok {
		that2, ok := that.(DiscoveryMetadata)
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

	if len(m.GetLabels()) != len(target.GetLabels()) {
		return false
	}
	for k, v := range m.GetLabels() {

		if strings.Compare(v, target.GetLabels()[k]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *HeaderValue) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HeaderValue)
	if !ok {
		that2, ok := that.(HeaderValue)
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

	if strings.Compare(m.GetKey(), target.GetKey()) != 0 {
		return false
	}

	if strings.Compare(m.GetValue(), target.GetValue()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *PreconnectPolicy) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*PreconnectPolicy)
	if !ok {
		that2, ok := that.(PreconnectPolicy)
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

	if h, ok := interface{}(m.GetPerUpstreamPreconnectRatio()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPerUpstreamPreconnectRatio()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPerUpstreamPreconnectRatio(), target.GetPerUpstreamPreconnectRatio()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetPredictivePreconnectRatio()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPredictivePreconnectRatio()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPredictivePreconnectRatio(), target.GetPredictivePreconnectRatio()) {
			return false
		}
	}

	return true
}
