package runner

import (
	"context"
	"net/http"
	"os"
	"time"

	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"

	"github.com/solo-io/ext-auth-service/pkg/config/oauth/token_validation/opaque"
	"github.com/solo-io/ext-auth-service/pkg/config/oauth/user_info"
	plugins "github.com/solo-io/ext-auth-service/pkg/config/plugin"
	"github.com/solo-io/ext-auth-service/pkg/server"
	"github.com/solo-io/ext-auth-service/pkg/service"
	"github.com/solo-io/gloo/pkg/utils/syncutil"
	xdsproto "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/extauth/v1"
	"github.com/solo-io/solo-projects/projects/extauth/pkg/config"

	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	_struct "github.com/golang/protobuf/ptypes/struct"
	"github.com/solo-io/go-utils/contextutils"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	extauthConnectedStateDescription = "zero indicates extauth is unable to connect to connect to gloo to get its configuration"
	mExtauthConnectedState           = stats.Int64("glooe.extauth/xds_client_connected_state", extauthConnectedStateDescription, "1")

	extauthConnectedStateView = &view.View{
		Name:        "glooe.extauth/connected_state",
		Measure:     mExtauthConnectedState,
		Description: extauthConnectedStateDescription,
		Aggregation: view.LastValue(),
		TagKeys:     []tag.Key{},
	}

	extauthConnectedStateCounterDescription = "number of client connections extauth has made to gloo to get its configuration"
	mExtauthConnectedStateCounter           = stats.Int64("glooe.extauth/xds_client_connect_counter", extauthConnectedStateCounterDescription, "1")

	extauthConnectedStateCounterView = &view.View{
		Name:        "glooe.extauth/connected_state_counter",
		Measure:     mExtauthConnectedStateCounter,
		Description: extauthConnectedStateCounterDescription,
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}
)

func init() {
	_ = view.Register(extauthConnectedStateView, extauthConnectedStateCounterView)
}

const (
	// The extauth server sends xDS discovery requests to Gloo to get its configuration from Gloo. This constant determines
	// the value of the nodeInfo.Metadata.role field that the server sends along to retrieve its configuration snapshot,
	// similarly to how the regular Gloo gateway-proxies do.
	ExtAuthServerRole = "extauth"
	ModuleName        = "xDS"
)

// This `RunnableModule` implementation uses xDS to get rate limit configuration from the GlooE control plane.
type configSource struct {
	settings Settings
}

func NewConfigSource(settings Settings) server.RunnableModule {
	return &configSource{
		settings: settings,
	}
}

func (*configSource) Name() string {
	return ModuleName
}

func (x *configSource) Run(ctx context.Context, service service.ExtAuthService) error {
	var nodeInfo core.Node
	var err error
	nodeInfo.Id, err = os.Hostname()
	// TODO(yuval-k): unhardcode this
	if err != nil {
		nodeInfo.Id = "extauth-unknown"
	}
	nodeInfo.Cluster = "extauth"
	role := ExtAuthServerRole
	nodeInfo.Metadata = &_struct.Struct{
		Fields: map[string]*_struct.Value{
			"role": {
				Kind: &_struct.Value_StringValue{
					StringValue: role,
				},
			},
		},
	}

	settings := x.settings

	generator := config.NewGenerator(
		ctx,
		[]byte(settings.ExtAuthSettings.SigningKey),
		settings.ExtAuthSettings.UserIdHeader,
		plugins.NewPluginLoader(settings.ExtAuthSettings.PluginDirectory),
		func(cacheTtl time.Duration, oauthEndpoints config.OAuthIntrospectionEndpoints) *config.OAuthIntrospectionClients {
			httpClient := &http.Client{
				Timeout: time.Second * 10,
			}

			introspectionClient := opaque.NewIntrospectionClient(httpClient, oauthEndpoints.IntrospectionUrl)
			userInfoClient := user_info.Client(nil)
			if oauthEndpoints.UserInfoUrl != "" {
				userInfoClient = user_info.NewClient(httpClient, oauthEndpoints.UserInfoUrl, cacheTtl)
			}

			opaqueTokenValidator := opaque.NewOpaqueTokenValidator(cacheTtl, introspectionClient)
			return &config.OAuthIntrospectionClients{
				TokenValidator: opaqueTokenValidator,
				UserInfoClient: userInfoClient,
			}
		},
	)

	protoRedactor := syncutil.NewProtoRedactor()

	xdsClientLoopFunc := func(ctx context.Context) error {

		stats.Record(ctx, mExtauthConnectedStateCounter.M(int64(1)))

		client := xdsproto.NewExtAuthConfigClient(
			&nodeInfo,
			func(version string, resources []*xdsproto.ExtAuthConfig) error {

				logger := contextutils.LoggerFrom(ctx)
				logger.Infof("got %d new configs", len(resources))
				for _, resource := range resources {
					redactedJson, err := protoRedactor.BuildRedactedJsonString(resource)
					if err == nil {
						logger.Info(redactedJson)
					} else {
						logger.Warnf("Error while converting config into redacted JSON for logging: %+v", err)
					}
				}

				serverState, err := generator.GenerateConfig(resources)
				if err != nil {
					logger.Errorw("failed to generate config", zap.Any("err", err))
					return err
				}
				service.UpdateConfig(serverState)
				return nil
			},
		)

		// We are using non secure gRPC to Gloo with the assumption that it will be secured by envoy.
		// If this assumption is not correct this needs to change.
		conn, err := grpc.DialContext(ctx, settings.GlooAddress, grpc.WithInsecure())
		if err != nil {
			contextutils.LoggerFrom(ctx).Errorw("failed to create gRPC client connection to Gloo", zap.Any("error", err))
			return err
		}

		stats.Record(ctx, mExtauthConnectedState.M(int64(1)))
		defer func() {
			err = conn.Close()
			if err != nil {
				contextutils.LoggerFrom(ctx).Errorw("failed to close grpc connection", zap.Any("error", err))
			} else {
				contextutils.LoggerFrom(ctx).Infow("closed grpc connection", zap.Any("address", settings.GlooAddress))
			}
			stats.Record(ctx, mExtauthConnectedState.M(int64(0)))
		}()

		err = client.Start(ctx, conn)
		if err != nil {
			contextutils.LoggerFrom(ctx).Errorw("failed to start xDS client", zap.Any("error", err))
		} else {
			contextutils.LoggerFrom(ctx).Info("successfully started xDS client")
		}
		return err
	}

	err = contextutils.NewExponentioalBackoff(contextutils.ExponentioalBackoff{}).Backoff(ctx, xdsClientLoopFunc)
	if err == context.Canceled {
		return nil
	}
	return err
}
