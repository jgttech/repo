package fs

import "os"

func (node *INode) Mode() os.FileMode {
	return node.mode
}
