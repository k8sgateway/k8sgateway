// Code generated by solo-kit. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"fmt"
	"time"

	"go.opencensus.io/stats"
	"go.uber.org/zap"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/go-utils/errutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
)

type EnterpriseSimpleEmitter interface {
	Snapshots(ctx context.Context) (<-chan *EnterpriseSnapshot, <-chan error, error)
}

func NewEnterpriseSimpleEmitter(aggregatedWatch clients.ResourceWatch) EnterpriseSimpleEmitter {
	return NewEnterpriseSimpleEmitterWithEmit(aggregatedWatch, make(chan struct{}))
}

func NewEnterpriseSimpleEmitterWithEmit(aggregatedWatch clients.ResourceWatch, emit <-chan struct{}) EnterpriseSimpleEmitter {
	return &enterpriseSimpleEmitter{
		aggregatedWatch: aggregatedWatch,
		forceEmit:       emit,
	}
}

type enterpriseSimpleEmitter struct {
	forceEmit       <-chan struct{}
	aggregatedWatch clients.ResourceWatch
}

func (c *enterpriseSimpleEmitter) Snapshots(ctx context.Context) (<-chan *EnterpriseSnapshot, <-chan error, error) {
	snapshots := make(chan *EnterpriseSnapshot)
	errs := make(chan error)

	untyped, watchErrs, err := c.aggregatedWatch(ctx)
	if err != nil {
		return nil, nil, err
	}

	go errutils.AggregateErrs(ctx, errs, watchErrs, "enterprise-emitter")

	go func() {
		currentSnapshot := EnterpriseSnapshot{}
		timer := time.NewTicker(time.Second * 1)
		var previousHash uint64
		sync := func() {
			currentHash, err := currentSnapshot.Hash(nil)
			if err != nil {
				contextutils.LoggerFrom(ctx).Panicw("error while hashing, this should never happen", zap.Error(err))
			}
			if previousHash == currentHash {
				return
			}

			previousHash = currentHash

			stats.Record(ctx, mEnterpriseSnapshotOut.M(1))
			sentSnapshot := currentSnapshot.Clone()
			snapshots <- &sentSnapshot
		}

		defer func() {
			close(snapshots)
			close(errs)
		}()

		for {
			record := func() { stats.Record(ctx, mEnterpriseSnapshotIn.M(1)) }

			select {
			case <-timer.C:
				sync()
			case <-ctx.Done():
				return
			case <-c.forceEmit:
				sentSnapshot := currentSnapshot.Clone()
				snapshots <- &sentSnapshot
			case untypedList := <-untyped:
				record()

				currentSnapshot = EnterpriseSnapshot{}
				for _, res := range untypedList {
					switch typed := res.(type) {
					case *GraphQLExtendedSchema:
						currentSnapshot.GraphqlSchemas = append(currentSnapshot.GraphqlSchemas, typed)
					default:
						select {
						case errs <- fmt.Errorf("EnterpriseSnapshotEmitter "+
							"cannot process resource %v of type %T", res.GetMetadata().Ref(), res):
						case <-ctx.Done():
							return
						}
					}
				}

			}
		}
	}()
	return snapshots, errs, nil
}
