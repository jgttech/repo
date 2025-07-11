package cfg

import (
	"crypto/sha512"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jgttech/repo/src/assert"
	"gopkg.in/yaml.v3"
)

type Timestamp int64

type State struct {
	Path         string                      `yaml:"-"`
	CreatedAt    Timestamp                   `yaml:"created_at"`
	UpdatedAt    Timestamp                   `yaml:"updated_at"`
	Backups      []*StateBackup              `yaml:"backups"`
	Repositories map[string]*StateRepository `yaml:"repositories"`
	Defaults     *StateDefaults              `yaml:"defaults"`
	History      *StateHistory               `yaml:"history"`
}

type StateBackup struct {
	From string `yaml:"from"`
	To   string `yaml:"to"`
}

type StateRepository struct {
	CreatedAt Timestamp `yaml:"created_at"`
	UpdatedAt Timestamp `yaml:"updated_at"`
	Dir       string    `yaml:"dir"`
	Name      string    `yaml:"name"`
	Alias     string    `yaml:"alias"`
	Primary   string    `yaml:"primary"`
	Branch    string    `yaml:"branch"`
	Ssh       string    `yaml:"ssh"`
	Config    string    `yaml:"config"`
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

func (this *State) update(timestamp Timestamp, id string) {
	this.UpdatedAt = timestamp

	if this.History.Max > 0 && id != "" {
		size := len(this.History.Entries)

		if size >= this.History.Max {
			history := []string{}

			for _, entry := range this.History.Entries[1:] {
				history = append(history, entry)
			}

			this.History.Entries = history
		}

		entry := fmt.Sprintf("%s@%s", strconv.Itoa(int(timestamp)), id)
		this.History.Entries = append(this.History.Entries, entry)
	}
}

func (this *State) Save() {
	this.update(Timestamp(time.Now().Unix()), "")

	filepath := this.Path
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

func (this *State) AddRepository(repo *StateRepository) {
	hasher := sha512.New()
	signature := fmt.Sprintf("%x", hasher.Sum([]byte(repo.Dir)))[0:32]

	this.Repositories[signature] = repo
	this.update(repo.UpdatedAt, signature)
}
