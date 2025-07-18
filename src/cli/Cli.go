package cli

import "github.com/jgttech/repo/src/fs"

type Cli struct {
	node         *fs.Node `yaml:"-"`
	Version      string   `yaml:"version"`
	Exports      string   `yaml:"exports"`
	Dependencies []string `yaml:"dependencies"`
}
