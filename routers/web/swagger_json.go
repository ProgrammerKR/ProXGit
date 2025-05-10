// Copyright 2020 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package web

import (
<<<<<<< HEAD
	"code.gitea.io/gitea/services/context"
=======
	"code.proxgit.io/proxgit/services/context"
>>>>>>> master
)

// SwaggerV1Json render swagger v1 json
func SwaggerV1Json(ctx *context.Context) {
	ctx.JSONTemplate("swagger/v1_json")
}
