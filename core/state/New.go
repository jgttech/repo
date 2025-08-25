package state

import (
	"time"

	"github.com/jgttech/repo/core/env"
)

// Creates a new instanes of a state struct.
// This does not load an existsing state struct.
// To load a state struct from a file, use the
// "Load" receiver function.
func New() (state *State) {
	createdAt := time.Now().UnixNano()
	updatedAt := createdAt
	version := env.VERSION

	state = &State{
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
		Version:      version,
		Import:       env.IMPORT,
		Export:       env.EXPORT,
		Upload:       env.UPLOAD,
		Download:     env.DOWNLOAD,
		Dependencies: env.DEPENDENCIES,
		Base:         env.OS_BASE,
		Repositories: []*Repository{},
	}

	return
}
