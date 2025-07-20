package main

import (
	"context"
	"log"
	"os"
	"strings"

	_env "github.com/jgttech/repo/cmds/env"
	_export "github.com/jgttech/repo/cmds/export"
	_get "github.com/jgttech/repo/cmds/get"
	_import "github.com/jgttech/repo/cmds/import"
	_install "github.com/jgttech/repo/cmds/install"
	_set "github.com/jgttech/repo/cmds/set"
	_uninstall "github.com/jgttech/repo/cmds/uninstall"
	_version "github.com/jgttech/repo/cmds/version"
	"github.com/jgttech/repo/src/cmd"
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
			_install.Command(),
			_import.Command(),
			cmd.Protected(_env.Command()),
			cmd.Protected(_uninstall.Command()),
			cmd.Protected(_export.Command()),
			cmd.Protected(_set.Command()),
			cmd.Protected(_get.Command()),
			cmd.Protected(_version.Command()),
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
