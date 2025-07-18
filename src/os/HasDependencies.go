package os

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/jgttech/repo/src/cli"
	"github.com/jgttech/repo/src/env"
)

func HasDependencies(conf *cli.Conf) (err error) {
	missing := []string{}

	for _, dep := range conf.Dependencies {
		if !PackageInstalled(dep) {
			missing = append(missing, dep)
		}
	}

	if len(missing) > 0 {
		msg := fmt.Sprintf("\n|\n| Missing (%d) required dependencies:\n", len(missing))

		for _, pkg := range missing {
			msg += fmt.Sprintf("| - %s\n", pkg)
		}

		msg += fmt.Sprintf("\n\n"+
			"|\n"+
			"| ERROR\n"+
			"| Please install required dependencies and\n"+
			"| run 'bash ~/%s/install' to complete\n"+
			"| the installation\n"+
			"|\n",
			filepath.Base(env.BASE),
		)

		err = errors.New(msg)
	}

	return
}
