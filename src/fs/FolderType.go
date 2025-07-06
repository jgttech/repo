package fs

import "github.com/jgttech/repo/src/env"

func FolderType(n *Node) {
	n.Type = env.FOLDER_NODE
}
