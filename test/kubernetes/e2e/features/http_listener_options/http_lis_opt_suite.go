package http_listener_options

import (
	"context"
	"net/http"

	"github.com/onsi/gomega"
	"github.com/stretchr/testify/suite"

	"github.com/solo-io/gloo/pkg/utils/kubeutils"
	"github.com/solo-io/gloo/pkg/utils/requestutils/curl"
	"github.com/solo-io/gloo/test/gomega/matchers"
	"github.com/solo-io/gloo/test/kubernetes/e2e"
	testdefaults "github.com/solo-io/gloo/test/kubernetes/e2e/defaults"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ e2e.NewSuiteFunc = NewTestingSuite

// testingSuite is the entire Suite of tests for the "HttpListenerOptions" feature
type testingSuite struct {
	//suite.Suite
	testdefaults.CommonTestSuiteImpl
	// ctx              context.Context
	// testInstallation *e2e.TestInstallation
	// // maps test name to a list of manifests to apply before the test
	manifests map[string][]string
}

// DO_NOT_SUBMIT: Better with embedding or writing the methods directly like intest/kubernetes/e2e/features/headless_svc/gloo_gateway_suite.go?
func NewTestingSuite(
	ctx context.Context,
	testInst *e2e.TestInstallation,
) suite.TestingSuite {
	return &testingSuite{
		CommonTestSuiteImpl: *testdefaults.NewCommonTestSuiteImpl(
			ctx,
			testInst,
			[]testdefaults.Resource{
				&testdefaults.CurlPodResource{},
			},
		),
	}
}

func (s *testingSuite) SetupSuite() {
	// Check that the common setup manifest is applied
	err := s.TestInstallation().Actions.Kubectl().ApplyFile(s.Ctx(), setupManifest)
	s.NoError(err, "can apply "+setupManifest)
	s.TestInstallation().Assertions.EventuallyObjectsExist(s.Ctx(), exampleSvc, nginxPod)
	// Check that test app is running
	s.TestInstallation().Assertions.EventuallyPodsRunning(s.Ctx(), nginxPod.ObjectMeta.GetNamespace(), metav1.ListOptions{
		LabelSelector: "app.kubernetes.io/name=nginx",
	})

	testdefaults.InstallResources(s)

	// include gateway manifests for the tests, so we recreate it for each test run
	s.manifests = map[string][]string{
		"TestConfigureHttpListenerOptions":            {gatewayManifest, basicLisOptManifest},
		"TestConfigureNotAttachedHttpListenerOptions": {gatewayManifest, notAttachedLisOptManifest},
	}
}

func (s *testingSuite) TearDownSuite() {
	// Check that the common setup manifest is deleted
	output, err := s.TestInstallation().Actions.Kubectl().DeleteFileWithOutput(s.Ctx(), setupManifest)
	s.TestInstallation().Assertions.ExpectObjectDeleted(setupManifest, err, output)

	testdefaults.DeleteResources(s)
}

func (s *testingSuite) BeforeTest(suiteName, testName string) {
	manifests, ok := s.manifests[testName]
	if !ok {
		s.FailNow("no manifests found for %s, manifest map contents: %v", testName, s.manifests)
	}

	for _, manifest := range manifests {
		err := s.TestInstallation().Actions.Kubectl().ApplyFile(s.Ctx(), manifest)
		s.Assert().NoError(err, "can apply manifest "+manifest)
	}

	// we recreate the `Gateway` resource (and thus dynamically provision the proxy pod) for each test run
	// so let's assert the proxy svc and pod is ready before moving on
	s.TestInstallation().Assertions.EventuallyObjectsExist(s.Ctx(), proxyService, proxyDeployment)
	s.TestInstallation().Assertions.EventuallyPodsRunning(s.Ctx(), proxyDeployment.ObjectMeta.GetNamespace(), metav1.ListOptions{
		LabelSelector: "app.kubernetes.io/name=gloo-proxy-gw",
	})
}

func (s *testingSuite) AfterTest(suiteName, testName string) {
	manifests, ok := s.manifests[testName]
	if !ok {
		s.FailNow("no manifests found for " + testName)
	}

	for _, manifest := range manifests {
		output, err := s.TestInstallation().Actions.Kubectl().DeleteFileWithOutput(s.Ctx(), manifest)
		s.TestInstallation().Assertions.ExpectObjectDeleted(manifest, err, output)
	}
}

func (s *testingSuite) TestConfigureHttpListenerOptions() {
	// Check healthy response and response headers contain server name override from HttpListenerOption
	s.TestInstallation().Assertions.AssertEventualCurlResponse(
		s.Ctx(),
		testdefaults.CurlPodExecOpt,
		[]curl.Option{
			curl.WithHost(kubeutils.ServiceFQDN(proxyService.ObjectMeta)),
			curl.WithHostHeader("example.com"),
		},
		&matchers.HttpResponse{
			StatusCode: http.StatusOK,
			Body:       gomega.ContainSubstring("Welcome to nginx!"),
			Headers: map[string]interface{}{
				"server": "unit-test v4.19",
			},
		})
}

func (s *testingSuite) TestConfigureNotAttachedHttpListenerOptions() {
	// Check healthy response and response headers contain default server name as HttpLisOpt isn't attached
	s.TestInstallation().Assertions.AssertEventualCurlResponse(
		s.Ctx(),
		testdefaults.CurlPodExecOpt,
		[]curl.Option{
			curl.WithHost(kubeutils.ServiceFQDN(proxyService.ObjectMeta)),
			curl.WithHostHeader("example.com"),
		},
		&matchers.HttpResponse{
			StatusCode: http.StatusOK,
			Body:       gomega.ContainSubstring("Welcome to nginx!"),
			Headers: map[string]interface{}{
				"server": "envoy",
			},
		})
}
