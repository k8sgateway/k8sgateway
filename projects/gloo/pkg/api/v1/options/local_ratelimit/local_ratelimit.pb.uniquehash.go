// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/local_ratelimit/local_ratelimit.proto

package local_ratelimit

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
func (m *TokenBucket) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("local_ratelimit.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/local_ratelimit.TokenBucket")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("MaxTokens")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetMaxTokens())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetTokensPerFill()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("TokensPerFill")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetTokensPerFill(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("TokensPerFill")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetFillInterval()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("FillInterval")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetFillInterval(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("FillInterval")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *Settings) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("local_ratelimit.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/local_ratelimit.Settings")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetDefaultLimit()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("DefaultLimit")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetDefaultLimit(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("DefaultLimit")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetLocalRateLimitPerDownstreamConnection()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("LocalRateLimitPerDownstreamConnection")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetLocalRateLimitPerDownstreamConnection(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("LocalRateLimitPerDownstreamConnection")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetEnableXRatelimitHeaders()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("EnableXRatelimitHeaders")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetEnableXRatelimitHeaders(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("EnableXRatelimitHeaders")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}
