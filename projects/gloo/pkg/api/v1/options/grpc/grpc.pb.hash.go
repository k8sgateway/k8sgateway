// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/grpc/grpc.proto

package grpc

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
func (m *ServiceSpec) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("grpc.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/grpc.ServiceSpec")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write(m.GetDescriptors()); err != nil {
		return 0, err
	}

	for _, v := range m.GetGrpcServices() {

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

	return hasher.Sum64(), nil
}

// Hash function
func (m *DestinationSpec) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("grpc.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/grpc.DestinationSpec")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetPackage())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetService())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetFunction())); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetParameters()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Parameters")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetParameters(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Parameters")); err != nil {
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
func (m *ServiceSpec_GrpcService) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("grpc.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/grpc.ServiceSpec_GrpcService")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetPackageName())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetServiceName())); err != nil {
		return 0, err
	}

	for _, v := range m.GetFunctionNames() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}
