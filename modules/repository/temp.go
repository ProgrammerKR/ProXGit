// Copyright 2019 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package repository

import (
	"context"
	"fmt"

	"code.proxgit.io/proxgit/modules/log"
	"code.proxgit.io/proxgit/modules/setting"
)

// CreateTemporaryPath creates a temporary path
func CreateTemporaryPath(prefix string) (string, context.CancelFunc, error) {
	basePath, cleanup, err := setting.AppDataTempDir("local-repo").MkdirTempRandom(prefix + ".git")
	if err != nil {
		log.Error("Unable to create temporary directory: %s-*.git (%v)", prefix, err)
		return "", nil, fmt.Errorf("failed to create dir %s-*.git: %w", prefix, err)
	}
	return basePath, cleanup, nil
}
