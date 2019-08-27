package plugins

import (
	"context"
	"sort"

	envoyapi "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	envoylistener "github.com/envoyproxy/go-control-plane/envoy/api/v2/listener"
	envoyroute "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	envoyhttp "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/http_connection_manager/v2"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
)

type InitParams struct {
	Ctx                context.Context
	ExtensionsSettings *v1.Extensions
	Settings           *v1.Settings
}

type Plugin interface {
	Init(params InitParams) error
}

type Params struct {
	Ctx      context.Context
	Snapshot *v1.ApiSnapshot
}

type VirtualHostParams struct {
	Params
	Proxy    *v1.Proxy
	Listener *v1.Listener
}

type RouteParams struct {
	VirtualHostParams
	VirtualHost *v1.VirtualHost
}

/*
	Upstream Plugins
*/

type UpstreamPlugin interface {
	Plugin
	ProcessUpstream(params Params, in *v1.Upstream, out *envoyapi.Cluster) error
}

/*
	Routing Plugins
*/

type RoutePlugin interface {
	Plugin
	ProcessRoute(params RouteParams, in *v1.Route, out *envoyroute.Route) error
}

// note: any route action plugin can be implemented as a route plugin
// suggestion: if your plugin requires configuration from a RoutePlugin field, implement the RoutePlugin interface
type RouteActionPlugin interface {
	Plugin
	ProcessRouteAction(params RouteParams, inAction *v1.RouteAction, out *envoyroute.RouteAction) error
}

type WeightedDestinationPlugin interface {
	Plugin
	ProcessWeightedDestination(params RouteParams, in *v1.WeightedDestination, out *envoyroute.WeightedCluster_ClusterWeight) error
}

/*
	Listener Plugins
*/

type ListenerPlugin interface {
	Plugin
	ProcessListener(params Params, in *v1.Listener, out *envoyapi.Listener) error
}

type ListenerFilterPlugin interface {
	Plugin
	ProcessListenerFilter(params Params, in *v1.Listener) ([]StagedListenerFilter, error)
}

type StagedListenerFilter struct {
	ListenerFilter envoylistener.Filter
	Stage          FilterStage
}

type StagedListenerFilterList []StagedListenerFilter

func (s StagedListenerFilterList) Len() int {
	return len(s)
}

// filters by Relative Stage, Weighting, Name, and (to ensure stability) index
func (s StagedListenerFilterList) Less(i, j int) bool {
	if FilterStageLess(s[i].Stage, s[j].Stage) {
		return true
	}
	if s[i].ListenerFilter.Name < s[j].ListenerFilter.Name {
		return true
	}
	// ensure stability
	return i < j
}

func (s StagedListenerFilterList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Currently only supported for TCP listeners, plan to change this in the future
type ListenerFilterChainPlugin interface {
	Plugin
	ProcessListenerFilterChain(params Params, in *v1.Listener) ([]envoylistener.FilterChain, error)
}

type HttpFilterPlugin interface {
	Plugin
	HttpFilters(params Params, listener *v1.HttpListener) ([]StagedHttpFilter, error)
}

type VirtualHostPlugin interface {
	Plugin
	ProcessVirtualHost(params VirtualHostParams, in *v1.VirtualHost, out *envoyroute.VirtualHost) error
}

type StagedHttpFilter struct {
	HttpFilter *envoyhttp.HttpFilter
	Stage      FilterStage
}

type StagedHttpFilterList []StagedHttpFilter

func (s StagedHttpFilterList) Len() int {
	return len(s)
}

// filters by Relative Stage, Weighting, Name, and (to ensure stability) index
func (s StagedHttpFilterList) Less(i, j int) bool {
	if FilterStageLess(s[i].Stage, s[j].Stage) {
		return true
	}
	if s[i].HttpFilter.Name < s[j].HttpFilter.Name {
		return true
	}
	// ensure stability
	return i < j
}

func (s StagedHttpFilterList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

var _ sort.Interface = StagedHttpFilterList{}

// WellKnownFilterStages are represented by an integer that reflects their relative ordering
type WellKnownFilterStage int

// If new well known filter stages are added, they should be inserted in a position corresponding to their order
const (
	FaultStage     WellKnownFilterStage = iota // Fault injection // First Filter Stage
	CorsStage                                  // Cors stage
	AuthNStage                                 // Authentication stage
	AuthZStage                                 // Authorization stage
	RateLimitStage                             // Rate limiting stage
	AcceptedStage                              // Request passed all the checks and will be forwarded upstream
	OutAuthStage                               // Add auth for the upstream (i.e. aws λ)
	RouteStage                                 // Request is going to upstream // Last Filter Stage
)

type FilterStage struct {
	RelativeTo WellKnownFilterStage
	Weight     int
}

// FilterStageLess implements the sort.Interface Less function for use in other implementations of sort.Interface
func FilterStageLess(a, b FilterStage) bool {
	if a.RelativeTo < b.RelativeTo {
		return true
	}
	if a.Weight < b.Weight {
		return true
	}
	return false
}

func BeforeStage(wellKnown WellKnownFilterStage) FilterStage {
	return RelativeToStage(wellKnown, -1)
}
func DuringStage(wellKnown WellKnownFilterStage) FilterStage {
	return RelativeToStage(wellKnown, 0)
}
func AfterStage(wellKnown WellKnownFilterStage) FilterStage {
	return RelativeToStage(wellKnown, 1)
}
func RelativeToStage(wellKnown WellKnownFilterStage, weight int) FilterStage {
	return FilterStage{
		RelativeTo: wellKnown,
		Weight:     weight,
	}
}

// The following FilterStages are preserved for backwards compatibility. They will be removed and should not be used
// going forward.
var (
	// Deprecated
	FaultFilter = DuringStage(FaultStage)
	// Deprecated
	PreInAuth = BeforeStage(AuthNStage)
	// Deprecated
	InAuth = DuringStage(AuthNStage)
	// Deprecated
	PostInAuth = AfterStage(AuthNStage)
	// Deprecated
	PreOutAuth = BeforeStage(OutAuthStage)
)

/*
	Generation plugins
*/
type ClusterGeneratorPlugin interface {
	Plugin
	GeneratedClusters(params Params) ([]*envoyapi.Cluster, error)
}
