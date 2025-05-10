// Copyright 2022 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

//go:build bindata

package options

import (
<<<<<<< HEAD
	"code.gitea.io/gitea/modules/assetfs"
=======
	"code.proxgit.io/proxgit/modules/assetfs"
>>>>>>> master
)

func BuiltinAssets() *assetfs.Layer {
	return assetfs.Bindata("builtin(bindata)", Assets)
}
