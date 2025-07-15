package tar

import (
	"fmt"
	"os"
)

func (self *Tar) Create() (err error) {
	_, err = os.Stat(self.Path)

	if err == nil {
		err = fmt.Errorf("Archive already exists: %s\n", self.Path)
		return
	}

	file, err := os.Create(self.Path)

	if file == nil {
		err = fmt.Errorf("File not found: %s\n", self.Path)
		return
	}

	defer file.Close()

	for _, asset := range self.Files {
		fmt.Println("ASSET:", asset)
	}

	return
}
