package tests

import (
	"github.com/solo-io/gloo/test/kubernetes/e2e/features/headless_svc"
	"github.com/solo-io/gloo/test/kubernetes/e2e/features/istio"
)

func AutomtlsIstioEdgeApiTests() TestRunner {
	automtlsIstioEdgeApiTests := new(UnorderedTests)

	automtlsIstioEdgeApiTests.Register("HeadlessSvc", headless_svc.NewEdgeGatewayHeadlessSvcSuite)
	automtlsIstioEdgeApiTests.Register("IstioIntegrationAutoMtls", istio.NewGlooIstioAutoMtlsSuite)

	return automtlsIstioEdgeApiTests
}
