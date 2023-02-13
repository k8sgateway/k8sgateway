// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/dynamic_forward_proxy/dynamic_forward_proxy.proto

package dynamic_forward_proxy

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
func (m *FilterConfig) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("dfp.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/dynamic_forward_proxy.FilterConfig")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetDnsCacheConfig()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("DnsCacheConfig")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetDnsCacheConfig(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("DnsCacheConfig")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetSaveUpstreamAddress())
	if err != nil {
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

	return hasher.Sum64(), nil
}

// Hash function
func (m *DnsCacheCircuitBreakers) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("dfp.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/dynamic_forward_proxy.DnsCacheCircuitBreakers")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetMaxPendingRequests()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("MaxPendingRequests")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMaxPendingRequests(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("MaxPendingRequests")); err != nil {
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
func (m *DnsCacheConfig) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("dfp.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/dynamic_forward_proxy.DnsCacheConfig")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetDnsLookupFamily())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetDnsRefreshRate()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("DnsRefreshRate")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetDnsRefreshRate(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("DnsRefreshRate")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetHostTtl()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("HostTtl")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetHostTtl(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("HostTtl")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetMaxHosts()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("MaxHosts")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMaxHosts(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("MaxHosts")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetDnsFailureRefreshRate()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("DnsFailureRefreshRate")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetDnsFailureRefreshRate(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("DnsFailureRefreshRate")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetDnsCacheCircuitBreaker()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("DnsCacheCircuitBreaker")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetDnsCacheCircuitBreaker(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("DnsCacheCircuitBreaker")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	for _, v := range m.GetPreresolveHostnames() {

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

	if h, ok := interface{}(m.GetDnsQueryTimeout()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("DnsQueryTimeout")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetDnsQueryTimeout(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("DnsQueryTimeout")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	switch m.DnsCacheType.(type) {

	case *DnsCacheConfig_CaresDns:

		if h, ok := interface{}(m.GetCaresDns()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("CaresDns")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetCaresDns(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("CaresDns")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *DnsCacheConfig_AppleDns:

		if h, ok := interface{}(m.GetAppleDns()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("AppleDns")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetAppleDns(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("AppleDns")); err != nil {
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
func (m *RefreshRate) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("dfp.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/dynamic_forward_proxy.RefreshRate")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetBaseInterval()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("BaseInterval")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetBaseInterval(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("BaseInterval")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetMaxInterval()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("MaxInterval")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMaxInterval(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("MaxInterval")); err != nil {
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
func (m *PerRouteConfig) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("dfp.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/dynamic_forward_proxy.PerRouteConfig")); err != nil {
		return 0, err
	}

	switch m.HostRewriteSpecifier.(type) {

	case *PerRouteConfig_HostRewrite:

		if _, err = hasher.Write([]byte(m.GetHostRewrite())); err != nil {
			return 0, err
		}

	case *PerRouteConfig_AutoHostRewriteHeader:

		if _, err = hasher.Write([]byte(m.GetAutoHostRewriteHeader())); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *DnsResolverOptions) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("dfp.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/dynamic_forward_proxy.DnsResolverOptions")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetUseTcpForDnsLookups())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetNoDefaultSearchDomain())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *CaresDnsResolverConfig) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("dfp.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/dynamic_forward_proxy.CaresDnsResolverConfig")); err != nil {
		return 0, err
	}

	for _, v := range m.GetResolvers() {

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

	if h, ok := interface{}(m.GetDnsResolverOptions()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("DnsResolverOptions")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetDnsResolverOptions(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("DnsResolverOptions")); err != nil {
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
func (m *AppleDnsResolverConfig) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("dfp.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/dynamic_forward_proxy.AppleDnsResolverConfig")); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}
