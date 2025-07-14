package export

import (
	"context"
	"strings"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name: "export",
		Description: strings.Join([]string{
			"Generates an archive of your CLI",
			"configuration, at the moment the",
			"command is run and places this export",
			"in your ~/Documents directory with",
			"the name 'repocli.<unix_timestamp>.tar.gz'.",
		}, "\n"),
		EnableShellCompletion: true,
		Action: func(ctx context.Context, c *cli.Command) (err error) {
			return
		},
	}
}
