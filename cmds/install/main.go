package _install

import (
	"context"
	"fmt"
	"os"
	"path"
	"time"

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
	HOME := env.HOME
	BUILD_DIR := env.BUILD_DIR

	// Context in which the "install" operates.
	cwd := env.BASE

	// We need to load the config from the repo
	// for doing the installation and check that
	// all the dependencies are installed.
	conf := &cli.Conf{}

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
			// Loads the CLI that comes from the the repo, itself.
			conf.Load(cli.File(env.LOCAL_CLI))

			if _os.SystemInstalled() {
				err = fmt.Errorf("Already installed.")
				return
			}

			// Check dependencies and error on missing deps.
			if err = _os.HasDependencies(conf); err != nil {
				return
			}

			// Where things are at and what they are called.
			timestamp := time.Now().Unix()
			gitBackupName := fmt.Sprintf(".gitconfig.%d.bak", timestamp)
			gitConfigPath := path.Join(HOME, ".gitconfig")
			gitBackupPath := path.Join(HOME, gitBackupName)
			gitConfig := fs.NewNode(path.Join(BUILD_DIR, ".gitconfig"), fs.File)

			err = os.MkdirAll(env.BUILD_CONF, 0755)
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
			_, err = os.Stat(env.LOCAL_CLI)

			if os.IsNotExist(err) {
				return
			}

			// Copy the CLI config into the build.
			err = cp.File(cp.From(env.LOCAL_CLI), cp.To(env.BUILD_CLI))

			if err == nil {
				conf.Load(cli.File(env.BUILD_CLI))
				conf.Exports = os.ExpandEnv(conf.Exports)
				conf.Save()
			}

			// Generate all the symlinks from the build.
			arg := fmt.Sprintf("stow -t %s %s", HOME, env.BUILD_DIR)
			cmd := exec.Cmd(arg, exec.Dir(cwd), exec.Stdio)
			err = cmd.Run()

			return
		},
	}
}
