package reports

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	gwv1 "sigs.k8s.io/gateway-api/apis/v1"
)

type ReportMap struct {
	Gateways map[types.NamespacedName]*GatewayReport
	Routes   map[types.NamespacedName]*RouteReport
}

func NewReportMap() ReportMap {
	gr := make(map[types.NamespacedName]*GatewayReport)
	rr := make(map[types.NamespacedName]*RouteReport)
	return ReportMap{
		Gateways: gr,
		Routes:   rr,
	}
}

type GatewayReport struct {
	Conditions []metav1.Condition
	Listeners  map[string]*ListenerReport
}

type RouteReport struct {
	Parents map[ParentRefKey]*ParentRefReport
}

type ParentRefReport struct {
	Conditions []metav1.Condition
}

type ParentRefKey struct {
	Group string
	Kind  string
	types.NamespacedName
}

type ListenerReport struct {
	Status gwv1.ListenerStatus
}

type reporter struct {
	report *ReportMap
}

func (r *reporter) Gateway(gateway *gwv1.Gateway) GatewayReporter {
	var gr *GatewayReport
	key := client.ObjectKeyFromObject(gateway)
	gr, ok := r.report.Gateways[key]
	if !ok {
		gr = &GatewayReport{}
		r.report.Gateways[key] = gr
	}
	return gr
}

func (r *reporter) Route(route *gwv1.HTTPRoute) HTTPRouteReporter {
	var rr *RouteReport
	//TODO use client.ObjectKeyFromObject
	key := types.NamespacedName{
		Namespace: route.Namespace,
		Name:      route.Name,
	}
	rr, ok := r.report.Routes[key]
	if !ok {
		rr = &RouteReport{}
		r.report.Routes[key] = rr
	}
	return rr
}

func GetParentRefKey(parentRef *gwv1.ParentReference) ParentRefKey {
	var kind string
	if parentRef.Kind != nil {
		kind = string(*parentRef.Kind)
	}
	var ns string
	if parentRef.Namespace != nil {
		kind = string(*parentRef.Namespace)
	}
	return ParentRefKey{
		Group: string(parentRef.Name),
		Kind:  kind,
		NamespacedName: types.NamespacedName{
			Namespace: ns,
			Name:      string(parentRef.Name),
		},
	}

}

func (r *RouteReport) ParentRef(parentRef *gwv1.ParentReference) ParentRefReporter {
	key := GetParentRefKey(parentRef)
	if r.Parents == nil {
		r.Parents = make(map[ParentRefKey]*ParentRefReport)
	}
	var prr *ParentRefReport
	prr, ok := r.Parents[key]
	if !ok {
		prr = &ParentRefReport{}
		r.Parents[key] = prr
	}
	return prr
}

func (prr *ParentRefReport) SetCondition(rc HTTPRouteCondition) {
	condition := metav1.Condition{
		Type:    string(rc.Type),
		Status:  rc.Status,
		Reason:  string(rc.Reason),
		Message: rc.Message,
	}
	prr.Conditions = append(prr.Conditions, condition)
}

func (r *GatewayReport) Listener(listener *gwv1.Listener) ListenerReporter {
	if r.Listeners == nil {
		r.Listeners = make(map[string]*ListenerReport)
	}
	var lr *ListenerReport
	lr, ok := r.Listeners[string(listener.Name)]
	if !ok {
		lr = &ListenerReport{}
		lr.Status.Name = listener.Name
		r.Listeners[string(listener.Name)] = lr
	}
	return lr
}

func (g *GatewayReport) SetCondition(gc GatewayCondition) {
	condition := metav1.Condition{
		Type:    string(gc.Type),
		Status:  gc.Status,
		Reason:  string(gc.Reason),
		Message: gc.Message,
	}
	g.Conditions = append(g.Conditions, condition)
}

func (l *ListenerReport) SetCondition(lc ListenerCondition) {
	condition := metav1.Condition{
		Type:    string(lc.Type),
		Status:  lc.Status,
		Reason:  string(lc.Reason),
		Message: lc.Message,
	}
	l.Status.Conditions = append(l.Status.Conditions, condition)
}

func (l *ListenerReport) SetSupportedKinds(rgks []gwv1.RouteGroupKind) {
	l.Status.SupportedKinds = rgks
}

func (l *ListenerReport) SetAttachedRoutes(n uint) {
	l.Status.AttachedRoutes = int32(n)
}

func NewReporter(reportMap *ReportMap) Reporter {
	return &reporter{report: reportMap}
}

// Reports errors for GW translation
type Reporter interface {
	// returns the object reporter for the given type
	// TODO(Law): use string here instead of Gateway type
	Gateway(gateway *gwv1.Gateway) GatewayReporter

	Route(route *gwv1.HTTPRoute) HTTPRouteReporter
}

type GatewayReporter interface {
	// report an error on the whole gateway
	// Err(format string, a ...any)

	// report an error on the given listener
	// TODO(Law): use string here instead of Listener type
	Listener(listener *gwv1.Listener) ListenerReporter

	SetCondition(condition GatewayCondition)
}

type GatewayCondition struct {
	Type    gwv1.GatewayConditionType
	Status  metav1.ConditionStatus
	Reason  gwv1.GatewayConditionReason
	Message string
}

type ListenerCondition struct {
	Type    gwv1.ListenerConditionType
	Status  metav1.ConditionStatus
	Reason  gwv1.ListenerConditionReason
	Message string
}

type ListenerReporter interface {
	// report an error on the listener
	// Err(format string, a ...any)

	SetCondition(ListenerCondition)

	SetSupportedKinds([]gwv1.RouteGroupKind)
	SetAttachedRoutes(n uint)
}

type HTTPRouteCondition struct {
	Type    gwv1.RouteConditionType
	Status  metav1.ConditionStatus
	Reason  gwv1.RouteConditionReason
	Message string
}

type HTTPRouteReporter interface {
	// Err(format string, a ...any)
	ParentRef(parentRef *gwv1.ParentReference) ParentRefReporter
}

type ParentRefReporter interface {
	SetCondition(condition HTTPRouteCondition)
}
