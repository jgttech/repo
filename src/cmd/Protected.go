package cmd

import (
	"context"
	"fmt"

	"github.com/jgttech/repo/src/cli"
	"github.com/jgttech/repo/src/os"
	v3 "github.com/urfave/cli/v3"
)

func Protected(cmd *v3.Command) *v3.Command {
	installed := os.SystemInstalled()

	// If the CLI is NOT installed we can't run the
	// protected commands as they require the CLI to
	// be properly configured.
	if !installed {
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

	conf := &cli.Conf{}
	conf.Load()

	// We need to ensure that all the dependencies are
	// installed in the system on ALL protected commands.
	if err := os.HasDependencies(conf); err != nil {
		cmd.Action = func(_ context.Context, _ *v3.Command) error {
			return err
		}
	}

	return cmd
}
