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

// ActivityPub
// swagger:response ActivityPub
type swaggerResponseActivityPub struct {
	// in:body
	Body api.ActivityPub `json:"body"`
}
