package istio

import (
	"context"

	"github.com/onsi/gomega"
	"github.com/solo-io/gloo/test/kubernetes/testutils/helm"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/stretchr/testify/suite"

	"github.com/solo-io/gloo/pkg/utils/kubeutils"
	"github.com/solo-io/gloo/pkg/utils/requestutils/curl"
	"github.com/solo-io/gloo/test/kubernetes/e2e"
)

// istioTestingSuite is the entire Suite of tests for the "Istio" integration cases
type istioTestingSuite struct {
	suite.Suite

	ctx context.Context

	// testInstallation contains all the metadata/utilities necessary to execute a series of tests
	// against an installation of Gloo Gateway
	testInstallation *e2e.TestInstallation

	// helmOptions contains the options that are passed to the helm command
	helmOptions helm.InstallOptions
}

func NewTestingSuite(ctx context.Context, testInst *e2e.TestInstallation, helpOptions helm.InstallOptions) suite.TestingSuite {
	return &istioTestingSuite{
		ctx:              ctx,
		testInstallation: testInst,
		helmOptions:      helpOptions,
	}
}

func (s *istioTestingSuite) SetupSuite() {
	err := s.testInstallation.Actions.Kubectl().ApplyFile(s.ctx, setupManifest)
	s.NoError(err, "can apply setup manifest")

	// Ensure that the proxy service and deployment are created
	err = s.testInstallation.Actions.Kubectl().ApplyFile(s.ctx, k8sRoutingManifest)
	s.NoError(err, "can apply k8s routing manifest")
	s.testInstallation.Assertions.EventuallyObjectsExist(s.ctx, proxyService, proxyDeployment)
}

func (s *istioTestingSuite) TearDownSuite() {
	err := s.testInstallation.Actions.Kubectl().DeleteFile(s.ctx, k8sRoutingManifest)
	s.NoError(err, "can apply k8s routing manifest")
	s.testInstallation.Assertions.EventuallyObjectsNotExist(s.ctx, proxyService, proxyDeployment)

	err = s.testInstallation.Actions.Kubectl().DeleteFile(s.ctx, setupManifest)
	s.NoError(err, "can delete setup manifest")
}

func (s *istioTestingSuite) BeforeEach() {
	// ensure that auto mtls is enabled
	upgradeOpts := s.helmOptions
	upgradeOpts.ExtraArgs = []string{
		"--set", "global.istioIntegration.enableAutoMtls=true",
		// TODO: why is this getting overwritten?
		"--set", "kubeGateway.enabled=true",
	}
	err := helm.HelmUpgradeInstallGloo(upgradeOpts)
	s.NoError(err, "can upgrade gloo with automtls disabled")

	gomega.Eventually(func(g gomega.Gomega) {
		settings, err := s.testInstallation.ResourceClients.SettingsClient().Read(s.testInstallation.Metadata.InstallNamespace, "default", clients.ReadOpts{})
		g.Expect(err).NotTo(gomega.HaveOccurred(), "can read settings")
		g.Expect(settings.GetGloo().GetIstioOptions().GetEnableAutoMtls().GetValue()).To(gomega.BeTrue(), "settings have auto mtls enabled")
	}).Should(gomega.Succeed(), "settings have auto mtls enabled")
}

func (s *istioTestingSuite) TestStrictPeerAuth() {
	s.T().Cleanup(func() {
		err := s.testInstallation.Actions.Kubectl().DeleteFile(s.ctx, strictPeerAuthManifest)
		s.NoError(err, "can delete manifest")
	})

	err := s.testInstallation.Actions.Kubectl().ApplyFile(s.ctx, strictPeerAuthManifest)
	s.NoError(err, "can apply strictPeerAuthManifest")

	s.testInstallation.Assertions.AssertEventualCurlResponse(
		s.ctx,
		curlPodExecOpt,
		[]curl.Option{
			curl.WithHost(kubeutils.ServiceFQDN(proxyService.ObjectMeta)),
			curl.WithHostHeader("httpbin"),
			curl.WithPath("/headers"),
		},
		expectedMtlsResponse)

	// Disable automtls
	upgradeOpts := s.helmOptions
	upgradeOpts.ExtraArgs = []string{
		"--set", "global.istioIntegration.enableAutoMtls=false",
		// TODO: why is this getting overwritten?
		"--set", "kubeGateway.enabled=true",
	}
	err = helm.HelmUpgradeInstallGloo(upgradeOpts)
	s.NoError(err, "can upgrade gloo with automtls disabled")

	gomega.Eventually(func(g gomega.Gomega) {
		settings, err := s.testInstallation.ResourceClients.SettingsClient().Read(s.testInstallation.Metadata.InstallNamespace, "default", clients.ReadOpts{})
		g.Expect(err).NotTo(gomega.HaveOccurred(), "can read settings")
		g.Expect(settings.GetGloo().GetIstioOptions().GetEnableAutoMtls().GetValue()).To(gomega.BeFalse(), "settings have auto mtls disabled")
	}).Should(gomega.Succeed(), "settings have auto mtls disabled")

	// With auto mtls disabled in the mesh, the request should fail when the strict peer auth policy is applied
	s.testInstallation.Assertions.AssertEventualCurlResponse(
		s.ctx,
		curlPodExecOpt,
		[]curl.Option{
			curl.WithHost(kubeutils.ServiceFQDN(proxyService.ObjectMeta)),
			curl.WithHostHeader("httpbin"),
			curl.WithPath("/headers"),
		},
		expectedServiceUnavailableResponse)
}

func (s *istioTestingSuite) TestPermissivePeerAuth() {
	s.T().Cleanup(func() {
		err := s.testInstallation.Actions.Kubectl().DeleteFile(s.ctx, permissivePeerAuthManifest)
		s.NoError(err, "can delete manifest")
	})

	err := s.testInstallation.Actions.Kubectl().ApplyFile(s.ctx, permissivePeerAuthManifest)
	s.NoError(err, "can apply permissivePeerAuth")

	// With auto mtls enabled in the mesh, the response should contain the X-Forwarded-Client-Cert header even with permissive mode
	s.testInstallation.Assertions.AssertEventualCurlResponse(
		s.ctx,
		curlPodExecOpt,
		[]curl.Option{
			curl.WithHost(kubeutils.ServiceFQDN(proxyService.ObjectMeta)),
			curl.WithHostHeader("httpbin"),
			curl.WithPath("/headers"),
		},
		expectedMtlsResponse)

	// Disable automtls
	upgradeOpts := s.helmOptions
	upgradeOpts.ExtraArgs = []string{
		"--set", "global.istioIntegration.enableAutoMtls=false",
		// TODO: why is this getting overwritten?
		"--set", "kubeGateway.enabled=true",
	}
	err = helm.HelmUpgradeInstallGloo(upgradeOpts)
	s.NoError(err, "can upgrade gloo with automtls disabled")

	gomega.Eventually(func(g gomega.Gomega) {
		settings, err := s.testInstallation.ResourceClients.SettingsClient().Read(s.testInstallation.Metadata.InstallNamespace, "default", clients.ReadOpts{})
		g.Expect(err).NotTo(gomega.HaveOccurred(), "can read settings")
		g.Expect(settings.GetGloo().GetIstioOptions().GetEnableAutoMtls().GetValue()).To(gomega.BeFalse(), "settings have auto mtls disabled")
	}).Should(gomega.Succeed(), "settings have auto mtls disabled")

	// With auto mtls disabled in the mesh, the response should not contain the X-Forwarded-Client-Cert header
	s.testInstallation.Assertions.AssertEventualCurlResponse(
		s.ctx,
		curlPodExecOpt,
		[]curl.Option{
			curl.WithHost(kubeutils.ServiceFQDN(proxyService.ObjectMeta)),
			curl.WithHostHeader("httpbin"),
			curl.WithPath("/headers"),
		},
		expectedPlaintextResponse)
}
