// Copyright 2022 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package activitypub

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

<<<<<<< HEAD
	"code.gitea.io/gitea/models/db"
	"code.gitea.io/gitea/models/unittest"
	user_model "code.gitea.io/gitea/models/user"
	"code.gitea.io/gitea/modules/setting"
=======
	"code.proxgit.io/proxgit/models/db"
	"code.proxgit.io/proxgit/models/unittest"
	user_model "code.proxgit.io/proxgit/models/user"
	"code.proxgit.io/proxgit/modules/setting"
>>>>>>> master

	"github.com/stretchr/testify/assert"
)

func TestActivityPubSignedPost(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())
	user := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 1})
	pubID := "https://example.com/pubID"
	c, err := NewClient(db.DefaultContext, user, pubID)
	assert.NoError(t, err)

	expected := "BODY"
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Regexp(t, "^"+setting.Federation.DigestAlgorithm, r.Header.Get("Digest"))
		assert.Contains(t, r.Header.Get("Signature"), pubID)
		assert.Equal(t, ActivityStreamsContentType, r.Header.Get("Content-Type"))
		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)
		assert.Equal(t, expected, string(body))
		fmt.Fprint(w, expected)
	}))
	defer srv.Close()

	r, err := c.Post([]byte(expected), srv.URL)
	assert.NoError(t, err)
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	assert.NoError(t, err)
	assert.Equal(t, expected, string(body))
}
