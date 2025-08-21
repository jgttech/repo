package fs

import (
	"os"
	"path/filepath"

	"github.com/jgttech/repo/core/log"
)

// Checks the node.path and synchronizes the
// status with the current node reference.
func nodeSync(node *INode) {
	stat, err := os.Stat(node.path)

	if err != nil && !os.IsNotExist(err) {
		log.Fatalln(err)
	}

	node.dir = filepath.Dir(node.path)
	node.exists = !os.IsNotExist(err)

	if stat != nil {
		node.name = stat.Name()
		node.size = stat.Size()
		node.isDir = stat.IsDir()
		node.mode = stat.Mode()
		node.modified = stat.ModTime()
	}
}
