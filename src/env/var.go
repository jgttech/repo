package env

import (
	"os"
	"path"

	"github.com/jgttech/repo/src/assert"
)

var (
	HOME       = assert.Must(os.Getwd())
	BASE_DIR   = path.Join(HOME, REPO_DIR)
	BASE_CONF  = path.Join(BASE_DIR, REPO_CONF)
	BASE_CLI   = path.Join(BASE_CONF, REPO_CLI)
	BASE_STATE = path.Join(BASE_CONF, REPO_STATE)
)
