// Copyright 2023 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package misc

import (
	"net/http"

	"code.proxgit.io/proxgit/modules/options"
	repo_module "code.proxgit.io/proxgit/modules/repository"
	"code.proxgit.io/proxgit/modules/structs"
	"code.proxgit.io/proxgit/modules/util"
	"code.proxgit.io/proxgit/services/context"
)

// Shows a list of all Gitignore templates
func ListGitignoresTemplates(ctx *context.APIContext) {
	// swagger:operation GET /gitignore/templates miscellaneous listGitignoresTemplates
	// ---
	// summary: Returns a list of all gitignore templates
	// produces:
	// - application/json
	// responses:
	//   "200":
	//     "$ref": "#/responses/GitignoreTemplateList"
	ctx.JSON(http.StatusOK, repo_module.Gitignores)
}

// SHows information about a gitignore template
func GetGitignoreTemplateInfo(ctx *context.APIContext) {
	// swagger:operation GET /gitignore/templates/{name} miscellaneous getGitignoreTemplateInfo
	// ---
	// summary: Returns information about a gitignore template
	// produces:
	// - application/json
	// parameters:
	// - name: name
	//   in: path
	//   description: name of the template
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/GitignoreTemplateInfo"
	//   "404":
	//     "$ref": "#/responses/notFound"
	name := util.PathJoinRelX(ctx.PathParam("name"))

	text, err := options.Gitignore(name)
	if err != nil {
		ctx.APIErrorNotFound()
		return
	}

	ctx.JSON(http.StatusOK, &structs.GitignoreTemplateInfo{Name: name, Source: string(text)})
}
