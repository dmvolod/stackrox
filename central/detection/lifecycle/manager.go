package lifecycle

import (
	"time"

	deploymentDatastore "github.com/stackrox/rox/central/deployment/datastore"
	"github.com/stackrox/rox/central/detection/alertmanager"
	"github.com/stackrox/rox/central/detection/deploytime"
	"github.com/stackrox/rox/central/detection/runtime"
	processDatastore "github.com/stackrox/rox/central/processindicator/datastore"
	whitelistDataStore "github.com/stackrox/rox/central/processwhitelist/datastore"
	"github.com/stackrox/rox/central/reprocessor"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/expiringcache"
	"github.com/stackrox/rox/pkg/logging"
	"github.com/stackrox/rox/pkg/process/filter"
	"github.com/stackrox/rox/pkg/set"
	"golang.org/x/time/rate"
)

const (
	rateLimitDuration            = 10 * time.Second
	indicatorFlushTickerDuration = 1 * time.Minute
)

var (
	log = logging.LoggerForModule()
)

// A Manager manages deployment/policy lifecycle updates.
//go:generate mockgen-wrapper
type Manager interface {
	IndicatorAdded(indicator *storage.ProcessIndicator) error
	UpsertPolicy(policy *storage.Policy) error
	HandleAlerts(deploymentID string, alerts []*storage.Alert, stage storage.LifecycleStage) error
	DeploymentRemoved(deployment *storage.Deployment) error
	RemovePolicy(policyID string) error
}

// newManager returns a new manager with the injected dependencies.
func newManager(deploytimeDetector deploytime.Detector, runtimeDetector runtime.Detector,
	deploymentDatastore deploymentDatastore.DataStore, processesDataStore processDatastore.DataStore, whitelists whitelistDataStore.DataStore,
	alertManager alertmanager.AlertManager, reprocessor reprocessor.Loop, deletedDeploymentsCache expiringcache.Cache, filter filter.Filter) *managerImpl {
	m := &managerImpl{
		deploytimeDetector:      deploytimeDetector,
		runtimeDetector:         runtimeDetector,
		alertManager:            alertManager,
		deploymentDataStore:     deploymentDatastore,
		processesDataStore:      processesDataStore,
		whitelists:              whitelists,
		reprocessor:             reprocessor,
		deletedDeploymentsCache: deletedDeploymentsCache,
		processFilter:           filter,

		queuedIndicators: make(map[string]*storage.ProcessIndicator),

		indicatorRateLimiter: rate.NewLimiter(rate.Every(rateLimitDuration), 5),
		indicatorFlushTicker: time.NewTicker(indicatorFlushTickerDuration),

		removedPolicies: set.NewStringSet(),
	}

	go m.flushQueuePeriodically()
	return m
}
