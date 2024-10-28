package krtcollection

import (
	"context"

	"github.com/solo-io/go-utils/contextutils"
	"istio.io/istio/pkg/kube"
	"istio.io/istio/pkg/kube/kclient"
	"istio.io/istio/pkg/kube/krt"
	"istio.io/istio/pkg/kube/kubetypes"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

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
	delayedClient := kclient.NewDelayedInformer[*unstructured.Unstructured](client, gvr, kubetypes.DynamicInformer, kclient.Filter{})
	mapper := krt.WrapClient(delayedClient, opts...)
	return krt.NewCollection(mapper, func(krtctx krt.HandlerContext, i *unstructured.Unstructured) **T {
		var empty T
		out := &empty
		err := runtime.DefaultUnstructuredConverter.FromUnstructured(i.UnstructuredContent(), out)
		if err != nil {
			contextutils.LoggerFrom(ctx).DPanic("failed converting unstructured into %T: %v", empty, i)
			return nil
		}
		return &out
	})
}
