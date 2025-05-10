// Copyright 2021 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package system

import (
	"context"

<<<<<<< HEAD
	"code.gitea.io/gitea/models/system"
	"code.gitea.io/gitea/modules/json"
	"code.gitea.io/gitea/modules/util"
=======
	"code.proxgit.io/proxgit/models/system"
	"code.proxgit.io/proxgit/modules/json"
	"code.proxgit.io/proxgit/modules/util"
>>>>>>> master
)

// DBStore can be used to store app state items in local filesystem
type DBStore struct{}

// Get reads the state item
func (f *DBStore) Get(ctx context.Context, item StateItem) error {
	content, err := system.GetAppStateContent(ctx, item.Name())
	if err != nil {
		return err
	}
	if content == "" {
		return nil
	}
	return json.Unmarshal(util.UnsafeStringToBytes(content), item)
}

// Set saves the state item
func (f *DBStore) Set(ctx context.Context, item StateItem) error {
	b, err := json.Marshal(item)
	if err != nil {
		return err
	}
	return system.SaveAppStateContent(ctx, item.Name(), util.UnsafeBytesToString(b))
}
