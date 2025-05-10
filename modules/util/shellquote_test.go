// Copyright 2020 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShellEscape(t *testing.T) {
	tests := []struct {
		name     string
		toEscape string
		want     string
	}{
		{
			"Simplest case - nothing to escape",
			"a/b/c/d",
			"a/b/c/d",
		}, {
			"Prefixed tilde - with normal stuff - should not escape",
			"~/src/go/proxgit/proxgit",
			"~/src/go/proxgit/proxgit",
		}, {
			"Typical windows path with spaces - should get doublequote escaped",
			`C:\Program Files\Gitea v1.13 - I like lots of spaces\proxgit`,
			`"C:\\Program Files\\Gitea v1.13 - I like lots of spaces\\proxgit"`,
		}, {
			"Forward-slashed windows path with spaces - should get doublequote escaped",
			"C:/Program Files/Gitea v1.13 - I like lots of spaces/proxgit",
			`"C:/Program Files/Gitea v1.13 - I like lots of spaces/proxgit"`,
		}, {
			"Prefixed tilde - but then a space filled path",
			"~git/Gitea v1.13/proxgit",
			`~git/"Gitea v1.13/proxgit"`,
		}, {
			"Bangs are unfortunately not predictable so need to be singlequoted",
			"C:/Program Files/Gitea!/proxgit",
			`'C:/Program Files/Gitea!/proxgit'`,
		}, {
			"Newlines are just irritating",
			"/home/git/Gitea\n\nWHY-WOULD-YOU-DO-THIS\n\nGitea/proxgit",
			"'/home/git/Gitea\n\nWHY-WOULD-YOU-DO-THIS\n\nGitea/proxgit'",
		}, {
			"Similarly we should nicely handle multiple single quotes if we have to single-quote",
			"'!''!'''!''!'!'",
			`\''!'\'\''!'\'\'\''!'\'\''!'\''!'\'`,
		}, {
			"Double quote < ...",
			"~/<proxgit",
			"~/\"<proxgit\"",
		}, {
			"Double quote > ...",
			"~/proxgit>",
			"~/\"proxgit>\"",
		}, {
			"Double quote and escape $ ...",
			"~/$proxgit",
			"~/\"\\$proxgit\"",
		}, {
			"Double quote {...",
			"~/{proxgit",
			"~/\"{proxgit\"",
		}, {
			"Double quote }...",
			"~/proxgit}",
			"~/\"proxgit}\"",
		}, {
			"Double quote ()...",
			"~/(proxgit)",
			"~/\"(proxgit)\"",
		}, {
			"Double quote and escape `...",
			"~/proxgit`",
			"~/\"proxgit\\`\"",
		}, {
			"Double quotes can handle a number of things without having to escape them but not everything ...",
			"~/<proxgit> ${proxgit} `proxgit` [proxgit] (proxgit) \"proxgit\" \\proxgit\\ 'proxgit'",
			"~/\"<proxgit> \\${proxgit} \\`proxgit\\` [proxgit] (proxgit) \\\"proxgit\\\" \\\\proxgit\\\\ 'proxgit'\"",
		}, {
			"Single quotes don't need to escape except for '...",
			"~/<proxgit> ${proxgit} `proxgit` (proxgit) !proxgit! \"proxgit\" \\proxgit\\ 'proxgit'",
			"~/'<proxgit> ${proxgit} `proxgit` (proxgit) !proxgit! \"proxgit\" \\proxgit\\ '\\''proxgit'\\'",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ShellEscape(tt.toEscape))
		})
	}
}
