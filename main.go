package main

import (
	"context"
	"log"
	"os"
	cliinstall "repo/cli/cmds/install"
	cliversion "repo/cli/cmds/version"

	"github.com/urfave/cli/v3"
)

func main() {
	app := cli.Command{
		Name:  "repo",
		Usage: "Git repository account manager",
		Commands: []*cli.Command{
			cliversion.Command(),
			cliinstall.Command(),
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatalln(err)
	}
}
