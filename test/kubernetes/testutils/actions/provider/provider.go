package provider

import (
	"testing"

	"github.com/solo-io/gloo/test/kubernetes/testutils/actions/glooctl"
	"github.com/solo-io/gloo/test/kubernetes/testutils/actions/kubectl"

	"github.com/solo-io/gloo/test/kubernetes/testutils/cluster"
	"github.com/solo-io/gloo/test/kubernetes/testutils/gloogateway"
)

// ActionsProvider is the entity that creates actions.
// These actions are executed against a running installation of Gloo Gateway, within a Kubernetes Cluster.
// This provider is just a wrapper around sub-providers, so it exposes methods to access those providers
type ActionsProvider struct {
	kubeCtlProvider *kubectl.Provider
	glooCtlProvider glooctl.Provider
}

// NewActionsProvider returns an ActionsProvider that will fail because it is not configured with a Kubernetes Cluster
func NewActionsProvider(testingFramework testing.TB) *ActionsProvider {
	return &ActionsProvider{
		kubeCtlProvider: kubectl.NewProvider(),
		glooCtlProvider: glooctl.NewProvider(testingFramework),
	}
}

// WithClusterContext sets the provider, and all of it's sub-providers, to point to the provided cluster
func (p *ActionsProvider) WithClusterContext(clusterContext *cluster.Context) *ActionsProvider {
	p.kubeCtlProvider.WithClusterCli(clusterContext.Cli)
	p.glooCtlProvider.WithClusterContext(clusterContext)
	return p
}

// WithGlooGatewayContext sets the provider, and all of it's sub-providers, to point to the provided installation
func (p *ActionsProvider) WithGlooGatewayContext(ggCtx *gloogateway.Context) *ActionsProvider {
	p.glooCtlProvider.WithGlooGatewayContext(ggCtx)
	return p
}

// WithGlooctlProvider sets the glooctl provider on this ActionsProvider
func (p *ActionsProvider) WithGlooctlProvider(provider glooctl.Provider) *ActionsProvider {
	p.glooCtlProvider = provider
	return p
}

func (p *ActionsProvider) KubeCtl() *kubectl.Provider {
	return p.kubeCtlProvider
}

func (p *ActionsProvider) GlooCtl() glooctl.Provider {
	return p.glooCtlProvider
}
