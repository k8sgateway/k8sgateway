package utils

import (
	"context"
	"fmt"
	"reflect"

	"github.com/solo-io/gloo/projects/gateway2/query"
	"github.com/solo-io/gloo/projects/gateway2/translator/plugins"
	skv2corev1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
	gwv1 "sigs.k8s.io/gateway-api/apis/v1"
)

// FindAppliedRouteFilters finds all instances of the supplied filterTypes for the Rule supplied in the RouteContext.
// Should only be used for plugins that support multiple filters as part of a single Rule
func FindAppliedRouteFilters(
	routeCtx *plugins.RouteContext,
	filterTypes ...gwv1.HTTPRouteFilterType,
) []gwv1.HTTPRouteFilter {
	if routeCtx.Rule == nil {
		return nil
	}
	var appliedFilters []gwv1.HTTPRouteFilter
	for _, filter := range routeCtx.Rule.Filters {
		for _, filterType := range filterTypes {
			if filter.Type == filterType {
				appliedFilters = append(appliedFilters, filter)
			}
		}
	}
	return appliedFilters
}

// FindAppliedRouteFilter finds the first instance of the filterType supplied in the Rule being processed.
// Returns nil if the Rule doesn't contain a filter of the provided Type
func FindAppliedRouteFilter(
	routeCtx *plugins.RouteContext,
	filterType gwv1.HTTPRouteFilterType,
) *gwv1.HTTPRouteFilter {
	if routeCtx.Rule == nil {
		return nil
	}
	// TODO: check full Filter list for duplicates and error?
	for _, filter := range routeCtx.Rule.Filters {
		if filter.Type == filterType {
			return &filter
		}
	}
	return nil
}

// FindExtensionRefFilter finds the first instance of an ExtensionRef filter that
// references the supplied GroupKind in the Rule being processed.
// Returns nil if the Rule doesn't contain a matching ExtensionRef filter
func FindExtensionRefFilter(
	rule *gwv1.HTTPRouteRule,
	gk schema.GroupKind,
) *gwv1.HTTPRouteFilter {
	if rule == nil {
		return nil
	}
	// TODO: check full Filter list for duplicates and error?
	for _, filter := range rule.Filters {
		if filter.Type == gwv1.HTTPRouteFilterExtensionRef {
			if filter.ExtensionRef.Group == gwv1.Group(gk.Group) && filter.ExtensionRef.Kind == gwv1.Kind(gk.Kind) {
				return &filter
			}
		}
	}
	return nil
}

var (
	ErrTypesNotEqual = fmt.Errorf("types not equal")
	ErrNotSettable   = fmt.Errorf("can't set value")
)

// GetExtensionRefObj uses the provided query engine to retrieve an ExtensionRef object
// and set the value of `obj` to point to it.
// The type of `obj` must match the type referenced in the extensionRef and must be a pointer.
// An error will be returned if the Get was unsuccessful or if the type passed is not valid.
// A nil error indicates success and `obj` should be usable as normal.
func GetExtensionRefObj(
	ctx context.Context,
	route *gwv1.HTTPRoute,
	queries query.GatewayQueries,
	extensionRef *gwv1.LocalObjectReference,
	obj client.Object,
) error {
	return GetExtensionRefObjFrom(ctx, queries.ObjToFrom(route), queries, extensionRef, obj)
}

func GetExtensionRefObjFrom(
	ctx context.Context,
	from query.From,
	queries query.GatewayQueries,
	extensionRef *gwv1.LocalObjectReference,
	obj client.Object,
) error {
	localObj, err := queries.GetLocalObjRef(ctx, from, *extensionRef)
	if err != nil {
		return err
	}
	if reflect.TypeOf(obj) != reflect.TypeOf(localObj) {
		return fmt.Errorf(
			"%w: passed Obj typeOf: '%v' localObj typeOf: '%v'",
			ErrTypesNotEqual,
			reflect.TypeOf(obj),
			reflect.TypeOf(localObj),
		)
	}
	elem := reflect.ValueOf(obj).Elem()
	if !elem.CanSet() {
		return ErrNotSettable
	}
	elem.Set(reflect.ValueOf(localObj).Elem())
	return nil
}

// PolicyWithSectionedTargetRefs is a wrapper type to represent policy objects
// that attach via TargetRefWtihSectionName
type PolicyWithSectionedTargetRefs[T client.Object] interface {
	GetTargetRefs() []*skv2corev1.PolicyTargetReferenceWithSectionName
	GetObject() T
}

// GetPrioritizedListenerPolicies accepts a slice of Gateway-attached policies (that may explicitly
// target a specific Listener and returns a slice of these policies (or a subset) resources.
// The returned policy list is sorted by specificity in the order of
//
// 1. older with section name
//
// 2. newer with section name
//
// 3. older without section name
//
// 4. newer without section name
func GetPrioritizedListenerPolicies[T client.Object](
	items []PolicyWithSectionedTargetRefs[T],
	listener *gwv1.Listener,
) []T {
	var optsWithSectionName, optsWithoutSectionName []T
	for i := range items {
		item := items[i]
		// only use the first targetRef in the list for now; user should be warned by caller of this function
		targetRef := item.GetTargetRefs()[0]
		if sectionName := targetRef.GetSectionName(); sectionName != nil && sectionName.GetValue() != "" {
			// we have a section name, now check if it matches the specific listener provided
			if sectionName.GetValue() == string(listener.Name) {
				optsWithSectionName = append(optsWithSectionName, item.GetObject())
			}
		} else {
			// attach all matched items that do not have a section name and let the caller be discerning
			optsWithoutSectionName = append(optsWithoutSectionName, item.GetObject())
		}
	}

	// this can happen if the policy list only contains items targeting other Listeners by section name
	if len(optsWithoutSectionName)+len(optsWithSectionName) == 0 {
		return nil
	}

	SortByCreationTime(optsWithSectionName)
	SortByCreationTime(optsWithoutSectionName)
	return append(optsWithSectionName, optsWithoutSectionName...)
}

// TODO: remove this as part of https://github.com/solo-io/solo-projects/issues/6286
const MultipleTargetRefErrStr = "found ListenerOption %s/%s that contains multiple targetRefs which is not currently supported, only the first targetRef will be used"
