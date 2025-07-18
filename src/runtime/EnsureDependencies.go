package runtime

import (
	"fmt"

	"github.com/jgttech/repo/src/assert"
	"github.com/jgttech/repo/src/cli"
	"github.com/jgttech/repo/src/os"
)

func EnsureDependencies() {
	conf := assert.Must(cli.New())
	fmt.Println("Checking dependencies, please wait...")

	for _, dep := range conf.Dependencies {
		if !os.PackageInstalled(dep) {
			fmt.Printf("Missing: %s", dep)
		}
	}

	fmt.Println("Done.")
}
