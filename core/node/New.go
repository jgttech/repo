package node

import "os"

func New(source string) (*Node, error) {
	var err error
	node := &Node{}

	node.nodeType = NODE_FILE
	node.path = source
	node.stat, err = os.Stat(source)

	if node.stat.IsDir() {
		node.nodeType = NODE_DIR
	}

	if os.IsNotExist(err) {
		return nil, err
	}

	return node, nil
}
