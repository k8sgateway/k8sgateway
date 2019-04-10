package common

import (
	"github.com/ghodss/yaml"
	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/helpers"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/go-utils/errors"
	"github.com/solo-io/go-utils/protoutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
)

func CreateAndPrintObject(yml []byte, outputType string) error {
	resource, err := resourceFromYaml(yml)
	if err != nil {
		return errors.Wrapf(err, "parsing resource from yaml")
	}
	switch res := resource.(type) {
	case *gloov1.Upstream:
		us, err := helpers.MustUpstreamClient().Write(res, clients.WriteOpts{})
		if err != nil {
			return errors.Wrapf(err, "saving Upstream to storage")
		}
		helpers.PrintUpstreams(gloov1.UpstreamList{us}, outputType)
	case *v1.VirtualService:
		vs, err := helpers.MustVirtualServiceClient().Write(res, clients.WriteOpts{})
		if err != nil {
			return errors.Wrapf(err, "saving VirtualService to storage")
		}
		helpers.PrintVirtualServices(v1.VirtualServiceList{vs}, outputType)
	default:
		return errors.Errorf("cli error: unimplemented resource type %v", resource)
	}
	return nil
}

func resourceFromYaml(yml []byte) (resources.Resource, error) {
	var untypedObj map[string]interface{}
	if err := yaml.Unmarshal(yml, &untypedObj); err != nil {
		return nil, err
	}
	// TODO ilackarms: find a better way. right now we rely on a required field being present in the yaml
	switch {
	case untypedObj["virtualHost"] != nil:
		var vs v1.VirtualService
		if err := protoutils.UnmarshalYaml(yml, &vs); err != nil {
			return nil, err
		}
		return &vs, nil
	case untypedObj["upstreamSpec"] != nil:
		var us gloov1.Upstream
		if err := protoutils.UnmarshalYaml(yml, &us); err != nil {
			return nil, err
		}
		return &us, nil
	}
	return nil, errors.Errorf("unknown object: %v", untypedObj)
}
