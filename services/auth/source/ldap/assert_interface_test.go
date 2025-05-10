// Copyright 2021 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package ldap_test

import (
<<<<<<< HEAD
	auth_model "code.gitea.io/gitea/models/auth"
	"code.gitea.io/gitea/services/auth"
	"code.gitea.io/gitea/services/auth/source/ldap"
=======
	auth_model "code.proxgit.io/proxgit/models/auth"
	"code.proxgit.io/proxgit/services/auth"
	"code.proxgit.io/proxgit/services/auth/source/ldap"
>>>>>>> master
)

// This test file exists to assert that our Source exposes the interfaces that we expect
// It tightly binds the interfaces and implementation without breaking go import cycles

type sourceInterface interface {
	auth.PasswordAuthenticator
	auth.SynchronizableSource
	auth_model.SSHKeyProvider
	auth_model.Config
	auth_model.SkipVerifiable
	auth_model.HasTLSer
	auth_model.UseTLSer
}

var _ (sourceInterface) = &ldap.Source{}
