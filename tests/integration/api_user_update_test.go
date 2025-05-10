// Copyright 2024 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package integration

import (
	"net/http"
	"testing"

<<<<<<< HEAD
	auth_model "code.gitea.io/gitea/models/auth"
	"code.gitea.io/gitea/tests"
=======
	auth_model "code.proxgit.io/proxgit/models/auth"
	"code.proxgit.io/proxgit/tests"
>>>>>>> master
)

func TestAPIUpdateUser(t *testing.T) {
	defer tests.PrepareTestEnv(t)()

	normalUsername := "user2"
	session := loginUser(t, normalUsername)
	token := getTokenForLoggedInUser(t, session, auth_model.AccessTokenScopeWriteUser)

	req := NewRequestWithJSON(t, "PATCH", "/api/v1/user/settings", map[string]string{
<<<<<<< HEAD
		"website": "https://gitea.com",
=======
		"website": "https://proxgit.com",
>>>>>>> master
	}).AddTokenAuth(token)
	MakeRequest(t, req, http.StatusOK)
}
