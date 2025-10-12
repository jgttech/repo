package cliversion

import (
	"context"
	"fmt"
	"repo/cli/core/env"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "version",
		Usage: "Display the CLI version",
		Action: func(ctx context.Context, c *cli.Command) error {
			fmt.Println(env.GetVersion())
			return nil
		},
	}
}
