// Copyright 2023 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package convert

import (
	"testing"

	"code.proxgit.io/proxgit/models/db"
	repo_model "code.proxgit.io/proxgit/models/repo"
	"code.proxgit.io/proxgit/models/unittest"

	"github.com/stretchr/testify/assert"
)

func TestRelease_ToRelease(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	repo1 := unittest.AssertExistsAndLoadBean(t, &repo_model.Repository{ID: 1})
	release1 := unittest.AssertExistsAndLoadBean(t, &repo_model.Release{ID: 1})
	release1.LoadAttributes(db.DefaultContext)

	apiRelease := ToAPIRelease(db.DefaultContext, repo1, release1)
	assert.NotNil(t, apiRelease)
	assert.EqualValues(t, 1, apiRelease.ID)
	assert.Equal(t, "https://try.proxgit.io/api/v1/repos/user2/repo1/releases/1", apiRelease.URL)
	assert.Equal(t, "https://try.proxgit.io/api/v1/repos/user2/repo1/releases/1/assets", apiRelease.UploadURL)
}
