package fs

import (
	"os"

	"github.com/jgttech/repo/src/assert"
	"github.com/jgttech/repo/src/env"
)

type Node struct {
	Name string
	Path string
	Base string
	Ext  string
	Type env.FsNode
}

func (node *Node) Exists() (exists bool) {
	_, err := os.Stat(node.Path)

	if os.IsNotExist(err) {
		exists = false
	} else if err == nil {
		exists = true
	}

	return
}

func (node *Node) Create() (success bool) {
	switch node.Type {
	case env.FILE_NODE:
		success = true
		file := assert.Must(os.Create(node.Path))
		file.Close()
	case env.FOLDER_NODE:
		success = true
		os.MkdirAll(node.Path, 0700)
	default:
		success = false
	}
	return
}
