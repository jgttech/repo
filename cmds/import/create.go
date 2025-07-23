package _import

import (
	"os"

	"github.com/jgttech/repo/src/env"
	"github.com/jgttech/repo/src/fs/cp"
)

// This gets run when there is not pre-existing install
// and, effectively, does a standard import of the config
// which is, basically, a new install.
func create(base string) (err error) {
	if err = os.MkdirAll(env.BUILD_DIR, 0755); err != nil {
		return
	}

	if err = cp.Dir(cp.From(base), cp.To(env.BUILD_DIR)); err == nil {
		err = os.RemoveAll(base)
	}

	return
}
