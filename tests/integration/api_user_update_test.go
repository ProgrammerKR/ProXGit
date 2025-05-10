// Copyright 2024 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package integration

import (
	"net/http"
	"testing"

	auth_model "code.proxgit.io/proxgit/models/auth"
	"code.proxgit.io/proxgit/tests"
)

func TestAPIUpdateUser(t *testing.T) {
	defer tests.PrepareTestEnv(t)()

	normalUsername := "user2"
	session := loginUser(t, normalUsername)
	token := getTokenForLoggedInUser(t, session, auth_model.AccessTokenScopeWriteUser)

	req := NewRequestWithJSON(t, "PATCH", "/api/v1/user/settings", map[string]string{
		"website": "https://proxgit.com",
	}).AddTokenAuth(token)
	MakeRequest(t, req, http.StatusOK)
}
