// Copyright 2020 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package swagger

import (
<<<<<<< HEAD
	api "code.gitea.io/gitea/modules/structs"
=======
	api "code.proxgit.io/proxgit/modules/structs"
>>>>>>> master
)

// CronList
// swagger:response CronList
type swaggerResponseCronList struct {
	// in:body
	Body []api.Cron `json:"body"`
}
