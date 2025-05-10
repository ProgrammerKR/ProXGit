// Copyright 2021 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package forms

import (
	"net/http"

	"code.proxgit.io/proxgit/modules/web/middleware"
	"code.proxgit.io/proxgit/services/context"

	"proxgit.com/go-chi/binding"
)

// ProtectTagForm form for changing protected tag settings
type ProtectTagForm struct {
	NamePattern    string `binding:"Required;GlobOrRegexPattern"`
	AllowlistUsers string
	AllowlistTeams string
}

// Validate validates the fields
func (f *ProtectTagForm) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	ctx := context.GetValidateContext(req)
	return middleware.Validate(errs, ctx.Data, f, ctx.Locale)
}
