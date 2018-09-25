package mocks

import (
	"context"

	"go.opencensus.io/trace"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/pkg/utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/utils/errutils"
)

type FestingSyncer interface {
	Sync(context.Context, *FestingSnapshot) error
}

type FestingEventLoop interface {
	Run(namespaces []string, opts clients.WatchOpts) (<-chan error, error)
}

type festingEventLoop struct {
	emitter FestingEmitter
	syncer  FestingSyncer
}

func NewFestingEventLoop(emitter FestingEmitter, syncer FestingSyncer) FestingEventLoop {
	return &festingEventLoop{
		emitter: emitter,
		syncer:  syncer,
	}
}

func (el *festingEventLoop) Run(namespaces []string, opts clients.WatchOpts) (<-chan error, error) {
	opts = opts.WithDefaults()
	opts.Ctx = contextutils.WithLogger(opts.Ctx, "mocks.event_loop")
	logger := contextutils.LoggerFrom(opts.Ctx)
	logger.Infof("event loop started")

	errs := make(chan error)

	watch, emitterErrs, err := el.emitter.Snapshots(namespaces, opts)
	if err != nil {
		return nil, errors.Wrapf(err, "starting snapshot watch")
	}
	go errutils.AggregateErrs(opts.Ctx, errs, emitterErrs, "mocks.emitter errors")
	go func() {

		// create a new context for each loop, cancel it before each loop
		var cancel context.CancelFunc = func() {}
		defer cancel()
		for {
			select {
			case snapshot, ok := <-watch:
				if !ok {
					return
				}
				// cancel any open watches from previous loop
				cancel()

				ctx, span := trace.StartSpan(opts.Ctx, "festing.solo.io.EventLoopSync")
				ctx, canc := context.WithCancel(ctx)
				cancel = canc
				err := el.syncer.Sync(ctx, snapshot)
				span.End()

				if err != nil {
					select {
					case errs <- err:
					default:
						logger.Errorf("write error channel is full! could not propagate err: %v", err)
					}
				}
			case <-opts.Ctx.Done():
				return
			}
		}
	}()
	return errs, nil
}
