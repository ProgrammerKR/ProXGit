// Copyright 2022 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package swagger

import (
<<<<<<< HEAD
	api "code.gitea.io/gitea/modules/structs"
=======
	api "code.proxgit.io/proxgit/modules/structs"
>>>>>>> master
)

// Package
// swagger:response Package
type swaggerResponsePackage struct {
	// in:body
	Body api.Package `json:"body"`
}

// PackageList
// swagger:response PackageList
type swaggerResponsePackageList struct {
	// in:body
	Body []api.Package `json:"body"`
}

// PackageFileList
// swagger:response PackageFileList
type swaggerResponsePackageFileList struct {
	// in:body
	Body []api.PackageFile `json:"body"`
}
