package cli

import (
	"os"

	"github.com/jgttech/repo/src/assert"
	"gopkg.in/yaml.v3"
)

func (self *Conf) Save() {
	filepath := self.node.Path
	bytes := assert.Must(yaml.Marshal(self))
	assert.Throws(os.WriteFile(filepath, bytes, 0700))
}
