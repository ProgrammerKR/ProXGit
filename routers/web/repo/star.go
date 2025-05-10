// Copyright 2025 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package repo

import (
	"net/http"

	repo_model "code.proxgit.io/proxgit/models/repo"
	"code.proxgit.io/proxgit/modules/templates"
	"code.proxgit.io/proxgit/services/context"
)

const tplStarUnstar templates.TplName = "repo/star_unstar"

func ActionStar(ctx *context.Context) {
	err := repo_model.StarRepo(ctx, ctx.Doer, ctx.Repo.Repository, ctx.PathParam("action") == "star")
	if err != nil {
		handleActionError(ctx, err)
		return
	}

	ctx.Data["IsStaringRepo"] = repo_model.IsStaring(ctx, ctx.Doer.ID, ctx.Repo.Repository.ID)
	ctx.Data["Repository"], err = repo_model.GetRepositoryByName(ctx, ctx.Repo.Repository.OwnerID, ctx.Repo.Repository.Name)
	if err != nil {
		ctx.ServerError("GetRepositoryByName", err)
		return
	}
	ctx.RespHeader().Add("hx-trigger", "refreshUserCards") // see the `hx-trigger="refreshUserCards ..."` comments in tmpl
	ctx.HTML(http.StatusOK, tplStarUnstar)
}
