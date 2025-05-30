// Copyright 2023 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package v1_22 //nolint

import (
<<<<<<< HEAD
	"code.gitea.io/gitea/modules/timeutil"
=======
	"code.proxgit.io/proxgit/modules/timeutil"
>>>>>>> master

	"xorm.io/xorm"
)

func CreateAuthTokenTable(x *xorm.Engine) error {
	type AuthToken struct {
		ID          string `xorm:"pk"`
		TokenHash   string
		UserID      int64              `xorm:"INDEX"`
		ExpiresUnix timeutil.TimeStamp `xorm:"INDEX"`
	}

	return x.Sync(new(AuthToken))
}
