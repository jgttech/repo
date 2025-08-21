package fs

func (node *INode) Sync() {
	nodeSync(node)
}
