package cli

import "github.com/jgttech/repo/src/fs"

type Cli struct {
	node         *fs.Node `yaml:"-"`
	Version      string   `yaml:"version"`
	Export       string   `yaml:"export"`
	Dependencies []string `yaml:"dependencies"`
}
