package fs

func AsDir(node *INode) {
	node.metadata.asDir = true
	node.metadata.permission = 0755
}
