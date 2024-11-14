// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/external/envoy/extensions/filters/http/ext_proc/v3/processing_mode.proto

package v3

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
func (m *ProcessingMode) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.extensions.filters.http.ext_proc.v3.github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/extensions/filters/http/ext_proc/v3.ProcessingMode")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetRequestHeaderMode())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetResponseHeaderMode())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetRequestBodyMode())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetResponseBodyMode())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetRequestTrailerMode())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetResponseTrailerMode())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}
