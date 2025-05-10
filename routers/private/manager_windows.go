// Copyright 2020 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

//go:build windows

package private

import (
	"net/http"

<<<<<<< HEAD
	"code.gitea.io/gitea/modules/graceful"
	"code.gitea.io/gitea/modules/private"
	"code.gitea.io/gitea/services/context"
=======
	"code.proxgit.io/proxgit/modules/graceful"
	"code.proxgit.io/proxgit/modules/private"
	"code.proxgit.io/proxgit/services/context"
>>>>>>> master
)

// Restart is not implemented for Windows based servers as they can't fork
func Restart(ctx *context.PrivateContext) {
	ctx.JSON(http.StatusNotImplemented, private.Response{
		UserMsg: "windows servers cannot be gracefully restarted - shutdown and restart manually",
	})
}

// Shutdown causes the server to perform a graceful shutdown
func Shutdown(ctx *context.PrivateContext) {
	graceful.GetManager().DoGracefulShutdown()
	ctx.PlainText(http.StatusOK, "success")
}
