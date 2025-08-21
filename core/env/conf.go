package env

import "path/filepath"

const (
	VERSION  = "0.0.1"
	UPLOAD   = "none"
	DOWNLOAD = "none"
)

var (
	CLI_ARCHIVES = filepath.Join("$HOME", BASE_NAME, "archives")

	IMPORT       = CLI_ARCHIVES
	EXPORT       = CLI_ARCHIVES
	DEPENDENCIES = []string{
		"stow",
		"git",
		"gh",
	}
)
