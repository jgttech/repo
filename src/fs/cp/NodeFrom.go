package cp

import "github.com/jgttech/repo/src/fs"

func NodeFrom(from *fs.Node) option {
	return func(state *state) {
		state.from = from.Path
	}
}
