package translator

import (
	"context"
	"fmt"
	"strings"

	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	"github.com/solo-io/go-utils/hashutils"

	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/api/v2/reporter"
)

//go:generate mockgen -destination mocks/mock_translator.go -package mocks github.com/solo-io/gloo/projects/gateway/pkg/translator Translator
type Translator interface {
	Translate(ctx context.Context, proxyName, namespace string, snap *v1.ApiSnapshot, filteredGateways v1.GatewayList) (*gloov1.Proxy, reporter.ResourceReports)
}

type GwTranslator struct {
	// listenerTranslators is the set of available translators that convert Gloo Gateways into Envoy Listeners
	listenerTranslators map[string]ListenerTranslator

	// writeNamespace is the namespace that all Proxy CRs will be written to
	writeNamespace string

	// predicate is used to determine which Gateways to process during translation
	predicate      Predicate
}

func NewDefaultTranslator(opts Opts) *GwTranslator {
	warnOnRouteShortCircuiting := false
	if opts.Validation != nil {
		warnOnRouteShortCircuiting = opts.Validation.WarnOnRouteShortCircuiting
	}

	httpTranslator := &HttpTranslator{
		WarnOnRouteShortCircuiting: warnOnRouteShortCircuiting,
	}
	tcpTranslator := &TcpTranslator{}
	hybridTranslator := &HybridTranslator{
		HttpTranslator: httpTranslator,
		TcpTranslator:  tcpTranslator,
	}

	return NewTranslator([]ListenerTranslator{httpTranslator, tcpTranslator, hybridTranslator}, opts)
}

func NewTranslator(listenerTranslators []ListenerTranslator, opts Opts) *GwTranslator {
	translatorsByName := make(map[string]ListenerTranslator)
	for _, t := range listenerTranslators {
		translatorsByName[t.Name()] = t
	}

	return &GwTranslator{
		listenerTranslators: translatorsByName,
		predicate:           GetPredicate(opts.WriteNamespace, opts.ReadGatewaysFromAllNamespaces),
		writeNamespace:      opts.WriteNamespace,
	}
}

func (t *GwTranslator) Translate(ctx context.Context, proxyName string, snap *v1.ApiSnapshot, gateways v1.GatewayList) (*gloov1.Proxy, reporter.ResourceReports) {
	logger := contextutils.LoggerFrom(ctx)

	reports := make(reporter.ResourceReports)
	reports.Accept(snap.Gateways.AsInputResources()...)
	reports.Accept(snap.VirtualServices.AsInputResources()...)
	reports.Accept(snap.RouteTables.AsInputResources()...)

	// NOTE: We could optimize this by removing this responsibility from the Translator
	//	that way it is called once per translation run, as opposed to once per proxy
	//	However, we must perform this both for translation syncs and validation syncs
	//	and it would be more challenging to keep those two in sync.
	filteredGateways := FilterGateways(gateways, t.predicate)
	if len(filteredGateways) == 0 {
		snapHash := hashutils.MustHash(snap)
		logger.Infof("%v had no gateways", snapHash)
		return nil, reports
	}

	params := NewTranslatorParams(ctx, snap, reports)
	validateGateways(filteredGateways, snap.VirtualServices, reports)

	listeners := make([]*gloov1.Listener, 0, len(filteredGateways))
	for _, gateway := range filteredGateways {
		listenerTranslator := t.getListenerTranslatorForGateway(gateway)
		listener := listenerTranslator.ComputeListener(params, proxyName, gateway)
		if listener != nil {
			listeners = append(listeners, listener)
		}
	}

	if len(listeners) == 0 {
		return nil, reports
	}

	return &gloov1.Proxy{
		Metadata: &core.Metadata{
			Name:      proxyName,
			Namespace: t.writeNamespace,
		},
		Listeners: listeners,
	}, reports
}

func (t *GwTranslator) getListenerTranslatorForGateway(gateway *v1.Gateway) ListenerTranslator {
	var listenerTranslatorImpl ListenerTranslator

	switch gateway.GetGatewayType().(type) {
	case *v1.Gateway_HttpGateway:
		listenerTranslatorImpl = t.listenerTranslators[HttpTranslatorName]

	case *v1.Gateway_TcpGateway:
		listenerTranslatorImpl = t.listenerTranslators[TcpTranslatorName]

	case *v1.Gateway_HybridGateway:
		listenerTranslatorImpl = t.listenerTranslators[HybridTranslatorName]
	}

	if listenerTranslatorImpl == nil {
		// This should not happen, but will occur when a user has not defined any GatewayType
		return &InvalidGatewayTypeTranslator{}
	}

	return listenerTranslatorImpl
}

func makeListener(gateway *v1.Gateway) *gloov1.Listener {
	return &gloov1.Listener{
		Name:          ListenerName(gateway),
		BindAddress:   gateway.GetBindAddress(),
		BindPort:      gateway.GetBindPort(),
		Options:       gateway.GetOptions(),
		UseProxyProto: gateway.GetUseProxyProto(),
		RouteOptions:  gateway.GetRouteOptions(),
	}
}

func ListenerName(gateway *v1.Gateway) string {
	return fmt.Sprintf("listener-%s-%d", gateway.GetBindAddress(), gateway.GetBindPort())
}

func validateGateways(gateways v1.GatewayList, virtualServices v1.VirtualServiceList, reports reporter.ResourceReports) {
	bindAddresses := map[string]v1.GatewayList{}
	// if two gateway (=listener) that belong to the same proxy share the same bind address,
	// they are invalid.
	for _, gw := range gateways {
		bindAddress := fmt.Sprintf("%s:%d", gw.GetBindAddress(), gw.GetBindPort())
		bindAddresses[bindAddress] = append(bindAddresses[bindAddress], gw)

		var gatewayVirtualServices []*core.ResourceRef
		switch gatewayType := gw.GetGatewayType().(type) {
		case *v1.Gateway_HttpGateway:
			gatewayVirtualServices = gatewayType.HttpGateway.GetVirtualServices()
		case *v1.Gateway_HybridGateway:
			for _, matchedGateway := range gatewayType.HybridGateway.GetMatchedGateways() {
				if httpGateway := matchedGateway.GetHttpGateway(); httpGateway != nil {
					gatewayVirtualServices = append(gatewayVirtualServices, httpGateway.GetVirtualServices()...)
				}
			}
		}

		for _, vs := range gatewayVirtualServices {
			if _, err := virtualServices.Find(vs.Strings()); err != nil {
				reports.AddError(gw, fmt.Errorf("invalid virtual service ref %v", vs))
			}
		}
	}

	for addr, gateways := range bindAddresses {
		if len(gateways) > 1 {
			for _, gw := range gateways {
				reports.AddError(gw, fmt.Errorf("bind-address %s is not unique in a proxy. gateways: %s", addr, strings.Join(gatewaysRefsToString(gateways), ",")))
			}
		}
	}
}

func gatewaysRefsToString(gateways v1.GatewayList) []string {
	var ret []string
	for _, gw := range gateways {
		ret = append(ret, gw.GetMetadata().Ref().Key())
	}
	return ret
}
