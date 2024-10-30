package krtcollections

import (
	"context"
	"errors"
	"fmt"
	"maps"
	"slices"
	"strings"
	"sync"
	"sync/atomic"

	envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	envoy_service_discovery_v3 "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/solo-io/gloo/projects/gateway2/utils"
	"github.com/solo-io/gloo/projects/gloo/pkg/xds"
	xdsserver "github.com/solo-io/solo-kit/pkg/api/v1/control-plane/server"
	"google.golang.org/protobuf/types/known/structpb"
	"istio.io/istio/pkg/kube/controllers"
	"istio.io/istio/pkg/kube/krt"
	"k8s.io/apimachinery/pkg/types"
)

type ConnectedClient struct {
	uniqueClientName string
}

func newConnectedClient(uniqueClientName string) ConnectedClient {
	return ConnectedClient{
		uniqueClientName: uniqueClientName,
	}
}

type UniqlyConnectedClient struct {
	// name namespace of gateway
	Role      string
	Labels    map[string]string
	Locality  PodLocality
	Namespace string

	snapshotKey  string
	resourceName string
}

func (c UniqlyConnectedClient) ResourceName() string {
	return c.resourceName
}

var _ krt.Equaler[UniqlyConnectedClient] = new(UniqlyConnectedClient)

func (c UniqlyConnectedClient) Equals(k UniqlyConnectedClient) bool {
	return c.Role == k.Role && c.Namespace == k.Namespace && c.Locality == k.Locality && maps.Equal(c.Labels, k.Labels)
}

func labeledRole(role string, labels map[string]string) string {
	return fmt.Sprintf("%s%s%d", role, xds.KeyDelimiter, utils.HashLabels(labels))
}

func newUniqlyConnectedClient(node *envoy_config_core_v3.Node, ns string, labels map[string]string, locality PodLocality) UniqlyConnectedClient {
	role := node.GetMetadata().GetFields()[xds.RoleKey].GetStringValue()
	snapshotKey := labeledRole(role, labels)
	return UniqlyConnectedClient{
		Role:         role,
		Namespace:    ns,
		Locality:     locality,
		Labels:       labels,
		snapshotKey:  snapshotKey,
		resourceName: fmt.Sprintf("%s%s%s", snapshotKey, xds.KeyDelimiter, ns),
	}
}

type callbacksCollection struct {
	ctx              context.Context
	clients          map[int64]ConnectedClient
	uniqClientsCount map[string]uint64
	uniqClients      map[string]UniqlyConnectedClient
	fanoutChan       chan krt.Event[UniqlyConnectedClient]
	stateLock        sync.RWMutex

	eventHandlers handlers[UniqlyConnectedClient]
	augmentedPods krt.Collection[LocalityPod]
}

type callbacks struct {
	collection atomic.Pointer[callbacksCollection]
}

type UniquelyConnectedClientsBulider func(ctx context.Context, augmentedPods krt.Collection[LocalityPod]) krt.Collection[UniqlyConnectedClient]

// THIS IS THE SET OF THINGS WE RUN TRANSLATION FOR
// add returned callbacks to the xds server.

func NewUniquelyConnectedClients() (xdsserver.Callbacks, UniquelyConnectedClientsBulider) {
	cb := &callbacks{}
	return cb, buildCollection(cb)
}

func buildCollection(callbacks *callbacks) UniquelyConnectedClientsBulider {
	return func(ctx context.Context, augmentedPods krt.Collection[LocalityPod]) krt.Collection[UniqlyConnectedClient] {
		cb := &callbacksCollection{
			ctx:              ctx,
			clients:          make(map[int64]ConnectedClient),
			uniqClientsCount: make(map[string]uint64),
			uniqClients:      make(map[string]UniqlyConnectedClient),
			fanoutChan:       make(chan krt.Event[UniqlyConnectedClient], 100),
			augmentedPods:    augmentedPods,
		}
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event := <-cb.fanoutChan:
					cb.eventHandlers.Notify(event)
				}
			}
		}()
		callbacks.collection.Store(cb)
		return cb
	}
}

var _ xdsserver.Callbacks = new(callbacks)
var _ krt.Collection[UniqlyConnectedClient] = new(callbacksCollection)

// OnStreamOpen is called once an xDS stream is open with a stream ID and the type URL (or "" for ADS).
// Returning an error will end processing and close the stream. OnStreamClosed will still be called.
func (x *callbacks) OnStreamOpen(_ context.Context, _ int64, _ string) error {
	return nil
}

// OnStreamClosed is called immediately prior to closing an xDS stream with a stream ID.
func (x *callbacks) OnStreamClosed(sid int64) {
	c := x.collection.Load()
	if c == nil {
		return
	}
	c.OnStreamClosed(sid)
}

func (x *callbacksCollection) OnStreamClosed(sid int64) {
	ucc := x.del(sid)
	if ucc != nil {
		x.Notify(krt.Event[UniqlyConnectedClient]{Old: ucc, Event: controllers.EventDelete})
	}
}

func (x *callbacksCollection) del(sid int64) *UniqlyConnectedClient {
	x.stateLock.Lock()
	defer x.stateLock.Unlock()

	c, ok := x.clients[sid]
	delete(x.clients, sid)
	if ok {
		resouceName := c.uniqueClientName
		current := x.uniqClientsCount[resouceName]
		x.uniqClientsCount[resouceName] = current - 1
		if current == 1 {
			ucc := x.uniqClients[resouceName]
			delete(x.uniqClientsCount, resouceName)
			delete(x.uniqClients, resouceName)
			return &ucc
		}
	}
	return nil
}

func (x *callbacksCollection) add(sid int64, r *envoy_service_discovery_v3.DiscoveryRequest) (*UniqlyConnectedClient, error) {

	var pod *LocalityPod
	if r.Node != nil {
		podRef := getRef(r.Node)
		k := krt.Key[LocalityPod](krt.Named{Name: podRef.Name, Namespace: podRef.Namespace}.ResourceName())
		pod = x.augmentedPods.GetKey(k)
	}

	x.stateLock.Lock()
	defer x.stateLock.Unlock()
	if _, ok := x.clients[sid]; !ok {
		if pod == nil {
			// error if we can't get the pod
			return nil, fmt.Errorf("pod not found for node %v", r.Node)
		}
		// TODO: modify request to include the label that are relevant for the client?
		ucc := newUniqlyConnectedClient(r.Node, pod.Namespace, pod.AugmentedLabels, pod.Locality)
		c := newConnectedClient(ucc.resourceName)
		x.clients[sid] = c
		x.uniqClientsCount[ucc.resourceName] += 1
		x.uniqClients[ucc.resourceName] = ucc
		return &ucc, nil
	}
	return nil, nil
}

func (x *callbacksCollection) Notify(e krt.Event[UniqlyConnectedClient]) {
	for {
		// note: do not use a default block here, we want to block if the channel is full, as otherwise we will have inconsistent state in krt.
		select {
		case x.fanoutChan <- e:
			return
		case <-x.ctx.Done():
			return
		}
	}
}

// OnStreamRequest is called once a request is received on a stream.
// Returning an error will end processing and close the stream. OnStreamClosed will still be called.
func (x *callbacks) OnStreamRequest(sid int64, r *envoy_service_discovery_v3.DiscoveryRequest) error {
	role := r.GetNode().GetMetadata().GetFields()[xds.RoleKey].GetStringValue()
	// as gloo-edge and ggv2 share a control plane, check that this collection only handles ggv2 clients
	if !xds.IsKubeGatewayCacheKey(role) {
		return nil
	}
	c := x.collection.Load()
	if c == nil {
		return errors.New("ggv2 not initialized")
	}
	return c.OnStreamRequest(sid, r)
}

func (x *callbacksCollection) OnStreamRequest(sid int64, r *envoy_service_discovery_v3.DiscoveryRequest) error {
	ucc, err := x.add(sid, r)
	if err != nil {
		return err
	}
	if ucc != nil {
		nodeMd := r.GetNode().GetMetadata()
		if nodeMd == nil {
			nodeMd = &structpb.Struct{}
		}
		if nodeMd.Fields == nil {
			nodeMd.Fields = map[string]*structpb.Value{}
		}
		role := nodeMd.Fields[xds.RoleKey].GetStringValue()
		if role != "" {
			nodeMd.Fields[xds.RoleKey] = structpb.NewStringValue(ucc.resourceName)
			r.Node.Metadata = nodeMd
		} else {
			// TODO: log warning
		}
		x.Notify(krt.Event[UniqlyConnectedClient]{New: ucc, Event: controllers.EventAdd})
	}
	return nil
}

func (x *callbacksCollection) getClients() []UniqlyConnectedClient {
	x.stateLock.RLock()
	defer x.stateLock.RUnlock()
	clients := make([]UniqlyConnectedClient, 0, len(x.uniqClients))
	for _, c := range x.uniqClients {
		clients = append(clients, c)
	}
	return clients
}

// OnStreamResponse is called immediately prior to sending a response on a stream.
func (x *callbacks) OnStreamResponse(_ int64, _ *envoy_service_discovery_v3.DiscoveryRequest, _ *envoy_service_discovery_v3.DiscoveryResponse) {
}

// OnFetchRequest is called for each Fetch request. Returning an error will end processing of the
// request and respond with an error.
func (x *callbacks) OnFetchRequest(_ context.Context, _ *envoy_service_discovery_v3.DiscoveryRequest) error {
	return nil
}

// OnFetchResponse is called immediately prior to sending a response.
func (x *callbacks) OnFetchResponse(_ *envoy_service_discovery_v3.DiscoveryRequest, _ *envoy_service_discovery_v3.DiscoveryResponse) {
}

func (x *callbacksCollection) Register(f func(o krt.Event[UniqlyConnectedClient])) krt.Syncer {
	return x.RegisterBatch(func(events []krt.Event[UniqlyConnectedClient], initialSync bool) {
		for _, o := range events {
			f(o)
		}
	}, true)
}

func (x *callbacksCollection) RegisterBatch(f func(o []krt.Event[UniqlyConnectedClient], initialSync bool), runExistingState bool) krt.Syncer {
	if runExistingState {
		f(nil, true)
	}
	// notify under lock to make sure that no regular events fire during the initial sync.
	x.eventHandlers.mu.Lock()
	defer x.eventHandlers.mu.Unlock()
	x.eventHandlers.insertLocked(f)
	if runExistingState {
		var events []krt.Event[UniqlyConnectedClient]
		for _, v := range x.getClients() {
			events = append(events, krt.Event[UniqlyConnectedClient]{
				New:   &v,
				Event: controllers.EventAdd,
			})

		}
		f(events, true)
	}

	return &simpleSyncer{}
}

func (x *callbacksCollection) Synced() krt.Syncer {
	return &simpleSyncer{}
}

// GetKey returns an object by its key, if present. Otherwise, nil is returned.

func (x *callbacksCollection) GetKey(k krt.Key[UniqlyConnectedClient]) *UniqlyConnectedClient {
	x.stateLock.RLock()
	defer x.stateLock.RUnlock()
	u, ok := x.uniqClients[string(k)]
	if ok {
		return &u
	}
	return nil
}

// List returns all objects in the collection.
// Order of the list is undefined.

func (x *callbacksCollection) List() []UniqlyConnectedClient { return x.getClients() }

type simpleSyncer struct{}

func (s *simpleSyncer) WaitUntilSynced(stop <-chan struct{}) bool {
	return true
}

func (s *simpleSyncer) HasSynced() bool {
	return true
}

type handlers[O any] struct {
	mu sync.RWMutex
	h  []func(o []krt.Event[O], initialSync bool)
}

func (o *handlers[O]) Insert(f func(o []krt.Event[O], initialSync bool)) {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.h = append(o.h, f)
}
func (o *handlers[O]) insertLocked(f func(o []krt.Event[O], initialSync bool)) {
	o.h = append(o.h, f)
}

func (o *handlers[O]) Get() []func(o []krt.Event[O], initialSync bool) {
	o.mu.RLock()
	defer o.mu.RUnlock()
	return slices.Clone(o.h)
}

func (o *handlers[O]) Notify(e krt.Event[O]) {
	cb := o.Get()
	for _, f := range cb {
		events := [1]krt.Event[O]{e}
		f(events[:], false)
	}
}

func getRef(node *envoy_config_core_v3.Node) types.NamespacedName {
	nns := node.Id
	split := strings.SplitN(nns, ".", 2)
	if len(split) != 2 {
		return types.NamespacedName{}
	}
	return types.NamespacedName{
		Name:      split[0],
		Namespace: split[1],
	}
}
