package fs

import (
	"os"

	"github.com/jgttech/repo/src/assert"
)

func (n *Node) Create() (success bool) {
	switch n.Type {
	case TYPE_FILE:
		success = true
		file := assert.Must(os.Create(n.Path))
		file.Close()
	case TYPE_FOLDER:
		success = true
		os.MkdirAll(n.Path, 0755)
	default:
		success = false
	}
	return
}
