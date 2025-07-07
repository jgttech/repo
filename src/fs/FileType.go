package fs

import "github.com/jgttech/repo/src/env"

func FileType(node *Node) {
	node.Type = env.FILE_NODE
}
