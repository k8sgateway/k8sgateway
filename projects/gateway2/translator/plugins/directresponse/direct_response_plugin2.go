package directresponse

import (
	"context"
	"fmt"
	"net/http"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	gwv1 "sigs.k8s.io/gateway-api/apis/v1"

	corev3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	envoy_config_listener_v3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	"github.com/solo-io/gloo/projects/controller/pkg/plugins"
	"github.com/solo-io/gloo/projects/gateway2/api/v1alpha1"
	extensionplug "github.com/solo-io/gloo/projects/gateway2/extensions2/plugin"
	"github.com/solo-io/gloo/projects/gateway2/model"
	"github.com/solo-io/gloo/projects/gateway2/reports"
	"github.com/solo-io/go-utils/contextutils"
	"istio.io/istio/pkg/kube"
	"istio.io/istio/pkg/kube/kclient"
	"istio.io/istio/pkg/kube/krt"
	"istio.io/istio/pkg/kube/kubetypes"
)

type directResponse struct {
	spec v1alpha1.DirectResponseSpec
}
type directResponseGwPass struct {
}

func NewPlugin2(ctx context.Context, istioClient kube.Client, dbg *krt.DebugHandler) *extensionplug.Plugin {

	col := SetupCollectionDynamic[v1alpha1.DirectResponse](
		ctx,
		istioClient,
		v1alpha1.GroupVersion.WithResource("directresponses"),
		krt.WithName("DirectResponse"), krt.WithDebugging(dbg),
	)
	gk := v1alpha1.DirectResponseGVK.GroupKind()
	policyCol := krt.NewCollection(col, func(krtctx krt.HandlerContext, i *v1alpha1.DirectResponse) *model.PolicyWrapper {
		var pol = &model.PolicyWrapper{
			ObjectSource: model.ObjectSource{
				Group:     gk.Group,
				Kind:      gk.Kind,
				Namespace: i.Namespace,
				Name:      i.Name,
			},
			Policy:   i,
			PolicyIR: &directResponse{spec: i.Spec},
			// no target refs for direct response
		}
		return pol
	})

	return &extensionplug.Plugin{
		ContributesPolicies: map[schema.GroupKind]model.PolicyImpl{
			v1alpha1.DirectResponseGVK.GroupKind(): {
				Name:             "directresponse",
				AttachmentPoints: []model.AttachmentPoints{model.HttpAttachmentPoint},
				//	NewGatewayTranslationPass: newPlug,
				Policies: policyCol,
				//				AttachmentPoints:          []model.AttachmentPoints{model.HttpAttachmentPoint},
				NewGatewayTranslationPass: NewGatewayTranslationPass,
			},
		},
	}
}

func NewGatewayTranslationPass(ctx context.Context, tctx model.GwTranslationCtx) model.ProxyTranslationPass {
	return &directResponseGwPass{}
}

// called 1 time for each listener
func (p *directResponseGwPass) ApplyListenerPlugin(ctx context.Context, pCtx *model.ListenerContext, out *envoy_config_listener_v3.Listener) {
}

func (p *directResponseGwPass) ApplyVhostPlugin(ctx context.Context, pCtx *model.VirtualHostContext, out *envoy_config_route_v3.VirtualHost) {
}

// called 0 or more times
func (p *directResponseGwPass) ApplyForRoute(ctx context.Context, pCtx *model.RouteContext, outputRoute *envoy_config_route_v3.Route) error {
	dr, ok := pCtx.Policy.(*directResponse)
	if !ok {
		return fmt.Errorf("internal error: expected *directResponse, got %T", pCtx.Policy)
	}
	// at this point, we have a valid DR reference that we should apply to the route.
	if outputRoute.GetAction() != nil {
		// the output route already has an action, which is incompatible with the DirectResponse,
		// so we'll return an error. note: the direct response plugin runs after other route plugins
		// that modify the output route (e.g. the redirect plugin), so this should be a rare case.
		errMsg := fmt.Sprintf("DirectResponse cannot be applied to route with existing action: %T", outputRoute.GetAction())
		pCtx.Reporter.SetCondition(reports.RouteCondition{
			Type:    gwv1.RouteConditionAccepted,
			Status:  metav1.ConditionFalse,
			Reason:  gwv1.RouteReasonIncompatibleFilters,
			Message: errMsg,
		})
		outputRoute.Action = &envoy_config_route_v3.Route_DirectResponse{
			DirectResponse: &envoy_config_route_v3.DirectResponseAction{
				Status: http.StatusInternalServerError,
			},
		}
		return fmt.Errorf(errMsg)
	}

	outputRoute.Action = &envoy_config_route_v3.Route_DirectResponse{
		DirectResponse: &envoy_config_route_v3.DirectResponseAction{
			Status: dr.spec.StatusCode,
			Body: &corev3.DataSource{
				Specifier: &corev3.DataSource_InlineString{
					InlineString: dr.spec.Body,
				},
			},
		},
	}
	return nil
}

func (p *directResponseGwPass) ApplyForRouteBackend(
	ctx context.Context,
	policy model.PolicyIR,
	pCtx *model.RouteBackendContext,
) error {
	return model.ErrNotAttachable
}

// called 1 time per listener
// if a plugin emits new filters, they must be with a plugin unique name.
// any filter returned from route config must be disabled, so it doesnt impact other routes.
func (p *directResponseGwPass) HttpFilters(ctx context.Context, fcc model.FilterChainCommon) ([]plugins.StagedHttpFilter, error) {
	return nil, nil
}

func (p *directResponseGwPass) UpstreamHttpFilters(ctx context.Context) ([]plugins.StagedUpstreamHttpFilter, error) {
	return nil, nil
}

func (p *directResponseGwPass) NetworkFilters(ctx context.Context) ([]plugins.StagedNetworkFilter, error) {
	return nil, nil
}

// called 1 time (per envoy proxy). replaces GeneratedResources
func (p *directResponseGwPass) ResourcesToAdd(ctx context.Context) model.Resources {
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
