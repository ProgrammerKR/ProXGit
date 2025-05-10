// Copyright 2017 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package utils

import (
<<<<<<< HEAD
	"code.gitea.io/gitea/models/db"
	"code.gitea.io/gitea/services/context"
	"code.gitea.io/gitea/services/convert"
=======
	"code.proxgit.io/proxgit/models/db"
	"code.proxgit.io/proxgit/services/context"
	"code.proxgit.io/proxgit/services/convert"
>>>>>>> master
)

// GetListOptions returns list options using the page and limit parameters
func GetListOptions(ctx *context.APIContext) db.ListOptions {
	return db.ListOptions{
		Page:     ctx.FormInt("page"),
		PageSize: convert.ToCorrectPageSize(ctx.FormInt("limit")),
	}
}
