// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/reconcile"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
)

// Option to copy anything from the original to the desired before writing. Return value of false means don't update
type TransitionUpstreamGroupFunc func(original, desired *UpstreamGroup) (bool, error)

type UpstreamGroupReconciler interface {
	Reconcile(namespace string, desiredResources UpstreamGroupList, transition TransitionUpstreamGroupFunc, opts clients.ListOpts) error
}

func upstreamGroupsToResources(list UpstreamGroupList) resources.ResourceList {
	var resourceList resources.ResourceList
	for _, upstreamGroup := range list {
		resourceList = append(resourceList, upstreamGroup)
	}
	return resourceList
}

func NewUpstreamGroupReconciler(client UpstreamGroupClient) UpstreamGroupReconciler {
	return &upstreamGroupReconciler{
		base: reconcile.NewReconciler(client.BaseClient()),
	}
}

type upstreamGroupReconciler struct {
	base reconcile.Reconciler
}

func (r *upstreamGroupReconciler) Reconcile(namespace string, desiredResources UpstreamGroupList, transition TransitionUpstreamGroupFunc, opts clients.ListOpts) error {
	opts = opts.WithDefaults()
	opts.Ctx = contextutils.WithLogger(opts.Ctx, "upstreamGroup_reconciler")
	var transitionResources reconcile.TransitionResourcesFunc
	if transition != nil {
		transitionResources = func(original, desired resources.Resource) (bool, error) {
			return transition(original.(*UpstreamGroup), desired.(*UpstreamGroup))
		}
	}
	return r.base.Reconcile(namespace, upstreamGroupsToResources(desiredResources), transitionResources, opts)
}
