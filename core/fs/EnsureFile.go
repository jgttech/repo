package fs

import (
	"github.com/jgttech/repo/core/log"
	"os"
)

func EnsureFile(node *INode) {
	EnsureDir(node)

	file := log.Must(os.Create(node.path))

	if err := file.Close(); err != nil {
		panic(err)
	}
}
