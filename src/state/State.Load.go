package state

import (
	"os"
	"time"

	"github.com/jgttech/repo/src/assert"
	"github.com/jgttech/repo/src/env"
	"github.com/jgttech/repo/src/fs"
)

func (self *Conf) Load() {
	cwd := assert.Must(os.Getwd())
	file := env.BUILD_STATE
	self.File = fs.NewNode(file, fs.File)
	_, err := os.Stat(file)

	if os.IsNotExist(err) {
		timestamp := time.Now().Unix()

		self.CreatedAt = timestamp
		self.UpdatedAt = timestamp
		self.Home = cwd
		self.Backups = []*Backup{}
		self.Repositories = make(map[string]*Repository)

		self.Defaults = &Defaults{
			Branches: &DefaultsBranches{
				Primary: "main",
			},
		}

		self.History = &History{
			Max:     5,
			Entries: []string{},
		}

		self.File.Create()
		self.Save()
	} else {
		self.read(file)
	}
}
