package v1helpers

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"

	"github.com/gogo/protobuf/proto"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	static_plugin_gloo "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/static"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

type ReceivedRequest struct {
	Method      string
	Body        []byte
	Host        string
	GRPCRequest proto.Message
}

func NewTestHttpUpstream(ctx context.Context, addr string) *TestUpstream {
	backendport, responses := RunTestServer(ctx)
	return newTestUpstream(addr, backendport, responses)
}

type TestUpstream struct {
	Upstream *gloov1.Upstream
	C        <-chan *ReceivedRequest
	Address  string
	Port     uint32
}

var id = 0

func newTestUpstream(addr string, port uint32, responses <-chan *ReceivedRequest) *TestUpstream {

	id += 1
	u := &gloov1.Upstream{
		Metadata: core.Metadata{
			Name:      fmt.Sprintf("local-%d", id),
			Namespace: "default",
		},
		UpstreamSpec: &gloov1.UpstreamSpec{
			UpstreamType: &gloov1.UpstreamSpec_Static{
				Static: &static_plugin_gloo.UpstreamSpec{
					Hosts: []*static_plugin_gloo.Host{{
						Addr: addr,
						Port: port,
					}},
				},
			},
		},
	}

	return &TestUpstream{
		Upstream: u,
		C:        responses,
		Address:  fmt.Sprintf("%s:%d", addr, port),
		Port:     port,
	}
}

func RunTestServer(ctx context.Context) (uint32, <-chan *ReceivedRequest) {
	bodychan := make(chan *ReceivedRequest, 100)
	handlerfunc := func(rw http.ResponseWriter, r *http.Request) {
		var rr ReceivedRequest
		rr.Method = r.Method
		if r.Body != nil {
			body, _ := ioutil.ReadAll(r.Body)
			r.Body.Close()
			if len(body) != 0 {
				rr.Body = body
				rw.Write(body)
			}
		}

		rr.Host = r.Host

		bodychan <- &rr
	}

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}

	addr := listener.Addr().String()
	_, portstr, err := net.SplitHostPort(addr)
	if err != nil {
		panic(err)
	}

	port, err := strconv.Atoi(portstr)
	if err != nil {
		panic(err)
	}

	handler := http.HandlerFunc(handlerfunc)
	go func() {
		defer GinkgoRecover()
		h := &http.Server{Handler: handler}
		go func() {
			defer GinkgoRecover()
			if err := h.Serve(listener); err != nil {
				if err != http.ErrServerClosed {
					panic(err)
				}
			}
		}()

		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		h.Shutdown(ctx)
		cancel()
		// close channel, the http handler may panic but this should be caught by the http code.
		close(bodychan)
	}()
	return uint32(port), bodychan
}

func TestUpstremReachable(envoyPort uint32, tu *TestUpstream, rootca *string) {

	body := []byte("solo.io test")

	EventuallyWithOffset(1, func() error {
		// send a request with a body
		var buf bytes.Buffer
		buf.Write(body)

		var client http.Client

		scheme := "http"
		if rootca != nil {
			scheme = "https"
			caCertPool := x509.NewCertPool()
			ok := caCertPool.AppendCertsFromPEM([]byte(*rootca))
			if !ok {
				return fmt.Errorf("ca cert is not OK")
			}

			client.Transport = &http.Transport{
				TLSClientConfig: &tls.Config{
					RootCAs:            caCertPool,
					InsecureSkipVerify: true,
				},
			}
		}

		res, err := client.Post(fmt.Sprintf(scheme+"://%s:%d/1", "localhost", envoyPort), "application/octet-stream", &buf)
		if err != nil {
			return err
		}
		if res.StatusCode != http.StatusOK {
			return fmt.Errorf("%v is not OK", res.StatusCode)
		}
		return nil
	}, "10s", ".5s").Should(BeNil())

	EventuallyWithOffset(1, tu.C).Should(Receive(PointTo(MatchFields(IgnoreExtras, Fields{
		"Method": Equal("POST"),
		"Body":   Equal(body),
	}))))
}
