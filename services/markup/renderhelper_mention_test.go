// Copyright 2022 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package markup

import (
	"net/http"
	"net/http/httptest"
	"testing"

<<<<<<< HEAD
	"code.gitea.io/gitea/models/db"
	"code.gitea.io/gitea/models/unittest"
	"code.gitea.io/gitea/models/user"
	gitea_context "code.gitea.io/gitea/services/context"
	"code.gitea.io/gitea/services/contexttest"
=======
	"code.proxgit.io/proxgit/models/db"
	"code.proxgit.io/proxgit/models/unittest"
	"code.proxgit.io/proxgit/models/user"
	proxgit_context "code.proxgit.io/proxgit/services/context"
	"code.proxgit.io/proxgit/services/contexttest"
>>>>>>> master

	"github.com/stretchr/testify/assert"
)

func TestRenderHelperMention(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	userPublic := "user1"
	userPrivate := "user31"
	userLimited := "user33"
	userNoSuch := "no-such-user"

	unittest.AssertCount(t, &user.User{Name: userPublic}, 1)
	unittest.AssertCount(t, &user.User{Name: userPrivate}, 1)
	unittest.AssertCount(t, &user.User{Name: userLimited}, 1)
	unittest.AssertCount(t, &user.User{Name: userNoSuch}, 0)

	// when using general context, use user's visibility to check
	assert.True(t, FormalRenderHelperFuncs().IsUsernameMentionable(t.Context(), userPublic))
	assert.False(t, FormalRenderHelperFuncs().IsUsernameMentionable(t.Context(), userLimited))
	assert.False(t, FormalRenderHelperFuncs().IsUsernameMentionable(t.Context(), userPrivate))
	assert.False(t, FormalRenderHelperFuncs().IsUsernameMentionable(t.Context(), userNoSuch))

	// when using web context, use user.IsUserVisibleToViewer to check
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	assert.NoError(t, err)
<<<<<<< HEAD
	base := gitea_context.NewBaseContextForTest(httptest.NewRecorder(), req)
	giteaCtx := gitea_context.NewWebContext(base, &contexttest.MockRender{}, nil)

	assert.True(t, FormalRenderHelperFuncs().IsUsernameMentionable(giteaCtx, userPublic))
	assert.False(t, FormalRenderHelperFuncs().IsUsernameMentionable(giteaCtx, userPrivate))

	giteaCtx.Doer, err = user.GetUserByName(db.DefaultContext, userPrivate)
	assert.NoError(t, err)
	assert.True(t, FormalRenderHelperFuncs().IsUsernameMentionable(giteaCtx, userPublic))
	assert.True(t, FormalRenderHelperFuncs().IsUsernameMentionable(giteaCtx, userPrivate))
=======
	base := proxgit_context.NewBaseContextForTest(httptest.NewRecorder(), req)
	proxgitCtx := proxgit_context.NewWebContext(base, &contexttest.MockRender{}, nil)

	assert.True(t, FormalRenderHelperFuncs().IsUsernameMentionable(proxgitCtx, userPublic))
	assert.False(t, FormalRenderHelperFuncs().IsUsernameMentionable(proxgitCtx, userPrivate))

	proxgitCtx.Doer, err = user.GetUserByName(db.DefaultContext, userPrivate)
	assert.NoError(t, err)
	assert.True(t, FormalRenderHelperFuncs().IsUsernameMentionable(proxgitCtx, userPublic))
	assert.True(t, FormalRenderHelperFuncs().IsUsernameMentionable(proxgitCtx, userPrivate))
>>>>>>> master
}
