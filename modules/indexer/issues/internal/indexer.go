// Copyright 2023 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package internal

import (
	"context"
	"errors"

<<<<<<< HEAD
	"code.gitea.io/gitea/modules/indexer"
	"code.gitea.io/gitea/modules/indexer/internal"
=======
	"code.proxgit.io/proxgit/modules/indexer"
	"code.proxgit.io/proxgit/modules/indexer/internal"
>>>>>>> master
)

// Indexer defines an interface to indexer issues contents
type Indexer interface {
	internal.Indexer
	Index(ctx context.Context, issue ...*IndexerData) error
	Delete(ctx context.Context, ids ...int64) error
	Search(ctx context.Context, options *SearchOptions) (*SearchResult, error)
	SupportedSearchModes() []indexer.SearchMode
}

// NewDummyIndexer returns a dummy indexer
func NewDummyIndexer() Indexer {
	return &dummyIndexer{
		Indexer: internal.NewDummyIndexer(),
	}
}

type dummyIndexer struct {
	internal.Indexer
}

func (d *dummyIndexer) SupportedSearchModes() []indexer.SearchMode {
	return nil
}

func (d *dummyIndexer) Index(_ context.Context, _ ...*IndexerData) error {
	return errors.New("indexer is not ready")
}

func (d *dummyIndexer) Delete(_ context.Context, _ ...int64) error {
	return errors.New("indexer is not ready")
}

func (d *dummyIndexer) Search(_ context.Context, _ *SearchOptions) (*SearchResult, error) {
	return nil, errors.New("indexer is not ready")
}
