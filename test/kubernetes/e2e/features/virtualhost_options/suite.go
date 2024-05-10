package virtualhost_options

import (
	"context"
	"strings"

	"github.com/stretchr/testify/suite"

	"github.com/solo-io/gloo/pkg/utils/kubeutils"
	"github.com/solo-io/gloo/pkg/utils/requestutils/curl"
	"github.com/solo-io/gloo/projects/gloo/pkg/defaults"
	"github.com/solo-io/gloo/test/helpers"
	"github.com/solo-io/gloo/test/kubernetes/e2e"
	testdefaults "github.com/solo-io/gloo/test/kubernetes/e2e/defaults"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// testingSuite is the entire Suite of tests for the "VirtualHostOptions" feature
type testingSuite struct {
	suite.Suite

	ctx context.Context

	// testInstallation contains all the metadata/utilities necessary to execute a series of tests
	// against an installation of Gloo Gateway
	testInstallation *e2e.TestInstallation

	// maps test name to a list of manifests to apply before the test
	manifests map[string][]string

	// validationEnabled tracks if the validating webhook was enabled and will reject invalid resources
	validationEnabled bool
}

func NewTestingSuite(
	ctx context.Context,
	testInst *e2e.TestInstallation,
	validationEnabled bool,
) suite.TestingSuite {
	return &testingSuite{
		ctx:               ctx,
		testInstallation:  testInst,
		validationEnabled: validationEnabled,
	}
}

func (s *testingSuite) SetupSuite() {
	// We include tests with manual setup here because the cleanup is still automated via AfterTest
	s.manifests = map[string][]string{
		"TestConfigureVirtualHostOptions":                           {setupManifest, basicVhOManifest},
		"TestConfigureInvalidVirtualHostOptions":                    {setupManifest, basicVhOManifest, badVhOManifest},
		"TestConfigureVirtualHostOptionsWithSectionNameManualSetup": {setupManifest, basicVhOManifest, extraVhOManifest, sectionNameVhOManifest},
		"TestMultipleVirtualHostOptionsManualSetup":                 {setupManifest, basicVhOManifest, extraVhOManifest},
	}
}

func (s *testingSuite) TearDownSuite() {}

func (s *testingSuite) BeforeTest(suiteName, testName string) {
	if strings.Contains(testName, "ManualSetup") {
		return
	}

	manifests, ok := s.manifests[testName]
	if !ok {
		s.FailNow("no manifests found for %s, manifest map contents: %v", testName, s.manifests)
	}

	for _, manifest := range manifests {
		err := s.testInstallation.Actions.Kubectl().ApplyFileWithRunError(s.ctx, manifest)
		if strings.Contains(manifest, "bad") {
			s.ErrorContains(err, "Validating *v1.VirtualHostOption failed")
		} else {
			s.NoError(err, "can apply "+manifest)
		}
	}
}

func (s *testingSuite) AfterTest(suiteName, testName string) {
	manifests, ok := s.manifests[testName]
	if !ok {
		s.Fail("no manifests found for " + testName)
	}

	for _, manifest := range manifests {
		if strings.Contains(manifest, "bad") {
			// this resource was rejected so no need to delete
			continue
		}
		err := s.testInstallation.Actions.Kubectl().DeleteFile(s.ctx, manifest)
		s.NoError(err, "can delete "+manifest)
	}
}

func (s *testingSuite) TestConfigureVirtualHostOptions() {
	// Check healthy response with no content-length header
	s.testInstallation.Assertions.AssertEventualCurlResponse(
		s.ctx,
		testdefaults.CurlPodExecOpt,
		[]curl.Option{
			curl.WithHost(kubeutils.ServiceFQDN(proxyService.ObjectMeta)),
			curl.WithHostHeader("example.com"),
		},
		expectedResponseWithoutContentLength)

	// Check status is accepted on VirtualHostOption
	s.testInstallation.Assertions.EventuallyResourceStatusMatchesState(
		s.getterForMeta(&basicVirtualHostOptionMeta),
		core.Status_Accepted,
		defaults.KubeGatewayReporter,
	)
}

func (s *testingSuite) TestConfigureInvalidVirtualHostOptions() {
	if s.validationEnabled {
		// TODO: assert that resource doesn't exist in cluster
	} else {
		// Check status is rejected on bad VirtualHostOption
		s.testInstallation.Assertions.EventuallyResourceStatusMatchesState(
			s.getterForMeta(&badVirtualHostOptionMeta),
			core.Status_Rejected,
			defaults.KubeGatewayReporter,
		)
	}
}

// The goal here is to test the behavior when multiple VHOs target a gateway with multiple listeners and only some
// conflict. This will generate a warning on the conflicted resource, but the VHO should be attached properly and
// options propagated for the listener.
func (s *testingSuite) TestConfigureVirtualHostOptionsWithSectionNameManualSetup() {
	// Manually apply our manifests so we can assert that basic vho exists before applying extra vho.
	// This is needed because our solo-kit clients currently do not return creationTimestamp
	err := s.testInstallation.Actions.Kubectl().ApplyFile(s.ctx, setupManifest)
	s.NoError(err, "can apply "+setupManifest)

	err = s.testInstallation.Actions.Kubectl().ApplyFile(s.ctx, basicVhOManifest)
	s.NoError(err, "can apply "+basicVhOManifest)
	// Check status is accepted before moving on to apply conflicting vho
	s.testInstallation.Assertions.EventuallyResourceStatusMatchesState(
		s.getterForMeta(&basicVirtualHostOptionMeta),
		core.Status_Accepted,
		defaults.KubeGatewayReporter,
	)

	err = s.testInstallation.Actions.Kubectl().ApplyFile(s.ctx, extraVhOManifest)
	s.NoError(err, "can apply "+extraVhOManifest)

	err = s.testInstallation.Actions.Kubectl().ApplyFile(s.ctx, sectionNameVhOManifest)
	s.NoError(err, "can apply "+sectionNameVhOManifest)

	// Check healthy response with added foo header to listener targeted by sectionName
	s.testInstallation.Assertions.AssertEventualCurlResponse(
		s.ctx,
		testdefaults.CurlPodExecOpt,
		[]curl.Option{
			curl.WithHost(kubeutils.ServiceFQDN(proxyService.ObjectMeta)),
			curl.WithHostHeader("example.com"),
			curl.WithPort(8080),
		},
		expectedResponseWithFooHeader)

	// Check healthy response with content-length removed to listener NOT targeted by sectionName
	s.testInstallation.Assertions.AssertEventualCurlResponse(
		s.ctx,
		testdefaults.CurlPodExecOpt,
		[]curl.Option{
			curl.WithHost(kubeutils.ServiceFQDN(proxyService.ObjectMeta)),
			curl.WithHostHeader("example.com"),
			curl.WithPort(8081),
		},
		expectedResponseWithoutContentLength)

	// Check status is accepted on VirtualHostOption with section name
	s.testInstallation.Assertions.EventuallyResourceStatusMatchesState(
		s.getterForMeta(&sectionNameVirtualHostOptionMeta),
		core.Status_Accepted,
		defaults.KubeGatewayReporter,
	)
	// Check status is warning on VirtualHostOption with conflicting attachment,
	// despite being properly attached to another listener
	s.testInstallation.Assertions.EventuallyResourceStatusMatchesWarningReasons(
		s.getterForMeta(&basicVirtualHostOptionMeta),
		[]string{"conflict with more-specific or older VirtualHostOption"},
		defaults.KubeGatewayReporter,
	)

	// Check status is warning on VirtualHostOption not selected for attachment
	s.testInstallation.Assertions.EventuallyResourceStatusMatchesWarningReasons(
		s.getterForMeta(&extraVirtualHostOptionMeta),
		[]string{"conflict with more-specific or older VirtualHostOption"},
		defaults.KubeGatewayReporter,
	)
}

// The goal here is to test the behavior when multiple VHOs are targeting a gateway without sectionName. The expected
// behavior is that the oldest resource is used
func (s *testingSuite) TestMultipleVirtualHostOptionsManualSetup() {
	// Manually apply our manifests so we can assert that basic vho exists before applying extra vho.
	// This is needed because our solo-kit clients currently do not return creationTimestamp
	err := s.testInstallation.Actions.Kubectl().ApplyFile(s.ctx, setupManifest)
	s.NoError(err, "can apply "+setupManifest)

	err = s.testInstallation.Actions.Kubectl().ApplyFile(s.ctx, basicVhOManifest)
	s.NoError(err, "can apply "+basicVhOManifest)
	// Check status is accepted before moving on to apply conflicting vho
	s.testInstallation.Assertions.EventuallyResourceStatusMatchesState(
		s.getterForMeta(&basicVirtualHostOptionMeta),
		core.Status_Accepted,
		defaults.KubeGatewayReporter,
	)

	err = s.testInstallation.Actions.Kubectl().ApplyFile(s.ctx, extraVhOManifest)
	s.NoError(err, "can apply "+extraVhOManifest)

	// Check healthy response with no content-length header
	s.testInstallation.Assertions.AssertEventualCurlResponse(
		s.ctx,
		testdefaults.CurlPodExecOpt,
		[]curl.Option{
			curl.WithHost(kubeutils.ServiceFQDN(proxyService.ObjectMeta)),
			curl.WithHostHeader("example.com"),
		},
		expectedResponseWithoutContentLength)

	// Check status is accepted on older VirtualHostOption
	s.testInstallation.Assertions.EventuallyResourceStatusMatchesState(
		s.getterForMeta(&basicVirtualHostOptionMeta),
		core.Status_Accepted,
		defaults.KubeGatewayReporter,
	)
	// Check status is warning on newer VirtualHostOption not selected for attachment
	s.testInstallation.Assertions.EventuallyResourceStatusMatchesWarningReasons(
		s.getterForMeta(&extraVirtualHostOptionMeta),
		[]string{"conflict with more-specific or older VirtualHostOption"},
		defaults.KubeGatewayReporter,
	)
}

func (s *testingSuite) getterForMeta(meta *metav1.ObjectMeta) helpers.InputResourceGetter {
	return func() (resources.InputResource, error) {
		return s.testInstallation.ResourceClients.VirtualHostOptionClient().Read(meta.GetNamespace(), meta.GetName(), clients.ReadOpts{})
	}
}
