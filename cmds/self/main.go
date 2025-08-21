package _self

import (
	"github.com/jgttech/repo/cmds/self/install"
	"github.com/jgttech/repo/cmds/self/uninstall"
	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name: "self",
		Commands: []*cli.Command{
			install.Command(),
			uninstall.Command(),
		},
	}
}
