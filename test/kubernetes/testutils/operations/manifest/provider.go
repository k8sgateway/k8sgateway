package manifest

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/solo-io/gloo/test/kubernetes/testutils/assertions"

	"github.com/solo-io/gloo/pkg/utils/kubeutils/kubectl"
	"github.com/solo-io/gloo/test/kubernetes/testutils/operations"
)

type OperationProvider struct {
	kubeCli *kubectl.Cli
}

func NewProvider() *OperationProvider {
	return &OperationProvider{
		kubeCli: nil,
	}
}

// WithClusterCli sets the OperationProvider to use a Cli
func (p *OperationProvider) WithClusterCli(kubeCli *kubectl.Cli) *OperationProvider {
	p.kubeCli = kubeCli
	return p
}

func (p *OperationProvider) NewApplyManifestOperation(manifest string, assertions ...assertions.DiscreteAssertion) operations.Operation {
	return &operations.BasicOperation{
		OpName: fmt.Sprintf("apply-manifest-%s", filepath.Base(manifest)),
		OpExecute: func(ctx context.Context) error {
			return p.kubeCli.ApplyFile(ctx, manifest)
		},
		OpAssertions: assertions,
	}
}

func (p *OperationProvider) NewDeleteManifestOperation(manifest string, assertions ...assertions.DiscreteAssertion) operations.Operation {
	return &operations.BasicOperation{
		OpName: fmt.Sprintf("delete-manifest-%s", filepath.Base(manifest)),
		OpExecute: func(ctx context.Context) error {
			return p.kubeCli.DeleteFile(ctx, manifest)
		},
		OpAssertions: assertions,
	}
}

func (p *OperationProvider) NewReversibleOperation() operations.ReversibleOperation {
	// Not implemented yet
	// It might be nice to make construction of this easier for developers

	return operations.ReversibleOperation{}
}
