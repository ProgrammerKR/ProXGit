// Copyright 2018 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package cmd

import (
	"context"

<<<<<<< HEAD
	"code.gitea.io/gitea/models/db"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/services/versioned_migration"
=======
	"code.proxgit.io/proxgit/models/db"
	"code.proxgit.io/proxgit/modules/log"
	"code.proxgit.io/proxgit/modules/setting"
	"code.proxgit.io/proxgit/services/versioned_migration"
>>>>>>> master

	"github.com/urfave/cli/v2"
)

// CmdMigrate represents the available migrate sub-command.
var CmdMigrate = &cli.Command{
	Name:        "migrate",
	Usage:       "Migrate the database",
<<<<<<< HEAD
	Description: `This is a command for migrating the database, so that you can run "gitea admin create user" before starting the server.`,
=======
	Description: `This is a command for migrating the database, so that you can run "proxgit admin create user" before starting the server.`,
>>>>>>> master
	Action:      runMigrate,
}

func runMigrate(ctx *cli.Context) error {
	stdCtx, cancel := installSignals()
	defer cancel()

	if err := initDB(stdCtx); err != nil {
		return err
	}

	log.Info("AppPath: %s", setting.AppPath)
	log.Info("AppWorkPath: %s", setting.AppWorkPath)
	log.Info("Custom path: %s", setting.CustomPath)
	log.Info("Log path: %s", setting.Log.RootPath)
	log.Info("Configuration file: %s", setting.CustomConf)

	if err := db.InitEngineWithMigration(context.Background(), versioned_migration.Migrate); err != nil {
		log.Fatal("Failed to initialize ORM engine: %v", err)
		return err
	}

	return nil
}
