// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/v1/connection.proto

package v1

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
func (m *ConnectionConfig) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1.ConnectionConfig")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("MaxRequestsPerConnection")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetMaxRequestsPerConnection())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetConnectTimeout()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("ConnectTimeout")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetConnectTimeout(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("ConnectTimeout")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetTcpKeepalive()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("TcpKeepalive")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetTcpKeepalive(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("TcpKeepalive")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetPerConnectionBufferLimitBytes()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("PerConnectionBufferLimitBytes")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetPerConnectionBufferLimitBytes(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("PerConnectionBufferLimitBytes")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetCommonHttpProtocolOptions()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("CommonHttpProtocolOptions")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetCommonHttpProtocolOptions(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("CommonHttpProtocolOptions")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetHttp1ProtocolOptions()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Http1ProtocolOptions")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetHttp1ProtocolOptions(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Http1ProtocolOptions")); err != nil {
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
func (m *ConnectionConfig_TcpKeepAlive) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1.ConnectionConfig_TcpKeepAlive")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("KeepaliveProbes")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetKeepaliveProbes())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetKeepaliveTime()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("KeepaliveTime")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetKeepaliveTime(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("KeepaliveTime")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetKeepaliveInterval()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("KeepaliveInterval")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetKeepaliveInterval(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("KeepaliveInterval")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}
