package fs

func (node *INode) Exists() bool {
	return node.exists
}
