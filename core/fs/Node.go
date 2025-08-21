package fs

import "slices"

type INodeOption func(node *INode)

func Node(path string, options ...INodeOption) (node *INode) {
	node = &INode{path: path}

	nodeSync(node)

	for option := range slices.Values(options) {
		option(node)
	}

	return
}
