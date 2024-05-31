package listener_options

import (
	"context"
	"time"

	adminv3 "github.com/envoyproxy/go-control-plane/envoy/admin/v3"
	listenerv3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	"github.com/onsi/gomega"
	"github.com/stretchr/testify/suite"

	"github.com/solo-io/gloo/pkg/utils/envoyutils/admincli"
	"github.com/solo-io/gloo/pkg/utils/kubeutils"
	"github.com/solo-io/gloo/pkg/utils/requestutils/curl"
	"github.com/solo-io/gloo/test/kubernetes/e2e"
	testdefaults "github.com/solo-io/gloo/test/kubernetes/e2e/defaults"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// testingSuite is the entire Suite of tests for the "ListenerOptions" feature
type testingSuite struct {
	suite.Suite
	ctx              context.Context
	testInstallation *e2e.TestInstallation
	// maps test name to a list of manifests to apply before the test
	manifests map[string][]string
}

func NewTestingSuite(
	ctx context.Context,
	testInst *e2e.TestInstallation,
) suite.TestingSuite {
	return &testingSuite{
		ctx:              ctx,
		testInstallation: testInst,
	}
}

func (s *testingSuite) SetupSuite() {
	// Check that the common setup manifest is applied
	err := s.testInstallation.Actions.Kubectl().ApplyFile(s.ctx, setupManifest)
	s.NoError(err, "can apply "+setupManifest)
	s.testInstallation.Assertions.EventuallyObjectsExist(s.ctx, proxyService, proxyDeployment, exampleSvc, nginxPod)
	// Check that test resources are running
	s.testInstallation.Assertions.EventuallyPodsRunning(s.ctx, nginxPod.ObjectMeta.GetNamespace(), metav1.ListOptions{
		LabelSelector: "app.kubernetes.io/name=nginx",
	})
	s.testInstallation.Assertions.EventuallyPodsRunning(s.ctx, proxyDeployment.ObjectMeta.GetNamespace(), metav1.ListOptions{
		LabelSelector: "app.kubernetes.io/name=gloo-proxy-gw",
	})

	s.manifests = map[string][]string{
		"TestConfigureListenerOptions": {basicLisOptManifest},
	}
}

func (s *testingSuite) TearDownSuite() {
	// Check that the common setup manifest is deleted
	output, err := s.testInstallation.Actions.Kubectl().DeleteFileWithOutput(s.ctx, setupManifest)
	s.testInstallation.Assertions.ExpectObjectDeleted(setupManifest, err, output)
}

func (s *testingSuite) BeforeTest(suiteName, testName string) {
	manifests, ok := s.manifests[testName]
	if !ok {
		s.FailNow("no manifests found for %s, manifest map contents: %v", testName, s.manifests)
	}

	for _, manifest := range manifests {
		err := s.testInstallation.Actions.Kubectl().ApplyFile(s.ctx, manifest)
		s.Assert().NoError(err, "can apply manifest "+manifest)
	}
}

func (s *testingSuite) AfterTest(suiteName, testName string) {
	manifests, ok := s.manifests[testName]
	if !ok {
		s.FailNow("no manifests found for " + testName)
	}

	for _, manifest := range manifests {
		output, err := s.testInstallation.Actions.Kubectl().DeleteFileWithOutput(s.ctx, manifest)
		s.testInstallation.Assertions.ExpectObjectDeleted(manifest, err, output)
	}
}

func (s *testingSuite) TestConfigureListenerOptions() {
	// Check healthy response
	s.testInstallation.Assertions.AssertEventualCurlResponse(
		s.ctx,
		testdefaults.CurlPodExecOpt,
		[]curl.Option{
			curl.WithHost(kubeutils.ServiceFQDN(proxyService.ObjectMeta)),
			curl.WithHostHeader("example.com"),
		},
		expectedHealthyResponse)

	// Check the buffer limit is set on the Listener via Envoy config dump
	s.testInstallation.Assertions.AssertEnvoyAdminApi(
		s.ctx,
		proxyDeployment.ObjectMeta,
		listenerBufferLimitAssertion(s.testInstallation),
	)
}

func listenerBufferLimitAssertion(testInstallation *e2e.TestInstallation) func(ctx context.Context, adminClient *admincli.Client) {
	return func(ctx context.Context, adminClient *admincli.Client) {
		testInstallation.Assertions.Gomega.Eventually(func(g gomega.Gomega) {
			queryParams := map[string]string{
				"resource":   "dynamic_listeners",
				"name_regex": "http",
			}
			cfgDump, err := adminClient.GetConfigDump(ctx, queryParams)
			g.Expect(err).NotTo(gomega.HaveOccurred(), "could not get envoy config_dump from adminClient")
			g.Expect(cfgDump.GetConfigs()).To(gomega.HaveLen(1))

			listenerDump := adminv3.ListenersConfigDump_DynamicListener{}
			err = cfgDump.GetConfigs()[0].UnmarshalTo(&listenerDump)
			g.Expect(err).NotTo(gomega.HaveOccurred(), "could not unmarshal envoy config_dump")

			listener := listenerv3.Listener{}
			err = listenerDump.GetActiveState().GetListener().UnmarshalTo(&listener)
			g.Expect(err).NotTo(gomega.HaveOccurred(), "could not unmarshal listener from listener dump")
			g.Expect(listener.GetPerConnectionBufferLimitBytes().GetValue()).To(gomega.BeEquivalentTo(42000))
		}).
			WithContext(ctx).
			WithTimeout(time.Second * 10).
			WithPolling(time.Millisecond * 200).
			Should(gomega.Succeed())
	}
}
