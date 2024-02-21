package controller

import (
	"os"

	"github.com/solo-io/gloo/projects/gloo/pkg/bootstrap"

	"github.com/solo-io/gloo/projects/gateway2/controller/scheme"
	"github.com/solo-io/gloo/projects/gateway2/discovery"
	"github.com/solo-io/gloo/projects/gateway2/secrets"
	"github.com/solo-io/gloo/projects/gateway2/xds"
	"github.com/solo-io/gloo/projects/gloo/pkg/syncer/sanitizer"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
	apiv1 "sigs.k8s.io/gateway-api/apis/v1"
)

var (
	setupLog = ctrl.Log.WithName("setup")
)

type ControllerConfig struct {
	// The name of the GatewayClass to watch for
	GatewayClassName      string
	GatewayControllerName string
	Release               string
	AutoProvision         bool
	Dev                   bool

	ControlPlane bootstrap.ControlPlane
}

func Start(cfg ControllerConfig) {
	setupLog.Info("xxxxx starting gw2 controller xxxxxx")
	var opts []zap.Opts
	if cfg.Dev {
		setupLog.Info("starting log in dev mode")
		opts = append(opts, zap.UseDevMode(true))
	}
	ctrl.SetLogger(zap.New(opts...))
	mgrOpts := ctrl.Options{
		Scheme: scheme.NewScheme(),
		// PprofBindAddress: "127.0.0.1:9099",
		// // if you change the port here, also change the port "health" in the helmchart.
		// HealthProbeBindAddress: ":9093",
		// Metrics: metricsserver.Options{
		// 	BindAddress: ":9092",
		// },
	}
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), mgrOpts)
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	// TODO: replace this with something that checks that we have xds snapshot ready (or that we don't need one).
	// mgr.AddReadyzCheck("ready-ping", healthz.Ping)

	ctx := signals.SetupSignalHandler()

	glooTranslator := newGlooTranslator(ctx)
	var sanz sanitizer.XdsSanitizers
	inputChannels := xds.NewXdsInputChannels()
	xdsSyncer := xds.NewXdsSyncer(
		cfg.GatewayControllerName,
		glooTranslator,
		sanz,
		cfg.ControlPlane.SnapshotCache,
		false,
		inputChannels,
		mgr.GetClient(),
		mgr.GetScheme(),
	)
	if err := mgr.Add(xdsSyncer); err != nil {
		setupLog.Error(err, "unable to add xdsSyncer runnable")
		os.Exit(1)
	}

	// sam-heilbron: I don't think this is necessary, as we should have a shared cache
	// if cfg.Dev {
	// 	go xdsSyncer.ServeXdsSnapshots()
	// }

	var gatewayClassName apiv1.ObjectName = apiv1.ObjectName(cfg.GatewayClassName)

	gwcfg := GatewayConfig{
		Mgr:            mgr,
		GWClass:        gatewayClassName,
		Dev:            cfg.Dev,
		ControllerName: cfg.GatewayControllerName,
		AutoProvision:  cfg.AutoProvision,
		ControlPlane:   cfg.ControlPlane,
		Kick:           inputChannels.Kick,
	}
	err = NewBaseGatewayController(ctx, gwcfg)

	if err != nil {
		setupLog.Error(err, "unable to create controller")
		os.Exit(1)
	}

	err = discovery.NewDiscoveryController(ctx, mgr, inputChannels)
	if err != nil {
		setupLog.Error(err, "unable to create controller")
		os.Exit(1)
	}

	err = secrets.NewSecretsController(ctx, mgr, inputChannels)
	if err != nil {
		setupLog.Error(err, "unable to create controller")
		os.Exit(1)
	}

	setupLog.Info("starting manager")
	if err := mgr.Start(ctx); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}

}
