// Copyright 2020 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package project

import (
	"testing"

<<<<<<< HEAD
	"code.gitea.io/gitea/models/unittest"

	_ "code.gitea.io/gitea/models/repo"
=======
	"code.proxgit.io/proxgit/models/unittest"

	_ "code.proxgit.io/proxgit/models/repo"
>>>>>>> master
)

func TestMain(m *testing.M) {
	unittest.MainTest(m, &unittest.TestOptions{
		FixtureFiles: []string{
			"project.yml",
			"project_board.yml",
			"project_issue.yml",
			"repository.yml",
		},
	})
}
