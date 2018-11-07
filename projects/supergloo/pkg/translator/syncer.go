package translator

import (
	"context"
	"fmt"
	"github.com/solo-io/solo-kit/pkg/errors"

	gloov1 "github.com/solo-io/solo-kit/projects/gloo/pkg/api/v1"

	"github.com/solo-io/solo-kit/projects/supergloo/pkg/api/external/istio/networking/v1alpha3"
	"github.com/solo-io/solo-kit/projects/supergloo/pkg/api/v1"
)

type Syncer struct{}

func (s *Syncer) Sync(context.Context, *v1.TranslatorSnapshot) error {
	panic("implement me")
}

type translator struct{}

func getHostsForUpstream(us *gloov1.Upstream) ([]string, error) {
	switch specType := us.UpstreamSpec.UpstreamType.(type) {
	case *gloov1.UpstreamSpec_Aws:
		return nil, errors.Errorf("aws not implemented")
	case *gloov1.UpstreamSpec_Azure:
		return nil, errors.Errorf("azure not implemented")
	case *gloov1.UpstreamSpec_Kube:
		return []string{
			specType.Kube.ServiceName,
			fmt.Sprintf("%.%v.svc.cluster.local", specType.Kube.ServiceName, specType.Kube.ServiceNamespace),
		}, nil
	case *gloov1.UpstreamSpec_Static:
		var hosts []string
		for _, h := range specType.Static.Hosts {
			hosts = append(hosts, h.Addr)
		}
		return hosts, nil
	}
	return nil, errors.Errorf("unsupported upstream type %v", us)
}

func (t *translator) translateIstioRouting(routing *v1.Routing, upstreams gloov1.UpstreamList) ([]*v1alpha3.VirtualService, error) {
	var virtualServices []*v1alpha3.VirtualService
	for i, dest := range routing.DestinationRules {
		upstream, err := upstreams.Find(dest.Destination.Upstream.Namespace, dest.Destination.Upstream.Namespace)
		if err != nil {
			return nil, errors.Wrapf(err, "invalid destination for rule %v", i)
		}
		hosts, err := getHostsForUpstream(upstream)
		if err != nil {
			return nil, errors.Wrapf(err, "cannot get hosts for dest rule %v", i)
		}
		routes, err := convertHttpRules(dest.MeshHttpRules, upstreams)
		if err != nil {
			return nil, errors.Wrapf(err, "cannot get hosts for dest rule %v", i)
		}
		vs := &v1alpha3.VirtualService{
			Gateways: []string{}, // equivalent to "mesh"
			Hosts:    hosts,
			Http:     routes,
		}
		virtualServices = append(virtualServices, vs)
	}
	return virtualServices, nil
}

func convertHttpRules(rules []*v1.HTTPRule, upstreams gloov1.UpstreamList) ([]*v1alpha3.HTTPRoute, error) {
	var istioRoutes []*v1alpha3.HTTPRoute
	for _, rule := range rules {
		istioRoutes = append(istioRoutes, &v1alpha3.HTTPRoute{
			Match: convertMatch(rule.Match),
		})
	}
	return istioRoutes, nil
}

func convertMatch(match []*v1.HTTPMatchRequest) []*v1alpha3.HTTPMatchRequest {
	var istioMatch []*v1alpha3.HTTPMatchRequest
	for _, m := range match {
		istioMatch = append(istioMatch, &v1alpha3.HTTPMatchRequest{
			Uri:     convertStringMatch(m.Uri),
			Method:  convertStringMatch(m.Method),
			Headers: convertHeaders(m.Headers),
		})
	}
	return istioMatch
}
func convertHeaders(headers map[string]*v1.StringMatch) map[string]*v1alpha3.StringMatch {
	out := make(map[string]*v1alpha3.StringMatch)
	for k, v := range headers {
		out[k] = convertStringMatch(v)
	}
	return out
}

func convertStringMatch(match *v1.StringMatch) *v1alpha3.StringMatch {
	switch strMatch := match.MatchType.(type) {
	case *v1.StringMatch_Exact:
		return &v1alpha3.StringMatch{MatchType: &v1alpha3.StringMatch_Exact{Exact: strMatch.Exact}}
	case *v1.StringMatch_Prefix:
		return &v1alpha3.StringMatch{MatchType: &v1alpha3.StringMatch_Prefix{Prefix: strMatch.Prefix}}
	case *v1.StringMatch_Regex:
		return &v1alpha3.StringMatch{MatchType: &v1alpha3.StringMatch_Regex{Regex: strMatch.Regex}}
	}
	return nil
}
