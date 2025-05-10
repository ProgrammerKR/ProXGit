// Copyright 2024 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package renderhelper

import (
	"context"
	"testing"

	"code.proxgit.io/proxgit/models/unittest"
	"code.proxgit.io/proxgit/modules/markup"
)

func TestMain(m *testing.M) {
	unittest.MainTest(m, &unittest.TestOptions{
		FixtureFiles: []string{"repository.yml", "user.yml"},
		SetUp: func() error {
			markup.RenderBehaviorForTesting.DisableAdditionalAttributes = true
			markup.Init(&markup.RenderHelperFuncs{
				IsUsernameMentionable: func(ctx context.Context, username string) bool {
					return username == "user2"
				},
			})
			return nil
		},
	})
}
