package controller

import (
	"context"

	"k8s.io/client-go/rest"
	"k8s.io/utils/ptr"
	"sigs.k8s.io/controller-runtime/pkg/config"

	glooschemes "github.com/solo-io/gloo/pkg/schemes"
	"github.com/solo-io/go-utils/contextutils"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/sets"
	ctrl "sigs.k8s.io/controller-runtime"

	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"
	gwv1a2 "sigs.k8s.io/gateway-api/apis/v1alpha2"

	"github.com/solo-io/gloo/projects/gateway2/deployer"
	"github.com/solo-io/gloo/projects/gateway2/extensions"
	ext "github.com/solo-io/gloo/projects/gateway2/extensions"
	"github.com/solo-io/gloo/projects/gateway2/extensions2/common"
	extensionsplug "github.com/solo-io/gloo/projects/gateway2/extensions2/plugin"
	"github.com/solo-io/gloo/projects/gateway2/extensions2/registry"
	"github.com/solo-io/gloo/projects/gateway2/ir"
	"github.com/solo-io/gloo/projects/gateway2/krtcollections"
	"github.com/solo-io/gloo/projects/gateway2/proxy_syncer"
	"github.com/solo-io/gloo/projects/gateway2/utils/krtutil"
	"github.com/solo-io/gloo/projects/gateway2/wellknown"
	glookubev1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/kube/apis/gloo.solo.io/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/bootstrap"
	"github.com/solo-io/gloo/projects/gloo/pkg/syncer"
	"github.com/solo-io/solo-kit/pkg/api/v2/reporter"
	uzap "go.uber.org/zap"
	istiokube "istio.io/istio/pkg/kube"
	"istio.io/istio/pkg/kube/kclient"
	"istio.io/istio/pkg/kube/krt"
	corev1 "k8s.io/api/core/v1"
)

const (
	// AutoProvision controls whether the controller will be responsible for provisioning dynamic
	// infrastructure for the Gateway API.
	AutoProvision = true
)

var setupLog = ctrl.Log.WithName("setup")

type StartConfig struct {
	Dev        bool
	SetupOpts  *bootstrap.SetupOpts
	RestConfig *rest.Config
	// ExtensionsFactory is the factory function which will return an extensions.K8sGatewayExtensions
	// This is responsible for producing the extension points that this controller requires
	ExtensionsFactory extensions.K8sGatewayExtensionsFactory

	// GlooStatusReporter is the shared reporter from setup_syncer that reports as 'gloo',
	// it is used to report on Upstreams and Proxies after xds translation.
	// this is required because various upstream tests expect a certain reporter for Upstreams
	// TODO: remove the other reporter and only use this one, no need for 2 different reporters
	GlooStatusReporter reporter.StatusReporter

	// KubeGwStatusReporter is used within any StatusPlugins that must persist a GE-classic style status
	// TODO: as mentioned above, this should be removed: https://github.com/solo-io/solo-projects/issues/7055
	KubeGwStatusReporter reporter.StatusReporter

	// SyncerExtensions is a list of extensions, the kube gw controller will use these to get extension-specific
	// errors & warnings for any Proxies it generates
	SyncerExtensions []syncer.TranslatorSyncerExtension

	Client istiokube.Client

	AugmentedPods krt.Collection[krtcollections.LocalityPod]
	UniqueClients krt.Collection[ir.UniqlyConnectedClient]

	InitialSettings *glookubev1.Settings
	Settings        krt.Singleton[glookubev1.Settings]

	KrtOptions krtutil.KrtOptions
}

// Start runs the controllers responsible for processing the K8s Gateway API objects
// It is intended to be run in a goroutine as the function will block until the supplied
// context is cancelled
type ControllerBuilder struct {
	proxySyncer     *proxy_syncer.ProxySyncer
	cfg             StartConfig
	k8sGwExtensions ext.K8sGatewayExtensions
	mgr             ctrl.Manager
}

func NewControllerBuilder(ctx context.Context, cfg StartConfig) (*ControllerBuilder, error) {
	var opts []zap.Opts
	if cfg.Dev {
		setupLog.Info("starting log in dev mode")
		opts = append(opts, zap.UseDevMode(true))
	}
	ctrl.SetLogger(zap.New(opts...))

	scheme := glooschemes.DefaultScheme()

	// Extend the scheme if the TCPRoute CRD exists.
	if err := glooschemes.AddGatewayV1A2Scheme(cfg.RestConfig, scheme); err != nil {
		return nil, err
	}

	mgrOpts := ctrl.Options{
		BaseContext:      func() context.Context { return ctx },
		Scheme:           scheme,
		PprofBindAddress: "127.0.0.1:9099",
		// if you change the port here, also change the port "health" in the helmchart.
		HealthProbeBindAddress: ":9093",
		Metrics: metricsserver.Options{
			BindAddress: ":9092",
		},
		Controller: config.Controller{
			// see https://github.com/kubernetes-sigs/controller-runtime/issues/2937
			// in short, our tests reuse the same name (reasonably so) and the controller-runtime
			// package does not reset the stack of controller names between tests, so we disable
			// the name validation here.
			SkipNameValidation: ptr.To(true),
		},
	}
	mgr, err := ctrl.NewManager(cfg.RestConfig, mgrOpts)
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		return nil, err
	}

	// TODO: replace this with something that checks that we have xds snapshot ready (or that we don't need one).
	mgr.AddReadyzCheck("ready-ping", healthz.Ping)

	//	virtualHostOptionCollection := proxy_syncer.SetupCollectionDynamic[gatewaykubev1.VirtualHostOption](
	//		ctx,
	//		cfg.Client,
	//		gatewaykubev1.SchemeGroupVersion.WithResource("virtualhostoptions"),
	//		krt.WithName("VirtualHostOption"))
	//	routeOptionCollection := proxy_syncer.SetupCollectionDynamic[gatewaykubev1.RouteOption](
	//		ctx,
	//		cfg.Client,
	//		gatewaykubev1.SchemeGroupVersion.WithResource("routeoptions"),
	//		krt.WithName("RouteOption"))
	//	authConfigCollection := proxy_syncer.SetupCollectionDynamic[extauthkubev1.AuthConfig](
	//		ctx,
	//		cfg.Client,
	//		gatewaykubev1.SchemeGroupVersion.WithResource("authconfigs"),
	//		krt.WithName("AuthConfig"))

	setupLog.Info("initializing k8sgateway extensions")
	secretClient := kclient.New[*corev1.Secret](cfg.Client)
	k8sSecretsRaw := krt.WrapClient(secretClient, krt.WithStop(ctx.Done()), krt.WithName("Secrets") /* no debug here - we don't want raw secrets printed*/)
	k8sSecrets := krt.NewCollection(k8sSecretsRaw, func(kctx krt.HandlerContext, i *corev1.Secret) *ir.Secret {
		res := ir.Secret{
			ObjectSource: ir.ObjectSource{
				Group:     "",
				Kind:      "Secret",
				Namespace: i.Namespace,
				Name:      i.Name,
			},
			Obj:  i,
			Data: i.Data,
		}
		return &res
	}, cfg.KrtOptions.ToOptions("secrets")...)
	secrets := map[schema.GroupKind]krt.Collection[ir.Secret]{
		{Group: "", Kind: "Secret"}: k8sSecrets,
	}
	commoncol := common.CommonCollections{
		Client:   cfg.Client,
		KrtOpts:  cfg.KrtOptions,
		Secrets:  krtcollections.NewSecretIndex(secrets),
		Pods:     cfg.AugmentedPods,
		Settings: cfg.Settings,
	}

	// Create the proxy syncer for the Gateway API resources
	setupLog.Info("initializing proxy syncer")
	proxySyncer := proxy_syncer.NewProxySyncer(
		ctx,
		cfg.InitialSettings,
		cfg.Settings,
		wellknown.GatewayControllerName,
		mgr,
		cfg.Client,
		cfg.AugmentedPods,
		cfg.UniqueClients,
		pluginFactoryWithBuiltin,
		commoncol,
		cfg.SetupOpts.Cache,
	)
	proxySyncer.Init(ctx, cfg.KrtOptions)
	if err := mgr.Add(proxySyncer); err != nil {
		setupLog.Error(err, "unable to add proxySyncer runnable")
		return nil, err
	}

	return &ControllerBuilder{
		proxySyncer: proxySyncer,
		cfg:         cfg,
		mgr:         mgr,
	}, nil
}

func pluginFactoryWithBuiltin(ctx context.Context, commoncol *common.CommonCollections) extensionsplug.Plugin {
	plugins := registry.Plugins(ctx, commoncol)
	plugins = append(plugins, krtcollections.NewBuiltinPlugin(ctx))
	return registry.MergePlugins(plugins...)
}

func (c *ControllerBuilder) Start(ctx context.Context) error {
	logger := contextutils.LoggerFrom(ctx)
	logger.Info("starting gateway controller")
	// GetXdsAddress waits for gloo-edge to populate the xds address of the server.
	// in the future this logic may move here and be duplicated.
	xdsHost, xdsPort := c.cfg.SetupOpts.GetXdsAddress(ctx)
	if xdsHost == "" {
		return ctx.Err()
	}

	logger.Infow("got xds address for deployer", uzap.String("xds_host", xdsHost), uzap.Int32("xds_port", xdsPort))

	integrationEnabled := c.cfg.InitialSettings.Spec.GetGloo().GetIstioOptions().GetEnableIntegration().GetValue()

	// copy over relevant aws options (if any) from Settings
	var awsInfo *deployer.AwsInfo
	awsOpts := c.cfg.InitialSettings.Spec.GetGloo().GetAwsOptions()
	if awsOpts != nil {
		credOpts := awsOpts.GetServiceAccountCredentials()
		if credOpts != nil {
			awsInfo = &deployer.AwsInfo{
				EnableServiceAccountCredentials: true,
				StsClusterName:                  credOpts.GetCluster(),
				StsUri:                          credOpts.GetUri(),
			}
		} else {
			awsInfo = &deployer.AwsInfo{
				EnableServiceAccountCredentials: false,
			}
		}
	}

	// Initialize the set of Gateway API CRDs we care about
	crds, err := getGatewayCRDs(c.cfg.RestConfig)
	if err != nil {
		return err
	}

	gwCfg := GatewayConfig{
		Mgr:            c.mgr,
		GWClasses:      sets.New(append(c.cfg.SetupOpts.ExtraGatewayClasses, wellknown.GatewayClassName)...),
		ControllerName: wellknown.GatewayControllerName,
		AutoProvision:  AutoProvision,
		ControlPlane: deployer.ControlPlaneInfo{
			XdsHost: xdsHost,
			XdsPort: xdsPort,
		},
		// TODO pass in the settings so that the deloyer can register to it for changes.
		IstioIntegrationEnabled: integrationEnabled,
		Aws:                     awsInfo,
		Kick:                    func(context.Context) {},
		CRDs:                    crds,
	}
	if err := NewBaseGatewayController(ctx, gwCfg); err != nil {
		setupLog.Error(err, "unable to create controller")
		return err
	}

	return c.mgr.Start(ctx)
}

func getGatewayCRDs(restConfig *rest.Config) (sets.Set[string], error) {
	crds := wellknown.GatewayStandardCRDs

	tcpRouteExists, err := glooschemes.CRDExists(restConfig, gwv1a2.GroupVersion.Group, gwv1a2.GroupVersion.Version, wellknown.TCPRouteKind)
	if err != nil {
		return nil, err
	}

	if tcpRouteExists {
		crds.Insert(wellknown.TCPRouteCRDName)
	}

	return crds, nil
}
