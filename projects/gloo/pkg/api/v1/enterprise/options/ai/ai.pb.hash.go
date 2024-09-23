// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/options/ai/ai.proto

package ai

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
func (m *SingleAuthToken) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.SingleAuthToken")); err != nil {
		return 0, err
	}

	switch m.AuthTokenSource.(type) {

	case *SingleAuthToken_Inline:

		if _, err = hasher.Write([]byte(m.GetInline())); err != nil {
			return 0, err
		}

	case *SingleAuthToken_SecretRef:

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

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *UpstreamSpec) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.UpstreamSpec")); err != nil {
		return 0, err
	}

	switch m.Llm.(type) {

	case *UpstreamSpec_Openai:

		if h, ok := interface{}(m.GetOpenai()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Openai")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetOpenai(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Openai")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *UpstreamSpec_Mistral_:

		if h, ok := interface{}(m.GetMistral()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Mistral")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetMistral(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Mistral")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *UpstreamSpec_Anthropic_:

		if h, ok := interface{}(m.GetAnthropic()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Anthropic")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetAnthropic(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Anthropic")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *UpstreamSpec_AzureOpenai:

		if h, ok := interface{}(m.GetAzureOpenai()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("AzureOpenai")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetAzureOpenai(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("AzureOpenai")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *UpstreamSpec_Multi_:

		if h, ok := interface{}(m.GetMulti()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Multi")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetMulti(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Multi")); err != nil {
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
func (m *RouteSettings) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.RouteSettings")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetPromptEnrichment()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("PromptEnrichment")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetPromptEnrichment(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("PromptEnrichment")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetPromptGuard()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("PromptGuard")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetPromptGuard(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("PromptGuard")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetRag()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Rag")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRag(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Rag")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetSemanticCache()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("SemanticCache")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetSemanticCache(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("SemanticCache")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	for _, v := range m.GetDefaults() {

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

	err = binary.Write(hasher, binary.LittleEndian, m.GetRouteType())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *FieldDefault) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.FieldDefault")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetField())); err != nil {
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

	err = binary.Write(hasher, binary.LittleEndian, m.GetOverride())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *Postgres) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.Postgres")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetConnectionString())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetCollectionName())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *Embedding) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.Embedding")); err != nil {
		return 0, err
	}

	switch m.Embedding.(type) {

	case *Embedding_Openai:

		if h, ok := interface{}(m.GetOpenai()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Openai")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetOpenai(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Openai")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *Embedding_AzureOpenai:

		if h, ok := interface{}(m.GetAzureOpenai()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("AzureOpenai")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetAzureOpenai(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("AzureOpenai")); err != nil {
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
func (m *SemanticCache) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.SemanticCache")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetDatastore()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Datastore")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetDatastore(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Datastore")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetEmbedding()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Embedding")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetEmbedding(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Embedding")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetTtl())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetMode())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *RAG) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.RAG")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetDatastore()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Datastore")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetDatastore(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Datastore")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetEmbedding()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Embedding")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetEmbedding(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Embedding")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte(m.GetPromptTemplate())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *RateLimiting) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.RateLimiting")); err != nil {
		return 0, err
	}

	for _, v := range m.GetRateLimitConfigs() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *AIPromptEnrichment) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.AIPromptEnrichment")); err != nil {
		return 0, err
	}

	for _, v := range m.GetPrepend() {

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

	for _, v := range m.GetAppend() {

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
func (m *AIPromptGaurd) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.AIPromptGaurd")); err != nil {
		return 0, err
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

	return hasher.Sum64(), nil
}

// Hash function
func (m *UpstreamSpec_CustomHost) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.UpstreamSpec_CustomHost")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetHost())); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetPort())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *UpstreamSpec_OpenAI) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.UpstreamSpec_OpenAI")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetAuthToken()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("AuthToken")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetAuthToken(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("AuthToken")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetCustomHost()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("CustomHost")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetCustomHost(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("CustomHost")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte(m.GetModel())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *UpstreamSpec_AzureOpenAI) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.UpstreamSpec_AzureOpenAI")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetEndpoint())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetDeploymentName())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetApiVersion())); err != nil {
		return 0, err
	}

	switch m.AuthTokenSource.(type) {

	case *UpstreamSpec_AzureOpenAI_AuthToken:

		if h, ok := interface{}(m.GetAuthToken()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("AuthToken")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetAuthToken(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("AuthToken")); err != nil {
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
func (m *UpstreamSpec_Mistral) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.UpstreamSpec_Mistral")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetAuthToken()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("AuthToken")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetAuthToken(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("AuthToken")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetCustomHost()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("CustomHost")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetCustomHost(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("CustomHost")); err != nil {
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
func (m *UpstreamSpec_Anthropic) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.UpstreamSpec_Anthropic")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetAuthToken()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("AuthToken")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetAuthToken(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("AuthToken")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetCustomHost()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("CustomHost")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetCustomHost(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("CustomHost")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte(m.GetVersion())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *UpstreamSpec_Multi) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.UpstreamSpec_Multi")); err != nil {
		return 0, err
	}

	for _, v := range m.GetPools() {

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
func (m *UpstreamSpec_Multi_Backend) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.UpstreamSpec_Multi_Backend")); err != nil {
		return 0, err
	}

	switch m.Llm.(type) {

	case *UpstreamSpec_Multi_Backend_Openai:

		if h, ok := interface{}(m.GetOpenai()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Openai")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetOpenai(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Openai")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *UpstreamSpec_Multi_Backend_Mistral:

		if h, ok := interface{}(m.GetMistral()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Mistral")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetMistral(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Mistral")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *UpstreamSpec_Multi_Backend_Anthropic:

		if h, ok := interface{}(m.GetAnthropic()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Anthropic")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetAnthropic(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Anthropic")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *UpstreamSpec_Multi_Backend_AzureOpenai:

		if h, ok := interface{}(m.GetAzureOpenai()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("AzureOpenai")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetAzureOpenai(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("AzureOpenai")); err != nil {
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
func (m *UpstreamSpec_Multi_Priority) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.UpstreamSpec_Multi_Priority")); err != nil {
		return 0, err
	}

	for _, v := range m.GetPool() {

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

	err = binary.Write(hasher, binary.LittleEndian, m.GetPriority())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *Embedding_OpenAI) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.Embedding_OpenAI")); err != nil {
		return 0, err
	}

	switch m.AuthTokenSource.(type) {

	case *Embedding_OpenAI_AuthToken:

		if h, ok := interface{}(m.GetAuthToken()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("AuthToken")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetAuthToken(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("AuthToken")); err != nil {
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
func (m *Embedding_AzureOpenAI) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.Embedding_AzureOpenAI")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetApiVersion())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetEndpoint())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetDeploymentName())); err != nil {
		return 0, err
	}

	switch m.AuthTokenSource.(type) {

	case *Embedding_AzureOpenAI_AuthToken:

		if h, ok := interface{}(m.GetAuthToken()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("AuthToken")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetAuthToken(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("AuthToken")); err != nil {
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
func (m *SemanticCache_Redis) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.SemanticCache_Redis")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetConnectionString())); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetScoreThreshold())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *SemanticCache_Weaviate) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.SemanticCache_Weaviate")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetHost())); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetHttpPort())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetGrpcPort())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetInsecure())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *SemanticCache_DataStore) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.SemanticCache_DataStore")); err != nil {
		return 0, err
	}

	switch m.Datastore.(type) {

	case *SemanticCache_DataStore_Redis:

		if h, ok := interface{}(m.GetRedis()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Redis")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetRedis(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Redis")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *SemanticCache_DataStore_Weaviate:

		if h, ok := interface{}(m.GetWeaviate()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Weaviate")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetWeaviate(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Weaviate")); err != nil {
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
func (m *RAG_DataStore) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.RAG_DataStore")); err != nil {
		return 0, err
	}

	switch m.Datastore.(type) {

	case *RAG_DataStore_Postgres:

		if h, ok := interface{}(m.GetPostgres()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Postgres")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetPostgres(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Postgres")); err != nil {
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
func (m *AIPromptEnrichment_Message) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.AIPromptEnrichment_Message")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetRole())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetContent())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *AIPromptGaurd_Request) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.AIPromptGaurd_Request")); err != nil {
		return 0, err
	}

	for _, v := range m.GetMatches() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	if _, err = hasher.Write([]byte(m.GetCustomResponseMessage())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *AIPromptGaurd_Response) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.AIPromptGaurd_Response")); err != nil {
		return 0, err
	}

	for _, v := range m.GetMatches() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	for _, v := range m.GetBuiltins() {

		err = binary.Write(hasher, binary.LittleEndian, v)
		if err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}
