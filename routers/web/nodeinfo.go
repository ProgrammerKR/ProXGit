// Copyright 2021 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package web

import (
	"net/http"

<<<<<<< HEAD
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/services/context"
=======
	"code.proxgit.io/proxgit/modules/setting"
	"code.proxgit.io/proxgit/services/context"
>>>>>>> master
)

type nodeInfoLinks struct {
	Links []nodeInfoLink `json:"links"`
}

type nodeInfoLink struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
}

// NodeInfoLinks returns links to the node info endpoint
func NodeInfoLinks(ctx *context.Context) {
	nodeinfolinks := &nodeInfoLinks{
		Links: []nodeInfoLink{{
			setting.AppURL + "api/v1/nodeinfo",
			"http://nodeinfo.diaspora.software/ns/schema/2.1",
		}},
	}
	ctx.JSON(http.StatusOK, nodeinfolinks)
}
