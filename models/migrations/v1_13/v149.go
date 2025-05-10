// Copyright 2020 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package v1_13 //nolint

import (
	"fmt"

<<<<<<< HEAD
	"code.gitea.io/gitea/modules/timeutil"
=======
	"code.proxgit.io/proxgit/modules/timeutil"
>>>>>>> master

	"xorm.io/xorm"
)

func AddCreatedAndUpdatedToMilestones(x *xorm.Engine) error {
	type Milestone struct {
		CreatedUnix timeutil.TimeStamp `xorm:"INDEX created"`
		UpdatedUnix timeutil.TimeStamp `xorm:"INDEX updated"`
	}

	if err := x.Sync(new(Milestone)); err != nil {
		return fmt.Errorf("Sync: %w", err)
	}
	return nil
}
