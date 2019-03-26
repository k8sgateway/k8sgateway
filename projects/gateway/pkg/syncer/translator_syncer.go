package syncer

import (
	"context"

	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gateway/pkg/propagator"
	"github.com/solo-io/gloo/projects/gateway/pkg/translator"
	"github.com/solo-io/gloo/projects/gateway/pkg/utils"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/reporter"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
)

type translatorSyncer struct {
	writeNamespace  string
	reporter        reporter.Reporter
	propagator      *propagator.Propagator
	proxyClient     gloov1.ProxyClient
	gwClient        v1.GatewayClient
	vsClient        v1.VirtualServiceClient
	proxyReconciler gloov1.ProxyReconciler
}

func NewTranslatorSyncer(writeNamespace string, proxyClient gloov1.ProxyClient, gwClient v1.GatewayClient, vsClient v1.VirtualServiceClient, reporter reporter.Reporter, propagator *propagator.Propagator) v1.ApiSyncer {
	return &translatorSyncer{
		writeNamespace:  writeNamespace,
		reporter:        reporter,
		propagator:      propagator,
		proxyClient:     proxyClient,
		gwClient:        gwClient,
		vsClient:        vsClient,
		proxyReconciler: gloov1.NewProxyReconciler(proxyClient),
	}
}

// TODO (ilackarms): make sure that sync happens if proxies get updated as well; may need to resync
func (s *translatorSyncer) Sync(ctx context.Context, snap *v1.ApiSnapshot) error {
	ctx = contextutils.WithLogger(ctx, "translatorSyncer")

	logger := contextutils.LoggerFrom(ctx)
	logger.Infof("begin sync %v (%v virtual services, %v gateways)", snap.Hash(),
		len(snap.VirtualServices), len(snap.Gateways))
	defer logger.Infof("end sync %v", snap.Hash())
	logger.Debugf("%v", snap)

	proxy, resourceErrs := translator.Translate(ctx, s.writeNamespace, snap)
	if err := resourceErrs.Validate(); err != nil {
		if err := s.reporter.WriteReports(ctx, resourceErrs, nil); err != nil {
			contextutils.LoggerFrom(ctx).Errorf("failed to write reports: %v", err)
		}
		logger.Warnf("snapshot %v was rejected due to invalid config: %v\nxDS cache will not be updated.", snap.Hash(), err)
		return err
	}

	labels := map[string]string{
		"created_by": "gateway",
	}

	var desiredResources gloov1.ProxyList
	if proxy != nil {
		logger.Infof("creating proxy %v", proxy.Metadata.Ref())
		proxy.Metadata.Labels = labels
		desiredResources = gloov1.ProxyList{proxy}
	}

	if err := s.proxyReconciler.Reconcile(s.writeNamespace, desiredResources, utils.TransitionFunction, clients.ListOpts{
		Ctx:      ctx,
		Selector: labels,
	}); err != nil {
		return err
	}

	// start propagating for new set of resources
	if err := s.propagateProxyStatus(ctx, snap, proxy, resourceErrs); err != nil {
		return err
	}

	return nil
}

func (s *translatorSyncer) propagateProxyStatus(ctx context.Context, snap *v1.ApiSnapshot, proxy *gloov1.Proxy, resourceErrs reporter.ResourceErrors) error {
	if proxy == nil {
		return nil
	}
	statuses, err := watchProxyStatus(ctx, s.proxyClient, proxy)
	if err != nil {
		return err
	}
	var lastStatus core.Status
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case status := <-statuses:
				if status.Equal(lastStatus) {
					continue
				}
				lastStatus = status
				subresourceStatuses := map[string]*core.Status{
					resources.Key(proxy): &status,
				}
				err := s.reporter.WriteReports(ctx, resourceErrs, subresourceStatuses)
				if err != nil {
					contextutils.LoggerFrom(ctx).Errorf("err: updating dependent statuses: %v", err)
				}
			}
		}
	}()
	return nil
}

func watchProxyStatus(ctx context.Context, proxyClient gloov1.ProxyClient, proxy *gloov1.Proxy) (<-chan core.Status, error) {
	ctx = contextutils.WithLogger(ctx, "proxy-err-propagator")
	proxies, errs, err := proxyClient.Watch(proxy.Metadata.Namespace, clients.WatchOpts{
		Ctx:      ctx,
		Selector: proxy.Metadata.Labels,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "creating watch for proxy %v", proxy.Metadata.Ref())
	}
	statuses := make(chan core.Status)
	go func() {
		defer close(statuses)
		for {
			select {
			case <-ctx.Done():
				return
			case err := <-errs:
				contextutils.LoggerFrom(ctx).Error(err)
			case list := <-proxies:
				proxy, err := list.Find(proxy.Metadata.Namespace, proxy.Metadata.Name)
				if err != nil {
					contextutils.LoggerFrom(ctx).Error(err)
					continue
				}
				select {
				case <-ctx.Done():
					return
				case statuses <- proxy.Status:
				}
			}
		}
	}()

	return statuses, nil
}
