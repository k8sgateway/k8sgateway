// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/retries/retries.proto

package retries

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
func (m *RetryBackOff) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("retries.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/retries.RetryBackOff")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetBaseInterval()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("BaseInterval")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetBaseInterval(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("BaseInterval")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetMaxInterval()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("MaxInterval")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMaxInterval(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("MaxInterval")); err != nil {
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
func (m *RetryPolicy) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("retries.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/retries.RetryPolicy")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetRetryOn())); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetNumRetries())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetPerTryTimeout()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("PerTryTimeout")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetPerTryTimeout(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("PerTryTimeout")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetRetryBackOff()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("RetryBackOff")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRetryBackOff(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("RetryBackOff")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	switch m.PriorityPredicate.(type) {

	case *RetryPolicy_PreviousPriorities_:

		if h, ok := interface{}(m.GetPreviousPriorities()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("PreviousPriorities")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetPreviousPriorities(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("PreviousPriorities")); err != nil {
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
func (m *RetryPolicy_PreviousPriorities) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("retries.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/retries.RetryPolicy_PreviousPriorities")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetUpdateFrequency()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("UpdateFrequency")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetUpdateFrequency(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("UpdateFrequency")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}
