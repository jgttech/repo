package setup

import "github.com/urfave/cli/v3"

func Command() *cli.Command {
	return &cli.Command{
		Name:   "setup",
		Action: action,
	}
}
