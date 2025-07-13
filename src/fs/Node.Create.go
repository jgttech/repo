package fs

import (
	"github.com/jgttech/repo/src/assert"
	"os"
)

func (n *Node) Create() (success bool) {
	switch n.Type {
	case TYPE_FILE:
		success = true
		file := assert.Must(os.Create(n.Path))
		file.Close()
	case TYPE_FOLDER:
		success = true
		os.MkdirAll(n.Path, 0700)
	default:
		success = false
	}
	return
}
