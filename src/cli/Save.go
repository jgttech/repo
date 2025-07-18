package cli

import (
	"gopkg.in/yaml.v3"
	"os"
)

func (self *Conf) Save() (err error) {
	filepath := self.node.Path

	bytes, err := yaml.Marshal(self)

	if err != nil {
		return
	}

	err = os.WriteFile(filepath, bytes, 0700)
	return
}
