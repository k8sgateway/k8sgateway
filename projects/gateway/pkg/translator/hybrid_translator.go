package translator

import (
	"github.com/imdario/mergo"
	errors "github.com/rotisserie/eris"
	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	hcm "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/hcm"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/go-utils/hashutils"
)

var _ ListenerTranslator = new(HybridTranslator)

const HybridTranslatorName = "hybrid"

var (
	EmptyHybridGatewayMessage = "hybrid gateway does not have any populated matched gateways"
)

type HybridTranslator struct {
	VirtualServiceTranslator *VirtualServiceTranslator
	TcpTranslator            *TcpTranslator
}

func (t *HybridTranslator) ComputeListener(params Params, proxyName string, gateway *v1.Gateway) *gloov1.Listener {
	hybridGateway := gateway.GetHybridGateway()
	if hybridGateway == nil {
		return nil
	}

	var hybridListener *gloov1.HybridListener

	matchedGateways := hybridGateway.GetMatchedGateways()
	delegatedGateways := hybridGateway.GetDelegatedHttpGateways()
	if matchedGateways == nil && delegatedGateways == nil {
		return nil
	}

	// MatchedGateways take precedence over DelegatedHttpGateways
	if matchedGateways != nil {
		hybridListener = t.computeHybridListenerFromMatchedGateways(params, proxyName, gateway, matchedGateways)
		if len(hybridListener.GetMatchedListeners()) == 0 {
			// matched gateways are define inline, and therefore if they don't produce
			// any matched listeners, there is an error on the gateway resource
			params.reports.AddError(gateway, errors.New(EmptyHybridGatewayMessage))
			return nil
		}
	} else {
		// DelegatedHttpGateways is only processed if there are no MatchedGateways defined
		hybridListener = t.computeHybridListenerFromDelegatedGateway(params, proxyName, gateway, delegatedGateways)
		if len(hybridListener.GetMatchedListeners()) == 0 {
			// missing refs should only result in a warning
			// this allows resources to be applied asynchronously
			params.reports.AddWarning(gateway, EmptyHybridGatewayMessage)
			return nil
		}
	}

	listener := makeListener(gateway)
	listener.ListenerType = &gloov1.Listener_HybridListener{
		HybridListener: hybridListener,
	}

	if err := appendSource(listener, gateway); err != nil {
		// should never happen
		params.reports.AddError(gateway, err)
	}

	return listener
}

func (t *HybridTranslator) computeHybridListenerFromMatchedGateways(
	params Params,
	proxyName string,
	gateway *v1.Gateway,
	matchedGateways []*v1.MatchedGateway,
) *gloov1.HybridListener {
	snap := params.snapshot
	hybridListener := &gloov1.HybridListener{}
	loggedError := false

	for _, matchedGateway := range matchedGateways {
		matchedListener := &gloov1.MatchedListener{
			Matcher: &gloov1.Matcher{
				SslConfig:          matchedGateway.GetMatcher().GetSslConfig(),
				SourcePrefixRanges: matchedGateway.GetMatcher().GetSourcePrefixRanges(),
			},
		}

		switch gt := matchedGateway.GetGatewayType().(type) {
		case *v1.MatchedGateway_HttpGateway:
			if len(snap.VirtualServices) == 0 {
				if !loggedError {
					snapHash := hashutils.MustHash(snap)
					contextutils.LoggerFrom(params.ctx).Debugf("%v had no virtual services", snapHash)
					loggedError = true // only log no virtual service error once
				}
				continue
			}

			httpGateway := matchedGateway.GetHttpGateway()
			sslGateway := matchedGateway.GetMatcher().GetSslConfig() != nil
			virtualServices := getVirtualServicesForHttpGateway(params, gateway, httpGateway, sslGateway)

			matchedListener.ListenerType = &gloov1.MatchedListener_HttpListener{
				HttpListener: &gloov1.HttpListener{
					VirtualHosts: t.VirtualServiceTranslator.ComputeVirtualHosts(params, gateway, virtualServices, proxyName),
					Options:      httpGateway.GetOptions(),
				},
			}

			if sslGateway {
				virtualServices.Each(func(vs *v1.VirtualService) {
					matchedListener.SslConfigurations = append(matchedListener.GetSslConfigurations(), vs.GetSslConfig())
				})
			}

		case *v1.MatchedGateway_TcpGateway:
			matchedListener.ListenerType = &gloov1.MatchedListener_TcpListener{
				TcpListener: t.TcpTranslator.ComputeTcpListener(gt.TcpGateway),
			}
		}

		hybridListener.MatchedListeners = append(hybridListener.GetMatchedListeners(), matchedListener)
	}

	return hybridListener
}

func (t *HybridTranslator) computeHybridListenerFromDelegatedGateway(
	params Params,
	proxyName string,
	gateway *v1.Gateway,
	delegatedGateway *v1.DelegatedHttpGateway,
) *gloov1.HybridListener {
	httpGatewaySelector := NewHttpGatewaySelector(params.snapshot.HttpGateways)
	onSelectionError := func(err error) {
		params.reports.AddError(gateway, err)
	}
	matchableHttpGateways := httpGatewaySelector.SelectMatchableHttpGateways(delegatedGateway, onSelectionError)
	if len(matchableHttpGateways) == 0 {
		return nil
	}

	hybridListener := &gloov1.HybridListener{}

	matchableHttpGateways.Each(func(element *v1.MatchableHttpGateway) {
		matchedListener := t.computeMatchedListener(params, proxyName, gateway, element)
		if matchedListener != nil {
			hybridListener.MatchedListeners = append(hybridListener.GetMatchedListeners(), matchedListener)
		}
	})

	return hybridListener
}

func (t *HybridTranslator) computeMatchedListener(
	params Params,
	proxyName string,
	parentGateway *v1.Gateway,
	matchableHttpGateway *v1.MatchableHttpGateway,
) *gloov1.MatchedListener {
	// v ---- inheritence logic ---- v
	preventChildOverrides := parentGateway.GetHybridGateway().GetDelegatedHttpGateways().GetPreventChildOverrides()

	// SslConfig
	parentSslConfig := parentGateway.GetHybridGateway().GetDelegatedHttpGateways().GetSslConfig()
	var childSslConfig *gloov1.SslConfig
	if matchableHttpGateway.GetMatcher() != nil {
		childSslConfig = matchableHttpGateway.GetMatcher().GetSslConfig()
	}

	if childSslConfig == nil {
		// use parentSslConfig exactly as-is
		childSslConfig = parentSslConfig
	} else if parentSslConfig == nil {
		// use childSslConfig exactly as-is
	} else if childSslConfig != nil && parentSslConfig != nil {
		if preventChildOverrides {
			// merge, preferring parentSslConfig
			mergo.Merge(childSslConfig, parentSslConfig, mergo.WithOverride)
		} else {
			// merge, preferring childSslConfig
			mergo.Merge(childSslConfig, parentSslConfig)
		}
	}

	if matchableHttpGateway.GetMatcher() != nil {
		matchableHttpGateway.GetMatcher().SslConfig = childSslConfig
	}
	parentGateway.GetHybridGateway().GetDelegatedHttpGateways().SslConfig = childSslConfig

	// HcmOptions
	parentHcmOptions := parentGateway.GetHybridGateway().GetDelegatedHttpGateways().GetHttpConnectionManagerSettings()
	var childHcmOptions *hcm.HttpConnectionManagerSettings
	if matchableHttpGateway.GetHttpGateway().GetOptions() != nil {
		childHcmOptions = matchableHttpGateway.GetHttpGateway().GetOptions().GetHttpConnectionManagerSettings()
	}

	if childHcmOptions == nil {
		// use parentHcmOptions exactly as-is
		childHcmOptions = parentHcmOptions
	} else if parentHcmOptions == nil {
		// use childHcmOptions exactly as-is
	} else if childHcmOptions != nil && parentHcmOptions != nil {
		if preventChildOverrides {
			// merge, preferring parentHcmOptions
			mergo.Merge(childHcmOptions, parentHcmOptions, mergo.WithOverride)
		} else {
			// merge, preferring childHcmOptions
			mergo.Merge(childHcmOptions, parentHcmOptions)
		}
	}

	if matchableHttpGateway.GetHttpGateway().GetOptions() != nil {
		matchableHttpGateway.GetHttpGateway().GetOptions().HttpConnectionManagerSettings = childHcmOptions
	} else {
		matchableHttpGateway.GetHttpGateway().Options = &gloov1.HttpListenerOptions{
			HttpConnectionManagerSettings: childHcmOptions,
		}
	}
	parentGateway.GetHybridGateway().GetDelegatedHttpGateways().HttpConnectionManagerSettings = childHcmOptions
	// ^ ---- inheritence logic ---- ^

	matchedListener := &gloov1.MatchedListener{
		Matcher: &gloov1.Matcher{
			SslConfig:          matchableHttpGateway.GetMatcher().GetSslConfig(),
			SourcePrefixRanges: matchableHttpGateway.GetMatcher().GetSourcePrefixRanges(),
		},
	}

	httpGateway := matchableHttpGateway.GetHttpGateway()
	sslGateway := matchableHttpGateway.GetMatcher().GetSslConfig() != nil
	virtualServices := getVirtualServicesForHttpGateway(params, parentGateway, httpGateway, sslGateway)

	matchedListener.ListenerType = &gloov1.MatchedListener_HttpListener{
		HttpListener: &gloov1.HttpListener{
			VirtualHosts: t.VirtualServiceTranslator.ComputeVirtualHosts(params, parentGateway, virtualServices, proxyName),
			Options:      httpGateway.GetOptions(),
		},
	}

	if sslGateway {
		virtualServices.Each(func(vs *v1.VirtualService) {
			matchedListener.SslConfigurations = append(matchedListener.GetSslConfigurations(), vs.GetSslConfig())
		})
	}

	return matchedListener
}
