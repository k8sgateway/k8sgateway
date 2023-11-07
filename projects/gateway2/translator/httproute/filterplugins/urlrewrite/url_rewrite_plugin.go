package urlrewrite

import (
	"context"

	routev3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	matcherv3 "github.com/envoyproxy/go-control-plane/envoy/type/matcher/v3"
	errors "github.com/rotisserie/eris"
	"github.com/solo-io/gloo/pkg/utils/regexutils"

	gwv1 "sigs.k8s.io/gateway-api/apis/v1"
)

type Plugin struct{}

func NewPlugin() *Plugin {
	return &Plugin{}
}

func (p *Plugin) ApplyFilter(
	ctx context.Context,
	filter gwv1.HTTPRouteFilter,
	outputRoute *routev3.Route,
) error {
	config := filter.URLRewrite
	if config == nil {
		return errors.Errorf("UrlRewrite filter supplied does not define urlRewrite config")
	}

	routeAction := outputRoute.GetRoute()
	if routeAction == nil {
		return errors.Errorf("UrlRewrite filter supplied to route without a route action")
	}

	if config.Hostname != nil {

		routeAction.HostRewriteSpecifier = &routev3.RouteAction_HostRewriteLiteral{
			HostRewriteLiteral: string(*config.Hostname),
		}
	}

	if config.Path != nil {
		switch config.Path.Type {
		case gwv1.FullPathHTTPPathModifier:
			if config.Path.ReplaceFullPath == nil {
				return errors.Errorf("UrlRewrite filter supplied with Full Path rewrite type, but no Full Path supplied")
			}

			routeAction.RegexRewrite = &matcherv3.RegexMatchAndSubstitute{
				Pattern:      regexutils.NewRegexWithProgramSize(".*", nil),
				Substitution: *config.Path.ReplaceFullPath,
			}
		case gwv1.PrefixMatchHTTPPathModifier:
			if config.Path.ReplacePrefixMatch == nil {
				return errors.Errorf("UrlRewrite filter supplied with prefix rewrite type, but no prefix supplied")
			}
			routeAction.PrefixRewrite = *config.Path.ReplacePrefixMatch
		}
	}

	return nil
}
