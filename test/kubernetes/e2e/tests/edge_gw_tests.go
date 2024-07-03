package tests

import (
	"github.com/solo-io/gloo/test/kubernetes/e2e"
	"github.com/solo-io/gloo/test/kubernetes/e2e/features/glooctl"
	"github.com/solo-io/gloo/test/kubernetes/e2e/features/headless_svc"
	"github.com/solo-io/gloo/test/kubernetes/e2e/features/port_routing"
)

func EdgeGwSuiteRunner() e2e.SuiteRunner {
	edgeGwSuiteRunner := e2e.NewSuiteRunner(false)

	edgeGwSuiteRunner.Register("HeadlessSvc", headless_svc.NewEdgeGatewayHeadlessSvcSuite)
	edgeGwSuiteRunner.Register("PortRouting", port_routing.NewEdgeGatewayApiTestingSuite)
	edgeGwSuiteRunner.Register("GlooctlCheck", glooctl.NewCheckSuite)
	edgeGwSuiteRunner.Register("GlooctlCheckCrds", glooctl.NewCheckCrdsSuite)
	edgeGwSuiteRunner.Register("GlooctlCheckCrds", glooctl.NewDebugSuite)

	return edgeGwSuiteRunner
}
