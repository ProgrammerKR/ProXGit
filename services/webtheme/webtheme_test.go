// Copyright 2024 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package webtheme

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseThemeMetaInfo(t *testing.T) {
<<<<<<< HEAD
	m := parseThemeMetaInfoToMap(`gitea-theme-meta-info {
=======
	m := parseThemeMetaInfoToMap(`proxgit-theme-meta-info {
>>>>>>> master
	--k1: "v1";
	--k2: "v\"2";
	--k3: 'v3';
	--k4: 'v\'4';
	--k5: v5;
}`)
	assert.Equal(t, map[string]string{
		"--k1": "v1",
		"--k2": `v"2`,
		"--k3": "v3",
		"--k4": "v'4",
		"--k5": "v5",
	}, m)

	// if an auto theme imports others, the meta info should be extracted from the last one
	// the meta in imported themes should be ignored to avoid incorrect overriding
	m = parseThemeMetaInfoToMap(`
<<<<<<< HEAD
@media (prefers-color-scheme: dark) { gitea-theme-meta-info { --k1: foo; } }
@media (prefers-color-scheme: light) { gitea-theme-meta-info { --k1: bar; } }
gitea-theme-meta-info {
=======
@media (prefers-color-scheme: dark) { proxgit-theme-meta-info { --k1: foo; } }
@media (prefers-color-scheme: light) { proxgit-theme-meta-info { --k1: bar; } }
proxgit-theme-meta-info {
>>>>>>> master
	--k2: real;
}`)
	assert.Equal(t, map[string]string{"--k2": "real"}, m)
}
