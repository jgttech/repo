package _version

import (
	"context"
	"fmt"

	"github.com/jgttech/repo/src/cli"
	v3 "github.com/urfave/cli/v3"
)

func Command() *v3.Command {
	return &v3.Command{
		Name: "version",
		Action: func(ctx context.Context, c *v3.Command) (err error) {
			conf := &cli.Conf{}
			conf.Load()

			fmt.Println(conf.Version)
			return
		},
	}
}
