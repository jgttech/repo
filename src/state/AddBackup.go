package state

func (self *State) AddBackup(from, to string) {
	self.updateTimestamp()
	self.Backups = append(self.Backups, &Backup{
		From: from,
		To:   to,
	})
}
