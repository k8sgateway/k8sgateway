package consul

import (
	"strings"

	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/discovery"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins/utils"
	"github.com/solo-io/gloo/projects/gloo/pkg/upstreams/consul"
	"github.com/solo-io/go-utils/errors"
	"github.com/solo-io/go-utils/errutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
)

var (
	InvalidSpecTypeError = func(us *v1.Upstream, name string) error {
		return errors.Errorf("internal error: invalid %s spec, "+
			"expected *v1.Upstream_Consul, got  %T", name, us)
	}
)

func (p *plugin) DiscoverUpstreams(_ []string, writeNamespace string, opts clients.WatchOpts, discOpts discovery.Opts) (chan v1.UpstreamList, chan error, error) {
	upstreams, errs, err := consul.NewConsulUpstreamClient(p.client).Watch("", opts)
	if err != nil {
		return nil, nil, err
	}

	realUpstreams := make(chan v1.UpstreamList)

	// need to do this because interface requires a bidirectional channel
	ourErrs := make(chan error)
	go errutils.AggregateErrs(opts.Ctx, ourErrs, errs, "consul-uds")

	// strip fake name prefix generated by upstream client
	go func() {
		defer close(realUpstreams)
		for {
			select {
			case upstreamList, ok := <-upstreams:
				if !ok {
					return
				}
				select {
				case <-opts.Ctx.Done():
					return
				case realUpstreams <- setRealName(upstreamList, writeNamespace):
				}
			case <-opts.Ctx.Done():
				return
			}
		}
	}()

	return realUpstreams, ourErrs, nil
}

// set namespace and name to be valid for writing to storage
func setRealName(list v1.UpstreamList, writeNamespace string) v1.UpstreamList {
	list.Each(func(element *v1.Upstream) {
		element.Metadata.Name = strings.TrimPrefix(element.Metadata.Name, consul.UpstreamNamePrefix)
		element.Metadata.Namespace = writeNamespace
	})
	return list
}
func (p *plugin) UpdateUpstream(original, desired *v1.Upstream) (bool, error) {
	return UpdateUpstream(original, desired)
}

func UpdateUpstream(original, desired *v1.Upstream) (bool, error) {
	originalSpec, ok := original.UpstreamType.(*v1.Upstream_Consul)
	if !ok {
		return false, InvalidSpecTypeError(original, "original")
	}
	desiredSpec, ok := desired.UpstreamType.(*v1.Upstream_Consul)
	if !ok {
		return false, InvalidSpecTypeError(desired, "desired")
	}

	// copy service spec, we don't want to overwrite that
	desiredSpec.Consul.ServiceSpec = originalSpec.Consul.ServiceSpec

	utils.UpdateUpstream(original, desired)

	if originalSpec.Equal(desiredSpec) {
		return false, nil
	}

	return true, nil
}
