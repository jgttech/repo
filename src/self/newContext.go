package self

import (
	"os"
	"path"
	"strconv"

	"github.com/jgttech/repo/src/assert"
	"github.com/jgttech/repo/src/env"
)

type context struct {
	Yaml *yml
	Pwd  string
	Home string
	Conf string
}

func newContext() (ctx *context) {
	file := env.REPO_FILE
	dir := env.REPO_DIR
	repo_mode := os.Getenv("REPO_MODE")

	if repo_mode == "" {
		repo_mode = strconv.Itoa(int(env.REPO_MODE_PROD))
	}

	mode := assert.Must(strconv.Atoi(repo_mode))
	prod := path.Join(os.Getenv("HOME"), dir)
	dev := path.Join(assert.Must(os.Getwd()))
	PROD := int(env.REPO_MODE_PROD)

	ctx = &context{}
	ctx.Pwd = assert.Must(os.Getwd())

	switch mode {
	case PROD:
		ctx.Home = prod
	default:
		ctx.Home = dev
	}

	switch mode {
	case PROD:
		ctx.Conf = path.Join(prod, file)
	default:
		ctx.Conf = path.Join(dev, file)
	}

	ctx.Yaml = newYaml(ctx.Conf)

	return
}
