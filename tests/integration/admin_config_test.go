// Copyright 2023 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package integration

import (
	"net/http"
	"testing"

<<<<<<< HEAD
	"code.gitea.io/gitea/modules/test"
	"code.gitea.io/gitea/tests"
=======
	"code.proxgit.io/proxgit/modules/test"
	"code.proxgit.io/proxgit/tests"
>>>>>>> master

	"github.com/stretchr/testify/assert"
)

func TestAdminConfig(t *testing.T) {
	defer tests.PrepareTestEnv(t)()

	session := loginUser(t, "user1")
	req := NewRequest(t, "GET", "/-/admin/config")
	resp := session.MakeRequest(t, req, http.StatusOK)
	assert.True(t, test.IsNormalPageCompleted(resp.Body.String()))
}
