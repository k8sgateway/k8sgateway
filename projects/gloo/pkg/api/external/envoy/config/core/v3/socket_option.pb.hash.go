// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/config/core/v3/socket_option.proto

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
func (m *SocketOption) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.core.v3.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/config/core/v3.SocketOption")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetDescription())); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetLevel())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetName())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetState())
	if err != nil {
		return 0, err
	}

	switch m.Value.(type) {

	case *SocketOption_IntValue:

		err = binary.Write(hasher, binary.LittleEndian, m.GetIntValue())
		if err != nil {
			return 0, err
		}

	case *SocketOption_BufValue:

		if _, err = hasher.Write(m.GetBufValue()); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}
