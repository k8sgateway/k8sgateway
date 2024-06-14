package tests

import (
	"github.com/solo-io/gloo/test/kubernetes/e2e"
	"github.com/solo-io/gloo/test/kubernetes/e2e/features/listener_options"
	"github.com/solo-io/gloo/test/kubernetes/e2e/features/port_routing"
	"github.com/solo-io/gloo/test/kubernetes/e2e/features/route_options"
	"github.com/solo-io/gloo/test/kubernetes/e2e/features/virtualhost_options"
)

func KubeGatewayNoValidationTests() e2e.SuiteRunner {
	kubeGatewayNoValidationTests := e2e.NewSuiteRunner(false)

	kubeGatewayNoValidationTests.Register("ListenerOptions", listener_options.NewTestingSuite)
	kubeGatewayNoValidationTests.Register("RouteOptions", route_options.NewTestingSuite)
	kubeGatewayNoValidationTests.Register("VirtualHostOptions", virtualhost_options.NewTestingSuite)
	kubeGatewayNoValidationTests.Register("PortRouting", port_routing.NewTestingSuite)

	return kubeGatewayNoValidationTests
}
