package _set

import (
	"context"
	"os"
	"strings"

	"github.com/jgttech/repo/src/cli"
	v3 "github.com/urfave/cli/v3"
)

func Command() *v3.Command {
	return &v3.Command{
		Name: "set",
		Description: strings.Join([]string{
			"Allows you to set a config setting with",
			"a simple update syntax that affects how",
			"the CLI works.",
		}, "\n"),
		EnableShellCompletion: true,
		Action: func(ctx context.Context, c *v3.Command) (err error) {
			conf := &cli.Conf{}
			conf.Load()

			key := c.Args().Get(0)
			value := c.Args().Get(1)

			if key != "" && value != "" {
				switch key {
				case "exports":
					conf.Exports = os.ExpandEnv(value)
					conf.Save()
				}
			}

			return
		},
	}
}
