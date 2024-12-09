package routeoptions

import (
	"context"
	"time"

	"google.golang.org/protobuf/types/known/durationpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	envoy_config_listener_v3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	"github.com/solo-io/gloo/projects/controller/pkg/plugins"
	"github.com/solo-io/gloo/projects/gateway2/api/v1alpha1"
	extensionplug "github.com/solo-io/gloo/projects/gateway2/extensions2/plugin"
	"github.com/solo-io/gloo/projects/gateway2/model"
	"github.com/solo-io/go-utils/contextutils"
	"istio.io/istio/pkg/kube"
	"istio.io/istio/pkg/kube/kclient"
	"istio.io/istio/pkg/kube/krt"
	"istio.io/istio/pkg/kube/kubetypes"
	gw1alpha2 "sigs.k8s.io/gateway-api/apis/v1alpha2"
)

type routeOptsPlugin struct {
	spec v1alpha1.RoutePolicySpec
}
type routeOptsPluginGwPass struct {
	dr *routeOptsPlugin
}

func NewPlugin2(ctx context.Context, istioClient kube.Client, dbg *krt.DebugHandler) *extensionplug.Plugin {

	col := SetupCollectionDynamic[v1alpha1.RoutePolicy](
		ctx,
		istioClient,
		v1alpha1.GroupVersion.WithResource("routepolicies"),
		krt.WithName("RoutePolicy"), krt.WithDebugging(dbg),
	)
	gk := v1alpha1.RoutePolicyGVK.GroupKind()
	policyCol := krt.NewCollection(col, func(krtctx krt.HandlerContext, i *v1alpha1.RoutePolicy) *model.PolicyWrapper {
		var pol = &model.PolicyWrapper{
			ObjectSource: model.ObjectSource{
				Group:     gk.Group,
				Kind:      gk.Kind,
				Namespace: i.Namespace,
				Name:      i.Name,
			},
			Policy:     i,
			PolicyIR:   &routeOptsPlugin{spec: i.Spec},
			TargetRefs: convert(i.Spec.TargetRef),
		}
		return pol
	})

	return &extensionplug.Plugin{
		ContributesPolicies: map[schema.GroupKind]model.PolicyImpl{
			v1alpha1.RoutePolicyGVK.GroupKind(): {
				AttachmentPoints: []model.AttachmentPoints{model.HttpAttachmentPoint},
				//NewGatewayTranslationPass: newPlug,
				Policies: policyCol,
			},
		},
	}
}

func convert(targetRef gw1alpha2.LocalPolicyTargetReference) []model.PolicyTargetRef {
	return []model.PolicyTargetRef{{
		Kind:  string(targetRef.Kind),
		Name:  string(targetRef.Name),
		Group: string(targetRef.Group),
	}}
}

func (d *routeOptsPlugin) NewGatewayTranslationPass(ctx context.Context, tctx model.GwTranslationCtx) model.ProxyTranslationPass {
	return &routeOptsPluginGwPass{dr: d}
}
func (p *routeOptsPlugin) Name() string {
	return "routepolicies"
}

// called 1 time for each listener
func (p *routeOptsPluginGwPass) ApplyListenerPlugin(ctx context.Context, pCtx *model.ListenerContext, out *envoy_config_listener_v3.Listener) {
}

func (p *routeOptsPluginGwPass) ApplyVhostPlugin(ctx context.Context, pCtx *model.VirtualHostContext, out *envoy_config_route_v3.VirtualHost) {
}

// called 0 or more times
func (p *routeOptsPluginGwPass) ApplyForRoute(ctx context.Context, pCtx *model.RouteContext, outputRoute *envoy_config_route_v3.Route) error {

	if p.dr.spec.Timeout > 0 && outputRoute.GetRoute() != nil {
		outputRoute.GetRoute().Timeout = durationpb.New(time.Second * time.Duration(p.dr.spec.Timeout))
	}

	return nil
}

func (p *routeOptsPluginGwPass) ApplyForRouteBackend(
	ctx context.Context,
	policy model.PolicyIR,
	pCtx *model.RouteBackendContext,
) error {
	return nil
}

// called 1 time per listener
// if a plugin emits new filters, they must be with a plugin unique name.
// any filter returned from route config must be disabled, so it doesnt impact other routes.
func (p *routeOptsPluginGwPass) HttpFilters(ctx context.Context, fcc model.FilterChainCommon) ([]plugins.StagedHttpFilter, error) {
	return nil, nil
}

func (p *routeOptsPluginGwPass) UpstreamHttpFilters(ctx context.Context) ([]plugins.StagedUpstreamHttpFilter, error) {
	return nil, nil
}

func (p *routeOptsPluginGwPass) NetworkFilters(ctx context.Context) ([]plugins.StagedNetworkFilter, error) {
	return nil, nil
}

// called 1 time (per envoy proxy). replaces GeneratedResources
func (p *routeOptsPluginGwPass) ResourcesToAdd(ctx context.Context) model.Resources {
	return model.Resources{}
}

// SetupCollectionDynamic uses the dynamic client to setup an informer for a resource
// and then uses an intermediate krt collection to type the unstructured resource.
// This is a temporary workaround until we update to the latest istio version and can
// uncomment the code below for registering types.
// HACK: we don't want to use this long term, but it's letting me push forward with deveopment
func SetupCollectionDynamic[T any](
	ctx context.Context,
	client kube.Client,
	gvr schema.GroupVersionResource,
	opts ...krt.CollectionOption,
) krt.Collection[*T] {
	logger := contextutils.LoggerFrom(ctx)
	logger.Infof("setting up dynamic collection for %s", gvr.String())
	delayedClient := kclient.NewDelayedInformer[*unstructured.Unstructured](client, gvr, kubetypes.DynamicInformer, kclient.Filter{})
	mapper := krt.WrapClient(delayedClient, opts...)
	return krt.NewCollection(mapper, func(krtctx krt.HandlerContext, i *unstructured.Unstructured) **T {
		var empty T
		out := &empty
		err := runtime.DefaultUnstructuredConverter.FromUnstructured(i.UnstructuredContent(), out)
		if err != nil {
			logger.DPanic("failed converting unstructured into %T: %v", empty, i)
			return nil
		}
		return &out
	})
}
