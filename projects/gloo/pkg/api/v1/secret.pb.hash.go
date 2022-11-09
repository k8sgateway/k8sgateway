// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/secret.proto

package v1

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
func (m *Secret) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1.Secret")); err != nil {
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

	switch m.Kind.(type) {

	case *Secret_Aws:

		if h, ok := interface{}(m.GetAws()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Aws")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetAws(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Aws")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *Secret_Azure:

		if h, ok := interface{}(m.GetAzure()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Azure")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetAzure(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Azure")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *Secret_Tls:

		if h, ok := interface{}(m.GetTls()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Tls")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetTls(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Tls")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *Secret_Oauth:

		if h, ok := interface{}(m.GetOauth()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Oauth")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetOauth(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Oauth")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *Secret_ApiKey:

		if h, ok := interface{}(m.GetApiKey()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("ApiKey")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetApiKey(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("ApiKey")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *Secret_Header:

		if h, ok := interface{}(m.GetHeader()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Header")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetHeader(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Header")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *Secret_Credentials:

		if h, ok := interface{}(m.GetCredentials()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Credentials")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetCredentials(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Credentials")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *Secret_Extensions:

		if h, ok := interface{}(m.GetExtensions()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Extensions")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetExtensions(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Extensions")); err != nil {
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
func (m *AwsSecret) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1.AwsSecret")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetAccessKey())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetSecretKey())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetSessionToken())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *AzureSecret) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1.AzureSecret")); err != nil {
		return 0, err
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetApiKeys() {
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
func (m *TlsSecret) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1.TlsSecret")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetCertChain())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetPrivateKey())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetRootCa())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *HeaderSecret) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1.HeaderSecret")); err != nil {
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

	return hasher.Sum64(), nil
}

// Hash function
func (m *AccountCredentialsSecret) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1.AccountCredentialsSecret")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetUsername())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetPassword())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}
