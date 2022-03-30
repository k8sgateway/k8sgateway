// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/options/graphql/v1alpha1/graphql.proto

package v1alpha1

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"hash/fnv"

	"github.com/mitchellh/hashstructure"
	safe_hasher "github.com/solo-io/protoc-gen-ext/pkg/hasher"
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
func (m *RequestTemplate) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("graphql.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/graphql/v1alpha1.RequestTemplate")); err != nil {
		return 0, err
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetHeaders() {
			innerHash.Reset()

			if _, err = innerHash.Write([]byte(v)); err != nil {
				return 0, err
			}

			if _, err = innerHash.Write([]byte(k)); err != nil {
				return 0, err
			}

			result = result ^ innerHash.Sum64()
		}
		err = binary.Write(hasher, binary.LittleEndian, result)
		if err != nil {
			return 0, err
		}

	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetQueryParams() {
			innerHash.Reset()

			if _, err = innerHash.Write([]byte(v)); err != nil {
				return 0, err
			}

			if _, err = innerHash.Write([]byte(k)); err != nil {
				return 0, err
			}

			result = result ^ innerHash.Sum64()
		}
		err = binary.Write(hasher, binary.LittleEndian, result)
		if err != nil {
			return 0, err
		}

	}

	if h, ok := interface{}(m.GetBody()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Body")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetBody(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Body")); err != nil {
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
func (m *ResponseTemplate) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("graphql.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/graphql/v1alpha1.ResponseTemplate")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetResultRoot())); err != nil {
		return 0, err
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetSetters() {
			innerHash.Reset()

			if _, err = innerHash.Write([]byte(v)); err != nil {
				return 0, err
			}

			if _, err = innerHash.Write([]byte(k)); err != nil {
				return 0, err
			}

			result = result ^ innerHash.Sum64()
		}
		err = binary.Write(hasher, binary.LittleEndian, result)
		if err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *GrpcRequestTemplate) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("graphql.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/graphql/v1alpha1.GrpcRequestTemplate")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetOutgoingMessageJson()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("OutgoingMessageJson")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetOutgoingMessageJson(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("OutgoingMessageJson")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte(m.GetServiceName())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetMethodName())); err != nil {
		return 0, err
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetRequestMetadata() {
			innerHash.Reset()

			if _, err = innerHash.Write([]byte(v)); err != nil {
				return 0, err
			}

			if _, err = innerHash.Write([]byte(k)); err != nil {
				return 0, err
			}

			result = result ^ innerHash.Sum64()
		}
		err = binary.Write(hasher, binary.LittleEndian, result)
		if err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *RESTResolver) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("graphql.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/graphql/v1alpha1.RESTResolver")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetUpstreamRef()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("UpstreamRef")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetUpstreamRef(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("UpstreamRef")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetRequest()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Request")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRequest(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Request")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetResponse()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Response")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetResponse(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Response")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte(m.GetSpanName())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *GrpcDescriptorRegistry) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("graphql.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/graphql/v1alpha1.GrpcDescriptorRegistry")); err != nil {
		return 0, err
	}

	switch m.DescriptorSet.(type) {

	case *GrpcDescriptorRegistry_ProtoDescriptor:

		if _, err = hasher.Write([]byte(m.GetProtoDescriptor())); err != nil {
			return 0, err
		}

	case *GrpcDescriptorRegistry_ProtoDescriptorBin:

		if _, err = hasher.Write(m.GetProtoDescriptorBin()); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *GrpcResolver) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("graphql.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/graphql/v1alpha1.GrpcResolver")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetUpstreamRef()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("UpstreamRef")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetUpstreamRef(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("UpstreamRef")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetRequestTransform()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("RequestTransform")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRequestTransform(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("RequestTransform")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte(m.GetSpanName())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *StitchedSchema) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("graphql.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/graphql/v1alpha1.StitchedSchema")); err != nil {
		return 0, err
	}

	for _, v := range m.GetSubschemas() {

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

	return hasher.Sum64(), nil
}

// Hash function
func (m *MockResolver) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("graphql.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/graphql/v1alpha1.MockResolver")); err != nil {
		return 0, err
	}

	switch m.Response.(type) {

	case *MockResolver_SyncResponse:

		if h, ok := interface{}(m.GetSyncResponse()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("SyncResponse")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetSyncResponse(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("SyncResponse")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *MockResolver_AsyncResponse_:

		if h, ok := interface{}(m.GetAsyncResponse()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("AsyncResponse")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetAsyncResponse(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("AsyncResponse")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *MockResolver_ErrorResponse:

		if _, err = hasher.Write([]byte(m.GetErrorResponse())); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *Resolution) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("graphql.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/graphql/v1alpha1.Resolution")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetStatPrefix()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("StatPrefix")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetStatPrefix(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("StatPrefix")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	switch m.Resolver.(type) {

	case *Resolution_RestResolver:

		if h, ok := interface{}(m.GetRestResolver()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("RestResolver")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetRestResolver(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("RestResolver")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *Resolution_GrpcResolver:

		if h, ok := interface{}(m.GetGrpcResolver()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("GrpcResolver")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetGrpcResolver(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("GrpcResolver")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *Resolution_MockResolver:

		if h, ok := interface{}(m.GetMockResolver()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("MockResolver")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetMockResolver(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("MockResolver")); err != nil {
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
func (m *GraphQLApi) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("graphql.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/graphql/v1alpha1.GraphQLApi")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetMetadata()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Metadata")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMetadata(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Metadata")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetStatPrefix()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("StatPrefix")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetStatPrefix(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("StatPrefix")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetPersistedQueryCacheConfig()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("PersistedQueryCacheConfig")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetPersistedQueryCacheConfig(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("PersistedQueryCacheConfig")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	for _, v := range m.GetAllowedQueryHashes() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	if h, ok := interface{}(m.GetOptions()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Options")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetOptions(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Options")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	switch m.Schema.(type) {

	case *GraphQLApi_ExecutableSchema:

		if h, ok := interface{}(m.GetExecutableSchema()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("ExecutableSchema")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetExecutableSchema(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("ExecutableSchema")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *GraphQLApi_StitchedSchema:

		if h, ok := interface{}(m.GetStitchedSchema()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("StitchedSchema")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetStitchedSchema(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("StitchedSchema")); err != nil {
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
func (m *PersistedQueryCacheConfig) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("graphql.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/graphql/v1alpha1.PersistedQueryCacheConfig")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetCacheSize())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ExecutableSchema) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("graphql.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/graphql/v1alpha1.ExecutableSchema")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetSchemaDefinition())); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetExecutor()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Executor")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetExecutor(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Executor")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetGrpcDescriptorRegistry()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("GrpcDescriptorRegistry")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetGrpcDescriptorRegistry(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("GrpcDescriptorRegistry")); err != nil {
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
func (m *Executor) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("graphql.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/graphql/v1alpha1.Executor")); err != nil {
		return 0, err
	}

	switch m.Executor.(type) {

	case *Executor_Local_:

		if h, ok := interface{}(m.GetLocal()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Local")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetLocal(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Local")); err != nil {
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
func (m *StitchedSchema_SubschemaConfig) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("graphql.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/graphql/v1alpha1.StitchedSchema_SubschemaConfig")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetName())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetNamespace())); err != nil {
		return 0, err
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetTypeMerge() {
			innerHash.Reset()

			if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
				if _, err = innerHash.Write([]byte("")); err != nil {
					return 0, err
				}
				if _, err = h.Hash(innerHash); err != nil {
					return 0, err
				}
			} else {
				if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
					return 0, err
				} else {
					if _, err = innerHash.Write([]byte("")); err != nil {
						return 0, err
					}
					if err := binary.Write(innerHash, binary.LittleEndian, fieldValue); err != nil {
						return 0, err
					}
				}
			}

			if _, err = innerHash.Write([]byte(k)); err != nil {
				return 0, err
			}

			result = result ^ innerHash.Sum64()
		}
		err = binary.Write(hasher, binary.LittleEndian, result)
		if err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *StitchedSchema_SubschemaConfig_TypeMergeConfig) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("graphql.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/graphql/v1alpha1.StitchedSchema_SubschemaConfig_TypeMergeConfig")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetSelectionSet())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetQueryName())); err != nil {
		return 0, err
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetArgs() {
			innerHash.Reset()

			if _, err = innerHash.Write([]byte(v)); err != nil {
				return 0, err
			}

			if _, err = innerHash.Write([]byte(k)); err != nil {
				return 0, err
			}

			result = result ^ innerHash.Sum64()
		}
		err = binary.Write(hasher, binary.LittleEndian, result)
		if err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *MockResolver_AsyncResponse) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("graphql.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/graphql/v1alpha1.MockResolver_AsyncResponse")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetResponse()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Response")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetResponse(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Response")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetDelay()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Delay")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetDelay(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Delay")); err != nil {
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
func (m *GraphQLApi_GraphqlApiOptions) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("graphql.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/graphql/v1alpha1.GraphQLApi_GraphqlApiOptions")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetLogSensitiveInfo())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *Executor_Local) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("graphql.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/graphql/v1alpha1.Executor_Local")); err != nil {
		return 0, err
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetResolutions() {
			innerHash.Reset()

			if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
				if _, err = innerHash.Write([]byte("")); err != nil {
					return 0, err
				}
				if _, err = h.Hash(innerHash); err != nil {
					return 0, err
				}
			} else {
				if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
					return 0, err
				} else {
					if _, err = innerHash.Write([]byte("")); err != nil {
						return 0, err
					}
					if err := binary.Write(innerHash, binary.LittleEndian, fieldValue); err != nil {
						return 0, err
					}
				}
			}

			if _, err = innerHash.Write([]byte(k)); err != nil {
				return 0, err
			}

			result = result ^ innerHash.Sum64()
		}
		err = binary.Write(hasher, binary.LittleEndian, result)
		if err != nil {
			return 0, err
		}

	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetEnableIntrospection())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetOptions()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Options")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetOptions(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Options")); err != nil {
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
func (m *Executor_Local_LocalExecutorOptions) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("graphql.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/graphql/v1alpha1.Executor_Local_LocalExecutorOptions")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetMaxDepth()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("MaxDepth")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMaxDepth(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("MaxDepth")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}
