package main

import (
	"context"
	"os"

	"github.com/solo-io/gloo/projects/gateway/pkg/conversion"
	"github.com/solo-io/gloo/projects/gateway/pkg/conversion/setup"
	"github.com/solo-io/go-utils/contextutils"
	"go.uber.org/zap"
)

func main() {
	ctx := contextutils.WithLogger(context.Background(), "gateway-conversion")
	clientSet := setup.MustClientSet(ctx)
	resourceConverter := conversion.NewResourceConverter(
		mustPodNamespace(ctx),
		clientSet.V1Gateway,
		clientSet.V2Gateway,
		conversion.NewGatewayConverter(),
	)

	if err := resourceConverter.ConvertAll(ctx); err != nil {
		contextutils.LoggerFrom(ctx).Fatalw("Failed to upgrade all existing gateway resources.", zap.Error(err))
	}
}

func mustPodNamespace(ctx context.Context) string {
	namespace := os.Getenv("POD_NAMESPACE")
	if namespace == "" {
		contextutils.LoggerFrom(ctx).Fatal("POD_NAMESPACE is not set.")
	}
	return namespace
}
