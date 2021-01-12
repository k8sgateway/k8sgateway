// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/type/metadata/v3/metadata.proto

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
func (m *MetadataKey) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.type.metadata.v3.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/type/metadata/v3.MetadataKey")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetKey())); err != nil {
		return 0, err
	}

	for _, v := range m.GetPath() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *MetadataKind) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.type.metadata.v3.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/type/metadata/v3.MetadataKind")); err != nil {
		return 0, err
	}

	switch m.Kind.(type) {

	case *MetadataKind_Request_:

		if h, ok := interface{}(m.GetRequest()).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(m.GetRequest(), nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	case *MetadataKind_Route_:

		if h, ok := interface{}(m.GetRoute()).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(m.GetRoute(), nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	case *MetadataKind_Cluster_:

		if h, ok := interface{}(m.GetCluster()).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(m.GetCluster(), nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	case *MetadataKind_Host_:

		if h, ok := interface{}(m.GetHost()).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(m.GetHost(), nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *MetadataKey_PathSegment) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.type.metadata.v3.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/type/metadata/v3.MetadataKey_PathSegment")); err != nil {
		return 0, err
	}

	switch m.Segment.(type) {

	case *MetadataKey_PathSegment_Key:

		if _, err = hasher.Write([]byte(m.GetKey())); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *MetadataKind_Request) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.type.metadata.v3.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/type/metadata/v3.MetadataKind_Request")); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *MetadataKind_Route) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.type.metadata.v3.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/type/metadata/v3.MetadataKind_Route")); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *MetadataKind_Cluster) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.type.metadata.v3.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/type/metadata/v3.MetadataKind_Cluster")); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *MetadataKind_Host) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.type.metadata.v3.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/type/metadata/v3.MetadataKind_Host")); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}
