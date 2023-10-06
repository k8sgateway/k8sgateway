package helpers

import v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"

// SecretBuilder contains options for building Endpoints to be included in scaled Snapshots
// there are no options currently configurable for the endpointBuilder
type SecretBuilder struct{}

func NewSecretBuilder() *SecretBuilder {
	return &SecretBuilder{}
}

func (b *SecretBuilder) Build(i int) *v1.Secret {
	return Secret(i)
}
