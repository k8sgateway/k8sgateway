package translator

import (
	v3 "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/config/core/v3"

	"github.com/golang/protobuf/ptypes/wrappers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
)

var _ = Describe("matcherID", func() {
	When("matchers are identical", func() {
		matcher1 := &v1.Matcher{
			SourcePrefixRanges: []*v3.CidrRange{
				{
					AddressPrefix: "foo",
					PrefixLen: &wrappers.UInt32Value{
						Value: 123,
					},
				},
				{
					AddressPrefix: "bar",
					PrefixLen: &wrappers.UInt32Value{
						Value: 456,
					},
				},
			},
			SslConfig: &v1.SslConfig{
				SniDomains: []string{"abc", "def"},
			},
		}
		matcher2 := &v1.Matcher{
			SourcePrefixRanges: []*v3.CidrRange{
				{
					AddressPrefix: "foo",
					PrefixLen: &wrappers.UInt32Value{
						Value: 123,
					},
				},
				{
					AddressPrefix: "bar",
					PrefixLen: &wrappers.UInt32Value{
						Value: 456,
					},
				},
			},
			SslConfig: &v1.SslConfig{
				SniDomains: []string{"abc", "def"},
			},
		}
		It("produces the same ID", func() {
			Expect(matcherID(matcher1)).To(Equal(matcherID(matcher2)))
		})
	})
	When("matchers are equivalent", func() {
		matcher1 := &v1.Matcher{
			SourcePrefixRanges: []*v3.CidrRange{
				{
					AddressPrefix: "bar",
					PrefixLen: &wrappers.UInt32Value{
						Value: 456,
					},
				},
				{
					AddressPrefix: "foo",
					PrefixLen: &wrappers.UInt32Value{
						Value: 123,
					},
				},
			},
			SslConfig: &v1.SslConfig{
				SniDomains: []string{"def","abc"},
			},
		}
		matcher2 := &v1.Matcher{
			SourcePrefixRanges: []*v3.CidrRange{
				{
					AddressPrefix: "foo",
					PrefixLen: &wrappers.UInt32Value{
						Value: 123,
					},
				},
				{
					AddressPrefix: "bar",
					PrefixLen: &wrappers.UInt32Value{
						Value: 456,
					},
				},
			},
			SslConfig: &v1.SslConfig{
				SniDomains: []string{"abc", "def"},
			},
		}
		It("produces the same ID", func() {
			Expect(matcherID(matcher1)).To(Equal(matcherID(matcher2)))
		})
	})
	When("matchers are different", func() {
		matcher1 := &v1.Matcher{
			SourcePrefixRanges: []*v3.CidrRange{
				{
					AddressPrefix: "foo",
					PrefixLen: &wrappers.UInt32Value{
						Value: 123,
					},
				},
			},
			SslConfig: &v1.SslConfig{
				SniDomains: []string{"abc"},
			},
		}
		matcher2 := &v1.Matcher{
			SourcePrefixRanges: []*v3.CidrRange{
				{
					AddressPrefix: "foo",
					PrefixLen: &wrappers.UInt32Value{
						Value: 123,
					},
				},
			},
			SslConfig: &v1.SslConfig{
				SniDomains: []string{"def"},
			},
		}
		It("produces different IDs", func() {
			Expect(matcherID(matcher1)).NotTo(Equal(matcherID(matcher2)))
		})
	})
})
