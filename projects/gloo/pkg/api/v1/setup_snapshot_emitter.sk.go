// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"sync"
	"time"

	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"

	"github.com/solo-io/go-utils/errutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/errors"
)

var (
	mSetupSnapshotIn  = stats.Int64("setup.gloo.solo.io/snap_emitter/snap_in", "The number of snapshots in", "1")
	mSetupSnapshotOut = stats.Int64("setup.gloo.solo.io/snap_emitter/snap_out", "The number of snapshots out", "1")

	setupsnapshotInView = &view.View{
		Name:        "setup.gloo.solo.io_snap_emitter/snap_in",
		Measure:     mSetupSnapshotIn,
		Description: "The number of snapshots updates coming in",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}
	setupsnapshotOutView = &view.View{
		Name:        "setup.gloo.solo.io/snap_emitter/snap_out",
		Measure:     mSetupSnapshotOut,
		Description: "The number of snapshots updates going out",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}
)

func init() {
	view.Register(setupsnapshotInView, setupsnapshotOutView)
}

type SetupEmitter interface {
	Register() error
	Settings() SettingsClient
	Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *SetupSnapshot, <-chan error, error)
}

func NewSetupEmitter(settingsClient SettingsClient) SetupEmitter {
	return NewSetupEmitterWithEmit(settingsClient, make(chan struct{}))
}

func NewSetupEmitterWithEmit(settingsClient SettingsClient, emit <-chan struct{}) SetupEmitter {
	return &setupEmitter{
		settings:  settingsClient,
		forceEmit: emit,
	}
}

type setupEmitter struct {
	forceEmit <-chan struct{}
	settings  SettingsClient
}

func (c *setupEmitter) Register() error {
	if err := c.settings.Register(); err != nil {
		return err
	}
	return nil
}

func (c *setupEmitter) Settings() SettingsClient {
	return c.settings
}

func (c *setupEmitter) Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *SetupSnapshot, <-chan error, error) {

	if len(watchNamespaces) == 0 {
		watchNamespaces = []string{""}
	}

	for _, ns := range watchNamespaces {
		if ns == "" && len(watchNamespaces) > 1 {
			return nil, nil, errors.Errorf("the \"\" namespace is used to watch all namespaces. Snapshots can either be tracked for " +
				"specific namespaces or \"\" AllNamespaces, but not both.")
		}
	}

	errs := make(chan error)
	var done sync.WaitGroup
	ctx := opts.Ctx
	/* Create channel for Settings */
	type settingsListWithNamespace struct {
		list      SettingsList
		namespace string
	}
	settingsChan := make(chan settingsListWithNamespace)

	for _, namespace := range watchNamespaces {
		/* Setup namespaced watch for Settings */
		settingsNamespacesChan, settingsErrs, err := c.settings.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting Settings watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, settingsErrs, namespace+"-settings")
		}(namespace)

		/* Watch for changes and update snapshot */
		go func(namespace string) {
			for {
				select {
				case <-ctx.Done():
					return
				case settingsList := <-settingsNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case settingsChan <- settingsListWithNamespace{list: settingsList, namespace: namespace}:
					}
				}
			}
		}(namespace)
	}

	snapshots := make(chan *SetupSnapshot)
	go func() {
		originalSnapshot := SetupSnapshot{}
		currentSnapshot := originalSnapshot.Clone()
		timer := time.NewTicker(time.Second * 1)
		sync := func() {
			if originalSnapshot.Hash() == currentSnapshot.Hash() {
				return
			}

			stats.Record(ctx, mSetupSnapshotOut.M(1))
			originalSnapshot = currentSnapshot.Clone()
			sentSnapshot := currentSnapshot.Clone()
			snapshots <- &sentSnapshot
		}
		settingsByNamespace := make(map[string]SettingsList)

		for {
			record := func() { stats.Record(ctx, mSetupSnapshotIn.M(1)) }

			select {
			case <-timer.C:
				sync()
			case <-ctx.Done():
				close(snapshots)
				done.Wait()
				close(errs)
				return
			case <-c.forceEmit:
				sentSnapshot := currentSnapshot.Clone()
				snapshots <- &sentSnapshot
			case settingsNamespacedList := <-settingsChan:
				record()

				namespace := settingsNamespacedList.namespace

				// merge lists by namespace
				settingsByNamespace[namespace] = settingsNamespacedList.list
				var settingsList SettingsList
				for _, settings := range settingsByNamespace {
					settingsList = append(settingsList, settings...)
				}
				currentSnapshot.Settings = settingsList.Sort()
			}
		}
	}()
	return snapshots, errs, nil
}
