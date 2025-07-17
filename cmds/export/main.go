package export

import (
	"context"
	"fmt"
	"os"

	"path"
	"strings"

	"time"

	"github.com/jgttech/repo/src/archive"
	"github.com/jgttech/repo/src/env"
	"github.com/jgttech/repo/src/state"
	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name: "export",
		Description: strings.Join([]string{
			"Generates an archive of your CLI",
			"configuration, at the moment the",
			"command is run and places this export",
			"in your ~/Documents directory with",
			"the name 'repocli.<unix_timestamp>.tar.gz'.",
		}, "\n"),
		EnableShellCompletion: true,
		Action: func(ctx context.Context, c *cli.Command) (err error) {
			s := state.New()
			home := os.Getenv("HOME")
			name := fmt.Sprintf("repocli.%d.tar.gz", time.Now().Unix())

			in := path.Join(s.Home, env.REPO_DIR)
			out := path.Join(home, name)

			err = archive.Create(
				archive.From(in),
				archive.To(out),
			)

			return
		},
	}
}
