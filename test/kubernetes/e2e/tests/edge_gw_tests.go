package tests

import (
	"github.com/solo-io/gloo/test/kubernetes/e2e"
	"github.com/solo-io/gloo/test/kubernetes/e2e/features/glooctl"
)

func EdgeGwSuiteRunner() e2e.SuiteRunner {
	edgeGwSuiteRunner := e2e.NewSuiteRunner(false)

	//edgeGwSuiteRunner.Register("HeadlessSvc", headless_svc.NewEdgeGatewayHeadlessSvcSuite)
	//edgeGwSuiteRunner.Register("PortRouting", port_routing.NewEdgeGatewayApiTestingSuite)
	edgeGwSuiteRunner.Register("GlooctlCheck", glooctl.NewCheckSuite)
	//edgeGwSuiteRunner.Register("GlooctlCheckCrds", glooctl.NewCheckCrdsSuite)
	//edgeGwSuiteRunner.Register("GlooctlDebug", glooctl.NewDebugSuite)

	return edgeGwSuiteRunner
}
