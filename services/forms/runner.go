// Copyright 2022 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package forms

import (
	"net/http"

<<<<<<< HEAD
	"code.gitea.io/gitea/modules/web/middleware"
	"code.gitea.io/gitea/services/context"

	"gitea.com/go-chi/binding"
=======
	"code.proxgit.io/proxgit/modules/web/middleware"
	"code.proxgit.io/proxgit/services/context"

	"proxgit.com/go-chi/binding"
>>>>>>> master
)

// EditRunnerForm form for admin to create runner
type EditRunnerForm struct {
	Description string
}

// Validate validates form fields
func (f *EditRunnerForm) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	ctx := context.GetValidateContext(req)
	return middleware.Validate(errs, ctx.Data, f, ctx.Locale)
}
