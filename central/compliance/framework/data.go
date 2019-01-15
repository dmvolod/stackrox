package framework

import (
	"github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
)

//go:generate mockgen-wrapper ComplianceDataRepository

// ComplianceDataRepository is the unified interface for accessing all the data that might be relevant for a compliance
// run. This provides check implementors with a unified view of all data objects regardless of their source (stored by
// central vs. obtained specifically for a compliance run), and also allows presenting a stable snapshot to all checks
// to reduce the risk of inconsistencies.
type ComplianceDataRepository interface {
	Cluster() *storage.Cluster
	Nodes() map[string]*storage.Node
	Deployments() map[string]*storage.Deployment

	NetworkPolicies() map[string]*storage.NetworkPolicy
	NetworkGraph() *v1.NetworkGraph
	Policies() map[string]*storage.Policy
	ImageIntegrations() []*storage.ImageIntegration
}
