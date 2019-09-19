package runner

import (
	"context"
	"fmt"
	"net"

	pb "github.com/envoyproxy/go-control-plane/envoy/service/metrics/v2"
	"github.com/solo-io/gloo/projects/metrics/pkg/metricsservice"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/go-utils/healthchecker"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func init() {
	view.Register(ocgrpc.DefaultServerViews...)
}

func RunE(parentCtx context.Context) error {
	clientSettings := NewSettings()
	ctx := contextutils.WithLogger(parentCtx, "metrics")

	opts := metricsservice.Options{
		Ctx: ctx,
	}
	service := metricsservice.NewServer(opts)

	return RunWithSettings(ctx, service, clientSettings)
}

func Run(ctx context.Context) {
	err := RunE(ctx)
	if err != nil {
		if ctx.Err() == nil {
			// not a context error - panic
			panic(err)
		}
	}
}

func RunWithSettings(ctx context.Context, service *metricsservice.Server, clientSettings Settings) error {
	err := StartMetricsService(ctx, clientSettings, service)
	if ctx.Err() != nil {
		return ctx.Err()
	}
	return err
}

func StartMetricsService(ctx context.Context, clientSettings Settings, service *metricsservice.Server) error {
	srv := grpc.NewServer(grpc.StatsHandler(&ocgrpc.ServerHandler{}))

	pb.RegisterMetricsServiceServer(srv, service)
	hc := healthchecker.NewGrpc(clientSettings.ServiceName, health.NewServer())
	healthpb.RegisterHealthServer(srv, hc.GetServer())
	reflection.Register(srv)

	logger := contextutils.LoggerFrom(ctx)
	logger.Infow("Starting metrics server")

	addr := fmt.Sprintf(":%d", clientSettings.ServerPort)
	runMode := "gRPC"
	network := "tcp"

	logger.Infof("metrics server running in [%s] mode, listening at [%s]", runMode, addr)
	lis, err := net.Listen(network, addr)
	if err != nil {
		logger.Errorw("Failed to announce on network", zap.Any("mode", runMode), zap.Any("address", addr), zap.Any("error", err))
		return err
	}
	go func() {
		<-ctx.Done()
		srv.Stop()
		_ = lis.Close()
	}()

	return srv.Serve(lis)
}
