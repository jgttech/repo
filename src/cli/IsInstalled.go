package cli

import (
	"os"

	"github.com/jgttech/repo/src/env"
)

func IsInstalled() bool {
	_, err := os.Stat(env.BASE_DIR)
	return err == nil
}
