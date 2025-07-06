package fs

import "github.com/jgttech/repo/src/env"

func FileType(n *Node) {
	n.Type = env.FILE_NODE
}
