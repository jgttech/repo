package cli

import "github.com/jgttech/repo/src/fs"

type cliOption func(*Conf)

type Conf struct {
	node         *fs.Node `yaml:"-"`
	Version      string   `yaml:"version"`
	Exports      string   `yaml:"exports"`
	Dependencies []string `yaml:"dependencies"`
}
