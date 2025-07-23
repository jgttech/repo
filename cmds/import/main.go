package _import

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jgttech/repo/src/archive"
	"github.com/jgttech/repo/src/env"
	"github.com/jgttech/repo/src/errors"
	"github.com/jgttech/repo/src/exec"
	"github.com/jgttech/repo/src/state"
	v3 "github.com/urfave/cli/v3"
)

func Command() *v3.Command {
	return &v3.Command{
		Name: "import",
		Action: func(ctx context.Context, c *v3.Command) (err error) {
			target := c.Args().First()

			if target == "" {
				err = errors.New(strings.Join([]string{
					"Missing the required argument pointing",
					"to the .tar.gz archive.",
				}, "\n"))

				return
			}

			if _, err = os.Stat(target); os.IsNotExist(err) {
				err = errors.Errorf("File does not exist: %s", target)
				return
			}

			base, err := os.MkdirTemp(os.TempDir(), "repocli-")

			if err != nil {
				return
			}

			archive.Extract(
				archive.From(target),
				archive.To(base),
			)

			_, err = os.Stat(env.BUILD_DIR)

			if os.IsNotExist(err) {
				err = create(base)
			} else {
				err = swap(base)
			}

			// TODO
			// Need to load the state and cli configs from the extracted archive.

			// TODO
			// Then need to iterate through the backups in the state config and
			// re-create the backup files. Consider generating new backup files
			// with updated timestamps.

			// TODO
			// Ensure the backup files that are re-created delete the originals
			// so the the symlink works.

			// fmt.Println("BUILD_CLI...:", env.BUILD_CLI)
			// fmt.Println("BUILD_STATE.:", env.BUILD_STATE)
			// fmt.Println("BUILD_CONF..:", env.BUILD_CONF)

			// dir := filepath.Join(env.BUILD_DIR, "..")
			// arg := fmt.Sprintf("stow -t %s %s", env.HOME, env.CONST_DIR)
			// cmd := exec.Cmd(arg, exec.Dir(dir), exec.Stdio)
			// err = cmd.Run()

			// WIP
			// conf := &state.Conf{}
			// conf.Load()

			return
		},
	}
}
