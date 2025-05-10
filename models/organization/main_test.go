// Copyright 2021 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package organization_test

import (
	"testing"

	"code.proxgit.io/proxgit/models/unittest"

	_ "code.proxgit.io/proxgit/models"
	_ "code.proxgit.io/proxgit/models/actions"
	_ "code.proxgit.io/proxgit/models/activities"
	_ "code.proxgit.io/proxgit/models/organization"
	_ "code.proxgit.io/proxgit/models/repo"
	_ "code.proxgit.io/proxgit/models/user"
)

func TestMain(m *testing.M) {
	unittest.MainTest(m)
}
