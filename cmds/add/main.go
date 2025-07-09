package add

import (
	"context"
	"os"
	"path"
	"time"

	"github.com/jgttech/repo/src/cfg"
	"github.com/jgttech/repo/src/self"
	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	var alias string
	primary := self.State.Defaults.Branches.Primary

	return &cli.Command{
		Name:  "add",
		Usage: "Adds the repository to the 'repo' CLI state config.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "alias",
				Aliases:     []string{"A"},
				Destination: &alias,
			},
			&cli.StringFlag{
				Name:        "primary",
				Aliases:     []string{"P"},
				Destination: &primary,
			},
		},
		Action: func(_ context.Context, c *cli.Command) (err error) {
			timestamp := time.Now().Unix()

			args := c.Args()
			gitdir := args.Get(0)
			branch := args.Get(1)
			repo := &cfg.StateRepository{}

			if gitdir == "" || gitdir == "." {
				cwd := self.Context.Pwd
				_, err = os.Stat(path.Join(cwd.Path, ".git"))

				if os.IsNotExist(err) {
					return
				}

				repo.Dir = cwd.Path
				repo.Name = cwd.Name
			} else {
				_, err = os.Stat(path.Join(gitdir, ".git"))

				if os.IsNotExist(err) {
					return
				}

				repo.Dir = gitdir
			}

			repo.CreatedAt = cfg.Timestamp(timestamp)
			repo.UpdatedAt = cfg.Timestamp(timestamp)
			repo.Branch = branch

			if primary == "" {
				repo.Primary = self.State.Defaults.Branches.Primary
			} else {
				repo.Primary = primary
			}

			if alias != "" {
				repo.Alias = alias
			}

			self.State.AddRepository(repo)
			self.State.Save()
			return
		},
	}
}
