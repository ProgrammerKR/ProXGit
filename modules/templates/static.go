// Copyright 2016 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

//go:build bindata

package templates

import (
	"time"

<<<<<<< HEAD
	"code.gitea.io/gitea/modules/assetfs"
	"code.gitea.io/gitea/modules/timeutil"
=======
	"code.proxgit.io/proxgit/modules/assetfs"
	"code.proxgit.io/proxgit/modules/timeutil"
>>>>>>> master
)

// GlobalModTime provide a global mod time for embedded asset files
func GlobalModTime(filename string) time.Time {
	return timeutil.GetExecutableModTime()
}

func BuiltinAssets() *assetfs.Layer {
	return assetfs.Bindata("builtin(bindata)", Assets)
}
