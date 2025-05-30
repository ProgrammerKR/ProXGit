// Copyright 2022 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package markdown

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestRenderConfig_UnmarshalYAML(t *testing.T) {
	tests := []struct {
		name     string
		expected *RenderConfig
		args     string
	}{
		{
			"empty", &RenderConfig{
				Meta: "table",
				Lang: "",
			}, "",
		},
		{
			"lang", &RenderConfig{
				Meta: "table",
				Lang: "test",
			}, "lang: test",
		},
		{
			"metatable", &RenderConfig{
				Meta: "table",
				Lang: "",
<<<<<<< HEAD
			}, "gitea: table",
=======
			}, "proxgit: table",
>>>>>>> master
		},
		{
			"metanone", &RenderConfig{
				Meta: "none",
				Lang: "",
<<<<<<< HEAD
			}, "gitea: none",
=======
			}, "proxgit: none",
>>>>>>> master
		},
		{
			"metadetails", &RenderConfig{
				Meta: "details",
				Lang: "",
<<<<<<< HEAD
			}, "gitea: details",
=======
			}, "proxgit: details",
>>>>>>> master
		},
		{
			"metawrong", &RenderConfig{
				Meta: "details",
				Lang: "",
<<<<<<< HEAD
			}, "gitea: wrong",
=======
			}, "proxgit: wrong",
>>>>>>> master
		},
		{
			"toc", &RenderConfig{
				TOC:  "true",
				Meta: "table",
				Lang: "",
			}, "include_toc: true",
		},
		{
			"tocfalse", &RenderConfig{
				TOC:  "false",
				Meta: "table",
				Lang: "",
			}, "include_toc: false",
		},
		{
			"toclang", &RenderConfig{
				Meta: "table",
				TOC:  "true",
				Lang: "testlang",
			}, `
				include_toc: true
				lang: testlang
				`,
		},
		{
			"complexlang", &RenderConfig{
				Meta: "table",
				Lang: "testlang",
			}, `
<<<<<<< HEAD
				gitea:
=======
				proxgit:
>>>>>>> master
					lang: testlang
				`,
		},
		{
			"complexlang2", &RenderConfig{
				Meta: "table",
				Lang: "testlang",
			}, `
	lang: notright
<<<<<<< HEAD
	gitea:
=======
	proxgit:
>>>>>>> master
		lang: testlang
`,
		},
		{
			"complexlang", &RenderConfig{
				Meta: "table",
				Lang: "testlang",
			}, `
<<<<<<< HEAD
	gitea:
=======
	proxgit:
>>>>>>> master
		lang: testlang
`,
		},
		{
			"complex2", &RenderConfig{
				Lang: "two",
				Meta: "table",
				TOC:  "true",
			}, `
	lang: one
	include_toc: true
<<<<<<< HEAD
	gitea:
=======
	proxgit:
>>>>>>> master
		details_icon: smiley
		meta: table
		include_toc: true
		lang: two
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := &RenderConfig{
				Meta: "table",
				Lang: "",
			}
			err := yaml.Unmarshal([]byte(strings.ReplaceAll(tt.args, "\t", "    ")), got)
			require.NoError(t, err)

			assert.Equal(t, tt.expected.Meta, got.Meta)
			assert.Equal(t, tt.expected.Lang, got.Lang)
			assert.Equal(t, tt.expected.TOC, got.TOC)
		})
	}
}
