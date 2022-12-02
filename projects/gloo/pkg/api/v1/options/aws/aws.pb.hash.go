// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/aws/aws.proto

package aws

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
func (m *UpstreamSpec) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("aws.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/aws.UpstreamSpec")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetRegion())); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetSecretRef()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("SecretRef")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetSecretRef(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("SecretRef")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	for _, v := range m.GetLambdaFunctions() {

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

	if _, err = hasher.Write([]byte(m.GetRoleArn())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetAwsAccountId())); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetDisableRoleChaining())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetDestinationOverrides()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("DestinationOverrides")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetDestinationOverrides(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("DestinationOverrides")); err != nil {
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
func (m *LambdaFunctionSpec) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("aws.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/aws.LambdaFunctionSpec")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetLogicalName())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetLambdaFunctionName())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetQualifier())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *DestinationSpec) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("aws.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/aws.DestinationSpec")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetLogicalName())); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetInvocationStyle())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetRequestTransformation())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetResponseTransformation())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetUnwrapAsAlb())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetUnwrapAsApiGateway())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetWrapAsApiGateway())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}
