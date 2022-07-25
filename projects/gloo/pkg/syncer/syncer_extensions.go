package syncer

import (
	"context"

	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ratelimit"
	v1snap "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/gloosnapshot"
	envoycache "github.com/solo-io/solo-kit/pkg/api/v1/control-plane/cache"
	"github.com/solo-io/solo-kit/pkg/api/v2/reporter"
)

// TranslatorSyncerExtension represents a custom sync behavior that updates an entry in the SnapshotCache
type TranslatorSyncerExtension interface {
	// ID returns the unique identifier for this TranslatorSyncerExtension
	// This represents the Key in the SnapshotCache where Sync() will store results
	ID() string

	// Sync processes an ApiSnapshot and updates reports with Errors/Warnings that it encounters
	// and updates the SnapshotCache entry if possible
	Sync(
		ctx context.Context,
		snap *v1snap.ApiSnapshot,
		settings *v1.Settings,
		xdsCache envoycache.SnapshotCache,
		reports reporter.ResourceReports)
}

type TranslatorSyncerExtensionParams struct {
	Hasher                   func(resources []envoycache.Resource) uint64
	RateLimitServiceSettings ratelimit.ServiceSettings
}

// TranslatorSyncerExtensionFactory generates TranslatorSyncerExtensions
type TranslatorSyncerExtensionFactory func(context.Context, TranslatorSyncerExtensionParams) (TranslatorSyncerExtension, error)
