package install

import (
	"context"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/jgttech/repo/src/cfg"
	"github.com/jgttech/repo/src/fs"
	"github.com/jgttech/repo/src/self"
	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:        "install",
		Description: "Does initial setup work for using 'repo' to manage your git repositories.",
		Action: func(ctx context.Context, c *cli.Command) (err error) {
			timestamp := time.Now().Unix()
			home := os.Getenv("HOME")
			gitconfig := path.Join(home, ".gitconfig")
			gitconfigBackup := path.Join(home, fmt.Sprintf(".gitconfig.%d.bak", timestamp))
			_, err = os.Stat(gitconfig)

			if err == nil {
				err = fs.CopyFile(gitconfig, gitconfigBackup)

				if err == nil {
					self.State.Backups = append(
						self.State.Backups,
						&cfg.StateBackup{
							From: gitconfig,
							To:   gitconfigBackup,
						},
					)
				}
			}

			return
		},
	}
}
