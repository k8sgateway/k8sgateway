// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/errors"
)

type ClusterIngressClient interface {
	BaseClient() clients.ResourceClient
	Register() error
	Read(namespace, name string, opts clients.ReadOpts) (*ClusterIngress, error)
	Write(resource *ClusterIngress, opts clients.WriteOpts) (*ClusterIngress, error)
	Delete(namespace, name string, opts clients.DeleteOpts) error
	List(namespace string, opts clients.ListOpts) (ClusterIngressList, error)
	Watch(namespace string, opts clients.WatchOpts) (<-chan ClusterIngressList, <-chan error, error)
}

type clusterIngressClient struct {
	rc clients.ResourceClient
}

func NewClusterIngressClient(rcFactory factory.ResourceClientFactory) (ClusterIngressClient, error) {
	return NewClusterIngressClientWithToken(rcFactory, "")
}

func NewClusterIngressClientWithToken(rcFactory factory.ResourceClientFactory, token string) (ClusterIngressClient, error) {
	rc, err := rcFactory.NewResourceClient(factory.NewResourceClientParams{
		ResourceType: &ClusterIngress{},
		Token:        token,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "creating base ClusterIngress resource client")
	}
	return NewClusterIngressClientWithBase(rc), nil
}

func NewClusterIngressClientWithBase(rc clients.ResourceClient) ClusterIngressClient {
	return &clusterIngressClient{
		rc: rc,
	}
}

func (client *clusterIngressClient) BaseClient() clients.ResourceClient {
	return client.rc
}

func (client *clusterIngressClient) Register() error {
	return client.rc.Register()
}

func (client *clusterIngressClient) Read(namespace, name string, opts clients.ReadOpts) (*ClusterIngress, error) {
	opts = opts.WithDefaults()

	resource, err := client.rc.Read(namespace, name, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*ClusterIngress), nil
}

func (client *clusterIngressClient) Write(clusterIngress *ClusterIngress, opts clients.WriteOpts) (*ClusterIngress, error) {
	opts = opts.WithDefaults()
	resource, err := client.rc.Write(clusterIngress, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*ClusterIngress), nil
}

func (client *clusterIngressClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	opts = opts.WithDefaults()

	return client.rc.Delete(namespace, name, opts)
}

func (client *clusterIngressClient) List(namespace string, opts clients.ListOpts) (ClusterIngressList, error) {
	opts = opts.WithDefaults()

	resourceList, err := client.rc.List(namespace, opts)
	if err != nil {
		return nil, err
	}
	return convertToClusterIngress(resourceList), nil
}

func (client *clusterIngressClient) Watch(namespace string, opts clients.WatchOpts) (<-chan ClusterIngressList, <-chan error, error) {
	opts = opts.WithDefaults()

	resourcesChan, errs, initErr := client.rc.Watch(namespace, opts)
	if initErr != nil {
		return nil, nil, initErr
	}
	clusteringressesChan := make(chan ClusterIngressList)
	go func() {
		for {
			select {
			case resourceList := <-resourcesChan:
				clusteringressesChan <- convertToClusterIngress(resourceList)
			case <-opts.Ctx.Done():
				close(clusteringressesChan)
				return
			}
		}
	}()
	return clusteringressesChan, errs, nil
}

func convertToClusterIngress(resources resources.ResourceList) ClusterIngressList {
	var clusterIngressList ClusterIngressList
	for _, resource := range resources {
		clusterIngressList = append(clusterIngressList, resource.(*ClusterIngress))
	}
	return clusterIngressList
}
