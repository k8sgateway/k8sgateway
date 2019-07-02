// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/errors"
)

type UpstreamGroupWatcher interface {
	// watch namespace-scoped UpstreamGroups
	Watch(namespace string, opts clients.WatchOpts) (<-chan UpstreamGroupList, <-chan error, error)
}

type UpstreamGroupClient interface {
	BaseClient() clients.ResourceClient
	Register() error
	Read(namespace, name string, opts clients.ReadOpts) (*UpstreamGroup, error)
	Write(resource *UpstreamGroup, opts clients.WriteOpts) (*UpstreamGroup, error)
	Delete(namespace, name string, opts clients.DeleteOpts) error
	List(namespace string, opts clients.ListOpts) (UpstreamGroupList, error)
	UpstreamGroupWatcher
}

type upstreamGroupClient struct {
	rc clients.ResourceClient
}

func NewUpstreamGroupClient(rcFactory factory.ResourceClientFactory) (UpstreamGroupClient, error) {
	return NewUpstreamGroupClientWithToken(rcFactory, "")
}

func NewUpstreamGroupClientWithToken(rcFactory factory.ResourceClientFactory, token string) (UpstreamGroupClient, error) {
	rc, err := rcFactory.NewResourceClient(factory.NewResourceClientParams{
		ResourceType: &UpstreamGroup{},
		Token:        token,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "creating base UpstreamGroup resource client")
	}
	return NewUpstreamGroupClientWithBase(rc), nil
}

func NewUpstreamGroupClientWithBase(rc clients.ResourceClient) UpstreamGroupClient {
	return &upstreamGroupClient{
		rc: rc,
	}
}

func (client *upstreamGroupClient) BaseClient() clients.ResourceClient {
	return client.rc
}

func (client *upstreamGroupClient) Register() error {
	return client.rc.Register()
}

func (client *upstreamGroupClient) Read(namespace, name string, opts clients.ReadOpts) (*UpstreamGroup, error) {
	opts = opts.WithDefaults()

	resource, err := client.rc.Read(namespace, name, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*UpstreamGroup), nil
}

func (client *upstreamGroupClient) Write(upstreamGroup *UpstreamGroup, opts clients.WriteOpts) (*UpstreamGroup, error) {
	opts = opts.WithDefaults()
	resource, err := client.rc.Write(upstreamGroup, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*UpstreamGroup), nil
}

func (client *upstreamGroupClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	opts = opts.WithDefaults()

	return client.rc.Delete(namespace, name, opts)
}

func (client *upstreamGroupClient) List(namespace string, opts clients.ListOpts) (UpstreamGroupList, error) {
	opts = opts.WithDefaults()

	resourceList, err := client.rc.List(namespace, opts)
	if err != nil {
		return nil, err
	}
	return convertToUpstreamGroup(resourceList), nil
}

func (client *upstreamGroupClient) Watch(namespace string, opts clients.WatchOpts) (<-chan UpstreamGroupList, <-chan error, error) {
	opts = opts.WithDefaults()

	resourcesChan, errs, initErr := client.rc.Watch(namespace, opts)
	if initErr != nil {
		return nil, nil, initErr
	}
	UpstreamGroupsChan := make(chan UpstreamGroupList)
	go func() {
		for {
			select {
			case resourceList := <-resourcesChan:
				UpstreamGroupsChan <- convertToUpstreamGroup(resourceList)
			case <-opts.Ctx.Done():
				close(UpstreamGroupsChan)
				return
			}
		}
	}()
	return UpstreamGroupsChan, errs, nil
}

func convertToUpstreamGroup(resources resources.ResourceList) UpstreamGroupList {
	var upstreamGroupList UpstreamGroupList
	for _, resource := range resources {
		upstreamGroupList = append(upstreamGroupList, resource.(*UpstreamGroup))
	}
	return upstreamGroupList
}
