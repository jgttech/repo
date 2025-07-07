package add

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/jgttech/repo/src/self"
	"github.com/urfave/cli/v3"
)

var Command = &cli.Command{
	Name: "add",
	Action: func(_ context.Context, _ *cli.Command) (err error) {
		cwd := self.Context.Pwd.Path
		gitdir := path.Join(cwd, ".git")
		_, err = os.Stat(gitdir)

		if os.IsNotExist(err) {
			return
		}

		fmt.Println(cwd)
		return
	},
}
