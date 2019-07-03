package translator

import (
	"fmt"
	"sort"
	"time"

	knativev1alpha1 "github.com/knative/serving/pkg/apis/networking/v1alpha1"
	"github.com/pkg/errors"
	v1alpha1 "github.com/solo-io/gloo/projects/clusteringress/pkg/api/external/knative"
	v1 "github.com/solo-io/gloo/projects/clusteringress/pkg/api/v1"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/retries"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/transformation"
	"github.com/solo-io/go-utils/log"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

const (
	bindPortHttp  = 80
	bindPortHttps = 443
)

func translateProxy(namespace string, snap *v1.TranslatorSnapshot) (*gloov1.Proxy, error) {
	clusterIngresses := snap.Clusteringresses
	secrets := snap.Secrets

	virtualHostsHttp, secureVirtualHosts, err := virtualHosts(clusterIngresses, secrets)
	if err != nil {
		return nil, errors.Wrapf(err, "computing virtual hosts")
	}
	var virtualHostsHttps []*gloov1.VirtualHost
	var sslConfigs []*gloov1.SslConfig
	for _, svh := range secureVirtualHosts {
		virtualHostsHttps = append(virtualHostsHttps, svh.vh)
		sslConfigs = append(sslConfigs, &gloov1.SslConfig{
			SslSecrets: &gloov1.SslConfig_SecretRef{
				SecretRef: &svh.secret,
			},
			SniDomains: svh.vh.Domains,
		})
	}
	var listeners []*gloov1.Listener
	if len(virtualHostsHttp) > 0 {
		listeners = append(listeners, &gloov1.Listener{
			Name:        "http",
			BindAddress: "::",
			BindPort:    bindPortHttp,
			ListenerType: &gloov1.Listener_HttpListener{
				HttpListener: &gloov1.HttpListener{
					VirtualHosts: virtualHostsHttp,
				},
			},
		})
	}
	if len(virtualHostsHttps) > 0 {
		listeners = append(listeners, &gloov1.Listener{
			Name:        "https",
			BindAddress: "::",
			BindPort:    bindPortHttps,
			ListenerType: &gloov1.Listener_HttpListener{
				HttpListener: &gloov1.HttpListener{
					VirtualHosts: virtualHostsHttps,
				},
			},
			SslConfigurations: sslConfigs,
		})
	}
	return &gloov1.Proxy{
		Metadata: core.Metadata{
			Name:      "clusteringress-proxy", // must match envoy role
			Namespace: namespace,
		},
		Listeners: listeners,
	}, nil
}

type secureVirtualHost struct {
	vh     *gloov1.VirtualHost
	secret core.ResourceRef
}

func virtualHosts(ingresses v1alpha1.ClusterIngressList, secrets gloov1.SecretList) ([]*gloov1.VirtualHost, []secureVirtualHost, error) {
	routesByHostHttp := make(map[string][]*gloov1.Route)
	routesByHostHttps := make(map[string][]*gloov1.Route)
	secretsByHost := make(map[string]*core.ResourceRef)
	for _, ing := range ingresses {
		spec := ing.Spec
		for _, tls := range spec.TLS {
			secret, err := secrets.Find(ing.Namespace, tls.SecretName)
			if err != nil {
				return nil, nil, errors.Wrapf(err, "invalid secret for knative ingress %v", ing.Name)
			}

			ref := secret.Metadata.Ref()
			for _, host := range tls.Hosts {
				if existing, alreadySet := secretsByHost[host]; alreadySet {
					if existing.Name != ref.Name || existing.Namespace != ref.Namespace {
						log.Warnf("a TLS secret for host %v was redefined in knative ingress %v, ignoring", ing.Name)
						continue
					}
				}
				secretsByHost[host] = &ref
			}
		}

		for i, rule := range spec.Rules {
			if rule.HTTP == nil {
				log.Warnf("rule %v in knative ingress %v is missing HTTP field", i, ing.Name)
				continue
			}
			for _, route := range rule.HTTP.Paths {
				pathRegex := route.Path
				if pathRegex == "" {
					pathRegex = ".*"
				}
				appendHeaders := make(map[string]*transformation.InjaTemplate)
				for name, value := range route.AppendHeaders {
					appendHeaders[name] = &transformation.InjaTemplate{Text: value}
				}
				var timeout *time.Duration
				if route.Timeout != nil {
					timeout = &route.Timeout.Duration
				}
				var retryPolicy *retries.RetryPolicy
				if route.Retries != nil {
					var perTryTimeout *time.Duration
					if route.Retries.PerTryTimeout != nil {
						perTryTimeout = &route.Retries.PerTryTimeout.Duration
					}
					retryPolicy = &retries.RetryPolicy{
						NumRetries:    uint32(route.Retries.Attempts),
						PerTryTimeout: perTryTimeout,
					}
				}
				appendHeadersTransformation := &transformation.RouteTransformations{
					RequestTransformation: &transformation.Transformation{
						TransformationType: &transformation.Transformation_TransformationTemplate{
							TransformationTemplate: &transformation.TransformationTemplate{
								Headers: appendHeaders,
								BodyTransformation: &transformation.TransformationTemplate_Passthrough{
									Passthrough: &transformation.Passthrough{},
								},
							},
						},
					},
				}

				action, err := routeActionFromSplits(route.Splits)
				if err != nil {
					return nil, nil, errors.Wrapf(err, "")
				}

				route := &gloov1.Route{
					Matcher: &gloov1.Matcher{
						PathSpecifier: &gloov1.Matcher_Regex{
							Regex: pathRegex,
						},
					},
					Action: &gloov1.Route_RouteAction{
						RouteAction: action,
					},
					RoutePlugins: &gloov1.RoutePlugins{
						Transformations: appendHeadersTransformation,
						Timeout:         timeout,
						Retries:         retryPolicy,
					},
				}
				for _, host := range rule.Hosts {
					if _, useTls := secretsByHost[host]; useTls {
						routesByHostHttps[host] = append(routesByHostHttps[host], route)
					} else {
						routesByHostHttp[host] = append(routesByHostHttp[host], route)
					}
				}
			}
		}
	}

	var virtualHostsHttp []*gloov1.VirtualHost
	var virtualHostsHttps []secureVirtualHost

	// TODO (ilackarms): support for VirtualHostPlugins on ingress?
	for host, routes := range routesByHostHttp {
		sortByLongestPathName(routes)
		virtualHostsHttp = append(virtualHostsHttp, &gloov1.VirtualHost{
			Name:    host + "-http",
			Domains: []string{host, fmt.Sprintf("%v:%v", host, bindPortHttp)},
			Routes:  routes,
		})
	}

	for host, routes := range routesByHostHttps {
		sortByLongestPathName(routes)
		secret, ok := secretsByHost[host]
		if !ok {
			return nil, nil, errors.Errorf("internal error: secret not found for host %v after processing knative ingresses", host)
		}
		virtualHostsHttps = append(virtualHostsHttps, secureVirtualHost{
			vh: &gloov1.VirtualHost{
				Name:    host + "-http",
				Domains: []string{host, fmt.Sprintf("%v:%v", host, bindPortHttps)},
				Routes:  routes,
			},
			secret: *secret,
		})
	}

	sort.SliceStable(virtualHostsHttp, func(i, j int) bool {
		return virtualHostsHttp[i].Name < virtualHostsHttp[j].Name
	})
	sort.SliceStable(virtualHostsHttps, func(i, j int) bool {
		return virtualHostsHttps[i].vh.Name < virtualHostsHttps[j].vh.Name
	})
	return virtualHostsHttp, virtualHostsHttps, nil
}

func routeActionFromSplits(splits []knativev1alpha1.IngressBackendSplit) (*gloov1.RouteAction, error) {
	switch len(splits) {
	case 0:
		return nil, errors.Errorf("invalid cluster ingress: must provide at least 1 split")
	case 1:
		split := splits[0]
		return &gloov1.RouteAction{
			Destination: &gloov1.RouteAction_Single{
				Single: &gloov1.Destination{
					DestinationType: serviceForSplit(split),
				},
			},
		}, nil
	}

	var destinations []*gloov1.WeightedDestination
	for _, split := range splits {
		destinations = append(destinations, &gloov1.WeightedDestination{
			Destination: &gloov1.Destination{
				DestinationType: serviceForSplit(split),
			},
			Weight: uint32(split.Percent),
		})
	}
	return &gloov1.RouteAction{
		Destination: &gloov1.RouteAction_Multi{
			Multi: &gloov1.MultiDestination{
				Destinations: destinations,
			},
		},
	}, nil
}

func serviceForSplit(split knativev1alpha1.IngressBackendSplit) *gloov1.Destination_Service {
	return &gloov1.Destination_Service{
		Service: &gloov1.ServiceDestination{
			Ref:  core.ResourceRef{Name: split.ServiceName, Namespace: split.ServiceNamespace},
			Port: uint32(split.ServicePort.IntValue()),
		},
	}
}

func sortByLongestPathName(routes []*gloov1.Route) {
	sort.SliceStable(routes, func(i, j int) bool {
		return routes[i].Matcher.PathSpecifier.(*gloov1.Matcher_Regex).Regex > routes[j].Matcher.PathSpecifier.(*gloov1.Matcher_Regex).Regex
	})
}
