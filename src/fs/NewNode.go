package fs

import (
	"os"
	"path"
	"strings"

	"github.com/jgttech/repo/src/env"
)

type nodeOption func(*Node)

func NewNode(nodePath string, opts ...nodeOption) (node *Node) {
	node = &Node{}

	if nodePath == "" {
		node.Base = "/"
		node.Path = "/"
	} else {
		node.Path = os.ExpandEnv(nodePath)

		tokens := strings.Split(nodePath, "/")
		size := len(tokens)
		name := tokens[size-1]

		node.Name = name
		node.Base = os.ExpandEnv(path.Join(tokens[0 : size-1]...))

		if strings.Contains(name, ".") {
			tokens := strings.Split(name, ".")
			size := len(tokens)
			ext := tokens[size-1]

			node.Ext = ext
		}

		if !node.Exists() {
			node.Type = env.FILE_NODE
		}
	}

	for _, fn := range opts {
		fn(node)
	}

	return
}
