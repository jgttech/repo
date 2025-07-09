package remove

import (
	"context"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name: "remove",
		Action: func(ctx context.Context, c *cli.Command) error {
			return nil
		},
	}
}
