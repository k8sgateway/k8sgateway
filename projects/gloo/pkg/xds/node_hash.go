package xds

import (
	"strings"

	envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	"github.com/solo-io/solo-kit/pkg/api/v1/control-plane/cache"
)

var _ cache.NodeHash = new(nodeRoleHasher)

const (
	// FallbackNodeCacheKey is used to let nodes know they have a bad config
	// we assign a "fix me" snapshot for bad nodes
	FallbackNodeCacheKey = "misconfigured-node"

	// KeyDelimiter is the character used to join segments of a cache key
	KeyDelimiter = "~"

	RoleKey = "role"
)

// OwnerNamespaceNameID returns the string identifier for an Envoy node in a provided namespace.
// Envoy proxies are assigned their configuration by Gloo based on their Node ID.
// Therefore, proxies must identify themselves using the same naming
// convention that we use to persist the Proxy resource in the snapshot cache.
// The naming convention that we follow is "OWNER~NAMESPACE~NAME"
func OwnerNamespaceNameID(owner, namespace, name string) string {
	return strings.Join([]string{owner, namespace, name}, KeyDelimiter)
}

func NewNodeRoleHasher() *nodeRoleHasher {
	return &nodeRoleHasher{}
}

// nodeRoleHasher identifies a node based on the values provided in the `node.metadata.role`
type nodeRoleHasher struct{}

func (h *nodeRoleHasher) ID(node *envoy_config_core_v3.Node) string {
	if node.GetMetadata() != nil {
		roleValue := node.GetMetadata().GetFields()[RoleKey]
		if roleValue != nil {
			return roleValue.GetStringValue()
		}
	}

	return FallbackNodeCacheKey
}
