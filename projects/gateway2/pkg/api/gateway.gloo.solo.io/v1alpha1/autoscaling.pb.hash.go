// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gateway2/api/v1alpha1/autoscaling.proto

package v1alpha1

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
func (m *Autoscaling) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gateway.gloo.solo.io.github.com/solo-io/gloo/projects/gateway2/pkg/api/gateway.gloo.solo.io/v1alpha1.Autoscaling")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetHorizontalPodAutoscaler()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("HorizontalPodAutoscaler")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetHorizontalPodAutoscaler(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("HorizontalPodAutoscaler")); err != nil {
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
func (m *HorizontalPodAutoscaler) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gateway.gloo.solo.io.github.com/solo-io/gloo/projects/gateway2/pkg/api/gateway.gloo.solo.io/v1alpha1.HorizontalPodAutoscaler")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetMinReplicas()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("MinReplicas")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMinReplicas(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("MinReplicas")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetMaxReplicas()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("MaxReplicas")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMaxReplicas(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("MaxReplicas")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetTargetCpuUtilizationPercentage()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("TargetCpuUtilizationPercentage")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetTargetCpuUtilizationPercentage(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("TargetCpuUtilizationPercentage")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetTargetMemoryUtilizationPercentage()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("TargetMemoryUtilizationPercentage")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetTargetMemoryUtilizationPercentage(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("TargetMemoryUtilizationPercentage")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}
