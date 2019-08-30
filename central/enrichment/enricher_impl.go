package enrichment

import (
	"context"

	"github.com/stackrox/rox/central/image/datastore"
	"github.com/stackrox/rox/central/role/resources"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/images/enricher"
	"github.com/stackrox/rox/pkg/images/types"
	"github.com/stackrox/rox/pkg/sac"
)

var (
	getImageContext = sac.WithGlobalAccessScopeChecker(context.Background(),
		sac.AllowFixedScopes(
			sac.AccessModeScopeKeys(storage.Access_READ_ACCESS),
			sac.ResourceScopeKeys(resources.Image)))
)

// enricherImpl enriches images with data from registries and scanners.
type enricherImpl struct {
	imageEnricher enricher.ImageEnricher
	images        datastore.DataStore
}

// EnrichDeployment enriches a deployment with data from registries and scanners.
func (e *enricherImpl) EnrichDeployment(ctx enricher.EnrichmentContext, deployment *storage.Deployment) (images []*storage.Image, updatedIndices []int, pendingEnrichment bool, err error) {
	for i, c := range deployment.GetContainers() {
		var imgToProcess *storage.Image
		if !ctx.IgnoreExisting && c.GetImage().GetId() != "" {
			var img *storage.Image
			img, _, err = e.images.GetImage(getImageContext, c.GetImage().GetId())
			if err != nil {
				return
			}
			imgToProcess = img
		}
		if imgToProcess == nil {
			imgToProcess = types.ToImage(c.GetImage())
		}
		images = append(images, imgToProcess)
		// If an ID was found and the image is not pullable, then don't try to get metadata because it won't
		// be available
		if imgToProcess.GetId() != "" && !imgToProcess.GetPullable() {
			continue
		}
		enrichmentResult := e.imageEnricher.EnrichImage(ctx, imgToProcess)
		if enrichmentResult.ImageUpdated && imgToProcess.GetId() != "" {
			updatedIndices = append(updatedIndices, i)
		}
		if enrichmentResult.ScanResult == enricher.ScanTriggered {
			pendingEnrichment = true
		}
	}
	return
}
