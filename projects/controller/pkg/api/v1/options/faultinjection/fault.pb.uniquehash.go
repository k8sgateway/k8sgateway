// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/v1/options/faultinjection/fault.proto

package faultinjection

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
func (m *RouteAbort) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fault.options.gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1/options/faultinjection.RouteAbort")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Percentage")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetPercentage())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("HttpStatus")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetHttpStatus())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *RouteDelay) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fault.options.gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1/options/faultinjection.RouteDelay")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Percentage")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetPercentage())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetFixedDelay()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("FixedDelay")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetFixedDelay(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("FixedDelay")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *RouteFaults) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fault.options.gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1/options/faultinjection.RouteFaults")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetAbort()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Abort")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetAbort(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Abort")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetDelay()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Delay")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetDelay(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Delay")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}
