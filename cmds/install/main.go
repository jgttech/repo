package _install

import (
	"context"
	"fmt"

	"github.com/jgttech/repo/core/fs/node"
	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:        "install",
		Description: "Ensures all configuration and repositories exist.",
		Action: func(ctx context.Context, c *cli.Command) error {
			baseDir := node.BaseDir
			baseDir.Ensure()

			fmt.Println("INSTALLING")
			return nil
		},
	}
}
