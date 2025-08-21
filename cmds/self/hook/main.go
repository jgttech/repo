package hook

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name: "hook",
		Action: func(ctx context.Context, c *cli.Command) (err error) {
			fmt.Println("HOOK")
			return
		},
	}
}
