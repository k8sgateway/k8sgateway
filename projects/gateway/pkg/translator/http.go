package translator

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/solo-io/gloo/projects/gloo/pkg/utils"

	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/core/matchers"

	"github.com/solo-io/go-utils/hashutils"

	errors "github.com/rotisserie/eris"

	"k8s.io/apimachinery/pkg/labels"

	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v2/reporter"
)

var (
	NoVirtualHostErr = func(vs *v1.VirtualService) error {
		return errors.Errorf("virtual service [%s] does not specify a virtual host", vs.Metadata.Ref().Key())
	}
	DomainInOtherVirtualServicesErr = func(domain string, conflictingVsRefs []string) error {
		if domain == "" {
			return errors.Errorf("domain conflict: other virtual services that belong to the same Gateway"+
				" as this one don't specify a domain (and thus default to '*'): %v", conflictingVsRefs)
		}
		return errors.Errorf("domain conflict: the [%s] domain is present in other virtual services "+
			"that belong to the same Gateway as this one: %v", domain, conflictingVsRefs)
	}
	GatewayHasConflictingVirtualServicesErr = func(conflictingDomains []string) error {
		var loggedDomains []string
		for _, domain := range conflictingDomains {
			if domain == "" {
				domain = "EMPTY_DOMAIN"
			}
			loggedDomains = append(loggedDomains, domain)
		}
		return errors.Errorf("domain conflict: the following domains are present in more than one of the "+
			"virtual services associated with this gateway: %v", loggedDomains)
	}
	ConflictingMatcherErr = func(vh string, matcher *matchers.Matcher) error {
		return errors.Errorf("virtual host [%s] has conflicting matcher: %v", vh, matcher)
	}
	UnorderedPrefixErr = func(vh, prefix string, matcher *matchers.Matcher) error {
		return errors.Errorf("virtual host [%s] has unordered prefix routes, earlier prefix [%s] short-circuited "+
			"later route [%v]", vh, prefix, matcher)
	}
	UnorderedRegexErr = func(vh, regex string, matcher *matchers.Matcher) error {
		return errors.Errorf("virtual host [%s] has unordered regex routes, earlier regex [%s] short-circuited "+
			"later route [%v]", vh, regex, matcher)
	}
)

type HttpTranslator struct {
	WarnOnRouteShortCircuiting bool
}

func (t *HttpTranslator) GenerateListeners(ctx context.Context, snap *v1.ApiSnapshot, filteredGateways []*v1.Gateway, reports reporter.ResourceReports) []*gloov1.Listener {
	if len(snap.VirtualServices) == 0 {
		snapHash := hashutils.MustHash(snap)
		contextutils.LoggerFrom(ctx).Debugf("%v had no virtual services", snapHash)
		return nil
	}
	var result []*gloov1.Listener
	for _, gateway := range filteredGateways {
		if gateway.GetHttpGateway() == nil {
			continue
		}

		virtualServices := getVirtualServicesForGateway(gateway, snap.VirtualServices)
		validateVirtualServiceDomains(gateway, virtualServices, reports)
		listener := t.desiredListenerForHttp(gateway, virtualServices, snap.RouteTables, reports)
		result = append(result, listener)
	}
	return result
}

// Errors will be added to the report object.
func validateVirtualServiceDomains(gateway *v1.Gateway, virtualServices v1.VirtualServiceList, reports reporter.ResourceReports) {

	// Index the virtual services for this gateway by the domain
	vsByDomain := map[string]v1.VirtualServiceList{}
	for _, vs := range virtualServices {

		// Add warning and skip if no virtual host
		if vs.VirtualHost == nil {
			reports.AddWarning(vs, NoVirtualHostErr(vs).Error())
			continue
		}

		// Not specifying any domains is not an error per se, but we need to check whether multiple virtual services
		// don't specify any, so we use the empty string as a placeholder in this function.
		domains := append([]string{}, vs.VirtualHost.Domains...)
		if len(domains) == 0 {
			domains = []string{""}
		}

		for _, domain := range domains {
			vsByDomain[domain] = append(vsByDomain[domain], vs)
		}
	}

	var conflictingDomains []string
	for domain, vsWithThisDomain := range vsByDomain {
		if len(vsWithThisDomain) > 1 {
			conflictingDomains = append(conflictingDomains, domain)
			for i, vs := range vsWithThisDomain {
				var conflictingVsNames []string
				for j, otherVs := range vsWithThisDomain {
					if i != j {
						conflictingVsNames = append(conflictingVsNames, otherVs.Metadata.Ref().Key())
					}
				}
				reports.AddError(vs, DomainInOtherVirtualServicesErr(domain, conflictingVsNames))
			}
		}
	}
	if len(conflictingDomains) > 0 {
		reports.AddError(gateway, GatewayHasConflictingVirtualServicesErr(conflictingDomains))
	}
}

func getVirtualServicesForGateway(gateway *v1.Gateway, virtualServices v1.VirtualServiceList) v1.VirtualServiceList {

	var virtualServicesForGateway v1.VirtualServiceList
	for _, vs := range virtualServices {
		if GatewayContainsVirtualService(gateway, vs) {
			virtualServicesForGateway = append(virtualServicesForGateway, vs)
		}
	}

	return virtualServicesForGateway
}

func GatewayContainsVirtualService(gateway *v1.Gateway, virtualService *v1.VirtualService) bool {
	httpGateway := gateway.GetHttpGateway()
	if httpGateway == nil {
		return false
	}

	if gateway.Ssl != hasSsl(virtualService) {
		return false
	}

	if len(httpGateway.VirtualServiceSelector) > 0 {
		// select virtual services by the label selector
		selector := labels.SelectorFromSet(httpGateway.VirtualServiceSelector)

		vsLabels := labels.Set(virtualService.Metadata.Labels)

		return virtualServiceNamespaceValidForGateway(gateway, virtualService) && selector.Matches(vsLabels)
	}
	// use individual refs to collect virtual services
	virtualServiceRefs := httpGateway.VirtualServices

	if len(virtualServiceRefs) == 0 {
		return virtualServiceNamespaceValidForGateway(gateway, virtualService)
	}

	vsRef := virtualService.Metadata.Ref()

	for _, ref := range virtualServiceRefs {
		if ref.Equal(vsRef) {
			return true
		}
	}

	return false
}

func virtualServiceNamespaceValidForGateway(gateway *v1.Gateway, virtualService *v1.VirtualService) bool {
	httpGateway := gateway.GetHttpGateway()
	if httpGateway == nil {
		return false
	}

	if len(httpGateway.VirtualServiceNamespaces) > 0 {
		for _, ns := range httpGateway.VirtualServiceNamespaces {
			if ns == "*" || virtualService.Metadata.Namespace == ns {
				return true
			}
		}
		return false
	}

	// by default, virtual services will be discovered in all namespaces
	return true
}

func hasSsl(vs *v1.VirtualService) bool {
	return vs.SslConfig != nil
}

func (t *HttpTranslator) desiredListenerForHttp(gateway *v1.Gateway, virtualServicesForGateway v1.VirtualServiceList, tables v1.RouteTableList, reports reporter.ResourceReports) *gloov1.Listener {
	var (
		virtualHosts []*gloov1.VirtualHost
		sslConfigs   []*gloov1.SslConfig
	)

	for _, virtualService := range virtualServicesForGateway.Sort() {
		if virtualService.VirtualHost == nil {
			virtualService.VirtualHost = &v1.VirtualHost{}
		}
		vh, err := t.virtualServiceToVirtualHost(virtualService, tables, reports)
		if err != nil {
			reports.AddError(virtualService, err)
			continue
		}
		virtualHosts = append(virtualHosts, vh)
		if virtualService.SslConfig != nil {
			sslConfigs = append(sslConfigs, virtualService.SslConfig)
		}
	}

	var httpPlugins *gloov1.HttpListenerOptions
	if httpGateway := gateway.GetHttpGateway(); httpGateway != nil {
		httpPlugins = httpGateway.Options
	}
	listener := makeListener(gateway)
	listener.ListenerType = &gloov1.Listener_HttpListener{
		HttpListener: &gloov1.HttpListener{
			VirtualHosts: virtualHosts,
			Options:      httpPlugins,
		},
	}
	listener.SslConfigurations = sslConfigs

	if err := appendSource(listener, gateway); err != nil {
		// should never happen
		reports.AddError(gateway, err)
	}

	return listener
}

func (t *HttpTranslator) virtualServiceToVirtualHost(vs *v1.VirtualService, tables v1.RouteTableList, reports reporter.ResourceReports) (*gloov1.VirtualHost, error) {
	converter := NewRouteConverter(NewRouteTableSelector(tables), NewRouteTableIndexer())
	routes, err := converter.ConvertVirtualService(vs, reports)
	if err != nil {
		// internal error, should never happen
		return nil, err
	}

	vh := &gloov1.VirtualHost{
		Name:    VirtualHostName(vs),
		Domains: vs.VirtualHost.Domains,
		Routes:  routes,
		Options: vs.VirtualHost.Options,
	}

	if t.WarnOnRouteShortCircuiting {
		validateRoutes(vs, vh, reports)
	}

	if err := appendSource(vh, vs); err != nil {
		// should never happen
		return nil, err
	}

	return vh, nil
}

func VirtualHostName(vs *v1.VirtualService) string {
	return fmt.Sprintf("%v.%v", vs.Metadata.Namespace, vs.Metadata.Name)
}

// this function is written with the assumption that the routes will not be modified afterwards,
// and are in their final sorted form
func validateRoutes(vs *v1.VirtualService, vh *gloov1.VirtualHost, reports reporter.ResourceReports) {
	validateAnyDuplicateMatchers(vs, vh, reports)
	validatePrefixHijacking(vs, vh, reports)
	validateRegexHijacking(vs, vh, reports)
}

func validateAnyDuplicateMatchers(vs *v1.VirtualService, vh *gloov1.VirtualHost, reports reporter.ResourceReports) {
	// warn on duplicate matchers
	seenMatchers := make(map[uint64]bool)
	for _, rt := range vh.Routes {
		for _, matcher := range rt.Matchers {
			hash := hashutils.MustHash(matcher)
			if _, ok := seenMatchers[hash]; ok == true {
				reports.AddWarning(vs, ConflictingMatcherErr(vh.GetName(), matcher).Error())
			} else {
				seenMatchers[hash] = true
			}
		}
	}
}

func validatePrefixHijacking(vs *v1.VirtualService, vh *gloov1.VirtualHost, reports reporter.ResourceReports) {
	// warn on early prefix matchers that short-circuit later routes

	var seenPrefixMatchers []*matchers.Matcher
	for _, rt := range vh.Routes {
		for _, matcher := range rt.Matchers {
			// make sure the current matcher doesn't match any previously defined prefix.
			// this code is written with the assumption that the routes are already in their final order;
			// we are trying to help users avoid misconfiguration and short-circuiting errors
			path := utils.PathAsString(matcher)
			for _, prefix := range seenPrefixMatchers {
				if strings.HasPrefix(path, prefix.GetPrefix()) && nonPathEarlyMatcherShortCircuitsLateMatcher(matcher, prefix) {
					reports.AddWarning(vs, UnorderedPrefixErr(vh.GetName(), prefix.GetPrefix(), matcher).Error())
				}
			}
			if matcher.GetPrefix() != "" {
				seenPrefixMatchers = append(seenPrefixMatchers, matcher)
			}
		}
	}
}

func validateRegexHijacking(vs *v1.VirtualService, vh *gloov1.VirtualHost, reports reporter.ResourceReports) {
	// warn on early regex matchers that short-circuit later routes

	var seenRegexMatchers []*matchers.Matcher
	for _, rt := range vh.Routes {
		for _, matcher := range rt.Matchers {
			if matcher.GetRegex() != "" {
				seenRegexMatchers = append(seenRegexMatchers, matcher)
			} else {
				// make sure the current matcher doesn't match any previously defined regex.
				// this code is written with the assumption that the routes are already in their final order;
				// we are trying to help users avoid misconfiguration and short-circuiting errors
				path := utils.PathAsString(matcher)
				for _, regex := range seenRegexMatchers {
					re := regexp.MustCompile(regex.GetRegex())
					foundIndex := re.FindStringIndex(path)
					if foundIndex != nil && nonPathEarlyMatcherShortCircuitsLateMatcher(matcher, regex) {
						reports.AddWarning(vs, UnorderedRegexErr(vh.GetName(), regex.GetRegex(), matcher).Error())
					}
				}
			}
		}
	}
}

// As future matcher APIs get added, this validation will need to be updated as well.
// If it gets too complex, consider modeling as a constraint satisfaction problem.

// This code is written with the assumption that header/query matchers for each header/query param shows up once
// If it shows up more than once, then we just use the latest condition. This may cause extra warnings that were
// unnecessary in rare cases.
func nonPathEarlyMatcherShortCircuitsLateMatcher(laterMatcher, earlierMatcher *matchers.Matcher) bool {

	// we play a trick here to validate the methods by writing them as header
	// matchers and just reusing the header matcher logic
	earlyMatcher := *earlierMatcher
	if len(earlyMatcher.Methods) > 0 {
		earlyMatcher.Headers = append(earlyMatcher.Headers, &matchers.HeaderMatcher{
			Name:  ":method",
			Value: fmt.Sprintf("(%s)", strings.Join(earlyMatcher.Methods, "|")),
			Regex: true,
		})
	}

	lateMatcher := *laterMatcher
	if len(lateMatcher.Methods) > 0 {
		lateMatcher.Headers = append(lateMatcher.Headers, &matchers.HeaderMatcher{
			Name:  ":method",
			Value: fmt.Sprintf("(%s)", strings.Join(lateMatcher.Methods, "|")),
			Regex: true,
		})
	}

	queryParamsShortCircuited := earlyQueryParametersShortCircuitedLaterOnes(lateMatcher, earlyMatcher)
	headersShortCircuited := earlyHeaderMatchersShortCircuitLaterOnes(lateMatcher, earlyMatcher)
	return queryParamsShortCircuited && headersShortCircuited
}

// returns true if the query parameter matcher conditions (or lack thereof) on the early matcher can short-circuit the
// query parameter matcher conditions of the latter. This can happen if every condition specified on the early matcher
// can also be satisfied by a condition on the same query parameter in a later matcher.
func earlyQueryParametersShortCircuitedLaterOnes(laterMatcher, earlyMatcher matchers.Matcher) bool {
	for _, earlyQpm := range earlyMatcher.QueryParameters {

		// for each early QPM, we see if there is a constraint on the later QPM that means we cannot satisfy
		// both at the same time. If we have an unsatisfiable constraint, then we know the earlier matcher cannot
		// short-circuit the later one.
		unsatisfiableConstraint := len(laterMatcher.QueryParameters) == 0

		for _, laterQpm := range laterMatcher.QueryParameters {
			if earlyQpm.Name == laterQpm.Name {
				// we found an overlapping condition

				// let's check if the early one is a subset of the later one
				if earlyQpm.Regex && !laterQpm.Regex {
					re := regexp.MustCompile(earlyQpm.Value)
					foundIndex := re.FindStringIndex(laterQpm.Value)
					if foundIndex == nil {
						// early regex doesn't capture the later matcher
						unsatisfiableConstraint = true
					}
				} else if !earlyQpm.Regex && !laterQpm.Regex {
					if earlyQpm.Value != laterQpm.Value {
						// early and late both have conditions on query parameter matcher
						unsatisfiableConstraint = true
					}
				}
			}
		}

		if unsatisfiableConstraint {
			// since both constraints can't be satisfied at the same time, we know that the
			// later route cannot be short-circuited by the earlier one
			return false
		}
	}

	// every single qpm matcher defined on the later matcher was short-circuited
	return true
}

// returns true if every header matcher specified on the later matcher is also specified on the earlier matcher,
// and the earlier header matcher doesn't have any extra header matchers for headers the later one lacks:
// thus, in terms of header matchers, the later header matcher is unreachable.
func earlyHeaderMatchersShortCircuitLaterOnes(laterMatcher, earlyMatcher matchers.Matcher) bool {
	earlyHeadersMap := map[string]*matchers.HeaderMatcher{}
	earlyHeadersSeen := map[string]bool{}
	for _, earlyHeader := range earlyMatcher.Headers {
		earlyHeadersMap[earlyHeader.Name] = earlyHeader
		earlyHeadersSeen[earlyHeader.Name] = false
	}

	laterHeadersMap := map[string]*matchers.HeaderMatcher{}
	for _, laterHeader := range laterMatcher.Headers {
		laterHeadersMap[laterHeader.Name] = laterHeader
	}

	for _, laterHeader := range laterHeadersMap {
		earlyHeader, ok := earlyHeadersMap[laterHeader.Name]
		if !ok {
			// later header matcher doesn't have an equivalent early one to short-circuit
			continue
		}
		earlyHeadersSeen[earlyHeader.Name] = true

		var match *bool
		if earlyHeader.Regex && !laterHeader.Regex {
			f := false
			match = &f
			re := regexp.MustCompile(earlyHeader.Value)
			foundIndex := re.FindStringIndex(laterHeader.Value)
			if foundIndex != nil {
				t := true
				match = &t
			}
		} else if !earlyHeader.Regex && laterHeader.Regex {
			// if the regex has format (A|B|C) ensure each condition can be reached
			// TODO(kdorosh) implement me
			continue
		} else if !earlyHeader.Regex && !laterHeader.Regex {
			f := false
			match = &f
			if earlyHeader.Value == laterHeader.Value {
				t := true
				match = &t
			}
		}
		if match != nil && earlyHeader.InvertMatch {
			// if we evaluated the header for match (non nil), then invert the match result
			tmp := *match
			swap := !tmp
			match = &swap
		}

		if match != nil && !*match {
			// early header matcher doesn't properly short-circuit the later one
			return false
		}
	}

	for _, seen := range earlyHeadersSeen {
		if seen == false {
			// early matcher had header condition more specific than the latter, doesn't short-circuit
			return false
		}
	}

	// every single header matcher defined on the later matcher was short-circuited
	return true
}
