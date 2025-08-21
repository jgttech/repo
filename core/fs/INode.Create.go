package fs

import "os"

func (node *INode) Create() (err error) {
	if node.isDir {
		err = os.MkdirAll(node.path, 0755)
	} else {
		file, createErr := os.Create(node.path)

		if createErr != nil {
			return createErr
		}

		err = file.Close()
	}

	return
}
