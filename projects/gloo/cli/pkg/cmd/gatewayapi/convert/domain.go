package convert

import (
	"bytes"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"strings"

	"github.com/itchyny/json2yaml"
	gatewaykube "github.com/solo-io/gloo/projects/gateway/pkg/api/v1/kube/apis/gateway.solo.io/v1"
	glookube "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/kube/apis/gloo.solo.io/v1"
	v1 "github.com/solo-io/solo-apis/pkg/api/enterprise.gloo.solo.io/v1"
	"github.com/spf13/pflag"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	gwv1 "sigs.k8s.io/gateway-api/apis/v1"
)

type Options struct {
	*options.Options

	InputFile                string
	Directory                string
	Overwrite                bool
	OverwriteSuffix          string
	Stats                    bool
	GCPRegex                 string
	RemoveGCPAUthConfig      bool
	RouteOptionStategicMerge bool
}

func (o *Options) addToFlags(flags *pflag.FlagSet) {
	flags.StringVar(&o.InputFile, "input-file", "", "File to convert")
	flags.BoolVar(&o.Overwrite, "overwrite", false, "Overwrite the existing files with the changes")
	flags.StringVar(&o.OverwriteSuffix, "suffix", "", "When writing to files add a suffix (to do side by side)")
	flags.BoolVar(&o.Stats, "stats", false, "Print stats about the conversion")
	flags.StringVar(&o.Directory, "dir", "", "Directory to read yaml/yml files")
}

type GlooEdgeInput struct {
	FileName           string
	YamlObjects        []string
	RouteTables        []*gatewaykube.RouteTable
	RouteOptions       []*gatewaykube.RouteOption
	VirtualHostOptions []*gatewaykube.VirtualHostOption
	Upstreams          []*glookube.Upstream
	VirtualServices    []*gatewaykube.VirtualService
	// Gateways           []*gatewaykube.Gateway // TODO do we need these?
	AuthConfigs []*v1.AuthConfig
}

type DelegateParentReference struct {
	Labels          map[string]string
	ParentName      string
	ParentNamespace string
}

type GatewayAPIOutput struct {
	FileName             string
	YamlObjects          []string
	DelegationReferences []*DelegateParentReference
	HTTPRoutes           []*gwv1.HTTPRoute
	RouteOptions         []*gatewaykube.RouteOption
	VirtualHostOptions   []*gatewaykube.VirtualHostOption
	Upstreams            []*glookube.Upstream
	AuthConfigs          []*v1.AuthConfig
	// Gateways             []*gwv1.Gateway
}

func (g *GatewayAPIOutput) HasItems() bool {

	if len(g.HTTPRoutes) > 0 {
		return true
	}
	if len(g.RouteOptions) > 0 {
		return true
	}
	if len(g.VirtualHostOptions) > 0 {
		return true
	}
	if len(g.Upstreams) > 0 {
		return true
	}
	if len(g.AuthConfigs) > 0 {
		return true
	}
	// if there are only yaml objects then skip because we didnt change anything in the file

	return false
}

func (g *GatewayAPIOutput) ToString() (string, error) {
	output := ""

	for _, y := range g.YamlObjects {
		output += "\n---\n" + y + "\n"
	}

	for _, obj := range g.Upstreams {
		o, err := runtime.Encode(codecs.LegacyCodec(corev1.SchemeGroupVersion, gwv1.SchemeGroupVersion, gatewaykube.SchemeGroupVersion, glookube.SchemeGroupVersion), obj)
		if err != nil {
			return "", err
		}

		var yaml strings.Builder
		if err := json2yaml.Convert(&yaml, bytes.NewReader(o)); err != nil {
			return "", err
		}

		output += "\n---\n" + yaml.String()
	}

	for _, obj := range g.HTTPRoutes {
		o, err := runtime.Encode(codecs.LegacyCodec(corev1.SchemeGroupVersion, gwv1.SchemeGroupVersion, gatewaykube.SchemeGroupVersion, glookube.SchemeGroupVersion), obj)
		if err != nil {
			return "", err
		}

		var yaml strings.Builder
		if err := json2yaml.Convert(&yaml, bytes.NewReader(o)); err != nil {
			return "", err
		}

		output += "\n---\n" + yaml.String()
	}
	for _, op := range g.RouteOptions {
		marshaller := YamlMarshaller{}
		yaml, err := marshaller.ToYaml(&op)
		if err != nil {
			return "", err
		}

		output += "\n---\n" + string(yaml)
	}
	for _, op := range g.AuthConfigs {
		marshaller := YamlMarshaller{}
		yaml, err := marshaller.ToYaml(&op)
		if err != nil {
			return "", err
		}

		output += "\n---\n" + string(yaml)
	}

	for _, op := range g.VirtualHostOptions {
		marshaller := YamlMarshaller{}
		yaml, err := marshaller.ToYaml(&op)
		if err != nil {
			return "", err
		}

		output += "\n---\n" + string(yaml)
	}

	// need to remove a few values
	//  creationTimestamp: null
	// status: {}
	// status:
	// parents: null
	output = strings.ReplaceAll(output, "  creationTimestamp: null\n", "")
	output = strings.ReplaceAll(output, "status:\n", "")
	output = strings.ReplaceAll(output, "parents: null\n", "")
	output = strings.ReplaceAll(output, "status: {}\n", "")
	output = strings.ReplaceAll(output, "\n\n\n", "\n")
	output = strings.ReplaceAll(output, "\n\n", "\n")
	output = strings.ReplaceAll(output, "spec: {}\n", "")

	// TODO remove leading and trailing ---
	// log.Printf("%s", output)
	return output, nil
}

func (g *GlooEdgeInput) ToString() (string, error) {
	output := ""

	for _, y := range g.YamlObjects {
		output += "\n---\n" + y + "\n"
	}

	for _, obj := range g.Upstreams {
		o, err := runtime.Encode(codecs.LegacyCodec(corev1.SchemeGroupVersion, gwv1.SchemeGroupVersion, gatewaykube.SchemeGroupVersion, glookube.SchemeGroupVersion), obj)
		if err != nil {
			return "", err
		}

		var yaml strings.Builder
		if err := json2yaml.Convert(&yaml, bytes.NewReader(o)); err != nil {
			return "", err
		}

		output += "\n---\n" + yaml.String()
	}

	for _, obj := range g.RouteTables {
		o, err := runtime.Encode(codecs.LegacyCodec(corev1.SchemeGroupVersion, gwv1.SchemeGroupVersion, gatewaykube.SchemeGroupVersion, glookube.SchemeGroupVersion), obj)
		if err != nil {
			return "", err
		}

		var yaml strings.Builder
		if err := json2yaml.Convert(&yaml, bytes.NewReader(o)); err != nil {
			return "", err
		}

		output += "\n---\n" + yaml.String()
	}
	for _, obj := range g.VirtualServices {
		o, err := runtime.Encode(codecs.LegacyCodec(corev1.SchemeGroupVersion, gwv1.SchemeGroupVersion, gatewaykube.SchemeGroupVersion, glookube.SchemeGroupVersion), obj)
		if err != nil {
			return "", err
		}

		var yaml strings.Builder
		if err := json2yaml.Convert(&yaml, bytes.NewReader(o)); err != nil {
			return "", err
		}

		output += "\n---\n" + yaml.String()
	}
	for _, op := range g.RouteOptions {
		marshaller := YamlMarshaller{}
		yaml, err := marshaller.ToYaml(&op)
		if err != nil {
			return "", err
		}

		output += "\n---\n" + string(yaml)
	}
	for _, op := range g.AuthConfigs {
		marshaller := YamlMarshaller{}
		yaml, err := marshaller.ToYaml(&op)
		if err != nil {
			return "", err
		}

		output += "\n---\n" + string(yaml)
	}

	for _, op := range g.VirtualHostOptions {
		marshaller := YamlMarshaller{}
		yaml, err := marshaller.ToYaml(&op)
		if err != nil {
			return "", err
		}

		output += "\n---\n" + string(yaml)
	}

	// need to remove a few values
	//  creationTimestamp: null
	// status: {}
	// status:
	// parents: null
	output = strings.ReplaceAll(output, "  creationTimestamp: null\n", "")
	output = strings.ReplaceAll(output, "status:\n", "")
	output = strings.ReplaceAll(output, "parents: null\n", "")
	output = strings.ReplaceAll(output, "status: {}\n", "")
	output = strings.ReplaceAll(output, "\n\n\n", "\n")
	output = strings.ReplaceAll(output, "\n\n", "\n")
	output = strings.ReplaceAll(output, "spec: {}\n", "")

	// TODO remove leading and trailing ---
	// log.Printf("%s", output)
	return output, nil
}
