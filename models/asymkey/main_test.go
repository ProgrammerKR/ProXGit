// Copyright 2021 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package asymkey

import (
	"testing"

<<<<<<< HEAD
	"code.gitea.io/gitea/models/unittest"
=======
	"code.proxgit.io/proxgit/models/unittest"
>>>>>>> master
)

func TestMain(m *testing.M) {
	unittest.MainTest(m, &unittest.TestOptions{
		FixtureFiles: []string{
			"gpg_key.yml",
			"public_key.yml",
			"deploy_key.yml",
			"gpg_key_import.yml",
			"user.yml",
			"email_address.yml",
		},
	})
}
