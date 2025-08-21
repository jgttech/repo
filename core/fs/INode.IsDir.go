package fs

func (node *INode) IsDir() bool {
	return node.isDir
}
