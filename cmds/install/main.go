package install

import (
	"context"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:        "install",
		Description: "Does initial setup work for using 'repo' to manage your git repositories.",
		Action: func(ctx context.Context, c *cli.Command) error {
			return nil
		},
	}
}
