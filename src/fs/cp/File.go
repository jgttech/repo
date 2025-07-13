package cp

import (
	"fmt"
	"github.com/jgttech/repo/src/assert"
	"io"
	"os"
)

func File(opts ...option) (err error) {
	data := &state{}

	for _, opt := range opts {
		opt(data)
	}

	fmt.Printf("%#v\n", data)

	from, fromErr := os.Stat(data.from)
	to, toErr := os.Stat(data.to)

	if os.IsNotExist(fromErr) {
		err = fromErr
		return
	}

	if toErr == nil {
		err = fmt.Errorf("File already exists: %s", data.to)
		return
	}

	if from != nil && from.IsDir() {
		err = fmt.Errorf("Target (%s) must be a file, not a directory.", data.from)
		return
	}

	if to != nil && to.IsDir() {
		err = fmt.Errorf("Target (%s) must be a file, not a directory.", data.to)
		return
	}

	src := assert.Must(os.Open(data.from))
	defer src.Close()

	dst := assert.Must(os.Create(data.to))
	defer dst.Close()

	_, err = io.Copy(dst, src)

	if err != nil {
		err = fmt.Errorf("Failed to copy file: %s", from)
	}

	err = dst.Sync()
	return
}
