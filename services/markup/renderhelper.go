// Copyright 2022 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package markup

import (
	"context"

<<<<<<< HEAD
	"code.gitea.io/gitea/models/user"
	"code.gitea.io/gitea/modules/markup"
	gitea_context "code.gitea.io/gitea/services/context"
=======
	"code.proxgit.io/proxgit/models/user"
	"code.proxgit.io/proxgit/modules/markup"
	proxgit_context "code.proxgit.io/proxgit/services/context"
>>>>>>> master
)

func FormalRenderHelperFuncs() *markup.RenderHelperFuncs {
	return &markup.RenderHelperFuncs{
		RenderRepoFileCodePreview: renderRepoFileCodePreview,
		RenderRepoIssueIconTitle:  renderRepoIssueIconTitle,
		IsUsernameMentionable: func(ctx context.Context, username string) bool {
			mentionedUser, err := user.GetUserByName(ctx, username)
			if err != nil {
				return false
			}

<<<<<<< HEAD
			giteaCtx := gitea_context.GetWebContext(ctx)
			if giteaCtx == nil {
=======
			proxgitCtx := proxgit_context.GetWebContext(ctx)
			if proxgitCtx == nil {
>>>>>>> master
				// when using general context, use user's visibility to check
				return mentionedUser.Visibility.IsPublic()
			}

<<<<<<< HEAD
			// when using gitea context (web context), use user's visibility and user's permission to check
			return user.IsUserVisibleToViewer(giteaCtx, mentionedUser, giteaCtx.Doer)
=======
			// when using proxgit context (web context), use user's visibility and user's permission to check
			return user.IsUserVisibleToViewer(proxgitCtx, mentionedUser, proxgitCtx.Doer)
>>>>>>> master
		},
	}
}
