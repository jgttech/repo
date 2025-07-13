package cp

func From(from string) option {
	return func(state *state) {
		state.from = from
	}
}
