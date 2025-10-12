package node

import (
	"os"
	"slices"
)

func New(source string, options ...nodeOption) (*Node, error) {
	var err error
	node := &Node{}

	node.nodeType = NODE_FILE
	node.path = source
	node.stat, err = os.Stat(source)

	if err == nil && node.stat.IsDir() {
		node.nodeType = NODE_DIR
	}

	for option := range slices.Values(options) {
		option(node)
	}

	return node, err
}
