package state

func (self *Conf) AddBackup(from, to string) {
	self.updateTimestamp()
	self.Backups = append(self.Backups, &Backup{
		From: from,
		To:   to,
	})
}
