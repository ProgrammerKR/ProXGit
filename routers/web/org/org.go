// Copyright 2014 The Gogs Authors. All rights reserved.
// Copyright 2018 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package org

import (
	"errors"
	"net/http"

	"code.proxgit.io/proxgit/models/db"
	"code.proxgit.io/proxgit/models/organization"
	user_model "code.proxgit.io/proxgit/models/user"
	"code.proxgit.io/proxgit/modules/log"
	"code.proxgit.io/proxgit/modules/setting"
	"code.proxgit.io/proxgit/modules/templates"
	"code.proxgit.io/proxgit/modules/web"
	"code.proxgit.io/proxgit/services/context"
	"code.proxgit.io/proxgit/services/forms"
)

const (
	// tplCreateOrg template path for create organization
	tplCreateOrg templates.TplName = "org/create"
)

// Create render the page for create organization
func Create(ctx *context.Context) {
	ctx.Data["Title"] = ctx.Tr("new_org")
	if !ctx.Doer.CanCreateOrganization() {
		ctx.ServerError("Not allowed", errors.New(ctx.Locale.TrString("org.form.create_org_not_allowed")))
		return
	}

	ctx.Data["visibility"] = setting.Service.DefaultOrgVisibilityMode
	ctx.Data["repo_admin_change_team_access"] = true

	ctx.HTML(http.StatusOK, tplCreateOrg)
}

// CreatePost response for create organization
func CreatePost(ctx *context.Context) {
	form := *web.GetForm(ctx).(*forms.CreateOrgForm)
	ctx.Data["Title"] = ctx.Tr("new_org")

	if !ctx.Doer.CanCreateOrganization() {
		ctx.ServerError("Not allowed", errors.New(ctx.Locale.TrString("org.form.create_org_not_allowed")))
		return
	}

	if ctx.HasError() {
		ctx.HTML(http.StatusOK, tplCreateOrg)
		return
	}

	org := &organization.Organization{
		Name:                      form.OrgName,
		IsActive:                  true,
		Type:                      user_model.UserTypeOrganization,
		Visibility:                form.Visibility,
		RepoAdminChangeTeamAccess: form.RepoAdminChangeTeamAccess,
	}

	if err := organization.CreateOrganization(ctx, org, ctx.Doer); err != nil {
		ctx.Data["Err_OrgName"] = true
		switch {
		case user_model.IsErrUserAlreadyExist(err):
			ctx.RenderWithErr(ctx.Tr("form.org_name_been_taken"), tplCreateOrg, &form)
		case db.IsErrNameReserved(err):
			ctx.RenderWithErr(ctx.Tr("org.form.name_reserved", err.(db.ErrNameReserved).Name), tplCreateOrg, &form)
		case db.IsErrNamePatternNotAllowed(err):
			ctx.RenderWithErr(ctx.Tr("org.form.name_pattern_not_allowed", err.(db.ErrNamePatternNotAllowed).Pattern), tplCreateOrg, &form)
		case organization.IsErrUserNotAllowedCreateOrg(err):
			ctx.RenderWithErr(ctx.Tr("org.form.create_org_not_allowed"), tplCreateOrg, &form)
		default:
			ctx.ServerError("CreateOrganization", err)
		}
		return
	}
	log.Trace("Organization created: %s", org.Name)

	ctx.Redirect(org.AsUser().DashboardLink())
}
