package _version

import (
	"context"
	"fmt"

	"github.com/jgttech/repo/core/env"
	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:    "version",
		Aliases: []string{"ver"},
		Action: func(ctx context.Context, c *cli.Command) (err error) {
			fmt.Println(env.VERSION)
			return
		},
	}
}
