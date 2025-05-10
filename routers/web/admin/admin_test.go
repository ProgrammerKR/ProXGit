// Copyright 2019 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package admin

import (
	"net/http"
	"testing"

<<<<<<< HEAD
	"code.gitea.io/gitea/modules/json"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/test"
	"code.gitea.io/gitea/services/contexttest"
=======
	"code.proxgit.io/proxgit/modules/json"
	"code.proxgit.io/proxgit/modules/setting"
	"code.proxgit.io/proxgit/modules/test"
	"code.proxgit.io/proxgit/services/contexttest"
>>>>>>> master

	"github.com/stretchr/testify/assert"
)

func TestShadowPassword(t *testing.T) {
	kases := []struct {
		Provider string
		CfgItem  string
		Result   string
	}{
		{
			Provider: "redis",
<<<<<<< HEAD
			CfgItem:  "network=tcp,addr=:6379,password=gitea,db=0,pool_size=100,idle_timeout=180",
=======
			CfgItem:  "network=tcp,addr=:6379,password=proxgit,db=0,pool_size=100,idle_timeout=180",
>>>>>>> master
			Result:   "network=tcp,addr=:6379,password=******,db=0,pool_size=100,idle_timeout=180",
		},
		{
			Provider: "mysql",
<<<<<<< HEAD
			CfgItem:  "root:@tcp(localhost:3306)/gitea?charset=utf8",
			Result:   "root:******@tcp(localhost:3306)/gitea?charset=utf8",
		},
		{
			Provider: "mysql",
			CfgItem:  "/gitea?charset=utf8",
			Result:   "/gitea?charset=utf8",
=======
			CfgItem:  "root:@tcp(localhost:3306)/proxgit?charset=utf8",
			Result:   "root:******@tcp(localhost:3306)/proxgit?charset=utf8",
		},
		{
			Provider: "mysql",
			CfgItem:  "/proxgit?charset=utf8",
			Result:   "/proxgit?charset=utf8",
>>>>>>> master
		},
		{
			Provider: "mysql",
			CfgItem:  "user:mypassword@/dbname",
			Result:   "user:******@/dbname",
		},
		{
			Provider: "postgres",
			CfgItem:  "user=pqgotest dbname=pqgotest sslmode=verify-full",
			Result:   "user=pqgotest dbname=pqgotest sslmode=verify-full",
		},
		{
			Provider: "postgres",
			CfgItem:  "user=pqgotest password= dbname=pqgotest sslmode=verify-full",
			Result:   "user=pqgotest password=****** dbname=pqgotest sslmode=verify-full",
		},
		{
			Provider: "postgres",
			CfgItem:  "postgres://user:pass@hostname/dbname",
			Result:   "postgres://user:******@hostname/dbname",
		},
		{
			Provider: "couchbase",
			CfgItem:  "http://dev-couchbase.example.com:8091/",
			Result:   "http://dev-couchbase.example.com:8091/",
		},
		{
			Provider: "couchbase",
			CfgItem:  "http://user:the_password@dev-couchbase.example.com:8091/",
			Result:   "http://user:******@dev-couchbase.example.com:8091/",
		},
	}

	for _, k := range kases {
		assert.Equal(t, k.Result, shadowPassword(k.Provider, k.CfgItem))
	}
}

func TestSelfCheckPost(t *testing.T) {
	defer test.MockVariableValue(&setting.AppURL, "http://config/sub/")()
	defer test.MockVariableValue(&setting.AppSubURL, "/sub")()

	ctx, resp := contexttest.MockContext(t, "GET http://host/sub/admin/self_check?location_origin=http://frontend")
	SelfCheckPost(ctx)
	assert.Equal(t, http.StatusOK, resp.Code)

	data := struct {
		Problems []string `json:"problems"`
	}{}
	err := json.Unmarshal(resp.Body.Bytes(), &data)
	assert.NoError(t, err)
	assert.Equal(t, []string{
		ctx.Locale.TrString("admin.self_check.location_origin_mismatch", "http://frontend/sub/", "http://config/sub/"),
	}, data.Problems)
}
