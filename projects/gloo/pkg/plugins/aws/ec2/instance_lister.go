package ec2

import (
	"context"

	"github.com/solo-io/gloo/projects/gloo/pkg/plugins/aws/ec2/awslister"

	"github.com/solo-io/go-utils/errors"

	"github.com/aws/aws-sdk-go/service/ec2"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"go.uber.org/zap"
)

type ec2InstanceLister struct {
}

func NewEc2InstanceLister() *ec2InstanceLister {
	return &ec2InstanceLister{}
}

var _ awslister.Ec2InstanceLister = &ec2InstanceLister{}

func (c *ec2InstanceLister) ListForCredentials(ctx context.Context, awsRegion string, secretRef core.ResourceRef, secrets v1.SecretList) ([]*ec2.Instance, error) {
	sess, err := getEc2SessionForCredentials(awsRegion, secretRef, secrets)
	if err != nil {
		return nil, GetClientError(err)
	}
	svc := ec2.New(sess)
	// pass an empty selector to get all instances that the session has access to
	result, err := svc.DescribeInstances(&ec2.DescribeInstancesInput{})
	if err != nil {
		return nil, DescribeInstancesError(err)
	}
	contextutils.LoggerFrom(ctx).Debugw("ec2Upstream result", zap.Any("value", result))
	return getInstancesFromDescription(result), nil
}

var (
	GetClientError = func(err error) error {
		return errors.Wrapf(err, "unable to get aws client")
	}

	DescribeInstancesError = func(err error) error {
		return errors.Wrapf(err, "unable to describe instances")
	}
)
