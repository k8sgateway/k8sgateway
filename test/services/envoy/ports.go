package envoy

import (
	"sync/atomic"

	"github.com/solo-io/gloo/projects/gloo/pkg/defaults"

	"github.com/solo-io/gloo/test/ginkgo/parallel"
)

var (
	bindPort = uint32(10080)

	adminPort  = defaults.EnvoyAdminPort
	httpPort   = defaults.HttpPort
	httpsPort  = defaults.HttpsPort
	tcpPort    = defaults.TcpPort
	hybridPort = defaults.HybridPort
)

func NextBindPort() uint32 {
	return advancePort(&bindPort)
}

func advanceRequestPorts() {
	httpPort = advancePort(&httpPort)
	httpsPort = advancePort(&httpsPort)
	tcpPort = advancePort(&tcpPort)
	hybridPort = advancePort(&hybridPort)
	adminPort = advancePort(&adminPort)

	// NOTE TO DEVELOPERS:
	// This file contains definitions for port values that the test suite will use
	// Ideally these ports would be owned exclusively by the envoy package.
	// However, the challenge is that we have some default resources, which are created using the defaults package.
	// Therefore, we need to keep the defaults package ports in sync with the envoy ports

	defaults.HttpPort = httpPort
	defaults.HttpsPort = httpsPort
	defaults.HybridPort = hybridPort
	defaults.TcpPort = tcpPort
	defaults.EnvoyAdminPort = adminPort
}

func advancePort(p *uint32) uint32 {
	return atomic.AddUint32(p, 1) + uint32(parallel.GetPortOffset())
}
