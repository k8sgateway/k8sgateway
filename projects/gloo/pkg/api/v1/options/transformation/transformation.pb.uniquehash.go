// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/transformation/transformation.proto

package transformation

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
func (m *ResponseMatch) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("transformation.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/transformation.ResponseMatch")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Matchers")); err != nil {
		return 0, err
	}
	for i, v := range m.GetMatchers() {
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

	if _, err = hasher.Write([]byte("ResponseCodeDetails")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetResponseCodeDetails())); err != nil {
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

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *RequestMatch) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("transformation.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/transformation.RequestMatch")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetMatcher()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Matcher")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMatcher(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Matcher")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte("ClearRouteCache")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetClearRouteCache())
	if err != nil {
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

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *Transformations) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("transformation.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/transformation.Transformations")); err != nil {
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

	if _, err = hasher.Write([]byte("ClearRouteCache")); err != nil {
		return 0, err
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

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *RequestResponseTransformations) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("transformation.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/transformation.RequestResponseTransformations")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("RequestTransforms")); err != nil {
		return 0, err
	}
	for i, v := range m.GetRequestTransforms() {
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

	if _, err = hasher.Write([]byte("ResponseTransforms")); err != nil {
		return 0, err
	}
	for i, v := range m.GetResponseTransforms() {
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
func (m *TransformationStages) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("transformation.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/transformation.TransformationStages")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetEarly()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Early")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetEarly(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Early")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetRegular()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Regular")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRegular(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Regular")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetPostRouting()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("PostRouting")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetPostRouting(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("PostRouting")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte("InheritTransformation")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetInheritTransformation())
	if err != nil {
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

	if h, ok := interface{}(m.GetEscapeCharacters()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("EscapeCharacters")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetEscapeCharacters(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("EscapeCharacters")); err != nil {
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
func (m *Transformation) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("transformation.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/transformation.Transformation")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("LogRequestResponseInfo")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetLogRequestResponseInfo())
	if err != nil {
		return 0, err
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

	case *Transformation_XsltTransformation:

		if h, ok := interface{}(m.GetXsltTransformation()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("XsltTransformation")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetXsltTransformation(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("XsltTransformation")); err != nil {
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
func (m *Extraction) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("transformation.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/transformation.Extraction")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Regex")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetRegex())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Subgroup")); err != nil {
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

	if _, err = hasher.Write([]byte("Mode")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetMode())
	if err != nil {
		return 0, err
	}

	switch m.Source.(type) {

	case *Extraction_Header:

		if _, err = hasher.Write([]byte("Header")); err != nil {
			return 0, err
		}
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

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *TransformationTemplate) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("transformation.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/transformation.TransformationTemplate")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("AdvancedTemplates")); err != nil {
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
				if _, err = innerHash.Write([]byte("v")); err != nil {
					return 0, err
				}
				if _, err = h.Hash(innerHash); err != nil {
					return 0, err
				}
			} else {
				if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
					return 0, err
				} else {
					if _, err = innerHash.Write([]byte("v")); err != nil {
						return 0, err
					}
					if err := binary.Write(innerHash, binary.LittleEndian, fieldValue); err != nil {
						return 0, err
					}
				}
			}

			if _, err = innerHash.Write([]byte("k")); err != nil {
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
		for k, v := range m.GetHeaders() {
			innerHash.Reset()

			if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
				if _, err = innerHash.Write([]byte("v")); err != nil {
					return 0, err
				}
				if _, err = h.Hash(innerHash); err != nil {
					return 0, err
				}
			} else {
				if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
					return 0, err
				} else {
					if _, err = innerHash.Write([]byte("v")); err != nil {
						return 0, err
					}
					if err := binary.Write(innerHash, binary.LittleEndian, fieldValue); err != nil {
						return 0, err
					}
				}
			}

			if _, err = innerHash.Write([]byte("k")); err != nil {
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

	if _, err = hasher.Write([]byte("HeadersToAppend")); err != nil {
		return 0, err
	}
	for i, v := range m.GetHeadersToAppend() {
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

	if _, err = hasher.Write([]byte("HeadersToRemove")); err != nil {
		return 0, err
	}
	for i, v := range m.GetHeadersToRemove() {
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

	if _, err = hasher.Write([]byte("ParseBodyBehavior")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetParseBodyBehavior())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("IgnoreErrorOnParse")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetIgnoreErrorOnParse())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("DynamicMetadataValues")); err != nil {
		return 0, err
	}
	for i, v := range m.GetDynamicMetadataValues() {
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

	if h, ok := interface{}(m.GetEscapeCharacters()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("EscapeCharacters")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetEscapeCharacters(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("EscapeCharacters")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
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

	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *InjaTemplate) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("transformation.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/transformation.InjaTemplate")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Text")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetText())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *Passthrough) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("transformation.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/transformation.Passthrough")); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *MergeExtractorsToBody) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("transformation.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/transformation.MergeExtractorsToBody")); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *HeaderBodyTransform) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("transformation.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/transformation.HeaderBodyTransform")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("AddRequestMetadata")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetAddRequestMetadata())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *TransformationTemplate_HeaderToAppend) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("transformation.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/transformation.TransformationTemplate_HeaderToAppend")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Key")); err != nil {
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

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *TransformationTemplate_DynamicMetadataValue) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("transformation.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/transformation.TransformationTemplate_DynamicMetadataValue")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("MetadataNamespace")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetMetadataNamespace())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Key")); err != nil {
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

	if _, err = hasher.Write([]byte("JsonToProto")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetJsonToProto())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}