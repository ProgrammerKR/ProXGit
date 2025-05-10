// Copyright 2025 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package v1_24 //nolint

import (
<<<<<<< HEAD
	"code.gitea.io/gitea/models/perm"
=======
	"code.proxgit.io/proxgit/models/perm"
>>>>>>> master

	"xorm.io/xorm"
)

func AddRepoUnitAnonymousAccessMode(x *xorm.Engine) error {
	type RepoUnit struct { //revive:disable-line:exported
		AnonymousAccessMode perm.AccessMode `xorm:"NOT NULL DEFAULT 0"`
	}
	_, err := x.SyncWithOptions(xorm.SyncOptions{
		IgnoreConstrains: true,
		IgnoreIndices:    true,
	}, new(RepoUnit))
	return err
}
