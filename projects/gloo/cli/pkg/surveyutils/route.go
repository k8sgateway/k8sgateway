package surveyutils

import (
	"fmt"
	"sort"

	"github.com/pkg/errors"
	"github.com/solo-io/gloo/pkg/cliutil"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/helpers"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/aws"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/rest"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

const (
	pathMatch_Prefix = "prefix"
	pathMatch_Regex  = "regex"
	pathMatch_Exact  = "exact"
)

var pathMatchOptions = []string{
	pathMatch_Prefix,
	pathMatch_Regex,
	pathMatch_Exact,
}

func getMatcherInteractive(match *options.RouteMatchers) error {
	var pathType string
	if err := cliutil.ChooseFromList(
		"Choose a path match type: ",
		&pathType,
		pathMatchOptions,
	); err != nil {
		return err
	}
	if pathType == "" {
		return errors.Errorf("must specify one of %v", pathMatchOptions)
	}

	var pathMatch string
	if err := cliutil.GetStringInputDefault(
		fmt.Sprintf("What path %v should we match? ", pathType),
		&pathMatch,
		"/",
	); err != nil {
		return err
	}

	switch pathType {
	case "exact":
		match.PathExact = pathMatch
	case "regex":
		match.PathRegex = pathMatch
	case "prefix":
		match.PathPrefix = pathMatch
	default:
		return errors.Errorf("must specify one of %v", pathMatchOptions)
	}

	if err := cliutil.GetStringSliceInput(
		fmt.Sprintf("Add a header matcher for this function (empty to skip)? %v", match.HeaderMatcher.Entries),
		&match.HeaderMatcher.Entries,
	); err != nil {
		return err
	}

	if err := cliutil.GetStringSliceInput(
		fmt.Sprintf("HTTP Method to match for this route (empty to skip)? %v", match.Methods),
		&match.Methods,
	); err != nil {
		return err
	}

	return nil
}

func getDestinationInteractive(dest *options.Destination) error {
	// collect secrets list
	usClient := helpers.MustUpstreamClient()
	ussByKey := make(map[string]*v1.Upstream)
	var usKeys []string
	for _, ns := range helpers.MustGetNamespaces() {
		usList, err := usClient.List(ns, clients.ListOpts{})
		if err != nil {
			return err
		}
		for _, us := range usList {
			ref := us.Metadata.Ref()
			ussByKey[ref.Key()] = us
			usKeys = append(usKeys, ref.Key())
		}
	}
	if len(usKeys) == 0 {
		return errors.Errorf("no upstreams found. create an upstream first or enable " +
			"discovery.")
	}
	var usKey string
	if err := cliutil.ChooseFromList(
		"Choose the upstream to route to: ",
		&usKey,
		usKeys,
	); err != nil {
		return err
	}
	us, ok := ussByKey[usKey]
	if !ok {
		return errors.Errorf("internal error: upstream map not populated")
	}
	dest.Upstream = us.Metadata.Ref()
	switch ut := us.UpstreamSpec.UpstreamType.(type) {
	case *v1.UpstreamSpec_Aws:
		if err := getAwsDestinationSpecInteractive(&dest.DestinationSpec.Aws, ut.Aws); err != nil {
			return err
		}
	case v1.ServiceSpecGetter:
		svcSpec := ut.GetServiceSpec()
		if svcSpec == nil {
			return nil
		}
		switch svcType := svcSpec.PluginType.(type) {
		case *plugins.ServiceSpec_Rest:
			if err := getRestDestinationSpecInteractive(&dest.DestinationSpec.Rest, svcType.Rest); err != nil {
				return err
			}
		}
	}
	return nil
}

func getPluginsInteractive(dest *options.RoutePlugins) error {
	yes, err := cliutil.GetYesInput("do you wish to add a prefix-rewrite transformation to the route [y/n]?\n" +
		"note that this will be overridden if your routes point to function destinations")
	if err != nil {
		return err
	}

	if !yes {
		return nil
	}

	var prefixRewrite string

	if err := cliutil.GetStringInput("rewrite the matched portion of HTTP requests with this prefix: ", &prefixRewrite); err != nil {
		return err
	}

	dest.PrefixRewrite.Value = &prefixRewrite

	return nil
}

func getAwsDestinationSpecInteractive(spec *options.AwsDestinationSpec, ut *aws.UpstreamSpec) error {
	var fnNames []string
	for _, fn := range ut.LambdaFunctions {
		fnNames = append(fnNames, fn.LogicalName)
	}
	if err := cliutil.ChooseFromList(
		"which function should this route invoke? ",
		&spec.LogicalName,
		fnNames,
	); err != nil {
		return err
	}

	return nil
}

func getRestDestinationSpecInteractive(spec *options.RestDestinationSpec, restSpec *rest.ServiceSpec) error {
	var fnNames []string
	for fn := range restSpec.Transformations {
		fnNames = append(fnNames, fn)
	}
	sort.Strings(fnNames)
	if err := cliutil.ChooseFromList(
		"which function should this route invoke? ",
		&spec.FunctionName,
		fnNames,
	); err != nil {
		return err
	}
	if err := cliutil.GetStringSliceInput(
		fmt.Sprintf("Add a header parameter for this function (empty to skip)? %v", spec.Parameters.Entries),
		&spec.Parameters.Entries,
	); err != nil {
		return err
	}

	return nil
}

func AddRouteFlagsInteractive(opts *options.Options) error {
	// collect vs list
	vsByKey := make(map[string]core.ResourceRef)
	vsKeys := []string{"create a new virtualservice"}
	var namespaces []string
	for _, ns := range helpers.MustGetNamespaces() {
		namespaces = append(namespaces, ns)
		vsList, err := helpers.MustVirtualServiceClient().List(ns, clients.ListOpts{})
		if err != nil {
			return err
		}
		for _, vs := range vsList {
			ref := vs.Metadata.Ref()
			vsByKey[ref.Key()] = ref
			vsKeys = append(vsKeys, ref.Key())
		}
	}

	var vsKey string
	if err := cliutil.ChooseFromList(
		"Choose a Virtual Service to add the route to: (empty to add "+
			"to a default virtual service. the default virtual service matches "+
			"the '*' domain and will be created if it does not exist) ",
		&vsKey,
		vsKeys,
	); err != nil {
		return err
	}
	opts.Metadata.Name = vsByKey[vsKey].Name
	opts.Metadata.Namespace = vsByKey[vsKey].Namespace

	if opts.Metadata.Name == "" || opts.Metadata.Namespace == "" {
		if err := cliutil.GetStringInput("name of the virtual service: ", &opts.Metadata.Name); err != nil {
			return err
		}
		if err := cliutil.ChooseFromList(
			"namespace of the virtual service: ",
			&opts.Metadata.Namespace,
			namespaces,
		); err != nil {
			return err
		}
	} else {
		// only get the insert index if the vs is predefined
		if err := cliutil.GetUint32InputDefault(
			fmt.Sprintf("where do you want to insert the route in the "+
				"virtual service's route list? "),
			&opts.Add.Route.InsertIndex,
			0,
		); err != nil {
			return err
		}
	}

	if err := getMatcherInteractive(&opts.Add.Route.Matcher); err != nil {
		return err
	}
	if err := getDestinationInteractive(&opts.Add.Route.Destination); err != nil {
		return err
	}
	if err := getPluginsInteractive(&opts.Add.Route.Plugins); err != nil {
		return err
	}

	return nil
}

func RemoveRouteFlagsInteractive(opts *options.Options) error {
	// collect vs list
	vsByKey := make(map[string]core.ResourceRef)
	var vsKeys []string
	var namespaces []string
	for _, ns := range helpers.MustGetNamespaces() {
		namespaces = append(namespaces, ns)
		vsList, err := helpers.MustVirtualServiceClient().List(ns, clients.ListOpts{})
		if err != nil {
			return err
		}
		for _, vs := range vsList {
			ref := vs.Metadata.Ref()
			vsByKey[ref.Key()] = ref
			vsKeys = append(vsKeys, ref.Key())
		}
	}

	if len(vsKeys) == 0 {
		return errors.Errorf("no virtual services found")
	}

	var vsKey string
	if err := cliutil.ChooseFromList(
		"Choose a Virtual Service from which to remove the route: ",
		&vsKey,
		vsKeys,
	); err != nil {
		return err
	}
	opts.Metadata.Name = vsByKey[vsKey].Name
	opts.Metadata.Namespace = vsByKey[vsKey].Namespace

	vs, err := helpers.MustVirtualServiceClient().Read(opts.Metadata.Namespace, opts.Metadata.Name,
		clients.ReadOpts{Ctx: opts.Top.Ctx})
	if err != nil {
		return err
	}
	if vs.VirtualHost == nil {
		return errors.Errorf("invalid virtual service %v", opts.Metadata.Ref())
	}
	if len(vs.VirtualHost.Routes) == 0 {
		return errors.Errorf("no routes defined for virtual service %v", opts.Metadata.Ref())
	}

	var routes []string
	for i, r := range vs.VirtualHost.Routes {
		routes = append(routes, fmt.Sprintf("%v: %+v", i, r.Matcher.PathSpecifier))
	}

	var chosenRoute string
	if err := cliutil.ChooseFromList(
		"Choose a Virtual Service from which to remove the route: ",
		&chosenRoute,
		routes,
	); err != nil {
		return err
	}

	for i, route := range routes {
		if route == chosenRoute {
			opts.Remove.Route.RemoveIndex = uint32(i)
			break
		}
	}

	return nil
}

func SelectVirtualServiceInteractive(opts *options.Options) error {
	// collect vs list
	vsByKey := make(map[string]core.ResourceRef)
	var vsKeys []string
	var namespaces []string
	for _, ns := range helpers.MustGetNamespaces() {
		namespaces = append(namespaces, ns)
		vsList, err := helpers.MustVirtualServiceClient().List(ns, clients.ListOpts{})
		if err != nil {
			return err
		}
		for _, vs := range vsList {
			ref := vs.Metadata.Ref()
			vsByKey[ref.Key()] = ref
			vsKeys = append(vsKeys, ref.Key())
		}
	}

	if len(vsKeys) == 0 {
		return errors.Errorf("no virtual services found")
	}

	var vsKey string
	if err := cliutil.ChooseFromList(
		"Choose a Virtual Service: ",
		&vsKey,
		vsKeys,
	); err != nil {
		return err
	}
	opts.Metadata.Name = vsByKey[vsKey].Name
	opts.Metadata.Namespace = vsByKey[vsKey].Namespace

	return nil
}
