package env

import (
	"github.com/jgttech/repo/src/assert"
	"os"
	"path"
)

const (
	// Mode types.
	REPO_MODE_PROD = 0
	REPO_MODE_DEV  = 1

	// Directory where things are installed.
	REPO_DIR = ".repocli"
)

var (
	REPO_HOME_PROD = path.Join(os.Getenv("HOME"), REPO_DIR)
	REPO_HOME_DEV  = path.Join(assert.Must(os.Getwd()), REPO_DIR)
	REPO_HOME      = (func() string {
		mode := os.Getenv("REPO_MODE")

		if mode == "0" {
			return REPO_HOME_PROD
		}

		return REPO_HOME_DEV
	})()
)
