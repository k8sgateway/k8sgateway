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
	mApiSnapshotIn     = stats.Int64("api.gloo.solo.io/snap_emitter/snap_in", "The number of snapshots in", "1")
	mApiSnapshotOut    = stats.Int64("api.gloo.solo.io/snap_emitter/snap_out", "The number of snapshots out", "1")
	mApiSnapshotMissed = stats.Int64("api.gloo.solo.io/snap_emitter/snap_missed", "The number of snapshots missed", "1")

	apisnapshotInView = &view.View{
		Name:        "api.gloo.solo.io_snap_emitter/snap_in",
		Measure:     mApiSnapshotIn,
		Description: "The number of snapshots updates coming in",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}
	apisnapshotOutView = &view.View{
		Name:        "api.gloo.solo.io/snap_emitter/snap_out",
		Measure:     mApiSnapshotOut,
		Description: "The number of snapshots updates going out",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}
	apisnapshotMissedView = &view.View{
		Name:        "api.gloo.solo.io/snap_emitter/snap_missed",
		Measure:     mApiSnapshotMissed,
		Description: "The number of snapshots updates going missed. this can happen in heavy load. missed snapshot will be re-tried after a second.",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}
)

func init() {
	view.Register(apisnapshotInView, apisnapshotOutView, apisnapshotMissedView)
}

type ApiEmitter interface {
	Register() error
	Artifact() ArtifactClient
	Endpoint() EndpointClient
	Proxy() ProxyClient
	UpstreamGroup() UpstreamGroupClient
	Secret() SecretClient
	Upstream() UpstreamClient
	Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *ApiSnapshot, <-chan error, error)
}

func NewApiEmitter(artifactClient ArtifactClient, endpointClient EndpointClient, proxyClient ProxyClient, upstreamGroupClient UpstreamGroupClient, secretClient SecretClient, upstreamClient UpstreamClient) ApiEmitter {
	return NewApiEmitterWithEmit(artifactClient, endpointClient, proxyClient, upstreamGroupClient, secretClient, upstreamClient, make(chan struct{}))
}

func NewApiEmitterWithEmit(artifactClient ArtifactClient, endpointClient EndpointClient, proxyClient ProxyClient, upstreamGroupClient UpstreamGroupClient, secretClient SecretClient, upstreamClient UpstreamClient, emit <-chan struct{}) ApiEmitter {
	return &apiEmitter{
		artifact:      artifactClient,
		endpoint:      endpointClient,
		proxy:         proxyClient,
		upstreamGroup: upstreamGroupClient,
		secret:        secretClient,
		upstream:      upstreamClient,
		forceEmit:     emit,
	}
}

type apiEmitter struct {
	forceEmit     <-chan struct{}
	artifact      ArtifactClient
	endpoint      EndpointClient
	proxy         ProxyClient
	upstreamGroup UpstreamGroupClient
	secret        SecretClient
	upstream      UpstreamClient
}

func (c *apiEmitter) Register() error {
	if err := c.artifact.Register(); err != nil {
		return err
	}
	if err := c.endpoint.Register(); err != nil {
		return err
	}
	if err := c.proxy.Register(); err != nil {
		return err
	}
	if err := c.upstreamGroup.Register(); err != nil {
		return err
	}
	if err := c.secret.Register(); err != nil {
		return err
	}
	if err := c.upstream.Register(); err != nil {
		return err
	}
	return nil
}

func (c *apiEmitter) Artifact() ArtifactClient {
	return c.artifact
}

func (c *apiEmitter) Endpoint() EndpointClient {
	return c.endpoint
}

func (c *apiEmitter) Proxy() ProxyClient {
	return c.proxy
}

func (c *apiEmitter) UpstreamGroup() UpstreamGroupClient {
	return c.upstreamGroup
}

func (c *apiEmitter) Secret() SecretClient {
	return c.secret
}

func (c *apiEmitter) Upstream() UpstreamClient {
	return c.upstream
}

func (c *apiEmitter) Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *ApiSnapshot, <-chan error, error) {

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
	/* Create channel for Artifact */
	type artifactListWithNamespace struct {
		list      ArtifactList
		namespace string
	}
	artifactChan := make(chan artifactListWithNamespace)

	var initialArtifactList ArtifactList
	/* Create channel for Endpoint */
	type endpointListWithNamespace struct {
		list      EndpointList
		namespace string
	}
	endpointChan := make(chan endpointListWithNamespace)

	var initialEndpointList EndpointList
	/* Create channel for Proxy */
	type proxyListWithNamespace struct {
		list      ProxyList
		namespace string
	}
	proxyChan := make(chan proxyListWithNamespace)

	var initialProxyList ProxyList
	/* Create channel for UpstreamGroup */
	type upstreamGroupListWithNamespace struct {
		list      UpstreamGroupList
		namespace string
	}
	upstreamGroupChan := make(chan upstreamGroupListWithNamespace)

	var initialUpstreamGroupList UpstreamGroupList
	/* Create channel for Secret */
	type secretListWithNamespace struct {
		list      SecretList
		namespace string
	}
	secretChan := make(chan secretListWithNamespace)

	var initialSecretList SecretList
	/* Create channel for Upstream */
	type upstreamListWithNamespace struct {
		list      UpstreamList
		namespace string
	}
	upstreamChan := make(chan upstreamListWithNamespace)

	var initialUpstreamList UpstreamList

	currentSnapshot := ApiSnapshot{}

	for _, namespace := range watchNamespaces {
		/* Setup namespaced watch for Artifact */
		{
			artifacts, err := c.artifact.List(namespace, clients.ListOpts{Ctx: opts.Ctx, Selector: opts.Selector})
			if err != nil {
				return nil, nil, errors.Wrapf(err, "initial Artifact list")
			}
			initialArtifactList = append(initialArtifactList, artifacts...)
		}
		artifactNamespacesChan, artifactErrs, err := c.artifact.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting Artifact watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, artifactErrs, namespace+"-artifacts")
		}(namespace)
		/* Setup namespaced watch for Endpoint */
		{
			endpoints, err := c.endpoint.List(namespace, clients.ListOpts{Ctx: opts.Ctx, Selector: opts.Selector})
			if err != nil {
				return nil, nil, errors.Wrapf(err, "initial Endpoint list")
			}
			initialEndpointList = append(initialEndpointList, endpoints...)
		}
		endpointNamespacesChan, endpointErrs, err := c.endpoint.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting Endpoint watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, endpointErrs, namespace+"-endpoints")
		}(namespace)
		/* Setup namespaced watch for Proxy */
		{
			proxies, err := c.proxy.List(namespace, clients.ListOpts{Ctx: opts.Ctx, Selector: opts.Selector})
			if err != nil {
				return nil, nil, errors.Wrapf(err, "initial Proxy list")
			}
			initialProxyList = append(initialProxyList, proxies...)
		}
		proxyNamespacesChan, proxyErrs, err := c.proxy.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting Proxy watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, proxyErrs, namespace+"-proxies")
		}(namespace)
		/* Setup namespaced watch for UpstreamGroup */
		{
			upstreamGroups, err := c.upstreamGroup.List(namespace, clients.ListOpts{Ctx: opts.Ctx, Selector: opts.Selector})
			if err != nil {
				return nil, nil, errors.Wrapf(err, "initial UpstreamGroup list")
			}
			initialUpstreamGroupList = append(initialUpstreamGroupList, upstreamGroups...)
		}
		upstreamGroupNamespacesChan, upstreamGroupErrs, err := c.upstreamGroup.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting UpstreamGroup watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, upstreamGroupErrs, namespace+"-upstreamGroups")
		}(namespace)
		/* Setup namespaced watch for Secret */
		{
			secrets, err := c.secret.List(namespace, clients.ListOpts{Ctx: opts.Ctx, Selector: opts.Selector})
			if err != nil {
				return nil, nil, errors.Wrapf(err, "initial Secret list")
			}
			initialSecretList = append(initialSecretList, secrets...)
		}
		secretNamespacesChan, secretErrs, err := c.secret.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting Secret watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, secretErrs, namespace+"-secrets")
		}(namespace)
		/* Setup namespaced watch for Upstream */
		{
			upstreams, err := c.upstream.List(namespace, clients.ListOpts{Ctx: opts.Ctx, Selector: opts.Selector})
			if err != nil {
				return nil, nil, errors.Wrapf(err, "initial Upstream list")
			}
			initialUpstreamList = append(initialUpstreamList, upstreams...)
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

		/* Watch for changes and update snapshot */
		go func(namespace string) {
			for {
				select {
				case <-ctx.Done():
					return
				case artifactList := <-artifactNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case artifactChan <- artifactListWithNamespace{list: artifactList, namespace: namespace}:
					}
				case endpointList := <-endpointNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case endpointChan <- endpointListWithNamespace{list: endpointList, namespace: namespace}:
					}
				case proxyList := <-proxyNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case proxyChan <- proxyListWithNamespace{list: proxyList, namespace: namespace}:
					}
				case upstreamGroupList := <-upstreamGroupNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case upstreamGroupChan <- upstreamGroupListWithNamespace{list: upstreamGroupList, namespace: namespace}:
					}
				case secretList := <-secretNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case secretChan <- secretListWithNamespace{list: secretList, namespace: namespace}:
					}
				case upstreamList := <-upstreamNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case upstreamChan <- upstreamListWithNamespace{list: upstreamList, namespace: namespace}:
					}
				}
			}
		}(namespace)
	}
	/* Initialize snapshot for Artifacts */
	currentSnapshot.Artifacts = initialArtifactList.Sort()
	/* Initialize snapshot for Endpoints */
	currentSnapshot.Endpoints = initialEndpointList.Sort()
	/* Initialize snapshot for Proxies */
	currentSnapshot.Proxies = initialProxyList.Sort()
	/* Initialize snapshot for UpstreamGroups */
	currentSnapshot.UpstreamGroups = initialUpstreamGroupList.Sort()
	/* Initialize snapshot for Secrets */
	currentSnapshot.Secrets = initialSecretList.Sort()
	/* Initialize snapshot for Upstreams */
	currentSnapshot.Upstreams = initialUpstreamList.Sort()

	snapshots := make(chan *ApiSnapshot)
	go func() {
		// sent initial snapshot to kick off the watch
		initialSnapshot := currentSnapshot.Clone()
		snapshots <- &initialSnapshot

		timer := time.NewTicker(time.Second * 1)
		previousHash := initialSnapshot.Hash()
		sync := func() {
			currentHash := currentSnapshot.Hash()
			if previousHash == currentHash {
				return
			}

			sentSnapshot := currentSnapshot.Clone()
			select {
			case snapshots <- &sentSnapshot:
				stats.Record(ctx, mApiSnapshotOut.M(1))
				previousHash = currentHash
			default:
				stats.Record(ctx, mApiSnapshotMissed.M(1))
			}
		}
		artifactsByNamespace := make(map[string]ArtifactList)
		endpointsByNamespace := make(map[string]EndpointList)
		proxiesByNamespace := make(map[string]ProxyList)
		upstreamGroupsByNamespace := make(map[string]UpstreamGroupList)
		secretsByNamespace := make(map[string]SecretList)
		upstreamsByNamespace := make(map[string]UpstreamList)

		for {
			record := func() { stats.Record(ctx, mApiSnapshotIn.M(1)) }

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
			case artifactNamespacedList := <-artifactChan:
				record()

				namespace := artifactNamespacedList.namespace

				// merge lists by namespace
				artifactsByNamespace[namespace] = artifactNamespacedList.list
				var artifactList ArtifactList
				for _, artifacts := range artifactsByNamespace {
					artifactList = append(artifactList, artifacts...)
				}
				currentSnapshot.Artifacts = artifactList.Sort()
			case endpointNamespacedList := <-endpointChan:
				record()

				namespace := endpointNamespacedList.namespace

				// merge lists by namespace
				endpointsByNamespace[namespace] = endpointNamespacedList.list
				var endpointList EndpointList
				for _, endpoints := range endpointsByNamespace {
					endpointList = append(endpointList, endpoints...)
				}
				currentSnapshot.Endpoints = endpointList.Sort()
			case proxyNamespacedList := <-proxyChan:
				record()

				namespace := proxyNamespacedList.namespace

				// merge lists by namespace
				proxiesByNamespace[namespace] = proxyNamespacedList.list
				var proxyList ProxyList
				for _, proxies := range proxiesByNamespace {
					proxyList = append(proxyList, proxies...)
				}
				currentSnapshot.Proxies = proxyList.Sort()
			case upstreamGroupNamespacedList := <-upstreamGroupChan:
				record()

				namespace := upstreamGroupNamespacedList.namespace

				// merge lists by namespace
				upstreamGroupsByNamespace[namespace] = upstreamGroupNamespacedList.list
				var upstreamGroupList UpstreamGroupList
				for _, upstreamGroups := range upstreamGroupsByNamespace {
					upstreamGroupList = append(upstreamGroupList, upstreamGroups...)
				}
				currentSnapshot.UpstreamGroups = upstreamGroupList.Sort()
			case secretNamespacedList := <-secretChan:
				record()

				namespace := secretNamespacedList.namespace

				// merge lists by namespace
				secretsByNamespace[namespace] = secretNamespacedList.list
				var secretList SecretList
				for _, secrets := range secretsByNamespace {
					secretList = append(secretList, secrets...)
				}
				currentSnapshot.Secrets = secretList.Sort()
			case upstreamNamespacedList := <-upstreamChan:
				record()

				namespace := upstreamNamespacedList.namespace

				// merge lists by namespace
				upstreamsByNamespace[namespace] = upstreamNamespacedList.list
				var upstreamList UpstreamList
				for _, upstreams := range upstreamsByNamespace {
					upstreamList = append(upstreamList, upstreams...)
				}
				currentSnapshot.Upstreams = upstreamList.Sort()
			}
		}
	}()
	return snapshots, errs, nil
}
