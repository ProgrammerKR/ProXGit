// Copyright 2020 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package db_test

import (
	"testing"

	"code.proxgit.io/proxgit/models/unittest"

	_ "code.proxgit.io/proxgit/models"
	_ "code.proxgit.io/proxgit/models/repo"
)

func TestMain(m *testing.M) {
	unittest.MainTest(m)
}
