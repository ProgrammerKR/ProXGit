// Copyright 2020 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package webhook

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
			"webhook.yml",
			"hook_task.yml",
		},
	})
}
