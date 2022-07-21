package ratelimit

import (
	"context"
	"fmt"

	"github.com/solo-io/gloo/projects/gloo/pkg/utils"

	"github.com/rotisserie/eris"

	"github.com/solo-io/solo-kit/pkg/api/v2/reporter"

	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	gloov1snap "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/gloosnapshot"
	"github.com/solo-io/gloo/projects/gloo/pkg/syncer"
	"github.com/solo-io/go-utils/contextutils"
	envoycache "github.com/solo-io/solo-kit/pkg/api/v1/control-plane/cache"
)

// Compile-time assertion
var (
	_ syncer.TranslatorSyncerExtension = new(TranslatorSyncerExtension)
)

const (
	Name       = "rate-limit"
	ServerRole = "ratelimit"
)

type TranslatorSyncerExtension struct {
}

func NewTranslatorSyncerExtension(_ context.Context, params syncer.TranslatorSyncerExtensionParams) (syncer.TranslatorSyncerExtension, error) {
	return &TranslatorSyncerExtension{}, nil
}

func (s *TranslatorSyncerExtension) ID() string {
	return ServerRole
}

func (s *TranslatorSyncerExtension) Sync(
	ctx context.Context,
	snap *gloov1snap.ApiSnapshot,
	_ *gloov1.Settings,
	_ envoycache.SnapshotCache,
	_ reporter.ResourceReports,
) error {
	ctx = contextutils.WithLogger(ctx, "rateLimitTranslatorSyncer")
	logger := contextutils.LoggerFrom(ctx)

	enterpriseOnlyError := func(enterpriseFeature string) error {
		errorMsg := createErrorMsg(enterpriseFeature)
		logger.Errorf(errorMsg)
		return eris.New(errorMsg)
	}

	for _, proxy := range snap.Proxies {
		for _, listener := range proxy.GetListeners() {
			virtualHosts := utils.GetVhostsFromListener(listener)

			for _, virtualHost := range virtualHosts {

				// RateLimitConfigs is an enterprise feature https://docs.solo.io/gloo-edge/latest/guides/security/rate_limiting/crds/
				if virtualHost.GetOptions().GetRateLimitConfigs() != nil {
					return enterpriseOnlyError("RateLimitConfig")
				}

				// ratelimitBasic is an enterprise feature https://docs.solo.io/gloo-edge/latest/guides/security/rate_limiting/simple/
				if virtualHost.GetOptions().GetRatelimitBasic() != nil {
					return enterpriseOnlyError("ratelimitBasic")
				}

				// check setActions on vhost
				rlactionsVhost := virtualHost.GetOptions().GetRatelimit().GetRateLimits()
				for _, rlaction := range rlactionsVhost {
					if rlaction.GetSetActions() != nil {
						return enterpriseOnlyError("setActions")
					}
				}

				// Staged RateLimiting is an enterprise feature
				if virtualHost.GetOptions().GetRateLimitEarlyConfigType() != nil {
					return enterpriseOnlyError("RateLimitEarly")
				}
				if virtualHost.GetOptions().GetRateLimitRegularConfigType() != nil {
					return enterpriseOnlyError("RateLimitRegular")
				}

				for _, route := range virtualHost.GetRoutes() {
					if route.GetOptions().GetRateLimitConfigs() != nil {
						return enterpriseOnlyError("RateLimitConfig")
					}

					if route.GetOptions().GetRatelimitBasic() != nil {
						return enterpriseOnlyError("ratelimitBasic")
					}

					// check setActions on route
					rlactionsRoute := route.GetOptions().GetRatelimit().GetRateLimits()
					for _, rlaction := range rlactionsRoute {
						if rlaction.GetSetActions() != nil {
							return enterpriseOnlyError("setActions")
						}
					}

					// Staged RateLimiting is an enterprise feature
					if route.GetOptions().GetRateLimitEarlyConfigType() != nil {
						return enterpriseOnlyError("RateLimitEarly")
					}
					if route.GetOptions().GetRateLimitRegularConfigType() != nil {
						return enterpriseOnlyError("RateLimitRegular")
					}
				}
			}
		}
	}

	return nil
}

func createErrorMsg(feature string) string {
	return fmt.Sprintf("The Gloo Advanced Rate limit API feature '%s' is enterprise-only, please upgrade or use the Envoy rate-limit API instead", feature)
}
