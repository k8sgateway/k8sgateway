// Code generated by solo-kit. DO NOT EDIT.

package v1alpha1

import (
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/errors"
)

type IngressWatcher interface {
	// watch namespace-scoped ingresses
	Watch(namespace string, opts clients.WatchOpts) (<-chan IngressList, <-chan error, error)
}

type IngressClient interface {
	BaseClient() clients.ResourceClient
	Register() error
	Read(namespace, name string, opts clients.ReadOpts) (*Ingress, error)
	Write(resource *Ingress, opts clients.WriteOpts) (*Ingress, error)
	Delete(namespace, name string, opts clients.DeleteOpts) error
	List(namespace string, opts clients.ListOpts) (IngressList, error)
	IngressWatcher
}

type ingressClient struct {
	rc clients.ResourceClient
}

func NewIngressClient(rcFactory factory.ResourceClientFactory) (IngressClient, error) {
	return NewIngressClientWithToken(rcFactory, "")
}

func NewIngressClientWithToken(rcFactory factory.ResourceClientFactory, token string) (IngressClient, error) {
	rc, err := rcFactory.NewResourceClient(factory.NewResourceClientParams{
		ResourceType: &Ingress{},
		Token:        token,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "creating base Ingress resource client")
	}
	return NewIngressClientWithBase(rc), nil
}

func NewIngressClientWithBase(rc clients.ResourceClient) IngressClient {
	return &ingressClient{
		rc: rc,
	}
}

func (client *ingressClient) BaseClient() clients.ResourceClient {
	return client.rc
}

func (client *ingressClient) Register() error {
	return client.rc.Register()
}

func (client *ingressClient) Read(namespace, name string, opts clients.ReadOpts) (*Ingress, error) {
	opts = opts.WithDefaults()

	resource, err := client.rc.Read(namespace, name, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*Ingress), nil
}

func (client *ingressClient) Write(ingress *Ingress, opts clients.WriteOpts) (*Ingress, error) {
	opts = opts.WithDefaults()
	resource, err := client.rc.Write(ingress, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*Ingress), nil
}

func (client *ingressClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	opts = opts.WithDefaults()

	return client.rc.Delete(namespace, name, opts)
}

func (client *ingressClient) List(namespace string, opts clients.ListOpts) (IngressList, error) {
	opts = opts.WithDefaults()

	resourceList, err := client.rc.List(namespace, opts)
	if err != nil {
		return nil, err
	}
	return convertToIngress(resourceList), nil
}

func (client *ingressClient) Watch(namespace string, opts clients.WatchOpts) (<-chan IngressList, <-chan error, error) {
	opts = opts.WithDefaults()

	resourcesChan, errs, initErr := client.rc.Watch(namespace, opts)
	if initErr != nil {
		return nil, nil, initErr
	}
	ingressesChan := make(chan IngressList)
	go func() {
		for {
			select {
			case resourceList := <-resourcesChan:
				ingressesChan <- convertToIngress(resourceList)
			case <-opts.Ctx.Done():
				close(ingressesChan)
				return
			}
		}
	}()
	return ingressesChan, errs, nil
}

func convertToIngress(resources resources.ResourceList) IngressList {
	var ingressList IngressList
	for _, resource := range resources {
		ingressList = append(ingressList, resource.(*Ingress))
	}
	return ingressList
}
