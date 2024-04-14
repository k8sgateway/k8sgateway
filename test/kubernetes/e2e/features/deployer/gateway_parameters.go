package deployer

import (
	"context"
	"time"

	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/pkg/utils/envoyutils/admincli"
	"github.com/solo-io/gloo/projects/gateway2/pkg/api/gateway.gloo.solo.io/v1alpha1"
	"github.com/solo-io/gloo/test/kubernetes/e2e"
	"github.com/solo-io/gloo/test/kubernetes/testutils/operations"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	gwParametersManifestFile = e2e.FeatureManifestFile("gateway-parameters.yaml")

	gwParams = &v1alpha1.GatewayParameters{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "gw-params",
			Namespace: "default",
		},
	}
)

var ConfigureProxiesFromGatewayParameters = e2e.Test{
	Name:        "Deployer.ConfigureProxiesFromGatewayParameters",
	Description: "the deployer will provision a deployment and service for a defined gateway, and configure it based on the GatewayParameters CR",

	Test: func(ctx context.Context, installation *e2e.TestInstallation) {
		provisionResourcesOp := operations.ReversibleOperation{
			Do: installation.OperationsProvider.KubeCtl().NewApplyManifestOperation(
				manifestFile,
				installation.AssertionsProvider.ObjectsExist(proxyService, proxyDeployment),
			),
			// We rely on the --ignore-not-found flag in the deletion command, because we have 2 manifests
			// that manage the same resource (manifestFile, gwParametersManifestFile).
			// So when we perform Undo of configureGatewayParametersOp, it will delete the Gateway CR,
			// and then this operation  will also attempt to delete the same resource.
			// Ideally, we do not include the same resource in multiple manifests that are used by a test
			// But this is an example of ways to solve that problem if it occurs
			Undo: installation.OperationsProvider.KubeCtl().NewDeleteManifestIgnoreNotFoundOperation(
				manifestFile,
				installation.AssertionsProvider.ObjectsNotExist(proxyService, proxyDeployment),
			),
		}

		configureGatewayParametersOp := operations.ReversibleOperation{
			Do: installation.OperationsProvider.KubeCtl().NewApplyManifestOperation(
				gwParametersManifestFile,

				// We applied a manifest containing the GatewayParameters CR
				installation.AssertionsProvider.ObjectsExist(gwParams),

				// We configure the GatewayParameters CR to provision workloads with a specific image that should exist
				installation.AssertionsProvider.RunningReplicas(proxyDeployment.ObjectMeta, 1),

				// We assert that we can port-forward requests to the proxy deployment, and then execute requests against the server
				installation.AssertionsProvider.EnvoyAdminApiAssertion(
					proxyDeployment.ObjectMeta,
					installation.OperationsProvider.KubeCtl().Client(),
					func(ctx context.Context, adminClient *admincli.Client) {
						Eventually(func(g Gomega) {
							serverInfo, err := adminClient.GetServerInfo(ctx)
							g.Expect(err).NotTo(HaveOccurred())
							g.Expect(serverInfo.GetCommandLineOptions().GetLogLevel()).To(Equal("debug"), "defined on the GatewayParameters CR")
						}).
							WithContext(ctx).
							WithTimeout(time.Second * 10).
							WithPolling(time.Millisecond * 200).
							Should(Succeed())
					},
				),
			),
			Undo: installation.OperationsProvider.KubeCtl().NewDeleteManifestOperation(
				gwParametersManifestFile,
				installation.AssertionsProvider.ObjectsNotExist(gwParams),
			),
		}

		err := installation.Operator.ExecuteReversibleOperations(ctx, provisionResourcesOp, configureGatewayParametersOp)
		Expect(err).NotTo(HaveOccurred())
	},
}