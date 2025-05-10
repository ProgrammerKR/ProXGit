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
<<<<<<< HEAD
			"~/src/go/gitea/gitea",
			"~/src/go/gitea/gitea",
		}, {
			"Typical windows path with spaces - should get doublequote escaped",
			`C:\Program Files\Gitea v1.13 - I like lots of spaces\gitea`,
			`"C:\\Program Files\\Gitea v1.13 - I like lots of spaces\\gitea"`,
		}, {
			"Forward-slashed windows path with spaces - should get doublequote escaped",
			"C:/Program Files/Gitea v1.13 - I like lots of spaces/gitea",
			`"C:/Program Files/Gitea v1.13 - I like lots of spaces/gitea"`,
		}, {
			"Prefixed tilde - but then a space filled path",
			"~git/Gitea v1.13/gitea",
			`~git/"Gitea v1.13/gitea"`,
		}, {
			"Bangs are unfortunately not predictable so need to be singlequoted",
			"C:/Program Files/Gitea!/gitea",
			`'C:/Program Files/Gitea!/gitea'`,
		}, {
			"Newlines are just irritating",
			"/home/git/Gitea\n\nWHY-WOULD-YOU-DO-THIS\n\nGitea/gitea",
			"'/home/git/Gitea\n\nWHY-WOULD-YOU-DO-THIS\n\nGitea/gitea'",
=======
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
>>>>>>> master
		}, {
			"Similarly we should nicely handle multiple single quotes if we have to single-quote",
			"'!''!'''!''!'!'",
			`\''!'\'\''!'\'\'\''!'\'\''!'\''!'\'`,
		}, {
			"Double quote < ...",
<<<<<<< HEAD
			"~/<gitea",
			"~/\"<gitea\"",
		}, {
			"Double quote > ...",
			"~/gitea>",
			"~/\"gitea>\"",
		}, {
			"Double quote and escape $ ...",
			"~/$gitea",
			"~/\"\\$gitea\"",
		}, {
			"Double quote {...",
			"~/{gitea",
			"~/\"{gitea\"",
		}, {
			"Double quote }...",
			"~/gitea}",
			"~/\"gitea}\"",
		}, {
			"Double quote ()...",
			"~/(gitea)",
			"~/\"(gitea)\"",
		}, {
			"Double quote and escape `...",
			"~/gitea`",
			"~/\"gitea\\`\"",
		}, {
			"Double quotes can handle a number of things without having to escape them but not everything ...",
			"~/<gitea> ${gitea} `gitea` [gitea] (gitea) \"gitea\" \\gitea\\ 'gitea'",
			"~/\"<gitea> \\${gitea} \\`gitea\\` [gitea] (gitea) \\\"gitea\\\" \\\\gitea\\\\ 'gitea'\"",
		}, {
			"Single quotes don't need to escape except for '...",
			"~/<gitea> ${gitea} `gitea` (gitea) !gitea! \"gitea\" \\gitea\\ 'gitea'",
			"~/'<gitea> ${gitea} `gitea` (gitea) !gitea! \"gitea\" \\gitea\\ '\\''gitea'\\'",
=======
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
>>>>>>> master
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ShellEscape(tt.toEscape))
		})
	}
}
