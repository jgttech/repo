package state

import "github.com/jgttech/repo/src/os"

func (self *Conf) AddBackup(from, to string) {
	from = os.ContractEnv(from)
	to = os.ContractEnv(to)

	self.updateTimestamp()
	self.Backups = append(self.Backups, &Backup{
		From: from,
		To:   to,
	})
}
