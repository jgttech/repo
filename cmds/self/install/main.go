package install

import (
	"context"
	"strings"

	"github.com/jgttech/repo/core/fs/node"
	"github.com/jgttech/repo/core/log"
	"github.com/jgttech/repo/tui/card"
	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	baseDir := node.BaseDir

	var force bool

	return &cli.Command{
		Name:        "install",
		Description: "Initializes the CLI configuration for installation.",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "force",
				Aliases:     []string{"f"},
				Destination: &force,
			},
		},
		Action: func(ctx context.Context, c *cli.Command) (err error) {
			log.Args()

			if baseDir.Exists() {
				err = log.Println(
					card.GetInfo(
						strings.Join([]string{
							"It appears that the CLI already has a config.",
							"If you wish to run the install process, regardless",
							"of the existing config, use '-f/--force' and the",
							"install will proceed, destroying the existing",
							"config.\n\nFound here: \n%s",
						}, " "),
						baseDir.Path(),
					),
				)

				return
			}

			err = baseDir.Create()

			return
		},
	}
}
