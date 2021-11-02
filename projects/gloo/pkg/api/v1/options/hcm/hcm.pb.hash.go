// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/hcm/hcm.proto

package hcm

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
func (m *HttpConnectionManagerSettings) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("hcm.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/hcm.HttpConnectionManagerSettings")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetSkipXffAppend())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetVia())); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetXffNumTrustedHops())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetUseRemoteAddress()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("UseRemoteAddress")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetUseRemoteAddress(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("UseRemoteAddress")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetGenerateRequestId()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("GenerateRequestId")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetGenerateRequestId(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("GenerateRequestId")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetProxy_100Continue())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetStreamIdleTimeout()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("StreamIdleTimeout")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetStreamIdleTimeout(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("StreamIdleTimeout")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetIdleTimeout()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("IdleTimeout")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetIdleTimeout(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("IdleTimeout")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetMaxRequestHeadersKb()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("MaxRequestHeadersKb")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMaxRequestHeadersKb(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("MaxRequestHeadersKb")); err != nil {
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

	if h, ok := interface{}(m.GetDrainTimeout()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("DrainTimeout")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetDrainTimeout(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("DrainTimeout")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetDelayedCloseTimeout()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("DelayedCloseTimeout")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetDelayedCloseTimeout(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("DelayedCloseTimeout")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte(m.GetServerName())); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetAcceptHttp_10())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetDefaultHostForHttp_10())); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetProperCaseHeaderKeyFormat())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetTracing()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Tracing")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetTracing(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Tracing")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetForwardClientCertDetails())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetSetCurrentClientCertDetails()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("SetCurrentClientCertDetails")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetSetCurrentClientCertDetails(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("SetCurrentClientCertDetails")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetPreserveExternalRequestId())
	if err != nil {
		return 0, err
	}

	for _, v := range m.GetUpgrades() {

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

	if h, ok := interface{}(m.GetMaxConnectionDuration()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("MaxConnectionDuration")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMaxConnectionDuration(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("MaxConnectionDuration")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetMaxStreamDuration()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("MaxStreamDuration")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMaxStreamDuration(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("MaxStreamDuration")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetMaxHeadersCount()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("MaxHeadersCount")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMaxHeadersCount(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("MaxHeadersCount")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetServerHeaderTransformation())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetPathWithEscapedSlashesAction())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetCodecType())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *HttpConnectionManagerSettings_SetCurrentClientCertDetails) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("hcm.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/hcm.HttpConnectionManagerSettings_SetCurrentClientCertDetails")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetSubject()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Subject")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetSubject(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Subject")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetCert())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetChain())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetDns())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetUri())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}
