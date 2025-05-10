// Copyright 2020 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package models

import (
	"testing"

	activities_model "code.proxgit.io/proxgit/models/activities"
	"code.proxgit.io/proxgit/models/organization"
	repo_model "code.proxgit.io/proxgit/models/repo"
	"code.proxgit.io/proxgit/models/unittest"
	user_model "code.proxgit.io/proxgit/models/user"

	_ "code.proxgit.io/proxgit/models/actions"
	_ "code.proxgit.io/proxgit/models/system"

	"github.com/stretchr/testify/assert"
)

// TestFixturesAreConsistent assert that test fixtures are consistent
func TestFixturesAreConsistent(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())
	unittest.CheckConsistencyFor(t,
		&user_model.User{},
		&repo_model.Repository{},
		&organization.Team{},
		&activities_model.Action{})
}

func TestMain(m *testing.M) {
	unittest.MainTest(m)
}
