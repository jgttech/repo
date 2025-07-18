package env

import (
	"os"
	"path"

	"github.com/jgttech/repo/src/assert"
)

var (
	BASE_DIR   = path.Join(assert.Must(os.Getwd()), REPO_DIR)
	BASE_CONF  = path.Join(BASE_DIR, REPO_CONF)
	BASE_CLI   = path.Join(BASE_CONF, REPO_CLI)
	BASE_STATE = path.Join(BASE_CONF, REPO_STATE)
)
