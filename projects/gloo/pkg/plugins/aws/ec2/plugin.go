package ec2

import (
	"reflect"

	envoyapi "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/solo-io/gloo/projects/gloo/pkg/xds"

	"github.com/pkg/errors"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/discovery"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
)

/*
Steps:
- User creates an EC2 upstream
  - describes the instances that should be made into Endpoints
- Discovery finds all instances that match the description with DescribeInstances
- Gloo plugin creates an endpoint for each instance
*/

type plugin struct {
	secretClient v1.SecretClient

	// pre-initialization only
	secretFactory factory.ResourceClientFactory
}

// checks to ensure interfaces are implemented
var _ plugins.Plugin = new(plugin)
var _ plugins.UpstreamPlugin = new(plugin)
var _ discovery.DiscoveryPlugin = new(plugin)

func NewPlugin(secretFactory factory.ResourceClientFactory) *plugin {
	return &plugin{secretFactory: secretFactory}
}

func (p *plugin) Init(params plugins.InitParams) error {
	var err error
	p.secretClient, err = v1.NewSecretClient(p.secretFactory)
	if err != nil {
		return err
	}
	if err := p.secretClient.Register(); err != nil {
		return err
	}
	return nil
}

// we do not need to update any fields, just check that the input is valid
func (p *plugin) UpdateUpstream(original, desired *v1.Upstream) (bool, error) {
	originalSpec, ok := original.UpstreamSpec.UpstreamType.(*v1.UpstreamSpec_AwsEc2)
	if !ok {
		return false, errors.Errorf("internal error: expected *v1.UpstreamSpec_AwsEc2, got %v", reflect.TypeOf(original.UpstreamSpec.UpstreamType).Name())
	}
	desiredSpec, ok := desired.UpstreamSpec.UpstreamType.(*v1.UpstreamSpec_AwsEc2)
	if !ok {
		return false, errors.Errorf("internal error: expected *v1.UpstreamSpec_AwsEc2, got %v", reflect.TypeOf(original.UpstreamSpec.UpstreamType).Name())
	}
	if !originalSpec.Equal(desiredSpec) {
		return false, errors.New("expected no difference between *v1.UpstreamSpec_AwsEc2 upstreams")
	}
	return false, nil
}

func (p *plugin) ProcessUpstream(params plugins.Params, in *v1.Upstream, out *envoyapi.Cluster) error {
	_, ok := in.UpstreamSpec.UpstreamType.(*v1.UpstreamSpec_AwsEc2)
	if !ok {
		return nil
	}

	// configure the cluster to use EDS:ADS and call it a day
	xds.SetEdsOnCluster(out)
	return nil
}
