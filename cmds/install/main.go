package _install

import (
	"context"
	"fmt"

	"github.com/jgttech/repo/src/cli"
	"github.com/jgttech/repo/src/env"
	"github.com/jgttech/repo/src/os"
	v3 "github.com/urfave/cli/v3"
)

func Command() *v3.Command {
	// HOME := env.HOME
	// BUILD_DIR := env.BUILD_DIR

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

			if os.SystemInstalled() {
				err = fmt.Errorf("Already installed.")
				return
			}

			// Check dependencies and error on missing deps.
			if err = os.HasDependencies(conf); err != nil {
				return
			}

			// Where things are at and what they are called.
			// timestamp := time.Now().Unix()
			// gitBackupName := fmt.Sprintf(".gitconfig.%d.bak", timestamp)
			// gitConfigPath := path.Join(HOME, ".gitconfig")
			// gitBackupPath := path.Join(HOME, gitBackupName)
			// gitConfig := fs.NewNode(path.Join(BUILD_DIR, ".gitconfig"), fs.File)

			// fmt.Println(env.BUILD_DIR)
			return
		},
	}
}

// func Command() *v3.Command {
// 	cwd := assert.Must(os.Getwd())
// 	timestamp := time.Now().Unix()
// 	home := os.Getenv("HOME")
// 	backupName := fmt.Sprintf(".gitconfig.%d.bak", timestamp)
// 	gitConfigPath := path.Join(home, ".gitconfig")
// 	gitBackupPath := path.Join(home, backupName)
// 	gitConfig := fs.NewNode(path.Join(env.BASE_DIR, ".gitconfig"), fs.File)
//
// 	return &v3.Command{
// 		Name:                  "install",
// 		EnableShellCompletion: true,
// 		Flags: []v3.Flag{
// 			&v3.StringFlag{
// 				Name:        "cwd",
// 				Destination: &cwd,
// 			},
// 		},
// 		Action: func(ctx context.Context, c *v3.Command) (err error) {
// 			base := []string{env.HOME, env.REPO_CLI}
// 			confPath := strings.Join(base, string(filepath.Separator))
// 			conf := assert.Must(cli.Load(confPath))
// 			missing := []string{}
//
// 			for _, dep := range conf.Dependencies {
// 				if !_os.PackageInstalled(dep) {
// 					missing = append(missing, dep)
// 				}
// 			}
//
// 			if len(missing) > 0 {
// 				fmt.Printf("\n|\n| Missing (%d) required dependencies:\n", len(missing))
//
// 				for _, pkg := range missing {
// 					fmt.Printf("| - %s\n", pkg)
// 				}
//
// 				fmt.Printf("|\n\n")
//
// 				err = fmt.Errorf("\n\n"+
// 					"|\n"+
// 					"| ERROR\n"+
// 					"| Please install required dependencies and\n"+
// 					"| run 'bash ~/%s/install' to complete\n"+
// 					"| the installation\n"+
// 					"|\n",
// 					filepath.Base(env.HOME),
// 				)
//
// 				return
// 			}
//
// 			_, err = os.Stat(env.BASE_CONF)
//
// 			if err == nil {
// 				err = fmt.Errorf("Already installed.")
// 				return
// 			}
//
// 			err = os.MkdirAll(env.BASE_CONF, 0755)
// 			err = cp.File(cp.From(gitConfigPath), cp.To(gitBackupPath))
// 			err = os.Remove(gitConfigPath)
//
// 			// Create the .gitconfig the CLI manages.
// 			gitConfig.Create()
//
// 			// Create the new state.
// 			s := state.New()
//
// 			// Add backed up .gitconfig to install config.
// 			s.AddBackup(gitConfigPath, gitBackupPath)
// 			s.Save()
//
// 			// Make sure the config exists.
// 			cliConfig := path.Join(s.Home, env.REPO_CLI)
// 			_, err = os.Stat(cliConfig)
//
// 			if os.IsNotExist(err) {
// 				return
// 			}
//
// 			// Copy the CLI config into the build.
// 			err = cp.File(cp.From(cliConfig), cp.To(env.BASE_CLI))
//
// 			if err == nil {
// 				base := assert.Must(cli.Load(env.BASE_CLI))
// 				base.Exports = os.ExpandEnv(base.Exports)
// 				base.Save()
// 			}
//
// 			// Generate all the symlinks from the build.
// 			arg := fmt.Sprintf("stow -t %s %s", home, env.REPO_DIR)
// 			cmd := exec.Cmd(arg, exec.Dir(cwd), exec.Stdio)
// 			err = cmd.Run()
//
// 			return
// 		},
// 	}
// }
