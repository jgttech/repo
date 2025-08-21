package fs

import "os"

func (node *INode) Remove() (err error) {
	if node.isDir {
		err = os.RemoveAll(node.path)
	} else {
		err = os.Remove(node.path)
	}

	return
}
