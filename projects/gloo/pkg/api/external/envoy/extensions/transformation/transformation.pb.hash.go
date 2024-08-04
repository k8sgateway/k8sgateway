// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/extensions/transformation/transformation.proto

package transformation

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
func (m *FilterTransformations) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.api.v2.filter.http.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/transformation.FilterTransformations")); err != nil {
		return 0, err
	}

	for _, v := range m.GetTransformations() {

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

	err = binary.Write(hasher, binary.LittleEndian, m.GetStage())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetLogRequestResponseInfo())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *TransformationRule) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.api.v2.filter.http.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/transformation.TransformationRule")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetMatch()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Match")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMatch(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Match")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetRouteTransformations()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("RouteTransformations")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRouteTransformations(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("RouteTransformations")); err != nil {
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
func (m *RouteTransformations) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.api.v2.filter.http.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/transformation.RouteTransformations")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetRequestTransformation()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("RequestTransformation")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRequestTransformation(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("RequestTransformation")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetResponseTransformation()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("ResponseTransformation")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetResponseTransformation(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("ResponseTransformation")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetClearRouteCache())
	if err != nil {
		return 0, err
	}

	for _, v := range m.GetTransformations() {

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
func (m *ResponseMatcher) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.api.v2.filter.http.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/transformation.ResponseMatcher")); err != nil {
		return 0, err
	}

	for _, v := range m.GetHeaders() {

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

	if h, ok := interface{}(m.GetResponseCodeDetails()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("ResponseCodeDetails")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetResponseCodeDetails(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("ResponseCodeDetails")); err != nil {
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
func (m *ResponseTransformationRule) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.api.v2.filter.http.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/transformation.ResponseTransformationRule")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetMatch()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Match")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMatch(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Match")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetResponseTransformation()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("ResponseTransformation")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetResponseTransformation(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("ResponseTransformation")); err != nil {
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
func (m *Transformation) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.api.v2.filter.http.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/transformation.Transformation")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetLogRequestResponseInfo()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("LogRequestResponseInfo")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetLogRequestResponseInfo(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("LogRequestResponseInfo")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	switch m.TransformationType.(type) {

	case *Transformation_TransformationTemplate:

		if h, ok := interface{}(m.GetTransformationTemplate()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("TransformationTemplate")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetTransformationTemplate(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("TransformationTemplate")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *Transformation_HeaderBodyTransform:

		if h, ok := interface{}(m.GetHeaderBodyTransform()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("HeaderBodyTransform")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetHeaderBodyTransform(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("HeaderBodyTransform")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *Transformation_TransformerConfig:

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

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *Extraction) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.api.v2.filter.http.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/transformation.Extraction")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetRegex())); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetSubgroup())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetReplacementText()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("ReplacementText")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetReplacementText(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("ReplacementText")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetMode())
	if err != nil {
		return 0, err
	}

	switch m.Source.(type) {

	case *Extraction_Header:

		if _, err = hasher.Write([]byte(m.GetHeader())); err != nil {
			return 0, err
		}

	case *Extraction_Body:

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

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *TransformationTemplate) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.api.v2.filter.http.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/transformation.TransformationTemplate")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetAdvancedTemplates())
	if err != nil {
		return 0, err
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetExtractors() {
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

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetHeaders() {
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

	for _, v := range m.GetHeadersToAppend() {

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

	for _, v := range m.GetHeadersToRemove() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetParseBodyBehavior())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetIgnoreErrorOnParse())
	if err != nil {
		return 0, err
	}

	for _, v := range m.GetDynamicMetadataValues() {

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

	err = binary.Write(hasher, binary.LittleEndian, m.GetEscapeCharacters())
	if err != nil {
		return 0, err
	}

	switch m.BodyTransformation.(type) {

	case *TransformationTemplate_Body:

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

	case *TransformationTemplate_Passthrough:

		if h, ok := interface{}(m.GetPassthrough()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Passthrough")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetPassthrough(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Passthrough")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *TransformationTemplate_MergeExtractorsToBody:

		if h, ok := interface{}(m.GetMergeExtractorsToBody()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("MergeExtractorsToBody")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetMergeExtractorsToBody(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("MergeExtractorsToBody")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *TransformationTemplate_MergeJsonKeys:

		if h, ok := interface{}(m.GetMergeJsonKeys()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("MergeJsonKeys")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetMergeJsonKeys(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("MergeJsonKeys")); err != nil {
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
func (m *InjaTemplate) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.api.v2.filter.http.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/transformation.InjaTemplate")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetText())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *Passthrough) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.api.v2.filter.http.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/transformation.Passthrough")); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *MergeExtractorsToBody) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.api.v2.filter.http.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/transformation.MergeExtractorsToBody")); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *MergeJsonKeys) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.api.v2.filter.http.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/transformation.MergeJsonKeys")); err != nil {
		return 0, err
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetJsonKeys() {
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
func (m *HeaderBodyTransform) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.api.v2.filter.http.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/transformation.HeaderBodyTransform")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetAddRequestMetadata())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *TransformationRule_Transformations) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.api.v2.filter.http.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/transformation.TransformationRule_Transformations")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetRequestTransformation()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("RequestTransformation")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRequestTransformation(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("RequestTransformation")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetClearRouteCache())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetResponseTransformation()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("ResponseTransformation")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetResponseTransformation(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("ResponseTransformation")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetOnStreamCompletionTransformation()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("OnStreamCompletionTransformation")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetOnStreamCompletionTransformation(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("OnStreamCompletionTransformation")); err != nil {
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
func (m *RouteTransformations_RouteTransformation) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.api.v2.filter.http.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/transformation.RouteTransformations_RouteTransformation")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetStage())
	if err != nil {
		return 0, err
	}

	switch m.Match.(type) {

	case *RouteTransformations_RouteTransformation_RequestMatch_:

		if h, ok := interface{}(m.GetRequestMatch()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("RequestMatch")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetRequestMatch(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("RequestMatch")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *RouteTransformations_RouteTransformation_ResponseMatch_:

		if h, ok := interface{}(m.GetResponseMatch()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("ResponseMatch")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetResponseMatch(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("ResponseMatch")); err != nil {
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
func (m *RouteTransformations_RouteTransformation_RequestMatch) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.api.v2.filter.http.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/transformation.RouteTransformations_RouteTransformation_RequestMatch")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetMatch()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Match")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMatch(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Match")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetRequestTransformation()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("RequestTransformation")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRequestTransformation(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("RequestTransformation")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetResponseTransformation()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("ResponseTransformation")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetResponseTransformation(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("ResponseTransformation")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetClearRouteCache())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *RouteTransformations_RouteTransformation_ResponseMatch) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.api.v2.filter.http.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/transformation.RouteTransformations_RouteTransformation_ResponseMatch")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetMatch()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Match")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMatch(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Match")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetResponseTransformation()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("ResponseTransformation")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetResponseTransformation(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("ResponseTransformation")); err != nil {
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
func (m *TransformationTemplate_HeaderToAppend) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.api.v2.filter.http.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/transformation.TransformationTemplate_HeaderToAppend")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetKey())); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetValue()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Value")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetValue(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Value")); err != nil {
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
func (m *TransformationTemplate_DynamicMetadataValue) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.api.v2.filter.http.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/transformation.TransformationTemplate_DynamicMetadataValue")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetMetadataNamespace())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetKey())); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetValue()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Value")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetValue(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Value")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetJsonToProto())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *MergeJsonKeys_OverridableTemplate) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.api.v2.filter.http.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/transformation.MergeJsonKeys_OverridableTemplate")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetTmpl()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Tmpl")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetTmpl(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Tmpl")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetOverrideEmpty())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}
