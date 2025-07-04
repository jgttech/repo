package install

import "github.com/urfave/cli/v3"

func Command() *cli.Command {
	return &cli.Command{
		Name:   "install",
		Action: action,
	}
}
