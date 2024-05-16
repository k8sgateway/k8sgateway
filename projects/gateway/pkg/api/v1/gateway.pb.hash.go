// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gateway/api/v1/gateway.proto

package v1

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"hash/fnv"

	safe_hasher "github.com/solo-io/protoc-gen-ext/pkg/hasher"
	"github.com/solo-io/protoc-gen-ext/pkg/hasher/hashstructure"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = new(hash.Hash64)
	_ = fnv.New64
	_ = hashstructure.Hash
	_ = new(safe_hasher.SafeHasher)
)

// Hash function
func (m *Gateway) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gateway.solo.io.github.com/solo-io/gloo/projects/gateway/pkg/api/v1.Gateway")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetSsl())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetBindAddress())); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetBindPort())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetOptions()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Options")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetOptions(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Options")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetMetadata()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Metadata")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMetadata(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Metadata")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetUseProxyProto()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("UseProxyProto")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetUseProxyProto(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("UseProxyProto")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	for _, v := range m.GetProxyNames() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	if h, ok := interface{}(m.GetRouteOptions()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("RouteOptions")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRouteOptions(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("RouteOptions")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	switch m.GatewayType.(type) {

	case *Gateway_HttpGateway:

		if h, ok := interface{}(m.GetHttpGateway()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("HttpGateway")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetHttpGateway(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("HttpGateway")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *Gateway_TcpGateway:

		if h, ok := interface{}(m.GetTcpGateway()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("TcpGateway")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetTcpGateway(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("TcpGateway")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *Gateway_HybridGateway:

		if h, ok := interface{}(m.GetHybridGateway()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("HybridGateway")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetHybridGateway(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("HybridGateway")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *TcpGateway) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gateway.solo.io.github.com/solo-io/gloo/projects/gateway/pkg/api/v1.TcpGateway")); err != nil {
		return 0, err
	}

	for _, v := range m.GetTcpHosts() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	if h, ok := interface{}(m.GetOptions()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Options")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetOptions(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Options")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *HybridGateway) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gateway.solo.io.github.com/solo-io/gloo/projects/gateway/pkg/api/v1.HybridGateway")); err != nil {
		return 0, err
	}

	for _, v := range m.GetMatchedGateways() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	if h, ok := interface{}(m.GetDelegatedHttpGateways()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("DelegatedHttpGateways")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetDelegatedHttpGateways(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("DelegatedHttpGateways")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetDelegatedTcpGateways()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("DelegatedTcpGateways")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetDelegatedTcpGateways(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("DelegatedTcpGateways")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *DelegatedHttpGateway) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gateway.solo.io.github.com/solo-io/gloo/projects/gateway/pkg/api/v1.DelegatedHttpGateway")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetPreventChildOverrides())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetHttpConnectionManagerSettings()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("HttpConnectionManagerSettings")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetHttpConnectionManagerSettings(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("HttpConnectionManagerSettings")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetSslConfig()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("SslConfig")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetSslConfig(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("SslConfig")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	switch m.SelectionType.(type) {

	case *DelegatedHttpGateway_Ref:

		if h, ok := interface{}(m.GetRef()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Ref")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetRef(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Ref")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *DelegatedHttpGateway_Selector:

		if h, ok := interface{}(m.GetSelector()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Selector")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetSelector(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Selector")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *DelegatedTcpGateway) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gateway.solo.io.github.com/solo-io/gloo/projects/gateway/pkg/api/v1.DelegatedTcpGateway")); err != nil {
		return 0, err
	}

	switch m.SelectionType.(type) {

	case *DelegatedTcpGateway_Ref:

		if h, ok := interface{}(m.GetRef()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Ref")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetRef(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Ref")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *DelegatedTcpGateway_Selector:

		if h, ok := interface{}(m.GetSelector()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Selector")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetSelector(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Selector")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *MatchedGateway) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gateway.solo.io.github.com/solo-io/gloo/projects/gateway/pkg/api/v1.MatchedGateway")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetMatcher()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Matcher")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMatcher(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Matcher")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	switch m.GatewayType.(type) {

	case *MatchedGateway_HttpGateway:

		if h, ok := interface{}(m.GetHttpGateway()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("HttpGateway")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetHttpGateway(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("HttpGateway")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *MatchedGateway_TcpGateway:

		if h, ok := interface{}(m.GetTcpGateway()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("TcpGateway")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetTcpGateway(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("TcpGateway")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *Matcher) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gateway.solo.io.github.com/solo-io/gloo/projects/gateway/pkg/api/v1.Matcher")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetSslConfig()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("SslConfig")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetSslConfig(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("SslConfig")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	for _, v := range m.GetSourcePrefixRanges() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	for _, v := range m.GetPrefixRanges() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetDestinationPort())
	if err != nil {
		return 0, err
	}

	for _, v := range m.GetPassthroughCipherSuites() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}
