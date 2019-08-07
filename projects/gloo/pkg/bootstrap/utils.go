package bootstrap

import (
	"context"
	"path/filepath"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/cache"

	consulapi "github.com/hashicorp/consul/api"
	vaultapi "github.com/hashicorp/vault/api"
	kubeconverters "github.com/solo-io/gloo/projects/gloo/pkg/api/converters/kube"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/go-utils/kubeutils"
	"github.com/solo-io/solo-kit/pkg/api/external/kubernetes/service"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/memory"
	skkube "github.com/solo-io/solo-kit/pkg/api/v1/resources/common/kubernetes"
	"github.com/solo-io/solo-kit/pkg/errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// used for vault and consul key-value storage
const DefaultRootKey = "gloo"

type ConfigFactoryParams struct {
	settings *v1.Settings
	memory   configFactoryParamsMemory
	kube     configFactoryParamsKube
	consul   configFactoryParamsConsul
}

func NewConfigFactoryParams(settings *v1.Settings,
	sharedCache memory.InMemoryResourceCache,
	cache kube.SharedCache,
	cfg **rest.Config,
	consulClient *consulapi.Client) ConfigFactoryParams {
	return ConfigFactoryParams{
		settings: settings,
		memory: configFactoryParamsMemory{
			sharedCache: sharedCache,
		},
		kube: configFactoryParamsKube{
			kubeCache: cache,
			restCfg:   cfg,
		},
		consul: configFactoryParamsConsul{
			consulClient: consulClient,
		},
	}
}

type configFactoryParamsMemory struct {
	sharedCache memory.InMemoryResourceCache
}

type configFactoryParamsKube struct {
	kubeCache kube.SharedCache
	restCfg   **rest.Config
}

type configFactoryParamsConsul struct {
	consulClient *consulapi.Client
}

// sharedCache, resourceCrd+cfg OR consulClient must be non-nil
func ConfigFactoryForSettings(params ConfigFactoryParams, resourceCrd crd.Crd) (factory.ResourceClientFactory, error) {
	settings := params.settings

	if settings.ConfigSource == nil {
		sharedCache := params.memory.sharedCache
		if sharedCache == nil {
			return nil, errors.Errorf("internal error: shared cache cannot be nil")
		}
		return &factory.MemoryResourceClientFactory{
			Cache: sharedCache,
		}, nil
	}

	switch source := settings.ConfigSource.(type) {
	// this is at trick to reuse the same cfg across multiple clients
	case *v1.Settings_KubernetesConfigSource:
		kubeCache := params.kube.kubeCache
		cfg := params.kube.restCfg
		if *cfg == nil {
			c, err := kubeutils.GetConfig("", "")
			if err != nil {
				return nil, err
			}

			if kube := settings.GetKubernetes(); kube != nil {
				if rl := kube.GetRateLimits(); rl != nil {
					c.QPS = rl.QPS
					c.Burst = int(rl.Burst)
				}
			}

			*cfg = c
		}

		return &factory.KubeResourceClientFactory{
			Crd:         resourceCrd,
			Cfg:         *cfg,
			SharedCache: kubeCache,
		}, nil
	case *v1.Settings_ConsulKvSource:
		consulClient := params.consul.consulClient
		rootKey := source.ConsulKvSource.GetRootKey()
		if rootKey == "" {
			rootKey = DefaultRootKey
		}
		return &factory.ConsulResourceClientFactory{
			Consul:  consulClient,
			RootKey: rootKey,
		}, nil
	case *v1.Settings_DirectoryConfigSource:
		return &factory.FileResourceClientFactory{
			RootDir: filepath.Join(source.DirectoryConfigSource.Directory, resourceCrd.Plural),
		}, nil
	}
	return nil, errors.Errorf("invalid config source type")
}

func KubeServiceClientForSettings(ctx context.Context,
	settings *v1.Settings,
	sharedCache memory.InMemoryResourceCache,
	cfg **rest.Config,
	clientset *kubernetes.Interface,
	kubeCoreCache *cache.KubeCoreCache) (skkube.ServiceClient, error) {

	// We are running in kubernetes
	switch settings.ConfigSource.(type) {
	case *v1.Settings_KubernetesConfigSource:
		if err := initializeForKube(ctx, cfg, clientset, kubeCoreCache); err != nil {
			return nil, errors.Wrapf(err, "initializing kube cfg clientset and core cache")
		}
		return service.NewServiceClient(*clientset, *kubeCoreCache), nil
	}

	// In all other cases, run in memory
	if sharedCache == nil {
		return nil, errors.Errorf("internal error: shared cache cannot be nil")
	}
	memoryRcFactory := &factory.MemoryResourceClientFactory{Cache: sharedCache}
	inMemoryClient, err := memoryRcFactory.NewResourceClient(factory.NewResourceClientParams{
		ResourceType: &skkube.Service{},
	})
	if err != nil {
		return nil, err
	}
	return skkube.NewServiceClientWithBase(inMemoryClient), nil
}

// sharedCach OR resourceCrd+cfg must be non-nil
func SecretFactoryForSettings(ctx context.Context,
	settings *v1.Settings,
	sharedCache memory.InMemoryResourceCache,
	cfg **rest.Config,
	clientset *kubernetes.Interface,
	kubeCoreCache *cache.KubeCoreCache,
	vaultClient *vaultapi.Client,
	pluralName string) (factory.ResourceClientFactory, error) {
	if settings.SecretSource == nil {
		if sharedCache == nil {
			return nil, errors.Errorf("internal error: shared cache cannot be nil")
		}
		return &factory.MemoryResourceClientFactory{
			Cache: sharedCache,
		}, nil
	}

	switch source := settings.SecretSource.(type) {
	case *v1.Settings_KubernetesSecretSource:
		if err := initializeForKube(ctx, cfg, clientset, kubeCoreCache); err != nil {
			return nil, errors.Wrapf(err, "initializing kube cfg clientset and core cache")
		}
		converterChain := kubeconverters.NewSecretConverterChain(new(kubeconverters.TLSSecretConverter))
		return &factory.KubeSecretClientFactory{
			Clientset:       *clientset,
			Cache:           *kubeCoreCache,
			SecretConverter: converterChain,
		}, nil
	case *v1.Settings_VaultSecretSource:
		rootKey := source.VaultSecretSource.GetRootKey()
		if rootKey == "" {
			rootKey = DefaultRootKey
		}
		return &factory.VaultSecretClientFactory{
			Vault:   vaultClient,
			RootKey: rootKey,
		}, nil
		return nil, errors.Errorf("vault configuration not implemented")
	case *v1.Settings_DirectorySecretSource:
		return &factory.FileResourceClientFactory{
			RootDir: filepath.Join(source.DirectorySecretSource.Directory, pluralName),
		}, nil
	}
	return nil, errors.Errorf("invalid config source type")
}

// sharedCach OR resourceCrd+cfg must be non-nil
func ArtifactFactoryForSettings(ctx context.Context,
	settings *v1.Settings,
	sharedCache memory.InMemoryResourceCache,
	cfg **rest.Config,
	clientset *kubernetes.Interface,
	kubeCoreCache *cache.KubeCoreCache,
	pluralName string) (factory.ResourceClientFactory, error) {
	if settings.SecretSource == nil {
		if sharedCache == nil {
			return nil, errors.Errorf("internal error: shared cache cannot be nil")
		}
		return &factory.MemoryResourceClientFactory{
			Cache: sharedCache,
		}, nil
	}

	switch source := settings.ArtifactSource.(type) {
	case *v1.Settings_KubernetesArtifactSource:
		if err := initializeForKube(ctx, cfg, clientset, kubeCoreCache); err != nil {
			return nil, errors.Wrapf(err, "initializing kube cfg clientset and core cache")
		}
		return &factory.KubeSecretClientFactory{
			Clientset: *clientset,
			Cache:     *kubeCoreCache,
		}, nil
	case *v1.Settings_DirectoryArtifactSource:
		return &factory.FileResourceClientFactory{
			RootDir: filepath.Join(source.DirectoryArtifactSource.Directory, pluralName),
		}, nil
	}
	return nil, errors.Errorf("invalid config source type")
}

func initializeForKube(ctx context.Context,
	cfg **rest.Config,
	clientset *kubernetes.Interface,
	kubeCoreCache *cache.KubeCoreCache) error {
	if cfg == nil {
		c, err := kubeutils.GetConfig("", "")
		if err != nil {
			return err
		}
		*cfg = c
	}

	if *clientset == nil {
		cs, err := kubernetes.NewForConfig(*cfg)
		if err != nil {
			return err
		}
		*clientset = cs
	}

	if *kubeCoreCache == nil {
		coreCache, err := cache.NewKubeCoreCache(ctx, *clientset)
		if err != nil {
			return err
		}
		*kubeCoreCache = coreCache
	}

	return nil

}
