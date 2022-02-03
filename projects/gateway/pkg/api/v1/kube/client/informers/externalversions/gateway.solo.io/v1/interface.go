/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/solo-io/gloo/projects/gateway/pkg/api/v1/kube/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Gateways returns a GatewayInformer.
	Gateways() GatewayInformer
	// MatchableHttpGateways returns a MatchableHttpGatewayInformer.
	MatchableHttpGateways() MatchableHttpGatewayInformer
	// RouteOptions returns a RouteOptionInformer.
	RouteOptions() RouteOptionInformer
	// RouteTables returns a RouteTableInformer.
	RouteTables() RouteTableInformer
	// VirtualHostOptions returns a VirtualHostOptionInformer.
	VirtualHostOptions() VirtualHostOptionInformer
	// VirtualServices returns a VirtualServiceInformer.
	VirtualServices() VirtualServiceInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Gateways returns a GatewayInformer.
func (v *version) Gateways() GatewayInformer {
	return &gatewayInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MatchableHttpGateways returns a MatchableHttpGatewayInformer.
func (v *version) MatchableHttpGateways() MatchableHttpGatewayInformer {
	return &matchableHttpGatewayInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// RouteOptions returns a RouteOptionInformer.
func (v *version) RouteOptions() RouteOptionInformer {
	return &routeOptionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// RouteTables returns a RouteTableInformer.
func (v *version) RouteTables() RouteTableInformer {
	return &routeTableInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VirtualHostOptions returns a VirtualHostOptionInformer.
func (v *version) VirtualHostOptions() VirtualHostOptionInformer {
	return &virtualHostOptionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VirtualServices returns a VirtualServiceInformer.
func (v *version) VirtualServices() VirtualServiceInformer {
	return &virtualServiceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
