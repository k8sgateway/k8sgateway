package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/pkg/errors"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"go.uber.org/zap"
)

// Ec2InstanceLister is a simple interface for calling the AWS API.
// This allows us to easily mock the API in our tests.
type Ec2InstanceLister interface {
	ListForCredentials(ctx context.Context, awsRegion string, secretRef core.ResourceRef, secrets v1.SecretList) ([]*ec2.Instance, error)
}

type ec2InstanceLister struct {
}

func NewEc2InstanceLister() *ec2InstanceLister {
	return &ec2InstanceLister{}
}

var _ Ec2InstanceLister = &ec2InstanceLister{}

func (c *ec2InstanceLister) ListForCredentials(ctx context.Context, awsRegion string, secretRef core.ResourceRef, secrets v1.SecretList) ([]*ec2.Instance, error) {
	logger := contextutils.LoggerFrom(ctx)
	sess, err := getEc2SessionForCredentials(awsRegion, secretRef, secrets)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get aws client")
	}
	svc := ec2.New(sess)
	result, err := svc.DescribeInstances(describeInstancesInputForAllInstances())
	if err != nil {
		return nil, errors.Wrapf(err, "unable to describe instances")
	}
	logger.Debugw("ec2Upstream result", zap.Any("value", result))
	return getInstancesFromDescription(result), nil
}
