// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/kubernetes/kubernetes.proto

package kubernetes

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
func (m *UpstreamSpec) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("kubernetes.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/kubernetes.UpstreamSpec")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("ServiceName")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetServiceName())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("ServiceNamespace")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetServiceNamespace())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("ServicePort")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetServicePort())
	if err != nil {
		return 0, err
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetSelector() {
			innerHash.Reset()

			if _, err = innerHash.Write([]byte("v")); err != nil {
				return 0, err
			}
			if _, err = innerHash.Write([]byte(v)); err != nil {
				return 0, err
			}

			if _, err = innerHash.Write([]byte("k")); err != nil {
				return 0, err
			}
			if _, err = innerHash.Write([]byte(k)); err != nil {
				return 0, err
			}

			result = result ^ innerHash.Sum64()
		}
		err = binary.Write(hasher, binary.LittleEndian, result)
		if err != nil {
			return 0, err
		}

	}

	if h, ok := interface{}(m.GetServiceSpec()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("ServiceSpec")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetServiceSpec(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("ServiceSpec")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetSubsetSpec()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("SubsetSpec")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetSubsetSpec(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("SubsetSpec")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}
