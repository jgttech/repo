package cliinstall

import (
	"context"
	"repo/cli/core"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	home := core.GetInstallHome()

	return &cli.Command{
		Name:  "install",
		Usage: "Install the 'repo' CLI",
		Action: func(ctx context.Context, c *cli.Command) error {
			if home.Exists() {
				return nil
			}

			home.Create()

			return nil
		},
	}
}
