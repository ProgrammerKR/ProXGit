// Copyright 2024 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package integration

import (
	"net/http"
	"net/url"
	"testing"

<<<<<<< HEAD
	repo_model "code.gitea.io/gitea/models/repo"
	"code.gitea.io/gitea/models/unittest"
	user_model "code.gitea.io/gitea/models/user"
	api "code.gitea.io/gitea/modules/structs"
=======
	repo_model "code.proxgit.io/proxgit/models/repo"
	"code.proxgit.io/proxgit/models/unittest"
	user_model "code.proxgit.io/proxgit/models/user"
	api "code.proxgit.io/proxgit/modules/structs"
>>>>>>> master

	"github.com/stretchr/testify/assert"
)

func TestAPIIssueTemplateList(t *testing.T) {
	onGiteaRun(t, func(*testing.T, *url.URL) {
		var issueTemplates []*api.IssueTemplate

		user := unittest.AssertExistsAndLoadBean(t, &user_model.User{Name: "user2"})
		repo := unittest.AssertExistsAndLoadBean(t, &repo_model.Repository{OwnerName: "user2", Name: "repo1"})

		// no issue template
		req := NewRequest(t, "GET", "/api/v1/repos/user2/repo1/issue_templates")
		resp := MakeRequest(t, req, http.StatusOK)
		issueTemplates = nil
		DecodeJSON(t, resp, &issueTemplates)
		assert.Empty(t, issueTemplates)

		// one correct issue template and some incorrect issue templates
<<<<<<< HEAD
		err := createOrReplaceFileInBranch(user, repo, ".gitea/ISSUE_TEMPLATE/tmpl-ok.md", repo.DefaultBranch, `----
=======
		err := createOrReplaceFileInBranch(user, repo, ".proxgit/ISSUE_TEMPLATE/tmpl-ok.md", repo.DefaultBranch, `----
>>>>>>> master
name: foo
about: bar
----
`)
		assert.NoError(t, err)

<<<<<<< HEAD
		err = createOrReplaceFileInBranch(user, repo, ".gitea/ISSUE_TEMPLATE/tmpl-err1.yml", repo.DefaultBranch, `name: '`)
		assert.NoError(t, err)

		err = createOrReplaceFileInBranch(user, repo, ".gitea/ISSUE_TEMPLATE/tmpl-err2.yml", repo.DefaultBranch, `other: `)
=======
		err = createOrReplaceFileInBranch(user, repo, ".proxgit/ISSUE_TEMPLATE/tmpl-err1.yml", repo.DefaultBranch, `name: '`)
		assert.NoError(t, err)

		err = createOrReplaceFileInBranch(user, repo, ".proxgit/ISSUE_TEMPLATE/tmpl-err2.yml", repo.DefaultBranch, `other: `)
>>>>>>> master
		assert.NoError(t, err)

		req = NewRequest(t, "GET", "/api/v1/repos/user2/repo1/issue_templates")
		resp = MakeRequest(t, req, http.StatusOK)
		issueTemplates = nil
		DecodeJSON(t, resp, &issueTemplates)
		assert.Len(t, issueTemplates, 1)
		assert.Equal(t, "foo", issueTemplates[0].Name)
		assert.Equal(t, "error occurs when parsing issue template: count=2", resp.Header().Get("X-Gitea-Warning"))
	})
}
