package matchers

import (
	"github.com/onsi/gomega"
	"github.com/onsi/gomega/gstruct"
	"github.com/onsi/gomega/types"
	"github.com/solo-io/gloo/test/gomega/transforms"
	"net/http"
)

// ContainHeaders produces a matcher that will only match if all provided headers
// are completely accounted for, including multi-value headers.
func ContainHeaders(headers http.Header) types.GomegaMatcher {
	if headers == nil {
		// If no headers are defined, we create a matcher that always succeeds
		// If we do not this we will create an And matcher for 0 objects, which leads to a panic
		return gstruct.Ignore()
	}
	headerMatchers := make([]types.GomegaMatcher, 0, len(headers))
	for k, v := range headers {
		vals := make([]interface{}, len(v))
		for i := range v {
			vals[i] = v[i]
		}
		headerMatchers = append(headerMatchers, gomega.WithTransform(transforms.WithHeaderValues(k), gomega.ContainElements(vals...)))
	}
	return gomega.And(headerMatchers...)
}
