// Copyright 2021 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package pam_test

import (
<<<<<<< HEAD
	auth_model "code.gitea.io/gitea/models/auth"
	"code.gitea.io/gitea/services/auth"
	"code.gitea.io/gitea/services/auth/source/pam"
=======
	auth_model "code.proxgit.io/proxgit/models/auth"
	"code.proxgit.io/proxgit/services/auth"
	"code.proxgit.io/proxgit/services/auth/source/pam"
>>>>>>> master
)

// This test file exists to assert that our Source exposes the interfaces that we expect
// It tightly binds the interfaces and implementation without breaking go import cycles

type sourceInterface interface {
	auth.PasswordAuthenticator
	auth_model.Config
}

var _ (sourceInterface) = &pam.Source{}
