package setup

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func action(ctx context.Context, c *cli.Command) error {
	fmt.Println("Init")
	return nil
}
