package cli

import "github.com/jgttech/repo/src/env"

func New() (*Cli, error) {
	file := env.BASE_CLI
	return Load(file)
}
