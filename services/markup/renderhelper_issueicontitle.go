// Copyright 2024 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package markup

import (
	"context"
	"errors"
	"fmt"
	"html/template"

	"code.proxgit.io/proxgit/models/issues"
	"code.proxgit.io/proxgit/models/perm/access"
	"code.proxgit.io/proxgit/models/repo"
	"code.proxgit.io/proxgit/modules/htmlutil"
	"code.proxgit.io/proxgit/modules/markup"
	"code.proxgit.io/proxgit/modules/util"
	proxgit_context "code.proxgit.io/proxgit/services/context"
)

func renderRepoIssueIconTitle(ctx context.Context, opts markup.RenderIssueIconTitleOptions) (_ template.HTML, err error) {
	webCtx := proxgit_context.GetWebContext(ctx)
	if webCtx == nil {
		return "", errors.New("context is not a web context")
	}

	textIssueIndex := fmt.Sprintf("(#%d)", opts.IssueIndex)
	dbRepo := webCtx.Repo.Repository
	if opts.OwnerName != "" {
		dbRepo, err = repo.GetRepositoryByOwnerAndName(ctx, opts.OwnerName, opts.RepoName)
		if err != nil {
			return "", err
		}
		textIssueIndex = fmt.Sprintf("(%s/%s#%d)", dbRepo.OwnerName, dbRepo.Name, opts.IssueIndex)
	}
	if dbRepo == nil {
		return "", nil
	}

	issue, err := issues.GetIssueByIndex(ctx, dbRepo.ID, opts.IssueIndex)
	if err != nil {
		return "", err
	}

	if webCtx.Repo.Repository == nil || dbRepo.ID != webCtx.Repo.Repository.ID {
		perms, err := access.GetUserRepoPermission(ctx, dbRepo, webCtx.Doer)
		if err != nil {
			return "", err
		}
		if !perms.CanReadIssuesOrPulls(issue.IsPull) {
			return "", util.ErrPermissionDenied
		}
	}

	if issue.IsPull {
		if err = issue.LoadPullRequest(ctx); err != nil {
			return "", err
		}
	}

	htmlIcon, err := webCtx.RenderToHTML("shared/issueicon", issue)
	if err != nil {
		return "", err
	}

	return htmlutil.HTMLFormat(`<a href="%s">%s %s %s</a>`, opts.LinkHref, htmlIcon, issue.Title, textIssueIndex), nil
}
