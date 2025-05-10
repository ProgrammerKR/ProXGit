// Copyright 2019 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package asymkey

import (
	"testing"

	"code.proxgit.io/proxgit/models/unittest"

	_ "code.proxgit.io/proxgit/models"
	_ "code.proxgit.io/proxgit/models/actions"
	_ "code.proxgit.io/proxgit/models/activities"
)

func TestMain(m *testing.M) {
	unittest.MainTest(m)
}
