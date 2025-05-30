// Copyright 2024 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package rpm

import (
	"context"

<<<<<<< HEAD
	packages_model "code.gitea.io/gitea/models/packages"
	rpm_module "code.gitea.io/gitea/modules/packages/rpm"
=======
	packages_model "code.proxgit.io/proxgit/models/packages"
	rpm_module "code.proxgit.io/proxgit/modules/packages/rpm"
>>>>>>> master
)

// GetGroups gets all available groups
func GetGroups(ctx context.Context, ownerID int64) ([]string, error) {
	return packages_model.GetDistinctPropertyValues(
		ctx,
		packages_model.TypeRpm,
		ownerID,
		packages_model.PropertyTypeFile,
		rpm_module.PropertyGroup,
		nil,
	)
}
