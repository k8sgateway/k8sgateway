// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/options/waf/waf.proto

package waf

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
func (m *Settings) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("waf.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/waf.Settings")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetDisabled())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetCustomInterventionMessage())); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetCoreRuleSet()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("CoreRuleSet")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetCoreRuleSet(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("CoreRuleSet")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	for _, v := range m.GetRuleSets() {

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

	for _, v := range m.GetCustomConfigMapRuleSets() {

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

	if h, ok := interface{}(m.GetAuditLogging()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("AuditLogging")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetAuditLogging(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("AuditLogging")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetRequestHeadersOnly())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetResponseHeadersOnly())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *CoreRuleSet) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("waf.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/waf.CoreRuleSet")); err != nil {
		return 0, err
	}

	switch m.CustomSettingsType.(type) {

	case *CoreRuleSet_CustomSettingsString:

		if _, err = hasher.Write([]byte(m.GetCustomSettingsString())); err != nil {
			return 0, err
		}

	case *CoreRuleSet_CustomSettingsFile:

		if _, err = hasher.Write([]byte(m.GetCustomSettingsFile())); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}
