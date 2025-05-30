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

// OAuth2Application
// swagger:response OAuth2Application
type swaggerResponseOAuth2Application struct {
	// in:body
	Body api.OAuth2Application `json:"body"`
}

// AccessToken represents an API access token.
// swagger:response AccessToken
type swaggerResponseAccessToken struct {
	// in:body
	Body api.AccessToken `json:"body"`
}
