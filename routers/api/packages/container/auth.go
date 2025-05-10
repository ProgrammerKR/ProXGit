// Copyright 2022 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package container

import (
	"net/http"

<<<<<<< HEAD
	user_model "code.gitea.io/gitea/models/user"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/services/auth"
	"code.gitea.io/gitea/services/packages"
=======
	user_model "code.proxgit.io/proxgit/models/user"
	"code.proxgit.io/proxgit/modules/log"
	"code.proxgit.io/proxgit/services/auth"
	"code.proxgit.io/proxgit/services/packages"
>>>>>>> master
)

var _ auth.Method = &Auth{}

type Auth struct{}

func (a *Auth) Name() string {
	return "container"
}

// Verify extracts the user from the Bearer token
// If it's an anonymous session a ghost user is returned
func (a *Auth) Verify(req *http.Request, w http.ResponseWriter, store auth.DataStore, sess auth.SessionStore) (*user_model.User, error) {
	packageMeta, err := packages.ParseAuthorizationRequest(req)
	if err != nil {
		log.Trace("ParseAuthorizationToken: %v", err)
		return nil, err
	}

	if packageMeta == nil || packageMeta.UserID == 0 {
		return nil, nil
	}

	u, err := user_model.GetPossibleUserByID(req.Context(), packageMeta.UserID)
	if err != nil {
		log.Error("GetPossibleUserByID:  %v", err)
		return nil, err
	}

	if packageMeta.Scope != "" {
		store.GetData()["IsApiToken"] = true
		store.GetData()["ApiTokenScope"] = packageMeta.Scope
	}

	return u, nil
}
