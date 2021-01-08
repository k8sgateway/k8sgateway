// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/config/trace/v3/zipkin.proto

package v3

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"hash/fnv"

	"github.com/mitchellh/hashstructure"
	safe_hasher "github.com/solo-io/protoc-gen-ext/pkg/hasher"
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
func (m *ZipkinConfig) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.trace.v3.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/config/trace/v3.ZipkinConfig")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetCollectorEndpoint())); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetTraceId_128Bit())
	if err != nil {
		return 0, err
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

		if _, err = hasher.Write([]byte(m.GetClusterName())); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}
