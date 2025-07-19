package _env

import (
	"context"
	"fmt"

	"github.com/jgttech/repo/src/env"
	v3 "github.com/urfave/cli/v3"
)

func Command() *v3.Command {
	return &v3.Command{
		Name: "env",
		Action: func(ctx context.Context, c *v3.Command) error {
			fmt.Printf("\n+--------------------------\n")
			fmt.Println("| Constants")
			fmt.Println("+--------------------------")
			fmt.Printf("| CONST_PRD\n| %s\n|\n", env.CONST_PRD)
			fmt.Printf("| CONST_DEV\n| %s\n|\n", env.CONST_DEV)
			fmt.Printf("| CONST_BIN\n| %s\n|\n", env.CONST_BIN)
			fmt.Printf("| CONST_LOCAL\n| %s\n|\n", env.CONST_LOCAL)
			fmt.Printf("| CONST_CONF\n| %s\n|\n", env.CONST_CONF)
			fmt.Printf("| CONST_STATE\n| %s\n|\n", env.CONST_STATE)

			fmt.Printf("\n+--------------------------\n")
			fmt.Println("| Variables")
			fmt.Println("+--------------------------")
			fmt.Printf("| BASE\n| %s\n|\n", env.BASE)
			fmt.Printf("| HOME\n| %s\n|\n", env.HOME)
			fmt.Printf("| LOCAL_CLI\n| %s\n|\n", env.LOCAL_CLI)
			fmt.Printf("| BUILD_DIR\n| %s\n|\n", env.BUILD_DIR)
			fmt.Printf("| BUILD_CONF\n| %s\n|\n", env.BUILD_CONF)
			fmt.Printf("| BUILD_BIN\n| %s\n|\n", env.BUILD_BIN)
			fmt.Printf("| BUILD_STATE\n| %s\n|\n", env.BUILD_STATE)
			fmt.Printf("| OS_CONF\n| %s\n|\n", env.OS_CONF)
			fmt.Printf("| OS_CLI\n| %s\n|\n", env.OS_CLI)
			fmt.Printf("| OS_STATE\n| %s\n|\n", env.OS_STATE)
			fmt.Printf("| OS_BIN\n| %s\n|\n", env.OS_BIN)

			return nil
		},
	}
}
