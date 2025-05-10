// Copyright 2020 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package repo_test

import (
	"testing"

	"code.proxgit.io/proxgit/models/unittest"

	_ "code.proxgit.io/proxgit/models" // register table model
	_ "code.proxgit.io/proxgit/models/actions"
	_ "code.proxgit.io/proxgit/models/activities"
	_ "code.proxgit.io/proxgit/models/perm/access" // register table model
	_ "code.proxgit.io/proxgit/models/repo"        // register table model
	_ "code.proxgit.io/proxgit/models/user"        // register table model
)

func TestMain(m *testing.M) {
	unittest.MainTest(m)
}
