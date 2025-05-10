// Copyright 2022 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

//nolint:forbidigo
package base

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"

<<<<<<< HEAD
	"code.gitea.io/gitea/models/unittest"
	"code.gitea.io/gitea/modules/git"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/tempdir"
	"code.gitea.io/gitea/modules/test"
	"code.gitea.io/gitea/modules/testlogger"
=======
	"code.proxgit.io/proxgit/models/unittest"
	"code.proxgit.io/proxgit/modules/git"
	"code.proxgit.io/proxgit/modules/setting"
	"code.proxgit.io/proxgit/modules/tempdir"
	"code.proxgit.io/proxgit/modules/test"
	"code.proxgit.io/proxgit/modules/testlogger"
>>>>>>> master

	"github.com/stretchr/testify/require"
	"xorm.io/xorm"
)

// FIXME: this file shouldn't be in a normal package, it should only be compiled for tests

// PrepareTestEnv prepares the test environment and reset the database. The skip parameter should usually be 0.
// Provide models to be sync'd with the database - in particular any models you expect fixtures to be loaded from.
//
// fixtures in `models/migrations/fixtures/<TestName>` will be loaded automatically
func PrepareTestEnv(t *testing.T, skip int, syncModels ...any) (*xorm.Engine, func()) {
	t.Helper()
	ourSkip := 2
	ourSkip += skip
	deferFn := testlogger.PrintCurrentTest(t, ourSkip)
<<<<<<< HEAD
	require.NoError(t, unittest.SyncDirs(filepath.Join(filepath.Dir(setting.AppPath), "tests/gitea-repositories-meta"), setting.RepoRootPath))
=======
	require.NoError(t, unittest.SyncDirs(filepath.Join(filepath.Dir(setting.AppPath), "tests/proxgit-repositories-meta"), setting.RepoRootPath))
>>>>>>> master

	if err := deleteDB(); err != nil {
		t.Fatalf("unable to reset database: %v", err)
		return nil, deferFn
	}

	x, err := newXORMEngine()
	require.NoError(t, err)
	if x != nil {
		oldDefer := deferFn
		deferFn = func() {
			oldDefer()
			if err := x.Close(); err != nil {
				t.Errorf("error during close: %v", err)
			}
			if err := deleteDB(); err != nil {
				t.Errorf("unable to reset database: %v", err)
			}
		}
	}
	if err != nil {
		return x, deferFn
	}

	if len(syncModels) > 0 {
		if err := x.Sync(syncModels...); err != nil {
			t.Errorf("error during sync: %v", err)
			return x, deferFn
		}
	}

	fixturesDir := filepath.Join(filepath.Dir(setting.AppPath), "models", "migrations", "fixtures", t.Name())

	if _, err := os.Stat(fixturesDir); err == nil {
		t.Logf("initializing fixtures from: %s", fixturesDir)
		if err := unittest.InitFixtures(
			unittest.FixturesOptions{
				Dir: fixturesDir,
			}, x); err != nil {
			t.Errorf("error whilst initializing fixtures from %s: %v", fixturesDir, err)
			return x, deferFn
		}
		if err := unittest.LoadFixtures(); err != nil {
			t.Errorf("error whilst loading fixtures from %s: %v", fixturesDir, err)
			return x, deferFn
		}
	} else if !os.IsNotExist(err) {
		t.Errorf("unexpected error whilst checking for existence of fixtures: %v", err)
	} else {
		t.Logf("no fixtures found in: %s", fixturesDir)
	}

	return x, deferFn
}

func MainTest(m *testing.M) {
	testlogger.Init()

<<<<<<< HEAD
	giteaRoot := test.SetupGiteaRoot()
	giteaBinary := "gitea"
	if runtime.GOOS == "windows" {
		giteaBinary += ".exe"
	}
	setting.AppPath = filepath.Join(giteaRoot, giteaBinary)
	if _, err := os.Stat(setting.AppPath); err != nil {
		testlogger.Fatalf("Could not find gitea binary at %s\n", setting.AppPath)
	}

	giteaConf := os.Getenv("GITEA_CONF")
	if giteaConf == "" {
		giteaConf = filepath.Join(filepath.Dir(setting.AppPath), "tests/sqlite.ini")
		fmt.Printf("Environment variable $GITEA_CONF not set - defaulting to %s\n", giteaConf)
	}

	if !filepath.IsAbs(giteaConf) {
		setting.CustomConf = filepath.Join(giteaRoot, giteaConf)
	} else {
		setting.CustomConf = giteaConf
	}

	tmpDataPath, cleanup, err := tempdir.OsTempDir("gitea-test").MkdirTempRandom("data")
=======
	proxgitRoot := test.SetupGiteaRoot()
	proxgitBinary := "proxgit"
	if runtime.GOOS == "windows" {
		proxgitBinary += ".exe"
	}
	setting.AppPath = filepath.Join(proxgitRoot, proxgitBinary)
	if _, err := os.Stat(setting.AppPath); err != nil {
		testlogger.Fatalf("Could not find proxgit binary at %s\n", setting.AppPath)
	}

	proxgitConf := os.Getenv("GITEA_CONF")
	if proxgitConf == "" {
		proxgitConf = filepath.Join(filepath.Dir(setting.AppPath), "tests/sqlite.ini")
		fmt.Printf("Environment variable $GITEA_CONF not set - defaulting to %s\n", proxgitConf)
	}

	if !filepath.IsAbs(proxgitConf) {
		setting.CustomConf = filepath.Join(proxgitRoot, proxgitConf)
	} else {
		setting.CustomConf = proxgitConf
	}

	tmpDataPath, cleanup, err := tempdir.OsTempDir("proxgit-test").MkdirTempRandom("data")
>>>>>>> master
	if err != nil {
		testlogger.Fatalf("Unable to create temporary data path %v\n", err)
	}
	defer cleanup()

	setting.CustomPath = filepath.Join(setting.AppWorkPath, "custom")
	setting.AppDataPath = tmpDataPath

	unittest.InitSettingsForTesting()
	if err = git.InitFull(context.Background()); err != nil {
		testlogger.Fatalf("Unable to InitFull: %v\n", err)
	}
	setting.LoadDBSetting()
	setting.InitLoggersForTest()

	exitStatus := m.Run()

	if err := removeAllWithRetry(setting.RepoRootPath); err != nil {
		fmt.Fprintf(os.Stderr, "os.RemoveAll: %v\n", err)
	}
	os.Exit(exitStatus)
}
