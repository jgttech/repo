package version

import (
	"context"
	"fmt"

	"github.com/jgttech/repo/src/env"
	"github.com/urfave/cli/v3"
)

func action(ctx context.Context, c *cli.Command) error {
	fmt.Println(env.REPO_HOME)
	return nil
}
