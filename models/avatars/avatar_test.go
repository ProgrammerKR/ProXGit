// Copyright 2020 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package avatars_test

import (
	"testing"

<<<<<<< HEAD
	avatars_model "code.gitea.io/gitea/models/avatars"
	"code.gitea.io/gitea/models/db"
	system_model "code.gitea.io/gitea/models/system"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/setting/config"
=======
	avatars_model "code.proxgit.io/proxgit/models/avatars"
	"code.proxgit.io/proxgit/models/db"
	system_model "code.proxgit.io/proxgit/models/system"
	"code.proxgit.io/proxgit/modules/setting"
	"code.proxgit.io/proxgit/modules/setting/config"
>>>>>>> master

	"github.com/stretchr/testify/assert"
)

const gravatarSource = "https://secure.gravatar.com/avatar/"

func disableGravatar(t *testing.T) {
	err := system_model.SetSettings(db.DefaultContext, map[string]string{setting.Config().Picture.EnableFederatedAvatar.DynKey(): "false"})
	assert.NoError(t, err)
	err = system_model.SetSettings(db.DefaultContext, map[string]string{setting.Config().Picture.DisableGravatar.DynKey(): "true"})
	assert.NoError(t, err)
}

func enableGravatar(t *testing.T) {
	err := system_model.SetSettings(db.DefaultContext, map[string]string{setting.Config().Picture.DisableGravatar.DynKey(): "false"})
	assert.NoError(t, err)
	setting.GravatarSource = gravatarSource
}

func TestHashEmail(t *testing.T) {
	assert.Equal(t,
		"d41d8cd98f00b204e9800998ecf8427e",
		avatars_model.HashEmail(""),
	)
	assert.Equal(t,
		"353cbad9b58e69c96154ad99f92bedc7",
<<<<<<< HEAD
		avatars_model.HashEmail("gitea@example.com"),
=======
		avatars_model.HashEmail("proxgit@example.com"),
>>>>>>> master
	)
}

func TestSizedAvatarLink(t *testing.T) {
	setting.AppSubURL = "/testsuburl"

	disableGravatar(t)
	config.GetDynGetter().InvalidateCache()
	assert.Equal(t, "/testsuburl/assets/img/avatar_default.png",
<<<<<<< HEAD
		avatars_model.GenerateEmailAvatarFastLink(db.DefaultContext, "gitea@example.com", 100))
=======
		avatars_model.GenerateEmailAvatarFastLink(db.DefaultContext, "proxgit@example.com", 100))
>>>>>>> master

	enableGravatar(t)
	config.GetDynGetter().InvalidateCache()
	assert.Equal(t,
		"https://secure.gravatar.com/avatar/353cbad9b58e69c96154ad99f92bedc7?d=identicon&s=100",
<<<<<<< HEAD
		avatars_model.GenerateEmailAvatarFastLink(db.DefaultContext, "gitea@example.com", 100),
=======
		avatars_model.GenerateEmailAvatarFastLink(db.DefaultContext, "proxgit@example.com", 100),
>>>>>>> master
	)
}
