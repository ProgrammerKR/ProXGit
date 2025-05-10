// Copyright 2020 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package auth_test

import (
	"testing"

	"code.proxgit.io/proxgit/models/unittest"

	_ "code.proxgit.io/proxgit/models"
	_ "code.proxgit.io/proxgit/models/actions"
	_ "code.proxgit.io/proxgit/models/activities"
	_ "code.proxgit.io/proxgit/models/auth"
	_ "code.proxgit.io/proxgit/models/perm/access"
)

func TestMain(m *testing.M) {
	unittest.MainTest(m)
}
