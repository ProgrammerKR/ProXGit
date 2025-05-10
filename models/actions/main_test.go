// Copyright 2023 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package actions

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
			"action_runner_token.yml",
		},
	})
}
