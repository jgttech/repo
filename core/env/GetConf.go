package env

import (
	"path/filepath"
	"repo/cli/core/node"
)

func GetConf() (*node.Node, error) {
	home, err := GetInstallHome()

	if err != nil {
		return &node.Node{}, err
	}

	source := filepath.Join(home.GetPath(), "repo.yml")
	return node.New(source, node.AsFile)
}
