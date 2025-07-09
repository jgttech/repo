package self

import (
	"os"
	"time"

	"github.com/jgttech/repo/src/assert"
	"github.com/jgttech/repo/src/cfg"
)

func newState(ctx *cfg.Context) (s *cfg.State) {
	file := ctx.State
	s = &cfg.State{Path: file.Path}

	if file.Exists() {
		assert.Throws(s.Load(file.Path))
		return
	}

	os.MkdirAll(file.Base, 0700)
	file.Create()

	timestamp := time.Now().Unix()

	s.CreatedAt = cfg.Timestamp(timestamp)
	s.UpdatedAt = cfg.Timestamp(timestamp)
	s.Repositories = make(map[string]*cfg.StateRepository)
	s.Defaults = &cfg.StateDefaults{
		Branches: &cfg.StateDefaultsBranches{
			Primary: "main",
		},
	}
	s.History = &cfg.StateHistory{
		Max:     5,
		Entries: []string{},
	}

	s.Save()
	return
}
