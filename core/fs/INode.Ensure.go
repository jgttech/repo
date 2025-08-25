package fs

import (
	"os"
	"path/filepath"

	"github.com/jgttech/repo/core/log"
)

func (node *INode) Ensure() {
	if node.Exists() {
		return
	}

	// Create the base directory for whatever
	// the node is. If the node is a directory
	// this will ensure that its parent exists.
	base := filepath.Dir(node.path)
	_, err := os.Stat(base)

	if os.IsNotExist(err) {
		log.Throws(os.MkdirAll(base, 0755))
	}

	perm := node.metadata.permission

	if node.metadata.asDir {
		log.Throws(os.MkdirAll(node.path, perm))
	} else {
		file, err := os.Create(node.path)

		if err != nil {
			log.Fatalln(err)
		}

		file.Close()
	}
}
