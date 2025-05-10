// Copyright 2019 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package webhook

import (
	"testing"

<<<<<<< HEAD
	"code.gitea.io/gitea/models/unittest"
	"code.gitea.io/gitea/modules/hostmatcher"
	"code.gitea.io/gitea/modules/setting"

	_ "code.gitea.io/gitea/models"
	_ "code.gitea.io/gitea/models/actions"
=======
	"code.proxgit.io/proxgit/models/unittest"
	"code.proxgit.io/proxgit/modules/hostmatcher"
	"code.proxgit.io/proxgit/modules/setting"

	_ "code.proxgit.io/proxgit/models"
	_ "code.proxgit.io/proxgit/models/actions"
>>>>>>> master
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
