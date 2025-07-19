package _export

import (
	"context"
	"fmt"
	"github.com/jgttech/repo/src/archive"
	"github.com/jgttech/repo/src/cli"
	"github.com/jgttech/repo/src/env"
	"github.com/jgttech/repo/src/state"
	v3 "github.com/urfave/cli/v3"
	"path"
	"strings"
	"time"
)

func Command() *v3.Command {
	return &v3.Command{
		Name: "export",
		Description: strings.Join([]string{
			"Generates an archive of your CLI",
			"configuration, at the moment the",
			"command is run and places this export",
			"in your ~/Documents directory with",
			"the name 'repocli.<unix_timestamp>.tar.gz'.",
		}, "\n"),
		EnableShellCompletion: true,
		Action: func(ctx context.Context, c *v3.Command) (err error) {
			sconf := &state.Conf{}
			sconf.Load()

			conf := &cli.Conf{}
			conf.Load()

			base := conf.Exports
			name := fmt.Sprintf("repocli.%d.tar.gz", time.Now().UnixNano())

			in := path.Join(sconf.Home, env.CONST_DIR)
			out := path.Join(base, name)

			err = archive.Create(
				archive.From(in),
				archive.To(out),
			)

			return
		},
	}
}
