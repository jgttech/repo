package _get

import (
	"context"
	"fmt"
	"strings"

	"github.com/jgttech/repo/src/cli"
	v3 "github.com/urfave/cli/v3"
)

func Command() *v3.Command {
	return &v3.Command{
		Name: "get",
		Description: strings.Join([]string{
			"Allows you to get a config setting with",
			"a key that matches the path to a known",
			"setting.",
		}, "\n"),
		EnableShellCompletion: true,
		Action: func(ctx context.Context, c *v3.Command) (err error) {
			key := c.Args().Get(0)

			conf := &cli.Conf{}
			err = conf.Load()

			if key != "" {
				switch key {
				case "exports":
					fmt.Println(conf.Exports)
				}
			}

			return
		},
	}
}
