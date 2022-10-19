// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/wasm/wasm.proto

package wasm

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
func (m *PluginSource) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("wasm.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/wasm.PluginSource")); err != nil {
		return 0, err
	}

	for _, v := range m.GetFilters() {

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
func (m *WasmFilter) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("wasm.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/wasm.WasmFilter")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetConfig()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Config")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetConfig(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Config")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetFilterStage()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("FilterStage")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetFilterStage(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("FilterStage")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte(m.GetName())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetRootId())); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetVmType())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetFailOpen())
	if err != nil {
		return 0, err
	}

	switch m.Src.(type) {

	case *WasmFilter_Image:

		if _, err = hasher.Write([]byte(m.GetImage())); err != nil {
			return 0, err
		}

	case *WasmFilter_FilePath:

		if _, err = hasher.Write([]byte(m.GetFilePath())); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *FilterStage) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("wasm.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/wasm.FilterStage")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetStage())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetPredicate())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}
