package archive

func To(to string) archiveOption {
	return func(a *Archive) {
		a.To = to
	}
}
