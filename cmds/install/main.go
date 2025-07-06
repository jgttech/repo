package install

import (
	"context"

	"github.com/jgttech/repo/src/fs"
	"github.com/jgttech/repo/src/self"
	"github.com/urfave/cli/v3"
)

var Command = &cli.Command{
	Name: "install",
	Action: func(ctx context.Context, c *cli.Command) (err error) {
		yaml := self.ExecutionContext.Yaml
		config := yaml.Config
		bin := yaml.Bin
		configDir := fs.NewFolder(config.Base)
		binDir := fs.NewFolder(bin.Base)

		if !configDir.Exists() {
			configDir.Create()
		}

		if !binDir.Exists() {
			binDir.Create()
		}

		if !config.Exists() {
			config.Create()
		}

		return
	},
}
