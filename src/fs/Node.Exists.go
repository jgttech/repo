package fs

import "os"

func (n *Node) Exists() (exists bool) {
	_, err := os.Stat(n.Path)

	if os.IsNotExist(err) {
		exists = false
	} else if err == nil {
		exists = true
	}

	return
}
