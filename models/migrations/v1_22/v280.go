// Copyright 2023 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package v1_22 //nolint

import (
	"xorm.io/xorm"
)

func RenameUserThemes(x *xorm.Engine) error {
	sess := x.NewSession()
	defer sess.Close()

	if err := sess.Begin(); err != nil {
		return err
	}

	if _, err := sess.Exec("UPDATE `user` SET `theme` = 'proxgit-light' WHERE `theme` = 'proxgit'"); err != nil {
		return err
	}
	if _, err := sess.Exec("UPDATE `user` SET `theme` = 'proxgit-dark' WHERE `theme` = 'arc-green'"); err != nil {
		return err
	}
	if _, err := sess.Exec("UPDATE `user` SET `theme` = 'proxgit-auto' WHERE `theme` = 'auto'"); err != nil {
		return err
	}

	return sess.Commit()
}
