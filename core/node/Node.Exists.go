package node

import "os"

func (node *Node) Exists() bool {
	_, err := os.Stat(node.path)

	if os.IsNotExist(err) {
		return false
	}

	return true
}
