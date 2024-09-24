// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/connection_limit/connection_limit.proto

package connection_limit

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
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *ConnectionLimit) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("connection_limit.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/connection_limit.ConnectionLimit")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetMaxActiveConnections()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("MaxActiveConnections")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMaxActiveConnections(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("MaxActiveConnections")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetDelayBeforeClose()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("DelayBeforeClose")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetDelayBeforeClose(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("DelayBeforeClose")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}
