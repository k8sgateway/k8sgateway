package plugins

import (
	"bytes"
	"sort"
	"strings"

	envoy_config_listener_v3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	envoyhttp "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	"github.com/golang/protobuf/proto"
	"github.com/solo-io/gloo/projects/gloo/pkg/utils"
)

var (
	_ sort.Interface = new(StagedHttpFilterList)
	_ sort.Interface = new(StagedNetworkFilterList)
)

// WellKnownFilterStages are represented by an integer that reflects their relative ordering
type WellKnownFilterStage int

// The set of WellKnownFilterStages, whose order corresponds to the order used to sort filters
// If new well known filter stages are added, they should be inserted in a position corresponding to their order
const (
	FaultStage     WellKnownFilterStage = iota // Fault injection // First Filter Stage
	CorsStage                                  // Cors stage
	WafStage                                   // Web application firewall stage
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

// FilterStageComparison helps implement the sort.Interface Less function for use in other implementations of sort.Interface
// returns -1 if less than, 0 if equal, 1 if greater than
// It is not sufficient to return a Less bool because calling functions need to know if equal or greater when Less is false
func FilterStageComparison(a, b FilterStage) int {
	if a.RelativeTo < b.RelativeTo {
		return -1
	} else if a.RelativeTo > b.RelativeTo {
		return 1
	}
	if a.Weight < b.Weight {
		return -1
	} else if a.Weight > b.Weight {
		return 1
	}
	return 0
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

type StagedHttpFilter struct {
	HttpFilter *envoyhttp.HttpFilter
	Stage      FilterStage
}

type StagedHttpFilterList []StagedHttpFilter

func (s StagedHttpFilterList) Len() int {
	return len(s)
}

// filters by Relative Stage, Weighting, Name, Config Type-Url, Config Value, and (to ensure stability) index.
// The assumption is that if two filters are in the same stage, their order doesn't matter, and we
// just need to make sure it is stable.
func (s StagedHttpFilterList) Less(i, j int) bool {
	if compare := FilterStageComparison(s[i].Stage, s[j].Stage); compare != 0 {
		return compare < 0
	}

	if compare := strings.Compare(s[i].HttpFilter.GetName(), s[j].HttpFilter.GetName()); compare != 0 {
		return compare < 0
	}

	if compare := strings.Compare(s[i].HttpFilter.GetTypedConfig().GetTypeUrl(), s[j].HttpFilter.GetTypedConfig().GetTypeUrl()); compare != 0 {
		return compare < 0
	}

	if compare := bytes.Compare(s[i].HttpFilter.GetTypedConfig().GetValue(), s[j].HttpFilter.GetTypedConfig().GetValue()); compare != 0 {
		return compare < 0
	}

	// ensure stability
	return i < j
}

func (s StagedHttpFilterList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type StagedNetworkFilter struct {
	NetworkFilter *envoy_config_listener_v3.Filter
	Stage         FilterStage
}

type StagedNetworkFilterList []StagedNetworkFilter

func (s StagedNetworkFilterList) Len() int {
	return len(s)
}

// filters by Relative Stage, Weighting, Name, and (to ensure stability) index
func (s StagedNetworkFilterList) Less(i, j int) bool {
	switch FilterStageComparison(s[i].Stage, s[j].Stage) {
	case -1:
		return true
	case 1:
		return false
	}
	if s[i].NetworkFilter.GetName() < s[j].NetworkFilter.GetName() {
		return true
	}
	if s[i].NetworkFilter.GetName() > s[j].NetworkFilter.GetName() {
		return false
	}
	if s[i].NetworkFilter.String() < s[j].NetworkFilter.String() {
		return true
	}
	if s[i].NetworkFilter.String() > s[j].NetworkFilter.String() {
		return false
	}
	// ensure stability
	return i < j
}

func (s StagedNetworkFilterList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func NewStagedFilter(name string, stage FilterStage) StagedHttpFilter {
	s, _ := NewStagedFilterWithConfig(name, nil, stage)
	return s
}

func NewStagedFilterWithConfig(name string, config proto.Message, stage FilterStage) (StagedHttpFilter, error) {

	s := StagedHttpFilter{
		HttpFilter: &envoyhttp.HttpFilter{
			Name: name,
		},
		Stage: stage,
	}

	if config != nil {

		marshalledConf, err := utils.MessageToAny(config)
		if err != nil {
			// this should NEVER HAPPEN!
			return StagedHttpFilter{}, err
		}

		s.HttpFilter.ConfigType = &envoyhttp.HttpFilter_TypedConfig{
			TypedConfig: marshalledConf,
		}
	}

	return s, nil
}

func StagedFilterListContainsName(filters StagedHttpFilterList, filterName string) bool {
	for _, filter := range filters {
		if filter.HttpFilter.GetName() == filterName {
			return true
		}
	}

	return false
}
