// Copyright 2023 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package setting

<<<<<<< HEAD
import "code.gitea.io/gitea/modules/log"
=======
import "code.proxgit.io/proxgit/modules/log"
>>>>>>> master

type OtherConfig struct {
	ShowFooterVersion          bool
	ShowFooterTemplateLoadTime bool
	ShowFooterPoweredBy        bool
	EnableFeed                 bool
	EnableSitemap              bool
}

var Other = OtherConfig{
	ShowFooterVersion:          true,
	ShowFooterTemplateLoadTime: true,
	ShowFooterPoweredBy:        true,
	EnableSitemap:              true,
	EnableFeed:                 true,
}

func loadOtherFrom(rootCfg ConfigProvider) {
	sec := rootCfg.Section("other")
	if err := sec.MapTo(&Other); err != nil {
		log.Fatal("Failed to map [other] settings: %v", err)
	}
}
