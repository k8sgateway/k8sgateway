// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"sync"
	"time"

	gloo_solo_io "github.com/solo-io/gloo/projects/controller/pkg/api/v1"

	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
	"go.uber.org/zap"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/errors"
	skstats "github.com/solo-io/solo-kit/pkg/stats"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/go-utils/errutils"
)

var (
	// Deprecated. See mTranslatorResourcesIn
	mTranslatorSnapshotIn = stats.Int64("translator.ingress.solo.io/emitter/snap_in", "Deprecated. Use translator.ingress.solo.io/emitter/resources_in. The number of snapshots in", "1")

	// metrics for emitter
	mTranslatorResourcesIn    = stats.Int64("translator.ingress.solo.io/emitter/resources_in", "The number of resource lists received on open watch channels", "1")
	mTranslatorSnapshotOut    = stats.Int64("translator.ingress.solo.io/emitter/snap_out", "The number of snapshots out", "1")
	mTranslatorSnapshotMissed = stats.Int64("translator.ingress.solo.io/emitter/snap_missed", "The number of snapshots missed", "1")

	// views for emitter
	// deprecated: see translatorResourcesInView
	translatorsnapshotInView = &view.View{
		Name:        "translator.ingress.solo.io/emitter/snap_in",
		Measure:     mTranslatorSnapshotIn,
		Description: "Deprecated. Use translator.ingress.solo.io/emitter/resources_in. The number of snapshots updates coming in.",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}

	translatorResourcesInView = &view.View{
		Name:        "translator.ingress.solo.io/emitter/resources_in",
		Measure:     mTranslatorResourcesIn,
		Description: "The number of resource lists received on open watch channels",
		Aggregation: view.Count(),
		TagKeys: []tag.Key{
			skstats.NamespaceKey,
			skstats.ResourceKey,
		},
	}
	translatorsnapshotOutView = &view.View{
		Name:        "translator.ingress.solo.io/emitter/snap_out",
		Measure:     mTranslatorSnapshotOut,
		Description: "The number of snapshots updates going out",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}
	translatorsnapshotMissedView = &view.View{
		Name:        "translator.ingress.solo.io/emitter/snap_missed",
		Measure:     mTranslatorSnapshotMissed,
		Description: "The number of snapshots updates going missed. this can happen in heavy load. missed snapshot will be re-tried after a second.",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}
)

func init() {
	view.Register(
		translatorsnapshotInView,
		translatorsnapshotOutView,
		translatorsnapshotMissedView,
		translatorResourcesInView,
	)
}

type TranslatorSnapshotEmitter interface {
	Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *TranslatorSnapshot, <-chan error, error)
}

type TranslatorEmitter interface {
	TranslatorSnapshotEmitter
	Register() error
	Upstream() gloo_solo_io.UpstreamClient
	KubeService() KubeServiceClient
	Ingress() IngressClient
}

func NewTranslatorEmitter(upstreamClient gloo_solo_io.UpstreamClient, kubeServiceClient KubeServiceClient, ingressClient IngressClient) TranslatorEmitter {
	return NewTranslatorEmitterWithEmit(upstreamClient, kubeServiceClient, ingressClient, make(chan struct{}))
}

func NewTranslatorEmitterWithEmit(upstreamClient gloo_solo_io.UpstreamClient, kubeServiceClient KubeServiceClient, ingressClient IngressClient, emit <-chan struct{}) TranslatorEmitter {
	return &translatorEmitter{
		upstream:    upstreamClient,
		kubeService: kubeServiceClient,
		ingress:     ingressClient,
		forceEmit:   emit,
	}
}

type translatorEmitter struct {
	forceEmit   <-chan struct{}
	upstream    gloo_solo_io.UpstreamClient
	kubeService KubeServiceClient
	ingress     IngressClient
}

func (c *translatorEmitter) Register() error {
	if err := c.upstream.Register(); err != nil {
		return err
	}
	if err := c.kubeService.Register(); err != nil {
		return err
	}
	if err := c.ingress.Register(); err != nil {
		return err
	}
	return nil
}

func (c *translatorEmitter) Upstream() gloo_solo_io.UpstreamClient {
	return c.upstream
}

func (c *translatorEmitter) KubeService() KubeServiceClient {
	return c.kubeService
}

func (c *translatorEmitter) Ingress() IngressClient {
	return c.ingress
}

func (c *translatorEmitter) Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *TranslatorSnapshot, <-chan error, error) {

	if len(watchNamespaces) == 0 {
		watchNamespaces = []string{""}
	}

	for _, ns := range watchNamespaces {
		if ns == "" && len(watchNamespaces) > 1 {
			return nil, nil, errors.Errorf("the \"\" namespace is used to watch all namespaces. Snapshots can either be tracked for " +
				"specific namespaces or \"\" AllNamespaces, but not both.")
		}
	}

	errs := make(chan error)
	var done sync.WaitGroup
	ctx := opts.Ctx
	/* Create channel for Upstream */
	type upstreamListWithNamespace struct {
		list      gloo_solo_io.UpstreamList
		namespace string
	}
	upstreamChan := make(chan upstreamListWithNamespace)

	var initialUpstreamList gloo_solo_io.UpstreamList
	/* Create channel for KubeService */
	type kubeServiceListWithNamespace struct {
		list      KubeServiceList
		namespace string
	}
	kubeServiceChan := make(chan kubeServiceListWithNamespace)

	var initialKubeServiceList KubeServiceList
	/* Create channel for Ingress */
	type ingressListWithNamespace struct {
		list      IngressList
		namespace string
	}
	ingressChan := make(chan ingressListWithNamespace)

	var initialIngressList IngressList

	currentSnapshot := TranslatorSnapshot{}
	upstreamsByNamespace := make(map[string]gloo_solo_io.UpstreamList)
	servicesByNamespace := make(map[string]KubeServiceList)
	ingressesByNamespace := make(map[string]IngressList)

	for _, namespace := range watchNamespaces {
		/* Setup namespaced watch for Upstream */
		{
			upstreams, err := c.upstream.List(namespace, clients.ListOpts{Ctx: opts.Ctx, Selector: opts.Selector})
			if err != nil {
				return nil, nil, errors.Wrapf(err, "initial Upstream list")
			}
			initialUpstreamList = append(initialUpstreamList, upstreams...)
			upstreamsByNamespace[namespace] = upstreams
		}
		upstreamNamespacesChan, upstreamErrs, err := c.upstream.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting Upstream watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, upstreamErrs, namespace+"-upstreams")
		}(namespace)
		/* Setup namespaced watch for KubeService */
		{
			services, err := c.kubeService.List(namespace, clients.ListOpts{Ctx: opts.Ctx, Selector: opts.Selector})
			if err != nil {
				return nil, nil, errors.Wrapf(err, "initial KubeService list")
			}
			initialKubeServiceList = append(initialKubeServiceList, services...)
			servicesByNamespace[namespace] = services
		}
		kubeServiceNamespacesChan, kubeServiceErrs, err := c.kubeService.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting KubeService watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, kubeServiceErrs, namespace+"-services")
		}(namespace)
		/* Setup namespaced watch for Ingress */
		{
			ingresses, err := c.ingress.List(namespace, clients.ListOpts{Ctx: opts.Ctx, Selector: opts.Selector})
			if err != nil {
				return nil, nil, errors.Wrapf(err, "initial Ingress list")
			}
			initialIngressList = append(initialIngressList, ingresses...)
			ingressesByNamespace[namespace] = ingresses
		}
		ingressNamespacesChan, ingressErrs, err := c.ingress.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting Ingress watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, ingressErrs, namespace+"-ingresses")
		}(namespace)

		/* Watch for changes and update snapshot */
		go func(namespace string) {
			for {
				select {
				case <-ctx.Done():
					return
				case upstreamList, ok := <-upstreamNamespacesChan:
					if !ok {
						return
					}
					select {
					case <-ctx.Done():
						return
					case upstreamChan <- upstreamListWithNamespace{list: upstreamList, namespace: namespace}:
					}
				case kubeServiceList, ok := <-kubeServiceNamespacesChan:
					if !ok {
						return
					}
					select {
					case <-ctx.Done():
						return
					case kubeServiceChan <- kubeServiceListWithNamespace{list: kubeServiceList, namespace: namespace}:
					}
				case ingressList, ok := <-ingressNamespacesChan:
					if !ok {
						return
					}
					select {
					case <-ctx.Done():
						return
					case ingressChan <- ingressListWithNamespace{list: ingressList, namespace: namespace}:
					}
				}
			}
		}(namespace)
	}
	/* Initialize snapshot for Upstreams */
	currentSnapshot.Upstreams = initialUpstreamList.Sort()
	/* Initialize snapshot for Services */
	currentSnapshot.Services = initialKubeServiceList.Sort()
	/* Initialize snapshot for Ingresses */
	currentSnapshot.Ingresses = initialIngressList.Sort()

	snapshots := make(chan *TranslatorSnapshot)
	go func() {
		// sent initial snapshot to kick off the watch
		initialSnapshot := currentSnapshot.Clone()
		snapshots <- &initialSnapshot

		timer := time.NewTicker(time.Second * 1)
		previousHash, err := currentSnapshot.Hash(nil)
		if err != nil {
			contextutils.LoggerFrom(ctx).Panicw("error while hashing, this should never happen", zap.Error(err))
		}
		sync := func() {
			currentHash, err := currentSnapshot.Hash(nil)
			// this should never happen, so panic if it does
			if err != nil {
				contextutils.LoggerFrom(ctx).Panicw("error while hashing, this should never happen", zap.Error(err))
			}
			if previousHash == currentHash {
				return
			}

			sentSnapshot := currentSnapshot.Clone()
			select {
			case snapshots <- &sentSnapshot:
				stats.Record(ctx, mTranslatorSnapshotOut.M(1))
				previousHash = currentHash
			default:
				stats.Record(ctx, mTranslatorSnapshotMissed.M(1))
			}
		}

		defer func() {
			close(snapshots)
			// we must wait for done before closing the error chan,
			// to avoid sending on close channel.
			done.Wait()
			close(errs)
		}()
		for {
			record := func() { stats.Record(ctx, mTranslatorSnapshotIn.M(1)) }

			select {
			case <-timer.C:
				sync()
			case <-ctx.Done():
				return
			case <-c.forceEmit:
				sentSnapshot := currentSnapshot.Clone()
				snapshots <- &sentSnapshot
			case upstreamNamespacedList, ok := <-upstreamChan:
				if !ok {
					return
				}
				record()

				namespace := upstreamNamespacedList.namespace

				skstats.IncrementResourceCount(
					ctx,
					namespace,
					"upstream",
					mTranslatorResourcesIn,
				)

				// merge lists by namespace
				upstreamsByNamespace[namespace] = upstreamNamespacedList.list
				var upstreamList gloo_solo_io.UpstreamList
				for _, upstreams := range upstreamsByNamespace {
					upstreamList = append(upstreamList, upstreams...)
				}
				currentSnapshot.Upstreams = upstreamList.Sort()
			case kubeServiceNamespacedList, ok := <-kubeServiceChan:
				if !ok {
					return
				}
				record()

				namespace := kubeServiceNamespacedList.namespace

				skstats.IncrementResourceCount(
					ctx,
					namespace,
					"kube_service",
					mTranslatorResourcesIn,
				)

				// merge lists by namespace
				servicesByNamespace[namespace] = kubeServiceNamespacedList.list
				var kubeServiceList KubeServiceList
				for _, services := range servicesByNamespace {
					kubeServiceList = append(kubeServiceList, services...)
				}
				currentSnapshot.Services = kubeServiceList.Sort()
			case ingressNamespacedList, ok := <-ingressChan:
				if !ok {
					return
				}
				record()

				namespace := ingressNamespacedList.namespace

				skstats.IncrementResourceCount(
					ctx,
					namespace,
					"ingress",
					mTranslatorResourcesIn,
				)

				// merge lists by namespace
				ingressesByNamespace[namespace] = ingressNamespacedList.list
				var ingressList IngressList
				for _, ingresses := range ingressesByNamespace {
					ingressList = append(ingressList, ingresses...)
				}
				currentSnapshot.Ingresses = ingressList.Sort()
			}
		}
	}()
	return snapshots, errs, nil
}
