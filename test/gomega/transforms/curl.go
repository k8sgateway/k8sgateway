package transforms

import (
	"bytes"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/solo-io/gloo/pkg/utils/kubeutils/kubectl"
)

const (
	requestHeaderPrefix  = "> "
	responseHeaderPrefix = "< "
	infoPrefix           = "* "
	bytesDataSuffix      = " bytes data]"
	responseStatusPrefix = "< HTTP/1.1 "
)

// WithCurlHttpResponse is a Gomega Transform that converts the string return by an exec.Curl
// and transforms it into an http.Response. This is useful to be used in tandem with matchers.HaveHttpResponse
// NOTE: This is not feature complete, as we do not convert the entire response.
// For now, we handle HTTP/1.1 response headers, status, and body.
// The curl must be executed with verbose=true to include both the response headers/status
// and response body.
func WithCurlHttpResponse(curlResponse string) *http.Response {
	headers := make(http.Header)
	statusCode := 0
	var bodyBuf bytes.Buffer

	found_non_body_line := false
	for _, line := range strings.Split(curlResponse, "\n") {
		k, v := processResponseHeader(line)
		if k != "" {
			found_non_body_line = true
			headers.Add(k, v)
			continue
		}

		code := processResponseCode(line)
		if code != 0 {
			found_non_body_line = true
			statusCode = code
			continue
		}

		// Once we've found a line that is a header or status code, we can assume we are done with the body
		if isResponseBody(line) && !found_non_body_line {
			if bodyBuf.Len() > 0 {
				bodyBuf.WriteString("\n")
			}
			bodyBuf.WriteString(line)
		}
	}

	return &http.Response{
		StatusCode: statusCode,
		Header:     headers,
		Body:       bytesBody(bodyBuf.Bytes()),
	}
}

func WithCurlResponse(curlResponse *kubectl.CurlResponse) *http.Response {
	headers := make(http.Header)
	statusCode := 0
	var bodyBuf bytes.Buffer

	// Headers/response code
	for _, line := range strings.Split(curlResponse.Headers, "\n") {
		k, v := processResponseHeader(line)
		if k != "" {
			headers.Add(k, v)
			continue
		}

		code := processResponseCode(line)
		if code != 0 {
			statusCode = code
		}
	}

	// Body
	bodyBuf.WriteString(curlResponse.Body)

	return &http.Response{
		StatusCode: statusCode,
		Header:     headers,
		Body:       bytesBody(bodyBuf.Bytes()),
	}
}

func WithCurlHttpResponseFixHeaders(curlResponse string) *http.Response {
	headers := make(http.Header)
	statusCode := 0
	var bodyBuf bytes.Buffer

	for _, line := range strings.Split(curlResponse, "\n") {
		k, v := processResponseHeader(line)
		if k != "" {
			headers.Add(k, v)
			continue
		}

		code := processResponseCode(line)
		if code != 0 {
			statusCode = code
			continue
		}

		if isResponseBody(line) {
			if bodyBuf.Len() > 0 {
				bodyBuf.WriteString("\n")
			}
			bodyBuf.WriteString(line)
		}
	}

	return &http.Response{
		StatusCode: statusCode,
		Header:     headers,
		Body:       bytesBody(bodyBuf.Bytes()),
	}
}

// processResponseHeader processes the current line if it's a response header.
// Returns header key and value if the line was processed, otherwise returns empty strings.
func processResponseHeader(line string) (string, string) {
	// check for response headers
	if strings.HasPrefix(line, responseHeaderPrefix) {
		headerParts := strings.Split(line[len(responseHeaderPrefix):], ": ")
		if len(headerParts) == 2 {
			// strip "\r" from the end of the value
			return strings.ToLower(headerParts[0]), strings.TrimSuffix(headerParts[1], "\r")
		}
	}
	return "", ""
}

// processResponseCode processes the current line if it's a response status code.
// Returns the status code if the line was processed, otherwise returns 0.
func processResponseCode(line string) int {
	// check for response status. the line with the response code will be in the format
	// `< HTTP/1.1 <code> <message>`
	if strings.HasPrefix(line, responseStatusPrefix) {
		restOfLine := line[len(responseStatusPrefix):]
		statusParts := strings.Split(restOfLine, " ")
		if len(statusParts) > 0 {
			statusCode, err := strconv.Atoi(statusParts[0])
			if err == nil {
				return statusCode
			}
		}
	}
	return 0
}

// isResponseBody returns true if the current line is part of the response body, false otherwise.
func isResponseBody(line string) bool {
	// if there is no special prefix/suffix, assume this is part of the response body
	// (this may not work reliably for all curl outputs)
	return !strings.HasPrefix(line, infoPrefix) &&
		!strings.HasPrefix(line, requestHeaderPrefix) &&
		!strings.HasPrefix(line, responseHeaderPrefix) &&
		!strings.HasSuffix(line, bytesDataSuffix)
}

func bytesBody(bodyBytes []byte) io.ReadCloser {
	return io.NopCloser(bytes.NewReader(bodyBytes))
}
