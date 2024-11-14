// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/external/envoy/config/trace/v3/zipkin.proto

package v3

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"hash/fnv"
	"strconv"

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
	_ = strconv.Itoa
	_ = hashstructure.Hash
	_ = new(safe_hasher.SafeHasher)
)

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *ZipkinConfig) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.trace.v3.github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/config/trace/v3.ZipkinConfig")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("CollectorEndpoint")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetCollectorEndpoint())); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetTraceId_128Bit()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("TraceId_128Bit")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetTraceId_128Bit(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("TraceId_128Bit")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetSharedSpanContext()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("SharedSpanContext")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetSharedSpanContext(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("SharedSpanContext")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte("CollectorEndpointVersion")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetCollectorEndpointVersion())
	if err != nil {
		return 0, err
	}

	switch m.CollectorCluster.(type) {

	case *ZipkinConfig_CollectorUpstreamRef:

		if h, ok := interface{}(m.GetCollectorUpstreamRef()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("CollectorUpstreamRef")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetCollectorUpstreamRef(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("CollectorUpstreamRef")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *ZipkinConfig_ClusterName:

		if _, err = hasher.Write([]byte("ClusterName")); err != nil {
			return 0, err
		}
		if _, err = hasher.Write([]byte(m.GetClusterName())); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}
