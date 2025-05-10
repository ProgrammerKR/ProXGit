// Copyright 2022 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package user

import (
	"context"

<<<<<<< HEAD
	"code.gitea.io/gitea/models/db"
=======
	"code.proxgit.io/proxgit/models/db"
>>>>>>> master
)

func IncrUserRepoNum(ctx context.Context, userID int64) error {
	_, err := db.GetEngine(ctx).Incr("num_repos").ID(userID).Update(new(User))
	return err
}
