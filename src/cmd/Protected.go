package cmd

import (
	"context"
	"fmt"

	"github.com/jgttech/repo/src/cli"
	v3 "github.com/urfave/cli/v3"
)

func Protected(cmd *v3.Command) *v3.Command {
	enabled := cli.IsInstalled()

	if !enabled {
		cmd.Action = func(_ context.Context, _ *v3.Command) error {
			fmt.Println("\n|")
			fmt.Println("| ERROR")
			fmt.Println("| Must install or import CLI config.")
			fmt.Println("| Please, run one of the following:")
			fmt.Println("|")
			fmt.Println("| 1. $ repo install")
			fmt.Println("| 2. $ repo import")
			fmt.Printf("|\n\n")
			return nil
		}
	}

	return cmd
}
