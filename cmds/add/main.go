package add

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

var Command = &cli.Command{
	Name: "add",
	Action: func(ctx context.Context, c *cli.Command) (err error) {
		fmt.Println("add <arg>:", c.Args().First())
		return
	},
}
