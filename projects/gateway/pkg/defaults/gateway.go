package defaults

import (
	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/defaults"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

func DefaultGateway(writeNamespace string) *v1.Gateway {
	return &v1.Gateway{
		Metadata: core.Metadata{
			Name:      "gateway",
			Namespace: writeNamespace,
		},
		BindAddress: "::",
		BindPort:    defaults.HttpPort,
		// all virtualservices
	}
}

func DefaultVirtualService(namespace, name string) *v1.VirtualService {
	return &v1.VirtualService{
		Metadata: core.Metadata{
			Name:      name,
			Namespace: namespace,
		},
		VirtualHost: &gloov1.VirtualHost{
			Name:    "routes",
			Domains: []string{"*"},
			Routes: []*gloov1.Route{{
				Matcher: &gloov1.Matcher{
					PathSpecifier: &gloov1.Matcher_Prefix{Prefix: "/"},
				},
				Action: &gloov1.Route_DirectResponseAction{DirectResponseAction: &gloov1.DirectResponseAction{
					Status: 200,
					Body: `Gloo and Envoy are configured correctly!

Delete the '` + name + ` Virtual Service to get started. 
`,
				}},
			}},
		},
	}
}

func LocalUpstream(writeNamespace string) *gloov1.Upstream {
	return &gloov1.Upstream{
		Metadata: core.Metadata{
			Name:      "upstream",
			Namespace: writeNamespace,
		},
		UpstreamSpec: &gloov1.UpstreamSpec{},
	}
}
