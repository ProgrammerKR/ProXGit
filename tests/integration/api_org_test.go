// Copyright 2018 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package integration

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

<<<<<<< HEAD
	auth_model "code.gitea.io/gitea/models/auth"
	"code.gitea.io/gitea/models/db"
	org_model "code.gitea.io/gitea/models/organization"
	"code.gitea.io/gitea/models/perm"
	unit_model "code.gitea.io/gitea/models/unit"
	"code.gitea.io/gitea/models/unittest"
	user_model "code.gitea.io/gitea/models/user"
	"code.gitea.io/gitea/modules/setting"
	api "code.gitea.io/gitea/modules/structs"
	"code.gitea.io/gitea/modules/test"
	"code.gitea.io/gitea/tests"
=======
	auth_model "code.proxgit.io/proxgit/models/auth"
	"code.proxgit.io/proxgit/models/db"
	org_model "code.proxgit.io/proxgit/models/organization"
	"code.proxgit.io/proxgit/models/perm"
	unit_model "code.proxgit.io/proxgit/models/unit"
	"code.proxgit.io/proxgit/models/unittest"
	user_model "code.proxgit.io/proxgit/models/user"
	"code.proxgit.io/proxgit/modules/setting"
	api "code.proxgit.io/proxgit/modules/structs"
	"code.proxgit.io/proxgit/modules/test"
	"code.proxgit.io/proxgit/tests"
>>>>>>> master

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAPIOrgCreateRename(t *testing.T) {
	defer tests.PrepareTestEnv(t)()
	token := getUserToken(t, "user1", auth_model.AccessTokenScopeWriteOrganization)

	org := api.CreateOrgOption{
		UserName:    "user1_org",
		FullName:    "User1's organization",
		Description: "This organization created by user1",
<<<<<<< HEAD
		Website:     "https://try.gitea.io",
=======
		Website:     "https://try.proxgit.io",
>>>>>>> master
		Location:    "Shanghai",
		Visibility:  "limited",
	}
	req := NewRequestWithJSON(t, "POST", "/api/v1/orgs", &org).AddTokenAuth(token)
	resp := MakeRequest(t, req, http.StatusCreated)

	var apiOrg api.Organization
	DecodeJSON(t, resp, &apiOrg)

	assert.Equal(t, org.UserName, apiOrg.Name)
	assert.Equal(t, org.FullName, apiOrg.FullName)
	assert.Equal(t, org.Description, apiOrg.Description)
	assert.Equal(t, org.Website, apiOrg.Website)
	assert.Equal(t, org.Location, apiOrg.Location)
	assert.Equal(t, org.Visibility, apiOrg.Visibility)

	unittest.AssertExistsAndLoadBean(t, &user_model.User{
		Name:      org.UserName,
		LowerName: strings.ToLower(org.UserName),
		FullName:  org.FullName,
	})

	// check org name
	req = NewRequestf(t, "GET", "/api/v1/orgs/%s", org.UserName).AddTokenAuth(token)
	resp = MakeRequest(t, req, http.StatusOK)
	DecodeJSON(t, resp, &apiOrg)
	assert.Equal(t, org.UserName, apiOrg.Name)

	t.Run("CheckPermission", func(t *testing.T) {
		// Check owner team permission
		ownerTeam, _ := org_model.GetOwnerTeam(db.DefaultContext, apiOrg.ID)
		for _, ut := range unit_model.AllRepoUnitTypes {
			up := perm.AccessModeOwner
			if ut == unit_model.TypeExternalTracker || ut == unit_model.TypeExternalWiki {
				up = perm.AccessModeRead
			}
			unittest.AssertExistsAndLoadBean(t, &org_model.TeamUnit{
				OrgID:      apiOrg.ID,
				TeamID:     ownerTeam.ID,
				Type:       ut,
				AccessMode: up,
			})
		}
	})

	t.Run("CheckMembers", func(t *testing.T) {
		req = NewRequestf(t, "GET", "/api/v1/orgs/%s/members", org.UserName).AddTokenAuth(token)
		resp = MakeRequest(t, req, http.StatusOK)

		// user1 on this org is public
		var users []*api.User
		DecodeJSON(t, resp, &users)
		assert.Len(t, users, 1)
		assert.Equal(t, "user1", users[0].UserName)
	})

	t.Run("RenameOrg", func(t *testing.T) {
		req = NewRequestWithJSON(t, "POST", "/api/v1/orgs/user1_org/rename", &api.RenameOrgOption{
			NewName: "renamed_org",
		}).AddTokenAuth(token)
		MakeRequest(t, req, http.StatusNoContent)
		unittest.AssertExistsAndLoadBean(t, &org_model.Organization{Name: "renamed_org"})
		org.UserName = "renamed_org" // update the variable so the following tests could still use it
	})

	t.Run("ListRepos", func(t *testing.T) {
		// FIXME: this test is wrong, there is no repository at all, so the for-loop is empty
		req = NewRequestf(t, "GET", "/api/v1/orgs/%s/repos", org.UserName).AddTokenAuth(token)
		resp = MakeRequest(t, req, http.StatusOK)
		var repos []*api.Repository
		DecodeJSON(t, resp, &repos)
		for _, repo := range repos {
			assert.False(t, repo.Private)
		}
	})
}

func TestAPIOrgGeneral(t *testing.T) {
	defer tests.PrepareTestEnv(t)()
	user1Session := loginUser(t, "user1")
	user1Token := getTokenForLoggedInUser(t, user1Session, auth_model.AccessTokenScopeWriteOrganization)

	t.Run("OrgGetAll", func(t *testing.T) {
		// accessing with a token will return all orgs
		req := NewRequest(t, "GET", "/api/v1/orgs").AddTokenAuth(user1Token)
		resp := MakeRequest(t, req, http.StatusOK)
		var apiOrgList []*api.Organization

		DecodeJSON(t, resp, &apiOrgList)
		assert.Len(t, apiOrgList, 13)
		assert.Equal(t, "Limited Org 36", apiOrgList[1].FullName)
		assert.Equal(t, "limited", apiOrgList[1].Visibility)

		// accessing without a token will return only public orgs
		req = NewRequest(t, "GET", "/api/v1/orgs")
		resp = MakeRequest(t, req, http.StatusOK)

		DecodeJSON(t, resp, &apiOrgList)
		assert.Len(t, apiOrgList, 9)
		assert.Equal(t, "org 17", apiOrgList[0].FullName)
		assert.Equal(t, "public", apiOrgList[0].Visibility)
	})

	t.Run("OrgEdit", func(t *testing.T) {
		org := api.EditOrgOption{
			FullName:    "Org3 organization new full name",
			Description: "A new description",
<<<<<<< HEAD
			Website:     "https://try.gitea.io/new",
=======
			Website:     "https://try.proxgit.io/new",
>>>>>>> master
			Location:    "Beijing",
			Visibility:  "private",
		}
		req := NewRequestWithJSON(t, "PATCH", "/api/v1/orgs/org3", &org).AddTokenAuth(user1Token)
		resp := MakeRequest(t, req, http.StatusOK)

		var apiOrg api.Organization
		DecodeJSON(t, resp, &apiOrg)

		assert.Equal(t, "org3", apiOrg.Name)
		assert.Equal(t, org.FullName, apiOrg.FullName)
		assert.Equal(t, org.Description, apiOrg.Description)
		assert.Equal(t, org.Website, apiOrg.Website)
		assert.Equal(t, org.Location, apiOrg.Location)
		assert.Equal(t, org.Visibility, apiOrg.Visibility)
	})

	t.Run("OrgEditBadVisibility", func(t *testing.T) {
		org := api.EditOrgOption{
			FullName:    "Org3 organization new full name",
			Description: "A new description",
<<<<<<< HEAD
			Website:     "https://try.gitea.io/new",
=======
			Website:     "https://try.proxgit.io/new",
>>>>>>> master
			Location:    "Beijing",
			Visibility:  "badvisibility",
		}
		req := NewRequestWithJSON(t, "PATCH", "/api/v1/orgs/org3", &org).AddTokenAuth(user1Token)
		MakeRequest(t, req, http.StatusUnprocessableEntity)
	})

	t.Run("OrgDeny", func(t *testing.T) {
		defer test.MockVariableValue(&setting.Service.RequireSignInViewStrict, true)()

		orgName := "user1_org"
		req := NewRequestf(t, "GET", "/api/v1/orgs/%s", orgName)
		MakeRequest(t, req, http.StatusNotFound)

		req = NewRequestf(t, "GET", "/api/v1/orgs/%s/repos", orgName)
		MakeRequest(t, req, http.StatusNotFound)

		req = NewRequestf(t, "GET", "/api/v1/orgs/%s/members", orgName)
		MakeRequest(t, req, http.StatusNotFound)
	})

	t.Run("OrgSearchEmptyTeam", func(t *testing.T) {
		orgName := "org_with_empty_team"
		// create org
		req := NewRequestWithJSON(t, "POST", "/api/v1/orgs", &api.CreateOrgOption{
			UserName: orgName,
		}).AddTokenAuth(user1Token)
		MakeRequest(t, req, http.StatusCreated)

		// create team with no member
		req = NewRequestWithJSON(t, "POST", fmt.Sprintf("/api/v1/orgs/%s/teams", orgName), &api.CreateTeamOption{
			Name:                    "Empty",
			IncludesAllRepositories: true,
			Permission:              "read",
			Units:                   []string{"repo.code", "repo.issues", "repo.ext_issues", "repo.wiki", "repo.pulls"},
		}).AddTokenAuth(user1Token)
		MakeRequest(t, req, http.StatusCreated)

		// case-insensitive search for teams that have no members
		req = NewRequest(t, "GET", fmt.Sprintf("/api/v1/orgs/%s/teams/search?q=%s", orgName, "empty")).
			AddTokenAuth(user1Token)
		resp := MakeRequest(t, req, http.StatusOK)
		data := struct {
			Ok   bool
			Data []*api.Team
		}{}
		DecodeJSON(t, resp, &data)
		assert.True(t, data.Ok)
		if assert.Len(t, data.Data, 1) {
			assert.Equal(t, "Empty", data.Data[0].Name)
		}
	})

	t.Run("User2ChangeStatus", func(t *testing.T) {
		user2Session := loginUser(t, "user2")
		user2Token := getTokenForLoggedInUser(t, user2Session, auth_model.AccessTokenScopeWriteOrganization)

		req := NewRequest(t, "PUT", "/api/v1/orgs/org3/public_members/user2").AddTokenAuth(user2Token)
		MakeRequest(t, req, http.StatusNoContent)
		req = NewRequest(t, "DELETE", "/api/v1/orgs/org3/public_members/user2").AddTokenAuth(user2Token)
		MakeRequest(t, req, http.StatusNoContent)

		// non admin but org owner could also change other member's status
		user2 := unittest.AssertExistsAndLoadBean(t, &user_model.User{Name: "user2"})
		require.False(t, user2.IsAdmin)
		req = NewRequest(t, "PUT", "/api/v1/orgs/org3/public_members/user1").AddTokenAuth(user2Token)
		MakeRequest(t, req, http.StatusNoContent)
		req = NewRequest(t, "DELETE", "/api/v1/orgs/org3/public_members/user1").AddTokenAuth(user2Token)
		MakeRequest(t, req, http.StatusNoContent)
	})

	t.Run("User4ChangeStatus", func(t *testing.T) {
		user4Session := loginUser(t, "user4")
		user4Token := getTokenForLoggedInUser(t, user4Session, auth_model.AccessTokenScopeWriteOrganization)

		// user4 is a normal team member, they could change their own status
		req := NewRequest(t, "PUT", "/api/v1/orgs/org3/public_members/user4").AddTokenAuth(user4Token)
		MakeRequest(t, req, http.StatusNoContent)
		req = NewRequest(t, "DELETE", "/api/v1/orgs/org3/public_members/user4").AddTokenAuth(user4Token)
		MakeRequest(t, req, http.StatusNoContent)
		req = NewRequest(t, "PUT", "/api/v1/orgs/org3/public_members/user1").AddTokenAuth(user4Token)
		MakeRequest(t, req, http.StatusForbidden)
		req = NewRequest(t, "DELETE", "/api/v1/orgs/org3/public_members/user1").AddTokenAuth(user4Token)
		MakeRequest(t, req, http.StatusForbidden)
	})
}
