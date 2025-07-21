package cp

import (
	"github.com/jgttech/repo/src/errors"
	cp "github.com/otiai10/copy"
	"os"
)

func Dir(opts ...option) (err error) {
	data := &state{}

	for _, opt := range opts {
		opt(data)
	}

	from, fromErr := os.Stat(data.from)
	to, toErr := os.Stat(data.to)

	if os.IsNotExist(fromErr) {
		err = fromErr
		return
	}

	if os.IsNotExist(toErr) {
		err = os.MkdirAll(data.to, 0755)
		return
	}

	if from != nil && !from.IsDir() {
		err = errors.Errorf("Target must be a directory:\n%s", data.from)
		return
	}

	if to != nil && !to.IsDir() {
		err = errors.Errorf("Target must be a file, not a directory:\n%s", data.to)
		return
	}

	err = cp.Copy(data.from, data.to)
	return
}
