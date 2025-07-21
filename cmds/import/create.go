package _import

import (
	// "fmt"
	"os"
	// "path/filepath"

	"github.com/jgttech/repo/src/env"
	// "github.com/jgttech/repo/src/exec"
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

	// TODO
	// Need to load the state and cli configs from the extracted archive.

	// TODO
	// Then need to iterate through the backups in the state config and
	// re-create the backup files. Consider generating new backup files
	// with updated timestamps.

	// TODO
	// Ensure the backup files that are re-created delete the originals
	// so the the symlink works.

	// dir := filepath.Join(env.BUILD_DIR, "..")
	// arg := fmt.Sprintf("stow -t %s %s", env.HOME, env.CONST_DIR)
	// cmd := exec.Cmd(arg, exec.Dir(dir), exec.Stdio)
	// err = cmd.Run()

	return
}
