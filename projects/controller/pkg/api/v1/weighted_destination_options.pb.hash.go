// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/v1/weighted_destination_options.proto

package v1

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
func (m *WeightedDestinationOptions) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1.WeightedDestinationOptions")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetHeaderManipulation()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("HeaderManipulation")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetHeaderManipulation(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("HeaderManipulation")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetTransformations()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Transformations")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetTransformations(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Transformations")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetExtensions()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Extensions")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetExtensions(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Extensions")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetExtauth()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Extauth")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetExtauth(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Extauth")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetBufferPerRoute()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("BufferPerRoute")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetBufferPerRoute(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("BufferPerRoute")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetCsrf()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Csrf")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetCsrf(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Csrf")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetStagedTransformations()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("StagedTransformations")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetStagedTransformations(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("StagedTransformations")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}
