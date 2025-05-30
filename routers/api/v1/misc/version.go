// Copyright 2017 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package misc

import (
	"net/http"

<<<<<<< HEAD
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/structs"
	"code.gitea.io/gitea/services/context"
=======
	"code.proxgit.io/proxgit/modules/setting"
	"code.proxgit.io/proxgit/modules/structs"
	"code.proxgit.io/proxgit/services/context"
>>>>>>> master
)

// Version shows the version of the Gitea server
func Version(ctx *context.APIContext) {
	// swagger:operation GET /version miscellaneous getVersion
	// ---
	// summary: Returns the version of the Gitea application
	// produces:
	// - application/json
	// responses:
	//   "200":
	//     "$ref": "#/responses/ServerVersion"
	ctx.JSON(http.StatusOK, &structs.ServerVersion{Version: setting.AppVer})
}
