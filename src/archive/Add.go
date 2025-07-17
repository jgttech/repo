package archive

func (self *Archive) Add(asset string) {
	self.Files = append(self.Files, asset)
}
