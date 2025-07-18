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
			fmt.Println("|")
			fmt.Println("| ERROR")
			fmt.Println("| Must install or import CLI config.")
			fmt.Println("| Please, run one of the following:")
			fmt.Println("|")
			fmt.Println("| 1. $ repo install")
			fmt.Println("| 2. $ repo import")
			fmt.Println("|")
			return nil
		}
	}

	return cmd
}
