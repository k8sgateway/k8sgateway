package convert

import (
	"fmt"
	"os"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	filesMetrics = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "files_evaluated",
		},
	)
	totalLines = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "total_lines_of_yaml",
		}, []string{"api"},
	)
	glooConfigMetric = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "gloo_config_count",
		}, []string{"type"},
	)
	gatewayAPIConfigMetrics = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "gatewayapi_config_count",
		}, []string{"type"},
	)
)

func init() {
	prometheus.MustRegister(filesMetrics)
	prometheus.MustRegister(totalLines)
	prometheus.MustRegister(glooConfigMetric)
	prometheus.MustRegister(gatewayAPIConfigMetrics)

	glooConfigMetric.WithLabelValues("AuthConfig").Inc()
	glooConfigMetric.WithLabelValues("RouteTable").Inc()
	glooConfigMetric.WithLabelValues("Upstream").Inc()
	glooConfigMetric.WithLabelValues("VirtualService").Inc()
	glooConfigMetric.WithLabelValues("RouteOption").Inc()
	glooConfigMetric.WithLabelValues("VirtualHostOption").Inc()

	gatewayAPIConfigMetrics.WithLabelValues("AuthConfig").Inc()
	gatewayAPIConfigMetrics.WithLabelValues("HTTPRoute").Inc()
	gatewayAPIConfigMetrics.WithLabelValues("Upstream").Inc()
	gatewayAPIConfigMetrics.WithLabelValues("RouteOption").Inc()
	gatewayAPIConfigMetrics.WithLabelValues("VirtualHostOption").Inc()
}

func printMetrics(outputs []*GatewayAPIOutput) {

	for _, output := range outputs {
		//we need to save the output to metrics
		gatewayAPIConfigMetrics.WithLabelValues("AuthConfig").Add(float64(len(output.AuthConfigs)))
		gatewayAPIConfigMetrics.WithLabelValues("HTTPRoute").Add(float64(len(output.HTTPRoutes)))
		gatewayAPIConfigMetrics.WithLabelValues("Upstream").Add(float64(len(output.Upstreams)))
		gatewayAPIConfigMetrics.WithLabelValues("RouteOption").Add(float64(len(output.RouteOptions)))
		gatewayAPIConfigMetrics.WithLabelValues("VirtualHostOption").Add(float64(len(output.VirtualHostOptions)))
	}

	metrics, _ := prometheus.DefaultGatherer.Gather()
	fmt.Fprintf(os.Stdout, "-------------------------------------\n")
	for _, m := range metrics {
		if *m.Name == "gloo_config_count" {
			var count float64
			for _, t := range m.Metric {
				_, _ = fmt.Fprintf(os.Stdout, "Gloo Config: Number of %s: %v\n", *t.Label[0].Value, *t.Counter.Value-1)
				count += *t.Counter.Value - 1
			}
			_, _ = fmt.Fprintf(os.Stdout, "Total Gloo Config: %v\n", count)
		}
	}
	fmt.Fprintf(os.Stdout, "-------------------------------------\n")
	for _, m := range metrics {
		if *m.Name == "gatewayapi_config_count" {
			var count float64
			for _, t := range m.Metric {
				_, _ = fmt.Fprintf(os.Stdout, "Gateway API Config: Number of %s: %v\n", *t.Label[0].Value, *t.Counter.Value-1)
				count += *t.Counter.Value - 1
			}
			_, _ = fmt.Fprintf(os.Stdout, "Total Gateway API Config: %v\n", count)
		}
	}
	fmt.Fprintf(os.Stdout, "-------------------------------------\n")
	for _, m := range metrics {
		if *m.Name == "total_lines_of_yaml" {
			for _, t := range m.Metric {
				_, _ = fmt.Fprintf(os.Stdout, "Lines of Yaml %s: %v\n", *t.Label[0].Value, *t.Counter.Value-1)
			}
		}
	}
	fmt.Fprintf(os.Stdout, "-------------------------------------\n")
	for _, m := range metrics {
		if *m.Name == "files_evaluated" {
			_, _ = fmt.Fprintf(os.Stdout, "Files evaluated: %v\n", *m.Metric[0].Counter.Value)
		}
	}
}
