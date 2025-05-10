// Copyright 2021 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package oauth2_test

import (
<<<<<<< HEAD
	auth_model "code.gitea.io/gitea/models/auth"
	"code.gitea.io/gitea/services/auth"
	"code.gitea.io/gitea/services/auth/source/oauth2"
=======
	auth_model "code.proxgit.io/proxgit/models/auth"
	"code.proxgit.io/proxgit/services/auth"
	"code.proxgit.io/proxgit/services/auth/source/oauth2"
>>>>>>> master
)

// This test file exists to assert that our Source exposes the interfaces that we expect
// It tightly binds the interfaces and implementation without breaking go import cycles

type sourceInterface interface {
	auth_model.Config
	auth_model.RegisterableSource
	auth.PasswordAuthenticator
}

var _ (sourceInterface) = &oauth2.Source{}
