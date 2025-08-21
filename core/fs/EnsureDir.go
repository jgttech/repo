package fs

import (
	"os"
	"path/filepath"

	"github.com/jgttech/repo/core/log"
)

func EnsureDir(node *INode) {
	if node.IsDir() {
		return
	}

	base := filepath.Dir(node.path)
	_, err := os.Stat(base)

	if os.IsNotExist(err) {
		log.Throws(os.MkdirAll(base, 0755))
	}
}
