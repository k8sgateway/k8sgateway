package serviceconverter

import (
	"reflect"

	errors "github.com/rotisserie/eris"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/solo-kit/pkg/utils/protoutils"
	kubev1 "k8s.io/api/core/v1"
)

func init() {
	DefaultServiceConverters = append(DefaultServiceConverters, &GeneralServiceConverter{})
}

const GlooAnnotationPrefix = "gloo.solo.io/UpstreamConfig"

// The GeneralServiceConverter does not write config to these fields, because we expect them to be written to elsewhere.
var ExcludedFields = map[string]bool{
	"NamespacedStatuses": true,
	"Metadata":           true,
	"DiscoveryMetadata":  true,
	"UpstreamType":       true,
}

type GeneralServiceConverter struct{}

func (s *GeneralServiceConverter) ConvertService(svc *kubev1.Service, port kubev1.ServicePort, us *v1.Upstream) error {
	upstreamConfigJson, ok := svc.Annotations[GlooAnnotationPrefix]
	if !ok {
		return nil
	}

	var spec v1.Upstream
	if err := protoutils.UnmarshalResource([]byte(upstreamConfigJson), &spec); err != nil {
		return err
	}

	// iterate over fields in upstream spec
	specType := reflect.TypeOf(spec)
	numFields := specType.NumField()
	for i := 0; i < numFields; i++ {
		field := specType.Field(i)
		// if field is exported and not explicitly excluded, consider setting it on the upstream
		if field.PkgPath == "" && !ExcludedFields[field.Name] {
			fieldValue, err := getAttr(&spec, field.Name)
			if err != nil {
				return err
			}

			currentValue, err := getAttr(us, field.Name)
			if err != nil {
				return err
			}

			if fieldValue.IsValid() && currentValue.CanSet() && currentValue.IsZero() {
				currentValue.Set(fieldValue)
			}
		}
	}

	return nil
}

func getAttr(obj interface{}, fieldName string) (reflect.Value, error) {
	pointToStruct := reflect.ValueOf(obj) // addressable
	curStruct := pointToStruct.Elem()
	if curStruct.Kind() != reflect.Struct {
		return reflect.ValueOf(nil), errors.Errorf("Error: not struct")
	}
	curField := curStruct.FieldByName(fieldName) // type: reflect.Value
	if !curField.IsValid() {
		return reflect.ValueOf(nil), errors.Errorf("Error: not found:" + fieldName)
	}
	return curField, nil
}
