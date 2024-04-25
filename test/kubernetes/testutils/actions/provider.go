package actions

import (
	"testing"

	"github.com/solo-io/gloo/pkg/utils/kubeutils/kubectl"
	"github.com/stretchr/testify/require"

	"github.com/solo-io/gloo/test/kubernetes/testutils/cluster"
	"github.com/solo-io/gloo/test/kubernetes/testutils/gloogateway"
)

// Provider is the entity that creates actions.
// These actions are executed against a running installation of Gloo Gateway, within a Kubernetes Cluster.
// This provider is just a wrapper around sub-providers, so it exposes methods to access those providers
type Provider struct {
	require *require.Assertions

	kubeCli *kubectl.Cli

	glooctl Glooctl
}

// NewActionsProvider returns an Provider
func NewActionsProvider(t *testing.T) *Provider {
	return &Provider{
		require: require.New(t),
		kubeCli: nil,
		glooctl: NewGlooctl(t),
	}
}

// WithClusterContext sets the provider, and all of it's sub-providers, to point to the provided cluster
func (p *Provider) WithClusterContext(clusterContext *cluster.Context) *Provider {
	p.kubeCli = clusterContext.Cli
	p.glooctl.WithClusterContext(clusterContext)
	return p
}

// WithGlooGatewayContext sets the provider, and all of it's sub-providers, to point to the provided installation
func (p *Provider) WithGlooGatewayContext(ggCtx *gloogateway.Context) *Provider {
	p.glooctl.WithGlooGatewayContext(ggCtx)
	return p
}

// WithGlooctl sets the glooctl on this Provider
func (p *Provider) WithGlooctl(glooctl Glooctl) *Provider {
	p.glooctl = glooctl
	return p
}

func (p *Provider) Kubectl() *kubectl.Cli {
	return p.kubeCli
}

func (p *Provider) Glooctl() Glooctl {
	return p.glooctl
}
