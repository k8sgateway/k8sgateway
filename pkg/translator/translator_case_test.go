package translator_test

import (
	"context"
	"fmt"

	"github.com/onsi/ginkgo/v2"
	errors "github.com/rotisserie/eris"
	"github.com/solo-io/gloo/v2/pkg/reports"
	. "github.com/solo-io/gloo/v2/pkg/translator"
	"github.com/solo-io/gloo/v2/pkg/translator/testutils"
	"google.golang.org/protobuf/reflect/protoreflect"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	gwv1 "sigs.k8s.io/gateway-api/apis/v1"
)

type TestCase struct {
	Name             string
	InputFiles       []string
	ResultsByGateway map[types.NamespacedName]ExpectedTestResult
}

type ActualTestResult struct {
	ProxyResult ProxyResult
	// Reports     map[types.NamespacedName]*reports.GatewayReport
	//TODO(Law): figure out how RouteReports fit in
}

type ExpectedTestResult struct {
	ProxyResult string
	// Reports     map[types.NamespacedName]*reports.GatewayReport
}

func (r ExpectedTestResult) Equals(actual ActualTestResult) (bool, error) {
	proxy, err := testutils.ReadProxyFromFile(r.ProxyResult)
	if err != nil {
		return false, err
	}

	if len(proxy.ListenerAndRoutes) != len(actual.ProxyResult.ListenerAndRoutes) {
		return false, nil
	}

	for i := range proxy.ListenerAndRoutes {
		v1 := protoreflect.ValueOf(proxy.ListenerAndRoutes[i].Listener.ProtoReflect())
		v2 := protoreflect.ValueOf(actual.ProxyResult.ListenerAndRoutes[i].Listener.ProtoReflect())
		if !v1.Equal(v2) {
			return false, nil
		}
		if len(proxy.ListenerAndRoutes[i].RouteConfigs) != len(actual.ProxyResult.ListenerAndRoutes[i].RouteConfigs) {
			return false, nil
		}
		for j := range proxy.ListenerAndRoutes[i].RouteConfigs {
			v1 := protoreflect.ValueOf(proxy.ListenerAndRoutes[i].RouteConfigs[j].ProtoReflect())
			v2 := protoreflect.ValueOf(actual.ProxyResult.ListenerAndRoutes[i].RouteConfigs[j].ProtoReflect())
			if !v1.Equal(v2) {
				return false, nil
			}
		}
	}

	//TODO: return reflect.DeepEqual(r.Reports, actual.Reports), nil
	return true, nil
}

// map of gwv1.GW namespace/name to translation result
func (tc TestCase) Run(ctx context.Context) (map[types.NamespacedName]bool, error) {
	// load inputs

	var (
		gateways     []*gwv1.Gateway
		dependencies []client.Object
	)
	for _, file := range tc.InputFiles {
		objs, err := testutils.LoadFromFiles(ctx, file)
		if err != nil {
			return nil, err
		}
		for _, obj := range objs {
			switch obj := obj.(type) {
			case *gwv1.Gateway:
				gateways = append(gateways, obj)
			default:
				dependencies = append(dependencies, obj)
			}
		}
	}

	queries := testutils.BuildGatewayQueries(dependencies)

	results := make(map[types.NamespacedName]bool)
	for _, gw := range gateways {

		ref := types.NamespacedName{
			Namespace: gw.Namespace,
			Name:      gw.Name,
		}
		reportsMap := reports.NewReportMap()
		reporter := reports.NewReporter(&reportsMap)

		// translate gateway
		proxyResult := NewTranslator().TranslateProxy(
			ctx,
			gw,
			queries,
			reporter,
		)

		act, _ := testutils.MarshalYamlProxyResult(*proxyResult)
		fmt.Fprintf(ginkgo.GinkgoWriter, "actual result:\n %s \n", act)

		actReport, _ := testutils.MarshalAnyYaml(reportsMap)
		fmt.Fprintf(ginkgo.GinkgoWriter, "actual reports:\n %s \n", actReport)

		actual := ActualTestResult{
			ProxyResult: *proxyResult,
			// Reports:     reportsMap.Gateways,
		}

		expected, ok := tc.ResultsByGateway[ref]
		if !ok {
			return nil, errors.Errorf("no expected result found for gateway %v", ref)
		}

		equal, err := expected.Equals(actual)
		if err != nil {
			return nil, err
		}

		results[ref] = equal
	}

	return results, nil
}
