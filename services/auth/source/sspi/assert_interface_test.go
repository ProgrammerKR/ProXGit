// Copyright 2021 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package sspi_test

import (
<<<<<<< HEAD
	"code.gitea.io/gitea/models/auth"
	"code.gitea.io/gitea/services/auth/source/sspi"
=======
	"code.proxgit.io/proxgit/models/auth"
	"code.proxgit.io/proxgit/services/auth/source/sspi"
>>>>>>> master
)

// This test file exists to assert that our Source exposes the interfaces that we expect
// It tightly binds the interfaces and implementation without breaking go import cycles

type sourceInterface interface {
	auth.Config
}

var _ (sourceInterface) = &sspi.Source{}
