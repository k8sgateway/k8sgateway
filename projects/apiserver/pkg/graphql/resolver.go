package graphql

import (
	"context"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/utils/contextutils"
	"os"
	"time"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/projects/apiserver/pkg/graphql/customtypes"
	"github.com/solo-io/solo-kit/projects/apiserver/pkg/graphql/graph"
	"github.com/solo-io/solo-kit/projects/apiserver/pkg/graphql/models"
	gatewayv1 "github.com/solo-io/solo-kit/projects/gateway/pkg/api/v1"
	"github.com/solo-io/solo-kit/projects/gloo/pkg/api/v1"
	"github.com/solo-io/solo-kit/projects/gloo/pkg/defaults"
	sqoopv1 "github.com/solo-io/solo-kit/projects/sqoop/pkg/api/v1"
)

type ApiResolver struct {
	Upstreams       v1.UpstreamClient
	Secrets         v1.SecretClient
	Artifacts       v1.ArtifactClient
	Settings        v1.SettingsClient
	VirtualServices gatewayv1.VirtualServiceClient
	ResolverMaps    sqoopv1.ResolverMapClient
	Schemas         sqoopv1.SchemaClient
}

func NewResolvers(upstreams v1.UpstreamClient,
	schemas sqoopv1.SchemaClient,
	artifacts v1.ArtifactClient,
	settings v1.SettingsClient,
	secrets v1.SecretClient,
	virtualServices gatewayv1.VirtualServiceClient,
	resolverMaps sqoopv1.ResolverMapClient) *ApiResolver {
	return &ApiResolver{
		Upstreams:       upstreams,
		VirtualServices: virtualServices,
		ResolverMaps:    resolverMaps,
		Schemas:         schemas,
		Artifacts:       artifacts,
		Settings:        settings,
		Secrets:         secrets,
		// TODO(ilackarms): just make these private functions, remove converter
	}
}

func (r *ApiResolver) Mutation() graph.MutationResolver {
	return &mutationResolver{r}
}
func (r *ApiResolver) Query() graph.QueryResolver {
	return &queryResolver{r}
}
func (r *ApiResolver) UpstreamMutation() graph.UpstreamMutationResolver {
	return &upstreamMutationResolver{r}
}
func (r *ApiResolver) UpstreamQuery() graph.UpstreamQueryResolver {
	return &upstreamQueryResolver{r}
}
func (r *ApiResolver) VirtualServiceMutation() graph.VirtualServiceMutationResolver {
	return &virtualServiceMutationResolver{r}
}
func (r *ApiResolver) VirtualServiceQuery() graph.VirtualServiceQueryResolver {
	return &virtualServiceQueryResolver{r}
}
func (r *ApiResolver) ResolverMapMutation() graph.ResolverMapMutationResolver {
	return &resolverMapMutationResolver{r}
}
func (r *ApiResolver) ResolverMapQuery() graph.ResolverMapQueryResolver {
	return &resolverMapQueryResolver{r}
}

func (r *ApiResolver) SchemaMutation() graph.SchemaMutationResolver {
	return &schemaMutationResolver{r}
}
func (r *ApiResolver) SchemaQuery() graph.SchemaQueryResolver {
	return &schemaQueryResolver{r}
}

func (r *ApiResolver) ArtifactMutation() graph.ArtifactMutationResolver {
	return &artifactMutationResolver{r}
}

func (r *ApiResolver) ArtifactQuery() graph.ArtifactQueryResolver {
	return &artifactQueryResolver{r}
}

func (r *ApiResolver) SettingsMutation() graph.SettingsMutationResolver {
	return &settingsMutationResolver{r}
}

func (r *ApiResolver) SettingsQuery() graph.SettingsQueryResolver {
	return &settingsQueryResolver{r}
}

func (r *ApiResolver) SecretMutation() graph.SecretMutationResolver {
	return &secretMutationResolver{r}
}

func (r *ApiResolver) SecretQuery() graph.SecretQueryResolver {
	return &secretQueryResolver{r}
}

type mutationResolver struct{ *ApiResolver }

func (r *mutationResolver) Upstreams(ctx context.Context, namespace string) (customtypes.UpstreamMutation, error) {
	return customtypes.UpstreamMutation{Namespace: namespace}, nil
}
func (r *mutationResolver) VirtualServices(ctx context.Context, namespace string) (customtypes.VirtualServiceMutation, error) {
	return customtypes.VirtualServiceMutation{Namespace: namespace}, nil
}
func (r *mutationResolver) ResolverMaps(ctx context.Context, namespace string) (customtypes.ResolverMapMutation, error) {
	return customtypes.ResolverMapMutation{Namespace: namespace}, nil
}
func (r *mutationResolver) Schemas(ctx context.Context, namespace string) (customtypes.SchemaMutation, error) {
	return customtypes.SchemaMutation{Namespace: namespace}, nil
}
func (r *mutationResolver) Secrets(ctx context.Context, namespace string) (customtypes.SecretMutation, error) {
	return customtypes.SecretMutation{Namespace: namespace}, nil
}
func (r *mutationResolver) Artifacts(ctx context.Context, namespace string) (customtypes.ArtifactMutation, error) {
	return customtypes.ArtifactMutation{Namespace: namespace}, nil
}
func (r *mutationResolver) Settings(ctx context.Context) (customtypes.SettingsMutation, error) {
	return customtypes.SettingsMutation{}, nil
}

type queryResolver struct{ *ApiResolver }

func (r *queryResolver) Resource(ctx context.Context, guid string) (models.Resource, error) {
	kind, namespace, name, err := resources.SplitKey(guid)
	if err != nil {
		return nil, err
	}
	switch kind {
	case resources.Kind(&v1.Upstream{}):
		return r.UpstreamQuery().Get(ctx, &customtypes.UpstreamQuery{Namespace: namespace}, name)
	case resources.Kind(&gatewayv1.VirtualService{}):
		return r.VirtualServiceQuery().Get(ctx, &customtypes.VirtualServiceQuery{Namespace: namespace}, name)
	case resources.Kind(&sqoopv1.ResolverMap{}):
		return r.ResolverMapQuery().Get(ctx, &customtypes.ResolverMapQuery{Namespace: namespace}, name)
	case resources.Kind(&sqoopv1.Schema{}):
		return r.SchemaQuery().Get(ctx, &customtypes.SchemaQuery{Namespace: namespace}, name)
	case resources.Kind(&v1.Secret{}):
		return r.SecretQuery().Get(ctx, &customtypes.SecretQuery{Namespace: namespace}, name)
	case resources.Kind(&v1.Artifact{}):
		return r.ArtifactQuery().Get(ctx, &customtypes.ArtifactQuery{Namespace: namespace}, name)
	case resources.Kind(&v1.Settings{}):
		return r.SettingsQuery().Get(ctx, &customtypes.SettingsQuery{})
	}
	return nil, errors.Errorf("unknown kind %v", kind)
}

func (r *queryResolver) GetOAuthEndpoint(ctx context.Context) (models.OAuthEndpoint, error) {
	oauthUrl := os.Getenv("OAUTH_SERVER") // ip:port of openshift server
	if oauthUrl == "" {
		return models.OAuthEndpoint{}, errors.Errorf("apiserver configured improperly, OAUTH_SERVER environment variable is not set")
	}
	oauthClient := os.Getenv("OAUTH_CLIENT") // ip:port of openshift server
	if oauthClient == "" {
		return models.OAuthEndpoint{}, errors.Errorf("apiserver configured improperly, OAUTH_CLIENT environment variable is not set")
	}
	return models.OAuthEndpoint{
		URL:        oauthUrl,
		ClientName: oauthClient,
	}, nil
}
func (r *queryResolver) Upstreams(ctx context.Context, namespace string) (customtypes.UpstreamQuery, error) {
	return customtypes.UpstreamQuery{Namespace: namespace}, nil
}
func (r *queryResolver) VirtualServices(ctx context.Context, namespace string) (customtypes.VirtualServiceQuery, error) {
	return customtypes.VirtualServiceQuery{Namespace: namespace}, nil
}
func (r *queryResolver) ResolverMaps(ctx context.Context, namespace string) (customtypes.ResolverMapQuery, error) {
	return customtypes.ResolverMapQuery{Namespace: namespace}, nil
}
func (r *queryResolver) Schemas(ctx context.Context, namespace string) (customtypes.SchemaQuery, error) {
	return customtypes.SchemaQuery{Namespace: namespace}, nil
}
func (r *queryResolver) Secrets(ctx context.Context, namespace string) (customtypes.SecretQuery, error) {
	return customtypes.SecretQuery{Namespace: namespace}, nil
}
func (r *queryResolver) Artifacts(ctx context.Context, namespace string) (customtypes.ArtifactQuery, error) {
	return customtypes.ArtifactQuery{Namespace: namespace}, nil
}
func (r *queryResolver) Settings(ctx context.Context) (customtypes.SettingsQuery, error) {
	return customtypes.SettingsQuery{}, nil
}

type upstreamMutationResolver struct{ *ApiResolver }

func (r *upstreamMutationResolver) write(overwrite bool, ctx context.Context, obj *customtypes.UpstreamMutation, upstream models.InputUpstream) (*models.Upstream, error) {
	ups, err := NewConverter(r.ApiResolver, ctx).ConvertInputUpstream(upstream)
	if err != nil {
		return nil, err
	}
	out, err := r.Upstreams.Write(ups, clients.WriteOpts{
		Ctx:               ctx,
		OverwriteExisting: overwrite,
	})
	if err != nil {
		return nil, err
	}
	return NewConverter(r.ApiResolver, ctx).ConvertOutputUpstream(out)
}

func (r *upstreamMutationResolver) Create(ctx context.Context, obj *customtypes.UpstreamMutation, upstream models.InputUpstream) (*models.Upstream, error) {
	return r.write(false, ctx, obj, upstream)
}
func (r *upstreamMutationResolver) Update(ctx context.Context, obj *customtypes.UpstreamMutation, upstream models.InputUpstream) (*models.Upstream, error) {
	return r.write(true, ctx, obj, upstream)
}
func (r *upstreamMutationResolver) Delete(ctx context.Context, obj *customtypes.UpstreamMutation, name string) (*models.Upstream, error) {
	upstream, err := r.Upstreams.Read(obj.Namespace, name, clients.ReadOpts{
		Ctx: ctx,
	})
	if err != nil {
		if errors.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	err = r.Upstreams.Delete(obj.Namespace, name, clients.DeleteOpts{Ctx: ctx})
	if err != nil {
		return nil, err
	}
	return NewConverter(r.ApiResolver, ctx).ConvertOutputUpstream(upstream)
}

type upstreamQueryResolver struct{ *ApiResolver }

func (r *upstreamQueryResolver) List(ctx context.Context, obj *customtypes.UpstreamQuery, selector *models.InputMapStringString) ([]*models.Upstream, error) {
	var convertedSelector map[string]string
	if selector != nil {
		convertedSelector = selector.GoType()
	}
	list, err := r.Upstreams.List(obj.Namespace, clients.ListOpts{
		Ctx:      ctx,
		Selector: convertedSelector,
	})
	if err != nil {
		return nil, err
	}
	return NewConverter(r.ApiResolver, ctx).ConvertOutputUpstreams(list)
}

func (r *upstreamQueryResolver) Get(ctx context.Context, obj *customtypes.UpstreamQuery, name string) (*models.Upstream, error) {
	upstream, err := r.Upstreams.Read(obj.Namespace, name, clients.ReadOpts{
		Ctx: ctx,
	})
	if err != nil {
		return nil, err
	}
	return NewConverter(r.ApiResolver, ctx).ConvertOutputUpstream(upstream)
}

type resolverMapMutationResolver struct{ *ApiResolver }

func (r *resolverMapMutationResolver) SetResolver(ctx context.Context, obj *customtypes.ResolverMapMutation, resolverMapName, resourceVersion, typeName, fieldName string, resolver models.InputGlooResolver) (*models.ResolverMap, error) {
	v1Resolver, err := ConvertInputResolver(models.InputResolver{GlooResolver: &resolver})
	if err != nil {
		return nil, err
	}

	resolverMap, err := r.ResolverMaps.Read(obj.Namespace, resolverMapName, clients.ReadOpts{Ctx: ctx})
	if err != nil {
		return nil, err
	}
	if resolverMap.Metadata.ResourceVersion != resourceVersion {
		return nil, errors.Errorf("resource version mismatch. received %v, want %v", resourceVersion, resolverMap.Metadata.ResourceVersion)
	}

	typResolver, ok := resolverMap.Types[typeName]
	if !ok {
		return nil, errors.Errorf("no type %v in resolver map %v", typeName, resolverMapName)
	}
	typResolver.Fields[fieldName] = v1Resolver

	out, err := r.ResolverMaps.Write(resolverMap, clients.WriteOpts{
		Ctx:               ctx,
		OverwriteExisting: true,
	})
	if err != nil {
		return nil, err
	}
	return NewConverter(r.ApiResolver, ctx).ConvertOutputResolverMap(out)
}

func (r *resolverMapMutationResolver) write(overwrite bool, ctx context.Context, obj *customtypes.ResolverMapMutation, resolverMap models.InputResolverMap) (*models.ResolverMap, error) {
	ups, err := NewConverter(r.ApiResolver, ctx).ConvertInputResolverMap(resolverMap)
	if err != nil {
		return nil, err
	}
	out, err := r.ResolverMaps.Write(ups, clients.WriteOpts{
		Ctx:               ctx,
		OverwriteExisting: overwrite,
	})
	if err != nil {
		return nil, err
	}
	return NewConverter(r.ApiResolver, ctx).ConvertOutputResolverMap(out)
}

func (r *resolverMapMutationResolver) Create(ctx context.Context, obj *customtypes.ResolverMapMutation, resolverMap models.InputResolverMap) (*models.ResolverMap, error) {
	return r.write(false, ctx, obj, resolverMap)
}
func (r *resolverMapMutationResolver) Update(ctx context.Context, obj *customtypes.ResolverMapMutation, resolverMap models.InputResolverMap) (*models.ResolverMap, error) {
	return r.write(true, ctx, obj, resolverMap)
}
func (r *resolverMapMutationResolver) Delete(ctx context.Context, obj *customtypes.ResolverMapMutation, name string) (*models.ResolverMap, error) {
	resolverMap, err := r.ResolverMaps.Read(obj.Namespace, name, clients.ReadOpts{
		Ctx: ctx,
	})
	if err != nil {
		if errors.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	err = r.ResolverMaps.Delete(obj.Namespace, name, clients.DeleteOpts{Ctx: ctx})
	if err != nil {
		return nil, err
	}
	return NewConverter(r.ApiResolver, ctx).ConvertOutputResolverMap(resolverMap)
}

type resolverMapQueryResolver struct{ *ApiResolver }

func (r *resolverMapQueryResolver) List(ctx context.Context, obj *customtypes.ResolverMapQuery, selector *models.InputMapStringString) ([]*models.ResolverMap, error) {
	var convertedSelector map[string]string
	if selector != nil {
		convertedSelector = selector.GoType()
	}
	list, err := r.ResolverMaps.List(obj.Namespace, clients.ListOpts{
		Ctx:      ctx,
		Selector: convertedSelector,
	})
	if err != nil {
		return nil, err
	}
	return NewConverter(r.ApiResolver, ctx).ConvertOutputResolverMaps(list)
}

func (r *resolverMapQueryResolver) Get(ctx context.Context, obj *customtypes.ResolverMapQuery, name string) (*models.ResolverMap, error) {
	resolverMap, err := r.ResolverMaps.Read(obj.Namespace, name, clients.ReadOpts{
		Ctx: ctx,
	})
	if err != nil {
		return nil, err
	}
	return NewConverter(r.ApiResolver, ctx).ConvertOutputResolverMap(resolverMap)
}

type schemaMutationResolver struct{ *ApiResolver }

func (r *schemaMutationResolver) write(overwrite bool, ctx context.Context, obj *customtypes.SchemaMutation, schema models.InputSchema) (*models.Schema, error) {
	ups, err := NewConverter(r.ApiResolver, ctx).ConvertInputSchema(schema)
	if err != nil {
		return nil, err
	}
	out, err := r.Schemas.Write(ups, clients.WriteOpts{
		Ctx:               ctx,
		OverwriteExisting: overwrite,
	})
	if err != nil {
		return nil, err
	}
	return NewConverter(r.ApiResolver, ctx).ConvertOutputSchema(out), nil
}

func (r *schemaMutationResolver) Create(ctx context.Context, obj *customtypes.SchemaMutation, schema models.InputSchema) (*models.Schema, error) {
	return r.write(false, ctx, obj, schema)
}
func (r *schemaMutationResolver) Update(ctx context.Context, obj *customtypes.SchemaMutation, schema models.InputSchema) (*models.Schema, error) {
	return r.write(true, ctx, obj, schema)
}
func (r *schemaMutationResolver) Delete(ctx context.Context, obj *customtypes.SchemaMutation, name string) (*models.Schema, error) {
	schema, err := r.Schemas.Read(obj.Namespace, name, clients.ReadOpts{
		Ctx: ctx,
	})
	if err != nil {
		if errors.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	err = r.Schemas.Delete(obj.Namespace, name, clients.DeleteOpts{Ctx: ctx})
	if err != nil {
		return nil, err
	}
	return NewConverter(r.ApiResolver, ctx).ConvertOutputSchema(schema), nil
}

type schemaQueryResolver struct{ *ApiResolver }

func (r *schemaQueryResolver) List(ctx context.Context, obj *customtypes.SchemaQuery, selector *models.InputMapStringString) ([]*models.Schema, error) {
	var convertedSelector map[string]string
	if selector != nil {
		convertedSelector = selector.GoType()
	}
	list, err := r.Schemas.List(obj.Namespace, clients.ListOpts{
		Ctx:      ctx,
		Selector: convertedSelector,
	})
	if err != nil {
		return nil, err
	}
	return NewConverter(r.ApiResolver, ctx).ConvertOutputSchemas(list), nil
}

func (r *schemaQueryResolver) Get(ctx context.Context, obj *customtypes.SchemaQuery, name string) (*models.Schema, error) {
	schema, err := r.Schemas.Read(obj.Namespace, name, clients.ReadOpts{
		Ctx: ctx,
	})
	if err != nil {
		return nil, err
	}
	return NewConverter(r.ApiResolver, ctx).ConvertOutputSchema(schema), nil
}

type secretMutationResolver struct{ *ApiResolver }

func (r *secretMutationResolver) write(overwrite bool, ctx context.Context, obj *customtypes.SecretMutation, secret models.InputSecret) (*models.Secret, error) {
	ups, err := NewConverter(r.ApiResolver, ctx).ConvertInputSecret(secret)
	if err != nil {
		return nil, err
	}
	out, err := r.Secrets.Write(ups, clients.WriteOpts{
		Ctx:               ctx,
		OverwriteExisting: overwrite,
	})
	if err != nil {
		return nil, err
	}
	return NewConverter(r.ApiResolver, ctx).ConvertOutputSecret(out), nil
}

func (r *secretMutationResolver) Create(ctx context.Context, obj *customtypes.SecretMutation, secret models.InputSecret) (*models.Secret, error) {
	return r.write(false, ctx, obj, secret)
}
func (r *secretMutationResolver) Update(ctx context.Context, obj *customtypes.SecretMutation, secret models.InputSecret) (*models.Secret, error) {
	return r.write(true, ctx, obj, secret)
}
func (r *secretMutationResolver) Delete(ctx context.Context, obj *customtypes.SecretMutation, name string) (*models.Secret, error) {
	secret, err := r.Secrets.Read(obj.Namespace, name, clients.ReadOpts{
		Ctx: ctx,
	})
	if err != nil {
		if errors.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	err = r.Secrets.Delete(obj.Namespace, name, clients.DeleteOpts{Ctx: ctx})
	if err != nil {
		return nil, err
	}
	return NewConverter(r.ApiResolver, ctx).ConvertOutputSecret(secret), nil
}

type secretQueryResolver struct{ *ApiResolver }

func (r *secretQueryResolver) List(ctx context.Context, obj *customtypes.SecretQuery, selector *models.InputMapStringString) ([]*models.Secret, error) {
	var convertedSelector map[string]string
	if selector != nil {
		convertedSelector = selector.GoType()
	}
	list, err := r.Secrets.List(obj.Namespace, clients.ListOpts{
		Ctx:      ctx,
		Selector: convertedSelector,
	})
	if err != nil {
		return nil, err
	}
	return NewConverter(r.ApiResolver, ctx).ConvertOutputSecrets(list), nil
}

func (r *secretQueryResolver) Get(ctx context.Context, obj *customtypes.SecretQuery, name string) (*models.Secret, error) {
	secret, err := r.Secrets.Read(obj.Namespace, name, clients.ReadOpts{
		Ctx: ctx,
	})
	if err != nil {
		return nil, err
	}
	return NewConverter(r.ApiResolver, ctx).ConvertOutputSecret(secret), nil
}

type artifactMutationResolver struct{ *ApiResolver }

func (r *artifactMutationResolver) write(overwrite bool, ctx context.Context, obj *customtypes.ArtifactMutation, artifact models.InputArtifact) (*models.Artifact, error) {
	ups, err := NewConverter(r.ApiResolver, ctx).ConvertInputArtifact(artifact)
	if err != nil {
		return nil, err
	}
	out, err := r.Artifacts.Write(ups, clients.WriteOpts{
		Ctx:               ctx,
		OverwriteExisting: overwrite,
	})
	if err != nil {
		return nil, err
	}
	return NewConverter(r.ApiResolver, ctx).ConvertOutputArtifact(out), nil
}

func (r *artifactMutationResolver) Create(ctx context.Context, obj *customtypes.ArtifactMutation, artifact models.InputArtifact) (*models.Artifact, error) {
	return r.write(false, ctx, obj, artifact)
}
func (r *artifactMutationResolver) Update(ctx context.Context, obj *customtypes.ArtifactMutation, artifact models.InputArtifact) (*models.Artifact, error) {
	return r.write(true, ctx, obj, artifact)
}
func (r *artifactMutationResolver) Delete(ctx context.Context, obj *customtypes.ArtifactMutation, name string) (*models.Artifact, error) {
	artifact, err := r.Artifacts.Read(obj.Namespace, name, clients.ReadOpts{
		Ctx: ctx,
	})
	if err != nil {
		if errors.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	err = r.Artifacts.Delete(obj.Namespace, name, clients.DeleteOpts{Ctx: ctx})
	if err != nil {
		return nil, err
	}
	return NewConverter(r.ApiResolver, ctx).ConvertOutputArtifact(artifact), nil
}

type settingsMutationResolver struct{ *ApiResolver }

func (r *settingsMutationResolver) write(overwrite bool, ctx context.Context, obj *customtypes.SettingsMutation, settings *v1.Settings) (*models.Settings, error) {
	out, err := r.Settings.Write(settings, clients.WriteOpts{
		Ctx:               ctx,
		OverwriteExisting: overwrite,
	})
	if err != nil {
		return nil, err
	}
	return NewConverter(r.ApiResolver, ctx).ConvertOutputSettings(out), nil
}
func (r *settingsMutationResolver) Update(ctx context.Context, obj *customtypes.SettingsMutation, rawUpdates models.InputSettings) (*models.Settings, error) {
	updates, err := NewConverter(r.ApiResolver, ctx).ConvertInputSettings(rawUpdates)
	if err != nil {
		return nil, err
	}

	namespace := updates.Metadata.Namespace
	name := updates.Metadata.Name
	settings, err := r.Settings.Read(namespace, name, clients.ReadOpts{
		Ctx: ctx,
	})
	if err != nil {
		return nil, err
	}

	// preserve the given metadata to ensure request was made with latest resourceVersion
	settings.Metadata = updates.Metadata

	// only apply changes to the provided fields
	if updates.RefreshRate != nil {
		settings.RefreshRate = updates.RefreshRate
	}
	if updates.WatchNamespaces != nil {
		settings.WatchNamespaces = updates.WatchNamespaces
	}
	return r.write(true, ctx, obj, settings)
}

type artifactQueryResolver struct{ *ApiResolver }

func (r *artifactQueryResolver) List(ctx context.Context, obj *customtypes.ArtifactQuery, selector *models.InputMapStringString) ([]*models.Artifact, error) {
	var convertedSelector map[string]string
	if selector != nil {
		convertedSelector = selector.GoType()
	}
	list, err := r.Artifacts.List(obj.Namespace, clients.ListOpts{
		Ctx:      ctx,
		Selector: convertedSelector,
	})
	if err != nil {
		return nil, err
	}
	return NewConverter(r.ApiResolver, ctx).ConvertOutputArtifacts(list), nil
}

func (r *artifactQueryResolver) Get(ctx context.Context, obj *customtypes.ArtifactQuery, name string) (*models.Artifact, error) {
	artifact, err := r.Artifacts.Read(obj.Namespace, name, clients.ReadOpts{
		Ctx: ctx,
	})
	if err != nil {
		return nil, err
	}
	return NewConverter(r.ApiResolver, ctx).ConvertOutputArtifact(artifact), nil
}

type settingsQueryResolver struct{ *ApiResolver }

func (r *settingsQueryResolver) Get(ctx context.Context, obj *customtypes.SettingsQuery) (*models.Settings, error) {
	namespace := defaults.GlooSystem
	name := defaults.SettingsName
	settings, err := r.Settings.Read(namespace, name, clients.ReadOpts{
		Ctx: ctx,
	})
	if err != nil {
		return nil, err
	}
	return NewConverter(r.ApiResolver, ctx).ConvertOutputSettings(settings), nil
}

func (r *ApiResolver) Subscription() graph.SubscriptionResolver {
	return &subscriptionResolver{r}
}

type subscriptionResolver struct{ *ApiResolver }

func (r subscriptionResolver) Upstreams(ctx context.Context, namespace string, selector *models.InputMapStringString) (<-chan []*models.Upstream, error) {
	var convertedSelector map[string]string
	if selector != nil {
		convertedSelector = selector.GoType()
	}
	watch, errs, err := r.ApiResolver.Upstreams.Watch(namespace, clients.WatchOpts{
		// TODO(ilackarms): refresh rate
		RefreshRate: time.Minute * 10,
		Ctx:         ctx,
		Selector:    convertedSelector,
	})
	if err != nil {
		return nil, err
	}
	upstreamsChan := make(chan []*models.Upstream)
	go func() {
		defer close(upstreamsChan)
		for {
			select {
			case list, ok := <-watch:
				if !ok {
					return
				}
				upstreams, err := NewConverter(r.ApiResolver, ctx).ConvertOutputUpstreams(list)
				if err != nil {
					// TODO(mitchdraft) log this
					return
				}
				select {
				case upstreamsChan <- upstreams:
				default:
					contextutils.LoggerFrom(ctx).Errorf("upstream channel full, cannot send list of %v upstreams", len(list))
				}
			case err, ok := <-errs:
				if !ok {
					return
				}
				contextutils.LoggerFrom(ctx).Errorf("error in upstream subscription: %v", err)
			case <-ctx.Done():
				return
			}
		}
	}()

	return upstreamsChan, nil
}

func (r subscriptionResolver) VirtualServices(ctx context.Context, namespace string, selector *models.InputMapStringString) (<-chan []*models.VirtualService, error) {
	var convertedSelector map[string]string
	if selector != nil {
		convertedSelector = selector.GoType()
	}
	watch, errs, err := r.ApiResolver.VirtualServices.Watch(namespace, clients.WatchOpts{
		// TODO(ilackarms): refresh rate
		RefreshRate: time.Minute * 10,
		Ctx:         ctx,
		Selector:    convertedSelector,
	})
	if err != nil {
		return nil, err
	}
	virtualServicesChan := make(chan []*models.VirtualService)
	go func() {
		defer close(virtualServicesChan)
		for {
			select {
			case list, ok := <-watch:
				if !ok {
					return
				}
				virtualServices, err := NewConverter(r.ApiResolver, ctx).ConvertOutputVirtualServices(list)
				if err != nil {
					// TODO(mitchdraft) log this
					return
				}
				select {
				case virtualServicesChan <- virtualServices:
				default:
					contextutils.LoggerFrom(ctx).Errorf("virtualService channel full, cannot send list of %v virtualServices", len(list))
				}
			case err, ok := <-errs:
				if !ok {
					return
				}
				contextutils.LoggerFrom(ctx).Errorf("error in virtualService subscription: %v", err)
			case <-ctx.Done():
				return
			}
		}
	}()

	return virtualServicesChan, nil
}
