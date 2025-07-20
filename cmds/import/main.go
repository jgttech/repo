package _import

import (
	"context"
	"strings"
	// "fmt"
	// "os"
	//
	// "github.com/jgttech/repo/src/archive"
	"github.com/jgttech/repo/src/errors"
	v3 "github.com/urfave/cli/v3"
)

func Command() *v3.Command {
	return &v3.Command{
		Name: "import",
		Action: func(ctx context.Context, c *v3.Command) (err error) {
			target := c.Args().First()

			if target == "" {
				err = errors.New(strings.Join([]string{
					"Missing the required argument pointing",
					"to the .tar.gz archive.",
				}, "\n"))

				return
			}

			// if _, err := os.Stat(target); os.IsNotExist(err) {
			// 	err = fmt.Errorf("The target file does not exist")
			// 	return
			// }

			// archive.Extract(
			//   archive.From(dir),
			//   archive.To()
			// )

			return
		},
	}
}
