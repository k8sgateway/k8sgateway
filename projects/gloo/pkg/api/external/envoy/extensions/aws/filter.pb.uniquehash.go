// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/extensions/aws/filter.proto

package aws

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
func (m *AWSLambdaPerRoute) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.config.filter.http.aws_lambda.v2.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/aws.AWSLambdaPerRoute")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Name")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetName())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Qualifier")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetQualifier())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Async")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetAsync())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetEmptyBodyOverride()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("EmptyBodyOverride")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetEmptyBodyOverride(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("EmptyBodyOverride")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte("UnwrapAsAlb")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetUnwrapAsAlb())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetTransformerConfig()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("TransformerConfig")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetTransformerConfig(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("TransformerConfig")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetRequestTransformerConfig()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("RequestTransformerConfig")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRequestTransformerConfig(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("RequestTransformerConfig")); err != nil {
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
func (m *AWSLambdaProtocolExtension) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.config.filter.http.aws_lambda.v2.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/aws.AWSLambdaProtocolExtension")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Host")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetHost())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Region")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetRegion())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("AccessKey")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetAccessKey())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("SecretKey")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetSecretKey())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("SessionToken")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetSessionToken())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("RoleArn")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetRoleArn())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("DisableRoleChaining")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetDisableRoleChaining())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *AWSLambdaConfig) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.config.filter.http.aws_lambda.v2.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/aws.AWSLambdaConfig")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("PropagateOriginalRouting")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetPropagateOriginalRouting())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetCredentialRefreshDelay()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("CredentialRefreshDelay")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetCredentialRefreshDelay(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("CredentialRefreshDelay")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	switch m.CredentialsFetcher.(type) {

	case *AWSLambdaConfig_UseDefaultCredentials:

		if h, ok := interface{}(m.GetUseDefaultCredentials()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("UseDefaultCredentials")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetUseDefaultCredentials(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("UseDefaultCredentials")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *AWSLambdaConfig_ServiceAccountCredentials_:

		if h, ok := interface{}(m.GetServiceAccountCredentials()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("ServiceAccountCredentials")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetServiceAccountCredentials(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("ServiceAccountCredentials")); err != nil {
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
func (m *ApiGatewayTransformation) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.config.filter.http.aws_lambda.v2.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/aws.ApiGatewayTransformation")); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *AWSLambdaConfig_ServiceAccountCredentials) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.config.filter.http.aws_lambda.v2.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/aws.AWSLambdaConfig_ServiceAccountCredentials")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Cluster")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetCluster())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Uri")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetUri())); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetTimeout()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Timeout")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetTimeout(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Timeout")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte("Region")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetRegion())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}
