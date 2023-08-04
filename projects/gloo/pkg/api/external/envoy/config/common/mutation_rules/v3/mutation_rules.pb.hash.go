// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/config/common/mutation_rules/v3/mutation_rules.proto

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
func (m *HeaderMutationRules) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.common.mutation_rules.v3.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/config/common/mutation_rules/v3.HeaderMutationRules")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetAllowAllRouting()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("AllowAllRouting")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetAllowAllRouting(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("AllowAllRouting")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetAllowEnvoy()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("AllowEnvoy")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetAllowEnvoy(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("AllowEnvoy")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetDisallowSystem()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("DisallowSystem")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetDisallowSystem(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("DisallowSystem")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetDisallowAll()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("DisallowAll")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetDisallowAll(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("DisallowAll")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetAllowExpression()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("AllowExpression")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetAllowExpression(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("AllowExpression")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetDisallowExpression()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("DisallowExpression")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetDisallowExpression(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("DisallowExpression")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetDisallowIsError()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("DisallowIsError")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetDisallowIsError(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("DisallowIsError")); err != nil {
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
func (m *HeaderMutation) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.common.mutation_rules.v3.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/config/common/mutation_rules/v3.HeaderMutation")); err != nil {
		return 0, err
	}

	switch m.Action.(type) {

	case *HeaderMutation_Remove:

		if _, err = hasher.Write([]byte(m.GetRemove())); err != nil {
			return 0, err
		}

	case *HeaderMutation_Append:

		if h, ok := interface{}(m.GetAppend()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Append")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetAppend(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Append")); err != nil {
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
