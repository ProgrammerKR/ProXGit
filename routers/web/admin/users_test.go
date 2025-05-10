// Copyright 2017 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package admin

import (
	"testing"

<<<<<<< HEAD
	"code.gitea.io/gitea/models/unittest"
	user_model "code.gitea.io/gitea/models/user"
	"code.gitea.io/gitea/modules/setting"
	api "code.gitea.io/gitea/modules/structs"
	"code.gitea.io/gitea/modules/web"
	"code.gitea.io/gitea/services/contexttest"
	"code.gitea.io/gitea/services/forms"
=======
	"code.proxgit.io/proxgit/models/unittest"
	user_model "code.proxgit.io/proxgit/models/user"
	"code.proxgit.io/proxgit/modules/setting"
	api "code.proxgit.io/proxgit/modules/structs"
	"code.proxgit.io/proxgit/modules/web"
	"code.proxgit.io/proxgit/services/contexttest"
	"code.proxgit.io/proxgit/services/forms"
>>>>>>> master

	"github.com/stretchr/testify/assert"
)

func TestNewUserPost_MustChangePassword(t *testing.T) {
	unittest.PrepareTestEnv(t)
	ctx, _ := contexttest.MockContext(t, "admin/users/new")

	u := unittest.AssertExistsAndLoadBean(t, &user_model.User{
		IsAdmin: true,
		ID:      2,
	})

	ctx.Doer = u

<<<<<<< HEAD
	username := "gitea"
	email := "gitea@gitea.io"
=======
	username := "proxgit"
	email := "proxgit@proxgit.io"
>>>>>>> master

	form := forms.AdminCreateUserForm{
		LoginType:          "local",
		LoginName:          "local",
		UserName:           username,
		Email:              email,
		Password:           "abc123ABC!=$",
		SendNotify:         false,
		MustChangePassword: true,
	}

	web.SetForm(ctx, &form)
	NewUserPost(ctx)

	assert.NotEmpty(t, ctx.Flash.SuccessMsg)

	u, err := user_model.GetUserByName(ctx, username)

	assert.NoError(t, err)
	assert.Equal(t, username, u.Name)
	assert.Equal(t, email, u.Email)
	assert.True(t, u.MustChangePassword)
}

func TestNewUserPost_MustChangePasswordFalse(t *testing.T) {
	unittest.PrepareTestEnv(t)
	ctx, _ := contexttest.MockContext(t, "admin/users/new")

	u := unittest.AssertExistsAndLoadBean(t, &user_model.User{
		IsAdmin: true,
		ID:      2,
	})

	ctx.Doer = u

<<<<<<< HEAD
	username := "gitea"
	email := "gitea@gitea.io"
=======
	username := "proxgit"
	email := "proxgit@proxgit.io"
>>>>>>> master

	form := forms.AdminCreateUserForm{
		LoginType:          "local",
		LoginName:          "local",
		UserName:           username,
		Email:              email,
		Password:           "abc123ABC!=$",
		SendNotify:         false,
		MustChangePassword: false,
	}

	web.SetForm(ctx, &form)
	NewUserPost(ctx)

	assert.NotEmpty(t, ctx.Flash.SuccessMsg)

	u, err := user_model.GetUserByName(ctx, username)

	assert.NoError(t, err)
	assert.Equal(t, username, u.Name)
	assert.Equal(t, email, u.Email)
	assert.False(t, u.MustChangePassword)
}

func TestNewUserPost_InvalidEmail(t *testing.T) {
	unittest.PrepareTestEnv(t)
	ctx, _ := contexttest.MockContext(t, "admin/users/new")

	u := unittest.AssertExistsAndLoadBean(t, &user_model.User{
		IsAdmin: true,
		ID:      2,
	})

	ctx.Doer = u

<<<<<<< HEAD
	username := "gitea"
	email := "gitea@gitea.io\r\n"
=======
	username := "proxgit"
	email := "proxgit@proxgit.io\r\n"
>>>>>>> master

	form := forms.AdminCreateUserForm{
		LoginType:          "local",
		LoginName:          "local",
		UserName:           username,
		Email:              email,
		Password:           "abc123ABC!=$",
		SendNotify:         false,
		MustChangePassword: false,
	}

	web.SetForm(ctx, &form)
	NewUserPost(ctx)

	assert.NotEmpty(t, ctx.Flash.ErrorMsg)
}

func TestNewUserPost_VisibilityDefaultPublic(t *testing.T) {
	unittest.PrepareTestEnv(t)
	ctx, _ := contexttest.MockContext(t, "admin/users/new")

	u := unittest.AssertExistsAndLoadBean(t, &user_model.User{
		IsAdmin: true,
		ID:      2,
	})

	ctx.Doer = u

<<<<<<< HEAD
	username := "gitea"
	email := "gitea@gitea.io"
=======
	username := "proxgit"
	email := "proxgit@proxgit.io"
>>>>>>> master

	form := forms.AdminCreateUserForm{
		LoginType:          "local",
		LoginName:          "local",
		UserName:           username,
		Email:              email,
		Password:           "abc123ABC!=$",
		SendNotify:         false,
		MustChangePassword: false,
	}

	web.SetForm(ctx, &form)
	NewUserPost(ctx)

	assert.NotEmpty(t, ctx.Flash.SuccessMsg)

	u, err := user_model.GetUserByName(ctx, username)

	assert.NoError(t, err)
	assert.Equal(t, username, u.Name)
	assert.Equal(t, email, u.Email)
	// As default user visibility
	assert.Equal(t, setting.Service.DefaultUserVisibilityMode, u.Visibility)
}

func TestNewUserPost_VisibilityPrivate(t *testing.T) {
	unittest.PrepareTestEnv(t)
	ctx, _ := contexttest.MockContext(t, "admin/users/new")

	u := unittest.AssertExistsAndLoadBean(t, &user_model.User{
		IsAdmin: true,
		ID:      2,
	})

	ctx.Doer = u

<<<<<<< HEAD
	username := "gitea"
	email := "gitea@gitea.io"
=======
	username := "proxgit"
	email := "proxgit@proxgit.io"
>>>>>>> master

	form := forms.AdminCreateUserForm{
		LoginType:          "local",
		LoginName:          "local",
		UserName:           username,
		Email:              email,
		Password:           "abc123ABC!=$",
		SendNotify:         false,
		MustChangePassword: false,
		Visibility:         api.VisibleTypePrivate,
	}

	web.SetForm(ctx, &form)
	NewUserPost(ctx)

	assert.NotEmpty(t, ctx.Flash.SuccessMsg)

	u, err := user_model.GetUserByName(ctx, username)

	assert.NoError(t, err)
	assert.Equal(t, username, u.Name)
	assert.Equal(t, email, u.Email)
	// As default user visibility
	assert.True(t, u.Visibility.IsPrivate())
}
