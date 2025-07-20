package _import

import (
	"context"
	"fmt"

	v3 "github.com/urfave/cli/v3"
)

func Command() *v3.Command {
	return &v3.Command{
		Name: "import",
		Action: func(ctx context.Context, c *v3.Command) (err error) {
			dir := c.Args().First()

			if dir == "" {
				fmt.Println("\n|")
				fmt.Println("| ERROR")
				fmt.Println("| Missing the required argument pointing")
				fmt.Printf("| to the .tar.gz archive.\n")
				fmt.Println("|")

				return
			}

			return
		},
	}
}
