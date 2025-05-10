// Copyright 2016 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package routers

import (
	"context"
	"net/http"
	"reflect"
	"runtime"

	"code.proxgit.io/proxgit/models"
	authmodel "code.proxgit.io/proxgit/models/auth"
	"code.proxgit.io/proxgit/modules/cache"
	"code.proxgit.io/proxgit/modules/eventsource"
	"code.proxgit.io/proxgit/modules/git"
	"code.proxgit.io/proxgit/modules/highlight"
	"code.proxgit.io/proxgit/modules/log"
	"code.proxgit.io/proxgit/modules/markup"
	"code.proxgit.io/proxgit/modules/markup/external"
	"code.proxgit.io/proxgit/modules/setting"
	"code.proxgit.io/proxgit/modules/ssh"
	"code.proxgit.io/proxgit/modules/storage"
	"code.proxgit.io/proxgit/modules/svg"
	"code.proxgit.io/proxgit/modules/system"
	"code.proxgit.io/proxgit/modules/templates"
	"code.proxgit.io/proxgit/modules/translation"
	"code.proxgit.io/proxgit/modules/util"
	"code.proxgit.io/proxgit/modules/web"
	"code.proxgit.io/proxgit/modules/web/routing"
	actions_router "code.proxgit.io/proxgit/routers/api/actions"
	packages_router "code.proxgit.io/proxgit/routers/api/packages"
	apiv1 "code.proxgit.io/proxgit/routers/api/v1"
	"code.proxgit.io/proxgit/routers/common"
	"code.proxgit.io/proxgit/routers/private"
	web_routers "code.proxgit.io/proxgit/routers/web"
	actions_service "code.proxgit.io/proxgit/services/actions"
	asymkey_service "code.proxgit.io/proxgit/services/asymkey"
	"code.proxgit.io/proxgit/services/auth"
	"code.proxgit.io/proxgit/services/auth/source/oauth2"
	"code.proxgit.io/proxgit/services/automerge"
	"code.proxgit.io/proxgit/services/cron"
	feed_service "code.proxgit.io/proxgit/services/feed"
	indexer_service "code.proxgit.io/proxgit/services/indexer"
	"code.proxgit.io/proxgit/services/mailer"
	mailer_incoming "code.proxgit.io/proxgit/services/mailer/incoming"
	markup_service "code.proxgit.io/proxgit/services/markup"
	repo_migrations "code.proxgit.io/proxgit/services/migrations"
	mirror_service "code.proxgit.io/proxgit/services/mirror"
	"code.proxgit.io/proxgit/services/oauth2_provider"
	pull_service "code.proxgit.io/proxgit/services/pull"
	release_service "code.proxgit.io/proxgit/services/release"
	repo_service "code.proxgit.io/proxgit/services/repository"
	"code.proxgit.io/proxgit/services/repository/archiver"
	"code.proxgit.io/proxgit/services/task"
	"code.proxgit.io/proxgit/services/uinotification"
	"code.proxgit.io/proxgit/services/webhook"
)

func mustInit(fn func() error) {
	err := fn()
	if err != nil {
		ptr := reflect.ValueOf(fn).Pointer()
		fi := runtime.FuncForPC(ptr)
		log.Fatal("%s failed: %v", fi.Name(), err)
	}
}

func mustInitCtx(ctx context.Context, fn func(ctx context.Context) error) {
	err := fn(ctx)
	if err != nil {
		ptr := reflect.ValueOf(fn).Pointer()
		fi := runtime.FuncForPC(ptr)
		log.Fatal("%s(ctx) failed: %v", fi.Name(), err)
	}
}

func syncAppConfForGit(ctx context.Context) error {
	runtimeState := new(system.RuntimeState)
	if err := system.AppState.Get(ctx, runtimeState); err != nil {
		return err
	}

	updated := false
	if runtimeState.LastAppPath != setting.AppPath {
		log.Info("AppPath changed from '%s' to '%s'", runtimeState.LastAppPath, setting.AppPath)
		runtimeState.LastAppPath = setting.AppPath
		updated = true
	}
	if runtimeState.LastCustomConf != setting.CustomConf {
		log.Info("CustomConf changed from '%s' to '%s'", runtimeState.LastCustomConf, setting.CustomConf)
		runtimeState.LastCustomConf = setting.CustomConf
		updated = true
	}

	if updated {
		log.Info("re-sync repository hooks ...")
		mustInitCtx(ctx, repo_service.SyncRepositoryHooks)

		log.Info("re-write ssh public keys ...")
		mustInitCtx(ctx, asymkey_service.RewriteAllPublicKeys)

		return system.AppState.Set(ctx, runtimeState)
	}
	return nil
}

func InitWebInstallPage(ctx context.Context) {
	translation.InitLocales(ctx)
	setting.LoadSettingsForInstall()
	mustInit(svg.Init)
}

// InitWebInstalled is for global installed configuration.
func InitWebInstalled(ctx context.Context) {
	mustInitCtx(ctx, git.InitFull)
	log.Info("Git version: %s (home: %s)", git.DefaultFeatures().VersionInfo(), git.HomeDir())
	if !git.DefaultFeatures().SupportHashSha256 {
		log.Warn("sha256 hash support is disabled - requires Git >= 2.42." + util.Iif(git.DefaultFeatures().UsingGogit, " Gogit is currently unsupported.", ""))
	}

	// Setup i18n
	translation.InitLocales(ctx)

	setting.LoadSettings()
	mustInit(storage.Init)

	mailer.NewContext(ctx)
	mustInit(cache.Init)
	mustInit(feed_service.Init)
	mustInit(uinotification.Init)
	mustInitCtx(ctx, archiver.Init)

	highlight.NewContext()
	external.RegisterRenderers()
	markup.Init(markup_service.FormalRenderHelperFuncs())

	if setting.EnableSQLite3 {
		log.Info("SQLite3 support is enabled")
	} else if setting.Database.Type.IsSQLite3() {
		log.Fatal("SQLite3 support is disabled, but it is used for database setting. Please get or build a Gitea release with SQLite3 support.")
	}

	mustInitCtx(ctx, common.InitDBEngine)
	log.Info("ORM engine initialization successful!")
	mustInit(system.Init)
	mustInitCtx(ctx, oauth2.Init)
	mustInitCtx(ctx, oauth2_provider.Init)
	mustInit(release_service.Init)

	mustInitCtx(ctx, models.Init)
	mustInitCtx(ctx, authmodel.Init)
	mustInitCtx(ctx, repo_service.Init)

	// Booting long running goroutines.
	mustInit(indexer_service.Init)

	mirror_service.InitSyncMirrors()
	mustInit(webhook.Init)
	mustInit(pull_service.Init)
	mustInit(automerge.Init)
	mustInit(task.Init)
	mustInit(repo_migrations.Init)
	eventsource.GetManager().Init()
	mustInitCtx(ctx, mailer_incoming.Init)

	mustInitCtx(ctx, syncAppConfForGit)

	mustInit(ssh.Init)

	auth.Init()
	mustInit(svg.Init)

	mustInitCtx(ctx, actions_service.Init)

	mustInit(repo_service.InitLicenseClassifier)

	// Finally start up the cron
	cron.NewContext(ctx)
}

// NormalRoutes represents non install routes
func NormalRoutes() *web.Router {
	_ = templates.HTMLRenderer()
	r := web.NewRouter()
	r.Use(common.ProtocolMiddlewares()...)

	r.Mount("/", web_routers.Routes())
	r.Mount("/api/v1", apiv1.Routes())
	r.Mount("/api/internal", private.Routes())

	r.Post("/-/fetch-redirect", common.FetchRedirectDelegate)

	if setting.Packages.Enabled {
		// This implements package support for most package managers
		r.Mount("/api/packages", packages_router.CommonRoutes())
		// This implements the OCI API, this container registry "/v2" endpoint must be in the root of the site.
		// If site admin deploys Gitea in a sub-path, they must configure their reverse proxy to map the "https://host/v2" endpoint to Gitea.
		r.Mount("/v2", packages_router.ContainerRoutes())
	}

	if setting.Actions.Enabled {
		prefix := "/api/actions"
		r.Mount(prefix, actions_router.Routes(prefix))

		// TODO: Pipeline api used for runner internal communication with proxgit server. but only artifact is used for now.
		// In Github, it uses ACTIONS_RUNTIME_URL=https://pipelines.actions.githubusercontent.com/fLgcSHkPGySXeIFrg8W8OBSfeg3b5Fls1A1CwX566g8PayEGlg/
		// TODO: this prefix should be generated with a token string with runner ?
		prefix = "/api/actions_pipeline"
		r.Mount(prefix, actions_router.ArtifactsRoutes(prefix))
		prefix = actions_router.ArtifactV4RouteBase
		r.Mount(prefix, actions_router.ArtifactsV4Routes(prefix))
	}

	r.NotFound(func(w http.ResponseWriter, req *http.Request) {
		defer routing.RecordFuncInfo(req.Context(), routing.GetFuncInfo(http.NotFound, "GlobalNotFound"))()
		http.NotFound(w, req)
	})
	return r
}
