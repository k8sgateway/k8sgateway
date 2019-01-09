package setuputils

import (
	"context"

	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/memory"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
)

// tell us how to setup
type SetupFunc func(ctx context.Context,
	kubeCache kube.SharedCache,
	inMemoryCache memory.InMemoryResourceCache,
	settings *v1.Settings) error

type SetupSyncer struct {
	settingsRef   core.ResourceRef
	setupFunc     SetupFunc
	kubeCache     kube.SharedCache
	inMemoryCache memory.InMemoryResourceCache
}

func NewSetupSyncer(settingsRef core.ResourceRef, setupFunc SetupFunc) *SetupSyncer {
	return &SetupSyncer{
		settingsRef:   settingsRef,
		setupFunc:     setupFunc,
		kubeCache:     kube.NewKubeCache(),
		inMemoryCache: memory.NewInMemoryResourceCache(),
	}
}

func (s *SetupSyncer) Sync(ctx context.Context, snap *v1.SetupSnapshot) error {
	settings, err := snap.Settings.List().Find(s.settingsRef.Strings())
	if err != nil {
		return errors.Wrapf(err, "finding bootstrap configuration")
	}
	return s.setupFunc(ctx, s.kubeCache, s.inMemoryCache, settings)
}
