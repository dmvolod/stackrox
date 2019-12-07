package datastore

import (
	"context"

	"github.com/blevesearch/bleve"
	"github.com/dgraph-io/badger"
	"github.com/pkg/errors"
	"github.com/stackrox/rox/central/image/datastore/internal/search"
	"github.com/stackrox/rox/central/image/datastore/internal/store"
	badgerStore "github.com/stackrox/rox/central/image/datastore/internal/store/badger"
	"github.com/stackrox/rox/central/image/index"
	riskDS "github.com/stackrox/rox/central/risk/datastore"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	searchPkg "github.com/stackrox/rox/pkg/search"
)

// DataStore is an intermediary to AlertStorage.
//go:generate mockgen-wrapper DataStore
type DataStore interface {
	SearchListImages(ctx context.Context, q *v1.Query) ([]*storage.ListImage, error)
	ListImage(ctx context.Context, sha string) (*storage.ListImage, bool, error)

	Search(ctx context.Context, q *v1.Query) ([]searchPkg.Result, error)
	SearchImages(ctx context.Context, q *v1.Query) ([]*v1.SearchResult, error)
	SearchRawImages(ctx context.Context, q *v1.Query) ([]*storage.Image, error)

	CountImages(ctx context.Context) (int, error)
	GetImage(ctx context.Context, sha string) (*storage.Image, bool, error)
	GetImagesBatch(ctx context.Context, shas []string) ([]*storage.Image, error)

	UpsertImage(ctx context.Context, image *storage.Image) error

	DeleteImages(ctx context.Context, ids ...string) error
	Exists(ctx context.Context, id string) (bool, error)
}

func newDatastore(storage store.Store, bleveIndex bleve.Index, noUpdateTimestamps bool, risks riskDS.DataStore) (DataStore, error) {
	indexer := index.New(bleveIndex)
	searcher := search.New(storage, indexer)

	ds, err := newDatastoreImpl(storage, indexer, searcher, risks)
	if err != nil {
		return nil, err
	}

	if err := ds.initializeRankers(); err != nil {
		return nil, errors.Wrap(err, "failed to initialize ranker")
	}

	return ds, nil
}

// NewBadger returns a new instance of DataStore using the input store, indexer, and searcher.
// noUpdateTimestamps controls whether timestamps are automatically updated when upserting images.
// This should be set to `false` except for some tests.
func NewBadger(db *badger.DB, bleveIndex bleve.Index, noUpdateTimestamps bool, risks riskDS.DataStore) (DataStore, error) {
	storage := badgerStore.New(db, noUpdateTimestamps)
	return newDatastore(storage, bleveIndex, noUpdateTimestamps, risks)
}
