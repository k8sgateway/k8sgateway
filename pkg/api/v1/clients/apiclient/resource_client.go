package apiclient

import (
	"io"
	"reflect"
	"sort"

	"strings"

	"github.com/gogo/protobuf/types"
	"github.com/solo-io/solo-kit/pkg/api/v1/apiserver"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/pkg/utils/protoutils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type ResourceClient struct {
	grpc         apiserver.ApiServerClient
	resourceType resources.Resource
}

func NewResourceClient(cc *grpc.ClientConn, resourceType resources.Resource) *ResourceClient {
	return &ResourceClient{
		grpc:         apiserver.NewApiServerClient(cc),
		resourceType: resourceType,
	}
}

var _ clients.ResourceClient = &ResourceClient{}

func (rc *ResourceClient) Kind() string {
	return reflect.TypeOf(rc.resourceType).String()
}

func (rc *ResourceClient) NewResource() resources.Resource {
	return resources.Clone(rc.resourceType)
}

func (rc *ResourceClient) Register() error {
	return nil
}

func (rc *ResourceClient) Read(name string, opts clients.ReadOpts) (resources.Resource, error) {
	if err := resources.ValidateName(name); err != nil {
		return nil, errors.Wrapf(err, "validation error")
	}
	opts = opts.WithDefaults()
	resp, err := rc.grpc.Read(opts.Ctx, &apiserver.ReadRequest{
		Name:      name,
		Namespace: opts.Namespace,
		Kind:      rc.Kind(),
	})
	if err != nil {
		if stat, ok := status.FromError(err); ok && strings.Contains(stat.Message(), "does not exist") {
			return nil, errors.NewNotExistErr(opts.Namespace, name)
		}
		return nil, errors.Wrapf(err, "performing grpc request")
	}
	resource := rc.NewResource()
	if err := protoutils.UnmarshalStruct(resp.Resource.Data, resource); err != nil {
		return nil, errors.Wrapf(err, "reading proto struct into %v", reflect.TypeOf(rc.resourceType))
	}
	return resource, nil
}

func (rc *ResourceClient) Write(resource resources.Resource, opts clients.WriteOpts) (resources.Resource, error) {
	opts = opts.WithDefaults()
	if err := resources.Validate(resource); err != nil {
		return nil, errors.Wrapf(err, "validation error")
	}
	data, err := protoutils.MarshalStruct(resource)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to marshal resource")
	}
	resp, err := rc.grpc.Write(opts.Ctx, &apiserver.WriteRequest{
		Resource: &apiserver.Resource{
			Data: data,
			Kind: rc.Kind(),
		},
		OverwriteExisting: opts.OverwriteExisting,
	})
	if err != nil {
		if stat, ok := status.FromError(err); ok && strings.Contains(stat.Message(), "exists") {
			return nil, errors.NewExistErr(resource.GetMetadata())
		}
		return nil, errors.Wrapf(err, "performing grpc request")
	}
	written := rc.NewResource()
	if err := protoutils.UnmarshalStruct(resp.Resource.Data, written); err != nil {
		return nil, errors.Wrapf(err, "reading proto struct into %v", reflect.TypeOf(rc.resourceType))
	}
	return written, nil
}

func (rc *ResourceClient) Delete(name string, opts clients.DeleteOpts) error {
	opts = opts.WithDefaults()
	_, err := rc.grpc.Delete(opts.Ctx, &apiserver.DeleteRequest{
		Name:           name,
		Namespace:      opts.Namespace,
		Kind:           rc.Kind(),
		IgnoreNotExist: opts.IgnoreNotExist,
	})
	if err != nil {
		if stat, ok := status.FromError(err); ok && strings.Contains(stat.Message(), "does not exist") {
			return errors.NewNotExistErr(opts.Namespace, name)
		}
		return errors.Wrapf(err, "deleting resource %v", name)
	}
	return nil
}

func (rc *ResourceClient) List(opts clients.ListOpts) ([]resources.Resource, error) {
	opts = opts.WithDefaults()
	resp, err := rc.grpc.List(opts.Ctx, &apiserver.ListRequest{
		Namespace: opts.Namespace,
		Kind:      rc.Kind(),
	})
	if err != nil {
		return nil, errors.Wrapf(err, "performing grpc request")
	}

	var resourceList []resources.Resource
	for _, resourceData := range resp.ResourceList {
		resource := rc.NewResource()
		if err := protoutils.UnmarshalStruct(resourceData.Data, resource); err != nil {
			return nil, errors.Wrapf(err, "reading proto struct into %v", reflect.TypeOf(rc.resourceType))
		}
		resourceList = append(resourceList, resource)
	}

	sort.SliceStable(resourceList, func(i, j int) bool {
		return resourceList[i].GetMetadata().Name < resourceList[j].GetMetadata().Name
	})

	return resourceList, nil
}

func (rc *ResourceClient) Watch(opts clients.WatchOpts) (<-chan []resources.Resource, <-chan error, error) {
	opts = opts.WithDefaults()
	resp, err := rc.grpc.Watch(opts.Ctx, &apiserver.WatchRequest{
		SyncFrequency: &types.Duration{
			Nanos: int32(opts.RefreshRate),
		},
		Namespace: opts.Namespace,
		Kind:      rc.Kind(),
	})
	if err != nil {
		return nil, nil, errors.Wrapf(err, "performing grpc request")
	}

	resourcesChan := make(chan []resources.Resource)
	errs := make(chan error)
	// watch should open up with an initial read
	go func() {
		list, err := rc.List(clients.ListOpts{
			Ctx:       opts.Ctx,
			Selector:  opts.Selector,
			Namespace: opts.Namespace,
		})
		if err != nil {
			errs <- err
			return
		}
		resourcesChan <- list
	}()
	go func() {
		for {
			resourceDataList, err := resp.Recv()
			if err == io.EOF {
				errs <- errors.Wrapf(err, "grpc stream closed")
				return
			}
			if err != nil {
				errs <- err
				continue
			}
			var resourceList []resources.Resource
			for _, resourceData := range resourceDataList.ResourceList {
				resource := rc.NewResource()
				if err := protoutils.UnmarshalStruct(resourceData.Data, resource); err != nil {
					errs <- errors.Wrapf(err, "reading proto struct into %v", reflect.TypeOf(rc.resourceType))
					continue
				}
				resourceList = append(resourceList, resource)
			}

			sort.SliceStable(resourceList, func(i, j int) bool {
				return resourceList[i].GetMetadata().Name < resourceList[j].GetMetadata().Name
			})
			resourcesChan <- resourceList
		}
	}()

	return resourcesChan, errs, nil
}
