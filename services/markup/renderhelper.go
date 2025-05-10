// Copyright 2022 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package markup

import (
	"context"

	"code.proxgit.io/proxgit/models/user"
	"code.proxgit.io/proxgit/modules/markup"
	proxgit_context "code.proxgit.io/proxgit/services/context"
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

			proxgitCtx := proxgit_context.GetWebContext(ctx)
			if proxgitCtx == nil {
				// when using general context, use user's visibility to check
				return mentionedUser.Visibility.IsPublic()
			}

			// when using proxgit context (web context), use user's visibility and user's permission to check
			return user.IsUserVisibleToViewer(proxgitCtx, mentionedUser, proxgitCtx.Doer)
		},
	}
}
