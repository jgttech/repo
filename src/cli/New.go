package cli

import (
	"os"

	"github.com/jgttech/repo/src/fs"
	"gopkg.in/yaml.v3"
)

func New(file string) (*Cli, error) {
	cli := &Cli{}
	_, err := os.Stat(file)

	if err != nil {
		return nil, err
	}

	bytes, err := os.ReadFile(file)
	err = yaml.Unmarshal(bytes, cli)
	cli.node = fs.NewNode(file)

	return cli, err
}
