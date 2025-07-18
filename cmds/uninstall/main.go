package _uninstall

import (
	"context"

	"github.com/jgttech/repo/src/env"
	v3 "github.com/urfave/cli/v3"
)

func Command() *v3.Command {
	BASE := env.BASE

	return &v3.Command{
		Name: "uninstall",
		Action: func(ctx context.Context, c *v3.Command) (err error) {
			return
		},
	}
}

// import (
// 	"context"
// 	"fmt"
// 	"os"
//
// 	"github.com/jgttech/repo/src/env"
// 	"github.com/jgttech/repo/src/exec"
// 	"github.com/jgttech/repo/src/fs/cp"
// 	"github.com/jgttech/repo/src/state"
// 	"github.com/urfave/cli/v3"
// )
//
// func Command() *cli.Command {
// 	home := os.Getenv("HOME")
//
// 	return &cli.Command{
// 		Name:                  "uninstall",
// 		EnableShellCompletion: true,
// 		Action: func(ctx context.Context, c *cli.Command) (err error) {
// 			s := state.New()
//
// 			// Revert all backups that exist.
// 			for _, backup := range s.Backups {
// 				_, err := os.Stat(backup.From)
//
// 				if err == nil {
// 					os.Remove(backup.From)
// 				}
//
// 				cp.File(cp.From(backup.To), cp.To(backup.From))
// 				os.Remove(backup.To)
// 			}
//
// 			// Remove the build symlinks.
// 			arg := fmt.Sprintf("stow -t %s -D %s", home, env.REPO_DIR)
// 			cmd := exec.Cmd(arg, exec.Dir(s.Home))
// 			err = cmd.Run()
//
// 			// Clean up the build directory.
// 			_, err = os.Stat(env.BASE_DIR)
//
// 			if err == nil {
// 				os.RemoveAll(env.BASE_DIR)
// 			}
//
// 			return
// 		},
// 	}
// }
