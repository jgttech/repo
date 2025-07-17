package cmdImport

import (
	"context"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name: "import",
		Action: func(ctx context.Context, c *cli.Command) error {
			return nil
		},
	}
}
