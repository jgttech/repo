package env

import (
	"os"
	"path/filepath"
)

var (
	OS_HOME  = os.Getenv("HOME")
	OS_BASE  = filepath.Join(OS_HOME, BASE_NAME)
	OS_STATE = filepath.Join(OS_BASE, STATE_NAME)
)
