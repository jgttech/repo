package archive

func From(from string) archiveOption {
	return func(a *Archive) {
		a.From = from
	}
}
