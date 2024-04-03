package helper

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/solo-io/gloo/test/gomega/matchers"
	"github.com/solo-io/gloo/test/gomega/transforms"
	"github.com/solo-io/gloo/test/testutils"

	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	"github.com/solo-io/go-utils/log"
)

type CurlOpts struct {
	Protocol          string
	Path              string
	Method            string
	Host              string
	Service           string
	CaFile            string
	Body              string
	Headers           map[string]string
	Port              int
	ReturnHeaders     bool
	ConnectionTimeout int
	Verbose           bool
	LogResponses      bool
	AllowInsecure     bool
	// WithoutStats sets the -s flag to prevent download stats from printing
	WithoutStats bool
	// Optional SNI name to resolve domain to when sending request
	Sni        string
	SelfSigned bool
}

var (
	ErrCannotCurl = errors.New("cannot curl")
	errCannotCurl = func(imageName, imageTag string) error {
		return errors.Wrapf(ErrCannotCurl, "testContainer from image %s:%s", imageName, imageTag)
	}
)

func getTimeouts(timeout ...time.Duration) (currentTimeout, pollingInterval time.Duration) {
	defaultTimeout := time.Second * 20
	defaultPollingTimeout := time.Second * 5
	switch len(timeout) {
	case 0:
		currentTimeout = defaultTimeout
		pollingInterval = defaultPollingTimeout
	default:
		fallthrough
	case 2:
		pollingInterval = timeout[1]
		if pollingInterval == 0 {
			pollingInterval = defaultPollingTimeout
		}
		fallthrough
	case 1:
		currentTimeout = timeout[0]
		if currentTimeout == 0 {
			// for backwards compatability, leave this zero check
			currentTimeout = defaultTimeout
		}
	}
	return currentTimeout, pollingInterval
}

func (t *testContainer) CurlEventuallyShouldOutput(opts CurlOpts, substr string, ginkgoOffset int, timeout ...time.Duration) {
	currentTimeout, pollingInterval := getTimeouts(timeout...)

	// for some useful-ish output
	tick := time.Tick(currentTimeout / 8)

	EventuallyWithOffset(ginkgoOffset+1, func(g Gomega) {
		g.Expect(t.CanCurl()).To(BeTrue())

		var res string

		bufChan, done, err := t.CurlAsyncChan(opts)
		if err != nil {
			// trigger an early exit if the pod has been deleted
			// if we return an error here, the Eventually will continue. By making an
			// assertion with the outer context's Gomega, we can trigger a failure at
			// that outer scope.
			g.Expect(err.Error()).NotTo(ContainSubstring(`pods "testserver" not found`))
			return
		}
		defer close(done)
		var buf io.Reader
		select {
		case <-tick:
			buf = bytes.NewBufferString("waiting for reply")
		case r, ok := <-bufChan:
			if ok {
				buf = r
			}
		}
		byt, err := io.ReadAll(buf)
		if err != nil {
			res = err.Error()
		} else {
			res = string(byt)
		}
		if strings.Contains(res, substr) {
			log.GreyPrintf("success: %v", res)
		}

		g.Expect(res).To(WithTransform(transforms.WithCurlHttpResponse, matchers.HaveHttpResponse(&matchers.HttpResponse{
			Body:       ContainSubstring(substr),
			StatusCode: http.StatusOK,
		})))

	}, currentTimeout, pollingInterval).Should(Succeed())
}

func (t *testContainer) CurlEventuallyShouldRespond(opts CurlOpts, substr string, ginkgoOffset int, timeout ...time.Duration) {
	currentTimeout, pollingInterval := getTimeouts(timeout...)
	// for some useful-ish output
	tick := time.Tick(currentTimeout / 8)

	EventuallyWithOffset(ginkgoOffset+1, func(g Gomega) {
		g.Expect(t.CanCurl()).To(BeTrue())

		res, err := t.Curl(opts)
		if err != nil {
			// trigger an early exit if the pod has been deleted.
			// if we return an error here, the Eventually will continue. By making an
			// assertion with the outer context's Gomega, we can trigger a failure at
			// that outer scope.
			g.Expect(err.Error()).NotTo(ContainSubstring(`pods "testserver" not found`))
			return
		}
		select {
		default:
			break
		case <-tick:
			if opts.LogResponses {
				log.GreyPrintf("running: %v\nwant %v\nhave: %s", opts, substr, res)
			}
		}
		if strings.Contains(res, substr) && opts.LogResponses {
			log.GreyPrintf("success: %v", res)
		}

		g.Expect(res).To(WithTransform(transforms.WithCurlHttpResponse, matchers.HaveHttpResponse(&matchers.HttpResponse{
			Body:       ContainSubstring(substr),
			StatusCode: http.StatusOK,
		})))

	}, currentTimeout, pollingInterval).Should(Succeed())
}

func (t *testContainer) buildCurlArgs(opts CurlOpts) []string {
	curlRequestBuilder := testutils.DefaultCurlRequestBuilder()

	if opts.Verbose {
		curlRequestBuilder.VerboseOutput()
	}
	if opts.WithoutStats {
		curlRequestBuilder.WithoutStats()
	}
	if opts.ConnectionTimeout > 0 {
		curlRequestBuilder.WithConnectionTimeout(opts.ConnectionTimeout)
	}
	if opts.ReturnHeaders {
		curlRequestBuilder.WithReturnHeaders()
	}

	curlRequestBuilder.WithMethod(opts.Method)
	curlRequestBuilder.WithHost(opts.Host)
	curlRequestBuilder.WithCaFile(opts.CaFile)

	if opts.Body != "" {
		curlRequestBuilder.WithPostBody(opts.Body)
	}
	for h, v := range opts.Headers {
		curlRequestBuilder.WithHeader(h, v)
	}
	if opts.AllowInsecure {
		curlRequestBuilder.AllowInsecure()
	}

	port := opts.Port
	if port == 0 {
		port = 8080
	}
	curlRequestBuilder.WithPort(port)

	if opts.Protocol != "" {
		curlRequestBuilder.WithScheme(opts.Protocol)
	}

	service := opts.Service
	if service == "" {
		service = "test-ingress"
	}
	curlRequestBuilder.WithService(service)

	if opts.SelfSigned {
		curlRequestBuilder.SelfSigned()
	}
	curlRequestBuilder.WithSni(opts.Sni)

	args := curlRequestBuilder.BuildArgs()
	log.Printf("running: %v", strings.Join(args, " "))
	return args
}

func (t *testContainer) Curl(opts CurlOpts) (string, error) {
	if !t.CanCurl() {
		return "", errCannotCurl(t.containerImageName, t.imageTag)
	}

	args := t.buildCurlArgs(opts)
	return t.Exec(args...)
}

func (t *testContainer) CurlAsync(opts CurlOpts) (io.Reader, chan struct{}, error) {
	if !t.CanCurl() {
		return nil, nil, errCannotCurl(t.containerImageName, t.imageTag)
	}

	args := t.buildCurlArgs(opts)
	return t.ExecAsync(args...)
}

func (t *testContainer) CurlAsyncChan(opts CurlOpts) (<-chan io.Reader, chan struct{}, error) {
	if !t.CanCurl() {
		return nil, nil, errCannotCurl(t.containerImageName, t.imageTag)
	}

	args := t.buildCurlArgs(opts)
	return t.ExecChan(&bytes.Buffer{}, args...)
}
