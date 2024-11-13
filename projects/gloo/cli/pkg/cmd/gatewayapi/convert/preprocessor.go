package convert

import (
	"github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/core/matchers"
)

type Preprocessor struct{}

func NewPreprocessor() *Preprocessor {
	return &Preprocessor{}
}

func (p *Preprocessor) Preprocess(input *GlooEdgeInput) {
	for _, rt := range input.RouteTables {
		var newRoutes []*v1.Route
		for _, route := range rt.Spec.Routes {
			editedRoute := generateRoutesForMethodMatchers(route)
			newRoutes = append(newRoutes, editedRoute)
		}
		rt.Spec.Routes = newRoutes
	}
}

func generateRoutesForMethodMatchers(route *v1.Route) *v1.Route {

	var newMatchers []*matchers.Matcher
	for _, m := range route.Matchers {
		if len(m.Methods) > 1 {
			// for each method we need to split out the matchers
			for _, method := range m.Methods {
				newMatcher := &matchers.Matcher{
					PathSpecifier:   m.PathSpecifier,
					CaseSensitive:   m.CaseSensitive,
					Headers:         m.Headers,
					QueryParameters: m.QueryParameters,
					Methods:         []string{method},
				}
				newMatchers = append(newMatchers, newMatcher)
			}
		} else {
			//it only has one so we just add it
			newMatchers = append(newMatchers, m)
		}
	}
	route.Matchers = newMatchers

	return route
}
