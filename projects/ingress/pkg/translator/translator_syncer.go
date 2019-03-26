package translator

import (
	"context"

	"github.com/solo-io/gloo/projects/gateway/pkg/utils"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	v1 "github.com/solo-io/gloo/projects/ingress/pkg/api/v1"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
)

type translatorSyncer struct {
	writeNamespace  string
	writeErrs       chan error
	proxyClient     gloov1.ProxyClient
	ingressClient   v1.IngressClient
	proxyReconciler gloov1.ProxyReconciler
}

func NewSyncer(writeNamespace string, proxyClient gloov1.ProxyClient, ingressClient v1.IngressClient, writeErrs chan error) v1.TranslatorSyncer {
	return &translatorSyncer{
		writeNamespace:  writeNamespace,
		writeErrs:       writeErrs,
		proxyClient:     proxyClient,
		ingressClient:   ingressClient,
		proxyReconciler: gloov1.NewProxyReconciler(proxyClient),
	}
}

// TODO (ilackarms): make sure that sync happens if proxies get updated as well; may need to resync
func (s *translatorSyncer) Sync(ctx context.Context, snap *v1.TranslatorSnapshot) error {
	ctx = contextutils.WithLogger(ctx, "translatorSyncer")

	logger := contextutils.LoggerFrom(ctx)
	logger.Infof("begin sync %v (%v ingresses)", snap.Hash(),
		len(snap.Ingresses.List()))
	defer logger.Infof("end sync %v", snap.Hash())
	logger.Debugf("%v", snap)

	proxy, err := translateProxy(s.writeNamespace, snap)
	if err != nil {
		logger.Warnf("snapshot %v was rejected due to invalid config: %v\n"+
			"ingress proxy will not be updated.", snap.Hash(), err)
		return err
	}

	labels := map[string]string{
		"created_by": "ingress",
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

	return nil
}
