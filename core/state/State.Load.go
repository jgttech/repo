package state

import (
	"github.com/jgttech/repo/core/fs"
)

func Load(node *fs.INode) (state *State) {
	state = &State{}

	if !node.Exists() {
		return
	}

	return
}
