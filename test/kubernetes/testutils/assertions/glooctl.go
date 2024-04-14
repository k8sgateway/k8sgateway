package assertions

import (
	"context"
	"time"

	"github.com/solo-io/gloo/test/kube2e"

	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/check"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/printers"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

func (p *Provider) CheckResources() ClusterAssertion {
	p.requiresGlooGatewayContext()

	return func(ctx context.Context) {
		p.testingFramework.Helper()

		Eventually(func(g Gomega) {
			contextWithCancel, cancel := context.WithCancel(ctx)
			defer cancel()
			opts := &options.Options{
				Metadata: core.Metadata{
					Namespace: p.glooGatewayContext.InstallNamespace,
				},
				Top: options.Top{
					Ctx: contextWithCancel,
				},
			}
			err := check.CheckResources(contextWithCancel, printers.P{}, opts)
			g.Expect(err).NotTo(HaveOccurred())
		}).
			WithContext(ctx).
			// These are some basic defaults that we expect to work in most cases
			// We can make these configurable if need be, though most installations
			// Should be able to become healthy within this window
			WithTimeout(time.Second * 90).
			WithPolling(time.Second).
			Should(Succeed())
	}
}

func (p *Provider) InstallationWasSuccessful() ClusterAssertion {
	p.requiresGlooGatewayContext()

	return func(ctx context.Context) {
		p.testingFramework.Helper()

		// Check that everything is OK
		p.CheckResources()(ctx)

		// Ensure gloo reaches valid state and doesn't continually resync
		// we can consider doing the same for leaking go-routines after resyncs
		kube2e.EventuallyReachesConsistentState(p.glooGatewayContext.InstallNamespace)
	}
}

func (p *Provider) UninstallationWasSuccessful() ClusterAssertion {
	p.requiresGlooGatewayContext()

	return func(ctx context.Context) {
		p.testingFramework.Helper()

		p.NamespaceNotExist(p.glooGatewayContext.InstallNamespace)(ctx)
	}
}
