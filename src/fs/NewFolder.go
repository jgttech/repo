package fs

func NewFolder(nodePath string, opts ...nodeOption) (node *Node) {
	node = NewNode(nodePath, FolderType)
	return
}
