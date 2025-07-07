package cfg

import (
	"os"

	"github.com/jgttech/repo/src/assert"
	"gopkg.in/yaml.v3"
)

type Timestamp int64

type State struct {
	CreatedAt    Timestamp                  `yaml:"created_at"`
	UpdatedAt    Timestamp                  `yaml:"updated_at"`
	Repositories map[string]StateRepository `yaml:"repositories"`
	Defaults     *StateDefaults             `yaml:"defaults"`
	History      *StateHistory              `yaml:"history"`
}

type StateRepository struct {
	CreatedAt Timestamp `yaml:"created_at"`
	UpdatedAt Timestamp `yaml:"updated_at"`
	Dir       string    `yaml:"dir"`
	Name      string    `yaml:"name"`
	Alias     string    `yaml:"alias"`
	Branch    string    `yaml:"branch"`
	Primary   string    `yaml:"primary"`
}

type StateDefaults struct {
	Branches *StateDefaultsBranches `yaml:"branches"`
}

type StateDefaultsBranches struct {
	Primary string `yaml:"primary"`
}

type StateHistory struct {
	Max     int      `yaml:"max"`
	Entries []string `yaml:"entries"`
}

func (this *State) Save(filepath string) {
	bytes := assert.Must(yaml.Marshal(this))
	assert.Throws(os.WriteFile(filepath, bytes, 0700))
}

func (this *State) Load(filepath string) (err error) {
	_, err = os.Stat(filepath)

	if os.IsNotExist(err) {
		return err
	}

	err = nil

	bytes := assert.Must(os.ReadFile(filepath))
	yaml.Unmarshal(bytes, this)

	return
}
