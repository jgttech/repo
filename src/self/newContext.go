package self

import (
	"github.com/jgttech/repo/src/assert"
	"github.com/jgttech/repo/src/cli"
)

func newContext() (c *ctx) {
	c = &ctx{
		Cli: assert.Must(cli.NewCli("")),
	}

	return
}
