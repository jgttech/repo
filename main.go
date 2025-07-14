package main

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/jgttech/repo/cmds/export"
	"github.com/jgttech/repo/cmds/install"
	"github.com/jgttech/repo/cmds/uninstall"
	"github.com/urfave/cli/v3"
)

func main() {
	app := cli.Command{
		Name:    "repo",
		Usage:   "Git Account Multiplexer",
		Authors: []any{"jgt.tech@protonmail.com"},
		Description: strings.Join([]string{
			"CLI manager for Git + SSH configurations",
			"across multiple repo's and accounts.",
		}, "\n"),
		Commands: []*cli.Command{
			install.Command(),
			uninstall.Command(),
			export.Command(),
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
