// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/core/matchers/matchers.proto

package matchers

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
func (m *Matcher) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("matchers.core.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/core/matchers.Matcher")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetCaseSensitive()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("CaseSensitive")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetCaseSensitive(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("CaseSensitive")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte("Headers")); err != nil {
		return 0, err
	}
	for i, v := range m.GetHeaders() {
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

	if _, err = hasher.Write([]byte("QueryParameters")); err != nil {
		return 0, err
	}
	for i, v := range m.GetQueryParameters() {
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

	if _, err = hasher.Write([]byte("Methods")); err != nil {
		return 0, err
	}
	for i, v := range m.GetMethods() {
		if _, err = hasher.Write([]byte(strconv.Itoa(i))); err != nil {
			return 0, err
		}

		if _, err = hasher.Write([]byte("v")); err != nil {
			return 0, err
		}
		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	switch m.PathSpecifier.(type) {

	case *Matcher_Prefix:

		if _, err = hasher.Write([]byte("Prefix")); err != nil {
			return 0, err
		}
		if _, err = hasher.Write([]byte(m.GetPrefix())); err != nil {
			return 0, err
		}

	case *Matcher_Exact:

		if _, err = hasher.Write([]byte("Exact")); err != nil {
			return 0, err
		}
		if _, err = hasher.Write([]byte(m.GetExact())); err != nil {
			return 0, err
		}

	case *Matcher_Regex:

		if _, err = hasher.Write([]byte("Regex")); err != nil {
			return 0, err
		}
		if _, err = hasher.Write([]byte(m.GetRegex())); err != nil {
			return 0, err
		}

	case *Matcher_ConnectMatcher_:

		if h, ok := interface{}(m.GetConnectMatcher()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("ConnectMatcher")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetConnectMatcher(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("ConnectMatcher")); err != nil {
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
func (m *HeaderMatcher) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("matchers.core.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/core/matchers.HeaderMatcher")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Name")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetName())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Value")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetValue())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Regex")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetRegex())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("InvertMatch")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetInvertMatch())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *QueryParameterMatcher) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("matchers.core.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/core/matchers.QueryParameterMatcher")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Name")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetName())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Value")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetValue())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Regex")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetRegex())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *Matcher_ConnectMatcher) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("matchers.core.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/core/matchers.Matcher_ConnectMatcher")); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}
