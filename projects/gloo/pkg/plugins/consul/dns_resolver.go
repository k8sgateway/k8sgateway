package consul

import (
	"context"
	"net"
	"sync"

	"github.com/rotisserie/eris"
)

//go:generate mockgen -destination ./mocks/dnsresolver_mock.go github.com/solo-io/gloo/projects/gloo/pkg/plugins/consul DnsResolver
//go:generate gofmt -w ./mocks/
//go:generate goimports -w ./mocks/

type DnsResolver interface {
	Resolve(ctx context.Context, address string) ([]net.IPAddr, error)
}

var (
	_ DnsResolver = new(ConsulDnsResolver)
	_ DnsResolver = new(DnsResolverWithFallback)
)

func NewConsulDnsResolver(address string) DnsResolver {
	basicResolver := &ConsulDnsResolver{
		DnsAddress: address,
	}

	return &DnsResolverWithFallback{
		resolver:            basicResolver,
		previousResolutions: make(map[string][]net.IPAddr),
	}
}

type ConsulDnsResolver struct {
	DnsAddress string
}

func (c *ConsulDnsResolver) Resolve(ctx context.Context, address string) ([]net.IPAddr, error) {
	res := net.Resolver{
		PreferGo: true, // otherwise we may use cgo which doesn't resolve on my mac in testing
		Dial: func(ctx context.Context, network, address string) (conn net.Conn, err error) {
			// DNS typically uses UDP and falls back to TCP if the response size is greater than one packet
			// (originally 512 bytes). we use TCP to ensure we receive all IPs in a large DNS response
			return net.Dial("tcp", c.DnsAddress)
		},
	}
	ipAddrs, err := res.LookupIPAddr(ctx, address)
	if err != nil {
		return nil, err
	}
	if len(ipAddrs) == 0 {
		return nil, eris.Errorf("Consul service returned an address that couldn't be parsed as an IP (%s), "+
			"resolved as a hostname at %s but the DNS server returned no results", address, c.DnsAddress)
	}
	return ipAddrs, nil
}

type DnsResolverWithFallback struct {
	resolver DnsResolver

	mutex               sync.RWMutex
	previousResolutions map[string][]net.IPAddr
}

func (d *DnsResolverWithFallback) Resolve(ctx context.Context, address string) ([]net.IPAddr, error) {
	ipAddrs, err := d.resolver.Resolve(ctx, address)

	// If we successfully resolved the addresses, update our last known state and return
	if err == nil {
		d.mutex.Lock()
		defer d.mutex.Unlock()
		d.previousResolutions[address] = ipAddrs
		return ipAddrs, nil
	}

	// If we did not successfully resolve the addresses, attempt to use the last known state
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	lastKnownIdAddrs, resolvedPreviously := d.previousResolutions[address]
	if !resolvedPreviously {
		return nil, err
	}
	return lastKnownIdAddrs, nil
}
