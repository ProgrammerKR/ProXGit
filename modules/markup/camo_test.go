// Copyright 2022 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package markup

import (
	"testing"

<<<<<<< HEAD
	"code.gitea.io/gitea/modules/setting"
=======
	"code.proxgit.io/proxgit/modules/setting"
>>>>>>> master

	"github.com/stretchr/testify/assert"
)

func TestCamoHandleLink(t *testing.T) {
<<<<<<< HEAD
	setting.AppURL = "https://gitea.com"
=======
	setting.AppURL = "https://proxgit.com"
>>>>>>> master
	// Test media proxy
	setting.Camo.Enabled = true
	setting.Camo.ServerURL = "https://image.proxy"
	setting.Camo.HMACKey = "geheim"

	assert.Equal(t,
<<<<<<< HEAD
		"https://gitea.com/img.jpg",
		camoHandleLink("https://gitea.com/img.jpg"))
=======
		"https://proxgit.com/img.jpg",
		camoHandleLink("https://proxgit.com/img.jpg"))
>>>>>>> master
	assert.Equal(t,
		"https://testimages.org/img.jpg",
		camoHandleLink("https://testimages.org/img.jpg"))
	assert.Equal(t,
		"https://image.proxy/eivin43gJwGVIjR9MiYYtFIk0mw/aHR0cDovL3Rlc3RpbWFnZXMub3JnL2ltZy5qcGc",
		camoHandleLink("http://testimages.org/img.jpg"))

	setting.Camo.Always = true
	assert.Equal(t,
<<<<<<< HEAD
		"https://gitea.com/img.jpg",
		camoHandleLink("https://gitea.com/img.jpg"))
=======
		"https://proxgit.com/img.jpg",
		camoHandleLink("https://proxgit.com/img.jpg"))
>>>>>>> master
	assert.Equal(t,
		"https://image.proxy/tkdlvmqpbIr7SjONfHNgEU622y0/aHR0cHM6Ly90ZXN0aW1hZ2VzLm9yZy9pbWcuanBn",
		camoHandleLink("https://testimages.org/img.jpg"))
	assert.Equal(t,
		"https://image.proxy/eivin43gJwGVIjR9MiYYtFIk0mw/aHR0cDovL3Rlc3RpbWFnZXMub3JnL2ltZy5qcGc",
		camoHandleLink("http://testimages.org/img.jpg"))

	// Restore previous settings
	setting.Camo.Enabled = false
}
