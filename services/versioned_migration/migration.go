// Copyright 2025 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package versioned_migration //nolint

import (
	"context"

<<<<<<< HEAD
	"code.gitea.io/gitea/models/migrations"
	"code.gitea.io/gitea/modules/globallock"
=======
	"code.proxgit.io/proxgit/models/migrations"
	"code.proxgit.io/proxgit/modules/globallock"
>>>>>>> master

	"xorm.io/xorm"
)

func Migrate(ctx context.Context, x *xorm.Engine) error {
	// only one instance can do the migration at the same time if there are multiple instances
<<<<<<< HEAD
	release, err := globallock.Lock(ctx, "gitea_versioned_migration")
=======
	release, err := globallock.Lock(ctx, "proxgit_versioned_migration")
>>>>>>> master
	if err != nil {
		return err
	}
	defer release()

	return migrations.Migrate(x)
}
