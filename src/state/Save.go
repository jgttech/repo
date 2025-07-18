package state

import (
	"os"

	"gopkg.in/yaml.v3"
)

func (self *State) Save() (err error) {
	self.updateTimestamp()
	filepath := self.File.Path

	bytes, err := yaml.Marshal(self)

	if err != nil {
		return
	}

	err = os.WriteFile(filepath, bytes, 0700)
	return
}
