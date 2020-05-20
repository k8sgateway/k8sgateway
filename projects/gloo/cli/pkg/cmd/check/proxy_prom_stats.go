package check

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	v1 "k8s.io/api/apps/v1"

	"github.com/solo-io/gloo/pkg/cliutil"
	"github.com/solo-io/gloo/projects/gloo/pkg/defaults"
)

const promStatsPath = "/stats/prometheus"

const metricsUpdateInterval = time.Millisecond * 250

func checkProxiesPromStats(ctx context.Context, glooNamespace string, deployments *v1.DeploymentList) (bool, error) {
	for _, deployment := range deployments.Items {
		if deployment.Name == "gateway-proxy" || deployment.Name == "ingress-proxy" || deployment.Name == "knative-external-proxy" || deployment.Name == "knative-internal-proxy" {
			if passed, err := checkProxyPromStats(ctx, glooNamespace, deployment.Name); !passed || err != nil {
				return passed, err
			}
		}
	}
	fmt.Printf("OK\n")
	return true, nil
}

func checkProxyPromStats(ctx context.Context, glooNamespace string, deploymentName string) (bool, error) {

	// check if any proxy instances are out of sync with the Gloo control plane
	errMessage := "Problem while checking for out of sync proxies"

	// port-forward proxy deployment and get prometheus metrics
	freePort, err := cliutil.GetFreePort()
	if err != nil {
		fmt.Println(errMessage)
		return false, err
	}
	localPort := strconv.Itoa(freePort)
	adminPort := strconv.Itoa(int(defaults.EnvoyAdminPort))
	// stats is the string containing all stats from /stats/prometheus
	stats, portFwdCmd, err := cliutil.PortForwardGet(ctx, glooNamespace, "deploy/"+deploymentName,
		localPort, adminPort, false, promStatsPath)
	if err != nil {
		fmt.Println(errMessage)
		return false, err
	}
	if portFwdCmd.Process != nil {
		defer portFwdCmd.Process.Release()
		defer portFwdCmd.Process.Kill()
	}

	if !checkProxyConnectedState(stats, deploymentName, errMessage,
		"Your "+deploymentName+" is out of sync with the Gloo control plane and is not receiving valid gloo config.\n"+
			"You may want to try using the `glooctl proxy logs` or `glooctl debug logs` commands.") {
		return false, nil
	}

	return checkProxyUpdate(stats, localPort, deploymentName, errMessage)
}

// checks that envoy_control_plane_connected_state metric has a value of 1
func checkProxyConnectedState(stats string, deploymentName string, genericErrMessage string, connectedStateErrMessage string) bool {

	if strings.TrimSpace(stats) == "" {
		fmt.Println(genericErrMessage+": could not find any metrics at", promStatsPath, "endpoint of the "+deploymentName+" deployment")
		return false
	}

	if !strings.Contains(stats, "envoy_control_plane_connected_state{} 1") {
		fmt.Println(connectedStateErrMessage)
		return false
	}

	return true
}

// checks that update_rejected and update_failure stats have not increased by getting stats from /stats/prometheus again
func checkProxyUpdate(stats string, localPort string, deploymentName string, errMessage string) (bool, error) {

	// wait for metrics to update
	time.Sleep(metricsUpdateInterval)

	// gather metrics again
	res, err := http.Get("http://localhost:" + localPort + promStatsPath)
	if err != nil {
		fmt.Println(errMessage)
		return false, err
	}
	if res.StatusCode != 200 {
		fmt.Println(errMessage+": received unexpected status code", res.StatusCode, "from", promStatsPath, "endpoint of the "+deploymentName+" deployment")
		return false, nil
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(errMessage)
		return false, err
	}
	res.Body.Close()
	newStats := string(b)

	if strings.TrimSpace(newStats) == "" {
		fmt.Println(errMessage+": could not find any metrics at", promStatsPath, "endpoint of the "+deploymentName+" deployment")
		return false, nil
	}

	// for example, look for stats like "envoy_http_rds_update_attempt" and "envoy_http_rds_update_rejected"
	// more info at https://www.envoyproxy.io/docs/envoy/latest/configuration/overview/mgmt_server#xds-subscription-statistics
	desiredMetricsSegments := []string{"update_rejected", "update_failure"}
	statsMap := parseMetrics(stats, desiredMetricsSegments, deploymentName)
	newStatsMap := parseMetrics(newStats, desiredMetricsSegments, deploymentName)

	if reflect.DeepEqual(newStatsMap, statsMap) {
		return true, nil
	}

	for metricName, oldVal := range statsMap {
		newVal, ok := newStatsMap[metricName]
		if ok && strings.Contains(metricName, "rejected") && newVal > oldVal {
			// for example, if envoy_http_rds_update_rejected{envoy_http_conn_manager_prefix="http",envoy_rds_route_config="listener-__-8080-routes"}
			// increases, which occurs if envoy cannot parse the config from gloo
			fmt.Printf("An update to your "+deploymentName+" deployment was rejected due to schema/validation errors. The %v metric increased.\n"+
				"You may want to try using the `glooctl proxy logs` or `glooctl debug logs` commands.\n", metricName)
			return false, nil
		} else if ok && strings.Contains(metricName, "failure") && newVal > oldVal {
			fmt.Printf("An update to your "+deploymentName+" deployment was rejected due to network errors. The %v metric increased.\n"+
				"You may want to try using the `glooctl proxy logs` or `glooctl debug logs` commands.\n", metricName)
			return false, nil
		}
	}

	return true, nil
}

// parseMetrics parses prometheus metrics and returns a map from the metric name and labels to its value.
// It expects to only look for int values!
func parseMetrics(stats string, desiredMetricSegments []string, deploymentName string) map[string]int {
	statsMap := make(map[string]int)
	statsLines := strings.Split(stats, "\n")
	for _, line := range statsLines {
		trimLine := strings.TrimSpace(line)
		if strings.HasPrefix(trimLine, "#") || trimLine == "" {
			continue // Ignore comments, help text, type info, empty lines (https://github.com/prometheus/docs/blob/master/content/docs/instrumenting/exposition_formats.md#comments-help-text-and-type-information)
		}
		desiredMetric := false
		for _, s := range desiredMetricSegments {
			if strings.Contains(trimLine, s) {
				desiredMetric = true
			}
		}
		// sample line: "envoy_http_rds_update_rejected{envoy_http_conn_manager_prefix="http",envoy_rds_route_config="listener-__-8080-routes"} 90"
		if desiredMetric {
			pieces := strings.Fields(trimLine)                    // split by white spaces
			metric := strings.Join(pieces[0:len(pieces)-1], "")   // get all but last piece (as one string)- this is metric name and labels
			metricVal, err := strconv.Atoi(pieces[len(pieces)-1]) // get last piece (as int)- this is metric value
			if err != nil {
				fmt.Printf("Found an unexpected format in metrics at %v endpoint of the "+deploymentName+" deployment. "+
					"Expected %v metric to have an int value but got value %v.\nContinuing check...", promStatsPath, metric, pieces[len(pieces)-1])
				continue
			}
			statsMap[metric] = metricVal
		}
	}
	return statsMap
}
