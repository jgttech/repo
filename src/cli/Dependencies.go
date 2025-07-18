package cli

import (
	"fmt"

	"github.com/jgttech/repo/src/assert"
	"github.com/jgttech/repo/src/os"
	"github.com/urfave/cli/v3"
)

func Dependencies(cmd *cli.Command) *cli.Command {
	conf := assert.Must(New())
	fmt.Println("Checking dependencies, please wait...")

	for _, dep := range conf.Dependencies {
		if !os.PackageInstalled(dep) {
			fmt.Printf("Missing: %s", dep)
		}
	}

	fmt.Println("Done.")

	return cmd
}
