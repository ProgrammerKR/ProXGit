// Copyright 2018 The Gogs Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package forms

import (
	"testing"

<<<<<<< HEAD
	"code.gitea.io/gitea/modules/setting"
=======
	"code.proxgit.io/proxgit/modules/setting"
>>>>>>> master

	"github.com/gobwas/glob"
	"github.com/stretchr/testify/assert"
)

func TestRegisterForm_IsDomainAllowed_Empty(t *testing.T) {
	oldService := setting.Service
	defer func() {
		setting.Service = oldService
	}()

	setting.Service.EmailDomainAllowList = nil

	form := RegisterForm{}

	assert.True(t, form.IsEmailDomainAllowed())
}

func TestRegisterForm_IsDomainAllowed_InvalidEmail(t *testing.T) {
	oldService := setting.Service
	defer func() {
		setting.Service = oldService
	}()

<<<<<<< HEAD
	setting.Service.EmailDomainAllowList = []glob.Glob{glob.MustCompile("gitea.io")}
=======
	setting.Service.EmailDomainAllowList = []glob.Glob{glob.MustCompile("proxgit.io")}
>>>>>>> master

	tt := []struct {
		email string
	}{
		{"invalid-email"},
<<<<<<< HEAD
		{"gitea.io"},
=======
		{"proxgit.io"},
>>>>>>> master
	}

	for _, v := range tt {
		form := RegisterForm{Email: v.email}

		assert.False(t, form.IsEmailDomainAllowed())
	}
}

func TestRegisterForm_IsDomainAllowed_AllowedEmail(t *testing.T) {
	oldService := setting.Service
	defer func() {
		setting.Service = oldService
	}()

<<<<<<< HEAD
	setting.Service.EmailDomainAllowList = []glob.Glob{glob.MustCompile("gitea.io"), glob.MustCompile("*.allow")}
=======
	setting.Service.EmailDomainAllowList = []glob.Glob{glob.MustCompile("proxgit.io"), glob.MustCompile("*.allow")}
>>>>>>> master

	tt := []struct {
		email string
		valid bool
	}{
<<<<<<< HEAD
		{"security@gitea.io", true},
=======
		{"security@proxgit.io", true},
>>>>>>> master
		{"security@gITea.io", true},
		{"invalid", false},
		{"seee@example.com", false},

		{"user@my.allow", true},
		{"user@my.allow1", false},
	}

	for _, v := range tt {
		form := RegisterForm{Email: v.email}

		assert.Equal(t, v.valid, form.IsEmailDomainAllowed())
	}
}

func TestRegisterForm_IsDomainAllowed_BlockedEmail(t *testing.T) {
	oldService := setting.Service
	defer func() {
		setting.Service = oldService
	}()

	setting.Service.EmailDomainAllowList = nil
<<<<<<< HEAD
	setting.Service.EmailDomainBlockList = []glob.Glob{glob.MustCompile("gitea.io"), glob.MustCompile("*.block")}
=======
	setting.Service.EmailDomainBlockList = []glob.Glob{glob.MustCompile("proxgit.io"), glob.MustCompile("*.block")}
>>>>>>> master

	tt := []struct {
		email string
		valid bool
	}{
<<<<<<< HEAD
		{"security@gitea.io", false},
		{"security@gitea.example", true},
=======
		{"security@proxgit.io", false},
		{"security@proxgit.example", true},
>>>>>>> master
		{"invalid", true},

		{"user@my.block", false},
		{"user@my.block1", true},
	}

	for _, v := range tt {
		form := RegisterForm{Email: v.email}

		assert.Equal(t, v.valid, form.IsEmailDomainAllowed())
	}
}
