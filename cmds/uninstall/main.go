package uninstall

import (
	"context"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name: "uninstall",
		Action: func(ctx context.Context, c *cli.Command) (err error) {
			return
		},
	}
}
