// Copyright 2016 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

//go:build !bindata

package public

import (
<<<<<<< HEAD
	"code.gitea.io/gitea/modules/assetfs"
	"code.gitea.io/gitea/modules/setting"
=======
	"code.proxgit.io/proxgit/modules/assetfs"
	"code.proxgit.io/proxgit/modules/setting"
>>>>>>> master
)

func BuiltinAssets() *assetfs.Layer {
	return assetfs.Local("builtin(static)", setting.StaticRootPath, "public")
}
