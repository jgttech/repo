package cli

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/jgttech/repo/src/env"
	"github.com/jgttech/repo/src/os"
)

func (self *Conf) HasDependencies() (err error) {
	missing := []string{}

	for _, dep := range self.Dependencies {
		if !os.PackageInstalled(dep) {
			missing = append(missing, dep)
		}
	}

	if len(missing) > 0 {
		msg := fmt.Sprintf("\n\n|\n| Missing (%d) required dependencies:\n", len(missing))

		for _, pkg := range missing {
			msg += fmt.Sprintf("| - %s\n", pkg)
		}

		msg += fmt.Sprintf(
			"|\n"+
				"| ERROR\n"+
				"| Please install required dependencies and\n"+
				"| run 'bash ~/%s/install' to complete\n"+
				"| the installation\n"+
				"|\n\n",
			filepath.Base(env.BASE),
		)

		err = errors.New(msg)
	}

	return
}
