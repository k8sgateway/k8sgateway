package cloudfoundry

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"

	"code.cloudfoundry.org/copilot"
	copilotapi "code.cloudfoundry.org/copilot/api"
	"github.com/solo-io/gloo/pkg/api/types/v1"
	"github.com/solo-io/gloo/pkg/endpointdiscovery"
)

/*

Based on this example: https://github.com/cloudfoundry/copilot/blob/master/cmd/copilot-clients/istio/main.go
	flag.StringVar(&address, "address", , "ip:port of copilot server")
	flag.StringVar(&caCert, "server-ca", "", "path to cert for the copilot server CA")
	flag.StringVar(&clientCert, "client-cert", "", "path to cert for the copilot client")
	flag.StringVar(&clientKey, "client-key", "", "path to key for the copilot client")

*/

const UpstreamTypeCF = "cloudfoundry"

func GetClient(address, caCert, clientCert, clientKey string) (copilot.IstioClient, error) {
	if address == "" {
		address = "127.0.0.1:9000"
	}

	caCertBytes, err := ioutil.ReadFile(caCert)
	if err != nil {
		return nil, err
	}

	rootCAs := x509.NewCertPool()
	if ok := rootCAs.AppendCertsFromPEM(caCertBytes); !ok {
		return nil, errors.New("parsing server CAs: invalid pem block")
	}

	tlsCert, err := tls.LoadX509KeyPair(clientCert, clientKey)
	if err != nil {
		return nil, err
	}

	tlsConfig := &tls.Config{
		RootCAs:      rootCAs,
		Certificates: []tls.Certificate{tlsCert},
	}

	return copilot.NewIstioClient(address, tlsConfig)
}

func GetEndpointsForUpstream(ctx context.Context, client copilot.IstioClient, u *v1.Upstream) ([]endpointdiscovery.Endpoint, error) {
	resp, err := client.Routes(ctx, new(copilotapi.RoutesRequest))
	if err != nil {
		return nil, err
	}
	return GetEndpointsFromResponse(resp, u)
}

func GetEndpointsFromResponse(resp *copilotapi.RoutesResponse, us *v1.Upstream) ([]endpointdiscovery.Endpoint, error) {

	if us.Type != UpstreamTypeCF {
		return nil, errors.New("wrong upstream type")
	}
	spec, err := DecodeUpstreamSpec(us.Spec)
	if err != nil {
		return nil, err
	}
	return convertBackendSet(resp.Backends[spec.Hostname]), nil

}

func convertBackendSet(set *copilotapi.BackendSet) []endpointdiscovery.Endpoint {
	var endpoints []endpointdiscovery.Endpoint

	for _, b := range set.Backends {
		endpoints = append(endpoints, endpointdiscovery.Endpoint{
			Address: b.Address,
			Port:    int32(b.Port),
		})
	}
	return endpoints
}

func GetUpstreams(ctx context.Context, client copilot.IstioClient) ([]*v1.Upstream, error) {
	resp, err := client.Routes(ctx, new(copilotapi.RoutesRequest))
	if err != nil {
		return nil, err
	}
	return GetUpstreamsFromResponse(resp)
}

func GetUpstreamsFromResponse(resp *copilotapi.RoutesResponse) ([]*v1.Upstream, error) {
	var uss []*v1.Upstream
	for hostname := range resp.Backends {
		uss = append(uss, &v1.Upstream{
			Name: upstreamName(hostname),
			Type: UpstreamTypeCF,
			Spec: EncodeUpstreamSpec(UpstreamSpec{
				Hostname: hostname,
			}),
		})
	}

	return uss, nil
}

func upstreamName(hostname string) string {
	return hostname
}
