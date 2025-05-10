// Copyright 2021 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package unittest

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

<<<<<<< HEAD
	"code.gitea.io/gitea/models/db"
	"code.gitea.io/gitea/models/system"
	"code.gitea.io/gitea/modules/auth/password/hash"
	"code.gitea.io/gitea/modules/cache"
	"code.gitea.io/gitea/modules/git"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/setting/config"
	"code.gitea.io/gitea/modules/storage"
	"code.gitea.io/gitea/modules/tempdir"
	"code.gitea.io/gitea/modules/test"
	"code.gitea.io/gitea/modules/util"
=======
	"code.proxgit.io/proxgit/models/db"
	"code.proxgit.io/proxgit/models/system"
	"code.proxgit.io/proxgit/modules/auth/password/hash"
	"code.proxgit.io/proxgit/modules/cache"
	"code.proxgit.io/proxgit/modules/git"
	"code.proxgit.io/proxgit/modules/log"
	"code.proxgit.io/proxgit/modules/setting"
	"code.proxgit.io/proxgit/modules/setting/config"
	"code.proxgit.io/proxgit/modules/storage"
	"code.proxgit.io/proxgit/modules/tempdir"
	"code.proxgit.io/proxgit/modules/test"
	"code.proxgit.io/proxgit/modules/util"
>>>>>>> master

	"github.com/stretchr/testify/assert"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

<<<<<<< HEAD
var giteaRoot string
=======
var proxgitRoot string
>>>>>>> master

func fatalTestError(fmtStr string, args ...any) {
	_, _ = fmt.Fprintf(os.Stderr, fmtStr, args...)
	os.Exit(1)
}

// InitSettingsForTesting initializes config provider and load common settings for tests
func InitSettingsForTesting() {
	setting.IsInTesting = true
	log.OsExiter = func(code int) {
		if code != 0 {
			// non-zero exit code (log.Fatal) shouldn't occur during testing, if it happens, show a full stacktrace for more details
			panic(fmt.Errorf("non-zero exit code during testing: %d", code))
		}
		os.Exit(0)
	}
	if setting.CustomConf == "" {
		setting.CustomConf = filepath.Join(setting.CustomPath, "conf/app-unittest-tmp.ini")
		_ = os.Remove(setting.CustomConf)
	}
	setting.InitCfgProvider(setting.CustomConf)
	setting.LoadCommonSettings()

	if err := setting.PrepareAppDataPath(); err != nil {
		log.Fatal("Can not prepare APP_DATA_PATH: %v", err)
	}
	// register the dummy hash algorithm function used in the test fixtures
	_ = hash.Register("dummy", hash.NewDummyHasher)

	setting.PasswordHashAlgo, _ = hash.SetDefaultPasswordHashAlgorithm("dummy")
	setting.InitGiteaEnvVarsForTesting()
}

// TestOptions represents test options
type TestOptions struct {
	FixtureFiles []string
	SetUp        func() error // SetUp will be executed before all tests in this package
	TearDown     func() error // TearDown will be executed after all tests in this package
}

// MainTest a reusable TestMain(..) function for unit tests that need to use a
// test database. Creates the test database, and sets necessary settings.
func MainTest(m *testing.M, testOptsArg ...*TestOptions) {
	testOpts := util.OptionalArg(testOptsArg, &TestOptions{})
<<<<<<< HEAD
	giteaRoot = test.SetupGiteaRoot()
	setting.CustomPath = filepath.Join(giteaRoot, "custom")
	InitSettingsForTesting()

	fixturesOpts := FixturesOptions{Dir: filepath.Join(giteaRoot, "models", "fixtures"), Files: testOpts.FixtureFiles}
=======
	proxgitRoot = test.SetupGiteaRoot()
	setting.CustomPath = filepath.Join(proxgitRoot, "custom")
	InitSettingsForTesting()

	fixturesOpts := FixturesOptions{Dir: filepath.Join(proxgitRoot, "models", "fixtures"), Files: testOpts.FixtureFiles}
>>>>>>> master
	if err := CreateTestEngine(fixturesOpts); err != nil {
		fatalTestError("Error creating test engine: %v\n", err)
	}

	setting.IsInTesting = true
<<<<<<< HEAD
	setting.AppURL = "https://try.gitea.io/"
	setting.Domain = "try.gitea.io"
=======
	setting.AppURL = "https://try.proxgit.io/"
	setting.Domain = "try.proxgit.io"
>>>>>>> master
	setting.RunUser = "runuser"
	setting.SSH.User = "sshuser"
	setting.SSH.BuiltinServerUser = "builtinuser"
	setting.SSH.Port = 3000
<<<<<<< HEAD
	setting.SSH.Domain = "try.gitea.io"
	setting.Database.Type = "sqlite3"
	setting.Repository.DefaultBranch = "master" // many test code still assume that default branch is called "master"
	repoRootPath, cleanup1, err := tempdir.OsTempDir("gitea-test").MkdirTempRandom("repos")
=======
	setting.SSH.Domain = "try.proxgit.io"
	setting.Database.Type = "sqlite3"
	setting.Repository.DefaultBranch = "master" // many test code still assume that default branch is called "master"
	repoRootPath, cleanup1, err := tempdir.OsTempDir("proxgit-test").MkdirTempRandom("repos")
>>>>>>> master
	if err != nil {
		fatalTestError("TempDir: %v\n", err)
	}
	defer cleanup1()

	setting.RepoRootPath = repoRootPath
<<<<<<< HEAD
	appDataPath, cleanup2, err := tempdir.OsTempDir("gitea-test").MkdirTempRandom("appdata")
=======
	appDataPath, cleanup2, err := tempdir.OsTempDir("proxgit-test").MkdirTempRandom("appdata")
>>>>>>> master
	if err != nil {
		fatalTestError("TempDir: %v\n", err)
	}
	defer cleanup2()

	setting.AppDataPath = appDataPath
<<<<<<< HEAD
	setting.AppWorkPath = giteaRoot
	setting.StaticRootPath = giteaRoot
=======
	setting.AppWorkPath = proxgitRoot
	setting.StaticRootPath = proxgitRoot
>>>>>>> master
	setting.GravatarSource = "https://secure.gravatar.com/avatar/"

	setting.Attachment.Storage.Path = filepath.Join(setting.AppDataPath, "attachments")

	setting.LFS.Storage.Path = filepath.Join(setting.AppDataPath, "lfs")

	setting.Avatar.Storage.Path = filepath.Join(setting.AppDataPath, "avatars")

	setting.RepoAvatar.Storage.Path = filepath.Join(setting.AppDataPath, "repo-avatars")

	setting.RepoArchive.Storage.Path = filepath.Join(setting.AppDataPath, "repo-archive")

	setting.Packages.Storage.Path = filepath.Join(setting.AppDataPath, "packages")

	setting.Actions.LogStorage.Path = filepath.Join(setting.AppDataPath, "actions_log")

	setting.Git.HomePath = filepath.Join(setting.AppDataPath, "home")

	setting.IncomingEmail.ReplyToAddress = "incoming+%{token}@localhost"

	config.SetDynGetter(system.NewDatabaseDynKeyGetter())

	if err = cache.Init(); err != nil {
		fatalTestError("cache.Init: %v\n", err)
	}
	if err = storage.Init(); err != nil {
		fatalTestError("storage.Init: %v\n", err)
	}
<<<<<<< HEAD
	if err = SyncDirs(filepath.Join(giteaRoot, "tests", "gitea-repositories-meta"), setting.RepoRootPath); err != nil {
=======
	if err = SyncDirs(filepath.Join(proxgitRoot, "tests", "proxgit-repositories-meta"), setting.RepoRootPath); err != nil {
>>>>>>> master
		fatalTestError("util.SyncDirs: %v\n", err)
	}

	if err = git.InitFull(context.Background()); err != nil {
		fatalTestError("git.Init: %v\n", err)
	}

	if testOpts.SetUp != nil {
		if err := testOpts.SetUp(); err != nil {
			fatalTestError("set up failed: %v\n", err)
		}
	}

	exitStatus := m.Run()

	if testOpts.TearDown != nil {
		if err := testOpts.TearDown(); err != nil {
			fatalTestError("tear down failed: %v\n", err)
		}
	}
	os.Exit(exitStatus)
}

// FixturesOptions fixtures needs to be loaded options
type FixturesOptions struct {
	Dir   string
	Files []string
}

// CreateTestEngine creates a memory database and loads the fixture data from fixturesDir
func CreateTestEngine(opts FixturesOptions) error {
	x, err := xorm.NewEngine("sqlite3", "file::memory:?cache=shared&_txlock=immediate")
	if err != nil {
		if strings.Contains(err.Error(), "unknown driver") {
			return fmt.Errorf(`sqlite3 requires: -tags sqlite,sqlite_unlock_notify%s%w`, "\n", err)
		}
		return err
	}
	x.SetMapper(names.GonicMapper{})
	db.SetDefaultEngine(context.Background(), x)

	if err = db.SyncAllTables(); err != nil {
		return err
	}
	switch os.Getenv("GITEA_UNIT_TESTS_LOG_SQL") {
	case "true", "1":
		x.ShowSQL(true)
	}

	return InitFixtures(opts)
}

// PrepareTestDatabase load test fixtures into test database
func PrepareTestDatabase() error {
	return LoadFixtures()
}

// PrepareTestEnv prepares the environment for unit tests. Can only be called
// by tests that use the above MainTest(..) function.
func PrepareTestEnv(t testing.TB) {
	assert.NoError(t, PrepareTestDatabase())
<<<<<<< HEAD
	metaPath := filepath.Join(giteaRoot, "tests", "gitea-repositories-meta")
=======
	metaPath := filepath.Join(proxgitRoot, "tests", "proxgit-repositories-meta")
>>>>>>> master
	assert.NoError(t, SyncDirs(metaPath, setting.RepoRootPath))
	test.SetupGiteaRoot() // Makes sure GITEA_ROOT is set
}
