// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"context"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/errors"
)

type RouteTableWatcher interface {
	// watch namespace-scoped RouteTables
	Watch(namespace string, opts clients.WatchOpts) (<-chan RouteTableList, <-chan error, error)
}

type RouteTableClient interface {
	BaseClient() clients.ResourceClient
	Register() error
	Read(namespace, name string, opts clients.ReadOpts) (*RouteTable, error)
	Write(resource *RouteTable, opts clients.WriteOpts) (*RouteTable, error)
	Delete(namespace, name string, opts clients.DeleteOpts) error
	List(namespace string, opts clients.ListOpts) (RouteTableList, error)
	RouteTableWatcher
}

type routeTableClient struct {
	rc clients.ResourceClient
}

func NewRouteTableClient(ctx context.Context, rcFactory factory.ResourceClientFactory) (RouteTableClient, error) {
	return NewRouteTableClientWithToken(ctx, rcFactory, "")
}

func NewRouteTableClientWithToken(ctx context.Context, rcFactory factory.ResourceClientFactory, token string) (RouteTableClient, error) {
	rc, err := rcFactory.NewResourceClient(ctx, factory.NewResourceClientParams{
		ResourceType: &RouteTable{},
		Token:        token,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "creating base RouteTable resource client")
	}
	return NewRouteTableClientWithBase(rc), nil
}

func NewRouteTableClientWithBase(rc clients.ResourceClient) RouteTableClient {
	return &routeTableClient{
		rc: rc,
	}
}

func (client *routeTableClient) BaseClient() clients.ResourceClient {
	return client.rc
}

func (client *routeTableClient) Register() error {
	return client.rc.Register()
}

func (client *routeTableClient) Read(namespace, name string, opts clients.ReadOpts) (*RouteTable, error) {
	opts = opts.WithDefaults()

	resource, err := client.rc.Read(namespace, name, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*RouteTable), nil
}

func (client *routeTableClient) Write(routeTable *RouteTable, opts clients.WriteOpts) (*RouteTable, error) {
	opts = opts.WithDefaults()
	resource, err := client.rc.Write(routeTable, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*RouteTable), nil
}

func (client *routeTableClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	opts = opts.WithDefaults()

	return client.rc.Delete(namespace, name, opts)
}

func (client *routeTableClient) List(namespace string, opts clients.ListOpts) (RouteTableList, error) {
	opts = opts.WithDefaults()

	resourceList, err := client.rc.List(namespace, opts)
	if err != nil {
		return nil, err
	}
	return convertToRouteTable(resourceList), nil
}

func (client *routeTableClient) Watch(namespace string, opts clients.WatchOpts) (<-chan RouteTableList, <-chan error, error) {
	opts = opts.WithDefaults()

	resourcesChan, errs, initErr := client.rc.Watch(namespace, opts)
	if initErr != nil {
		return nil, nil, initErr
	}
	routeTablesChan := make(chan RouteTableList)
	go func() {
		for {
			select {
			case resourceList := <-resourcesChan:
				select {
				case routeTablesChan <- convertToRouteTable(resourceList):
				case <-opts.Ctx.Done():
					close(routeTablesChan)
					return
				}
			case <-opts.Ctx.Done():
				close(routeTablesChan)
				return
			}
		}
	}()
	return routeTablesChan, errs, nil
}

func convertToRouteTable(resources resources.ResourceList) RouteTableList {
	var routeTableList RouteTableList
	for _, resource := range resources {
		routeTableList = append(routeTableList, resource.(*RouteTable))
	}
	return routeTableList
}
