package common

import (
	"context"
	"io"
	"math"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/solo-io/gloo/pkg/utils/kubeutils/portforward"

	"github.com/solo-io/solo-kit/pkg/errors"

	gloodebug "github.com/solo-io/gloo/projects/gloo/pkg/debug"

	"github.com/avast/retry-go/v4"
	"github.com/solo-io/gloo/pkg/utils/kubeutils"

	"github.com/solo-io/gloo/pkg/cliutil"

	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/grpc/debug"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"google.golang.org/grpc"
)

// GetProxies retrieves the proxies from the Control Plane via the ProxyEndpointServer API
// This is utilized by `glooctl get proxy` to return the content of Proxies
func GetProxies(name string, opts *options.Options) (gloov1.ProxyList, error) {
	settings, err := GetSettings(opts)
	if err != nil {
		return nil, err
	}

	proxyEndpointPort, err := computeProxyEndpointPort(settings)
	if err != nil {
		return nil, err
	}

	return getProxiesFromControlPlane(opts, name, proxyEndpointPort)
}

// ListProxiesFromSettings retrieves the proxies from the Control Plane via the ProxyEndpointServer API
// This is utilized by `glooctl check` to report the statuses of Proxies
func ListProxiesFromSettings(namespace string, opts *options.Options, settings *gloov1.Settings) (gloov1.ProxyList, error) {
	proxyEndpointPort, err := computeProxyEndpointPort(settings)
	if err != nil {
		return nil, err
	}

	return listProxiesFromControlPlane(opts, namespace, proxyEndpointPort)
}

func computeProxyEndpointPort(settings *gloov1.Settings) (string, error) {
	proxyEndpointAddress := settings.GetGloo().GetProxyDebugBindAddr()
	if proxyEndpointAddress == "" {
		// This can occur if you are querying a Settings object that was created before the ProxyDebugBindAddr API
		// was introduced to the Settings CR. In practice, this should never occur, as the API has existed for many releases.
		return "", errors.Errorf("ProxyDebugBindAddr is empty. Consider upgrading the version of Gloo")
	}

	_, proxyEndpointPort, err := net.SplitHostPort(proxyEndpointAddress)
	return proxyEndpointPort, err
}

func getProxiesFromControlPlane(opts *options.Options, name string, proxyEndpointPort string) (gloov1.ProxyList, error) {
	proxyRequest := &debug.ProxyEndpointRequest{
		Name: name,
		// It is important that we use the Proxy.Namespace here, as opposed to the opts.Metadata.Namespace
		// The former is where Proxies will be searched, the latter is where Gloo is installed
		Namespace: opts.Get.Proxy.Namespace,
		Source:    getProxySource(opts.Get.Proxy),
		Selector:  opts.Get.Selector.MustMap(),
	}

	if opts.Get.Proxy.All {
		// Each of the properties in a request act as a filter of the requested proxies.
		// By supplying no filters, we are requesting all the proxies
		proxyRequest = &debug.ProxyEndpointRequest{
			Name:      "",
			Namespace: "",
			Source:    "",
			Selector:  nil,
		}
	}

	return requestProxiesFromControlPlane(opts, proxyRequest, proxyEndpointPort)
}

func listProxiesFromControlPlane(opts *options.Options, namespace, proxyEndpointPort string) (gloov1.ProxyList, error) {
	proxyRequest := &debug.ProxyEndpointRequest{
		Name:      "",
		Namespace: namespace,
		Source:    getProxySource(opts.Get.Proxy),
		Selector:  opts.Get.Selector.MustMap(),
	}

	return requestProxiesFromControlPlane(opts, proxyRequest, proxyEndpointPort)
}

// getProxySource returns the value of the ProxySource based on the request
// Proxies may either be "sourced" by the K8s Gateway translation or the Edge Gateway translation.
// The ProxyEndpointServer exposes an API that accepts a source which is a string, however to improve the UX
// of the CLI, we expose boolean flags to end users. The proxy source is set to the string value if
// ONLY one of the flags is enabled, or an empty string meaning "look at all sources"
// Example:
//
//	glooctl get proxies --kube
//		This will only return proxies generated by the k8s gateway translation
//	glooctl get proxies --edge
//		This will only return proxies generated by the edge gateway translation
//	glooctl get proxies --kube --edge
//		This will return proxies generated by EITHER the k8s gateway or edge gateway translations
func getProxySource(proxy options.GetProxy) string {
	proxySource := "" // empty string is considered "all"
	if proxy.EdgeGatewaySource && !proxy.K8sGatewaySource {
		proxySource = gloodebug.EdgeGatewaySourceName
	}
	if !proxy.EdgeGatewaySource && proxy.K8sGatewaySource {
		proxySource = gloodebug.K8sGatewaySourceName
	}
	return proxySource
}

// requestProxiesFromControlPlane executes a gRPC request against the Control Plane (Gloo) against a given port (proxyEndpointPort).
// Proxies are an intermediate resource that are often persisted in-memory in the Control Plane.
// To improve debuggability, we expose an API to return the current proxies, and rely on this CLI method to expose that to users
func requestProxiesFromControlPlane(opts *options.Options, request *debug.ProxyEndpointRequest, proxyEndpointPort string) (gloov1.ProxyList, error) {
	remotePort, err := strconv.Atoi(proxyEndpointPort)
	if err != nil {
		return nil, err
	}

	logger := cliutil.GetLogger()
	var outWriter, errWriter io.Writer
	errWriter = io.MultiWriter(logger, os.Stderr)
	if opts.Top.Verbose {
		outWriter = io.MultiWriter(logger, os.Stdout)
	} else {
		outWriter = logger
	}

	requestCtx, cancel := context.WithTimeout(opts.Top.Ctx, 30*time.Second)
	defer cancel()

	portForwarder := portforward.NewPortForwarder(
		portforward.WithDeployment(kubeutils.GlooDeploymentName, opts.Metadata.GetNamespace()),
		portforward.WithRemotePort(remotePort),
		portforward.WithWriters(outWriter, errWriter),
	)
	if err := portForwarder.Start(
		requestCtx,
		retry.LastErrorOnly(true),
		retry.Delay(100*time.Millisecond),
		retry.DelayType(retry.BackOffDelay),
		retry.Attempts(5),
	); err != nil {
		return nil, err
	}
	defer func() {
		portForwarder.Close()
		portForwarder.WaitForStop()
	}()

	var proxyEndpointResponse *debug.ProxyEndpointResponse
	requestErr := retry.Do(func() error {
		cc, err := grpc.DialContext(requestCtx, portForwarder.Address(), grpc.WithInsecure())
		if err != nil {
			return err
		}
		pxClient := debug.NewProxyEndpointServiceClient(cc)
		r, err := pxClient.GetProxies(requestCtx, request,
			// Some proxies can become very large and exceed the default 100Mb limit
			// For this reason we want remove the limit but will settle for a limit of MaxInt32
			// as we don't anticipate proxies to exceed this
			grpc.MaxCallRecvMsgSize(math.MaxInt32),
		)
		proxyEndpointResponse = r
		return err
	},
		retry.LastErrorOnly(true),
		retry.Delay(100*time.Millisecond),
		retry.DelayType(retry.BackOffDelay),
		retry.Attempts(5),
	)

	if requestErr != nil {
		return nil, requestErr
	}

	return proxyEndpointResponse.GetProxies(), nil
}
