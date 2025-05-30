// Copyright 2017 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

<<<<<<< HEAD
package structs // import "code.gitea.io/gitea/modules/structs"
=======
package structs // import "code.proxgit.io/proxgit/modules/structs"
>>>>>>> master

import (
	"time"
)

// Attachment a generic attachment
// swagger:model
type Attachment struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	Size          int64  `json:"size"`
	DownloadCount int64  `json:"download_count"`
	// swagger:strfmt date-time
	Created     time.Time `json:"created_at"`
	UUID        string    `json:"uuid"`
	DownloadURL string    `json:"browser_download_url"`
}

// EditAttachmentOptions options for editing attachments
// swagger:model
type EditAttachmentOptions struct {
	Name string `json:"name"`
}
