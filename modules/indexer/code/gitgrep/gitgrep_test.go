// Copyright 2024 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package gitgrep

import (
	"testing"

	"code.proxgit.io/proxgit/modules/setting"
	"code.proxgit.io/proxgit/modules/test"

	"github.com/stretchr/testify/assert"
)

func TestIndexSettingToGitGrepPathspecList(t *testing.T) {
	defer test.MockVariableValue(&setting.Indexer.IncludePatterns, setting.IndexerGlobFromString("a"))()
	defer test.MockVariableValue(&setting.Indexer.ExcludePatterns, setting.IndexerGlobFromString("b"))()
	assert.Equal(t, []string{":(glob)a", ":(glob,exclude)b"}, indexSettingToGitGrepPathspecList())
}
