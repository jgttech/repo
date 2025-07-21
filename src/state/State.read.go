package state

import (
	"os"

	"github.com/jgttech/repo/src/assert"
	"github.com/jgttech/repo/src/fs"
	"gopkg.in/yaml.v3"
)

func (self *Conf) read(filepath string) (err error) {
	_, err = os.Stat(filepath)

	if os.IsNotExist(err) {
		return err
	}

	bytes := assert.Must(os.ReadFile(filepath))
	err = yaml.Unmarshal(bytes, self)

	if err == nil {
		self.File = fs.NewNode(filepath, fs.File)
		self.Home = os.ExpandEnv(self.Home)

		for _, backup := range self.Backups {
			backup.From = os.ExpandEnv(backup.From)
			backup.To = os.ExpandEnv(backup.To)
		}
	}

	return
}
