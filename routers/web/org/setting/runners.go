// Copyright 2022 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package setting

import (
<<<<<<< HEAD
	"code.gitea.io/gitea/services/context"
=======
	"code.proxgit.io/proxgit/services/context"
>>>>>>> master
)

func RedirectToDefaultSetting(ctx *context.Context) {
	ctx.Redirect(ctx.Org.OrgLink + "/settings/actions/runners")
}
