package fs

import (
	"fmt"
	"os"
)

func (self *Node) Remove() (err error) {
	if _, err = os.Stat(self.Path); err != nil {
		err = fmt.Errorf("File does not exist: %s", self.Path)
		return
	}

	err = os.Remove(self.Path)
	return
}
