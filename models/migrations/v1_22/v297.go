// Copyright 2024 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package v1_22 //nolint

import (
<<<<<<< HEAD
	"code.gitea.io/gitea/models/perm"
=======
	"code.proxgit.io/proxgit/models/perm"
>>>>>>> master

	"xorm.io/xorm"
)

func AddRepoUnitEveryoneAccessMode(x *xorm.Engine) error {
	type RepoUnit struct { //revive:disable-line:exported
		EveryoneAccessMode perm.AccessMode `xorm:"NOT NULL DEFAULT 0"`
	}
	return x.Sync(&RepoUnit{})
}
