package cliinstall

import (
	"context"
	"repo/cli/core/env"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	home, _ := env.GetInstallHome()

	return &cli.Command{
		Name:  "install",
		Usage: "Install the 'repo' CLI",
		Action: func(ctx context.Context, c *cli.Command) error {
			if home.Exists() {
				return nil
			}

			if err := home.Create(); err != nil {
				return err
			}

			conf, err := env.GetConf()

			if err != nil {
				return err
			}

			if err = conf.Create(); err != nil {
				return err
			}

			return nil
		},
	}
}
