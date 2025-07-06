package self

import (
	"os"

	"github.com/jgttech/repo/src/assert"
	"gopkg.in/yaml.v3"
)

type yml struct {
	Version string `yaml:"version"`
}

func newYaml(file string) (cfg *yml) {
	cfg = &yml{}
	yaml.Unmarshal(assert.Must(os.ReadFile(file)), cfg)
	return
}
