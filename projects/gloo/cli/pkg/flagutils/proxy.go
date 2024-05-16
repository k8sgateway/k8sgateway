package flagutils

import (
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/spf13/pflag"
)

const (
	k8sGatewaySourceFlag  = "kube"
	edgeGatewaySourceFlag = "edge"
	allProxies            = "all"
)

func AddGetProxyFlags(set *pflag.FlagSet, proxy *options.GetProxy) {
	addK8sGatewaySourceFlag(set, &proxy.K8sGatewaySource)
	addEdgeGatewaySourceFlag(set, &proxy.EdgeGatewaySource)
	addAllFlag(set, &proxy.All)
}

func addK8sGatewaySourceFlag(set *pflag.FlagSet, k8sGatewaySource *bool) {
	set.BoolVarP(k8sGatewaySource, k8sGatewaySourceFlag, "", false, "include proxies produced from k8s gateway resources")
}

func addEdgeGatewaySourceFlag(set *pflag.FlagSet, edgeGatewaySource *bool) {
	set.BoolVarP(edgeGatewaySource, edgeGatewaySourceFlag, "", false, "include proxies produced from edge gateway resources")
}

func addAllFlag(set *pflag.FlagSet, all *bool) {
	set.BoolVar(all, allProxies, false, "get all proxies")
}
