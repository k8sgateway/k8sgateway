package pluginutils

import (
	"context"
	"reflect"

	envoyroute "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"

	"github.com/envoyproxy/go-control-plane/pkg/util"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	"github.com/pkg/errors"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/solo-kit/pkg/utils/contextutils"
)

func SetRoutePerFilterConfig(out *envoyroute.Route, filterName string, protoext proto.Message) error {
	if out.PerFilterConfig == nil {
		out.PerFilterConfig = make(map[string]*types.Struct)
	}
	return setConfig(out.PerFilterConfig, filterName, protoext)
}
func SetVhostPerFilterConfig(out *envoyroute.VirtualHost, filterName string, protoext proto.Message) error {
	if out.PerFilterConfig == nil {
		out.PerFilterConfig = make(map[string]*types.Struct)
	}
	return setConfig(out.PerFilterConfig, filterName, protoext)
}
func SetWeightedClusterPerFilterConfig(out *envoyroute.WeightedCluster_ClusterWeight, filterName string, protoext proto.Message) error {
	if out.PerFilterConfig == nil {
		out.PerFilterConfig = make(map[string]*types.Struct)
	}
	return setConfig(out.PerFilterConfig, filterName, protoext)
}

// Return Per-Filter config for destinations, we put them on the Route (single dest) or WeightedCluster (multi dest)
type PerFilterConfigFunc func(spec *v1.Destination) (proto.Message, error)

// call this from
func MarkPerFilterConfig(ctx context.Context, in *v1.Route, out *envoyroute.Route, filterName string, perFilterConfig PerFilterConfigFunc) error {
	inAction, outAction, err := getRouteActions(in, out)
	if err != nil {
		return err
	}

	switch dest := inAction.Destination.(type) {
	case *v1.RouteAction_Multi:
		multiClusterSpecifier, ok := outAction.ClusterSpecifier.(*envoyroute.RouteAction_WeightedClusters)
		if !ok {
			return errors.Errorf("input destination Multi but output destination was not")
		}
		return configureMultiDest(dest.Multi, multiClusterSpecifier.WeightedClusters, filterName, perFilterConfig)
	case *v1.RouteAction_Single:
		if out.PerFilterConfig == nil {
			out.PerFilterConfig = make(map[string]*types.Struct)
		}
		return configureSingleDest(dest.Single, out.PerFilterConfig, filterName, perFilterConfig)
	}

	err = errors.Errorf("unexpected destination type %v", reflect.TypeOf(inAction.Destination).Name())
	logger := contextutils.LoggerFrom(ctx)
	logger.DPanic("error: %v", err)
	return err
}

func configureMultiDest(in *v1.MultiDestination, out *envoyroute.WeightedCluster, filterName string, perFilterConfig PerFilterConfigFunc) error {
	if len(in.Destinations) != len(out.Clusters) {
		return errors.Errorf("number of input destinations did not match number of destination weighted clusters")
	}
	for i := range in.Destinations {
		if out.Clusters[i].PerFilterConfig == nil {
			out.Clusters[i].PerFilterConfig = make(map[string]*types.Struct)
		}
		err := configureSingleDest(in.Destinations[i].Destination, out.Clusters[i].PerFilterConfig, filterName, perFilterConfig)
		if err != nil {
			return err
		}
	}

	return nil
}

func configureSingleDest(in *v1.Destination, out map[string]*types.Struct, filterName string, perFilterConfig PerFilterConfigFunc) error {
	config, err := perFilterConfig(in)
	if err != nil {
		return err
	}
	return setConfig(out, filterName, config)
}

func setConfig(out map[string]*types.Struct, filterName string, config proto.Message) error {
	if config == nil {
		return nil
	}
	configStruct, err := util.MessageToStruct(config)
	if err != nil {
		return err
	}
	out[filterName] = configStruct
	return nil
}
