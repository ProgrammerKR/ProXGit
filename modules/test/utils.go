// Copyright 2017 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"strings"

<<<<<<< HEAD
	"code.gitea.io/gitea/modules/json"
	"code.gitea.io/gitea/modules/util"
=======
	"code.proxgit.io/proxgit/modules/json"
	"code.proxgit.io/proxgit/modules/util"
>>>>>>> master
)

// RedirectURL returns the redirect URL of a http response.
// It also works for JSONRedirect: `{"redirect": "..."}`
func RedirectURL(resp http.ResponseWriter) string {
	loc := resp.Header().Get("Location")
	if loc != "" {
		return loc
	}
	if r, ok := resp.(*httptest.ResponseRecorder); ok {
		m := map[string]any{}
		err := json.Unmarshal(r.Body.Bytes(), &m)
		if err == nil {
			if loc, ok := m["redirect"].(string); ok {
				return loc
			}
		}
	}
	return ""
}

func IsNormalPageCompleted(s string) bool {
	return strings.Contains(s, `<footer class="page-footer"`) && strings.Contains(s, `</html>`)
}

func MockVariableValue[T any](p *T, v ...T) (reset func()) {
	old := *p
	if len(v) > 0 {
		*p = v[0]
	}
	return func() { *p = old }
}

// SetupGiteaRoot Sets GITEA_ROOT if it is not already set and returns the value
func SetupGiteaRoot() string {
<<<<<<< HEAD
	giteaRoot := os.Getenv("GITEA_ROOT")
	if giteaRoot != "" {
		return giteaRoot
	}
	_, filename, _, _ := runtime.Caller(0)
	giteaRoot = filepath.Dir(filepath.Dir(filepath.Dir(filename)))
	fixturesDir := filepath.Join(giteaRoot, "models", "fixtures")
	if exist, _ := util.IsDir(fixturesDir); !exist {
		panic("fixtures directory not found: " + fixturesDir)
	}
	_ = os.Setenv("GITEA_ROOT", giteaRoot)
	return giteaRoot
=======
	proxgitRoot := os.Getenv("GITEA_ROOT")
	if proxgitRoot != "" {
		return proxgitRoot
	}
	_, filename, _, _ := runtime.Caller(0)
	proxgitRoot = filepath.Dir(filepath.Dir(filepath.Dir(filename)))
	fixturesDir := filepath.Join(proxgitRoot, "models", "fixtures")
	if exist, _ := util.IsDir(fixturesDir); !exist {
		panic("fixtures directory not found: " + fixturesDir)
	}
	_ = os.Setenv("GITEA_ROOT", proxgitRoot)
	return proxgitRoot
>>>>>>> master
}
