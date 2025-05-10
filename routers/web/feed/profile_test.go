// Copyright 2025 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT
package feed_test

import (
	"testing"

	"code.proxgit.io/proxgit/models/unittest"
	user_model "code.proxgit.io/proxgit/models/user"
	"code.proxgit.io/proxgit/routers/web/feed"
	"code.proxgit.io/proxgit/services/contexttest"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	unittest.MainTest(m)
}

func TestCheckGetOrgFeedsAsOrgMember(t *testing.T) {
	unittest.PrepareTestEnv(t)
	t.Run("OrgMember", func(t *testing.T) {
		ctx, resp := contexttest.MockContext(t, "org3.atom")
		ctx.ContextUser = unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 3})
		contexttest.LoadUser(t, ctx, 2)
		ctx.IsSigned = true
		feed.ShowUserFeedAtom(ctx)
		assert.Contains(t, resp.Body.String(), "<entry>") // Should contain 1 private entry
	})
	t.Run("NonOrgMember", func(t *testing.T) {
		ctx, resp := contexttest.MockContext(t, "org3.atom")
		ctx.ContextUser = unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 3})
		contexttest.LoadUser(t, ctx, 5)
		ctx.IsSigned = true
		feed.ShowUserFeedAtom(ctx)
		assert.NotContains(t, resp.Body.String(), "<entry>") // Should not contain any entries
	})
}
