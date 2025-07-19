package _uninstall

import (
	"context"
	"fmt"
	"os"

	"github.com/jgttech/repo/src/env"
	"github.com/jgttech/repo/src/exec"
	"github.com/jgttech/repo/src/fs/cp"
	_os "github.com/jgttech/repo/src/os"
	"github.com/jgttech/repo/src/state"
	v3 "github.com/urfave/cli/v3"
)

func Command() *v3.Command {
	HOME := env.HOME

	return &v3.Command{
		Name: "uninstall",
		Action: func(ctx context.Context, c *v3.Command) (err error) {
			sconf := &state.Conf{}
			sconf.Load()

			// Revert all backups that exist.
			for _, backup := range sconf.Backups {
				_, err = os.Stat(backup.From)

				if err == nil {
					if err = _os.SafeRemove(backup.From); err != nil {
						return
					}
				}

				if err = cp.File(cp.From(backup.To), cp.To(backup.From)); err != nil {
					return
				}

				if err = _os.SafeRemove(backup.To); err != nil {
					return
				}
			}

			// Remove the build symlinks.
			arg := fmt.Sprintf("stow -t %s -D %s", HOME, env.CONST_DIR)
			cmd := exec.Cmd(arg, exec.Dir(sconf.Home))

			if err = cmd.Run(); err != nil {
				return
			}

			// Clean up the build directory.
			_, err = os.Stat(env.BUILD_DIR)

			if err == nil {
				os.RemoveAll(env.BUILD_DIR)
			}

			return
		},
	}
}
