// Copyright 2021 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package v1_16 //nolint

import (
<<<<<<< HEAD
	"code.gitea.io/gitea/models/migrations/base"
=======
	"code.proxgit.io/proxgit/models/migrations/base"
>>>>>>> master

	"xorm.io/xorm"
)

func RecreateIssueResourceIndexTable(x *xorm.Engine) error {
	type IssueIndex struct {
		GroupID  int64 `xorm:"pk"`
		MaxIndex int64 `xorm:"index"`
	}

	return base.RecreateTables(new(IssueIndex))(x)
}
