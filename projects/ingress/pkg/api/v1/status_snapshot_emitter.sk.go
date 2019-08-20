// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"sync"
	"time"

	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"

	"github.com/solo-io/go-utils/errutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/errors"
)

var (
	mStatusSnapshotIn     = stats.Int64("status.ingress.solo.io/snap_emitter/snap_in", "The number of snapshots in", "1")
	mStatusSnapshotOut    = stats.Int64("status.ingress.solo.io/snap_emitter/snap_out", "The number of snapshots out", "1")
	mStatusSnapshotMissed = stats.Int64("status.ingress.solo.io/snap_emitter/snap_missed", "The number of snapshots missed", "1")

	statussnapshotInView = &view.View{
		Name:        "status.ingress.solo.io_snap_emitter/snap_in",
		Measure:     mStatusSnapshotIn,
		Description: "The number of snapshots updates coming in",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}
	statussnapshotOutView = &view.View{
		Name:        "status.ingress.solo.io/snap_emitter/snap_out",
		Measure:     mStatusSnapshotOut,
		Description: "The number of snapshots updates going out",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}
	statussnapshotMissedView = &view.View{
		Name:        "status.ingress.solo.io/snap_emitter/snap_missed",
		Measure:     mStatusSnapshotMissed,
		Description: "The number of snapshots updates going missed. this can happen in heavy load. missed snapshot will be re-tried after a second.",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}
)

func init() {
	view.Register(statussnapshotInView, statussnapshotOutView, statussnapshotMissedView)
}

type StatusEmitter interface {
	Register() error
	KubeService() KubeServiceClient
	Ingress() IngressClient
	Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *StatusSnapshot, <-chan error, error)
}

func NewStatusEmitter(kubeServiceClient KubeServiceClient, ingressClient IngressClient) StatusEmitter {
	return NewStatusEmitterWithEmit(kubeServiceClient, ingressClient, make(chan struct{}))
}

func NewStatusEmitterWithEmit(kubeServiceClient KubeServiceClient, ingressClient IngressClient, emit <-chan struct{}) StatusEmitter {
	return &statusEmitter{
		kubeService: kubeServiceClient,
		ingress:     ingressClient,
		forceEmit:   emit,
	}
}

type statusEmitter struct {
	forceEmit   <-chan struct{}
	kubeService KubeServiceClient
	ingress     IngressClient
}

func (c *statusEmitter) Register() error {
	if err := c.kubeService.Register(); err != nil {
		return err
	}
	if err := c.ingress.Register(); err != nil {
		return err
	}
	return nil
}

func (c *statusEmitter) KubeService() KubeServiceClient {
	return c.kubeService
}

func (c *statusEmitter) Ingress() IngressClient {
	return c.ingress
}

func (c *statusEmitter) Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *StatusSnapshot, <-chan error, error) {

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

	currentSnapshot := StatusSnapshot{}

	for _, namespace := range watchNamespaces {
		/* Setup namespaced watch for KubeService */
		{
			services, err := c.kubeService.List(namespace, clients.ListOpts{Ctx: opts.Ctx, Selector: opts.Selector})
			if err != nil {
				return nil, nil, errors.Wrapf(err, "initial KubeService list")
			}
			initialKubeServiceList = append(initialKubeServiceList, services...)
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
				case kubeServiceList := <-kubeServiceNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case kubeServiceChan <- kubeServiceListWithNamespace{list: kubeServiceList, namespace: namespace}:
					}
				case ingressList := <-ingressNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case ingressChan <- ingressListWithNamespace{list: ingressList, namespace: namespace}:
					}
				}
			}
		}(namespace)
	}
	/* Initialize snapshot for Services */
	currentSnapshot.Services = initialKubeServiceList.Sort()
	/* Initialize snapshot for Ingresses */
	currentSnapshot.Ingresses = initialIngressList.Sort()

	snapshots := make(chan *StatusSnapshot)
	go func() {
		// sent initial snapshot to kick off the watch
		initialSnapshot := currentSnapshot.Clone()
		snapshots <- &initialSnapshot

		timer := time.NewTicker(time.Second * 1)
		var previousHash uint64
		sync := func() {
			currentHash := currentSnapshot.Hash()
			if previousHash == currentHash {
				return
			}

			sentSnapshot := currentSnapshot.Clone()
			select {
			case snapshots <- &sentSnapshot:
				stats.Record(ctx, mStatusSnapshotOut.M(1))
				previousHash = currentHash
			default:
				stats.Record(ctx, mStatusSnapshotMissed.M(1))
			}
		}
		servicesByNamespace := make(map[string]KubeServiceList)
		ingressesByNamespace := make(map[string]IngressList)

		for {
			record := func() { stats.Record(ctx, mStatusSnapshotIn.M(1)) }

			select {
			case <-timer.C:
				sync()
			case <-ctx.Done():
				close(snapshots)
				done.Wait()
				close(errs)
				return
			case <-c.forceEmit:
				sentSnapshot := currentSnapshot.Clone()
				snapshots <- &sentSnapshot
			case kubeServiceNamespacedList := <-kubeServiceChan:
				record()

				namespace := kubeServiceNamespacedList.namespace

				// merge lists by namespace
				servicesByNamespace[namespace] = kubeServiceNamespacedList.list
				var kubeServiceList KubeServiceList
				for _, services := range servicesByNamespace {
					kubeServiceList = append(kubeServiceList, services...)
				}
				currentSnapshot.Services = kubeServiceList.Sort()
			case ingressNamespacedList := <-ingressChan:
				record()

				namespace := ingressNamespacedList.namespace

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
