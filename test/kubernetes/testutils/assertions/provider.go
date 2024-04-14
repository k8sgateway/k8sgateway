package assertions

import (
	"testing"

	"github.com/solo-io/gloo/test/kubernetes/testutils/gloogateway"

	"github.com/solo-io/gloo/test/kubernetes/testutils/cluster"
)

// Provider is the entity that creates operations.DiscreteAssertion
// These assertions occur against a running instance of Gloo Gateway, within a Kubernetes Cluster.
// So this provider maintains state about the install/cluster it is using, and then provides
// operations.DiscreteAssertion to match
type Provider struct {
	testingFramework testing.TB

	clusterContext     *cluster.Context
	glooGatewayContext *gloogateway.Context
}

// NewProvider returns a Provider that will provide Assertions that can be executed against an
// installation of Gloo Gateway
func NewProvider(testingFramework testing.TB) *Provider {
	return &Provider{
		testingFramework:   testingFramework,
		clusterContext:     nil,
		glooGatewayContext: nil,
	}
}

// WithClusterContext sets the provider to point to the provided cluster
func (p *Provider) WithClusterContext(clusterContext *cluster.Context) *Provider {
	p.clusterContext = clusterContext
	return p
}

// WithGlooGatewayContext sets the providers to point to a particualr installation of Gloo Gateway
func (p *Provider) WithGlooGatewayContext(ggCtx *gloogateway.Context) *Provider {
	p.glooGatewayContext = ggCtx
	return p
}

// requiresGlooGatewayContext is invoked by methods on the Provider that can only be invoked
// if the provider has been configured to point to a Gloo Gateway installation
// There are certain Assertions that can be invoked that do not require that Gloo Gateway be installed for them to be invoked
func (p *Provider) requiresGlooGatewayContext() {
	if p.glooGatewayContext == nil {
		p.testingFramework.Fatal("Provider attempted to create an Assertion that requires a Gloo Gateway installation, but none was configured")
	}
}