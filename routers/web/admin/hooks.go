// Copyright 2018 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package admin

import (
	"net/http"

<<<<<<< HEAD
	"code.gitea.io/gitea/models/webhook"
	"code.gitea.io/gitea/modules/optional"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/templates"
	"code.gitea.io/gitea/services/context"
=======
	"code.proxgit.io/proxgit/models/webhook"
	"code.proxgit.io/proxgit/modules/optional"
	"code.proxgit.io/proxgit/modules/setting"
	"code.proxgit.io/proxgit/modules/templates"
	"code.proxgit.io/proxgit/services/context"
>>>>>>> master
)

const (
	// tplAdminHooks template path to render hook settings
	tplAdminHooks templates.TplName = "admin/hooks"
)

// DefaultOrSystemWebhooks renders both admin default and system webhook list pages
func DefaultOrSystemWebhooks(ctx *context.Context) {
	var err error

	ctx.Data["Title"] = ctx.Tr("admin.hooks")
	ctx.Data["PageIsAdminSystemHooks"] = true
	ctx.Data["PageIsAdminDefaultHooks"] = true

	def := make(map[string]any, len(ctx.Data))
	sys := make(map[string]any, len(ctx.Data))
	for k, v := range ctx.Data {
		def[k] = v
		sys[k] = v
	}

	sys["Title"] = ctx.Tr("admin.systemhooks")
<<<<<<< HEAD
	sys["Description"] = ctx.Tr("admin.systemhooks.desc", "https://docs.gitea.com/usage/webhooks")
=======
	sys["Description"] = ctx.Tr("admin.systemhooks.desc", "https://docs.proxgit.com/usage/webhooks")
>>>>>>> master
	sys["Webhooks"], err = webhook.GetSystemWebhooks(ctx, optional.None[bool]())
	sys["BaseLink"] = setting.AppSubURL + "/-/admin/hooks"
	sys["BaseLinkNew"] = setting.AppSubURL + "/-/admin/system-hooks"
	if err != nil {
		ctx.ServerError("GetWebhooksAdmin", err)
		return
	}

	def["Title"] = ctx.Tr("admin.defaulthooks")
<<<<<<< HEAD
	def["Description"] = ctx.Tr("admin.defaulthooks.desc", "https://docs.gitea.com/usage/webhooks")
=======
	def["Description"] = ctx.Tr("admin.defaulthooks.desc", "https://docs.proxgit.com/usage/webhooks")
>>>>>>> master
	def["Webhooks"], err = webhook.GetDefaultWebhooks(ctx)
	def["BaseLink"] = setting.AppSubURL + "/-/admin/hooks"
	def["BaseLinkNew"] = setting.AppSubURL + "/-/admin/default-hooks"
	if err != nil {
		ctx.ServerError("GetWebhooksAdmin", err)
		return
	}

	ctx.Data["DefaultWebhooks"] = def
	ctx.Data["SystemWebhooks"] = sys

	ctx.HTML(http.StatusOK, tplAdminHooks)
}

// DeleteDefaultOrSystemWebhook handler to delete an admin-defined system or default webhook
func DeleteDefaultOrSystemWebhook(ctx *context.Context) {
	if err := webhook.DeleteDefaultSystemWebhook(ctx, ctx.FormInt64("id")); err != nil {
		ctx.Flash.Error("DeleteDefaultWebhook: " + err.Error())
	} else {
		ctx.Flash.Success(ctx.Tr("repo.settings.webhook_deletion_success"))
	}

	ctx.JSONRedirect(setting.AppSubURL + "/-/admin/hooks")
}
