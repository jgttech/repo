package tmp

import (
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/jgttech/repo/core/env"
)

func Dir() (dir string, err error) {
	tmp := os.TempDir()
	timestamp := time.Now().UnixNano()
	dir = filepath.Join(tmp, env.TMP_NAME, strconv.Itoa(int(timestamp)))
	_, err = os.Stat(dir)

	if !os.IsNotExist(err) {
		panic(err)
	}

	err = os.MkdirAll(dir, 0755)
	return
}
