// Code generated by solo-kit. DO NOT EDIT.

package extauth

import (
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/errors"
)

type AuthConfigWatcher interface {
	// watch namespace-scoped AuthConfigs
	Watch(namespace string, opts clients.WatchOpts) (<-chan AuthConfigList, <-chan error, error)
}

type AuthConfigClient interface {
	BaseClient() clients.ResourceClient
	Register() error
	Read(namespace, name string, opts clients.ReadOpts) (*AuthConfig, error)
	Write(resource *AuthConfig, opts clients.WriteOpts) (*AuthConfig, error)
	Delete(namespace, name string, opts clients.DeleteOpts) error
	List(namespace string, opts clients.ListOpts) (AuthConfigList, error)
	AuthConfigWatcher
}

type authConfigClient struct {
	rc clients.ResourceClient
}

func NewAuthConfigClient(rcFactory factory.ResourceClientFactory) (AuthConfigClient, error) {
	return NewAuthConfigClientWithToken(rcFactory, "")
}

func NewAuthConfigClientWithToken(rcFactory factory.ResourceClientFactory, token string) (AuthConfigClient, error) {
	rc, err := rcFactory.NewResourceClient(factory.NewResourceClientParams{
		ResourceType: &AuthConfig{},
		Token:        token,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "creating base AuthConfig resource client")
	}
	return NewAuthConfigClientWithBase(rc), nil
}

func NewAuthConfigClientWithBase(rc clients.ResourceClient) AuthConfigClient {
	return &authConfigClient{
		rc: rc,
	}
}

func (client *authConfigClient) BaseClient() clients.ResourceClient {
	return client.rc
}

func (client *authConfigClient) Register() error {
	return client.rc.Register()
}

func (client *authConfigClient) Read(namespace, name string, opts clients.ReadOpts) (*AuthConfig, error) {
	opts = opts.WithDefaults()

	resource, err := client.rc.Read(namespace, name, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*AuthConfig), nil
}

func (client *authConfigClient) Write(authConfig *AuthConfig, opts clients.WriteOpts) (*AuthConfig, error) {
	opts = opts.WithDefaults()
	resource, err := client.rc.Write(authConfig, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*AuthConfig), nil
}

func (client *authConfigClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	opts = opts.WithDefaults()

	return client.rc.Delete(namespace, name, opts)
}

func (client *authConfigClient) List(namespace string, opts clients.ListOpts) (AuthConfigList, error) {
	opts = opts.WithDefaults()

	resourceList, err := client.rc.List(namespace, opts)
	if err != nil {
		return nil, err
	}
	return convertToAuthConfig(resourceList), nil
}

func (client *authConfigClient) Watch(namespace string, opts clients.WatchOpts) (<-chan AuthConfigList, <-chan error, error) {
	opts = opts.WithDefaults()

	resourcesChan, errs, initErr := client.rc.Watch(namespace, opts)
	if initErr != nil {
		return nil, nil, initErr
	}
	authConfigsChan := make(chan AuthConfigList)
	go func() {
		for {
			select {
			case resourceList := <-resourcesChan:
				authConfigsChan <- convertToAuthConfig(resourceList)
			case <-opts.Ctx.Done():
				close(authConfigsChan)
				return
			}
		}
	}()
	return authConfigsChan, errs, nil
}

func convertToAuthConfig(resources resources.ResourceList) AuthConfigList {
	var authConfigList AuthConfigList
	for _, resource := range resources {
		authConfigList = append(authConfigList, resource.(*AuthConfig))
	}
	return authConfigList
}
