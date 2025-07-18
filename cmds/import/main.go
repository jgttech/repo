package _import

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name: "import",
		Action: func(ctx context.Context, c *cli.Command) error {
			fmt.Println("IMPORT")
			return nil
		},
	}
}
