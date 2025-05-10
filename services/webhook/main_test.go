// Copyright 2019 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package webhook

import (
	"testing"

	"code.proxgit.io/proxgit/models/unittest"
	"code.proxgit.io/proxgit/modules/hostmatcher"
	"code.proxgit.io/proxgit/modules/setting"

	_ "code.proxgit.io/proxgit/models"
	_ "code.proxgit.io/proxgit/models/actions"
)

func TestMain(m *testing.M) {
	// for tests, allow only loopback IPs
	setting.Webhook.AllowedHostList = hostmatcher.MatchBuiltinLoopback
	unittest.MainTest(m, &unittest.TestOptions{
		SetUp: func() error {
			setting.LoadQueueSettings()
			return Init()
		},
	})
}
