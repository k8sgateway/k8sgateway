package metricsservice

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8s "k8s.io/client-go/kubernetes/typed/core/v1"
)

//go:generate mockgen -destination mocks/mock_storage.go -package mocks github.com/solo-io/gloo/projects/metrics/pkg/metricsservice Storage

type Storage interface {
	RecordUsage(ctx context.Context, usage *GlobalUsage) error
	GetUsage(ctx context.Context) (*GlobalUsage, error)
}

type EnvoyMetrics struct {
	HttpRequests   uint64
	TcpConnections uint64
	Uptime         time.Duration
}

type EnvoyUsage struct {
	EnvoyMetrics    *EnvoyMetrics
	LastRecordedAt  time.Time
	FirstRecordedAt time.Time
	Active          bool // whether or not we believe this envoy to be active
}

type GlobalUsage struct {
	EnvoyIdToUsage map[string]*EnvoyUsage
}

type configMapStorage struct {
	configMapClient     k8s.ConfigMapInterface
	podNamespace        string
	currentTimeProvider CurrentTimeProvider

	// we may be receiving metrics from several envoys at the same time
	// be sure to lock appropriately to prevent data loss
	mutex sync.RWMutex
}

var _ Storage = &configMapStorage{}

const (
	metricsConfigMapName = "gloo-usage"
	usageDataKey         = "USAGE_DATA"

	// allow this much time between what we estimate for envoy's uptime and what it actually reports
	uptimeDiffThreshold = time.Second * 1

	// envoy should do a stats push every five seconds
	// if we go ten cycles without a stats push, then consider that envoy inactive
	envoyExpiryDuration = time.Second * 50
)

//go:generate mockgen -destination mocks/mock_config_map_client.go -package mocks k8s.io/client-go/kubernetes/typed/core/v1 ConfigMapInterface

func NewConfigMapStorage(podNamespace string, configMapClient k8s.ConfigMapInterface) Storage {
	return &configMapStorage{
		configMapClient:     configMapClient,
		podNamespace:        podNamespace,
		currentTimeProvider: time.Now,
		mutex:               sync.RWMutex{},
	}
}

// visible for testing
// provide a way to get the current time to make unit tests easier to write and more deterministic
func newConfigMapStorageWithTime(podNamespace string, configMapClient k8s.ConfigMapInterface, currentTimeProvider CurrentTimeProvider) Storage {
	return &configMapStorage{
		configMapClient:     configMapClient,
		podNamespace:        podNamespace,
		currentTimeProvider: currentTimeProvider,
		mutex:               sync.RWMutex{},
	}
}

// Record a new set of metrics for the given envoy instance id
// The envoy instance id template is set in the gateway proxy configmap: `envoy.yaml`.node.id
func (s *configMapStorage) RecordUsage(ctx context.Context, usage *GlobalUsage) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	_, configMap, err := s.getExistingUsage(ctx)
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(usage)
	if err != nil {
		return err
	}
	configMap.Data = map[string]string{usageDataKey: string(bytes)}

	_, err = s.configMapClient.Update(configMap)
	return err
}

func (s *configMapStorage) GetUsage(ctx context.Context) (*GlobalUsage, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	existingUsage, _, err := s.getExistingUsage(ctx)
	if err != nil {
		return nil, err
	}

	return existingUsage, nil
}

// returns the old usage, the config map it came from, and any error
// the config map is nil if and only if an error occurs
// the old usage is nil if it has not been written yet or if there was an error reading it
func (s *configMapStorage) getExistingUsage(ctx context.Context) (*GlobalUsage, *corev1.ConfigMap, error) {
	cm, err := s.configMapClient.Get(metricsConfigMapName, v1.GetOptions{})
	if err != nil {
		return nil, nil, err
	}

	usageJson, ok := cm.Data[usageDataKey]

	if !ok || usageJson == "" {
		return nil, cm, nil
	}

	usage := &GlobalUsage{}

	err = json.Unmarshal([]byte(usageJson), &usage)
	if err != nil {
		return nil, nil, err
	}

	return usage, cm, nil
}
