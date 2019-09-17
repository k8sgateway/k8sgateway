package proxysvc

import (
	"context"

	"github.com/solo-io/solo-projects/projects/grpcserver/server/internal/client"
	"github.com/solo-io/solo-projects/projects/grpcserver/server/internal/settings"

	"github.com/solo-io/solo-projects/projects/grpcserver/server/helpers/status"

	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	v1 "github.com/solo-io/solo-projects/projects/grpcserver/api/v1"
	"github.com/solo-io/solo-projects/projects/grpcserver/server/helpers/rawgetter"
	"go.uber.org/zap"
)

type proxyGrpcService struct {
	ctx             context.Context
	clientCache     client.ClientCache
	rawGetter       rawgetter.RawGetter
	statusConverter status.InputResourceStatusGetter
	settingsValues  settings.ValuesClient
}

func NewProxyGrpcService(ctx context.Context, clientCache client.ClientCache, rawGetter rawgetter.RawGetter, statusConverter status.InputResourceStatusGetter, settingsValues settings.ValuesClient) v1.ProxyApiServer {
	return &proxyGrpcService{
		ctx:             ctx,
		clientCache:     clientCache,
		rawGetter:       rawGetter,
		statusConverter: statusConverter,
		settingsValues:  settingsValues,
	}
}

func (s *proxyGrpcService) GetProxy(ctx context.Context, request *v1.GetProxyRequest) (*v1.GetProxyResponse, error) {
	proxy, err := s.clientCache.GetProxyClient().Read(request.GetRef().GetNamespace(), request.GetRef().GetName(), clients.ReadOpts{Ctx: s.ctx})
	if err != nil {
		wrapped := FailedToGetProxyError(err, request.GetRef())
		contextutils.LoggerFrom(s.ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
		return nil, wrapped
	}
	return &v1.GetProxyResponse{ProxyDetails: s.getDetails(proxy)}, nil
}

func (s *proxyGrpcService) ListProxies(ctx context.Context, request *v1.ListProxiesRequest) (*v1.ListProxiesResponse, error) {
	var proxyDetailsList []*v1.ProxyDetails
	for _, ns := range s.settingsValues.GetWatchNamespaces() {
		proxiesInNamespace, err := s.clientCache.GetProxyClient().List(ns, clients.ListOpts{Ctx: s.ctx})
		if err != nil {
			wrapped := FailedToListProxiesError(err, ns)
			contextutils.LoggerFrom(s.ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
			return nil, wrapped
		}
		for _, p := range proxiesInNamespace {
			proxyDetailsList = append(proxyDetailsList, s.getDetails(p))
		}
	}
	return &v1.ListProxiesResponse{ProxyDetails: proxyDetailsList}, nil
}

func (s *proxyGrpcService) getDetails(proxy *gloov1.Proxy) *v1.ProxyDetails {
	return &v1.ProxyDetails{
		Proxy:  proxy,
		Raw:    s.rawGetter.GetRaw(s.ctx, proxy, gloov1.ProxyCrd),
		Status: s.statusConverter.GetApiStatusFromResource(proxy),
	}
}
