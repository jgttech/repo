package _install

import (
	"context"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/jgttech/repo/src/assert"
	"github.com/jgttech/repo/src/cli"
	"github.com/jgttech/repo/src/env"
	"github.com/jgttech/repo/src/exec"
	"github.com/jgttech/repo/src/fs"
	"github.com/jgttech/repo/src/fs/cp"
	_os "github.com/jgttech/repo/src/os"
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
			base := []string{env.HOME, env.REPO_CLI}
			dir := strings.Join(base, string(filepath.Separator))
			conf := assert.Must(cli.Load(dir))
			missing := []string{}

			for _, dep := range conf.Dependencies {
				if !_os.PackageInstalled(dep) {
					missing = append(missing, dep)
				}
			}

			if len(missing) > 0 {
				fmt.Println("\n|\n| Missing required dependencies:")

				for _, pkg := range missing {
					fmt.Printf("| - %s\n", pkg)
				}

				fmt.Printf("|\n\n")

				err = fmt.Errorf("\n\n"+
					"|\n"+
					"| ERROR\n"+
					"| Please install required dependencies and\n"+
					"| run 'bash ~/%s/install' to complete\n"+
					"| the installation\n"+
					"|\n",
					filepath.Base(env.HOME),
				)

				return
			}

			_, err = os.Stat(env.BASE_CONF)

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
