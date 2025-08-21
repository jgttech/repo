package log

import (
	"os"
	"path/filepath"
)

func Cleanup() {
	tmpdir := filepath.Dir(tmpfile)
	stat, err := os.Stat(tmpdir)

	if os.IsNotExist(err) {
		return
	}

	if stat.IsDir() {
		if err := os.RemoveAll(tmpdir); err != nil {
			panic(err)
		}
	} else {
		if err := os.Remove(tmpdir); err != nil {
			panic(err)
		}
	}
}
