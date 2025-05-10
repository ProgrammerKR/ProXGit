// Copyright 2016 The Gogs Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package context

<<<<<<< HEAD
import "code.gitea.io/gitea/models/organization"
=======
import "code.proxgit.io/proxgit/models/organization"
>>>>>>> master

// APIOrganization contains organization and team
type APIOrganization struct {
	Organization *organization.Organization
	Team         *organization.Team
}
