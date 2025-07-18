package env

import (
	"os"
	"path"

	"github.com/jgttech/repo/src/assert"
)

var (
	// Determine the runtime mode.
	MODE = (func() string {
		mode := os.Getenv("REPO_MODE")

		if mode == "" {
			mode = CONST_PRD
		}

		return mode
	})()

	// CLI & User specific locations.
	BASE = assert.Must(os.Getwd())
	HOME = os.Getenv("HOME")

	// Where in the local repo, things are located.
	LOCAL_CLI = path.Join(BASE, CONST_CLI)

	// Where in the build, things are located.
	BUILD_DIR   = path.Join(BASE, CONST_DIR)
	BUILD_CONF  = path.Join(BUILD_DIR, CONST_CONF)
	BUILD_BIN   = path.Join(BUILD_DIR, CONST_LOCAL, CONST_BIN)
	BUILD_CLI   = path.Join(BUILD_CONF, CONST_CLI)
	BUILD_STATE = path.Join(BUILD_CONF, CONST_STATE)

	// Where on the OS, things are located.
	OS_CONF  = path.Join(HOME, CONST_CONF)
	OS_CLI   = path.Join(OS_CONF, CONST_CLI)
	OS_STATE = path.Join(OS_CONF, CONST_STATE)
	OS_BIN   = path.Join(HOME, CONST_LOCAL, CONST_BIN)
)
