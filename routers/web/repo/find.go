// Copyright 2022 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package repo

import (
	"net/http"

	"code.proxgit.io/proxgit/modules/templates"
	"code.proxgit.io/proxgit/modules/util"
	"code.proxgit.io/proxgit/services/context"
)

const (
	tplFindFiles templates.TplName = "repo/find/files"
)

// FindFiles render the page to find repository files
func FindFiles(ctx *context.Context) {
	path := ctx.PathParam("*")
	ctx.Data["TreeLink"] = ctx.Repo.RepoLink + "/src/" + util.PathEscapeSegments(path)
	ctx.Data["DataLink"] = ctx.Repo.RepoLink + "/tree-list/" + util.PathEscapeSegments(path)
	ctx.HTML(http.StatusOK, tplFindFiles)
}
