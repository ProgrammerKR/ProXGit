// Copyright 2022 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package user

import (
	"net/http"

	"code.proxgit.io/proxgit/models/db"
	issues_model "code.proxgit.io/proxgit/models/issues"
	"code.proxgit.io/proxgit/services/context"
	"code.proxgit.io/proxgit/services/convert"
)

// GetStopwatches get all stopwatches
func GetStopwatches(ctx *context.Context) {
	sws, err := issues_model.GetUserStopwatches(ctx, ctx.Doer.ID, db.ListOptions{
		Page:     ctx.FormInt("page"),
		PageSize: convert.ToCorrectPageSize(ctx.FormInt("limit")),
	})
	if err != nil {
		ctx.HTTPError(http.StatusInternalServerError, err.Error())
		return
	}

	count, err := issues_model.CountUserStopwatches(ctx, ctx.Doer.ID)
	if err != nil {
		ctx.HTTPError(http.StatusInternalServerError, err.Error())
		return
	}

	apiSWs, err := convert.ToStopWatches(ctx, sws)
	if err != nil {
		ctx.HTTPError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.SetTotalCountHeader(count)
	ctx.JSON(http.StatusOK, apiSWs)
}
