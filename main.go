package main

import (
	"context"
	"os"

	"github.com/jgttech/repo/cmds/install"
	"github.com/jgttech/repo/cmds/self"
	"github.com/jgttech/repo/cmds/version"
	// "github.com/jgttech/repo/core/fs/node"
	"github.com/jgttech/repo/core/log"
	// "github.com/jgttech/repo/core/state"
	"github.com/jgttech/repo/core/sys"
	"github.com/urfave/cli/v3"
)

func main() {
	log.Configure()
	defer log.Cleanup()

	// ctx := state.Load(node.StateFile)
	sys.Healthcheck()

	app := cli.Command{
		EnableShellCompletion: true,
		Name:                  "repo",
		Commands: []*cli.Command{
			_version.Command(),
			_install.Command(),
			_self.Command(),
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatalln(err)
	}
}
