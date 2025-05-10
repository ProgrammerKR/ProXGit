// Copyright 2021 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package org

import (
	"testing"

	"code.proxgit.io/proxgit/models/db"
	"code.proxgit.io/proxgit/models/organization"
	repo_model "code.proxgit.io/proxgit/models/repo"
	"code.proxgit.io/proxgit/models/unittest"
	user_model "code.proxgit.io/proxgit/models/user"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	unittest.MainTest(m)
}

func TestDeleteOrganization(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())
	org := unittest.AssertExistsAndLoadBean(t, &organization.Organization{ID: 6})
	assert.NoError(t, DeleteOrganization(db.DefaultContext, org, false))
	unittest.AssertNotExistsBean(t, &organization.Organization{ID: 6})
	unittest.AssertNotExistsBean(t, &organization.OrgUser{OrgID: 6})
	unittest.AssertNotExistsBean(t, &organization.Team{OrgID: 6})

	org = unittest.AssertExistsAndLoadBean(t, &organization.Organization{ID: 3})
	err := DeleteOrganization(db.DefaultContext, org, false)
	assert.Error(t, err)
	assert.True(t, repo_model.IsErrUserOwnRepos(err))

	user := unittest.AssertExistsAndLoadBean(t, &organization.Organization{ID: 5})
	assert.Error(t, DeleteOrganization(db.DefaultContext, user, false))
	unittest.CheckConsistencyFor(t, &user_model.User{}, &organization.Team{})
}
