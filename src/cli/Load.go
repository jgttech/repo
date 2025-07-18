package cli

import (
	"os"

	"github.com/jgttech/repo/src/fs"
	"gopkg.in/yaml.v3"
)

func (self *Conf) Load(options ...cliOption) (err error) {
	for _, fn := range options {
		fn(self)
	}

	file := self.node.Path
	bytes, err := os.ReadFile(file)
	err = yaml.Unmarshal(bytes, self)

	self.node = fs.NewNode(file)
	self.Exports = os.ExpandEnv(self.Exports)

	return
}
