package discovery_watchlabels

import (
	"context"
	"fmt"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/kube/apis/gloo.solo.io/v1"
	"github.com/solo-io/skv2/codegen/util"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"path/filepath"
	"time"

	"github.com/solo-io/gloo/projects/gloo/pkg/defaults"
	"github.com/solo-io/gloo/test/kubernetes/e2e"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/stretchr/testify/suite"
)

var _ e2e.NewSuiteFunc = NewDiscoveryWatchlabelsSuite

// discoveryWatchlabelsSuite is the Suite of tests for validating Upstream discovery behavior when watchLabels are enabled
type discoveryWatchlabelsSuite struct {
	suite.Suite

	ctx context.Context

	// testInstallation contains all the metadata/utilities necessary to execute a series of tests
	// against an installation of Gloo Gateway
	testInstallation *e2e.TestInstallation
}

func NewDiscoveryWatchlabelsSuite(ctx context.Context, testInst *e2e.TestInstallation) suite.TestingSuite {
	return &discoveryWatchlabelsSuite{
		ctx:              ctx,
		testInstallation: testInst,
	}
}

func (s *discoveryWatchlabelsSuite) TestDiscoverUpstreamMatchingWatchLabels() {
	s.T().Cleanup(func() {
		err := s.testInstallation.Actions.Kubectl().DeleteFile(s.ctx, filepath.Join(util.MustGetThisDir(), "testdata/service-with-labels.yaml"), "-n", s.testInstallation.Metadata.InstallNamespace)
		s.Assertions.NoError(err, "can delete service")

		err = s.testInstallation.Actions.Kubectl().DeleteFile(s.ctx, filepath.Join(util.MustGetThisDir(), "testdata/service-without-labels.yaml"), "-n", s.testInstallation.Metadata.InstallNamespace)
		s.Assertions.NoError(err, "can delete service")
	})

	// add one service with labels matching our watchLabels
	err := s.testInstallation.Actions.Kubectl().ApplyFile(s.ctx, filepath.Join(util.MustGetThisDir(), "testdata/service-with-labels.yaml"), "-n", s.testInstallation.Metadata.InstallNamespace)
	s.Assert().NoError(err, "can apply service")

	// add one service without labels matching our watchLabels
	err = s.testInstallation.Actions.Kubectl().ApplyFile(s.ctx, filepath.Join(util.MustGetThisDir(), "testdata/service-without-labels.yaml"), "-n", s.testInstallation.Metadata.InstallNamespace)
	s.Assert().NoError(err, "can apply service")

	// eventually an Upstream should be created for the Service with labels
	labeledUsName := fmt.Sprintf("%s-%s-%d", s.testInstallation.Metadata.InstallNamespace, "example-svc", 8000)
	s.testInstallation.Assertions.EventuallyResourceStatusMatchesState(
		func() (resources.InputResource, error) {
			return s.testInstallation.ResourceClients.UpstreamClient().Read(s.testInstallation.Metadata.InstallNamespace, labeledUsName, clients.ReadOpts{Ctx: s.ctx})
		},
		core.Status_Accepted,
		defaults.GlooReporter,
	)

	// the Upstream should have DiscoveryMetadata labels matching the parent Service
	us, err := s.testInstallation.ResourceClients.UpstreamClient().Read(s.testInstallation.Metadata.InstallNamespace, labeledUsName, clients.ReadOpts{Ctx: s.ctx})
	s.Assert().NoError(err, "can read upstream")

	s.Assert().Equal(map[string]string{
		"watchedKey": "watchedValue",
		"bonusKey":   "bonusValue",
	}, us.DiscoveryMetadata.Labels)

	// no Upstream should be created for the Service that does not have the watchLabels
	noLabelsUsName := fmt.Sprintf("%s-%s-%d", s.testInstallation.Metadata.InstallNamespace, "example-svc-no-labels", 8000)
	s.testInstallation.Assertions.ConsistentlyObjectsNotExist(
		s.ctx, &v1.Upstream{
			ObjectMeta: metav1.ObjectMeta{
				Name:      noLabelsUsName,
				Namespace: s.testInstallation.Metadata.InstallNamespace,
			},
		},
	)

	// modify the non-watched label on the labeled service
	err = s.testInstallation.Actions.Kubectl().ApplyFile(s.ctx, filepath.Join(util.MustGetThisDir(), "testdata/service-with-modified-labels.yaml"), "-n", s.testInstallation.Metadata.InstallNamespace)
	s.Assert().NoError(err, "can re-apply service")

	// expect the Upstream's DiscoveryMeta to eventually match the modified labels from the parent Service
	time.Sleep(10 * time.Second) // would rather do an Eventually

	us, err = s.testInstallation.ResourceClients.UpstreamClient().Read(s.testInstallation.Metadata.InstallNamespace, labeledUsName, clients.ReadOpts{Ctx: s.ctx})
	s.Assert().NoError(err, "can read upstream")

	s.Assert().Equal(map[string]string{
		"watchedKey": "watchedValue",
		"bonusKey":   "bonusValue-modified",
	}, us.DiscoveryMetadata.Labels)
}
