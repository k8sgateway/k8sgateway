package testutils

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"net"
	"net/http"
	"time"

	"github.com/onsi/ginkgo/v2"
)

var DefaultHttpClient = &http.Client{
	Timeout: time.Second * 1,
}

type HttpClientBuilder struct {
	timeout time.Duration

	rootCaCert         string
	proxyProtocolBytes []byte
}

func DefaultClientBuilder() *HttpClientBuilder {
	return &HttpClientBuilder{
		timeout: DefaultHttpClient.Timeout,
	}
}

func (c *HttpClientBuilder) WithTimeout(timeout time.Duration) *HttpClientBuilder {
	c.timeout = timeout
	return c
}

func (c *HttpClientBuilder) WithProxyProtocolBytes(bytes []byte) *HttpClientBuilder {
	c.proxyProtocolBytes = bytes
	return c
}

func (c *HttpClientBuilder) WithTLS(rootCaCert string) *HttpClientBuilder {
	c.rootCaCert = rootCaCert
	return c
}

func (c *HttpClientBuilder) Build() *http.Client {
	var (
		client          http.Client
		tlsClientConfig *tls.Config
		dialContext     func(ctx context.Context, network, addr string) (net.Conn, error)
	)

	// If the rootCACert is provided, configure the client to use TLS
	if c.rootCaCert != "" {
		caCertPool := x509.NewCertPool()
		ok := caCertPool.AppendCertsFromPEM([]byte(c.rootCaCert))
		if !ok {
			ginkgo.Fail("CA Cert is not ok")
		}

		tlsClientConfig = &tls.Config{
			InsecureSkipVerify: false,
			ServerName:         "gateway-proxy",
			RootCAs:            caCertPool,
		}
	}

	// If the proxyProtocolBytes are provided, configure the dialContext to prepend
	// the bytes at the beginning of the connection
	// https://www.haproxy.org/download/1.9/doc/proxy-protocol.txt
	if len(c.proxyProtocolBytes) > 0 {
		dialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
			var zeroDialer net.Dialer
			connection, err := zeroDialer.DialContext(ctx, network, addr)
			if err != nil {
				return nil, err
			}

			// inject proxy protocol bytes
			// example: []byte("PROXY TCP4 1.2.3.4 1.2.3.5 443 443\r\n")
			_, err = connection.Write(c.proxyProtocolBytes)
			if err != nil {
				_ = connection.Close()
				return nil, err
			}

			return connection, nil
		}
	}

	client.Transport = &http.Transport{
		TLSClientConfig: tlsClientConfig,
		DialContext:     dialContext,
	}

	return &client
}
