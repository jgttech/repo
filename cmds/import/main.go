package _import

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:    "import",
		Aliases: []string{"imp"},
		Action: func(ctx context.Context, c *cli.Command) (err error) {
			fmt.Println("import")
			return
		},
	}
}
