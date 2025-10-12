package env

import (
	"path/filepath"
	"repo/cli/core/node"
)

func GetInstallHome() (*node.Node, error) {
	return node.New(filepath.Join(GetHome(), ".repo"), node.AsDir)
}
