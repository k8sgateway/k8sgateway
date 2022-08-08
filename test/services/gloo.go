package services

import (
	"context"
	"fmt"
	"net"
	"sync/atomic"

	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/imdario/mergo"
	. "github.com/onsi/gomega"
	fdsrunner "github.com/solo-io/gloo/projects/discovery/pkg/fds/runner"
	udsrunner "github.com/solo-io/gloo/projects/discovery/pkg/uds/runner"
	"github.com/solo-io/gloo/projects/gloo/pkg/runner"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube"

	"github.com/solo-io/gloo/pkg/utils/settingsutil"

	"github.com/solo-io/solo-kit/test/helpers"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients/memory"

	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
	"github.com/solo-io/gloo/projects/gloo/pkg/defaults"
	"k8s.io/client-go/kubernetes"
)

var glooPortBase = int32(30400)

func AllocateGlooPort() int32 {
	return atomic.AddInt32(&glooPortBase, 1) + int32(config.GinkgoConfig.ParallelNode*1000)
}

func RunGateway(ctx context.Context, justGloo bool) TestClients {
	ns := defaults.GlooSystem
	ro := &RunOptions{
		NsToWrite: ns,
		NsToWatch: []string{"default", ns},
		WhatToRun: What{
			DisableGateway: justGloo,
		},
		KubeClient: helpers.MustKubeClient(),
	}
	return RunGlooGatewayUdsFds(ctx, ro)
}

type What struct {
	DisableGateway bool
	DisableUds     bool
	DisableFds     bool
}

type RunOptions struct {
	NsToWrite      string
	NsToWatch      []string
	WhatToRun      What
	GlooPort       int32
	ValidationPort int32
	RestXdsPort    int32
	Settings       *gloov1.Settings
	KubeClient     kubernetes.Interface
}

//noinspection GoUnhandledErrorResult
func RunGlooGatewayUdsFds(ctx context.Context, runOptions *RunOptions) TestClients {
	// Allocate any required ports which were not explicitly set
	if runOptions.GlooPort == 0 {
		runOptions.GlooPort = AllocateGlooPort()
	}
	if runOptions.ValidationPort == 0 {
		runOptions.ValidationPort = AllocateGlooPort()
	}
	if runOptions.RestXdsPort == 0 {
		runOptions.RestXdsPort = AllocateGlooPort()
	}
	if runOptions.Settings == nil {
		runOptions.Settings = &gloov1.Settings{}
	}

	// Initialize the Settings based on the RunOptions
	settings := &gloov1.Settings{
		WatchNamespaces:    runOptions.NsToWatch,
		DiscoveryNamespace: runOptions.NsToWrite,
		DevMode:            true,
		RefreshRate: &duration.Duration{
			Seconds: 1,
		},
		Gloo: &gloov1.GlooOptions{
			RestXdsBindAddr:    fmt.Sprintf("%s:%d", net.IPv4zero.String(), runOptions.RestXdsPort),
			ValidationBindAddr: fmt.Sprintf("%s:%d", net.IPv4zero.String(), runOptions.ValidationPort),
			XdsBindAddr:        fmt.Sprintf("%s:%d", net.IPv4zero.String(), runOptions.GlooPort),
			ProxyDebugBindAddr: fmt.Sprintf("%s:%d", net.IPv4zero.String(), AllocateGlooPort()),
			RemoveUnusedFilters: &wrappers.BoolValue{
				Value: true,
			},
		},
		Gateway: &gloov1.GatewayOptions{
			PersistProxySpec: &wrappers.BoolValue{
				Value: true,
			},
			EnableGatewayController: &wrappers.BoolValue{
				Value: !runOptions.WhatToRun.DisableGateway,
			},
		},
	}

	// Initialize the Cache used by the Runners
	inMemoryCache := memory.NewInMemoryResourceCache()
	var kubeCache kube.SharedCache
	if runOptions.KubeClient != nil {
		kubeCache = kube.NewKubeCache(ctx)
	}

	// Override any Settings explicitly defined by a test
	mergo.Merge(settings, runOptions.Settings, mergo.WithOverride)

	ctx = settingsutil.WithSettings(ctx, settings)

	// Run Gloo
	glooRunner := runner.NewGlooRunner()
	runErr := glooRunner.Run(ctx, kubeCache, inMemoryCache, settings)
	Expect(runErr).NotTo(HaveOccurred())
	resourceClientset := glooRunner.GetResourceClientset()
	typedClientset := glooRunner.GetTypedClientset()

	// Run FDS (if necessary)
	if !runOptions.WhatToRun.DisableFds {
		go func() {
			defer GinkgoRecover()

			fdsRunner := fdsrunner.NewFDSRunner()
			err := fdsRunner.Run(ctx, kubeCache, inMemoryCache, settings)
			Expect(err).NotTo(HaveOccurred())
		}()
	}

	// Run UDS (if necessary)
	if !runOptions.WhatToRun.DisableUds {
		go func() {
			defer GinkgoRecover()

			udsRunner := udsrunner.NewUDSRunner()
			err := udsRunner.Run(ctx, kubeCache, inMemoryCache, settings)
			Expect(err).NotTo(HaveOccurred())

		}()
	}

	return TestClients{
		GatewayClient:        resourceClientset.Gateways,
		HttpGatewayClient:    resourceClientset.MatchableHttpGateways,
		VirtualServiceClient: resourceClientset.VirtualServices,
		UpstreamClient:       resourceClientset.Upstreams,
		SecretClient:         resourceClientset.Secrets,
		ProxyClient:          resourceClientset.Proxies,
		ServiceClient:        typedClientset.KubeServiceClient,
		GlooPort:             int(runOptions.GlooPort),
		RestXdsPort:          int(runOptions.RestXdsPort),
	}
}
