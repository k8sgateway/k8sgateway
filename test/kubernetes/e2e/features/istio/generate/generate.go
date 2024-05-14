package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/solo-io/gloo/projects/gloo/pkg/defaults"
	"github.com/solo-io/gloo/test/kubernetes/e2e/features/istio"
	"github.com/solo-io/gloo/test/kubernetes/e2e/utils"
	"github.com/solo-io/skv2/codegen/util"
)

const (
	// exampleNs is the namespace where the resources will be created. Change this to the namespace where you want to create the resources
	exampleNs = defaults.GlooSystem
)

// Dev tool to generate the manifest files for the test suite for demo and docs purposes
//
//go:generate go run ./generate.go
func main() {
	log.Println("starting generate for istio examples")

	// use the Gloo Edge Gateway api resources with automtls enabled
	edgeGatewayApiResources := istio.GetGlooGatewayEdgeResources(exampleNs, istio.UpstreamConfigOpts{})
	automtlsGeneratedExample := filepath.Join(util.MustGetThisDir(), "generated_example", fmt.Sprintf("automtls-enabled-%s", istio.EdgeApisRoutingFileName))
	err := utils.WriteResourcesToFile(edgeGatewayApiResources, automtlsGeneratedExample)
	if err != nil {
		panic(err)
	}

	// automtls disabled
	edgeGatewayApiResources = istio.GetGlooGatewayEdgeResources(exampleNs, istio.UpstreamConfigOpts{DisableIstioAutoMtls: true})
	disableAutomtlsGeneratedExample := filepath.Join(util.MustGetThisDir(), "generated_example", fmt.Sprintf("automtls-disabled-%s", istio.EdgeApisRoutingFileName))
	err = utils.WriteResourcesToFile(edgeGatewayApiResources, disableAutomtlsGeneratedExample)
	if err != nil {
		panic(err)
	}

	// Upstream sslConfig is set
	edgeGatewayApiResources = istio.GetGlooGatewayEdgeResources(exampleNs, istio.UpstreamConfigOpts{SetSslConfig: true})
	upstreamSslConfigGeneratedExample := filepath.Join(util.MustGetThisDir(), "generated_example", fmt.Sprintf("automtls-disabled-%s", istio.EdgeApisRoutingFileName))
	err = utils.WriteResourcesToFile(edgeGatewayApiResources, upstreamSslConfigGeneratedExample)
	if err != nil {
		panic(err)
	}

	// Upstream sslConfig is set and automtls is disabled
	edgeGatewayApiResources = istio.GetGlooGatewayEdgeResources(exampleNs, istio.UpstreamConfigOpts{SetSslConfig: true, DisableIstioAutoMtls: true})
	sslConfigAndDisableAutomtlsGeneratedExample := filepath.Join(util.MustGetThisDir(), "generated_example", fmt.Sprintf("sslconfig-and-automtls-disabled-%s", istio.EdgeApisRoutingFileName))
	err = utils.WriteResourcesToFile(edgeGatewayApiResources, sslConfigAndDisableAutomtlsGeneratedExample)
	if err != nil {
		panic(err)
	}

	log.Println("finished generate for istio examples")
}
