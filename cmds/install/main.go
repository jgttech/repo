package install

import (
	"context"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/jgttech/repo/src/env"
	"github.com/jgttech/repo/src/fs/cp"
	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	home := os.Getenv("HOME")
	timestamp := time.Now().Unix()
	cliconfig := path.Join(env.BASE_SYMLINK, ".gitconfig")
	backupName := fmt.Sprintf(".gitconfig.%d.bak", timestamp)
	gitconfig := path.Join(home, ".gitconfig")
	backup := path.Join(home, backupName)

	return &cli.Command{
		Name:                  "install",
		EnableShellCompletion: true,
		Action: func(ctx context.Context, c *cli.Command) (err error) {
			_, err = os.Stat(env.BASE_DIR)

			if err == nil {
				err = fmt.Errorf("Already installed.")
				return
			} else {
				err = os.MkdirAll(env.BASE_SYMLINK, 0755)
				err = os.MkdirAll(env.BASE_DATA, 0755)
			}

			err = cp.File(cp.From(gitconfig), cp.To(backup))
			err = os.Remove(gitconfig)
			file, err := os.Create(cliconfig)
			file.Close()

			return
		},
	}
}
