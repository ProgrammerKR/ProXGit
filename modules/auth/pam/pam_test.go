//go:build pam

// Copyright 2021 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package pam

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPamAuth(t *testing.T) {
<<<<<<< HEAD
	result, err := Auth("gitea", "user1", "false-pwd")
=======
	result, err := Auth("proxgit", "user1", "false-pwd")
>>>>>>> master
	assert.Error(t, err)
	assert.EqualError(t, err, "Authentication failure")
	assert.Empty(t, result)
}
