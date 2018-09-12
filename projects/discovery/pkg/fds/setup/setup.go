package setup

import (
	"time"

	"github.com/solo-io/solo-kit/pkg/utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/utils/errutils"
	"github.com/solo-io/solo-kit/projects/discovery/pkg/fds"
	"github.com/solo-io/solo-kit/projects/discovery/pkg/fds/discoveries/aws"
	"github.com/solo-io/solo-kit/projects/discovery/pkg/fds/discoveries/swagger"
	"github.com/solo-io/solo-kit/projects/discovery/pkg/fds/syncer"
	"github.com/solo-io/solo-kit/projects/gloo/pkg/api/v1"
	"github.com/solo-io/solo-kit/projects/gloo/pkg/bootstrap"
	"github.com/solo-io/solo-kit/projects/gloo/pkg/plugins/registry"
)

func Setup(opts bootstrap.Opts) error {
	// TODO: Ilackarms: move this to multi-eventloop
	namespaces, errs, err := opts.Namespacer.Namespaces(opts.WatchOpts)
	if err != nil {
		return err
	}
	for {
		select {
		case err := <-errs:
			return err
		case watchNamespaces := <-namespaces:
			err := setupForNamespaces(watchNamespaces, opts)
			if err != nil {
				return err
			}
		}
	}
}

func setupForNamespaces(discoveredNamespaces []string, opts bootstrap.Opts) error {
	watchOpts := opts.WatchOpts.WithDefaults()

	watchOpts.Ctx = contextutils.WithLogger(watchOpts.Ctx, "setup")

	upstreamClient, err := v1.NewUpstreamClient(opts.Upstreams)
	if err != nil {
		return err
	}
	if err := upstreamClient.Register(); err != nil {
		return err
	}
	secretClient, err := v1.NewSecretClient(opts.Secrets)
	if err != nil {
		return err
	}
	if err := secretClient.Register(); err != nil {
		return err
	}

	cache := v1.NewDiscoveryEmitter(secretClient, upstreamClient)

	var resolvers fds.Resolvers
	for _, plug := range registry.Plugins(opts) {
		resolver, ok := plug.(fds.Resolver)
		if ok {
			resolvers = append(resolvers, resolver)
		}
	}

	// TODO: unhardcode
	functionalPlugins := []fds.FunctionDiscoveryFactory{
		&aws.AWSLambdaFunctionDiscoveryFactory{
			PollingTime: time.Second,
		},
		&swagger.SwaggerFunctionDiscoveryFactory{
			DetectionTimeout: time.Minute,
			FunctionPollTime: time.Second * 15,
		},
	}

	// TODO(yuval-k): max Concurrency here
	updater := fds.NewUpdater(watchOpts.Ctx, resolvers, upstreamClient, 0, functionalPlugins)
	disc := fds.NewFunctionDiscovery(updater)

	sync := syncer.NewSyncer(disc)
	eventLoop := v1.NewDiscoveryEventLoop(cache, sync)

	errs := make(chan error)

	eventLoopErrs, err := eventLoop.Run(discoveredNamespaces, watchOpts)
	if err != nil {
		return err
	}
	go errutils.AggregateErrs(watchOpts.Ctx, errs, eventLoopErrs, "event_loop.gloo")

	logger := contextutils.LoggerFrom(watchOpts.Ctx)

	for {
		select {
		case err := <-errs:
			logger.Errorf("error: %v", err)
		case <-watchOpts.Ctx.Done():
			return nil
		}
	}
}
