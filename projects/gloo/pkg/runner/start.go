package runner

import (
	"context"
	gatewayv1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	ratelimitv1 "github.com/solo-io/gloo/projects/gloo/pkg/api/external/solo/ratelimit"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	extauthv1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/extauth/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/graphql/v1beta1"
	"github.com/solo-io/gloo/projects/gloo/pkg/debug"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	"github.com/solo-io/gloo/projects/gloo/pkg/syncer"
	"github.com/solo-io/gloo/projects/gloo/pkg/upstreams/consul"
	"github.com/solo-io/gloo/projects/gloo/pkg/validation"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	corecache "github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/cache"
	"github.com/solo-io/solo-kit/pkg/api/v1/control-plane/cache"
	"github.com/solo-io/solo-kit/pkg/api/v1/control-plane/server"
	skkube "github.com/solo-io/solo-kit/pkg/api/v1/resources/common/kubernetes"
	"github.com/solo-io/solo-kit/pkg/api/v2/reporter"
	"google.golang.org/grpc"
	"k8s.io/client-go/kubernetes"
	"net"
)

type StartFunc func(opts StartOpts) error

type StartOpts struct {
	WriteNamespace               string
	WatchNamespaces              []string

	Settings         *gloov1.Settings
	WatchOpts        clients.WatchOpts

	ResourceClientset ResourceClientset
	TypedClientset TypedClientset

	GatewayControllerEnabled     bool

	ControlPlane     ControlPlane
	ValidationServer ValidationServer
	ProxyDebugServer ProxyDebugServer
}

// A PluginRegistryFactory generates a PluginRegistry
// It is executed each translation loop, ensuring we have up to date configuration of all plugins
type PluginRegistryFactory func(ctx context.Context, opts StartOpts) plugins.PluginRegistry


type StartExtensions struct {
	PluginRegistryFactory PluginRegistryFactory
	SyncerExtensions      []syncer.TranslatorSyncerExtensionFactory
	XdsCallbacks          server.Callbacks
	ApiEmitterChannel     chan struct{}
}

type ResourceClientset struct {
	// Gateway resources
	VirtualServices       gatewayv1.VirtualServiceClient
	RouteTables           gatewayv1.RouteTableClient
	Gateways              gatewayv1.GatewayClient
	MatchableHttpGateways gatewayv1.MatchableHttpGatewayClient
	VirtualHostOptions    gatewayv1.VirtualHostOptionClient
	RouteOptions          gatewayv1.RouteOptionClient

	// Gloo resources
	Endpoints             gloov1.EndpointClient
	Upstreams             gloov1.UpstreamClient
	UpstreamGroups        gloov1.UpstreamGroupClient
	Proxies               gloov1.ProxyClient
	Secrets               gloov1.SecretClient
	Artifacts             gloov1.ArtifactClient

	// Gloo Enterprise resources
	AuthConfigs           extauthv1.AuthConfigClient
	GraphQLApis           v1beta1.GraphQLApiClient
	RateLimitConfigs      ratelimitv1.RateLimitConfigClient
	RateLimitReporter reporter.ReporterResourceClient
}

type TypedClientset struct {
	// Kubernetes clients
	KubeClient                   kubernetes.Interface
	KubeServiceClient            skkube.ServiceClient
	KubeCoreCache                corecache.KubeCoreCache

	// Consul clients
	ConsulWatcher      consul.ConsulWatcher
}

type ControlPlane struct {
	*GrpcService
	SnapshotCache cache.SnapshotCache
	XDSServer     server.Server
}

// ValidationServer validates proxies generated by controllers outside the gloo pod
type ValidationServer struct {
	*GrpcService
	Server validation.ValidationServer
}

// ProxyDebugServer returns proxies to callers outside the gloo pod - this is only necessary for UI/debugging purposes.
type ProxyDebugServer struct {
	*GrpcService
	Server debug.ProxyEndpointServer
}
type GrpcService struct {
	Ctx             context.Context
	BindAddr        net.Addr
	GrpcServer      *grpc.Server
	StartGrpcServer bool
}
