// Copyright 2024 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package lfstransfer

import (
	"context"
	"fmt"
	"os"

<<<<<<< HEAD
	"code.gitea.io/gitea/modules/lfstransfer/backend"
=======
	"code.proxgit.io/proxgit/modules/lfstransfer/backend"
>>>>>>> master

	"github.com/charmbracelet/git-lfs-transfer/transfer"
)

func Main(ctx context.Context, repo, verb, token string) error {
	logger := newLogger()
	pktline := transfer.NewPktline(os.Stdin, os.Stdout, logger)
<<<<<<< HEAD
	giteaBackend, err := backend.New(ctx, repo, verb, token, logger)
=======
	proxgitBackend, err := backend.New(ctx, repo, verb, token, logger)
>>>>>>> master
	if err != nil {
		return err
	}

	for _, cap := range backend.Capabilities {
		if err := pktline.WritePacketText(cap); err != nil {
			logger.Log("error sending capability due to error:", err)
		}
	}
	if err := pktline.WriteFlush(); err != nil {
		logger.Log("error flushing capabilities:", err)
	}
<<<<<<< HEAD
	p := transfer.NewProcessor(pktline, giteaBackend, logger)
=======
	p := transfer.NewProcessor(pktline, proxgitBackend, logger)
>>>>>>> master
	defer logger.Log("done processing commands")
	switch verb {
	case "upload":
		return p.ProcessCommands(transfer.UploadOperation)
	case "download":
		return p.ProcessCommands(transfer.DownloadOperation)
	default:
		return fmt.Errorf("unknown operation %q", verb)
	}
}
