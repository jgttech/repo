package core

import (
	"log"
	"path/filepath"
	"repo/cli/core/node"
)

func GetInstallHome() *node.Node {
	home, err := node.New(filepath.Join(GetHome(), ".repo"))

	if err != nil {
		log.Fatalln(err)
	}

	return home
}
