package state

import (
	"github.com/jgttech/repo/src/fs"
)

type Timestamp int64

type Conf struct {
	File         *fs.Node               `yaml:"-"`
	CreatedAt    int64                  `yaml:"created_at"`
	UpdatedAt    int64                  `yaml:"updated_at"`
	Home         string                 `yaml:"home"`
	Backups      []*Backup              `yaml:"backups"`
	Repositories map[string]*Repository `yaml:"repositories"`
	Defaults     *Defaults              `yaml:"defaults"`
	History      *History               `yaml:"history"`
}

type Backup struct {
	From string `yaml:"from"`
	To   string `yaml:"to"`
}

type Repository struct {
	CreatedAt int64  `yaml:"created_at"`
	UpdatedAt int64  `yaml:"updated_at"`
	Dir       string `yaml:"dir"`
	Name      string `yaml:"name"`
	Alias     string `yaml:"alias"`
	Primary   string `yaml:"primary"`
	Branch    string `yaml:"branch"`
	Ssh       string `yaml:"ssh"`
	Config    string `yaml:"config"`
}

type Defaults struct {
	Branches *DefaultsBranches `yaml:"branches"`
}

type DefaultsBranches struct {
	Primary string `yaml:"primary"`
}

type History struct {
	Max     int      `yaml:"max"`
	Entries []string `yaml:"entries"`
}
