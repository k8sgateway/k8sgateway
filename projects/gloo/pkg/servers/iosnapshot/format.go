package iosnapshot

import (
	"encoding/json"
	"fmt"

	v1snap "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/gloosnapshot"
)

func getGenericMaps(snapshot map[string]*v1snap.ApiSnapshot) (map[string]interface{}, error) {
	genericMaps := map[string]interface{}{}
	for id, obj := range snapshot {
		jsn, err := obj.MarshalJSON()
		if err != nil {
			return nil, err
		}
		genericMap := map[string]interface{}{}
		if err := json.Unmarshal(jsn, &genericMap); err != nil {
			return nil, err
		}
		genericMaps[id] = genericMap
	}
	return genericMaps, nil
}

func formatMap(format string, genericMaps map[string]interface{}) ([]byte, error) {
	switch format {
	case "json":
		return json.MarshalIndent(genericMaps, "", "    ")
	case "", "json_compact":
		return json.Marshal(genericMaps)
	case "yaml":
		// There may be a case in the future, where yaml formatting is necessary
		// Since it is not required yet, we do not add support
		return nil, fmt.Errorf("%s format is not yet supported", format)
	default:
		return nil, fmt.Errorf("invalid format of %s", format)
	}

}
