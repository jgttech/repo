package _export

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:    "export",
		Aliases: []string{"exp"},
		Action: func(ctx context.Context, c *cli.Command) (err error) {
			fmt.Println("export")
			return
		},
	}
}
