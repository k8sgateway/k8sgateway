package headless_svc

import (
	"context"

	"github.com/stretchr/testify/suite"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/solo-io/gloo/pkg/utils/kubeutils"
	"github.com/solo-io/gloo/pkg/utils/requestutils/curl"
	"github.com/solo-io/gloo/test/kubernetes/e2e"
	"github.com/solo-io/gloo/test/kubernetes/e2e/utils"
)

type k8sGatewaySuite struct {
	suite.Suite

	ctx context.Context

	// testInstallation contains all the metadata/utilities necessary to execute a series of tests
	// against an installation of Gloo Gateway
	testInstallation *e2e.TestInstallation
}

func NewK8sGatewayHeadlessSvcSuite(ctx context.Context, testInst *e2e.TestInstallation) suite.TestingSuite {
	return &k8sGatewaySuite{
		ctx:              ctx,
		testInstallation: testInst,
	}
}

// SetupSuite generates manifest files for the test suite
func (s *k8sGatewaySuite) SetupSuite() {
	// use the k8s gateway api resources
	resources := []client.Object{gw, headlessSvcHTTPRoute}
	err := utils.WriteResourcesToFile(resources, k8sApiRoutingManifest)
	s.Require().NoError(err, "can write resources to file")
}

func (s *k8sGatewaySuite) TestConfigureRoutingHeadlessSvc() {
	s.T().Cleanup(func() {
		err := s.testInstallation.Actions.Kubectl().DeleteFile(s.ctx, headlessSvcSetupManifest)
		s.NoError(err, "can delete setup manifest")
		s.testInstallation.Assertions.EventuallyObjectsNotExist(s.ctx, headlessService)

		err = s.testInstallation.Actions.Kubectl().DeleteFile(s.ctx, k8sApiRoutingManifest)
		s.NoError(err, "can delete setup k8s routing manifest")
		s.testInstallation.Assertions.EventuallyObjectsNotExist(s.ctx, k8sApiProxyDeployment, k8sApiProxyService)
	})

	err := s.testInstallation.Actions.Kubectl().ApplyFile(s.ctx, headlessSvcSetupManifest)
	s.Assert().NoError(err, "can apply setup manifest")
	s.testInstallation.Assertions.EventuallyObjectsExist(s.ctx, headlessService)

	err = s.testInstallation.Actions.Kubectl().ApplyFile(s.ctx, k8sApiRoutingManifest)
	s.NoError(err, "can setup k8s routing manifest")

	s.testInstallation.Assertions.EventuallyObjectsExist(s.ctx, k8sApiProxyDeployment, k8sApiProxyService)

	s.testInstallation.Assertions.AssertEventualCurlResponse(
		s.ctx,
		curlPodExecOpt,
		[]curl.Option{
			curl.WithHost(kubeutils.ServiceFQDN(k8sApiProxyService.ObjectMeta)),
			// The host header must match the hostnames in the HTTPRoute
			curl.WithHostHeader(headlessSvcDomain),
			curl.WithPort(80),
		},
		expectedHealthyResponse)
}
