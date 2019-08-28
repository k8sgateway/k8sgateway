// Code generated by solo-kit. DO NOT EDIT.

package extauth

import (
	"context"
	"errors"
	"fmt"

	discovery "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/solo-io/solo-kit/pkg/api/v1/control-plane/cache"
	"github.com/solo-io/solo-kit/pkg/api/v1/control-plane/client"
	"github.com/solo-io/solo-kit/pkg/api/v1/control-plane/server"
)

// Type Definitions:

const ExtAuthConfigType = cache.TypePrefix + "/extauth.plugins.gloo.solo.io.ExtAuthConfig"

/* Defined a resource - to be used by snapshot */
type ExtAuthConfigXdsResourceWrapper struct {
	// TODO(yuval-k): This is public for mitchellh hashstructure to work properly. consider better alternatives.
	Resource *ExtAuthConfig
}

// Make sure the Resource interface is implemented
var _ cache.Resource = &ExtAuthConfigXdsResourceWrapper{}

func NewExtAuthConfigXdsResourceWrapper(resourceProto *ExtAuthConfig) *ExtAuthConfigXdsResourceWrapper {
	return &ExtAuthConfigXdsResourceWrapper{
		Resource: resourceProto,
	}
}

func (e *ExtAuthConfigXdsResourceWrapper) Self() cache.XdsResourceReference {
	return cache.XdsResourceReference{Name: e.Resource.Vhost, Type: ExtAuthConfigType}
}

func (e *ExtAuthConfigXdsResourceWrapper) ResourceProto() cache.ResourceProto {
	return e.Resource
}
func (e *ExtAuthConfigXdsResourceWrapper) References() []cache.XdsResourceReference {
	return nil
}

// Define a type record. This is used by the generic client library.
var ExtAuthConfigTypeRecord = client.NewTypeRecord(
	ExtAuthConfigType,

	// Return an empty message, that can be used to deserialize bytes into it.
	func() cache.ResourceProto { return &ExtAuthConfig{} },

	// Covert the message to a resource suitable for use for protobuf's Any.
	func(r cache.ResourceProto) cache.Resource {
		return &ExtAuthConfigXdsResourceWrapper{Resource: r.(*ExtAuthConfig)}
	},
)

// Server Implementation:

// Wrap the generic server and implement the type sepcific methods:
type extAuthDiscoveryServiceServer struct {
	server.Server
}

func NewExtAuthDiscoveryServiceServer(genericServer server.Server) ExtAuthDiscoveryServiceServer {
	return &extAuthDiscoveryServiceServer{Server: genericServer}
}

func (s *extAuthDiscoveryServiceServer) StreamExtAuthConfig(stream ExtAuthDiscoveryService_StreamExtAuthConfigServer) error {
	return s.Server.Stream(stream, ExtAuthConfigType)
}

func (s *extAuthDiscoveryServiceServer) FetchExtAuthConfig(ctx context.Context, req *discovery.DiscoveryRequest) (*discovery.DiscoveryResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.Unavailable, "empty request")
	}
	req.TypeUrl = ExtAuthConfigType
	return s.Server.Fetch(ctx, req)
}

func (s *extAuthDiscoveryServiceServer) DeltaExtAuthConfig(_ ExtAuthDiscoveryService_DeltaExtAuthConfigServer) error {
	return errors.New("not implemented")
}

// Client Implementation: Generate a strongly typed client over the generic client

// The apply functions receives resources and returns an error if they were applied correctly.
// In theory the configuration can become valid in the future (i.e. eventually consistent), but I don't think we need to worry about that now
// As our current use cases only have one configuration resource, so no interactions are expected.
type ApplyExtAuthConfig func(version string, resources []*ExtAuthConfig) error

// Convert the strongly typed apply to a generic apply.
func applyExtAuthConfig(typedApply ApplyExtAuthConfig) func(cache.Resources) error {
	return func(resources cache.Resources) error {

		var configs []*ExtAuthConfig
		for _, r := range resources.Items {
			if proto, ok := r.ResourceProto().(*ExtAuthConfig); !ok {
				return fmt.Errorf("resource %s of type %s incorrect", r.Self().Name, r.Self().Type)
			} else {
				configs = append(configs, proto)
			}
		}

		return typedApply(resources.Version, configs)
	}
}

func NewExtAuthConfigClient(nodeinfo *core.Node, typedApply ApplyExtAuthConfig) client.Client {
	return client.NewClient(nodeinfo, ExtAuthConfigTypeRecord, applyExtAuthConfig(typedApply))
}
