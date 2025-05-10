// Copyright 2021 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package git_test

import (
	"testing"

<<<<<<< HEAD
	"code.gitea.io/gitea/models/db"
	git_model "code.gitea.io/gitea/models/git"
	"code.gitea.io/gitea/models/unittest"
=======
	"code.proxgit.io/proxgit/models/db"
	git_model "code.proxgit.io/proxgit/models/git"
	"code.proxgit.io/proxgit/models/unittest"
>>>>>>> master

	"github.com/stretchr/testify/assert"
)

func TestIsUserAllowed(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	pt := &git_model.ProtectedTag{}
	allowed, err := git_model.IsUserAllowedModifyTag(db.DefaultContext, pt, 1)
	assert.NoError(t, err)
	assert.False(t, allowed)

	pt = &git_model.ProtectedTag{
		AllowlistUserIDs: []int64{1},
	}
	allowed, err = git_model.IsUserAllowedModifyTag(db.DefaultContext, pt, 1)
	assert.NoError(t, err)
	assert.True(t, allowed)

	allowed, err = git_model.IsUserAllowedModifyTag(db.DefaultContext, pt, 2)
	assert.NoError(t, err)
	assert.False(t, allowed)

	pt = &git_model.ProtectedTag{
		AllowlistTeamIDs: []int64{1},
	}
	allowed, err = git_model.IsUserAllowedModifyTag(db.DefaultContext, pt, 1)
	assert.NoError(t, err)
	assert.False(t, allowed)

	allowed, err = git_model.IsUserAllowedModifyTag(db.DefaultContext, pt, 2)
	assert.NoError(t, err)
	assert.True(t, allowed)

	pt = &git_model.ProtectedTag{
		AllowlistUserIDs: []int64{1},
		AllowlistTeamIDs: []int64{1},
	}
	allowed, err = git_model.IsUserAllowedModifyTag(db.DefaultContext, pt, 1)
	assert.NoError(t, err)
	assert.True(t, allowed)

	allowed, err = git_model.IsUserAllowedModifyTag(db.DefaultContext, pt, 2)
	assert.NoError(t, err)
	assert.True(t, allowed)
}

func TestIsUserAllowedToControlTag(t *testing.T) {
	cases := []struct {
		name    string
		userid  int64
		allowed bool
	}{
		{
			name:    "test",
			userid:  1,
			allowed: true,
		},
		{
			name:    "test",
			userid:  3,
			allowed: true,
		},
		{
<<<<<<< HEAD
			name:    "gitea",
=======
			name:    "proxgit",
>>>>>>> master
			userid:  1,
			allowed: true,
		},
		{
<<<<<<< HEAD
			name:    "gitea",
=======
			name:    "proxgit",
>>>>>>> master
			userid:  3,
			allowed: false,
		},
		{
<<<<<<< HEAD
			name:    "test-gitea",
=======
			name:    "test-proxgit",
>>>>>>> master
			userid:  1,
			allowed: true,
		},
		{
<<<<<<< HEAD
			name:    "test-gitea",
=======
			name:    "test-proxgit",
>>>>>>> master
			userid:  3,
			allowed: false,
		},
		{
<<<<<<< HEAD
			name:    "gitea-test",
=======
			name:    "proxgit-test",
>>>>>>> master
			userid:  1,
			allowed: true,
		},
		{
<<<<<<< HEAD
			name:    "gitea-test",
=======
			name:    "proxgit-test",
>>>>>>> master
			userid:  3,
			allowed: true,
		},
		{
			name:    "v-1",
			userid:  1,
			allowed: false,
		},
		{
			name:    "v-1",
			userid:  2,
			allowed: true,
		},
		{
			name:    "release",
			userid:  1,
			allowed: false,
		},
	}

	t.Run("Glob", func(t *testing.T) {
		protectedTags := []*git_model.ProtectedTag{
			{
<<<<<<< HEAD
				NamePattern:      `*gitea`,
=======
				NamePattern:      `*proxgit`,
>>>>>>> master
				AllowlistUserIDs: []int64{1},
			},
			{
				NamePattern:      `v-*`,
				AllowlistUserIDs: []int64{2},
			},
			{
				NamePattern: "release",
			},
		}

		for n, c := range cases {
			isAllowed, err := git_model.IsUserAllowedToControlTag(db.DefaultContext, protectedTags, c.name, c.userid)
			assert.NoError(t, err)
			assert.Equal(t, c.allowed, isAllowed, "case %d: error should match", n)
		}
	})

	t.Run("Regex", func(t *testing.T) {
		protectedTags := []*git_model.ProtectedTag{
			{
<<<<<<< HEAD
				NamePattern:      `/gitea\z/`,
=======
				NamePattern:      `/proxgit\z/`,
>>>>>>> master
				AllowlistUserIDs: []int64{1},
			},
			{
				NamePattern:      `/\Av-/`,
				AllowlistUserIDs: []int64{2},
			},
			{
				NamePattern: "/release/",
			},
		}

		for n, c := range cases {
			isAllowed, err := git_model.IsUserAllowedToControlTag(db.DefaultContext, protectedTags, c.name, c.userid)
			assert.NoError(t, err)
			assert.Equal(t, c.allowed, isAllowed, "case %d: error should match", n)
		}
	})
}
