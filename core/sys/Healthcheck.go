package sys

import (
	"fmt"
	"github.com/jgttech/repo/core/fs/node"
)

// The goal of the healthcheck is to ensure that
// all system requirements are met before running
// the CLI. This means that app dependencies are
// installed and the CLI checks it's own install
// status. Like an initial diagnostic check before
// doing anything.
func Healthcheck() {
	baseDir := node.BaseDir
	stateFile := node.StateFile

	// If you don't have a config and a state
	// file, then we can't check anything.
	if !baseDir.Exists() || !stateFile.Exists() {
		return
	}

	fmt.Println("HEALTH CHECK!!!")
}
