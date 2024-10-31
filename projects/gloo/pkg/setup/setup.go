package setup

import (
	"context"

	"github.com/solo-io/gloo/pkg/bootstrap/leaderelector"
	"github.com/solo-io/gloo/pkg/utils/envutils"
	"github.com/solo-io/gloo/pkg/utils/namespaces"
	"github.com/solo-io/gloo/pkg/utils/setuputils"
	"github.com/solo-io/gloo/pkg/version"
	"github.com/solo-io/gloo/projects/gateway2/extensions"
	ggv2setup "github.com/solo-io/gloo/projects/gateway2/setup"
	ggv2utils "github.com/solo-io/gloo/projects/gateway2/utils"
	"github.com/solo-io/gloo/projects/gloo/constants"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/bootstrap"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins/registry"
	"github.com/solo-io/gloo/projects/gloo/pkg/syncer/setup"
	"github.com/solo-io/gloo/projects/gloo/pkg/xds"
	"github.com/solo-io/go-utils/contextutils"
)

const (
	glooComponentName = "gloo"
)

func Main(customCtx context.Context) error {
	setuputils.SetupLogging(customCtx, glooComponentName)
	return startSetupLoop(customCtx)
}

func startSetupLoop(ctx context.Context) error {

	// get settings:

	setupOpts := bootstrap.NewSetupOpts(xds.NewAdsSnapshotCache(ctx))
	// start gw if needed, get the proxy reconcile q
	// pass that in to the setup func
	k8sgw := envutils.IsEnvTruthy(constants.GlooGatewayEnableK8sGwControllerEnv)
	if k8sgw {
		setupOpts.ProxyReconcileQueue = ggv2utils.NewAsyncQueue[gloov1.ProxyList]()
		go ggv2setup.StartGGv2(ctx, setupOpts, extensions.NewK8sGatewayExtensions, registry.GetPluginRegistryFactory)
	}

	return setuputils.Main(setuputils.SetupOpts{
		LoggerName:  glooComponentName,
		Version:     version.Version,
		SetupFunc:   newSetupFunc(setupOpts),
		ExitOnError: true,
		CustomCtx:   ctx,

		ElectionConfig: &leaderelector.ElectionConfig{
			Id:        glooComponentName,
			Namespace: namespaces.GetPodNamespace(),
			// no-op all the callbacks for now
			// at the moment, leadership functionality is performed within components
			// in the future we could pull that out and let these callbacks change configuration
			OnStartedLeading: func(c context.Context) {
				contextutils.LoggerFrom(c).Info("starting leadership")
			},
			OnNewLeader: func(leaderId string) {
				contextutils.LoggerFrom(ctx).Infof("new leader elected with ID: %s", leaderId)
			},
			OnStoppedLeading: func() {
				// Don't die if we fall from grace. Instead we can retry leader election
				// Ref: https://github.com/solo-io/gloo/issues/7346
				contextutils.LoggerFrom(ctx).Errorf("lost leadership")
			},
		},
	})
}

func newSetupFunc(setupOpts *bootstrap.SetupOpts) setuputils.SetupFunc {

	runFunc := func(opts bootstrap.Opts) error {
		return setup.RunGloo(opts)
	}

	return setup.NewSetupFuncWithRunAndExtensions(runFunc, setupOpts, nil)
}
