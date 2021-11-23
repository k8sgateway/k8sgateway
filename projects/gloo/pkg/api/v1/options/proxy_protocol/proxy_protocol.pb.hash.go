// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/proxy_protocol/proxy_protocol.proto

package proxy_protocol

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
func (m *ProxyProtocol) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("proxy_protocol.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/proxy_protocol.ProxyProtocol")); err != nil {
		return 0, err
	}

	for _, v := range m.GetRules() {

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

	err = binary.Write(hasher, binary.LittleEndian, m.GetAllowRequestsWithoutProxyProtocol())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ProxyProtocol_KeyValuePair) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("proxy_protocol.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/proxy_protocol.ProxyProtocol_KeyValuePair")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetMetadataNamespace())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetKey())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ProxyProtocol_Rule) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("proxy_protocol.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/proxy_protocol.ProxyProtocol_Rule")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetTlvType())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetOnTlvPresent()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("OnTlvPresent")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetOnTlvPresent(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("OnTlvPresent")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}
