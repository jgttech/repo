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

	"github.com/jgttech/repo/src/env"
	"github.com/jgttech/repo/src/state"
	"github.com/jgttech/repo/src/tar"
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
			home := os.Getenv("HOME")
			s := state.New()
			dir := path.Join(s.Home, env.REPO_DIR)
			archiveName := fmt.Sprintf("repocli.%d.tar.gz", time.Now().Unix())
			archivePath := path.Join(home, archiveName)
			archive := tar.New(archivePath)

			filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
				if err != nil {
					return err
				}

				if path != dir {
					archive.Add(path)
				}

				return nil
			})

			archive.Create()

			return
		},
	}
}
