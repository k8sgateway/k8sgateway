// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/ssl/ssl.proto

package ssl

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
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *SslConfig) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/ssl.SslConfig")); err != nil {
		return 0, err
	}

	for _, v := range m.GetSniDomains() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	for _, v := range m.GetVerifySubjectAltName() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	if h, ok := interface{}(m.GetParameters()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Parameters")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetParameters(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Parameters")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	for _, v := range m.GetAlpnProtocols() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	if h, ok := interface{}(m.GetOneWayTls()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("OneWayTls")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetOneWayTls(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("OneWayTls")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetDisableTlsSessionResumption()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("DisableTlsSessionResumption")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetDisableTlsSessionResumption(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("DisableTlsSessionResumption")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetTransportSocketConnectTimeout()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("TransportSocketConnectTimeout")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetTransportSocketConnectTimeout(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("TransportSocketConnectTimeout")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetOcspStaplePolicy())
	if err != nil {
		return 0, err
	}

	switch m.SslSecrets.(type) {

	case *SslConfig_SecretRef:

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

	case *SslConfig_SslFiles:

		if h, ok := interface{}(m.GetSslFiles()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("SslFiles")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetSslFiles(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("SslFiles")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *SslConfig_Sds:

		if h, ok := interface{}(m.GetSds()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Sds")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetSds(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Sds")); err != nil {
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
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *SSLFiles) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/ssl.SSLFiles")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetTlsCert())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetTlsKey())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetRootCa())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetOcspStaple())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *UpstreamSslConfig) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/ssl.UpstreamSslConfig")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetSni())); err != nil {
		return 0, err
	}

	for _, v := range m.GetVerifySubjectAltName() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	if h, ok := interface{}(m.GetParameters()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Parameters")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetParameters(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Parameters")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	for _, v := range m.GetAlpnProtocols() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	if h, ok := interface{}(m.GetAllowRenegotiation()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("AllowRenegotiation")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetAllowRenegotiation(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("AllowRenegotiation")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetOneWayTls()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("OneWayTls")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetOneWayTls(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("OneWayTls")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	switch m.SslSecrets.(type) {

	case *UpstreamSslConfig_SecretRef:

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

	case *UpstreamSslConfig_SslFiles:

		if h, ok := interface{}(m.GetSslFiles()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("SslFiles")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetSslFiles(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("SslFiles")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *UpstreamSslConfig_Sds:

		if h, ok := interface{}(m.GetSds()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Sds")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetSds(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Sds")); err != nil {
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
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *SDSConfig) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/ssl.SDSConfig")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetTargetUri())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetCertificatesSecretName())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetValidationContextName())); err != nil {
		return 0, err
	}

	switch m.SdsBuilder.(type) {

	case *SDSConfig_CallCredentials:

		if h, ok := interface{}(m.GetCallCredentials()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("CallCredentials")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetCallCredentials(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("CallCredentials")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *SDSConfig_ClusterName:

		if _, err = hasher.Write([]byte(m.GetClusterName())); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *CallCredentials) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/ssl.CallCredentials")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetFileCredentialSource()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("FileCredentialSource")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetFileCredentialSource(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("FileCredentialSource")); err != nil {
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
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *SslParameters) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/ssl.SslParameters")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetMinimumProtocolVersion())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetMaximumProtocolVersion())
	if err != nil {
		return 0, err
	}

	for _, v := range m.GetCipherSuites() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	for _, v := range m.GetEcdhCurves() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *CallCredentials_FileCredentialSource) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/ssl.CallCredentials_FileCredentialSource")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetTokenFileName())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetHeader())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}
