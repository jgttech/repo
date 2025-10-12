package node

import "os"

func (node *Node) Exists() bool {
	_, err := os.Stat(node.path)
	return !os.IsNotExist(err)
}
