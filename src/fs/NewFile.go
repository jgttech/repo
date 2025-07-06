package fs

func NewFile(nodePath string, opts ...nodeOption) (node *Node) {
	node = NewNode(nodePath, FileType)
	return
}
