package cp

import "github.com/jgttech/repo/src/fs"

func NodeTo(to *fs.Node) option {
	return func(state *state) {
		state.to = to.Path
	}
}
