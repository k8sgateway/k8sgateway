package services

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/solo-io/gloo/projects/gloo/cli/pkg/helpers"

	"github.com/imdario/mergo"

	"github.com/golang/protobuf/ptypes/wrappers"

	"github.com/solo-io/gloo/pkg/bootstrap/leaderelector/singlereplica"

	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/gloosnapshot"

	"github.com/solo-io/gloo/pkg/utils/settingsutil"

	"github.com/solo-io/gloo/pkg/utils/statusutils"

	"github.com/solo-io/gloo/projects/gloo/pkg/syncer/setup"

	"github.com/solo-io/gloo/projects/gateway/pkg/translator"

	"github.com/solo-io/gloo/projects/gloo/pkg/upstreams/consul"

	"github.com/solo-io/solo-kit/pkg/api/external/kubernetes/service"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/cache"

	skkube "github.com/solo-io/solo-kit/pkg/api/v1/resources/common/kubernetes"

	corecache "github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/cache"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/memory"

	gatewayv1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/bootstrap"
	"google.golang.org/grpc"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
	. "github.com/onsi/gomega"
	fds_syncer "github.com/solo-io/gloo/projects/discovery/pkg/fds/syncer"
	uds_syncer "github.com/solo-io/gloo/projects/discovery/pkg/uds/syncer"
	"github.com/solo-io/gloo/projects/gloo/pkg/defaults"
	"k8s.io/client-go/kubernetes"
)

type TestClients struct {
	GatewayClient        gatewayv1.GatewayClient
	HttpGatewayClient    gatewayv1.MatchableHttpGatewayClient
	VirtualServiceClient gatewayv1.VirtualServiceClient
	ProxyClient          gloov1.ProxyClient
	UpstreamClient       gloov1.UpstreamClient
	SecretClient         gloov1.SecretClient
	ServiceClient        skkube.ServiceClient
	GlooPort             int
	RestXdsPort          int
}

// WriteSnapshot writes all resources in the ApiSnapshot to the cache
func (c TestClients) WriteSnapshot(ctx context.Context, snapshot *gloosnapshot.ApiSnapshot) error {
	// We intentionally create child resources first to avoid having the validation webhook reject
	// the parent resource

	writeOptions := clients.WriteOpts{
		Ctx:               ctx,
		OverwriteExisting: false,
	}
	for _, secret := range snapshot.Secrets {
		if _, writeErr := c.SecretClient.Write(secret, writeOptions); writeErr != nil {
			return writeErr
		}
	}
	for _, us := range snapshot.Upstreams {
		if _, writeErr := c.UpstreamClient.Write(us, writeOptions); writeErr != nil {
			return writeErr
		}
	}
	for _, vs := range snapshot.VirtualServices {
		if _, writeErr := c.VirtualServiceClient.Write(vs, writeOptions); writeErr != nil {
			return writeErr
		}
	}
	for _, hgw := range snapshot.HttpGateways {
		if _, writeErr := c.HttpGatewayClient.Write(hgw, writeOptions); writeErr != nil {
			return writeErr
		}
	}
	for _, gw := range snapshot.Gateways {
		if _, writeErr := c.GatewayClient.Write(gw, writeOptions); writeErr != nil {
			return writeErr
		}
	}
	for _, proxy := range snapshot.Proxies {
		if _, writeErr := c.ProxyClient.Write(proxy, writeOptions); writeErr != nil {
			return writeErr
		}
	}

	return nil
}

// DeleteSnapshot deletes all resources in the ApiSnapshot from the cache
func (c TestClients) DeleteSnapshot(ctx context.Context, snapshot *gloosnapshot.ApiSnapshot) error {
	// We intentionally delete resources in the reverse order that we create resources
	// If we delete child resources first, the validation webhook may reject the change

	deleteOptions := clients.DeleteOpts{
		Ctx:            ctx,
		IgnoreNotExist: true,
	}

	for _, gw := range snapshot.Gateways {
		gwNamespace, gwName := gw.GetMetadata().Ref().Strings()
		if deleteErr := c.GatewayClient.Delete(gwNamespace, gwName, deleteOptions); deleteErr != nil {
			return deleteErr
		}
	}
	for _, hgw := range snapshot.HttpGateways {
		hgwNamespace, hgwName := hgw.GetMetadata().Ref().Strings()
		if deleteErr := c.HttpGatewayClient.Delete(hgwNamespace, hgwName, deleteOptions); deleteErr != nil {
			return deleteErr
		}
	}
	for _, vs := range snapshot.VirtualServices {
		vsNamespace, vsName := vs.GetMetadata().Ref().Strings()
		if deleteErr := c.VirtualServiceClient.Delete(vsNamespace, vsName, deleteOptions); deleteErr != nil {
			return deleteErr
		}
	}
	for _, us := range snapshot.Upstreams {
		usNamespace, usName := us.GetMetadata().Ref().Strings()
		if deleteErr := c.UpstreamClient.Delete(usNamespace, usName, deleteOptions); deleteErr != nil {
			return deleteErr
		}
	}
	for _, secret := range snapshot.Secrets {
		secretNamespace, secretName := secret.GetMetadata().Ref().Strings()
		if deleteErr := c.SecretClient.Delete(secretNamespace, secretName, deleteOptions); deleteErr != nil {
			return deleteErr
		}
	}

	// Proxies are auto generated by Gateway resources
	// Therefore we delete Proxies after we have deleted the resources that may regenerate a Proxy
	for _, proxy := range snapshot.Proxies {
		proxyNamespace, proxyName := proxy.GetMetadata().Ref().Strings()
		if deleteErr := c.ProxyClient.Delete(proxyNamespace, proxyName, deleteOptions); deleteErr != nil {
			return deleteErr
		}
	}

	return nil
}

var glooPortBase = int32(30400)

func AllocateGlooPort() int32 {
	return atomic.AddInt32(&glooPortBase, 1) + int32(config.GinkgoConfig.ParallelNode*1000)
}

type What struct {
	DisableGateway bool
	DisableUds     bool
	DisableFds     bool
}

type RunOptions struct {
	NsToWrite        string
	NsToWatch        []string
	WhatToRun        What
	GlooPort         int32
	ValidationPort   int32
	RestXdsPort      int32
	Settings         *gloov1.Settings
	KubeClient       kubernetes.Interface
	ConsulClient     consul.ConsulWatcher
	ConsulDnsAddress string
}

func RunGateway(ctx context.Context, justGloo bool) TestClients {
	ns := defaults.GlooSystem
	ro := &RunOptions{
		NsToWrite: ns,
		NsToWatch: []string{"default", ns},
		WhatToRun: What{
			DisableGateway: justGloo,
		},
		KubeClient: helpers.MustKubeClient(),
	}
	return RunGlooGatewayUdsFds(ctx, ro)
}

//noinspection GoUnhandledErrorResult
func RunGlooGatewayUdsFds(ctx context.Context, runOptions *RunOptions) TestClients {
	if runOptions.GlooPort == 0 {
		runOptions.GlooPort = AllocateGlooPort()
	}
	if runOptions.ValidationPort == 0 {
		runOptions.ValidationPort = AllocateGlooPort()
	}
	if runOptions.RestXdsPort == 0 {
		runOptions.RestXdsPort = AllocateGlooPort()
	}

	settings := constructTestSettings(runOptions)
	ctx = settingsutil.WithSettings(ctx, settings)

	// All Gloo Edge components run using a Bootstrap.Opts object
	// These values are extracted from the Settings object and as part of our SetupSyncer
	// we pull values off the Settings object to build the Bootstrap.Opts. It would be ideal if we
	// could use the same setup code, but in the meantime, we use constructTestOpts to mirror the functionality
	bootstrapOpts := constructTestOpts(ctx, runOptions, settings)

	go setup.RunGloo(bootstrapOpts)

	if !runOptions.WhatToRun.DisableFds {
		go func() {
			defer GinkgoRecover()
			fds_syncer.RunFDS(bootstrapOpts)
		}()
	}
	if !runOptions.WhatToRun.DisableUds {
		go func() {
			defer GinkgoRecover()
			uds_syncer.RunUDS(bootstrapOpts)
		}()
	}

	testClients := getTestClients(ctx, bootstrapOpts)
	testClients.GlooPort = int(runOptions.GlooPort)
	testClients.RestXdsPort = int(runOptions.RestXdsPort)
	return testClients
}

func getTestClients(ctx context.Context, bootstrapOpts bootstrap.Opts) TestClients {
	gatewayClient, err := gatewayv1.NewGatewayClient(ctx, bootstrapOpts.Gateways)
	Expect(err).NotTo(HaveOccurred())
	httpGatewayClient, err := gatewayv1.NewMatchableHttpGatewayClient(ctx, bootstrapOpts.MatchableHttpGateways)
	Expect(err).NotTo(HaveOccurred())
	virtualServiceClient, err := gatewayv1.NewVirtualServiceClient(ctx, bootstrapOpts.VirtualServices)
	Expect(err).NotTo(HaveOccurred())
	upstreamClient, err := gloov1.NewUpstreamClient(ctx, bootstrapOpts.Upstreams)
	Expect(err).NotTo(HaveOccurred())
	secretClient, err := gloov1.NewSecretClient(ctx, bootstrapOpts.Secrets)
	Expect(err).NotTo(HaveOccurred())
	proxyClient, err := gloov1.NewProxyClient(ctx, bootstrapOpts.Proxies)
	Expect(err).NotTo(HaveOccurred())

	return TestClients{
		GatewayClient:        gatewayClient,
		HttpGatewayClient:    httpGatewayClient,
		VirtualServiceClient: virtualServiceClient,
		UpstreamClient:       upstreamClient,
		SecretClient:         secretClient,
		ProxyClient:          proxyClient,
		ServiceClient:        bootstrapOpts.KubeServiceClient,
	}
}

func constructTestSettings(runOptions *RunOptions) *gloov1.Settings {
	// Define default Settings that all tests will inherit
	settings := &gloov1.Settings{
		WatchNamespaces:    runOptions.NsToWatch,
		DiscoveryNamespace: runOptions.NsToWrite,
		Gloo: &gloov1.GlooOptions{
			RemoveUnusedFilters: &wrappers.BoolValue{Value: true},
			RestXdsBindAddr:     fmt.Sprintf("%s:%d", net.IPv4zero.String(), runOptions.RestXdsPort),

			// Invalid Routes can be difficult to track down
			// By creating a Response Code and Body that are unique, hopefully it is easier to identify situations
			// Where invalid route replacement is taking effect
			InvalidConfigPolicy: &gloov1.GlooOptions_InvalidConfigPolicy{
				ReplaceInvalidRoutes:     true,
				InvalidRouteResponseCode: http.StatusTeapot,
				InvalidRouteResponseBody: "Invalid Route Replacement Encountered In Test",
			},
		},
		Gateway: &gloov1.GatewayOptions{
			Validation: &gloov1.GatewayOptions_ValidationOptions{
				// To validate transformations, we call out to an Envoy binary running in validate mode
				// https://github.com/solo-io/gloo/blob/01d04751f72c168e304977c4f67fdbcbf30232a9/projects/gloo/pkg/bootstrap/bootstrap_validation.go#L28
				// This binary is present in our CI/CD pipeline. But when running locally it is not, so we fallback to the Upstream Envoy binary
				// which doesn't have the custom Solo.io types registered with the deserializer. Therefore, when running locally tests will fail,
				// and the logs will contain:
				//	"Invalid type URL, unknown type: envoy.api.v2.filter.http.RouteTransformations for type Any)"
				DisableTransformationValidation: &wrappers.BoolValue{
					Value: true,
				},
			},
			EnableGatewayController: &wrappers.BoolValue{
				Value: !runOptions.WhatToRun.DisableGateway,
			},
			// To make debugging slightly easier
			PersistProxySpec: &wrappers.BoolValue{
				Value: true,
			},
			// For now we default this to false, and have explicit tests (aggregate_listener_test), which validate
			// the behavior when the setting is configured to true
			IsolateVirtualHostsBySslConfig: &wrappers.BoolValue{
				Value: false,
			},
		},
	}

	// Allow tests to override the default Settings
	if runOptions.Settings != nil {
		err := mergo.Merge(settings, runOptions.Settings)
		ExpectWithOffset(1, err).NotTo(HaveOccurred())
	}

	return settings
}

// constructTestOpts mirrors constructOpts in our SetupSyncer
func constructTestOpts(ctx context.Context, runOptions *RunOptions, settings *gloov1.Settings) bootstrap.Opts {
	ctx = contextutils.WithLogger(ctx, "gloo")
	logger := contextutils.LoggerFrom(ctx)
	grpcServer := grpc.NewServer(grpc.StreamInterceptor(
		grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(zap.NewNop()),
			func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
				logger.Infof("gRPC call: %v", info.FullMethod)
				return handler(srv, ss)
			},
		)),
	)
	grpcServerValidation := grpc.NewServer(grpc.StreamInterceptor(
		grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(zap.NewNop()),
			func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
				logger.Infof("gRPC call: %v", info.FullMethod)
				return handler(srv, ss)
			},
		)),
	)
	f := &factory.MemoryResourceClientFactory{
		Cache: memory.NewInMemoryResourceCache(),
	}
	var kubeCoreCache corecache.KubeCoreCache
	if runOptions.KubeClient != nil {
		var err error
		kubeCoreCache, err = cache.NewKubeCoreCacheWithOptions(ctx, runOptions.KubeClient, time.Hour, settings.GetWatchNamespaces())
		Expect(err).NotTo(HaveOccurred())
	}
	var validationOpts *translator.ValidationOpts
	if settings.GetGateway().GetValidation().GetProxyValidationServerAddr() != "" {
		if validationOpts == nil {
			validationOpts = &translator.ValidationOpts{}
		}
		validationOpts.ProxyValidationServerAddress = settings.GetGateway().GetValidation().GetProxyValidationServerAddr()
	}
	if settings.GetGateway().GetValidation().GetAllowWarnings() != nil {
		if validationOpts == nil {
			validationOpts = &translator.ValidationOpts{}
		}
		validationOpts.AllowWarnings = settings.GetGateway().GetValidation().GetAllowWarnings().GetValue()
	}
	if settings.GetGateway().GetValidation().GetAlwaysAccept() != nil {
		if validationOpts == nil {
			validationOpts = &translator.ValidationOpts{}
		}
		validationOpts.AlwaysAcceptResources = settings.GetGateway().GetValidation().GetAlwaysAccept().GetValue()
	}
	return bootstrap.Opts{
		Settings:                settings,
		WriteNamespace:          settings.GetDiscoveryNamespace(),
		StatusReporterNamespace: statusutils.GetStatusReporterNamespaceOrDefault(defaults.GlooSystem),
		Upstreams:               f,
		UpstreamGroups:          f,
		Proxies:                 f,
		Secrets:                 f,
		Artifacts:               f,
		AuthConfigs:             f,
		RateLimitConfigs:        f,
		GraphQLApis:             f,
		Gateways:                f,
		MatchableHttpGateways:   f,
		VirtualServices:         f,
		RouteTables:             f,
		RouteOptions:            f,
		VirtualHostOptions:      f,
		KubeServiceClient:       newServiceClient(ctx, f, runOptions),
		WatchNamespaces:         settings.GetWatchNamespaces(),
		WatchOpts: clients.WatchOpts{
			Ctx:         ctx,
			RefreshRate: time.Second / 10,
		},
		ControlPlane: setup.NewControlPlane(ctx, grpcServer, &net.TCPAddr{
			IP:   net.IPv4zero,
			Port: int(runOptions.GlooPort),
		}, nil, true),
		ValidationServer: setup.NewValidationServer(ctx, grpcServerValidation, &net.TCPAddr{
			IP:   net.IPv4zero,
			Port: int(runOptions.ValidationPort),
		}, true),
		ProxyDebugServer: setup.NewProxyDebugServer(ctx, grpcServer, &net.TCPAddr{
			IP:   net.IPv4zero,
			Port: 8001,
		}, false),
		KubeClient:    runOptions.KubeClient,
		KubeCoreCache: kubeCoreCache,
		DevMode:       true,
		Consul: bootstrap.Consul{
			ConsulWatcher: runOptions.ConsulClient,
			DnsServer:     runOptions.ConsulDnsAddress,
		},
		GatewayControllerEnabled: settings.GetGateway().GetEnableGatewayController().GetValue(),
		ValidationOpts:           validationOpts,
		Identity:                 singlereplica.Identity(),
	}
}

func newServiceClient(ctx context.Context, memFactory *factory.MemoryResourceClientFactory, runOpts *RunOptions) skkube.ServiceClient {

	// If the KubeClient option is set, the kubernetes discovery plugin will be activated and we must provide a
	// kubernetes service client in order for service-derived upstreams to be included in the snapshot
	if kube := runOpts.KubeClient; kube != nil {
		kubeCache, err := cache.NewKubeCoreCache(ctx, kube)
		if err != nil {
			panic(err)
		}
		return service.NewServiceClient(kube, kubeCache)
	}

	// Else return in-memory client
	client, err := skkube.NewServiceClient(ctx, memFactory)
	if err != nil {
		panic(err)
	}
	return client
}
