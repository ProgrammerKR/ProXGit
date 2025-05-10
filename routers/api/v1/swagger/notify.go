// Copyright 2019 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package swagger

import (
<<<<<<< HEAD
	api "code.gitea.io/gitea/modules/structs"
=======
	api "code.proxgit.io/proxgit/modules/structs"
>>>>>>> master
)

// NotificationThread
// swagger:response NotificationThread
type swaggerNotificationThread struct {
	// in:body
	Body api.NotificationThread `json:"body"`
}

// NotificationThreadList
// swagger:response NotificationThreadList
type swaggerNotificationThreadList struct {
	// in:body
	Body []api.NotificationThread `json:"body"`
}

// Number of unread notifications
// swagger:response NotificationCount
type swaggerNotificationCount struct {
	// in:body
	Body api.NotificationCount `json:"body"`
}
