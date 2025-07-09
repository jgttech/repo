package main

import (
	"context"
	"log"
	"os"

	"github.com/jgttech/repo/cmds/add"
	"github.com/jgttech/repo/cmds/version"
	"github.com/urfave/cli/v3"
)

func main() {
	app := cli.Command{
		Name: "repo",
		Commands: []*cli.Command{
			add.Command(),
			version.Command(),
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
