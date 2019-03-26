package generate

type Config struct {
	Namespace      *Namespace              `json:"namespace,omitempty"`
	Rbac           *Rbac                   `json:"rbac,omitempty"`
	Crds           *Crds                   `json:"crds,omitempty"`
	Settings       *Settings               `json:"settings,omitempty"`
	Gloo           *Gloo                   `json:"gloo,omitempty"`
	Discovery      *Discovery              `json:"discovery,omitempty"`
	Gateway        *Gateway                `json:"gateway,omitempty"`
	GatewayProxies map[string]GatewayProxy `json:"gatewayProxies,omitempty"`
	Ingress        *Ingress                `json:"ingress,omitempty"`
	IngressProxy   *IngressProxy           `json:"ingressProxy,omitempty"`
}

type Namespace struct {
	Create bool `json:"create"`
}

type Rbac struct {
	Create bool `json:"create"`
}

type Crds struct {
	Create bool `json:"create"`
}

// Common
type Image struct {
	Tag        string `json:"tag"`
	Repository string `json:"repository"`
	PullPolicy string `json:"pullPolicy"`
	PullSecret string `json:"pullSecret,omitempty"`
}

type DeploymentSpec struct {
	Replicas int  `json:"replicas"`
	Stats    bool `json:"stats"`
}

type Integrations struct {
	Knative *Knative `json:"knative"`
}
type Knative struct {
	Enabled *bool         `json:"enabled"`
	Proxy   *KnativeProxy `json:"proxy,omitempty"`
}

type KnativeProxy struct {
	Image     *Image `json:"image,omitempty"`
	HttpPort  string `json:"httpPort,omitempty"`
	HttpsPort string `json:"httpsPort,omitempty"`
	*DeploymentSpec
}

type Settings struct {
	WatchNamespaces []string      `json:"watchNamespaces,omitempty"`
	WriteNamespace  string        `json:"writeNamespace,omitempty"`
	Integrations    *Integrations `json:"integrations,omitempty"`
	Create          bool          `json:"create,omitempty"`
	Extensions      interface{}   `json:"extensions,omitempty"`
}

type Gloo struct {
	Deployment *GlooDeployment `json:"deployment,omitempty"`
}

type GlooDeployment struct {
	Image   *Image `json:"image,omitempty"`
	XdsPort string `json:"xdsPort,omitempty"`
	*DeploymentSpec
}

type Discovery struct {
	Deployment *DiscoveryDeployment `json:"deployment,omitempty"`
}

type DiscoveryDeployment struct {
	Image *Image `json:"image,omitempty"`
	*DeploymentSpec
}

type Gateway struct {
	Enabled    *bool              `json:"enabled"`
	Deployment *GatewayDeployment `json:"deployment,omitempty"`
}

type GatewayDeployment struct {
	Image *Image `json:"image,omitempty"`
	*DeploymentSpec
}

type GatewayProxy struct {
	Deployment *GatewayProxyDeployment `json:"deployment,omitempty"`
	ConfigMap  *GatewayProxyConfigMap  `json:"configMap,omitempty"`
	Service    *GatewayProxyService    `json:"service,omitempty"`
}

type GatewayProxyDeployment struct {
	Image            *Image            `json:"image,omitempty"`
	HttpPort         string            `json:"httpPort,omitempty"`
	HttpsPort        string            `json:"httpsPort,omitempty"`
	ExtraPorts       []interface{}     `json:"extraPorts,omitempty"`
	ExtraAnnotations map[string]string `json:"extraAnnotations,omitempty"`
	*DeploymentSpec
}
type GatewayProxyService struct {
	Type             string            `json:"type,omitempty"`
	HttpPort         string            `json:"httpPort,omitempty"`
	HttpsPort        string            `json:"httpsPort,omitempty"`
	ClusterIP        string            `json:"clusterIP,omitempty"`
	ExtraAnnotations map[string]string `json:"extraAnnotations,omitempty"`
	ExternalTrafficPolicy string       `json:"externalTrafficPolicy,omitempty"`
}

type GatewayProxyConfigMap struct {
	Data map[string]string `json:"data"`
}

type Ingress struct {
	Enabled    *bool              `json:"enabled"`
	Deployment *IngressDeployment `json:"deployment,omitempty"`
}

type IngressDeployment struct {
	Image *Image `json:"image,omitempty"`
	*DeploymentSpec
}

type IngressProxy struct {
	Deployment *IngressProxyDeployment `json:"deployment,omitempty"`
	ConfigMap  *IngressProxyConfigMap  `json:"configMap,omitempty"`
}

type IngressProxyDeployment struct {
	Image            *Image            `json:"image,omitempty"`
	HttpPort         string            `json:"httpPort,omitempty"`
	HttpsPort        string            `json:"httpsPort,omitempty"`
	ExtraPorts       []interface{}     `json:"extraPorts,omitempty"`
	ExtraAnnotations map[string]string `json:"extraAnnotations,omitempty"`
	*DeploymentSpec
}

type IngressProxyConfigMap struct {
	Data map[string]string `json:"data,omitempty"`
}
