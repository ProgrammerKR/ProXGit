// Copyright 2017 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package repo_test

import (
	"path/filepath"
	"testing"

	repo_model "code.proxgit.io/proxgit/models/repo"
	"code.proxgit.io/proxgit/models/unittest"
	"code.proxgit.io/proxgit/modules/setting"

	"github.com/stretchr/testify/assert"
)

func TestRepository_WikiCloneLink(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	repo := unittest.AssertExistsAndLoadBean(t, &repo_model.Repository{ID: 1})
	cloneLink := repo.WikiCloneLink(t.Context(), nil)
	assert.Equal(t, "ssh://sshuser@try.proxgit.io:3000/user2/repo1.wiki.git", cloneLink.SSH)
	assert.Equal(t, "https://try.proxgit.io/user2/repo1.wiki.git", cloneLink.HTTPS)
}

func TestWikiPath(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())
	expected := filepath.Join(setting.RepoRootPath, "user2/repo1.wiki.git")
	assert.Equal(t, expected, repo_model.WikiPath("user2", "repo1"))
}

func TestRepository_WikiPath(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())
	repo := unittest.AssertExistsAndLoadBean(t, &repo_model.Repository{ID: 1})
	expected := filepath.Join(setting.RepoRootPath, "user2/repo1.wiki.git")
	assert.Equal(t, expected, repo.WikiPath())
}

func TestRepository_HasWiki(t *testing.T) {
	unittest.PrepareTestEnv(t)
	repo1 := unittest.AssertExistsAndLoadBean(t, &repo_model.Repository{ID: 1})
	assert.True(t, repo1.HasWiki())
	repo2 := unittest.AssertExistsAndLoadBean(t, &repo_model.Repository{ID: 2})
	assert.False(t, repo2.HasWiki())
}
