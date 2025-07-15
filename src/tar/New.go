package tar

func New(path string) (self *Tar) {
	self = &Tar{Path: path}
	return
}
