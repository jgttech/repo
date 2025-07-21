package _import

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/jgttech/repo/src/archive"
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

			if _, err = os.Stat(target); os.IsNotExist(err) {
				err = errors.Errorf("File does not exist: %s", target)
				return
			}

			base, err := os.MkdirTemp(os.TempDir(), "repocli-")

			if err != nil {
				return
			}

			archive.Extract(
				archive.From(target),
				archive.To(base),
			)

			fmt.Println(base)
			return
		},
	}
}
