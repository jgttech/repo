package cli

import (
	"os"

	"github.com/jgttech/repo/src/fs"
)

func File(path string) cliOption {
	var valid bool
	_, err := os.Stat(path)

	if err == nil {
		valid = true
	}

	return func(c *Conf) {
		if valid {
			c.node = fs.NewNode(path, fs.File)
		}
	}
}
