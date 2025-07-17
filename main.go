package main

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/jgttech/repo/cmds/export"
	"github.com/jgttech/repo/cmds/install"
	"github.com/jgttech/repo/cmds/uninstall"
	"github.com/jgttech/repo/src/cli"
	v3 "github.com/urfave/cli/v3"
)

func main() {
	app := v3.Command{
		Name:    "repo",
		Usage:   "Git Account Multiplexer",
		Authors: []any{"jgt.tech@protonmail.com"},
		Description: strings.Join([]string{
			"CLI manager for Git + SSH configurations",
			"across multiple repo's and accounts.",
		}, "\n"),
		Commands: []*v3.Command{
			install.Command(),
			cli.Protected(uninstall.Command()),
			cli.Protected(export.Command()),
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
