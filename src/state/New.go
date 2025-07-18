package state

import (
	"os"
	"time"

	"github.com/jgttech/repo/src/assert"
	"github.com/jgttech/repo/src/env"
	"github.com/jgttech/repo/src/fs"
)

func New() (state *State) {
	cwd := assert.Must(os.Getwd())
	file := env.BUILD_STATE
	state = &State{File: fs.NewNode(file, fs.File)}
	_, err := os.Stat(file)

	if os.IsNotExist(err) {
		timestamp := time.Now().Unix()

		state.CreatedAt = timestamp
		state.UpdatedAt = timestamp
		state.Home = cwd
		state.Backups = []*Backup{}
		state.Repositories = make(map[string]*Repository)

		state.Defaults = &Defaults{
			Branches: &DefaultsBranches{
				Primary: "main",
			},
		}

		state.History = &History{
			Max:     5,
			Entries: []string{},
		}

		state.File.Create()
		state.Save()
	} else {
		state.Load(file)
	}

	return
}
