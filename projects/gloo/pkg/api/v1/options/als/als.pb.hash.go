// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/als/als.proto

package als

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
func (m *AccessLoggingService) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("als.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/als.AccessLoggingService")); err != nil {
		return 0, err
	}

	for _, v := range m.GetAccessLog() {

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
func (m *AccessLog) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("als.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/als.AccessLog")); err != nil {
		return 0, err
	}

	switch m.OutputDestination.(type) {

	case *AccessLog_FileSink:

		if h, ok := interface{}(m.GetFileSink()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("FileSink")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetFileSink(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("FileSink")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *AccessLog_GrpcService:

		if h, ok := interface{}(m.GetGrpcService()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("GrpcService")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetGrpcService(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("GrpcService")); err != nil {
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
func (m *FileSink) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("als.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/als.FileSink")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetPath())); err != nil {
		return 0, err
	}

	switch m.OutputFormat.(type) {

	case *FileSink_StringFormat:

		if _, err = hasher.Write([]byte(m.GetStringFormat())); err != nil {
			return 0, err
		}

	case *FileSink_JsonFormat:

		if h, ok := interface{}(m.GetJsonFormat()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("JsonFormat")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetJsonFormat(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("JsonFormat")); err != nil {
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
func (m *GrpcService) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("als.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/als.GrpcService")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetLogName())); err != nil {
		return 0, err
	}

	for _, v := range m.GetAdditionalRequestHeadersToLog() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	for _, v := range m.GetAdditionalResponseHeadersToLog() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	for _, v := range m.GetAdditionalResponseTrailersToLog() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	switch m.ServiceRef.(type) {

	case *GrpcService_StaticClusterName:

		if _, err = hasher.Write([]byte(m.GetStaticClusterName())); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}
