package gateway

import (
	"context"

	"github.com/solo-io/gloo/pkg/utils/statsutils"
	extensionsplug "github.com/solo-io/gloo/projects/gateway2/extensions2/plugin"
	"github.com/solo-io/gloo/projects/gateway2/ir"
	"github.com/solo-io/go-utils/contextutils"
	"istio.io/istio/pkg/kube/krt"

	"github.com/solo-io/gloo/projects/gateway2/query"
	"github.com/solo-io/gloo/projects/gateway2/reports"
	"github.com/solo-io/gloo/projects/gateway2/translator/listener"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	gwv1 "sigs.k8s.io/gateway-api/apis/v1"
)

func NewTranslator(queries query.GatewayQueries) extensionsplug.KGwTranslator {
	return &translator{
		queries: queries,
	}
}

type translator struct {
	queries query.GatewayQueries
}

func (t *translator) Translate(
	kctx krt.HandlerContext,
	ctx context.Context,
	gateway *ir.Gateway,
	reporter reports.Reporter,
) *ir.GatewayIR {
	stopwatch := statsutils.NewTranslatorStopWatch("TranslateProxy")
	stopwatch.Start()
	defer stopwatch.Stop(ctx)

	ctx = contextutils.WithLogger(ctx, "k8s-gateway-translator")
	logger := contextutils.LoggerFrom(ctx)
	routesForGw, err := t.queries.GetRoutesForGateway(kctx, ctx, gateway.Obj)
	if err != nil {
		logger.Errorf("failed to get routes for gateway %.%ss: %v", gateway.Namespace, gateway.Name, err)
		// TODO: decide how/if to report this error on Gateway
		// reporter.Gateway(gateway).Err(err.Error())
		return nil
	}

	for _, rErr := range routesForGw.RouteErrors {
		reporter.Route(rErr.Route.GetSourceObject()).ParentRef(&rErr.ParentRef).SetCondition(reports.RouteCondition{
			Type:   gwv1.RouteConditionAccepted,
			Status: metav1.ConditionFalse,
			Reason: rErr.Error.Reason,
			// TODO message
		})
	}

	for _, listener := range gateway.Listeners {
		availRoutes := 0
		if res, ok := routesForGw.ListenerResults[string(listener.Name)]; ok {
			// TODO we've never checked if the ListenerResult has an error.. is it already on RouteErrors?
			availRoutes = len(res.Routes)
		}
		reporter.Gateway(gateway.Obj).ListenerName(string(listener.Name)).SetAttachedRoutes(uint(availRoutes))
	}

	listeners := listener.TranslateListeners(
		kctx,
		ctx,
		t.queries,
		gateway,
		routesForGw,
		reporter,
	)

	//	func() { panic("TODO: handle gw policy attachment") }()
	return &ir.GatewayIR{
		SourceObject:         gateway.Obj,
		Listeners:            listeners,
		AttachedPolicies:     gateway.AttachedListenerPolicies,
		AttachedHttpPolicies: gateway.AttachedHttpPolicies,
	}
}
