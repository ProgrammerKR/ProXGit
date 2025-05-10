// Copyright 2024 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package markup_test

import (
	"os"
	"testing"

<<<<<<< HEAD
	"code.gitea.io/gitea/modules/markup"
	"code.gitea.io/gitea/modules/setting"
=======
	"code.proxgit.io/proxgit/modules/markup"
	"code.proxgit.io/proxgit/modules/setting"
>>>>>>> master
)

func TestMain(m *testing.M) {
	setting.IsInTesting = true
	markup.RenderBehaviorForTesting.DisableAdditionalAttributes = true
	os.Exit(m.Run())
}
