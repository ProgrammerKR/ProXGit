// Copyright 2017 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package integration

import (
	"net/http"
	"testing"

	issues_model "code.proxgit.io/proxgit/models/issues"
	"code.proxgit.io/proxgit/models/organization"
	access_model "code.proxgit.io/proxgit/models/perm/access"
	repo_model "code.proxgit.io/proxgit/models/repo"
	"code.proxgit.io/proxgit/models/unittest"
	user_model "code.proxgit.io/proxgit/models/user"
	"code.proxgit.io/proxgit/tests"
)

func assertUserDeleted(t *testing.T, userID int64) {
	unittest.AssertNotExistsBean(t, &user_model.User{ID: userID})
	unittest.AssertNotExistsBean(t, &user_model.Follow{UserID: userID})
	unittest.AssertNotExistsBean(t, &user_model.Follow{FollowID: userID})
	unittest.AssertNotExistsBean(t, &repo_model.Repository{OwnerID: userID})
	unittest.AssertNotExistsBean(t, &access_model.Access{UserID: userID})
	unittest.AssertNotExistsBean(t, &organization.OrgUser{UID: userID})
	unittest.AssertNotExistsBean(t, &issues_model.IssueUser{UID: userID})
	unittest.AssertNotExistsBean(t, &organization.TeamUser{UID: userID})
	unittest.AssertNotExistsBean(t, &repo_model.Star{UID: userID})
}

func TestUserDeleteAccount(t *testing.T) {
	defer tests.PrepareTestEnv(t)()

	session := loginUser(t, "user8")
	csrf := GetUserCSRFToken(t, session)
	urlStr := "/user/settings/account/delete?password=" + userPassword
	req := NewRequestWithValues(t, "POST", urlStr, map[string]string{
		"_csrf": csrf,
	})
	session.MakeRequest(t, req, http.StatusSeeOther)

	assertUserDeleted(t, 8)
	unittest.CheckConsistencyFor(t, &user_model.User{})
}

func TestUserDeleteAccountStillOwnRepos(t *testing.T) {
	defer tests.PrepareTestEnv(t)()

	session := loginUser(t, "user2")
	csrf := GetUserCSRFToken(t, session)
	urlStr := "/user/settings/account/delete?password=" + userPassword
	req := NewRequestWithValues(t, "POST", urlStr, map[string]string{
		"_csrf": csrf,
	})
	session.MakeRequest(t, req, http.StatusSeeOther)

	// user should not have been deleted, because the user still owns repos
	unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 2})
}
