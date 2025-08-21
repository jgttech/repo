package env

import (
	"os"
	"path/filepath"
)

var (
	// MODE = os.Getenv("REPO_MODE")

	OS_HOME  = os.Getenv("HOME")
	OS_BASE  = filepath.Join(OS_HOME, BASE_NAME)
	OS_STATE = filepath.Join(OS_BASE, STATE_NAME)
)

// import (
// 	"os"
// 	"path"
// 	"path/filepath"
//
// 	"github.com/jgttech/repo/core/assert"
// )
//
// var (
// 	SEP  = string(filepath.Separator)
// 	MODE = os.Getenv("REPO_MODE")
//
// 	// CLI & User specific locations.
// 	HOME = os.Getenv("HOME")
// 	BASE = assert.Must(os.Getwd())
//
// 	// Where in the local repo, things are located.
// 	LOCAL_CLI = path.Join(BASE, CONST_CLI)
//
// 	// Where in the build, things are located.
// 	BUILD_DIR   = path.Join(BASE, CONST_BUILD)
// 	BUILD_CONF  = path.Join(BUILD_DIR, CONST_CONFIG)
// 	BUILD_BIN   = path.Join(BUILD_DIR, CONST_LOCAL, CONST_BIN)
// 	BUILD_CLI   = path.Join(BUILD_CONF, CONST_CLI)
// 	BUILD_STATE = path.Join(BUILD_CONF, CONST_STATE)
//
// 	// Where on the OS, things are located.
// 	OS_CONF  = path.Join(HOME, CONST_CONFIG)
// 	OS_CLI   = path.Join(OS_CONF, CONST_CLI)
// 	OS_STATE = path.Join(OS_CONF, CONST_STATE)
// 	OS_BIN   = path.Join(HOME, CONST_LOCAL, CONST_BIN)
// )
