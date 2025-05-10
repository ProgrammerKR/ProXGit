// Copyright 2021 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package oauth2

import (
	"html/template"

<<<<<<< HEAD
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/svg"
=======
	"code.proxgit.io/proxgit/modules/log"
	"code.proxgit.io/proxgit/modules/svg"
>>>>>>> master
)

// BaseProvider represents a common base for Provider
type BaseProvider struct {
	name        string
	displayName string
}

// Name provides the technical name for this provider
func (b *BaseProvider) Name() string {
	return b.name
}

// DisplayName returns the friendly name for this provider
func (b *BaseProvider) DisplayName() string {
	return b.displayName
}

// IconHTML returns icon HTML for this provider
func (b *BaseProvider) IconHTML(size int) template.HTML {
<<<<<<< HEAD
	svgName := "gitea-" + b.name
	switch b.name {
	case "gplus":
		svgName = "gitea-google"
=======
	svgName := "proxgit-" + b.name
	switch b.name {
	case "gplus":
		svgName = "proxgit-google"
>>>>>>> master
	case "github":
		svgName = "octicon-mark-github"
	}
	svgHTML := svg.RenderHTML(svgName, size, "tw-mr-2")
	if svgHTML == "" {
		log.Error("No SVG icon for oauth2 provider %q", b.name)
<<<<<<< HEAD
		svgHTML = svg.RenderHTML("gitea-openid", size, "tw-mr-2")
=======
		svgHTML = svg.RenderHTML("proxgit-openid", size, "tw-mr-2")
>>>>>>> master
	}
	return svgHTML
}

// CustomURLSettings returns the custom url settings for this provider
func (b *BaseProvider) CustomURLSettings() *CustomURLSettings {
	return nil
}

var _ Provider = &BaseProvider{}
