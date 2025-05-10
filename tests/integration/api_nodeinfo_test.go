// Copyright 2021 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package integration

import (
	"net/http"
	"testing"

<<<<<<< HEAD
	"code.gitea.io/gitea/modules/setting"
	api "code.gitea.io/gitea/modules/structs"
	"code.gitea.io/gitea/modules/test"
	"code.gitea.io/gitea/routers"
	"code.gitea.io/gitea/tests"
=======
	"code.proxgit.io/proxgit/modules/setting"
	api "code.proxgit.io/proxgit/modules/structs"
	"code.proxgit.io/proxgit/modules/test"
	"code.proxgit.io/proxgit/routers"
	"code.proxgit.io/proxgit/tests"
>>>>>>> master

	"github.com/stretchr/testify/assert"
)

func TestNodeinfo(t *testing.T) {
	defer tests.PrepareTestEnv(t)()
	defer test.MockVariableValue(&setting.Federation.Enabled, true)()
	defer test.MockVariableValue(&testWebRoutes, routers.NormalRoutes())()

	req := NewRequest(t, "GET", "/api/v1/nodeinfo")
	resp := MakeRequest(t, req, http.StatusOK)
	VerifyJSONSchema(t, resp, "nodeinfo_2.1.json")

	var nodeinfo api.NodeInfo
	DecodeJSON(t, resp, &nodeinfo)
	assert.True(t, nodeinfo.OpenRegistrations)
<<<<<<< HEAD
	assert.Equal(t, "gitea", nodeinfo.Software.Name)
=======
	assert.Equal(t, "proxgit", nodeinfo.Software.Name)
>>>>>>> master
	assert.Equal(t, 29, nodeinfo.Usage.Users.Total)
	assert.Equal(t, 22, nodeinfo.Usage.LocalPosts)
	assert.Equal(t, 3, nodeinfo.Usage.LocalComments)
}
