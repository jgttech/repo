package fs

import (
	"fmt"
	"os"

	"github.com/jgttech/repo/src/assert"
)

func (n *Node) Create() (err error) {
	switch n.Type {
	case TYPE_FILE:
		file := assert.Must(os.Create(n.Path))
		file.Close()
	case TYPE_FOLDER:
		os.MkdirAll(n.Path, 0755)
	default:
		err = fmt.Errorf("Failed to create node: '%s'", n.Path)
	}
	return
}
