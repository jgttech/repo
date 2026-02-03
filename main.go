package main

import (
	"context"
	"os"

	"log"
	CommandVersion "securacore/repo/cmds/version"
	"securacore/repo/core/env"
	"securacore/repo/core/sys"

	"github.com/urfave/cli/v3"
)

func main() {
	sys.Healthcheck()

	app := cli.Command{
		Name:    "repo",
		Version: env.REPO_VERSION,
		Usage:   "Repo CLI the Git Account Multiplexer",
		Commands: []*cli.Command{
			CommandVersion.Command,
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
