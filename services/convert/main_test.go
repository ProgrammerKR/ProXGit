// Copyright 2020 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package convert

import (
	"testing"

	"code.proxgit.io/proxgit/models/unittest"

	_ "code.proxgit.io/proxgit/models/actions"
)

func TestMain(m *testing.M) {
	unittest.MainTest(m)
}
