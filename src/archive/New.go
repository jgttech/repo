package archive

func New(path string) (self *Archive) {
	self = &Archive{Path: path}
	return
}
