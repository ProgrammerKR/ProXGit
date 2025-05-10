// Copyright 2024 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package markup_test

import (
	"context"
	"html/template"
	"strings"
	"testing"

	"code.proxgit.io/proxgit/modules/htmlutil"
	"code.proxgit.io/proxgit/modules/markup"
	"code.proxgit.io/proxgit/modules/markup/markdown"
	testModule "code.proxgit.io/proxgit/modules/test"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRender_IssueList(t *testing.T) {
	defer testModule.MockVariableValue(&markup.RenderBehaviorForTesting.DisableAdditionalAttributes, true)()
	markup.Init(&markup.RenderHelperFuncs{
		RenderRepoIssueIconTitle: func(ctx context.Context, opts markup.RenderIssueIconTitleOptions) (template.HTML, error) {
			return htmlutil.HTMLFormat("<div>issue #%d</div>", opts.IssueIndex), nil
		},
	})

	test := func(input, expected string) {
		rctx := markup.NewTestRenderContext(markup.TestAppURL, map[string]string{
			"user": "test-user", "repo": "test-repo",
			"markupAllowShortIssuePattern": "true",
		})
		out, err := markdown.RenderString(rctx, input)
		require.NoError(t, err)
		assert.Equal(t, strings.TrimSpace(expected), strings.TrimSpace(string(out)))
	}

	t.Run("NormalIssueRef", func(t *testing.T) {
		test(
			"#12345",
			`<p><a href="/test-user/test-repo/issues/12345" class="ref-issue" rel="nofollow">#12345</a></p>`,
		)
	})

	t.Run("ListIssueRef", func(t *testing.T) {
		test(
			"* #12345",
			`<ul>
<li><div>issue #12345</div></li>
</ul>`,
		)
	})

	t.Run("ListIssueRefNormal", func(t *testing.T) {
		test(
			"* foo #12345 bar",
			`<ul>
<li>foo <a href="/test-user/test-repo/issues/12345" class="ref-issue" rel="nofollow">#12345</a> bar</li>
</ul>`,
		)
	})

	t.Run("ListTodoIssueRef", func(t *testing.T) {
		test(
			"* [ ] #12345",
			`<ul>
<li class="task-list-item"><input type="checkbox" disabled="" data-source-position="2"/><div>issue #12345</div></li>
</ul>`,
		)
	})
}
