package main

import (
	"context"
	"log"
	"os"

	"github.com/jgttech/repo/cmds/install"
	"github.com/jgttech/repo/cmds/uninstall"
	"github.com/urfave/cli/v3"
)

func main() {
	app := cli.Command{
		Name: "repo",
		Commands: []*cli.Command{
			install.Command(),
			uninstall.Command(),
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
