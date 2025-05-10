// Copyright 2014 The Gogs Authors. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package main

import (
	"os"
	"runtime"
	"strings"
	"time"

	"code.proxgit.io/proxgit/cmd"
	"code.proxgit.io/proxgit/modules/log"
	"code.proxgit.io/proxgit/modules/setting"

	// register supported doc types
	_ "code.proxgit.io/proxgit/modules/markup/asciicast"
	_ "code.proxgit.io/proxgit/modules/markup/console"
	_ "code.proxgit.io/proxgit/modules/markup/csv"
	_ "code.proxgit.io/proxgit/modules/markup/markdown"
	_ "code.proxgit.io/proxgit/modules/markup/orgmode"

	"github.com/urfave/cli/v2"
)

// these flags will be set by the build flags
var (
	Version     = "development" // program version for this build
	Tags        = ""            // the Golang build tags
	MakeVersion = ""            // "make" program version if built with make
)

func init() {
	setting.AppVer = Version
	setting.AppBuiltWith = formatBuiltWith()
	setting.AppStartTime = time.Now().UTC()
}

func main() {
	cli.OsExiter = func(code int) {
		log.GetManager().Close()
		os.Exit(code)
	}
	app := cmd.NewMainApp(cmd.AppVersion{Version: Version, Extra: formatBuiltWith()})
	_ = cmd.RunMainApp(app, os.Args...) // all errors should have been handled by the RunMainApp
	log.GetManager().Close()
}

func formatBuiltWith() string {
	version := runtime.Version()
	if len(MakeVersion) > 0 {
		version = MakeVersion + ", " + runtime.Version()
	}
	if len(Tags) == 0 {
		return " built with " + version
	}

	return " built with " + version + " : " + strings.ReplaceAll(Tags, " ", ", ")
}
