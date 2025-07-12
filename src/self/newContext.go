package self

import (
	"os"
	"path"
	"strconv"

	"github.com/jgttech/repo/src/assert"
	"github.com/jgttech/repo/src/cfg"
	"github.com/jgttech/repo/src/env"
	"github.com/jgttech/repo/src/fs"

	yml "gopkg.in/yaml.v3"
)

func newContext() (ctx *cfg.Context) {
	ctx = &cfg.Context{}
	cwd := assert.Must(os.Getwd())
	repo_mode := os.Getenv("REPO_MODE")

	if repo_mode == "" {
		repo_mode = strconv.Itoa(int(env.REPO_MODE_PROD))
	}

	mode := assert.Must(strconv.Atoi(repo_mode))
	filename := env.REPO_FILE
	dirname := env.REPO_DIR

	prod_mode := int(env.REPO_MODE_PROD)
	prod_path := path.Join(os.Getenv("HOME"), dirname)
	dev_path := path.Join(assert.Must(os.Getwd()))

	ctx.Home = fs.NewFolder(prod_path)

	switch mode {
	case prod_mode:
		ctx.Home = fs.NewFolder(prod_path)
	default:
		ctx.Home = fs.NewFolder(dev_path)
	}

	switch mode {
	case prod_mode:
		ctx.Conf = fs.NewFile(path.Join(prod_path, filename))
	default:
		ctx.Conf = fs.NewFile(path.Join(dev_path, filename))
	}

	ctx.Pwd = fs.NewFolder(cwd)

	yaml := &cfg.Yaml{}
	yml.Unmarshal(assert.Must(os.ReadFile(ctx.Conf.Path)), yaml)
	bin := fs.NewFile(yaml.Bin)
	state := fs.NewFile(yaml.State)

	switch mode {
	case prod_mode:
		ctx.Bin = bin
		ctx.State = state
	default:
		ctx.Bin = fs.NewFile(path.Join(cwd, bin.Name))
		ctx.State = fs.NewFile(path.Join(cwd, state.Name))
	}

	ctx.Version = yaml.Version

	return
}
