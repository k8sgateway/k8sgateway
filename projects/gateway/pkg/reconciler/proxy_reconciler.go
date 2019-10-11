package reconciler

import (
	"context"
	"sort"

	"github.com/solo-io/go-utils/contextutils"

	"github.com/solo-io/gloo/projects/gateway/pkg/reporting"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/grpc/validation"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"

	"github.com/solo-io/gloo/projects/gateway/pkg/utils"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v2/reporter"
	"github.com/solo-io/solo-kit/pkg/errors"
)

type GeneratedProxies map[*gloov1.Proxy]reporter.ResourceReports

type ProxyReconciler interface {
	ReconcileProxies(ctx context.Context, proxiesToWrite GeneratedProxies, writeNamespace string, labels map[string]string) error
}

type proxyReconciler struct {
	proxyValidator validation.ProxyValidationServiceClient
	baseReconciler gloov1.ProxyReconciler
}

func NewProxyReconciler(proxyValidator validation.ProxyValidationServiceClient, proxyClient gloov1.ProxyClient) *proxyReconciler {
	return &proxyReconciler{proxyValidator: proxyValidator, baseReconciler: gloov1.NewProxyReconciler(proxyClient)}
}

const proxyValidationErrMsg = "internal err: communication with proxy validation (gloo) failed"

func (s *proxyReconciler) ReconcileProxies(ctx context.Context, proxiesToWrite GeneratedProxies, writeNamespace string, labels map[string]string) error {
	if err := s.addProxyValidationResults(ctx, proxiesToWrite); err != nil {
		return errors.Wrapf(err, "failed to add proxy validation results to reports")
	}

	proxiesToWrite, err := stripInvalidListenersAndVirtualHosts(proxiesToWrite)
	if err != nil {
		return err
	}

	var allProxies gloov1.ProxyList
	for proxy := range proxiesToWrite {
		allProxies = append(allProxies, proxy)
	}

	sort.SliceStable(allProxies, func(i, j int) bool {
		return allProxies[i].Metadata.Less(allProxies[j].Metadata)
	})

	if err := s.baseReconciler.Reconcile(writeNamespace, allProxies, transitionFunc(proxiesToWrite), clients.ListOpts{
		Ctx:      ctx,
		Selector: labels,
	}); err != nil {
		return err
	}

	return nil
}

func transitionFunc(proxiesToWrite GeneratedProxies) gloov1.TransitionProxyFunc {
	return func(original, desired *gloov1.Proxy) (b bool, e error) {
		// if any listeners from the old proxy were rejected in the new reports, preserve those
		if err := forEachListener(original, proxiesToWrite[desired], func(listener *gloov1.Listener, accepted bool) {
			// old listener was rejected, preserve it on the desired proxy
			if !accepted {
				desired.Listeners = append(desired.Listeners, listener)
			}
		}); err != nil {
			// should never happen
			return false, err
		}

		// preserve previous vhosts if new vservice was errored
		for _, desiredListener := range desired.Listeners {

			desiredHttpListener := desiredListener.GetHttpListener()
			if desiredHttpListener == nil {
				continue
			}

			// find the original listener by its name
			// if it does not exist in the original, skip
			var originalListener *gloov1.Listener
			for _, origLis := range original.Listeners {
				if origLis.Name == desiredListener.Name {
					originalListener = origLis
					break
				}
			}
			if originalListener == nil {
				continue
			}

			// find any rejected vhosts on the original listener and copy them over
			if err := forEachVhost(originalListener, proxiesToWrite[desired], func(vhost *gloov1.VirtualHost, accepted bool) {
				// old vhost was rejected, preserve it on the desired proxy
				if !accepted {
					desiredHttpListener.VirtualHosts = append(desiredHttpListener.VirtualHosts, vhost)
				}
			}); err != nil {
				// should never happen
				return false, err
			}

			sort.SliceStable(desiredHttpListener.VirtualHosts, func(i, j int) bool {
				return desiredHttpListener.VirtualHosts[i].Name < desiredHttpListener.VirtualHosts[j].Name
			})

		}

		sort.SliceStable(desired.Listeners, func(i, j int) bool {
			return desired.Listeners[i].Name < desired.Listeners[j].Name
		})

		return utils.TransitionFunction(original, desired)
	}
}

// validate generated proxies and add reports for the owner resources
// this function makes a gRPC call to gloo validation server
func (s *proxyReconciler) addProxyValidationResults(ctx context.Context, proxiesToWrite GeneratedProxies) error {

	if s.proxyValidator == nil {
		contextutils.LoggerFrom(ctx).Warnf("proxy validation is not configured, skipping proxy validation check")
		return nil
	}

	for proxy, reports := range proxiesToWrite {

		proxyRpt, err := s.proxyValidator.ValidateProxy(ctx, &validation.ProxyValidationServiceRequest{
			Proxy: proxy,
		})
		if err != nil {
			return errors.Wrapf(err, proxyValidationErrMsg)
		}

		// add the proxy validation result to the existing resource reports
		if err := reporting.AddProxyValidationResult(reports, proxy, proxyRpt.GetProxyReport()); err != nil {
			//should never happen
			return err
		}
	}

	return nil
}

func stripInvalidListenersAndVirtualHosts(proxiesToWrite GeneratedProxies) (GeneratedProxies, error) {
	strippedProxies := GeneratedProxies{}

	for proxy, reports := range proxiesToWrite {

		// clone because mutations occur
		proxy := resources.Clone(proxy).(*gloov1.Proxy)

		var validListeners []*gloov1.Listener

		if err := forEachListener(proxy, reports, func(listener *gloov1.Listener, accepted bool) {
			if accepted {
				validListeners = append(validListeners, listener)
			}
		}); err != nil {
			return nil, err
		}

		for _, lis := range proxy.Listeners {

			if httpListenerType, ok := lis.ListenerType.(*gloov1.Listener_HttpListener); ok {
				var validVhosts []*gloov1.VirtualHost

				if err := forEachVhost(lis, reports, func(vhost *gloov1.VirtualHost, accepted bool) {
					if accepted {
						validVhosts = append(validVhosts, vhost)
					}
				}); err != nil {
					return nil, err
				}

				sort.SliceStable(validVhosts, func(i, j int) bool {
					return validVhosts[i].Name < validVhosts[j].Name
				})

				httpListenerType.HttpListener.VirtualHosts = validVhosts
			}
		}

		sort.SliceStable(validListeners, func(i, j int) bool {
			return validListeners[i].Name < validListeners[j].Name
		})

		proxy.Listeners = validListeners

		// update the map with the copy
		strippedProxies[proxy] = reports
	}

	return strippedProxies, nil
}

func forEachListener(proxy *gloov1.Proxy, reports reporter.ResourceReports, fn func(*gloov1.Listener, bool)) error {
	for _, lis := range proxy.Listeners {
		accepted, err := reporting.AllSourcesAccepted(reports, lis)
		if err != nil {
			return err
		}

		fn(lis, accepted)
	}
	return nil
}

func forEachVhost(lis *gloov1.Listener, reports reporter.ResourceReports, fn func(*gloov1.VirtualHost, bool)) error {
	if httpListenerType, ok := lis.ListenerType.(*gloov1.Listener_HttpListener); ok {

		for _, vhost := range httpListenerType.HttpListener.GetVirtualHosts() {
			accepted, err := reporting.AllSourcesAccepted(reports, vhost)
			if err != nil {
				return err
			}

			fn(vhost, accepted)
		}
	}
	return nil
}
