package CommandVersion

import (
	"context"
	"fmt"
	"securacore/repo/core/env"

	"github.com/urfave/cli/v3"
)

var Command = &cli.Command{
	Name:  "version",
	Usage: "Display the CLI version",
	Action: func(ctx context.Context, c *cli.Command) error {
		fmt.Println(env.REPO_VERSION)
		return nil
	},
}
