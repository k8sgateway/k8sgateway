// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/external/envoy/api/v2/cluster/outlier_detection.proto

package cluster

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
func (m *OutlierDetection) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.api.v2.cluster.github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/api/v2/cluster.OutlierDetection")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetConsecutive_5Xx()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Consecutive_5Xx")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetConsecutive_5Xx(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Consecutive_5Xx")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetInterval()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Interval")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetInterval(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Interval")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetBaseEjectionTime()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("BaseEjectionTime")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetBaseEjectionTime(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("BaseEjectionTime")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetMaxEjectionPercent()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("MaxEjectionPercent")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMaxEjectionPercent(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("MaxEjectionPercent")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetEnforcingConsecutive_5Xx()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("EnforcingConsecutive_5Xx")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetEnforcingConsecutive_5Xx(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("EnforcingConsecutive_5Xx")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetEnforcingSuccessRate()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("EnforcingSuccessRate")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetEnforcingSuccessRate(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("EnforcingSuccessRate")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetSuccessRateMinimumHosts()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("SuccessRateMinimumHosts")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetSuccessRateMinimumHosts(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("SuccessRateMinimumHosts")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetSuccessRateRequestVolume()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("SuccessRateRequestVolume")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetSuccessRateRequestVolume(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("SuccessRateRequestVolume")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetSuccessRateStdevFactor()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("SuccessRateStdevFactor")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetSuccessRateStdevFactor(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("SuccessRateStdevFactor")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetConsecutiveGatewayFailure()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("ConsecutiveGatewayFailure")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetConsecutiveGatewayFailure(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("ConsecutiveGatewayFailure")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetEnforcingConsecutiveGatewayFailure()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("EnforcingConsecutiveGatewayFailure")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetEnforcingConsecutiveGatewayFailure(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("EnforcingConsecutiveGatewayFailure")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte("SplitExternalLocalOriginErrors")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetSplitExternalLocalOriginErrors())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetConsecutiveLocalOriginFailure()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("ConsecutiveLocalOriginFailure")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetConsecutiveLocalOriginFailure(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("ConsecutiveLocalOriginFailure")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetEnforcingConsecutiveLocalOriginFailure()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("EnforcingConsecutiveLocalOriginFailure")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetEnforcingConsecutiveLocalOriginFailure(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("EnforcingConsecutiveLocalOriginFailure")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetEnforcingLocalOriginSuccessRate()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("EnforcingLocalOriginSuccessRate")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetEnforcingLocalOriginSuccessRate(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("EnforcingLocalOriginSuccessRate")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}
