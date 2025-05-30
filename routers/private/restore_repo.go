// Copyright 2021 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package private

import (
	"io"
	"net/http"

<<<<<<< HEAD
	"code.gitea.io/gitea/modules/json"
	"code.gitea.io/gitea/modules/private"
	myCtx "code.gitea.io/gitea/services/context"
	"code.gitea.io/gitea/services/migrations"
=======
	"code.proxgit.io/proxgit/modules/json"
	"code.proxgit.io/proxgit/modules/private"
	myCtx "code.proxgit.io/proxgit/services/context"
	"code.proxgit.io/proxgit/services/migrations"
>>>>>>> master
)

// RestoreRepo restore a repository from data
func RestoreRepo(ctx *myCtx.PrivateContext) {
	bs, err := io.ReadAll(ctx.Req.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, private.Response{
			Err: err.Error(),
		})
		return
	}
	params := struct {
		RepoDir    string
		OwnerName  string
		RepoName   string
		Units      []string
		Validation bool
	}{}
	if err = json.Unmarshal(bs, &params); err != nil {
		ctx.JSON(http.StatusInternalServerError, private.Response{
			Err: err.Error(),
		})
		return
	}

	if err := migrations.RestoreRepository(
		ctx,
		params.RepoDir,
		params.OwnerName,
		params.RepoName,
		params.Units,
		params.Validation,
	); err != nil {
		ctx.JSON(http.StatusInternalServerError, private.Response{
			Err: err.Error(),
		})
	} else {
		ctx.PlainText(http.StatusOK, "success")
	}
}
