// Copyright 2019 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package repo

import (
	"net/http"
	"testing"

<<<<<<< HEAD
	repo_model "code.gitea.io/gitea/models/repo"
	"code.gitea.io/gitea/models/unittest"
	api "code.gitea.io/gitea/modules/structs"
	"code.gitea.io/gitea/modules/web"
	"code.gitea.io/gitea/services/contexttest"
=======
	repo_model "code.proxgit.io/proxgit/models/repo"
	"code.proxgit.io/proxgit/models/unittest"
	api "code.proxgit.io/proxgit/modules/structs"
	"code.proxgit.io/proxgit/modules/web"
	"code.proxgit.io/proxgit/services/contexttest"
>>>>>>> master

	"github.com/stretchr/testify/assert"
)

func TestRepoEdit(t *testing.T) {
	unittest.PrepareTestEnv(t)

	ctx, _ := contexttest.MockAPIContext(t, "user2/repo1")
	contexttest.LoadRepo(t, ctx, 1)
	contexttest.LoadUser(t, ctx, 2)
	ctx.Repo.Owner = ctx.Doer
	description := "new description"
	website := "http://wwww.newwebsite.com"
	private := true
	hasIssues := false
	hasWiki := false
	defaultBranch := "master"
	hasPullRequests := true
	ignoreWhitespaceConflicts := true
	allowMerge := false
	allowRebase := false
	allowRebaseMerge := false
	allowSquashMerge := false
	allowFastForwardOnlyMerge := false
	archived := true
	opts := api.EditRepoOption{
		Name:                      &ctx.Repo.Repository.Name,
		Description:               &description,
		Website:                   &website,
		Private:                   &private,
		HasIssues:                 &hasIssues,
		HasWiki:                   &hasWiki,
		DefaultBranch:             &defaultBranch,
		HasPullRequests:           &hasPullRequests,
		IgnoreWhitespaceConflicts: &ignoreWhitespaceConflicts,
		AllowMerge:                &allowMerge,
		AllowRebase:               &allowRebase,
		AllowRebaseMerge:          &allowRebaseMerge,
		AllowSquash:               &allowSquashMerge,
		AllowFastForwardOnly:      &allowFastForwardOnlyMerge,
		Archived:                  &archived,
	}

	web.SetForm(ctx, &opts)
	Edit(ctx)

	assert.Equal(t, http.StatusOK, ctx.Resp.WrittenStatus())
	unittest.AssertExistsAndLoadBean(t, &repo_model.Repository{
		ID: 1,
	}, unittest.Cond("name = ? AND is_archived = 1", *opts.Name))
}

func TestRepoEditNameChange(t *testing.T) {
	unittest.PrepareTestEnv(t)

	ctx, _ := contexttest.MockAPIContext(t, "user2/repo1")
	contexttest.LoadRepo(t, ctx, 1)
	contexttest.LoadUser(t, ctx, 2)
	ctx.Repo.Owner = ctx.Doer
	name := "newname"
	opts := api.EditRepoOption{
		Name: &name,
	}

	web.SetForm(ctx, &opts)
	Edit(ctx)
	assert.Equal(t, http.StatusOK, ctx.Resp.WrittenStatus())

	unittest.AssertExistsAndLoadBean(t, &repo_model.Repository{
		ID: 1,
	}, unittest.Cond("name = ?", opts.Name))
}
