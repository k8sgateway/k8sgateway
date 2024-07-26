package iosnapshot

import (
	"cmp"
	"encoding/json"
	"fmt"
	"slices"

	v1snap "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/gloosnapshot"

	crdv1 "github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd/solo.io/v1"
)

// formatResources sorts the resources and formats them into json output
func formatResources(resources []crdv1.Resource) ([]byte, error) {
	sortResources(resources)
	return formatOutput("json_compact", resources)
}

// formatOutput formats a generic object into the specified output format
func formatOutput(format string, genericOutput interface{}) ([]byte, error) {
	switch format {
	case "json":
		return json.MarshalIndent(genericOutput, "", "    ")
	case "", "json_compact":
		return json.Marshal(genericOutput)
	case "yaml":
		// There may be a case in the future, where yaml formatting is necessary
		// Since it is not required yet, we do not add support
		return nil, fmt.Errorf("%s format is not yet supported", format)
	default:
		return nil, fmt.Errorf("invalid format of %s", format)
	}
}

// sortResources sorts resources by gvk, namespace, and name
func sortResources(resources []crdv1.Resource) {
	slices.SortStableFunc(resources, func(a, b crdv1.Resource) int {
		return cmp.Or(
			cmp.Compare(a.APIVersion, b.APIVersion),
			cmp.Compare(a.Kind, b.Kind),
			cmp.Compare(a.GetNamespace(), b.GetNamespace()),
			cmp.Compare(a.GetName(), b.GetName()),
		)
	})
}

// apiSnapshotToGenericMap converts an ApiSnapshot into a generic map
// Since maps do not guarantee ordering, we do not attempt to sort these resources, as we do four []crdv1.Resource
func apiSnapshotToGenericMap(snap *v1snap.ApiSnapshot) (map[string]interface{}, error) {
	genericMap := map[string]interface{}{}

	if snap == nil {
		return genericMap, nil
	}

	jsn, err := json.Marshal(snap)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(jsn, &genericMap); err != nil {
		return nil, err
	}
	return genericMap, nil
}
