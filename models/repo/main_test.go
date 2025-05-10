// Copyright 2020 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package repo_test

import (
	"testing"

<<<<<<< HEAD
	"code.gitea.io/gitea/models/unittest"

	_ "code.gitea.io/gitea/models" // register table model
	_ "code.gitea.io/gitea/models/actions"
	_ "code.gitea.io/gitea/models/activities"
	_ "code.gitea.io/gitea/models/perm/access" // register table model
	_ "code.gitea.io/gitea/models/repo"        // register table model
	_ "code.gitea.io/gitea/models/user"        // register table model
=======
	"code.proxgit.io/proxgit/models/unittest"

	_ "code.proxgit.io/proxgit/models" // register table model
	_ "code.proxgit.io/proxgit/models/actions"
	_ "code.proxgit.io/proxgit/models/activities"
	_ "code.proxgit.io/proxgit/models/perm/access" // register table model
	_ "code.proxgit.io/proxgit/models/repo"        // register table model
	_ "code.proxgit.io/proxgit/models/user"        // register table model
>>>>>>> master
)

func TestMain(m *testing.M) {
	unittest.MainTest(m)
}
