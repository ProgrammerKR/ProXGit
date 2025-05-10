// Copyright 2025 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package fileicon_test

import (
	"testing"

<<<<<<< HEAD
	"code.gitea.io/gitea/models/unittest"
	"code.gitea.io/gitea/modules/fileicon"
	"code.gitea.io/gitea/modules/git"
=======
	"code.proxgit.io/proxgit/models/unittest"
	"code.proxgit.io/proxgit/modules/fileicon"
	"code.proxgit.io/proxgit/modules/git"
>>>>>>> master

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	unittest.MainTest(m, &unittest.TestOptions{FixtureFiles: []string{}})
}

func TestFindIconName(t *testing.T) {
	unittest.PrepareTestEnv(t)
	p := fileicon.DefaultMaterialIconProvider()
	assert.Equal(t, "php", p.FindIconName(&fileicon.EntryInfo{FullName: "foo.php", EntryMode: git.EntryModeBlob}))
	assert.Equal(t, "php", p.FindIconName(&fileicon.EntryInfo{FullName: "foo.PHP", EntryMode: git.EntryModeBlob}))
	assert.Equal(t, "javascript", p.FindIconName(&fileicon.EntryInfo{FullName: "foo.js", EntryMode: git.EntryModeBlob}))
	assert.Equal(t, "visualstudio", p.FindIconName(&fileicon.EntryInfo{FullName: "foo.vba", EntryMode: git.EntryModeBlob}))
}
