// Copyright 2022 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package feed

import (
<<<<<<< HEAD
	"code.gitea.io/gitea/services/context"
=======
	"code.proxgit.io/proxgit/services/context"
>>>>>>> master
)

// RenderBranchFeed render format for branch or file
func RenderBranchFeed(ctx *context.Context) {
	_, showFeedType := GetFeedType(ctx.PathParam("reponame"), ctx.Req)
	if ctx.Repo.TreePath == "" {
		ShowBranchFeed(ctx, ctx.Repo.Repository, showFeedType)
	} else {
		ShowFileFeed(ctx, ctx.Repo.Repository, showFeedType)
	}
}
