// Copyright 2017 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package swagger

import (
<<<<<<< HEAD
	api "code.gitea.io/gitea/modules/structs"
=======
	api "code.proxgit.io/proxgit/modules/structs"
>>>>>>> master
)

// Organization
// swagger:response Organization
type swaggerResponseOrganization struct {
	// in:body
	Body api.Organization `json:"body"`
}

// OrganizationList
// swagger:response OrganizationList
type swaggerResponseOrganizationList struct {
	// in:body
	Body []api.Organization `json:"body"`
}

// Team
// swagger:response Team
type swaggerResponseTeam struct {
	// in:body
	Body api.Team `json:"body"`
}

// TeamList
// swagger:response TeamList
type swaggerResponseTeamList struct {
	// in:body
	Body []api.Team `json:"body"`
}

// OrganizationPermissions
// swagger:response OrganizationPermissions
type swaggerResponseOrganizationPermissions struct {
	// in:body
	Body api.OrganizationPermissions `json:"body"`
}
