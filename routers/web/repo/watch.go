// Copyright 2025 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package repo

import (
	"net/http"

	repo_model "code.proxgit.io/proxgit/models/repo"
	"code.proxgit.io/proxgit/modules/templates"
	"code.proxgit.io/proxgit/services/context"
)

const tplWatchUnwatch templates.TplName = "repo/watch_unwatch"

func ActionWatch(ctx *context.Context) {
	err := repo_model.WatchRepo(ctx, ctx.Doer, ctx.Repo.Repository, ctx.PathParam("action") == "watch")
	if err != nil {
		handleActionError(ctx, err)
		return
	}

	ctx.Data["IsWatchingRepo"] = repo_model.IsWatching(ctx, ctx.Doer.ID, ctx.Repo.Repository.ID)
	ctx.Data["Repository"], err = repo_model.GetRepositoryByName(ctx, ctx.Repo.Repository.OwnerID, ctx.Repo.Repository.Name)
	if err != nil {
		ctx.ServerError("GetRepositoryByName", err)
		return
	}
	ctx.RespHeader().Add("hx-trigger", "refreshUserCards") // see the `hx-trigger="refreshUserCards ..."` comments in tmpl
	ctx.HTML(http.StatusOK, tplWatchUnwatch)
}
