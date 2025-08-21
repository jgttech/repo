package tmp

import (
	"os"
	"path/filepath"

	"github.com/jgttech/repo/core/assert"
)

// Accepts a string for the name of the
// file to create. This generates a temporary
// directory for the file and then creates the
// file and returns a path to the file.
func File(name string) (file string, err error) {
	dir, err := Dir()

	// Ensure that the file path the is returned
	// does not contain an invalid path if there
	// is an error. So ensure that, if there is
	// an error, that the file path is reset.
	defer func() {
		if err != nil {
			file = ""
		}
	}()

	if err != nil {
		return
	}

	_, err = os.Stat(dir)

	if os.IsNotExist(err) {
		return
	}

	tmpfile := filepath.Join(dir, name)
	tmp, err := os.Create(tmpfile)

	if err != nil {
		return
	}

	defer assert.Throws(tmp.Close())
	file = tmpfile

	return
}
