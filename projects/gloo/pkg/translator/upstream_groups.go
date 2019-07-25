package translator

import (
	"github.com/pkg/errors"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	usconversions "github.com/solo-io/gloo/projects/gloo/pkg/upstreams"
	"github.com/solo-io/solo-kit/pkg/api/v1/reporter"
)

func (t *translator) verifyUpstreamGroups(params plugins.Params, resourceErrs reporter.ResourceErrors) {

	upstreams := params.Snapshot.Upstreams
	upstreamGroups := params.Snapshot.UpstreamGroups

	for _, ug := range upstreamGroups {
		for i, dest := range ug.Destinations {
			if dest.Destination == nil {
				resourceErrs.AddError(ug, errors.Errorf("destination # %d: destination is nil", i+1))
				continue
			}

			upRef, err := usconversions.DestinationToUpstreamRef(dest.Destination)
			if err != nil {
				resourceErrs.AddError(ug, err)
				continue
			}

			if _, err := upstreams.Find(upRef.Namespace, upRef.Name); err != nil {
				resourceErrs.AddError(ug, errors.Wrapf(err, "destination # %d: upstream not found", i+1))
				continue
			}
		}

	}

}
