package cli

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func Protected(cmd *cli.Command) *cli.Command {
	enabled := IsInstalled()

	if !enabled {
		cmd.Action = func(_ context.Context, _ *cli.Command) error {
			fmt.Println("[ERROR]")
			fmt.Println("Must install or import CLI config. Please,")
			fmt.Println("run 'repo install' or 'repo import'.")
			return nil
		}
	}

	return cmd
}
