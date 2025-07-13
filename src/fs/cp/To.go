package cp

func To(to string) option {
	return func(state *state) {
		state.to = to
	}
}
