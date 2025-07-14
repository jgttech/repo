package state

import (
	"os"

	"github.com/jgttech/repo/src/assert"
	"gopkg.in/yaml.v3"
)

func (self *State) Save() {
	self.updateTimestamp()
	filepath := self.File.Path
	bytes := assert.Must(yaml.Marshal(self))
	assert.Throws(os.WriteFile(filepath, bytes, 0700))
}
