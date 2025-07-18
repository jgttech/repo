package os

import (
	"os"

	"github.com/jgttech/repo/src/env"
)

func SystemInstalled() (installed bool) {
	missing := []string{}
	targets := []string{
		env.OS_BIN,
		env.OS_CONF,
		env.OS_CLI,
		env.OS_STATE,
	}

	for _, target := range targets {
		_, err := os.Stat(target)

		if os.IsNotExist(err) {
			missing = append(missing, target)
		}
	}

	installed = len(missing) == 0
	return
}
