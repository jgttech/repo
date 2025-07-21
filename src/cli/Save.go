package cli

import (
	_os "github.com/jgttech/repo/src/os"
	"gopkg.in/yaml.v3"
	"os"
)

func (self *Conf) Save() (err error) {
	self.Exports = _os.ContractEnv(self.Exports)
	filepath := self.node.Path

	bytes, err := yaml.Marshal(self)

	if err != nil {
		return
	}

	err = os.WriteFile(filepath, bytes, 0700)
	return
}
