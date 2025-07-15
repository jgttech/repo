package tar

func (self *Tar) Add(asset string) {
	self.Files = append(self.Files, asset)
}
