package fs

import "github.com/jgttech/repo/src/env"

func FolderType(node *Node) {
	node.Type = env.FOLDER_NODE
}
