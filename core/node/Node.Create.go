package node

import "os"

func (node *Node) Create() error {
	if node.nodeType == NODE_DIR {
		err := os.MkdirAll(node.path, 0755)
		return err
	} else {
		file, err := os.Create(node.path)

		if err != nil {
			return err
		}

		err = file.Close()
		return err
	}
}
