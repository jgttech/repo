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
			fmt.Printf("| REPO_PRD\n| %d\n|\n", env.REPO_PRD)
			fmt.Printf("| REPO_DEV\n| %d\n|\n", env.REPO_DEV)
			fmt.Printf("| REPO_DIR\n| %s\n|\n", env.REPO_DIR)
			fmt.Printf("| REPO_CONF\n| %s\n|\n", env.REPO_CONF)
			fmt.Printf("| REPO_CLI\n| %s\n|\n", env.REPO_CLI)
			fmt.Printf("| REPO_STATE\n| %s\n|\n", env.REPO_STATE)

			fmt.Printf("\n+--------------------------\n")
			fmt.Println("| Variables")
			fmt.Println("+--------------------------")
			fmt.Printf("| HOME\n| %s\n|\n", env.HOME)
			fmt.Printf("| USER\n| %s\n|\n", env.USER)
			fmt.Printf("| BASE_DIR\n| %s\n|\n", env.BASE_DIR)
			fmt.Printf("| BASE_CONF\n| %s\n|\n", env.BASE_CONF)
			fmt.Printf("| BASE_CLI\n| %s\n|\n", env.BASE_CLI)
			fmt.Printf("| BASE_STATE\n| %s\n|\n", env.BASE_STATE)
			fmt.Printf("| OS_CONF\n| %s\n|\n", env.OS_CONF)

			return nil
		},
	}
}
