package version

import (
	"context"
	"fmt"

	"github.com/jgttech/repo/src/self"
	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name: "version",
		Action: func(ctx context.Context, c *cli.Command) (err error) {
			fmt.Println(self.Context.Version)
			return
		},
	}
}
