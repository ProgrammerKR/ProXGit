// Copyright 2021 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package repo

import (
	"net/http"
	"time"

	git_model "code.proxgit.io/proxgit/models/git"
	repo_model "code.proxgit.io/proxgit/models/repo"
	"code.proxgit.io/proxgit/modules/git"
	api "code.proxgit.io/proxgit/modules/structs"
	"code.proxgit.io/proxgit/modules/web"
	"code.proxgit.io/proxgit/services/context"
	pull_service "code.proxgit.io/proxgit/services/pull"
	"code.proxgit.io/proxgit/services/repository/files"
)

// ApplyDiffPatch handles API call for applying a patch
func ApplyDiffPatch(ctx *context.APIContext) {
	// swagger:operation POST /repos/{owner}/{repo}/diffpatch repository repoApplyDiffPatch
	// ---
	// summary: Apply diff patch to repository
	// consumes:
	// - application/json
	// produces:
	// - application/json
	// parameters:
	// - name: owner
	//   in: path
	//   description: owner of the repo
	//   type: string
	//   required: true
	// - name: repo
	//   in: path
	//   description: name of the repo
	//   type: string
	//   required: true
	// - name: body
	//   in: body
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/UpdateFileOptions"
	// responses:
	//   "200":
	//     "$ref": "#/responses/FileResponse"
	//   "404":
	//     "$ref": "#/responses/notFound"
	//   "423":
	//     "$ref": "#/responses/repoArchivedError"
	apiOpts := web.GetForm(ctx).(*api.ApplyDiffPatchFileOptions)

	opts := &files.ApplyDiffPatchOptions{
		Content:   apiOpts.Content,
		SHA:       apiOpts.SHA,
		Message:   apiOpts.Message,
		OldBranch: apiOpts.BranchName,
		NewBranch: apiOpts.NewBranchName,
		Committer: &files.IdentityOptions{
			GitUserName:  apiOpts.Committer.Name,
			GitUserEmail: apiOpts.Committer.Email,
		},
		Author: &files.IdentityOptions{
			GitUserName:  apiOpts.Author.Name,
			GitUserEmail: apiOpts.Author.Email,
		},
		Dates: &files.CommitDateOptions{
			Author:    apiOpts.Dates.Author,
			Committer: apiOpts.Dates.Committer,
		},
		Signoff: apiOpts.Signoff,
	}
	if opts.Dates.Author.IsZero() {
		opts.Dates.Author = time.Now()
	}
	if opts.Dates.Committer.IsZero() {
		opts.Dates.Committer = time.Now()
	}

	if opts.Message == "" {
		opts.Message = "apply-patch"
	}

	if !canWriteFiles(ctx, apiOpts.BranchName) {
		ctx.APIErrorInternal(repo_model.ErrUserDoesNotHaveAccessToRepo{
			UserID:   ctx.Doer.ID,
			RepoName: ctx.Repo.Repository.LowerName,
		})
		return
	}

	fileResponse, err := files.ApplyDiffPatch(ctx, ctx.Repo.Repository, ctx.Doer, opts)
	if err != nil {
		if files.IsErrUserCannotCommit(err) || pull_service.IsErrFilePathProtected(err) {
			ctx.APIError(http.StatusForbidden, err)
			return
		}
		if git_model.IsErrBranchAlreadyExists(err) || files.IsErrFilenameInvalid(err) || pull_service.IsErrSHADoesNotMatch(err) ||
			files.IsErrFilePathInvalid(err) || files.IsErrRepoFileAlreadyExists(err) {
			ctx.APIError(http.StatusUnprocessableEntity, err)
			return
		}
		if git_model.IsErrBranchNotExist(err) || git.IsErrBranchNotExist(err) {
			ctx.APIError(http.StatusNotFound, err)
			return
		}
		ctx.APIErrorInternal(err)
	} else {
		ctx.JSON(http.StatusCreated, fileResponse)
	}
}
