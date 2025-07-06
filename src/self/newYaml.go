package self

import (
	"os"
	"path"

	"github.com/jgttech/repo/src/assert"
	"github.com/jgttech/repo/src/fs"
	"gopkg.in/yaml.v3"
)

type ymlFile struct {
	Path string `yaml:"path"`
	Name string `yaml:"name"`
}

type ymlConf struct {
	Version string  `yaml:"version"`
	Config  ymlFile `yaml:"config"`
	Bin     ymlFile `yaml:"bin"`
}

type yml struct {
	Version string
	Config  *fs.Node
	Bin     *fs.Node
}

func newYamlConf(file string) (cfg *yml) {
	cfg = &yml{}

	conf := ymlConf{}
	yaml.Unmarshal(assert.Must(os.ReadFile(file)), &conf)

	configFile := path.Join(conf.Config.Path, conf.Config.Name)
	binFile := path.Join(conf.Bin.Path, conf.Bin.Name)

	cfg.Version = conf.Version
	cfg.Config = fs.NewNode(configFile)
	cfg.Bin = fs.NewNode(binFile)

	return
}
