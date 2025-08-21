package uninstall

import (
	"context"
	"github.com/jgttech/repo/core/fs/node"
	"github.com/jgttech/repo/core/log"
	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	baseDir := node.BaseDir

	return &cli.Command{
		Name: "uninstall",
		Action: func(ctx context.Context, c *cli.Command) (err error) {
			log.Args()

			if baseDir.Exists() {
				err = baseDir.Remove()
			}

			return
		},
	}
}
