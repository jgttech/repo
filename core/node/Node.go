package node

import "os"

type NodeType int

const (
	NODE_FILE NodeType = iota
	NODE_DIR
)

type Node struct {
	nodeType NodeType
	path     string
	stat     os.FileInfo
}
