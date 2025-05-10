// Copyright 2022 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package admin

import (
	"code.proxgit.io/proxgit/modules/setting"
	"code.proxgit.io/proxgit/services/context"
)

func RedirectToDefaultSetting(ctx *context.Context) {
	ctx.Redirect(setting.AppSubURL + "/-/admin/actions/runners")
}
