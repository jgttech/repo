package cli

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func ProtectedCommand(enabled bool, cmd *cli.Command) *cli.Command {
	if !enabled {
		cmd.Action = func(_ context.Context, _ *cli.Command) error {
			fmt.Println("Must install CLI. Please run 'repo install'.")
			return nil
		}
	}

	return cmd
}
