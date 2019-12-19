// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/healthcheck/healthcheck.proto

package healthcheck

import (
	"errors"
	"fmt"
	"hash"
	"hash/fnv"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
)

// Hash function
func (m *HealthCheck) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error

	if _, err = hasher.Write([]byte(m.GetPath())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}
