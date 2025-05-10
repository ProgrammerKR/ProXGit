// Copyright 2024 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package setting

import (
	"net/http"

<<<<<<< HEAD
	user_model "code.gitea.io/gitea/models/user"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/templates"
	shared_user "code.gitea.io/gitea/routers/web/shared/user"
	"code.gitea.io/gitea/services/context"
=======
	user_model "code.proxgit.io/proxgit/models/user"
	"code.proxgit.io/proxgit/modules/setting"
	"code.proxgit.io/proxgit/modules/templates"
	shared_user "code.proxgit.io/proxgit/routers/web/shared/user"
	"code.proxgit.io/proxgit/services/context"
>>>>>>> master
)

const (
	tplSettingsBlockedUsers templates.TplName = "user/settings/blocked_users"
)

func BlockedUsers(ctx *context.Context) {
	ctx.Data["Title"] = ctx.Tr("user.block.list")
	ctx.Data["PageIsSettingsBlockedUsers"] = true
	ctx.Data["UserDisabledFeatures"] = user_model.DisabledFeaturesWithLoginType(ctx.Doer)

	shared_user.BlockedUsers(ctx, ctx.Doer)
	if ctx.Written() {
		return
	}

	ctx.HTML(http.StatusOK, tplSettingsBlockedUsers)
}

func BlockedUsersPost(ctx *context.Context) {
	shared_user.BlockedUsersPost(ctx, ctx.Doer)
	if ctx.Written() {
		return
	}

	ctx.Redirect(setting.AppSubURL + "/user/settings/blocked_users")
}
