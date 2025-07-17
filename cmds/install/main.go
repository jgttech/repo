package install

import (
	"context"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/jgttech/repo/src/assert"
	"github.com/jgttech/repo/src/env"
	"github.com/jgttech/repo/src/exec"
	"github.com/jgttech/repo/src/fs"
	"github.com/jgttech/repo/src/fs/cp"
	"github.com/jgttech/repo/src/state"
	v3 "github.com/urfave/cli/v3"
)

func Command() *v3.Command {
	cwd := assert.Must(os.Getwd())
	timestamp := time.Now().Unix()
	home := os.Getenv("HOME")
	backupName := fmt.Sprintf(".gitconfig.%d.bak", timestamp)
	gitConfigPath := path.Join(home, ".gitconfig")
	gitBackupPath := path.Join(home, backupName)
	gitConfig := fs.NewNode(path.Join(env.BASE_DIR, ".gitconfig"), fs.File)

	return &v3.Command{
		Name:                  "install",
		EnableShellCompletion: true,
		Flags: []v3.Flag{
			&v3.StringFlag{
				Name:        "cwd",
				Destination: &cwd,
			},
		},
		Action: func(ctx context.Context, c *v3.Command) (err error) {
			_, err = os.Stat(env.BASE_DIR)

			if err == nil {
				err = fmt.Errorf("Already installed.")
				return
			}

			err = os.MkdirAll(env.BASE_CONF, 0755)
			err = cp.File(cp.From(gitConfigPath), cp.To(gitBackupPath))
			err = os.Remove(gitConfigPath)

			// Create the .gitconfig the CLI manages.
			gitConfig.Create()

			// Create the new state.
			s := state.New()

			// Add backed up .gitconfig to install config.
			s.AddBackup(gitConfigPath, gitBackupPath)
			s.Save()

			// Make sure the config exists.
			cliConfig := path.Join(s.Home, env.REPO_CLI)
			_, err = os.Stat(cliConfig)

			if os.IsNotExist(err) {
				return
			}

			// Copy the CLI config into the build.
			err = cp.File(cp.From(cliConfig), cp.To(env.BASE_CLI))

			// Generate all the symlinks from the build.
			arg := fmt.Sprintf("stow -t %s %s", home, env.REPO_DIR)
			cmd := exec.Cmd(arg, exec.Dir(cwd), exec.Stdio)
			err = cmd.Run()

			return
		},
	}
}
