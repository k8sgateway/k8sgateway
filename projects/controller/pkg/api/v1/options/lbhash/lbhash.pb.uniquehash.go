// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/v1/options/lbhash/lbhash.proto

package lbhash

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
func (m *RouteActionHashConfig) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("lbhash.options.gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1/options/lbhash.RouteActionHashConfig")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("HashPolicies")); err != nil {
		return 0, err
	}
	for i, v := range m.GetHashPolicies() {
		if _, err = hasher.Write([]byte(strconv.Itoa(i))); err != nil {
			return 0, err
		}

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("v")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("v")); err != nil {
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

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *Cookie) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("lbhash.options.gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1/options/lbhash.Cookie")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Name")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetName())); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetTtl()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Ttl")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetTtl(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Ttl")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte("Path")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetPath())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *HashPolicy) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("lbhash.options.gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1/options/lbhash.HashPolicy")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Terminal")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetTerminal())
	if err != nil {
		return 0, err
	}

	switch m.KeyType.(type) {

	case *HashPolicy_Header:

		if _, err = hasher.Write([]byte("Header")); err != nil {
			return 0, err
		}
		if _, err = hasher.Write([]byte(m.GetHeader())); err != nil {
			return 0, err
		}

	case *HashPolicy_Cookie:

		if h, ok := interface{}(m.GetCookie()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Cookie")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetCookie(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Cookie")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *HashPolicy_SourceIp:

		if _, err = hasher.Write([]byte("SourceIp")); err != nil {
			return 0, err
		}
		err = binary.Write(hasher, binary.LittleEndian, m.GetSourceIp())
		if err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}
