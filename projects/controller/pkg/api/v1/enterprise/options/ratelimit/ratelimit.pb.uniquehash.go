// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/v1/enterprise/options/ratelimit/ratelimit.proto

package ratelimit

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
func (m *IngressRateLimit) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ratelimit.options.gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1/enterprise/options/ratelimit.IngressRateLimit")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetAuthorizedLimits()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("AuthorizedLimits")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetAuthorizedLimits(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("AuthorizedLimits")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetAnonymousLimits()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("AnonymousLimits")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetAnonymousLimits(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("AnonymousLimits")); err != nil {
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
	if _, err = hasher.Write([]byte("ratelimit.options.gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1/enterprise/options/ratelimit.Settings")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetRatelimitServerRef()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("RatelimitServerRef")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRatelimitServerRef(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("RatelimitServerRef")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetRequestTimeout()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("RequestTimeout")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRequestTimeout(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("RequestTimeout")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte("DenyOnFail")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetDenyOnFail())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("EnableXRatelimitHeaders")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetEnableXRatelimitHeaders())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("RateLimitBeforeAuth")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetRateLimitBeforeAuth())
	if err != nil {
		return 0, err
	}

	switch m.ServiceType.(type) {

	case *Settings_GrpcService:

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

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *GrpcService) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ratelimit.options.gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1/enterprise/options/ratelimit.GrpcService")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Authority")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetAuthority())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *ServiceSettings) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ratelimit.options.gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1/enterprise/options/ratelimit.ServiceSettings")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Descriptors")); err != nil {
		return 0, err
	}
	for i, v := range m.GetDescriptors() {
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

	if _, err = hasher.Write([]byte("SetDescriptors")); err != nil {
		return 0, err
	}
	for i, v := range m.GetSetDescriptors() {
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

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *RateLimitConfigRefs) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ratelimit.options.gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1/enterprise/options/ratelimit.RateLimitConfigRefs")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Refs")); err != nil {
		return 0, err
	}
	for i, v := range m.GetRefs() {
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

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *RateLimitConfigRef) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ratelimit.options.gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1/enterprise/options/ratelimit.RateLimitConfigRef")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Name")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetName())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Namespace")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetNamespace())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *RateLimitVhostExtension) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ratelimit.options.gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1/enterprise/options/ratelimit.RateLimitVhostExtension")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("RateLimits")); err != nil {
		return 0, err
	}
	for i, v := range m.GetRateLimits() {
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

	if h, ok := interface{}(m.GetLocalRatelimit()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("LocalRatelimit")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetLocalRatelimit(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("LocalRatelimit")); err != nil {
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
func (m *RateLimitRouteExtension) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ratelimit.options.gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1/enterprise/options/ratelimit.RateLimitRouteExtension")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("IncludeVhRateLimits")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetIncludeVhRateLimits())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("RateLimits")); err != nil {
		return 0, err
	}
	for i, v := range m.GetRateLimits() {
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

	if h, ok := interface{}(m.GetLocalRatelimit()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("LocalRatelimit")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetLocalRatelimit(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("LocalRatelimit")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}
