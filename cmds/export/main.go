package export

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/jgttech/repo/src/archive"
	"github.com/jgttech/repo/src/cli"
	"github.com/jgttech/repo/src/env"
	"github.com/jgttech/repo/src/state"
	v3 "github.com/urfave/cli/v3"
)

func Command() *v3.Command {
	return cli.ProtectedCommand(cli.IsInstalled(), &v3.Command{
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
			home := os.Getenv("HOME")
			s := state.New()
			dir := path.Join(s.Home, env.REPO_DIR)
			archiveName := fmt.Sprintf("repocli.%d.tar.gz", time.Now().Unix())
			archivePath := path.Join(home, archiveName)
			export := archive.New(archivePath)

			filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
				if err != nil {
					return err
				}

				if path != dir {
					export.Add(path)
				}

				return nil
			})

			export.Write()
			return
		},
	})
}
