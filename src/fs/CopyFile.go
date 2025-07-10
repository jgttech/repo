package fs

import (
	"fmt"
	"io"
	"os"

	"github.com/jgttech/repo/src/assert"
)

func CopyFile(from, to string) (err error) {
	_, err = os.Stat(to)

	if err == nil {
		err = fmt.Errorf("File already exists: %s", to)
		return
	}

	src := assert.Must(os.Open(from))
	defer src.Close()

	dst := assert.Must(os.Create(to))
	defer dst.Close()

	_, err = io.Copy(dst, src)

	if err != nil {
		err = fmt.Errorf("Failed to copy file: %s", from)
	}

	err = dst.Sync()
	return
}
